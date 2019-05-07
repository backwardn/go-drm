//+build linux dragonfly freebsd netbsd openbsd
//+build amd64 arm64

package drm

import (
	"syscall"
	"unsafe"
)

const (
	ioctlVersion = 0xC0406400
)

type versionResp struct {
	major, minor, patch int32
	nameLen uint64
	name *byte
	dateLen uint64
	date *byte
	descLen uint64
	desc *byte
}

func ioctl(fd uintptr, nr int, ptr unsafe.Pointer) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(nr), uintptr(ptr))
	if errno != 0 {
		return errno
	}
	return nil
}

func alloc(ptr **byte, len uint64) []byte {
	b := make([]byte, len)
	*ptr = (*byte)(unsafe.Pointer(&b[0]))
	return b
}

func version(fd uintptr, v *versionResp) error {
	return ioctl(fd, ioctlVersion, unsafe.Pointer(v))
}

type Device struct {
	fd uintptr
}

func NewDevice(fd uintptr) *Device {
	return &Device{fd}
}

type VersionResp struct {
	Major, Minor, Patch int32
	Name, Date, Desc string
}

func (d *Device) Version() (*VersionResp, error) {
	var v versionResp
	if err := version(d.fd, &v); err != nil {
		return nil, err
	}

	name := alloc(&v.name, v.nameLen)
	date := alloc(&v.date, v.dateLen)
	desc := alloc(&v.desc, v.descLen)

	if err := version(d.fd, &v); err != nil {
		return nil, err
	}

	return &VersionResp{
		Major: v.major,
		Minor: v.minor,
		Patch: v.patch,
		Name: string(name),
		Date: string(date),
		Desc: string(desc),
	}, nil
}
