package binders

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type key int

const ID key = 0

func Id(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strID, ok := vars["id"]
		if !ok {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, fmt.Sprintf("bad id: %s", err), http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPut {
			if err := r.ParseForm(); err != nil {
				message := fmt.Sprint("error on parsing form: ", err)
				http.Error(w, message, http.StatusInternalServerError)
				return
			}
		}

		context.Set(r, ID, id)
		h.ServeHTTP(w, r)
	}
}
