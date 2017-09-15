package zpool

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"workers/zpool"
)

func NewHandler(r *mux.Router) {

	d := zpool.NewDaemon()

	r.Methods("GET").Path("/list").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	list,_ := json.Marshal(d.List())
		w.Write([]byte(list))
	})
}
