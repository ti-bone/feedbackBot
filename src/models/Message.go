/*
 * Message.go
 * Copyright (c) ti-bone 2023-2024
 */

package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID           int64
	UserMessageId    int64
	SupportMessageId int64
	SupportChatId    int64
	IsOutgoing       bool
}
