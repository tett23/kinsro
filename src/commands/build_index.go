package commands

import (
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

var videoRe = regexp.MustCompile("\\d{4}/\\d{2}/\\d{2}/.+\\.mp4$")

func filterPaths(items []string) []string {
	ret := []string{}
	for i := range items {
		if !videoRe.Match([]byte(items[i])) {
			continue
		}
		ret = append(ret, items[i])
	}

	return ret
}

// BuildVIndex BuildVIndex
func BuildVIndex(storagePaths []string) (vindexdata.VIndex, error) {
	ret := vindexdata.VIndex{}
	for _, path := range storagePaths {
		fs := filesystem.GetFs()
		matches, err := afero.Glob(fs, filepath.Join(path, "/*/*/*/*.mp4"))
		if err != nil {
			return nil, errors.Wrap(err, "Glob failed")
		}

		filtered := filterPaths(matches)
		ret := make(vindexdata.VIndex, len(filtered))
		for i := range filtered {
			vindexItem, err := vindexdata.ParseFilepath([]string{path}, filtered[i])
			if err != nil {
				return nil, errors.Wrapf(err, "ParseFilepath failed. videoPath=%+v", filtered[i])
			}

			ret[i] = *vindexItem
		}

		return ret, nil
	}

	return ret, nil
}
