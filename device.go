package drm

import (
	"syscall"
)

func modeIsChr(mode uint32) bool {
	return mode&syscall.S_IFMT == syscall.S_IFCHR
}

func devIsDRM(dev uint64) bool {
	return major(dev) == drmMajor
}

func isDRM(stat *syscall.Stat_t) bool {
	return devIsDRM(stat.Rdev) && modeIsChr(stat.Mode)
}
