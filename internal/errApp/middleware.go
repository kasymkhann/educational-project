package errapp

import (
	"errors"
	"net/http"
)

type appHandler func(write http.ResponseWriter, request *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ea *ErrApp
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/type")
			if errors.As(err, &ea) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
					// } else if errors.Is(err, noAuthErr) {
					// 	w.WriteHeader(http.StatusUnauthorized)
					// 	w.Write(noAuthErr...)
					// 	return
					// }

				}

				err = err.(*ErrApp)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(ErrNotFound.Marshal())
				return
			}
			w.WriteHeader(http.StatusTeapot)
			w.Write(systemErr(err).Marshal())

		}
	}
}
