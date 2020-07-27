package model

import "time"

// Experiment struct
type Experiment struct {
	ID            uint      `gorm:"column:id;primary_key"`
	Name          string    `gorm:"column:name;unique:uniq_name;not null;size:200"`
	LayerID       uint      `gorm:"column:layer_id;not null;default:0"`
	LayerUsed     uint      `gorm:"column:layer_used;not null;default:0"`
	Groups        string    `gorm:"column:groups;type:longtext"`
	CurrentStatus uint      `gorm:"column:current_status;default:0"`
	BeginTime     time.Time `gorm:"column:begin_time"`
	EndTime       time.Time `gorm:"column:end_time"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

// TableName db table name
func (e *Experiment) TableName() string {
	return "abucket_experiment"
}
