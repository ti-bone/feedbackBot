/*
 * User.go
 * Copyright (c) ti-bone 2023
 */

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID       int64 `gorm:"unique"`
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	TopicID      int64
	IsBanned     bool
	IsAdmin      bool
}
