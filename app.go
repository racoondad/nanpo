package nanpo

import (
	"github.com/racoondad/nanpo/internal"
)

var app internal.Application

// NewApplication Create an application.
func NewApplication() Application {
	return internal.NewApplication()
}
