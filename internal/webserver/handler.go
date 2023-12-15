package webserver

import (
	"net/http"
	"text/template"

	"github.com/Kanbenn/mywbgonats/internal/app"
	"github.com/go-playground/pure/v5"
)

type handler struct {
	app *app.App
}

func newHandler(app *app.App) *handler {
	return &handler{app}
}

func (h *handler) getOrder(w http.ResponseWriter, r *http.Request) {
	oid := pure.RequestVars(r).URLParam("id")

	o, found := h.app.Ch.Get(oid)
	if !found {
		http.Error(w, "order_id not found "+oid, http.StatusBadRequest)
		return
	}

	pure.JSONBytes(w, http.StatusOK, o)
}

func (h *handler) getIndex(w http.ResponseWriter, r *http.Request) {
	indexPage := `
	<h3>Список поступивших заказов:</h3>
	<ul>
		{{range .}}
			<li><a href="/order/{{.}}">{{.}}</a></li>
		{{end}}
	</ul>`

	w.Header().Set("Content-Type", "text/html")

	templ := template.Must(template.New("index").Parse(indexPage))
	templ.Execute(w, h.app.Ch.GetAllKeys())
}
