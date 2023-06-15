/*
 * CheckAdmin.go
 * Copyright (c) ti-bone 2023
 */

package middlewares

import (
	"feedbackBot/src/db"
	"feedbackBot/src/models"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func CheckAdmin(_ *gotgbot.Bot, ctx *ext.Context) error {
	var user *models.User

	res := db.Connection.Where("user_id = ?", ctx.EffectiveSender.Id()).First(&user)

	if res.Error != nil || res.RowsAffected == 0 || !user.IsAdmin {
		return ext.EndGroups
	}

	return ext.ContinueGroups
}
