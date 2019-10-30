package filelock

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// Filelock Filelock
func Filelock(fs afero.Fs, filepath string, fn func() error) error {
	if err := Lock(fs, filepath); err != nil {
		return err
	}
	defer free(fs, filepath)

	if err := fn(); err != nil {
		return err
	}

	return nil
}

// IsFree IsFree
func IsFree(fs afero.Fs, path string) (bool, error) {
	ok, _ := afero.Exists(fs, path+".lock")
	if !ok {
		return true, nil
	}

	lockTo, err := readLockTimestamp(fs, path)
	if err != nil {
		return false, err
	}

	now := time.Now().Unix()

	return lockTo < now, nil
}

func free(fs afero.Fs, path string) error {
	ok, _ := afero.Exists(fs, path+".lock")
	if !ok {
		return nil
	}

	return fs.Remove(path + ".lock")
}

func readLockTimestamp(fs afero.Fs, path string) (int64, error) {
	data, err := afero.ReadFile(fs, path+".lock")
	if err != nil {
		return -1, err
	}

	lockTo, err := strconv.Atoi(string(data))
	if err != nil {
		return -1, err
	}

	return int64(lockTo), nil
}

// Lock Lock
func Lock(fs afero.Fs, path string) error {
	ok, err := IsFree(fs, path)
	if err != nil {
		return err
	}
	if !ok {
		lockTo, _ := readLockTimestamp(fs, path)
		return errors.Errorf("file locked. path=%v %v", path, lockTo)
	}

	now := time.Now().Unix()
	lockTo := now + 60*60*24

	timestampString := strconv.Itoa(int(lockTo))

	return afero.WriteFile(fs, path+".lock", []byte(timestampString), 0744)
}
