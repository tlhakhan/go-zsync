package zfs

import (
	"github.com/gorilla/mux"
	"net/http"
	"workers/zfs"
)

func NewHandler(r *mux.Router, zpool string) {

	zfsD := zfs.NewDaemon(zpool)

	r.Methods("GET").Path("/list").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write(zfsD.List())
	})
  
}
