package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"handlers/zfs"
	"handlers/zpool"
	"log"
	"net/http"
)

func main() {

  // get a port number to start on
	var port = flag.Int("p", 8089, "api listen port")
	flag.Parse()

	r := mux.NewRouter()

	zpoolRouter := r.PathPrefix("/api/zpool").Subrouter()
	zfsRouter := r.PathPrefix("/api/zfs").Subrouter()

	zfs.NewHandler(zfsrouter)
	zpool.NewHandler(zpoolRouter)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))

}
