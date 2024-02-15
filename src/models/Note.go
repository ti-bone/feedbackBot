/*
 * Note.go
 * Copyright (c) ti-bone 2023-2024
 */

package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID    int64
	User      User `gorm:"foreignKey:UserID"`
	AddedByID int64
	AddedBy   User `gorm:"foreignKey:UserID"`
	Text      string
}
