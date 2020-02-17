package sysstat

import (
	"strings"
	"syscall"

	"github.com/mackerelio/go-osstat/memory"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/syscalls"
)

// SysStat SysStat
type SysStat struct {
	// 時刻とサーバー名いるかな？
	Memory *MemoryStat
	Disks  []DiskStat
}

// MemoryStat MemoryStat
type MemoryStat struct {
	Total uint64
	Used  uint64
	Cache uint64
	Free  uint64
}

// DiskStat DiskStat
type DiskStat struct {
	Path  string
	Total uint64
	Free  uint64
}

// NewSysStat NewSysStat
func NewSysStat(statfs syscalls.Statfs, fs afero.Fs) (*SysStat, error) {
	memStat, err := NewMemoryStat()
	if err != nil {
		return nil, err
	}

	diskStat, err := NewDiskStat(statfs, fs)
	if err != nil {
		return nil, err
	}

	ret := SysStat{
		Memory: memStat,
		Disks:  diskStat,
	}

	return &ret, nil
}

// NewMemoryStat NewMemoryStat
func NewMemoryStat() (*MemoryStat, error) {
	memory, err := memory.Get()
	if err != nil {
		return nil, errors.Wrapf(err, "get memory stat failed.")
	}

	ret := MemoryStat{
		Total: memory.Total,
		Used:  memory.Used,
		Cache: memory.Cached,
		Free:  memory.Free,
	}

	return &ret, nil
}

// NewDiskStat NewDiskStat
func NewDiskStat(statfs syscalls.Statfs, fs afero.Fs) ([]DiskStat, error) {
	data, err := afero.ReadFile(fs, "/proc/mounts")
	if err != nil {
		return nil, errors.Wrapf(err, "read /proc/mounts faild")
	}
	str := string(data)
	lines := strings.Split(strings.TrimSpace(str), "\n")

	var diskStats []DiskStat

	for _, line := range lines {
		fields := strings.Fields(line)
		path := fields[1]
		fsName := fields[2]
		if !(fsName == "ext4" || fsName == "xfs" || fsName == "nfs4") {
			continue
		}

		var dstat syscall.Statfs_t
		err := statfs(path, &dstat)
		if err != nil {
			continue
		}

		ds := DiskStat{
			Path:  path,
			Total: dstat.Blocks * uint64(dstat.Bsize),
			Free:  dstat.Bfree * uint64(dstat.Bsize),
		}

		diskStats = append(diskStats, ds)
	}

	return diskStats, nil
}
