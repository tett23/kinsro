package main

import (
	"github.com/tett23/kinsro/scripts/encode_old_files/src/videotmp"
)

func main2() {
	queue, err := videotmp.ReadQueue(QueuePath)
	if err != nil {
		panic(err)
	}

	queue.Deque()
}
