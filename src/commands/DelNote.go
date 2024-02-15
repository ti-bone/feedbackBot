/*
 * DelNote.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"feedbackBot/src/notes"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func DelNote(b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()

	if len(args) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, messages.NoteIdInvalid.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	noteId, err := helpers.ParseNoteId(args[1])

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	err = notes.DeleteNoteById(noteId)

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf(messages.NoteDeleted, noteId), &gotgbot.SendMessageOpts{})

	return err
}
