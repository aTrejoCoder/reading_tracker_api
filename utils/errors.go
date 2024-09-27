package utils

import "errors"

var ErrNotFound = errors.New("not found")
var ErrDatabase = errors.New("database error")
var ErrValidation = errors.New("validation error")
