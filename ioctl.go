//+build linux dragonfly freebsd netbsd openbsd
//+build amd64 arm64

package drm

import (
	"syscall"
	"unsafe"
)

const (
	ioctlVersion = 0xC0406400
	ioctlGetCap = 0xC010640C
)

func ioctl(fd uintptr, nr int, ptr unsafe.Pointer) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(nr), uintptr(ptr))
	if errno != 0 {
		return errno
	}
	return nil
}

func allocBytes(ptr **byte, len uint64) []byte {
	b := make([]byte, len)
	*ptr = (*byte)(unsafe.Pointer(&b[0]))
	return b
}

type versionResp struct {
	major, minor, patch int32
	nameLen uint64
	name *byte
	dateLen uint64
	date *byte
	descLen uint64
	desc *byte
}

func version(fd uintptr, v *versionResp) error {
	return ioctl(fd, ioctlVersion, unsafe.Pointer(v))
}

type getCapArg struct {
	cap uint64
	ret uint64
}

func getCap(fd uintptr, cap uint64) (uint64, error) {
	arg := getCapArg{cap: cap}
	err := ioctl(fd, ioctlGetCap, unsafe.Pointer(&arg))
	return arg.ret, err
}
