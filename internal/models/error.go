package models

import "errors"

var (
	ErrNoOrder = errors.New("this order was not found: ")
)
