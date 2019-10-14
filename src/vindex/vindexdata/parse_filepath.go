package vindexdata

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var filenameRe = regexp.MustCompile("\\d{4}/\\d{2}/\\d{2}/.+\\.mp4$")

// ParseFilepath ParseFilepath
func ParseFilepath(storagePaths []string, path string) (*VIndexItem, error) {
	if !filenameRe.MatchString(path) {
		return nil, errors.Errorf("MatchString failed. path=%v", path)
	}

	storagePath, err := parseStoragePath(storagePaths, path)
	if err != nil {
		errors.Wrapf(err, "")
	}

	date, err := toDate(storagePath, path)
	if err != nil {
		return nil, errors.Wrapf(err, "toDate failed. path=%+v", path)
	}

	return &VIndexItem{
		Storage:  filepath.Base(storagePath),
		Filename: path,
		Date:     date,
	}, nil
}

func toDate(base, path string) (uint64, error) {
	pathItems := strings.Split(path[len(base):], "/")
	year, err := strconv.Atoi(pathItems[1])
	if err != nil {
		return 0, errors.Wrapf(err, "strconv.Atoi failed. value=%v", pathItems[1])
	}
	month, err := strconv.Atoi(pathItems[2])
	if err != nil {
		return 0, errors.Wrapf(err, "strconv.Atoi failed. value=%v", pathItems[2])
	}
	day, err := strconv.Atoi(pathItems[3])
	if err != nil {
		return 0, errors.Wrapf(err, "strconv.Atoi failed. value=%v", pathItems[3])
	}

	return uint64(year*10000 + month*100 + day), nil
}

func parseStoragePath(storagePaths []string, path string) (string, error) {
	var storagePath string
	for i := range storagePaths {
		storagePath = storagePaths[i]
		if strings.HasPrefix(path, storagePath) {
			break
		}
	}
	if storagePath == "" {
		return "", errors.Errorf("")
	}

	return storagePath, nil
}
