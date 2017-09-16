package zfs

import (
	"github.com/gorilla/mux"
	"net/http"
	"workers/zfs"
)

func NewHandler(r *mux.Router, zpool string) {

	zfsD := zfs.NewDaemon(zpool)

	r.Methods("GET").Path("/snapshots").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		name := req.FormValue("name")
		res.Write(zfsD.ListSnapshots(name))
	})

	r.Methods("GET").Path("/filesystems").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		name := req.FormValue("name")
		res.Write(zfsD.ListFileSystems(name))
	})
}
