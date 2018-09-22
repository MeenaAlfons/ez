// +build dev

package data

import (
	"net/http"
)

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("assets")

//go:generate go run assets_generate.go
