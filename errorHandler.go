package phi

import (
	"errors"
	"net/http"
)

var (
	handler = func(w http.ResponseWriter, r *http.Request) error { return nil }
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request) error

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		handler(w, r)
	}
}

func SetErrorHandler(fn ErrorHandler) error {
	if fn == nil {
		return errors.New("couldn't set empty error handling function")
	}

	handler = fn
	return nil
}
