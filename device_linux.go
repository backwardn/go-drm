package drm

import (
	"fmt"
	"strings"
	"os"
)

const drmMajor = 226

func major(dev uint64) uint32 {
	return uint32((dev & 0xfff00) >> 8)
}

func minor(dev uint64) uint32 {
	return uint32((dev & 0xff) | ((dev >> 12) & 0xfff00))
}

func getMinorType(min uint32) (NodeType, error) {
	t := NodeType(min >> 6)
	switch t {
	case NodePrimary, NodeControl, NodeRender:
		return t, nil
	default:
		return 0, fmt.Errorf("drm: unknown node type 0x%X", t)
	}
}

func getPCIPath(dev uint64) string {
	return fmt.Sprintf("/sys/dev/char/%d:%d/device", major(dev), minor(dev))
}

func getSubsystemType(dev uint64) (BusType, error) {
	subsystemPath := getPCIPath(dev) + "/subsystem"
	subsystemTarget, err := os.Readlink(subsystemPath)
	if err != nil {
		return 0, err
	}

	switch {
	case strings.HasSuffix(subsystemTarget, "/pci"):
		return BusPCI, nil
	case strings.HasSuffix(subsystemTarget, "/usb"):
		return BusUSB, nil
	case strings.HasSuffix(subsystemTarget, "/platform"):
		return BusPlatform, nil
	case strings.HasSuffix(subsystemTarget, "/spi"):
		return BusPlatform, nil
	case strings.HasSuffix(subsystemTarget, "/host1x"):
		return BusHost1x, nil
	case strings.HasSuffix(subsystemTarget, "/virtio"):
		return BusType(0x10), nil
	default:
		return 0, fmt.Errorf("drm: failed to get subsystem type from path: %v", subsystemTarget)
	}
}

func getPCIDevice(devID uint64) (*PCIDevice, error) {
	p := getPCIPath(devID)

	var dev PCIDevice
	props := []struct{
		name string
		dst *uint32
	}{
		{"vendor", &dev.Vendor},
		{"device", &dev.Device},
		{"subsystem_vendor", &dev.SubVendor},
		{"subsystem_device", &dev.SubDevice},
	}
	for _, prop := range props {
		f, err := os.Open(p + "/" + prop.name)
		if err != nil {
			return nil, err
		}

		_, err = fmt.Fscanf(f, "0x%x", prop.dst)
		f.Close()
		if err != nil {
			return nil, err
		}
	}

	return &dev, nil
}
