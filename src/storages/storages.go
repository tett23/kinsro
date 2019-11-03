package storages

import (
	"sort"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/syscalls"
)

// Storages Storages
type Storages []Storage

// NewStoragesFromPaths NewStoragesFromPaths
func NewStoragesFromPaths(paths []string) Storages {
	ret := make(Storages, len(paths))
	for i := range paths {
		item := NewStorage(paths[i])

		ret[i] = *item
	}

	return ret
}

// FindByName FindByName
func (storages Storages) FindByName(name string) (*Storage, bool) {
	for i := range storages {
		if storages[i].Name() == name {
			return &storages[i], true
		}
	}

	return nil, false
}

// MostSpacefulStorage MostSpacefulStorage
func (storages Storages) MostSpacefulStorage(statfs syscalls.Statfs) (*Storage, error) {
	var stats []struct {
		*Storage
		uint64
	}
	for i := range storages {
		free, err := storages[i].Free(statfs)
		if err != nil {
			continue
		}

		stats = append(stats, struct {
			*Storage
			uint64
		}{
			&storages[i],
			free,
		})
	}

	if len(stats) == 0 {
		return nil, errors.Errorf("no storages.")
	}

	sort.Slice(stats, func(a, b int) bool {
		return stats[a].uint64 > stats[b].uint64
	})

	return stats[0].Storage, nil
}
