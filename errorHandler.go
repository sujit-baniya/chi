package phi

import (
	"errors"
	"net/http"
)

var (
	handler = func(w http.ResponseWriter, r *http.Request, e error) {
		w.Write([]byte("unknown error"))
	}
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request) error

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		handler(w, r, err)
	}
}

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, e error)

func SetErrorHandler(fn ErrorHandlerFunc) error {
	if fn == nil {
		return errors.New("couldn't set empty error handling function")
	}

	handler = fn
	return nil
}
