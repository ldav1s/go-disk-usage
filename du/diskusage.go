// +build !windows

package du

import "syscall"

type DiskUsage struct {
	stat *syscall.Statfs_t
}

// Returns an object holding the disk usage of volumePath
// This function assumes volumePath is a valid path
func NewDiskUsage(volumePath string) (dusage *DiskUsage, err error) {

	var stat syscall.Statfs_t
	err = syscall.Statfs(volumePath, &stat)
	if err != nil {
		return
	}
	dusage = &DiskUsage{&stat}
	return
}

// Total free bytes on file system
func (this *DiskUsage) Free() uint64 {
	return this.stat.Bfree * uint64(this.stat.Bsize)
}

// Total available bytes on file system to an unprivileged user
func (this *DiskUsage) Available() uint64 {
	return this.stat.Bavail * uint64(this.stat.Bsize)
}

// Total size of the file system
func (this *DiskUsage) Size() uint64 {
	return this.stat.Blocks * uint64(this.stat.Bsize)
}

// Total bytes used in file system
func (this *DiskUsage) Used() uint64 {
	return this.Size() - this.Free()
}

// Percentage of use on the file system
func (this *DiskUsage) Usage() float32 {
	return float32(this.Used()) / float32(this.Size())
}
