package api

import (
	"github.com/gorilla/mux"
)

type Routable interface {
	RegisterRoutes(router *mux.Router)
}
