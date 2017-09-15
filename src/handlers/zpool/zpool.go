package zpool

import (
	"github.com/gorilla/mux"
	"net/http"
	"workers/zpool"
)

func NewHandler(r *mux.Router) {

	d := zpool.NewDaemon()

	r.Methods("GET").Path("/list").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(d.List())
	})
}
