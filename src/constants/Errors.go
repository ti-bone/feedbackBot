/*
 * Errors.go
 * Copyright (c) ti-bone 2023-2024
 */

package constants

import "errors"

var UserAlreadyBanned = errors.New("-400: user is already banned")
var UserNotBanned = errors.New("-400: user is not banned")
var UserNotSpecified = errors.New("-400: user not specified")
var NoteIdInvalid = errors.New("-400: invalid noteid")
var NoteTextNotSpecified = errors.New("-400: invalid note text")
var UserIdInvalid = errors.New("-400: invalid userid")
var UserInvalid = errors.New("-400: invalid userid or username")
var NoteTooLong = errors.New("-400: note too long")

var UserNotFound = errors.New("-404: user not found")
var NotesNotFound = errors.New("-404: notes for this user not found")
var NoteNotFound = errors.New("-404: note with such id not found")

var InternalError = errors.New("-500: internal error")
