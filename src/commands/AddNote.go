/*
 * AddNote.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"feedbackBot/src/notes"
	"feedbackBot/src/users"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"strings"
)

func AddNote(b *gotgbot.Bot, ctx *ext.Context) error {
	// Check if the user is provided
	user, isTopicMessage, err := helpers.ResolveUserWithSource(ctx, b)

	if err != nil || user == nil {
		return err
	}

	var noteText string

	switch {
	case !isTopicMessage && len(ctx.Args()) > 2:
		noteText = strings.Join(ctx.Args()[2:], " ")
	case isTopicMessage && len(ctx.Args()) > 1:
		noteText = strings.Join(ctx.Args()[1:], " ")
	}

	if len(noteText) == 0 {
		_, err := ctx.EffectiveMessage.Reply(b, messages.NoteTextNotSpecified.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	addedBy, err := users.GetUserById(ctx.EffectiveMessage.From.Id)

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	noteId, err := notes.AddNote(user, noteText, addedBy)

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(
		b,
		fmt.Sprintf(messages.NoteAdded, noteId, user.UserID),
		&gotgbot.SendMessageOpts{},
	)

	return err
}
