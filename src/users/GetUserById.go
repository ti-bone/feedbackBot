/*
 * GetUserById.go
 * Copyright (c) ti-bone 2023-2024
 */

package users

import "feedbackBot/src/models"

// GetUserById - resolves a user by their userId
func GetUserById(userId int64) (*models.User, error) {
	return getUser("user_id = ?", userId)
}
