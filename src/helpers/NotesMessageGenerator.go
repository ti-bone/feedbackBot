/*
 * NotesMessageGenerator.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/constants"
	"feedbackBot/src/models"
	"fmt"
	"html"
)

func GenerateNotesMessage(notes []*models.Note) string {
	if len(notes) == 0 {
		return constants.InternalError.Error()
	}

	text := fmt.Sprintf("Notes for #u%d:", notes[0].UserID)
	for _, note := range notes {
		text += fmt.Sprintf(
			"\n#n%d(by #u%d): <code>%s</code>",
			note.ID, note.AddedByID,
			html.EscapeString(note.Text),
		)
	}

	return text
}
