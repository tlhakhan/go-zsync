package zfs

import (
	"github.com/gorilla/mux"
  "git"
	"net/http"
)

func NewHandler(r *mux.Router) {
	r.Methods("GET").Path("{poolName}/list").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("zfs list"))
	})
}
