package drm

import (
	"fmt"
)

type ObjectType uint32

const (
	ObjectAny       ObjectType = 0
	ObjectCRTC      ObjectType = 0xCCCCCCCC
	ObjectConnector ObjectType = 0xC0C0C0C0
	ObjectEncoder   ObjectType = 0xE0E0E0E0
	ObjectMode      ObjectType = 0xDEDEDEDE
	ObjectProperty  ObjectType = 0xB0B0B0B0
	ObjectFB        ObjectType = 0xFBFBFBFB
	ObjectBlob      ObjectType = 0xBBBBBBBB
	ObjectPlane     ObjectType = 0xEEEEEEEE
)

func (t ObjectType) String() string {
	switch t {
	case ObjectAny:
		return "any"
	case ObjectCRTC:
		return "CRTC"
	case ObjectConnector:
		return "connector"
	case ObjectEncoder:
		return "encoder"
	case ObjectMode:
		return "mode"
	case ObjectProperty:
		return "property"
	case ObjectFB:
		return "FB"
	case ObjectBlob:
		return "blob"
	case ObjectPlane:
		return "plane"
	default:
		return "unknown"
	}
}

type (
	ObjectID    uint32
	CRTCID      ObjectID
	ConnectorID ObjectID
	EncoderID   ObjectID
	ModeID      ObjectID
	PropertyID  ObjectID
	FBID        ObjectID
	BlobID      ObjectID
	PlaneID     ObjectID
)

type AnyID interface {
	Type() ObjectType
	Object() ObjectID
}

func NewAnyID(id ObjectID, t ObjectType) AnyID {
	switch t {
	case ObjectAny:
		return id
	case ObjectCRTC:
		return CRTCID(id)
	case ObjectConnector:
		return ConnectorID(id)
	case ObjectEncoder:
		return EncoderID(id)
	case ObjectMode:
		return ModeID(id)
	case ObjectProperty:
		return PropertyID(id)
	case ObjectFB:
		return FBID(id)
	case ObjectBlob:
		return BlobID(id)
	case ObjectPlane:
		return PlaneID(id)
	default:
		panic(fmt.Sprintf("drm: unknown object type %v", t))
	}
}

func (id ObjectID) Type() ObjectType {
	return ObjectAny
}

func (id ObjectID) Object() ObjectID {
	return id
}

func (id CRTCID) Type() ObjectType {
	return ObjectCRTC
}

func (id CRTCID) Object() ObjectID {
	return ObjectID(id)
}

func (id ConnectorID) Type() ObjectType {
	return ObjectConnector
}

func (id ConnectorID) Object() ObjectID {
	return ObjectID(id)
}

func (id EncoderID) Type() ObjectType {
	return ObjectEncoder
}

func (id EncoderID) Object() ObjectID {
	return ObjectID(id)
}

func (id ModeID) Type() ObjectType {
	return ObjectMode
}

func (id ModeID) Object() ObjectID {
	return ObjectID(id)
}

func (id PropertyID) Type() ObjectType {
	return ObjectProperty
}

func (id PropertyID) Object() ObjectID {
	return ObjectID(id)
}

func (id FBID) Type() ObjectType {
	return ObjectFB
}

func (id FBID) Object() ObjectID {
	return ObjectID(id)
}

func (id BlobID) Type() ObjectType {
	return ObjectBlob
}

func (id BlobID) Object() ObjectID {
	return ObjectID(id)
}

func (id PlaneID) Type() ObjectType {
	return ObjectPlane
}

func (id PlaneID) Object() ObjectID {
	return ObjectID(id)
}
