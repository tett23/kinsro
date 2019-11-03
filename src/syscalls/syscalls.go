package syscalls

import (
	"os"
	"syscall"
)

// Statfs syscall.Statfs
type Statfs func(path string, stat *syscall.Statfs_t) error

// Symlink os.Symlink
type Symlink func(oldname, newname string) error

// Chdir os.Chdir
type Chdir func(dir string) error

// Syscalls Syscalls
type Syscalls struct {
	Statfs  Statfs
	Symlink Symlink
	Chdir   Chdir
}

// NewOSSyscalls NewOSSyscalls
func NewOSSyscalls() *Syscalls {
	ret := Syscalls{
		Statfs:  syscall.Statfs,
		Symlink: os.Symlink,
		Chdir:   os.Chdir,
	}

	return &ret
}
