/*
 * ParseInputNote.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/constants"
)

func ParseNoteId(input string) (uint, error) {
	if noteId, err := ParseUint(input); err == nil {
		return noteId, nil
	}

	if len(input) > 2 && input[0] == '#' && input[1] == 'n' {
		noteId, err := ParseUint(input[2:])

		if err != nil {
			return 0, constants.NoteIdInvalid
		}

		return noteId, nil
	}

	return 0, constants.NoteIdInvalid
}
