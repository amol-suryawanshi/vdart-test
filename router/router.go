package router

import (
	"net/http"
	"vdart-test/router/muxrouter"
)

func NewRouter() http.Handler {
	return muxrouter.GetMuxRouter()
}
