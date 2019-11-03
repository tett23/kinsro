package storages

import (
	"syscall"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/syscalls"
)

func buildStatfs() syscalls.Statfs {
	count := 0

	return func(path string, stat *syscall.Statfs_t) error {
		stat.Bsize = 4086
		stat.Bfree = uint64(1000 * (count + 1))
		count++

		return nil
	}
}

func TestStorages__Storage__Name(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns storage name", func(t *testing.T) {
			storage := NewStorage("/media/video1")
			actual := storage.Name()
			assert.Equal(t, actual, "video1")
		})
	})
}

func TestStorages__Storage__Free(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns size of free spaces", func(t *testing.T) {
			storage := NewStorage("/media/video1")
			_, err := storage.Free(buildStatfs())
			assert.NoError(t, err)
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("statfs returns error", func(t *testing.T) {
			statfs := func(path string, stat *syscall.Statfs_t) error {
				return errors.New("")
			}
			storage := NewStorage("/media/video1")

			free, err := storage.Free(statfs)
			assert.Error(t, err)
			assert.Equal(t, free, uint64(0))
		})
	})
}
