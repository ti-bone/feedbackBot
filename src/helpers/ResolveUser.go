/*
 * ResolveUser.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/models"
	"feedbackBot/src/users"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// ResolveUser resolves a user.
// Returns the resolved user and an error if any.
func ResolveUser(ctx *ext.Context, b *gotgbot.Bot) (*models.User, error) {
	user, _, err := ResolveUserWithSource(ctx, b)
	return user, err
}

// ResolveUserWithSource resolves a user.
// Returns the resolved user, whether the user is resolved from topic(source), and an error if any.
func ResolveUserWithSource(ctx *ext.Context, b *gotgbot.Bot) (*models.User, bool, error) {
	args := ctx.Args()

	topicId := ctx.EffectiveMessage.MessageThreadId

	// IsResolvable checks args length, so it's safe to access args[0-1]
	isResolvable, isValidTopicMessage, err := IsResolvable(ctx, b)

	if !isResolvable || err != nil {
		return nil, false, err
	}

	var user *models.User

	// Try to resolve user by topic ID
	if isValidTopicMessage {
		user, err = users.GetUserByTopicId(topicId)

		if err != nil {
			_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
			return nil, false, err
		}
	} else {
		user, err = ParseInputUser(args[1])

		if err != nil {
			_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
			return nil, false, err
		}
	}

	return user, isValidTopicMessage, nil
}
