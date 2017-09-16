package zfs

import (
	"encoding/json"
	"github.com/tlhakhan/golib/cmd"
	"log"
	"strings"
)

const (
	FILESYSTEM = iota
	SNAPSHOT   = iota
)

type Daemon struct {
	Pool string `json:pool`
	// all zfs strings
	FileSystems []string
	Snapshots   []string
}

func NewDaemon(pool string) *Daemon {
	d := &Daemon{Pool: pool}
	go d.run()
	return d
}

func (d *Daemon) run() {

	// zfs list -Hro name,origin -t filesystem clusters
	fsWorker := cmd.NewWorker([]string{"zfs", "list", "-Hro", "name", "-t", "filesystem", d.Pool}, 1)
	snapWorker := cmd.NewWorker([]string{"zfs", "list", "-Hro", "name", "-t", "snapshot", d.Pool}, 10)

	// listens for new output sent on worker channels
	for {
		select {
		case fsOut := <-fsWorker:
			d.processOutput(fsOut, FILESYSTEM)
		case snapOut := <-snapWorker:
			d.processOutput(snapOut, SNAPSHOT)
		default:
		}
	}
}

func (d *Daemon) processOutput(work string, fsType int) {
	switch fsType {
	case FILESYSTEM:
		log.Println("Adding zfs filesystems to Daemon struct.")
		d.FileSystems = strings.Split(work, "\n")
	case SNAPSHOT:
		log.Println("Adding zfs snapshots to Daemon struct.")
		d.Snapshots = strings.Split(work, "\n")
	}
}

func (d *Daemon) ListFileSystems(name string) []byte {

	if len(name) > 0 {
		found := false
		for _, value := range d.FileSystems {
			if value == name {
				found = true
				break
			}
		}
		if found == true {
			return ([]byte("true"))
		} else {
			return ([]byte("false"))
		}

	} else {
		j, _ := json.Marshal(d.FileSystems)
		return ([]byte(j))
	}

}

func (d *Daemon) ListSnapshots(name string) []byte {

	tmpSnapshots := make([]string, 0, 10)
	if len(name) > 0 {
		for _, val := range d.Snapshots {
			if strings.Split(val, "@")[0] == name {
				tmpSnapshots = append(tmpSnapshots, val)
			}
		}
		j, _ := json.Marshal(tmpSnapshots)
		return ([]byte(j))
	} else {
		j, _ := json.Marshal(d.Snapshots)
		return ([]byte(j))
	}

}
