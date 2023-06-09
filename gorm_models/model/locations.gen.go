// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLocation = "locations"

// Location mapped from table <locations>
type Location struct {
	Longitude        float64   `gorm:"column:longitude;not null" json:"longitude"`
	Latitude         float64   `gorm:"column:latitude;not null" json:"latitude"`
	City             string    `gorm:"column:city;not null" json:"city"`
	Address          string    `gorm:"column:address;not null" json:"address"`
	Postalcode       string    `gorm:"column:postalcode;not null" json:"postalcode"`
	Name             string    `gorm:"column:name;not null" json:"name"`
	ID               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Shortdescription string    `gorm:"column:shortdescription;not null" json:"shortdescription"`
	Text             string    `gorm:"column:text" json:"text"`
	Createdat        time.Time `gorm:"column:createdat;not null;default:now()" json:"createdat"`
	Updatedat        time.Time `gorm:"column:updatedat;not null;default:now()" json:"updatedat"`
	ImageID          int64     `gorm:"column:image_id;not null" json:"image_id"`
}

// TableName Location's table name
func (*Location) TableName() string {
	return TableNameLocation
}
