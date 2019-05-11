package drm

import (
	"fmt"
	"syscall"
)

func (n *Node) getDevice() (Device, error) {
	var stat syscall.Stat_t
	if err := syscall.Fstat(int(n.fd), &stat); err != nil {
		return nil, err
	}

	if !isDRM(&stat) {
		return nil, fmt.Errorf("drm: not a DRM device")
	}

	bus, err := getSubsystemType(stat.Rdev)
	if err != nil {
		return nil, err
	}

	switch bus {
	case BusPCI:
		return getPCIDevice(stat.Rdev)
	default:
		return &unknownDevice{bus}, nil
	}
}
