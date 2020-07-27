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
package model

import (
	"time"

	"github.com/outman/abucket/internal/pkg"
)

// Experiment struct
type Experiment struct {
	ID            uint      `gorm:"Column:id;PRIMARY_KEY"`
	Name          string    `gorm:"Column:name;UNIQUE_INDEX:uniq_name;NOT NULL;Size:200"`
	LayerID       uint      `gorm:"Column:layer_id;NOT NULL;default:0;Index:idx_layer_id"`
	LayerUsed     uint      `gorm:"Column:layer_used;NOT NULL;default:0"`
	Groups        string    `gorm:"Column:groups;Type:longtext"`
	CurrentStatus uint      `gorm:"Column:current_status;DEFAULT:0"`
	BeginTime     time.Time `gorm:"Column:begin_time"`
	EndTime       time.Time `gorm:"Column:end_time"`
	CreatedAt     time.Time `gorm:"Column:created_at"`
	UpdatedAt     time.Time `gorm:"Column:updated_at"`
}

// TableName db table name
func (e *Experiment) TableName() string {
	return "abucket_experiment"
}

func (e *Experiment) Index() []Experiment {
	var experiments []Experiment
	pkg.NewMySQL().DB.Find(&experiments)
	return experiments
}
