package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

var muxDispatcher = mux.NewRouter()

type MuxRouter struct{}

func NewMuxRouter() Router {
	return &MuxRouter{}
}

func (*MuxRouter) Serve() http.Handler {
	return muxDispatcher
}

func (*MuxRouter) Get(path string, f func(rw http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(path, f).Methods(http.MethodGet)
}

func (*MuxRouter) Post(path string, f func(rw http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(path, f).Methods(http.MethodPost)
}

func (*MuxRouter) Put(path string, f func(rw http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(path, f).Methods(http.MethodPut)
}

func (*MuxRouter) Delete(path string, f func(rw http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(path, f).Methods(http.MethodDelete)
}