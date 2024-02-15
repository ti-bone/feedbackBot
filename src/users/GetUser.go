/*
 * GetUser.go
 * Copyright (c) ti-bone 2023-2024
 */

package users

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/db"
	"feedbackBot/src/models"
	"gorm.io/gorm"
	"log"
	"os"
)

// getUser - resolves a user by a query, used internally in package
func getUser(query interface{}, value ...interface{}) (*models.User, error) {
	var user models.User

	res := db.Connection.Where(query, value...).First(&user)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, constants.UserNotFound
	} else if res.Error != nil {
		log.SetOutput(os.Stderr)
		log.Println(res.Error)

		return nil, constants.InternalError
	}

	if user.UserID <= 0 {
		return nil, constants.UserNotFound
	}

	return &user, nil
}
