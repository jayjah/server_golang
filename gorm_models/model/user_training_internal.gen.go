// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserTrainingInternal = "user_training_internal"

// UserTrainingInternal mapped from table <user_training_internal>
type UserTrainingInternal struct {
	ID         int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Counter    int32 `gorm:"column:counter;not null;default:50" json:"counter"`
	UserID     int64 `gorm:"column:user_id;not null" json:"user_id"`
	TrainingID int64 `gorm:"column:training_id;not null" json:"training_id"`
}

// TableName UserTrainingInternal's table name
func (*UserTrainingInternal) TableName() string {
	return TableNameUserTrainingInternal
}