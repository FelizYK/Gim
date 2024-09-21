package logic

import (
	"Gim/internal/server"

	"gorm.io/gorm"
)

type MessageInfo struct {
	gorm.Model
	Sender   string
	Receiver string
	Type     string // private / group / broadcast
	Media    string // text / image / file
	Content  string
}

func (table *MessageInfo) TableName() string {
	return "message_info"
}

func InitMessageTable() {
	server.DB.AutoMigrate(&MessageInfo{})
}
