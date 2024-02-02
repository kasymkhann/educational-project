package sort

import (
	"context"
	"net/http"
	"strings"
)

const (
	ASK               = "ask"
	DESC              = "desc"
	OptionsContextKey = "order_key"
)

func Middleware(handler http.HandlerFunc, defaultSortFiled, defaultSortOrder string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sortBy := r.URL.Query().Get("sort_by")
		sortOrder := r.URL.Query().Get("sort_order")

		if sortBy == "" {
			sortBy = defaultSortFiled
		}

		if sortOrder == "" {
			sortOrder = defaultSortOrder
		} else {
			upper := strings.ToUpper(sortOrder)
			if upper != ASK && upper != DESC {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad sort order"))
				
				return
			}
		}
		opt := Options{
			Filed: sortBy,
			Order: sortOrder,
		}
		ctx := context.WithValue(r.Context(), OptionsContextKey, opt)
		r = r.WithContext(ctx)

		handler(w, r)
	}

}

type Options struct {
	Filed string
	Order string
}
