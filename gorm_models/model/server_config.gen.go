// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameServerConfig = "server_config"

// ServerConfig mapped from table <server_config>
type ServerConfig struct {
	_ID int64  `gorm:"column:_id;primaryKey;autoIncrement:true" json:"_id"`
	ID string `gorm:"column:id;not null" json:"id"`
}

// TableName ServerConfig's table name
func (*ServerConfig) TableName() string {
	return TableNameServerConfig
}
