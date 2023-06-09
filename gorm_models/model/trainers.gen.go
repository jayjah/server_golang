// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTrainer = "trainers"

// Trainer mapped from table <trainers>
type Trainer struct {
	Agelabel         string    `gorm:"column:agelabel" json:"agelabel"`
	ID               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	Shortdescription string    `gorm:"column:shortdescription;not null" json:"shortdescription"`
	Text             string    `gorm:"column:text" json:"text"`
	Createdat        time.Time `gorm:"column:createdat;not null;default:now()" json:"createdat"`
	Updatedat        time.Time `gorm:"column:updatedat;not null;default:now()" json:"updatedat"`
	ImageID          int64     `gorm:"column:image_id;not null" json:"image_id"`
	UserID           int64     `gorm:"column:user_id" json:"user_id"`
}

// TableName Trainer's table name
func (*Trainer) TableName() string {
	return TableNameTrainer
}
