// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameConduitVersionPgsql = "_conduit_version_pgsql"

// ConduitVersionPgsql mapped from table <_conduit_version_pgsql>
type ConduitVersionPgsql struct {
	Versionnumber int32     `gorm:"column:versionnumber;not null" json:"versionnumber"`
	Dateofupgrade time.Time `gorm:"column:dateofupgrade;not null" json:"dateofupgrade"`
}

// TableName ConduitVersionPgsql's table name
func (*ConduitVersionPgsql) TableName() string {
	return TableNameConduitVersionPgsql
}