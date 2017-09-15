package zpool

import (
	"bufio"
	"bytes"
	"github.com/tlhakhan/golib/cmd"
	"io"
	"log"
)

type pool string

type ZpoolDaemon struct {
	pools []pool `json:"pools"`
}

func NewDaemon() *ZpoolDaemon {
	daemon := &ZpoolDaemon{}
	go daemon.run()
	return daemon
}

func (z *ZpoolDaemon) run()  {
	// Create a new command worker, sleep interval set to 10 seconds
	work := cmd.NewWorker([]string{"zpool", "list", "-Ho", "name"}, 10)

	// Listens for new output sent on work channel
	for {
		select {
		case output := <-work:
			z.processWork(output)
		}
	}
}

// Process zpool command output
func (z *ZpoolDaemon) processWork(work []byte) {

	tmpData := make([]pool, 2)
	reader := bufio.NewReader(bytes.NewBuffer(work))

	for {
		line, err := reader.ReadString('\n')
		switch err {
		case nil:
			tmpData = append(tmpData, pool(line))
		case io.EOF:
			z.pools = tmpData
			break
		default:
			log.Fatal(err)
		}
	}

}

// List
func (z *ZpoolDaemon) List() []pool {
	return z.pools
}
