package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func ConfigRoute(r *mux.Router) *mux.Router {
	routes := routerCreateKey

	routes = append(routes, getKeys)

	for _, rota := range routes {
		r.HandleFunc(rota.URI, rota.Func).Methods(rota.Method)
	}
	return r
}
