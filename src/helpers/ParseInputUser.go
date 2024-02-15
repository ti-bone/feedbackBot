/*
 * ParseInputUser.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/constants"
	"feedbackBot/src/models"
	users2 "feedbackBot/src/users"
	"strconv"
)

func ParseInputUser(input string) (*models.User, error) {
	if userId, err := strconv.ParseInt(input, 10, 64); err == nil {
		user, err := users2.GetUserById(userId)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	if len(input) < 2 {
		return nil, constants.UserInvalid
	}

	if input[0] == '#' && input[1] == 'u' {
		userId, err := strconv.ParseInt(input[2:], 10, 64)

		if err != nil {
			return nil, constants.UserIdInvalid
		}

		user, err := users2.GetUserById(userId)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	if input[0] == '@' {
		username := input[1:]

		user, err := users2.GetUserByUsername(username)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, constants.UserInvalid
}
