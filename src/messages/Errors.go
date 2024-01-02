/*
 * Errors.go
 * Copyright (c) ti-bone 2023-2024
 */

package messages

import "errors"

var UserAlreadyBanned = errors.New("-400: user is already banned")
var UserNotBanned = errors.New("-400: user is not banned")
var UserNotSpecified = errors.New("-400: user not specified")
var UserIdInvalid = errors.New("-400: invalid userid")
var UserInvalid = errors.New("-400: invalid userid or username")

var UserNotFound = errors.New("-404: user not found")

var InternalError = errors.New("-500: internal error")
