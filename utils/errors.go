package utils

import "errors"

var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrDatabase = errors.New("database error")
var ErrValidation = errors.New("validation error")
var ErrDuplicated = errors.New("duplicated error")
var ErrForbidden = errors.New("forbbiden error")
