/*
 * UserNotes.go
 * Copyright (c) ti-bone 2023-2024
 */

package notes

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/db"
	"feedbackBot/src/models"
	"gorm.io/gorm"
	"log"
	"os"
)

func AddNote(user *models.User, text string, addedBy *models.User) (uint, error) {
	if user == nil || addedBy == nil {
		return 0, constants.InternalError
	}

	// Check note length
	if len(text) > 1024 {
		return 0, constants.NoteTooLong
	}

	// Add note
	note := models.Note{
		UserID:    user.UserID,
		AddedByID: addedBy.UserID,
		Text:      text,
	}

	res := db.Connection.Create(&note)

	if res.Error != nil {
		log.SetOutput(os.Stderr)
		log.Println(res.Error)

		return 0, constants.InternalError
	}

	return note.ID, nil
}

func GetNotes(user *models.User) ([]*models.Note, error) {
	var notes []*models.Note

	res := db.Connection.Where("user_id = ?", user.UserID).Find(&notes)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, constants.NotesNotFound
	}

	if res.Error != nil {
		log.SetOutput(os.Stderr)
		log.Println(res.Error)

		return nil, constants.InternalError
	}

	return notes, nil
}

func DeleteNoteById(noteId uint) error {
	res := db.Connection.Where("id = ?", noteId).Delete(&models.Note{})

	if res.RowsAffected == 0 {
		return constants.NoteNotFound
	}

	if res.Error != nil {
		log.SetOutput(os.Stderr)
		log.Println(res.Error)

		return constants.InternalError
	}

	return nil
}
