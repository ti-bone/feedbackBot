/*
 * GetUserByTopicId.go
 * Copyright (c) ti-bone 2023-2024
 */

package users

import "feedbackBot/src/models"

// GetUserByTopicId - resolves a user by id of the topic they're currently assigned to
func GetUserByTopicId(topicId int64) (*models.User, error) {
	return getUser("topic_id = ?", topicId)
}
