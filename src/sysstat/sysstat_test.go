package sysstat

import (
	"syscall"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestNewSysStat(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns storage", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, "/proc/mounts", []byte(procContent), 0777)
			actual, err := NewSysStat(dummyStatfs, fs)
			assert.NoError(t, err)
			assert.NotNil(t, actual)
		})
	})
}

func TestSysStat__DiskStat(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns storage", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, "/proc/mounts", []byte(procContent), 0777)

			actual, err := NewDiskStat(dummyStatfs, fs)
			assert.NoError(t, err)
			assert.NotEmpty(t, actual)
		})
	})
}

func dummyStatfs(path string, stat *syscall.Statfs_t) error {
	stat.Blocks = 2048
	stat.Bsize = 4086
	stat.Bfree = uint64(1000)

	return nil
}

const procContent = `
/dev/root / ext4 rw,noatime 0 0
devtmpfs /dev devtmpfs rw,relatime,size=1867796k,nr_inodes=117771,mode=755 0 0
sysfs /sys sysfs rw,nosuid,nodev,noexec,relatime 0 0
proc /proc proc rw,relatime 0 0
securityfs /sys/kernel/security securityfs rw,nosuid,nodev,noexec,relatime 0 0
tmpfs /dev/shm tmpfs rw,nosuid,nodev 0 0
devpts /dev/pts devpts rw,nosuid,noexec,relatime,gid=5,mode=620,ptmxmode=000 0 0
tmpfs /run tmpfs rw,nosuid,nodev,mode=755 0 0
tmpfs /run/lock tmpfs rw,nosuid,nodev,noexec,relatime,size=5120k 0 0
tmpfs /sys/fs/cgroup tmpfs ro,nosuid,nodev,noexec,mode=755 0 0
cgroup2 /sys/fs/cgroup/unified cgroup2 rw,nosuid,nodev,noexec,relatime,nsdelegate 0 0
cgroup /sys/fs/cgroup/systemd cgroup rw,nosuid,nodev,noexec,relatime,xattr,name=systemd 0 0
cgroup /sys/fs/cgroup/cpu,cpuacct cgroup rw,nosuid,nodev,noexec,relatime,cpu,cpuacct 0 0
cgroup /sys/fs/cgroup/pids cgroup rw,nosuid,nodev,noexec,relatime,pids 0 0
cgroup /sys/fs/cgroup/devices cgroup rw,nosuid,nodev,noexec,relatime,devices 0 0
cgroup /sys/fs/cgroup/net_cls cgroup rw,nosuid,nodev,noexec,relatime,net_cls 0 0
cgroup /sys/fs/cgroup/memory cgroup rw,nosuid,nodev,noexec,relatime,memory 0 0
cgroup /sys/fs/cgroup/blkio cgroup rw,nosuid,nodev,noexec,relatime,blkio 0 0
cgroup /sys/fs/cgroup/freezer cgroup rw,nosuid,nodev,noexec,relatime,freezer 0 0
cgroup /sys/fs/cgroup/cpuset cgroup rw,nosuid,nodev,noexec,relatime,cpuset 0 0
sunrpc /run/rpc_pipefs rpc_pipefs rw,relatime 0 0
mqueue /dev/mqueue mqueue rw,relatime 0 0
debugfs /sys/kernel/debug debugfs rw,relatime 0 0
systemd-1 /proc/sys/fs/binfmt_misc autofs rw,relatime,fd=37,pgrp=1,timeout=0,minproto=5,maxproto=5,direct 0 0
configfs /sys/kernel/config configfs rw,relatime 0 0
/dev/mmcblk0p1 /boot vfat rw,relatime,fmask=0022,dmask=0022,codepage=437,iocharset=ascii,shortname=mixed,errors=remount-ro 0 0
/dev/sda /media/video_tmp ext4 rw,nosuid,nodev,noexec,relatime 0 0
10.0.1.110:/media/video1 /media/video1 nfs4 rw,relatime,vers=4.1,rsize=8192,wsize=8192,namlen=255,hard,proto=tcp,timeo=600,retrans=2,sec=sys,clientaddr=10.0.1.100,local_lock=none,addr=10.0.1.110 0 0
tmpfs /run/user/1000 tmpfs rw,nosuid,nodev,relatime,size=399976k,mode=700,uid=1000,gid=1000 0 0
`
