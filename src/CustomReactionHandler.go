/*
 * CustomReactionHandler.go
 * Copyright (c) ti-bone 2023-2024
 */

package main

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Reaction struct {
	Filter   bool
	Response func(b *gotgbot.Bot, ctx *ext.Context) error
}

func NewReaction(f bool, r func(b *gotgbot.Bot, ctx *ext.Context) error) Reaction {
	return Reaction{
		Filter:   f,
		Response: r,
	}
}

func (r Reaction) CheckUpdate(b *gotgbot.Bot, ctx *ext.Context) bool {
	if ctx.MessageReaction == nil {
		return false
	}
	return true
}

func (r Reaction) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	return r.Response(b, ctx)
}

func (r Reaction) Name() string {
	return fmt.Sprintf("reaction_%p", r.Response)
}
