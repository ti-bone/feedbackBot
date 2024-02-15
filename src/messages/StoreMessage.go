/*
 * StoreMessage.go
 * Copyright (c) ti-bone 2023-2024
 */

package messages

import (
	"feedbackBot/src/db"
	"feedbackBot/src/models"
)

// StoreMessage - stores a message in the database
// This function stores a message in the database, it is used to store messages identifiers sent by the user and the support
// More info: http://youtrack.hub/issue/GL-11 (only for employees, internal resource)
func StoreMessage(message models.Message) error {
	err := db.Connection.Create(&message).Error

	return err
}
