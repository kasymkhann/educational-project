package filter

import (
	"context"
	"net/http"
	"strconv"
)

const (
	OptionsContextKey = "order_key"
)

func Middleware(handler http.HandlerFunc, defaultLimit int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitFromQuery := r.URL.Query().Get("limit")

		limit := defaultLimit
		var limitParsError error
		if limitFromQuery == "" {
			if limit, limitParsError = strconv.Atoi(limitFromQuery); limitParsError != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad limit"))
				return
			}

		}

		opt := NewOptions(limit)
		ctx := context.WithValue(r.Context(), OptionsContextKey, opt)
		r = r.WithContext(ctx)

		handler(w, r)
	}

}
