/*
 * UserActions.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/db"
	"feedbackBot/src/messages"
	"feedbackBot/src/models"
)

func BanUser(user *models.User) error {
	// Get user from the DB
	db.Connection.Where("user_id = ?", user.UserID).First(&user)

	// Check if user is already banned
	if user.IsBanned {
		return messages.UserAlreadyBanned
	}

	db.Connection.Model(&user).Update("is_banned", true)

	return nil
}

func UnbanUser(user *models.User) error {
	// Get user from the DB
	db.Connection.Where("user_id = ?", user.UserID).First(&user)

	// Check if user is not banned
	if !user.IsBanned {
		return messages.UserNotBanned
	}

	db.Connection.Model(&user).Update("is_banned", false)

	return nil
}
