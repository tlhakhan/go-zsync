package zfs

import(
  "github.com/gorilla/mux"
  "net/http"
)

func NewHandler(r *mux.Router) {
  r.Methods("GET").Path("{poolName}/list").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("zfs list"));
  })
}
