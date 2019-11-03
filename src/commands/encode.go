package commands

import (
	"log"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/encode"
	"github.com/tett23/kinsro/src/fileentry/moveentries"
	"github.com/tett23/kinsro/src/fileentry/mpegts"
	"github.com/tett23/kinsro/src/storages"
	"github.com/tett23/kinsro/src/syscalls"
)

// EncodeOptions EncodeOptions
type EncodeOptions struct {
	RemoveTS bool
}

// EncodeTS EncodeTS
func EncodeTS(conf *config.Config, ss *syscalls.Syscalls, fs afero.Fs, tsPath string, options EncodeOptions) error {
	ts, err := mpegts.NewMpegTS(tsPath)
	if err != nil {
		return err
	}

	if !ts.IsProcessable(fs) {
		return nil
	}

	if err := encode.Encode(conf, fs, ts); err != nil {
		return err
	}

	group, err := ts.ToEntryGroup(fs)
	if err != nil {
		return err
	}

	storage, err := storages.NewStoragesFromPaths(conf.StoragePaths).MostSpacefulStorage(ss.Statfs)
	if err != nil {
		return err
	}

	err = moveentries.MoveEntryGroup(fs, group, storage)
	if err != nil {
		return err
	}

	vindexItem, err := ts.ToVIndexItem(storage)
	if err != nil {
		return err
	}

	if err := AppendToIndex(conf, fs, vindexItem); err != nil {
		return err
	}

	if err := Symlink(conf, ss, fs, vindexItem.HexDigest()); err != nil {
		return err
	}

	if options.RemoveTS {
		if err := ts.Remove(fs); err != nil {
			return err
		}
	}

	return nil
}

// EncodeTSAll EncodeTSAll
func EncodeTSAll(conf *config.Config, ss *syscalls.Syscalls, fs afero.Fs, tmpPath string, options EncodeOptions) error {
	var ignorePaths []string
	for true {
		tsPath, ok, err := headTS(fs, tmpPath, ignorePaths)
		log.Println(tsPath)
		if err != nil {
			return err
		}
		if !ok {
			break
		}

		ignorePaths = append(ignorePaths, tsPath)

		if err := EncodeTS(conf, ss, fs, tsPath, options); err != nil {
			log.Printf("%v\npath=%v\n"+err.Error(), tsPath)
			return err
		}
	}

	return nil
}

func headTS(fs afero.Fs, tsDir string, ignorePaths []string) (string, bool, error) {
	matches, err := afero.Glob(fs, filepath.Join(tsDir, "*.ts"))
	if err != nil {
		return "", false, err
	}
	for i := range matches {
		path := matches[i]
		if _, err := mpegts.NewMpegTS(path); err != nil {
			continue
		}
		if isMatchIgnorePath(path, ignorePaths) {
			continue
		}

		return path, true, nil
	}

	return "", false, nil
}

func isMatchIgnorePath(path string, ignorePaths []string) bool {
	for i := range ignorePaths {
		if path == ignorePaths[i] {
			return true
		}
	}

	return false
}
