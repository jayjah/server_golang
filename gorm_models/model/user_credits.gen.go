// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserCredit = "user_credits"

// UserCredit mapped from table <user_credits>
type UserCredit struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Createdat time.Time `gorm:"column:createdat;not null;default:now()" json:"createdat"`
	Updatedat time.Time `gorm:"column:updatedat;not null;default:now()" json:"updatedat"`
	UserID    int64     `gorm:"column:user_id;not null" json:"user_id"`
	CreditID  int64     `gorm:"column:credit_id;not null" json:"credit_id"`
}

// TableName UserCredit's table name
func (*UserCredit) TableName() string {
	return TableNameUserCredit
}
