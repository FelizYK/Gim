package logic

import (
	"Gim/internal/server"

	"gorm.io/gorm"
)

type GroupInfo struct {
	gorm.Model
	Name  string
	Owner string
	Type  string
	Desc  string
}

func (table *GroupInfo) TableName() string {
	return "group_info"
}

func InitGroupTable() {
	server.DB.AutoMigrate(&GroupInfo{})
}
