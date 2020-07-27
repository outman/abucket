package model

import "time"

// Layer struct
type Layer struct {
	ID        uint      `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name;not null;unique:uniq_name;size:200"`
	Left      uint      `gorm:"column:left;not null;default:0"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName db table name
func (e *Layer) TableName() string {
	return "abucket_layer"
}
