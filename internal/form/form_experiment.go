package form

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

import "time"

type FormSearchExperiment struct {
	CurrentStatus uint `form:"status"`
}

type FormCreateExperiment struct {
	Name      string    `form:"name" binding:"required"`
	UniqKey   string    `form:"uniq_key" binding:"required"`
	LayerID   uint      `form:"layer_id" binding:"required"`
	LayerUsed uint      `form:"layer_used" binding:"required"`
	BeginTime time.Time `form:"begin_time" binding:"required" time_format:"2006-01-02"`
	EndTime   time.Time `form:"end_time" binding:"required" time_format:"2006-01-02"`
}

type FormUpdateExperiment struct {
	ID            uint      `form:"id" binding:"required"`
	CurrentStatus uint      `form:"status" binding:"required"`
	BeginTime     time.Time `form:"begin_time" binding:"required" time_format:"2006-01-02"`
	EndTime       time.Time `form:"end_time" binding:"required" time_format:"2006-01-02"`
}

type FormDeleteExperiment struct {
	ID            uint `form:"id" binding:"required"`
	CurrentStatus uint `form:"status" binding:"required"`
}

type Group struct {
	Name    string `json:"name" binding:"required"`
	Percent *uint  `json:"percent" binding:"required"`
}

type FormExperimentGroups struct {
	ID     uint    `json:"id" binding:"required"`
	Groups []Group `json:"groups" binding:"required"`
}
