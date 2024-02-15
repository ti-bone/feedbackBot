/*
 * GetUserByUsername.go
 * Copyright (c) ti-bone 2023-2024
 */

package users

import "feedbackBot/src/models"

// GetUserByUsername - resolves a user by their username
func GetUserByUsername(username string) (*models.User, error) {
	return getUser("lower(username) = lower(?)", username)
}
