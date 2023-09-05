package router

import (
	"example/hello/src/router/routers"

	"github.com/gorilla/mux"
)

func RunRouteConfig() *mux.Router {
	r := mux.NewRouter()
	return routers.ConfigRoute(r)
}
