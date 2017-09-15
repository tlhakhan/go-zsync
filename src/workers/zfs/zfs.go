package zfs

import(
	"encoding/json"
)

type Daemon struct {
  pool string `json:pool`
}

func NewDaemon(pool string) *Daemon {
  return &Daemon{pool:pool}
}

func (d *Daemon) List() []byte {
  j,_ := json.Marshal(d.pool);
  return([]byte(j));
}
