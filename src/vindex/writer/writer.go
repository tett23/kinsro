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
	item       *vindexdata.VIndexItem
	config     *config.Config
}

var ch chan command

func init() {
	ch = make(chan command)

	go processor()
}

// Append Append
func Append(conf *config.Config, item *vindexdata.VIndexItem) chan bool {
	callbackCh := make(chan bool)
	data := command{callbackCh: callbackCh, item: item, config: conf}
	ch <- data

	return callbackCh
}

// CreateNewIndexFile CreateNewIndexFile
func CreateNewIndexFile(conf *config.Config, vindex vindexdata.VIndex) error {
	indexPath := conf.VIndexPath
	fs := filesystem.GetFs()

	file, err := fs.OpenFile(indexPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return errors.Wrapf(err, "fs.OpenFile failed. path=%v", indexPath)
	}
	defer func() {
		file.Close()
	}()

	bin, err := vindex.ToBinary()
	if err != nil {
		return errors.Wrapf(err, "vindex.ToBinary failed")
	}

	_, err = file.Write(bin)
	if err != nil {
		return errors.Wrapf(err, "file.Write failed")
	}

	return nil
}

func lockFilePath(indexPath string) string {
	return indexPath + ".lock"
}

func processor() {
	for {
		command := <-ch

		ok, err := isLocked(command.config.VIndexPath)
		if err != nil {
			panic(errors.Wrap(err, "lock error"))
		}
		if ok {
			command.callbackCh <- <-Append(command.config, command.item)
			continue
		}

		ok, err = appendRecord(command.config, command.item)
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

func isLocked(vindexPath string) (bool, error) {
	fs := filesystem.GetFs()

	stat, err := fs.Stat(lockFilePath(vindexPath))
	if err == nil {
		return false, nil
	}
	if stat == nil {
		return false, nil
	}

	if stat.IsDir() {
		return false, errors.New("stat type error")
	}

	lockFile, err := fs.OpenFile(lockFilePath(vindexPath), os.O_RDONLY, 0644)
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

func appendRecord(conf *config.Config, vindexItem *vindexdata.VIndexItem) (bool, error) {
	fs := filesystem.GetFs()

	file, err := fs.OpenFile(conf.VIndexPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return false, err
	}

	if _, err = file.Write(vindexItem.ToBinary()); err != nil {
		return false, err
	}

	return true, nil
}
