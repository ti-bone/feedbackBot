/*
 * Main.go
 * Copyright (c) ti-bone 2023-2024
 */

package main

import (
	"feedbackBot/src/botHandlers"
	"feedbackBot/src/commands"
	"feedbackBot/src/config"
	"feedbackBot/src/db"
	"feedbackBot/src/helpers"
	"feedbackBot/src/middlewares"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"log"
	"net/http"
	"time"
)

func main() {
	config.LoadConfig("config.json")

	db.Init()

	b, err := gotgbot.NewBot(config.CurrentConfig.BotToken, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: 5 * time.Second,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})

	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update:", err.Error())

				// Log error to chat
				err = helpers.LogError(err.Error(), b, ctx)
				if err != nil {
					log.Println("an error occurred while logging error:", err.Error())
				}

				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})

	// Handlers
	dispatcher := updater.Dispatcher

	// Start command
	dispatcher.AddHandler(handlers.NewCommand("start", commands.Start))
	dispatcher.AddHandler(handlers.NewCommand("id", commands.Id))

	// Middleware for syncing user in DB for any update from a user
	dispatcher.AddHandlerToGroup(handlers.NewMessage(message.All, middlewares.SyncUser), -1)

	dispatcher.AddHandler(handlers.NewMessage(message.Private, botHandlers.Message))

	// Admin commands
	dispatcher.AddHandlerToGroup(handlers.NewMessage(message.All, middlewares.CheckAdmin), 1)
	dispatcher.AddHandlerToGroup(handlers.NewCommand("ban", commands.Ban), 1)
	dispatcher.AddHandlerToGroup(handlers.NewCommand("unban", commands.Unban), 1)
	dispatcher.AddHandlerToGroup(handlers.NewCommand("protect", commands.Protect), 1)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(message.Supergroup, botHandlers.Response), 1)

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: false,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 0,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}

	fmt.Printf("@%s has been started...\n", b.User.Username)
	helpers.LogMessage(fmt.Sprintf("#SYSTEM\n@%s has been started...", b.User.Username), b)

	updater.Idle()
}
