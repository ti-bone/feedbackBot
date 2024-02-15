/*
 * GetMessageBySupportId.go
 * Copyright (c) ti-bone 2023-2024
 */

package messages

import "feedbackBot/src/models"

// GetMessageByUserId - gets a message by its user-side ID and chatID
func GetMessageByUserId(messageId int64, userId int64) (*models.Message, error) {
	return getMessage("user_message_id = ? and user_id = ?", messageId, userId)
}
