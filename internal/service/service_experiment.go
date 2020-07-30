package service

/*
Copyright © 2020 pochonlee@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/minio/highwayhash"
	"github.com/outman/abucket/internal/form"
	"github.com/outman/abucket/internal/model"
	"github.com/outman/abucket/internal/pkg"
)

// Cache with memory
var cacheExperiment = &pkg.CacheStore{
	Item:     sync.Map{},
	Duration: 600,
}

var cacheLayer = &pkg.CacheStore{
	Item:     sync.Map{},
	Duration: 600,
}

type serviceExperiment struct{}

func NewExperimentService() *serviceExperiment {
	return &serviceExperiment{}
}

// Index
// search with status and order by end_time desc
func (s *serviceExperiment) Index(f *form.FormSearchExperiment) []model.Experiment {
	var experiments []model.Experiment
	if f.CurrentStatus > 0 {
		pkg.NewMySQL().DB.Where("current_status = ?", f.CurrentStatus).Order("end_time desc").Find(&experiments)
	} else {
		pkg.NewMySQL().DB.Order("end_time desc").Find(&experiments)
	}
	return experiments
}

// Create new experiment record.
func (s *serviceExperiment) Create(f *form.FormCreateExperiment) (int, model.Experiment) {

	// check name uniq
	var experiment model.Experiment
	pkg.NewMySQL().DB.Where("name = ?", f.Name).First(&experiment)

	if experiment.ID > 0 {
		return ServiceOptionRecordNameDuplicate, experiment
	}

	// check experiment uniq key
	pkg.NewMySQL().DB.Where("uniq_key = ?", f.UniqKey).First(&experiment)

	if experiment.ID > 0 {
		return ServiceOptionRecordKeyDuplicate, experiment
	}

	// check layer_id
	var layer model.Layer
	pkg.NewMySQL().DB.Where("id = ?", f.LayerID).First(&layer)
	if layer.ID == 0 || layer.Left < f.LayerUsed {
		return ServiceOptionLayerNoSpace, experiment
	}

	experiment = model.Experiment{
		Name:      f.Name,
		UniqKey:   f.UniqKey,
		LayerID:   f.LayerID,
		LayerUsed: f.LayerUsed,
		BeginTime: f.BeginTime,
		EndTime:   f.EndTime,
		CreatedAt: time.Now(),
	}

	// Create experiment and reduce layer left
	result := pkg.NewMySQL().DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&experiment).Error; err != nil {
			return err
		}

		if (layer.Left - experiment.LayerUsed) < 0 {
			return fmt.Errorf("Layer left %d, need %d", layer.Left, experiment.LayerUsed)
		}

		layer.Left = layer.Left - experiment.LayerUsed
		layer.UpdatedAt = time.Now()
		if err := tx.Save(&layer).Error; err != nil {
			return err
		}
		return nil
	})

	if result != nil {
		return ServiceOptionDbError, experiment
	}

	return ServiceOptionSuccess, experiment
}

// Update status begin_time and end_time
func (s *serviceExperiment) Update(f *form.FormUpdateExperiment) (int, model.Experiment) {
	var experiment model.Experiment
	pkg.NewMySQL().DB.Where("id = ?", f.ID).First(&experiment)

	if experiment.ID == 0 {
		return ServiceOptionRecordNotFound, experiment
	}

	experiment.BeginTime = f.BeginTime
	experiment.EndTime = f.EndTime
	experiment.CurrentStatus = f.CurrentStatus
	experiment.UpdatedAt = time.Now()

	if err := pkg.NewMySQL().DB.Save(&experiment).Error; err != nil {
		return ServiceOptionSQLError, experiment
	}

	return ServiceOptionSuccess, experiment
}

// Delete
// set status = delete and update_at = time.Now()
func (s *serviceExperiment) Delete(f *form.FormDeleteExperiment) (int, model.Experiment) {
	var experiment model.Experiment
	pkg.NewMySQL().DB.Where("id = ?", f.ID).First(&experiment)

	if experiment.ID == 0 {
		return ServiceOptionRecordNotFound, experiment
	}

	experiment.CurrentStatus = f.CurrentStatus
	experiment.UpdatedAt = time.Now()

	if err := pkg.NewMySQL().DB.Save(&experiment).Error; err != nil {
		return ServiceOptionSQLError, experiment
	}

	return ServiceOptionSuccess, experiment
}

func (s *serviceExperiment) UpdateGroup(f *form.FormExperimentGroups) (int, model.Experiment) {
	var experiment model.Experiment
	pkg.NewMySQL().DB.Where("id = ?", f.ID).First(&experiment)
	if experiment.ID == 0 {
		return ServiceOptionRecordNotFound, experiment
	}

	data, err := json.Marshal(f)
	if err != nil {
		return ServiceOptionParameterError, experiment
	}

	experiment.Groups = string(data)
	experiment.UpdatedAt = time.Now()

	if err := pkg.NewMySQL().DB.Save(&experiment).Error; err != nil {
		return ServiceOptionSQLError, experiment
	}

	return ServiceOptionSuccess, experiment
}

func (s *serviceExperiment) HitGroup(f *form.RequestFormGroup) (int, string, uint) {

	var experiment model.Experiment
	v, ok := cacheExperiment.Get(f.UniqKey)
	if ok {
		experiment = v.(model.Experiment)
	} else {
		if err := pkg.NewMySQL().DB.Where("uniq_key = ?", f.UniqKey).First(&experiment).Error; err != nil {
			return ServiceOptionRecordNotFound, "", 0
		}
		cacheExperiment.Set(f.UniqKey, experiment)
	}

	// Experiment not working
	if experiment.CurrentStatus != model.StWorking {
		return ServiceOptionSuccess, "", 0
	}

	// Experiment out of date range
	current := time.Now()
	if current.Before(experiment.BeginTime) || current.After(experiment.EndTime) {
		return ServiceOptionSuccess, "", 0
	}
	uuid := []byte(f.UUID)
	// Experiment layer 100 percent layer
	if experiment.LayerUsed == 100 {
		return selectBucket(uuid, experiment)
	}

	// Experiment layer less than 100
	// process layer hits
	var experiments []model.Experiment
	a, ok := cacheLayer.Get(experiment.LayerID)
	if ok {
		experiments = a.([]model.Experiment)
	} else {
		if err := pkg.NewMySQL().DB.Order("id asc").Where("layer_id = ?", experiment.LayerID).Find(&experiments).Error; err != nil {
			return ServiceOptionSQLError, "", 0
		}
		cacheLayer.Set(experiment.LayerID, experiments)
	}

	hitLayer, bucket := selectLayer(uuid, experiment, experiments)
	if !hitLayer {
		return ServiceOptionSuccess, "", bucket
	}
	return selectBucket(uuid, experiment)
}

func selectLayer(data []byte, e model.Experiment, es []model.Experiment) (bool, uint) {
	key := []byte("14a1c4b4595739b653ffe832bcd50928")
	num := highwayhash.Sum64(data, key)
	bucket := uint(num%100 + 1)
	var total uint
	for _, v := range es {
		total += v.LayerUsed
		if bucket <= total {
			if v.UniqKey == e.UniqKey {
				return true, bucket
			}
			return false, bucket
		}
	}
	return false, bucket
}

func selectBucket(data []byte, e model.Experiment) (int, string, uint) {
	var f form.FormExperimentGroups
	err := json.Unmarshal([]byte(e.Groups), &f)
	if err != nil {
		return ServiceOptionGroupJSONError, "", 0
	}

	key := []byte("cff46a76d99c63c8a405eb539b781cab")
	num := highwayhash.Sum64(data, key)
	bucket := uint(num%100 + 1)

	var total uint
	for _, v := range f.Groups {
		total += *v.Percent
		if bucket <= total {
			return ServiceOptionSuccess, v.Name, bucket
		}
	}
	return ServiceOptionSuccess, "", bucket
}
