/*
 * GetMessageBySupportId.go
 * Copyright (c) ti-bone 2023-2024
 */

package messages

import "feedbackBot/src/models"

// GetMessageBySupportId - gets a message by its ID and chatID
func GetMessageBySupportId(messageId int64, chatId int64) (*models.Message, error) {
	return getMessage("support_message_id = ? and support_chat_id = ?", messageId, chatId)
}
