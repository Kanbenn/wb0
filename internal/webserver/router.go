package webserver

import (
	"net/http"

	"github.com/go-playground/pure/v5"
	mw "github.com/go-playground/pure/v5/_examples/middleware/logging-recovery"
)

func newRouter(h *handler) http.Handler {
	p := pure.New()
	p.Use(mw.LoggingAndRecovery(true))

	p.Get("/", h.getIndex)
	p.Get("/order/:id", h.getOrder)

	return p.Serve()
}
