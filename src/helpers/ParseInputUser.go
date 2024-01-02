/*
 * ParseInputUser.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/messages"
	"feedbackBot/src/models"
	"strconv"
)

func ParseInputUser(input string) (*models.User, error) {
	if userId, err := strconv.ParseInt(input, 10, 64); err == nil {
		user, err := GetUserById(userId)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	if input[0] == '#' && input[1] == 'u' {
		userId, err := strconv.ParseInt(input[2:], 10, 64)

		if err != nil {
			return nil, messages.UserIdInvalid
		}

		user, err := GetUserById(userId)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	if input[0] == '@' {
		username := input[1:]

		user, err := GetUserByUsername(username)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, messages.UserInvalid
}
