package drm

const (
	propertyPending   uint32 = 1 << 0 // deprecated
	propertyImmutable uint32 = 1 << 2
	propertyAtomic    uint32 = 1 << 31

	propertyLegacyTypeMask   uint32 = 0x0000003A
	propertyExtendedTypeMask uint32 = 0x0000FFC0
)

type PropertyType uint32

func newPropertyType(flags uint32) PropertyType {
	return PropertyType(flags & (propertyLegacyTypeMask | propertyExtendedTypeMask))
}

const (
	// Legacy types
	PropertyRange   PropertyType = 1 << 1
	PropertyEnum    PropertyType = 1 << 3
	PropertyBlob    PropertyType = 1 << 4
	PropertyBitmask PropertyType = 1 << 5

	// Extended types
	PropertyObject      PropertyType = 1 << 6
	PropertySignedRange PropertyType = 2 << 6
)

func (t PropertyType) String() string {
	switch t {
	case PropertyRange:
		return "range"
	case PropertyEnum:
		return "enum"
	case PropertyBlob:
		return "blob"
	case PropertyBitmask:
		return "bitmask"
	case PropertyObject:
		return "object"
	case PropertySignedRange:
		return "signed range"
	default:
		return "unknown"
	}
}

type ModePropertyEnum struct {
	Name  string
	Value uint64
}

func newModePropertyEnum(e modePropertyEnum) ModePropertyEnum {
	return ModePropertyEnum{
		Name:  newString(e.name[:]),
		Value: e.value,
	}
}

func newModePropertyEnumList(enums []modePropertyEnum) []ModePropertyEnum {
	l := make([]ModePropertyEnum, len(enums))
	for i, e := range enums {
		l[i] = newModePropertyEnum(e)
	}
	return l
}

type ModePropertyBlob struct {
	ID   BlobID
	Size uint32
}

func newModePropertyBlobList(blobIDs []BlobID, blobSizes []uint32) []ModePropertyBlob {
	if len(blobSizes) != len(blobIDs) {
		panic("drm: blob sizes and IDs length mismatch")
	}
	l := make([]ModePropertyBlob, len(blobSizes))
	for i := 0; i < len(blobIDs); i++ {
		l[i] = ModePropertyBlob{
			ID:   blobIDs[i],
			Size: blobSizes[i],
		}
	}
	return l
}

type ModeProperty struct {
	ID   PropertyID
	Name string

	flags  uint32
	values []uint64
	enums  []ModePropertyEnum
	blobs  []ModePropertyBlob
}

func (prop *ModeProperty) Type() PropertyType {
	return newPropertyType(prop.flags)
}

func (prop *ModeProperty) Immutable() bool {
	return prop.flags&propertyImmutable != 0
}

func (prop *ModeProperty) Atomic() bool {
	return prop.flags&propertyAtomic != 0
}

func (prop *ModeProperty) Range() (low, high uint64, ok bool) {
	if prop.Type() != PropertyRange || len(prop.values) != 2 {
		return 0, 0, false
	}
	return prop.values[0], prop.values[1], true
}

func (prop *ModeProperty) Enums() ([]ModePropertyEnum, bool) {
	switch prop.Type() {
	case PropertyEnum, PropertyBitmask:
		return prop.enums, true
	default:
		return nil, false
	}
}

func (prop *ModeProperty) Blobs() ([]ModePropertyBlob, bool) {
	switch prop.Type() {
	case PropertyBlob:
		return prop.blobs, true
	default:
		return nil, false
	}
}

func (prop *ModeProperty) ObjectType() (ObjectType, bool) {
	if prop.Type() != PropertyObject || len(prop.values) != 1 {
		return 0, false
	}
	return ObjectType(prop.values[0]), true
}

func (prop *ModeProperty) SignedRange() (low, high int64, ok bool) {
	if prop.Type() != PropertyRange || len(prop.values) != 2 {
		return 0, 0, false
	}
	return int64(prop.values[0]), int64(prop.values[1]), true
}
