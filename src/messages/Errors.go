/*
 * Errors.go
 * Copyright (c) ti-bone 2023-2024
 */

package messages

import "errors"

var UserHasNoTopic = errors.New("-400: user has no topic")
var UserAlreadyBanned = errors.New("-400: user is already banned")
var UserNotBanned = errors.New("-400: user is not banned")
var UserNotSpecified = errors.New("-400: user not specified")
var UserNotFound = errors.New("-404: user not found")
var UserIdInvalid = errors.New("-400: invalid userid")
var UserInvalid = errors.New("-400: invalid userid or username")
var InternalError = errors.New("-500: internal error")

const Protected = "protected mode is now %s for #%d;\nthereafter, they %s be able to forward new messages, received from the bot"
const Enabled = "enabled"
const Disabled = "disabled"
const Will = "will"
const Wont = "won't"
