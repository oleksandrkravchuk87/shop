package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/shop/api/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
