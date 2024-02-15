/*
 * GetMessage.go
 * Copyright (c) ti-bone 2023-2024
 */

package messages

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/db"
	"feedbackBot/src/models"
	"gorm.io/gorm"
	"log"
	"os"
)

// getMessage - gets info for message by a custom query, used internally in package
func getMessage(query interface{}, value ...interface{}) (*models.Message, error) {
	var message models.Message

	res := db.Connection.Where(query, value...).First(&message)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, constants.MessageNotFound
	} else if res.Error != nil {
		log.SetOutput(os.Stderr)
		log.Println(res.Error)

		return nil, constants.InternalError
	}

	if message.UserId <= 0 {
		return nil, constants.MessageNotFound
	}

	return &message, nil
}
