package auth

import (
	"bar-code/bcs/x/web"
)

const (
	errUserNotFound = web.Unauthorized("USER_NOT_FOUND")
)
