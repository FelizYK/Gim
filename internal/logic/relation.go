package logic

import (
	"Gim/internal/server"

	"gorm.io/gorm"
)

type RelationInfo struct {
	gorm.Model
	Owner  string
	Target string
	Type   string
	Desc   string
}

func (table *RelationInfo) TableName() string {
	return "relation_info"
}

func InitRelationTable() {
	server.DB.AutoMigrate(&RelationInfo{})
}
