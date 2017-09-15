package zfs

import(
  "fmt"
)

type Daemon struct {
  Pool: string `json:pool`
}

func NewDaemon(pool string) {
  return &Daemon{
    pool: pool
  }
}

func (*d Daemon) List() {
  j,_ := json.Marshal(d.Pool);
  return([]byte(j));
}
