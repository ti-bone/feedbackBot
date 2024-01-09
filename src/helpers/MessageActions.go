/*
 * MessageActions.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

func IsServiceMessage(ctx *ext.Context) bool {
	return ctx.EffectiveMessage.IsAutomaticForward ||
		ctx.EffectiveMessage.ForumTopicReopened != nil ||
		ctx.EffectiveMessage.ForumTopicClosed != nil ||
		ctx.EffectiveMessage.ForumTopicEdited != nil ||
		ctx.EffectiveMessage.ForumTopicCreated != nil
}
