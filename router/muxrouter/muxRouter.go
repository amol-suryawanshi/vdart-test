package muxrouter

import (
	"net/http"
	"vdart-test/handlers"
	"vdart-test/repositories"
	"vdart-test/usecases"

	"github.com/gorilla/mux"
)

const (
	symbolHandler = "symbol"
	allHandler    = "all"
)

//GetMuxRouter returns instance of mux router
func GetMuxRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/api/v1/currency", getHandler(allHandler)).Methods(http.MethodGet)
	router.Handle("/api/v1/currency/{symbol}", getHandler(symbolHandler)).Methods(http.MethodGet)
	return router
}

func getHandler(key string) http.Handler {
	switch key {

	case symbolHandler:
		repo := repositories.CreateSymbolReader()
		uc := usecases.CreateSymbolUC(repo)
		return handlers.CreateHandler(uc)

	case allHandler:
		repo := repositories.CreateAllSymbolImpl()
		uc := usecases.CreateAllSymbolUC(repo)
		return handlers.CreateAllHandler(uc)
	}
	return nil
}
