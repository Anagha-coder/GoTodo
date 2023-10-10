package controller

import (
	"net/http"
)

func Register() *http.ServeMux {

	mux := http.NewServeMux() //creats a new multiplexer
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", crud()) //HandleFunc used while registering endpoints
	return mux

}
