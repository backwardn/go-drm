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

type Version struct {
	Major, Minor, Patch int32
	Name, Date, Desc string
}

func (d *Node) Version() (*Version, error) {
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

	return &Version{
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

type ModeCard struct {
	FBs, CRTCs, Connectors, Encoders []ObjectID
	MinWidth, MaxWidth, MinHeight, MaxHeight uint32
}

func (d *Node) ModeGetResources() (*ModeCard, error) {
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

		return &ModeCard{
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

func newString(b []byte) string {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return string(b[:i])
		}
	}
	return string(b)
}

type ModeModeInfo struct {
	Clock uint32
	HDisplay, HSyncStart, HSyncEnd, HTotal, HSkew uint16
	VDisplay, VSyncStart, VSyncEnd, VTotal, VScan uint16

	VRefresh uint32

	Flags uint32
	Type uint32
	Name string
}

func newModeModeInfo(info *modeModeInfo) *ModeModeInfo {
	return &ModeModeInfo{
		Clock: info.clock,
		HDisplay: info.hDisplay,
		HSyncStart: info.hSyncStart,
		HSyncEnd: info.hSyncEnd,
		HTotal: info.hTotal,
		HSkew: info.hSkew,
		VDisplay: info.vDisplay,
		VSyncStart: info.vSyncStart,
		VSyncEnd: info.vSyncEnd,
		VTotal: info.vTotal,
		VScan: info.vScan,
		VRefresh: info.vRefresh,
		Flags: info.flags,
		Type: info.typ,
		Name: newString(info.name[:]),
	}
}

type ModeCRTC struct {
	ID ObjectID
	FB ObjectID
	X, Y uint32
	GammaSize uint32
	Mode *ModeModeInfo
}

func (d *Node) ModeGetCRTC(id ObjectID) (*ModeCRTC, error) {
	r := modeCRTCResp{id: uint32(id)}
	if err := modeGetCRTC(d.fd, &r); err != nil {
		return nil, err
	}

	var mode *ModeModeInfo
	if r.modeValid != 0 {
		mode = newModeModeInfo(&r.mode)
	}

	return &ModeCRTC{
		ID: ObjectID(r.id),
		FB: ObjectID(r.fb),
		X: r.x,
		Y: r.y,
		GammaSize: r.gammaSize,
		Mode: mode,
	}, nil
}
