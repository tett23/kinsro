package commands

import (
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

// func EncodeTSAll(conf *config.Config, fs afero.Fs, tmpPath string) error {
// 	for ts := headTs(); ts != nil; headTS() {
// 		EncodeTS(conf, fs, ts)
// 	}

// 	return nil
// }

// func headTS(fs afero.Fs, tsDir string) (string, bool, error) {
// 	matches, err := afero.Glob(fs, "*.ts")
// 	if err != nil {
// 		return "", false, err
// 	}
// 	if len(matches) == 0 {
// 		return "", false, nil
// 	}

// 	ts := matches[0]
// 	for i := range matches {
// 	&& filelock.IsFree(fs, matches[i])
// 	}

// 	return ts, true, nil
// }
