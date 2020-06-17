package router

import (
	"net/http"
)

type Logger interface {
	Error(a ...interface{})
}

type Router interface {
	Route(r *http.Request) error
}

func NewHandler(
	logger Logger,
	router Router,
	next http.Handler,
) *handler {
	return &handler{
		logger,
		router,
		next,
	}
}

type handler struct {
	Logger
	Router
	http.Handler
}

func (self *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := self.Router.Route(r); err != nil {
		w.WriteHeader(http.StatusNotFound)
		self.Logger.Error(err)
		return
	}

	self.Handler.ServeHTTP(w, r)
}
