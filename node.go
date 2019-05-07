package drm

import (
	"unsafe"
)

type ObjectID uint32

type Node struct {
	fd uintptr
}

func NewNode(fd uintptr) *Node {
	return &Node{fd}
}

type VersionResp struct {
	Major, Minor, Patch int32
	Name, Date, Desc string
}

func (d *Node) Version() (*VersionResp, error) {
	var v versionResp
	if err := version(d.fd, &v); err != nil {
		return nil, err
	}

	name := allocBytes(&v.name, v.nameLen)
	date := allocBytes(&v.date, v.dateLen)
	desc := allocBytes(&v.desc, v.descLen)

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

func (d *Node) GetCap(cap Cap) (uint64, error) {
	return getCap(d.fd, uint64(cap))
}

type ModeCardResp struct {
	FBs, CRTCs, Connectors, Encoders []ObjectID
	MinWidth, MaxWidth, MinHeight, MaxHeight uint32
}

func (d *Node) ModeGetResources() (*ModeCardResp, error) {
	for {
		var r modeCardResp
		if err := modeGetResources(d.fd, &r); err != nil {
			return nil, err
		}
		count := r

		var fbs, crtcs, connectors, encoders []ObjectID
		if r.fbsLen > 0 {
			fbs = make([]ObjectID, r.fbsLen)
			r.fbs = (*uint32)(unsafe.Pointer(&fbs[0]))
		}
		if r.crtcsLen > 0 {
			crtcs = make([]ObjectID, r.crtcsLen)
			r.crtcs = (*uint32)(unsafe.Pointer(&crtcs[0]))
		}
		if r.connectorsLen > 0 {
			connectors = make([]ObjectID, r.connectorsLen)
			r.connectors = (*uint32)(unsafe.Pointer(&connectors[0]))
		}
		if r.encodersLen > 0 {
			encoders = make([]ObjectID, r.encodersLen)
			r.encoders = (*uint32)(unsafe.Pointer(&encoders[0]))
		}

		if err := modeGetResources(d.fd, &r); err != nil {
			return nil, err
		}

		if r.fbsLen != count.fbsLen || r.crtcsLen != count.crtcsLen || r.connectorsLen != count.connectorsLen || r.encodersLen != count.encodersLen {
			continue
		}

		return &ModeCardResp{
			FBs: fbs,
			CRTCs: crtcs,
			Connectors: connectors,
			Encoders: encoders,
			MinWidth: r.minWidth,
			MaxWidth: r.maxWidth,
			MinHeight: r.minHeight,
			MaxHeight: r.maxHeight,
		}, nil
	}
}
