package commands

import (
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// BuildVIndex BuildVIndex
func BuildVIndex(storagePaths []string) (vindexdata.VIndex, error) {
	ret := vindexdata.VIndex{}
	for _, path := range storagePaths {
		if !isExistPathEntry(path) {
			return nil, errors.Errorf("path does not exists. path=%v", path)
		}

		isDir, err := isDirectory(path)
		if err != nil {
			return nil, errors.Wrap(err, "isDirectory failed")
		}
		if !isDir {
			return nil, errors.Errorf("%v is not a directory", path)
		}

		dirEntries, err := directoryEntries(path)
		if err != nil {
			return nil, errors.Wrap(err, "directoryEntries failed")
		}

		for _, entry := range dirEntries {
			items, err := pathRecursivly(entry.Path)
			if err != nil {
				return nil, errors.Wrap(err, "pathRecursively failed")
			}

			for i := range items {
				vindexItem, err := vindexdata.ParseFilepath([]string{path}, items[i])
				if err != nil {
					return nil, errors.Wrapf(err, "ParseFilepath failed. videoPath=%+v", items[i])
				}

				ret = append(ret, *vindexItem)
			}
		}
	}

	return ret, nil
}

func pathRecursivly(path string) ([]string, error) {
	if !isExistPathEntry(path) {
		return nil, errors.Errorf("path does not exists. path=%v", path)
	}

	isDir, err := isDirectory(path)
	if err != nil {
		return nil, errors.Wrap(err, "isDirectory failed")
	}
	if !isDir {
		return filterFilePaths([]string{path}), nil
	}

	dirEntries, err := directoryEntries(path)
	if err != nil {
		return nil, errors.Wrap(err, "directoryEntries failed")
	}

	ret := []string{}
	for i := range dirEntries {
		entry := dirEntries[i]
		entryPath := entry.Path

		if entry.EntryType == File {
			ret = append(ret, filterFilePaths([]string{entryPath})...)
			continue
		}

		filtered := filterDirPaths([]string{entryPath})
		if len(filtered) == 0 {
			continue
		}

		newPaths, err := pathRecursivly(filtered[0])
		if err != nil {
			return nil, err
		}

		ret = append(ret, newPaths...)
	}

	return ret, nil
}

func isExistPathEntry(path string) bool {
	fs := filesystem.GetFs()
	_, err := fs.Stat(path)

	return err == nil
}

func isDirectory(path string) (bool, error) {
	fs := filesystem.GetFs()
	stat, err := fs.Stat(path)
	if err != nil {
		return false, errors.Wrapf(err, "path does not exists. path=%v", path)
	}

	return stat.IsDir(), nil
}

func directoryEntries(path string) ([]directoryEntry, error) {
	fs := filesystem.GetFs()
	f, err := fs.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "fs.Open failed")
	}

	itemPaths, err := f.Readdir(-1)
	if err != nil {
		return nil, errors.Wrap(err, "f.Readdir failed")
	}

	ret := make([]directoryEntry, len(itemPaths))
	for i := range itemPaths {
		itemPath := filepath.Join(path, itemPaths[i].Name())
		stat, err := fs.Stat(itemPath)
		if err != nil {
			return nil, errors.Wrap(err, "fs.Stat failed")
		}

		var entType entryType
		if stat.IsDir() {
			entType = Directory
		} else {
			entType = File
		}

		ret[i] = directoryEntry{
			Path:      itemPath,
			EntryType: entType,
		}
	}

	return ret, nil
}

type entryType int

const (
	// File File
	File = iota + 1
	// Directory Directory
	Directory
)

type directoryEntry struct {
	Path      string
	EntryType entryType
}

var dirnameRe = regexp.MustCompile("\\d+$")
var tsRe = regexp.MustCompile("\\.mp4$")

func filterDirPaths(paths []string) []string {
	var ret []string = []string{}
	for i := range paths {
		if !dirnameRe.MatchString(paths[i]) {
			continue
		}

		ret = append(ret, paths[i])
	}

	return ret
}

func filterFilePaths(paths []string) []string {
	ret := []string{}
	for i := range paths {
		path := paths[i]
		if !tsRe.MatchString(path) {
			continue
		}

		ret = append(ret, path)
	}

	return ret
}
