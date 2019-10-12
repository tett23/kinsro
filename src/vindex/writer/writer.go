package writer

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

type command struct {
	callbackCh chan bool
	item       vindexdata.VIndexItem
}

var ch chan command

func init() {
	ch = make(chan command)

	go processor()
}

// Append Append
func Append(item vindexdata.VIndexItem) chan bool {
	callbackCh := make(chan bool)
	data := command{callbackCh: callbackCh, item: item}
	ch <- data

	return callbackCh
}

func lockFilePath(indexPath string) string {
	return indexPath + ".lock"
}

func processor() {
	for {
		command := <-ch

		ok, err := isLocked()
		if err != nil {
			panic(errors.Wrap(err, "lock error"))
		}
		if ok {
			command.callbackCh <- <-Append(command.item)
			continue
		}

		ok, err = appendRecord(command.item)
		if err != nil {
			panic(errors.Wrap(err, "append error"))
		}
		if !ok {
			command.callbackCh <- false
			continue
		}

		command.callbackCh <- true
	}
}

func isLocked() (bool, error) {
	config := config.GetConfig()
	fs := filesystem.GetFs()

	stat, err := fs.Stat(lockFilePath(config.VIndexPath))
	if err == nil {
		return false, nil
	}
	if stat == nil {
		return false, nil
	}

	if stat.IsDir() {
		return false, errors.New("stat type error")
	}

	lockFile, err := fs.OpenFile(lockFilePath(config.VIndexPath), os.O_RDONLY, 0644)
	if err != nil {
		return true, err
	}

	var timestampBytes []byte
	_, err = lockFile.Read(timestampBytes)
	if err != nil {
		return true, err
	}

	timestampStr := string(timestampBytes)
	timestamp, err := strconv.Atoi(timestampStr)
	if err != nil {
		return true, err
	}

	now := time.Now().Unix()

	return !(int64(timestamp) < now), nil
}

func appendRecord(vindexItem vindexdata.VIndexItem) (bool, error) {
	config := config.GetConfig()
	fs := filesystem.GetFs()

	file, err := fs.OpenFile(config.VIndexPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return false, err
	}

	if _, err = file.Write(vindexItem.ToBinary()); err != nil {
		return false, err
	}

	return true, nil
}
