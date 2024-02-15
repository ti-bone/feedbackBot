/*
 * Main.go
 * Copyright (c) ti-bone 2023-2024
 */

package main

import (
	"feedbackBot/src/commands"
	"feedbackBot/src/config"
	"feedbackBot/src/constants"
	"feedbackBot/src/db"
	"feedbackBot/src/handlers"
	"feedbackBot/src/helpers"
	"feedbackBot/src/middlewares"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	updateHandlers "github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"log"
	"net/http"
	"time"
)

func main() {
	config.LoadConfig("config.json")

	db.Init()

	b, err := gotgbot.NewBot(config.CurrentConfig.BotToken, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			Client: http.Client{},
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: 5 * time.Second,
				APIURL:  gotgbot.DefaultAPIURL,
			},
		},
	})

	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())

			// Log error to chat
			errorText := constants.InternalError.Error()

			if config.CurrentConfig.DiscloseErrorInternals {
				errorText = err.Error()
			}

			err = helpers.LogError(errorText, b, ctx)
			if err != nil {
				log.Println("an error occurred while logging error:", err.Error())
			}

			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	// Handlers
	updater := ext.NewUpdater(dispatcher, nil)

	// Middleware for syncing user in DB for any update from a user
	dispatcher.AddHandlerToGroup(updateHandlers.NewMessage(message.All, middlewares.SyncUser), -1)

	/*
	 * User handlers
	 */

	// Middleware for language filtering
	dispatcher.AddHandlerToGroup(updateHandlers.NewMessage(message.Private, middlewares.CheckLanguage), 0)

	// Command handlers
	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("start", commands.Start), 0)
	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("id", commands.Id), 0)

	// Message handlers
	dispatcher.AddHandlerToGroup(updateHandlers.NewMessage(message.Private, handlers.Message), 0)
	dispatcher.AddHandlerToGroup(NewReaction(true, handlers.Reaction), 0)

	/*
	 * Admin handlers
	 */

	// Middleware for admin checking
	dispatcher.AddHandlerToGroup(updateHandlers.NewMessage(message.All, middlewares.CheckAdmin), 1)

	// Command handlers
	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("ban", commands.Ban), 1)
	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("unban", commands.Unban), 1)
	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("protect", commands.Protect), 1)

	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("add", commands.AddNote), 1)
	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("del", commands.DelNote), 1)
	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("get", commands.GetNotes), 1)

	dispatcher.AddHandlerToGroup(updateHandlers.NewCommand("mdel", commands.DelMessage), 1)

	// Topic-related handlers
	dispatcher.AddHandlerToGroup(updateHandlers.NewMessage(message.TopicReopened, handlers.TopicReopened), 1)
	dispatcher.AddHandlerToGroup(updateHandlers.NewMessage(message.TopicClosed, handlers.TopicClosed), 1)

	// Response handler
	dispatcher.AddHandlerToGroup(updateHandlers.NewMessage(message.Supergroup, handlers.Response), 1)

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: false,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 60,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 70,
			},
			AllowedUpdates: []string{"message", "edited_message", "message_reaction", "my_chat_member"},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	fmt.Printf("@%s has been started...\n", b.User.Username)
	helpers.LogMessage(fmt.Sprintf("#SYSTEM\n@%s has been started...", b.User.Username), b)

	updater.Idle()
}
