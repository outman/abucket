package api

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
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outman/abucket/internal/form"
	"github.com/outman/abucket/internal/service"
)

// actionExperiment struct
type actionExperiment struct {
}

// NewActionExperiment return experiment action for route
func NewActionExperiment() *actionExperiment {
	return &actionExperiment{}
}

// Create experiment
// binding form.FormCreateExperiment
func (a *actionExperiment) Create(c *gin.Context) {
	var formParams form.FormCreateExperiment
	if err := c.ShouldBind(&formParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}

	s := service.NewExperimentService()
	code, data := s.Create(&formParams)
	if code == service.ServiceOptionSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code": actionSuccess,
			"data": data,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": actionParameterError,
		"text": service.CodeMessage(code),
	})
	return
}

// Index all experiments
// binding form.FormSearchExperiment
func (a *actionExperiment) Index(c *gin.Context) {
	var query form.FormSearchExperiment
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}
	s := service.NewExperimentService()
	data := s.Index(&query)
	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": data,
	})
	return
}

// Update
// Update experiment status, begin_time and end_time
func (a *actionExperiment) Update(c *gin.Context) {
	var formParams form.FormUpdateExperiment
	if err := c.ShouldBind(&formParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}
	s := service.NewExperimentService()
	code, data := s.Update(&formParams)
	if code != service.ServiceOptionSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code": actionServerError,
			"text": service.CodeMessage(code),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": data,
	})
	return
}

func (a *actionExperiment) Delete(c *gin.Context) {
	var formParams form.FormDeleteExperiment
	if err := c.ShouldBind(&formParams); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}

	s := service.NewExperimentService()
	code, data := s.Delete(&formParams)
	if code != service.ServiceOptionSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code": actionServerError,
			"text": service.CodeMessage(code),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": data,
	})
	return
}

func (a *actionExperiment) CreateGroup(c *gin.Context) {
	var group form.FormExperimentGroups
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": err.Error(),
		})
		return
	}

	var total uint
	names := make(map[string]uint)
	for _, v := range group.Groups {
		total += *v.Percent
		names[v.Name] = *v.Percent
	}

	if total > 100 || len(group.Groups) == 0 || len(names) < len(group.Groups) {
		c.JSON(http.StatusOK, gin.H{
			"code": actionParameterError,
			"text": service.CodeMessage(service.ServiceOptionGroupError),
		})
		return
	}

	s := service.NewExperimentService()
	code, data := s.UpdateGroup(&group)

	if code != service.ServiceOptionSuccess {
		c.JSON(http.StatusOK, gin.H{
			"code": actionServerError,
			"text": service.CodeMessage(code),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": actionSuccess,
		"data": data,
	})
	return
}
