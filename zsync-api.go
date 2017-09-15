package main

import (
	"flag"
  "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var port = flag.Int("p", 8089, "api listen port")
	flag.Parse()

	r := mux.NewRouter()
	zpoolRouter := r.PathPrefix("/api/zpool").Subrouter()
	zfsRouter := r.PathPrefix("/api/zfs").Subrouter()
	zpoolRouter.Methods("GET").Path("/list").HandlerFunc(zpoolHandler)
	zfsRouter.Methods("GET").Path("/{poolName}/list").HandlerFunc(zfsHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func zpoolHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Zpool Test"))
}

func zfsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Zfs Test"))
}
