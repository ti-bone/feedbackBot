/*
 * ParseInputUser.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/db"
	"feedbackBot/src/messages"
	"feedbackBot/src/models"
	"strconv"
)

func ParseInputUser(input string) (*models.User, error) {
	if userID, err := strconv.ParseInt(input, 10, 64); err == nil {
		user, err := getUserByID(userID)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	if input[0] == '#' && input[1] == 'u' {
		userID, err := strconv.ParseInt(input[2:], 10, 64)

		if err != nil {
			return nil, messages.UserIdInvalid
		}

		user, err := getUserByID(userID)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	if input[0] == '@' {
		username := input[1:]

		user, err := getUserByUsername(username)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, messages.UserInvalid
}

func getUserByID(userID int64) (*models.User, error) {
	var user *models.User

	db.Connection.Where("user_id = ?", userID).First(&user)

	if user == nil || user.UserID <= 0 {
		return nil, messages.UserNotFound
	}

	return user, nil
}

func getUserByUsername(username string) (*models.User, error) {
	var user *models.User

	db.Connection.Where("username = ?", username).First(&user)

	if user == nil || user.UserID <= 0 {
		return nil, messages.UserNotFound
	}

	return user, nil
}
