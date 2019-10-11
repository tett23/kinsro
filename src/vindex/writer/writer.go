package writer

import (
	"os"

	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

type command struct {
	callbackCh chan bool
	item       vindexdata.BinaryIndexItem
}

var ch chan command

func init() {
	ch = make(chan command)

	go processor()
}

// Append Append
func Append(item vindexdata.BinaryIndexItem) chan bool {
	callbackCh := make(chan bool)
	data := command{callbackCh: callbackCh, item: item}
	ch <- data

	return callbackCh
}

func lockFilePath(indexPath string) string {
	return indexPath + ".lock"
}

func processor() {
	config := config.GetConfig()
	fs := filesystem.GetFs()

	for {
		command := <-ch

		// なんかロック関連の処理する
		stat, err := fs.Stat(lockFilePath(config.VIndexPath))
		if err != nil {
			// command.callbackCh <- false
			// panic(err)
		}

		file, err := fs.OpenFile(config.VIndexPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			command.callbackCh <- false
			panic(err)
		}

		stat, err = file.Stat()
		if err != nil {
			command.callbackCh <- false
			panic(err)
		}

		var offset int64
		if stat.Size() == 0 {
			offset = 0
		} else {
			offset = stat.Size()
		}

		bytes := []byte(command.item.Filename)
		_, err = file.WriteAt(bytes, offset)
		if err != nil {
			command.callbackCh <- false
			panic(err)
		}

		command.callbackCh <- true
	}
}
