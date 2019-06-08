package drm

//go:generate ./fourcc.py

type Format uint32

type ModifierVendor uint8

type Modifier uint64

func (mod Modifier) Vendor() ModifierVendor {
	return ModifierVendor(mod >> 56)
}
