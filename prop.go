package drm

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

const formatModifierCurrentVersion = 1

type formatModifierHeader struct {
	version         uint32
	flags           uint32
	formatsLen      uint32
	formatsOffset   uint32 // bytes
	modifiersLen    uint32
	modifiersOffset uint32 // bytes
}

type formatModifier struct {
	formats uint64
	offset  uint32
	_       uint32

	modifier uint64
}

type FormatModifierSet struct {
	h         formatModifierHeader
	formats   []uint32
	modifiers []formatModifier
}

func ParseFormatModifierSet(b []byte) (*FormatModifierSet, error) {
	if len(b) < int(unsafe.Sizeof(formatModifierHeader{})) {
		return nil, fmt.Errorf("drm: format modifier blob too short")
	}

	hPtr := (*formatModifierHeader)(unsafe.Pointer(&b[0]))
	h := *hPtr
	runtime.KeepAlive(&b)

	if h.version != formatModifierCurrentVersion {
		return nil, fmt.Errorf("drm: unsupported format modifier blob version")
	}

	var formats []uint32
	if h.formatsLen > 0 {
		formats = make([]uint32, h.formatsLen)
		size := len(formats) * int(unsafe.Sizeof(formats[0]))
		if len(b) < int(h.formatsOffset)+size {
			return nil, fmt.Errorf("drm: format modifier blob too short")
		}
		sh := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(&formats[0])),
			Len:  size,
			Cap:  size,
		}
		copy(*(*[]byte)(unsafe.Pointer(&sh)), b[h.formatsOffset:])
	}

	var modifiers []formatModifier
	if h.modifiersLen > 0 {
		modifiers = make([]formatModifier, h.modifiersLen)
		size := len(modifiers) * int(unsafe.Sizeof(modifiers[0]))
		if len(b) < int(h.modifiersOffset)+size {
			return nil, fmt.Errorf("drm: format modifier blob too short")
		}
		sh := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(&modifiers[0])),
			Len:  size,
			Cap:  size,
		}
		copy(*(*[]byte)(unsafe.Pointer(&sh)), b[h.modifiersOffset:])
	}

	return &FormatModifierSet{
		h:         h,
		formats:   formats,
		modifiers: modifiers,
	}, nil
}

func (set *FormatModifierSet) Map() map[Modifier][]Format {
	m := make(map[Modifier][]Format, len(set.modifiers))
	for _, mod := range set.modifiers {
		var formats []Format
		for i := 0; i < 64; i++ {
			if mod.formats&(1<<uint(i)) != 0 {
				fmt := set.formats[int(mod.offset)+i]
				formats = append(formats, Format(fmt))
			}
		}
		m[Modifier(mod.modifier)] = formats
	}
	return m
}

func ParseModeModeInfo(b []byte) (*ModeModeInfo, error) {
	if len(b) < int(unsafe.Sizeof(modeModeInfo{})) {
		return nil, fmt.Errorf("drm: mode info blob too short")
	}

	infoPtr := (*modeModeInfo)(unsafe.Pointer(&b[0]))
	info := newModeModeInfo(infoPtr)
	runtime.KeepAlive(&b)
	return info, nil
}

func ParseFormats(b []byte) ([]Format, error) {
	formatSize := int(unsafe.Sizeof(uint32(0)))
	if len(b)%formatSize != 0 {
		return nil, fmt.Errorf("drm: malformed writeback pixel formats blob")
	}
	formatsLen := len(b) / formatSize

	var formats []Format
	if formatsLen > 0 {
		formats = make([]Format, formatsLen)
		sh := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(&formats[0])),
			Len:  len(b),
			Cap:  len(b),
		}
		copy(*(*[]byte)(unsafe.Pointer(&sh)), b)
	}

	return formats, nil
}

func ParsePath(b []byte) (string, error) {
	return newString(b), nil
}
