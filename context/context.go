package context

import (
	"context"
	"fmt"
	_ "log"
	"net/http"
)

type Store interface {
	Fetch(ctxt context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			// log.Fatal("error: ", err)
			return
		}

		fmt.Fprint(w, data)
	}
}