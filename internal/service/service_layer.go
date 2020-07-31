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
	"time"

	"github.com/outman/abucket/internal/form"
	"github.com/outman/abucket/internal/model"
	"github.com/outman/abucket/internal/pkg"
)

type serviceLayer struct{}

// NewLayerService *serviceLayer
func NewLayerService() *serviceLayer {
	return &serviceLayer{}
}

func (s *serviceLayer) Create(f *form.FormCreateLayer) (int, model.Layer) {
	var layer model.Layer
	pkg.NewMySQL().DB.Where("name = ?", f.Name).First(&layer)

	if layer.ID > 0 {
		return ServiceOptionRecordExists, layer
	}

	layer = model.Layer{
		Name:      f.Name,
		Left:      100,
		CreatedAt: time.Now(),
	}

	if err := pkg.NewMySQL().DB.Create(&layer).Error; err != nil {
		pkg.NewZapLogger().Error(err.Error())
		return ServiceOptionDbError, layer
	}

	return ServiceOptionSuccess, layer
}

func (s *serviceLayer) Index() []model.Layer {
	var layers []model.Layer
	pkg.NewMySQL().DB.Order("id asc").Where("`left` > ?", 0).Find(&layers)
	return layers
}
