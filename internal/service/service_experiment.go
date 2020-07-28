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
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/outman/abucket/internal/form"
	"github.com/outman/abucket/internal/model"
	"github.com/outman/abucket/internal/pkg"
)

type serviceExperiment struct{}

func NewExperimentService() *serviceExperiment {
	return &serviceExperiment{}
}

// Index 根据条件获取数据
func (s *serviceExperiment) Index(f *form.FormSearchExperiment) []model.Experiment {
	var experiments []model.Experiment
	if f.CurrentStatus > 0 {
		pkg.NewMySQL().DB.Where("current_status = ?", f.CurrentStatus).Order("end_time desc").Find(&experiments)
	} else {
		pkg.NewMySQL().DB.Order("end_time desc").Find(&experiments)
	}
	return experiments
}

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
