/*
 * Limiter.go
 * Copyright (c) ti-bone 2023
 */

package rateLimiter

import "time"

// Chats is a map of chat ids to the last time of request, they have made
var chats = make(map[int64]int64)

// Check - checks if the request has made in the last 10 seconds in this chat, if so, returns false, else true
func Check(chatId int64) bool {
	lastRequest, exists := chats[chatId]

	if exists && (lastRequest+10 > time.Now().Unix()) {
		chats[chatId] = time.Now().Unix()

		return false
	} else {
		chats[chatId] = time.Now().Unix()

		return true
	}
}
