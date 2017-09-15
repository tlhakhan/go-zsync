package zpool

import(
  "github.com/gorilla/mux"
  "net/http"
)

func NewHandler(r *mux.Router) {
  r.Methods("GET").Path("/list").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("zpool list"));
  })
}
