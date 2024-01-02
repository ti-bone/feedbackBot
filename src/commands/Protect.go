/*
 * Protect.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/config"
	"feedbackBot/src/db"
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"feedbackBot/src/models"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"os"
)

func Protect(b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()
	topicId := ctx.EffectiveMessage.MessageThreadId
	isValidTopicMessage := topicId != 0 && topicId != config.CurrentConfig.LogsTopicID

	if !isValidTopicMessage && len(args) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(
			b,
			messages.UserNotSpecified.Error()+"\nHint: you can send this command to any user-related topic.",
			&gotgbot.SendMessageOpts{},
		)

		return err
	}

	var user *models.User
	var err error

	// Try to resolve user by topic ID
	if isValidTopicMessage {
		user, err = helpers.GetUserByTopicId(topicId)

		if err != nil {
			_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
			return err
		}
	} else {
		// args[1] is 146% existing, because we checked it above
		user, err = helpers.ParseInputUser(args[1])

		if err != nil {
			_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
			return err
		}
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
