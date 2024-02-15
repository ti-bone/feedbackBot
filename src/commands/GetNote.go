/*
 * GetNote.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/helpers"
	"feedbackBot/src/notes"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GetNotes(b *gotgbot.Bot, ctx *ext.Context) error {
	user, err := helpers.ResolveUser(ctx, b)

	if err != nil || user == nil {
		return err
	}

	userNotes, err := notes.GetNotes(user)

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	message := helpers.GenerateNotesMessage(userNotes)

	return helpers.SendLargeText(ctx, b, message, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
}
