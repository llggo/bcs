package auth

import (
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

const (
	errUserNotFound = web.Unauthorized("USER_NOT_FOUND")
)
