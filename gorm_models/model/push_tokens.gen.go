// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePushToken = "push_tokens"

// PushToken mapped from table <push_tokens>
type PushToken struct {
	ID    int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Token string `gorm:"column:token;not null" json:"token"`
	Hms   bool   `gorm:"column:hms;not null" json:"hms"`
	Apn   bool   `gorm:"column:apn;not null" json:"apn"`
	Fcm   bool   `gorm:"column:fcm;not null;default:true" json:"fcm"`
}

// TableName PushToken's table name
func (*PushToken) TableName() string {
	return TableNamePushToken
}
