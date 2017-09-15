package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"handlers/zfs"
	"log"
	"net/http"
)

func main() {

  // get a port number to start on
	var port = flag.Int("p", 8089, "api listen port")
  var zpool = flag.String("Z", "zones", "zpool for zsync-api access")

	flag.Parse()

	r := mux.NewRouter()
	zfsRouter := r.PathPrefix("/api/zfs").Subrouter()
	zfs.NewHandler(zfsRouter, *zpool)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))

}
