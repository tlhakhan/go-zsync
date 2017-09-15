package zfs

import (
	"encoding/json"
	"bytes"
	"bufio"
	"github.com/tlhakhan/golib/cmd"
)

type Daemon struct {
	Pool string `json:pool`
	// all zfs datasets
	FileSystems []dataset
	Snapshots   []dataset
}

type dataset string

func NewDaemon(pool string) *Daemon {
	d := &Daemon{Pool: pool}
	go d.run()
	return d
}

func (d *Daemon) run() {

	// zfs list -Hro name,origin -t filesystem clusters
	fsWorker := cmd.NewWorker([]string{"zfs", "list", "-Hro", "name", "-t", "filesystem", d.Pool}, 10)
	snapWorker := cmd.NewWorker([]string{"zfs", "list", "-Hro", "name", "-t", "snapshot", d.Pool}, 10)

	// listens for new output sent on worker channels
	for {
		select {
		case fsOut := <-fsWorker:
			d.processFsOut(fsOut)
		case snapOut := <-snapWorker:
			d.processSnapOut(snapOut)
		default:
		}
	}
}

func (d *Daemon) processFsOut(work []byte) {

	tmpData := make([]dataset, 0, 50)
	scanner := bufio.NewScanner(bufio.NewReader(bytes.NewBuffer(work)))
	for scanner.Scan() {
		tmpData = append(tmpData, dataset(scanner.Text()))
	}
	d.FileSystems = tmpData
}

func (d *Daemon) processSnapOut(work []byte) {
	tmpData := make([]dataset, 50)
	scanner := bufio.NewScanner(bufio.NewReader(bytes.NewBuffer(work)))
	for scanner.Scan() {
		tmpData = append(tmpData, dataset(scanner.Text()))
	}
	d.Snapshots = tmpData
}

func (d *Daemon) List() []byte {
	j, _ := json.Marshal(d.FileSystems)
	return ([]byte(j))
}
