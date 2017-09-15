package zfs

import (
	"encoding/json"
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

	d := &Daemon{pool: pool}
	d.run()
	return d
}

func (d *Daemon) run() {

	// zfs list -Hro name,origin -t filesystem clusters
	fsWorker := cmd.NewWorker([]string{"zfs", "list", "-Hro", "name,origin", "-t", "filesystem", d.Pool}, 10)
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

	//tmpData := make([]dataset, 50)

	//reader := bufio.NewReader(bytes.NewBuffer(work))

	scanner := bufio.NewScanner(work)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func (d *Daemon) processSnapOut(work []byte) {
	scanner := bufio.NewScanner(work)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func (d *Daemon) List() []byte {
	j, _ := json.Marshal(d.pool)
	return ([]byte(j))
}
