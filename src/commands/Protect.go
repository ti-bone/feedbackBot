/*
 * Protect.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/db"
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"os"
)

func Protect(b *gotgbot.Bot, ctx *ext.Context) error {
	// Resolve user
	user, err := helpers.ResolveUser(ctx, b)

	if err != nil {
		return err
	}

	var result string

	// Set the result text depending on user's is_protected field
	if user.IsProtected {
		result = messages.Disabled
	} else {
		result = messages.Enabled
	}

	// Update user's is_protected field in the database
	res := db.Connection.Model(&user).Update("is_protected", !user.IsProtected)
	err = res.Error

	if err != nil {
		log.SetOutput(os.Stderr)
		log.Println(err)

		_, err := ctx.EffectiveMessage.Reply(b, messages.InternalError.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	// Send the result text to the user
	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf(messages.Protected, result, user.UserID), &gotgbot.SendMessageOpts{})

	return err
}
