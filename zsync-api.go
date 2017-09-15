package main

import (
	"flag"
  "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
  	"handlers/zfs"
  	"handlers/zpool"
)

func main() {
	var port = flag.Int("p", 8089, "api listen port")
	flag.Parse()

	r := mux.NewRouter()

	zpoolRouter := r.PathPrefix("/api/zpool").Subrouter()
	zfsRouter := r.PathPrefix("/api/zfs").Subrouter()

 zfs.NewHandler(zpoolRouter)
 zpool.NewHandler(zfsRouter)

  endpoint := fmt.Sprintf(":%d", *port);
  log.Println(endpoint)
	log.Fatal(http.ListenAndServe(endpoint, r))

}
