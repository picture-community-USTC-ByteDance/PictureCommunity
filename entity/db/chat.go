package db

import "time"

//每条私聊消息
type ChatMessage struct {
	MessageId uint      `gorm:"primarykey"`
	FromId    uint      `gorm:"not null;"` // 发送者
	ToId      uint      `gorm:"not null;"` // 接受者
	Content   string    `gorm:"size:255;"` //聊天内容
	CreatedAt time.Time `gorm:"not null;"` //内容发送时间
	IsRead    int8      `gorm:"default:0"` //1为已读状态
}
