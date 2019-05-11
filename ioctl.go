//+build linux dragonfly freebsd netbsd openbsd
//+build amd64 arm64

package drm

import (
	"syscall"
	"unsafe"
)

const (
	ioctlVersion      = 0xC0406400
	ioctlGetCap       = 0xC010640C
	ioctlSetClientCap = 0x4010640D

	ioctlModeGetResources        = 0xC04064A0
	ioctlModeGetCRTC             = 0xC06864A1
	ioctlModeGetEncoder          = 0xC01464A6
	ioctlModeGetConnector        = 0xC05064A7
	ioctlModeGetPlaneResources   = 0xC01064B5
	ioctlModeGetPlane            = 0xC02064B6
	ioctlModeObjectGetProperties = 0xC02064B9
	ioctlModeGetProperty         = 0xC04064AA
	ioctlModeGetBlob             = 0xC01064AC
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
	nameLen             uint64
	name                *byte
	dateLen             uint64
	date                *byte
	descLen             uint64
	desc                *byte
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

type setCapArg struct {
	cap uint64
	val uint64
}

func setClientCap(fd uintptr, cap, val uint64) error {
	arg := setCapArg{cap: cap, val: val}
	return ioctl(fd, ioctlSetClientCap, unsafe.Pointer(&arg))
}

type modeCardResp struct {
	fbs, crtcs, connectors, encoders             *uint32
	fbsLen, crtcsLen, connectorsLen, encodersLen uint32
	minWidth, maxWidth, minHeight, maxHeight     uint32
}

func modeGetResources(fd uintptr, r *modeCardResp) error {
	return ioctl(fd, ioctlModeGetResources, unsafe.Pointer(r))
}

type modeModeInfo struct {
	clock                                         uint32
	hDisplay, hSyncStart, hSyncEnd, hTotal, hSkew uint16
	vDisplay, vSyncStart, vSyncEnd, vTotal, vScan uint16

	vRefresh uint32

	flags uint32
	typ   uint32
	name  [32]byte
}

type modeCRTCResp struct {
	// For ioctlModeSetCRTC
	setConnectors    *uint32
	setConnectorsLen uint32

	id uint32
	fb uint32

	x, y uint32

	gammaSize uint32
	modeValid uint32
	mode      modeModeInfo
}

func modeGetCRTC(fd uintptr, r *modeCRTCResp) error {
	return ioctl(fd, ioctlModeGetCRTC, unsafe.Pointer(r))
}

type modeEncoderResp struct {
	id  uint32
	typ uint32

	crtc uint32

	possibleCRTCs, possibleClones uint32
}

func modeGetEncoder(fd uintptr, r *modeEncoderResp) error {
	return ioctl(fd, ioctlModeGetEncoder, unsafe.Pointer(r))
}

type modeConnectorResp struct {
	encoders   *uint32
	modes      *modeModeInfo
	propIDs    *uint32
	propValues *uint64

	modesLen    uint32
	propsLen    uint32
	encodersLen uint32

	encoder uint32
	id      uint32
	typ     uint32
	typeID  uint32

	status              uint32
	phyWidth, phyHeight uint32
	subpixel            uint32

	_ uint32
}

func modeGetConnector(fd uintptr, r *modeConnectorResp) error {
	return ioctl(fd, ioctlModeGetConnector, unsafe.Pointer(r))
}

type modePlaneResourcesResp struct {
	planes    *uint32
	planesLen uint32
}

func modeGetPlaneResources(fd uintptr, r *modePlaneResourcesResp) error {
	return ioctl(fd, ioctlModeGetPlaneResources, unsafe.Pointer(r))
}

type modePlaneResp struct {
	id uint32

	crtc uint32
	fb   uint32

	possibleCRTCs uint32
	gammaSize     uint32

	formatsLen uint32
	formats    *uint32
}

func modeGetPlane(fd uintptr, r *modePlaneResp) error {
	return ioctl(fd, ioctlModeGetPlane, unsafe.Pointer(r))
}

type modeObjectGetPropertiesResp struct {
	propIDs    *uint32
	propValues *uint64
	propsLen   uint32

	id  uint32
	typ uint32
}

func modeObjectGetProperties(fd uintptr, r *modeObjectGetPropertiesResp) error {
	return ioctl(fd, ioctlModeObjectGetProperties, unsafe.Pointer(r))
}

type modePropertyEnum struct {
	value uint64
	name  [32]byte
}

type modeGetPropertyResp struct {
	values    uintptr // values and blob lengths
	enumBlobs uintptr // enum items and blob IDs

	id    uint32
	flags uint32
	name  [32]byte

	valuesLen, enumBlobsLen uint32
}

func modeGetProperty(fd uintptr, r *modeGetPropertyResp) error {
	return ioctl(fd, ioctlModeGetProperty, unsafe.Pointer(r))
}

type modeGetBlobResp struct {
	id   uint32
	size uint32
	data *byte
}

func modeGetBlob(fd uintptr, r *modeGetBlobResp) error {
	return ioctl(fd, ioctlModeGetBlob, unsafe.Pointer(r))
}
