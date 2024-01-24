/*
 * CheckLanguage.go
 * Copyright (c) ti-bone 2023-2024
 */

package middlewares

import (
	"feedbackBot/src/config"
	"feedbackBot/src/rateLimiter"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gorm.io/gorm/utils"
)

func CheckLanguage(b *gotgbot.Bot, ctx *ext.Context) error {
	filterConfig := config.CurrentConfig.LanguageFilter

	if !filterConfig.Enabled {
		return ext.ContinueGroups // Skip this middleware
	}

	userLanguage := ctx.EffectiveUser.LanguageCode

	if utils.Contains(filterConfig.ForbiddenLanguages, userLanguage) {
		// If the filter matches, rate-limit the user for config-specified time
		if !rateLimiter.Check(ctx.EffectiveChat.Id, filterConfig.ErrorRateLimit) {
			return ext.EndGroups // Stop handling this update
		}

		_, err := ctx.EffectiveMessage.Reply(
			b,
			filterConfig.Message,
			&gotgbot.SendMessageOpts{
				ParseMode: "HTML",
			},
		)

		if err != nil {
			return err // Return error if something went wrong
		}

		return ext.EndGroups // Stop handling this update
	}

	return ext.ContinueGroups
}
