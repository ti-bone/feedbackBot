/*
 * Limiter.go
 * Copyright (c) ti-bone 2023-2024
 */

package rateLimiter

import "time"

// chats - is a map of chat ids to the last time of request, they have made
var chats = make(map[int64]int64)

// Check - checks if the request has made in the last N(controlled by delay parameter) seconds in this chat,
// if so, returns false, else true
//
// chatId - id of the chat
// delay - delay in seconds
func Check(chatId int64, delay int64) bool {
	lastRequest, exists := chats[chatId]

	if exists && (lastRequest+delay > time.Now().Unix()) {
		return false
	}

	chats[chatId] = time.Now().Unix()

	return true
}
