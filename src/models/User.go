/*
 * User.go
 * Copyright (c) ti-bone 2023-2024
 */

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId       int64 `gorm:"unique"`
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	TopicId      int64
	IsBanned     bool
	IsAdmin      bool
	IsProtected  bool
}
