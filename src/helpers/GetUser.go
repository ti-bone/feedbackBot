/*
 * GetUser.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"errors"
	"feedbackBot/src/db"
	"feedbackBot/src/messages"
	"feedbackBot/src/models"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetUserByUsername(username string) (*models.User, error) {
	return getUser("lower(username) = lower(?)", username)
}

func GetUserById(userId int64) (*models.User, error) {
	return getUser("user_id = ?", userId)
}

func GetUserByTopicId(topicId int64) (*models.User, error) {
	return getUser("topic_id = ?", topicId)
}

func getUser(query interface{}, value ...interface{}) (*models.User, error) {
	var user models.User

	res := db.Connection.Where(query, value).First(&user)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, messages.UserNotFound
	} else if res.Error != nil {
		log.SetOutput(os.Stderr)
		log.Println(res.Error)

		return nil, messages.InternalError
	}

	if user.UserID <= 0 {
		return nil, messages.UserNotFound
	}

	return &user, nil
}
