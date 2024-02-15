/*
 * Note.go
 * Copyright (c) ti-bone 2023-2024
 */

package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID    int64
	AddedByID int64
	Text      string
}
