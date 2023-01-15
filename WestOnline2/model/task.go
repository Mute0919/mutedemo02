package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	status    int    `gorm:"default:'0'"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
	Status    int
}
