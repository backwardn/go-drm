package drm

import (
	"unsafe"
)

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

type Node struct {
	fd uintptr
}

func NewNode(fd uintptr) *Node {
	return &Node{fd}
}

type Version struct {
	Major, Minor, Patch int32
	Name, Date, Desc    string
}

func (n *Node) Version() (*Version, error) {
	var v versionResp
	if err := version(n.fd, &v); err != nil {
		return nil, err
	}

	name := allocBytes(&v.name, v.nameLen)
	date := allocBytes(&v.date, v.dateLen)
	desc := allocBytes(&v.desc, v.descLen)

	if err := version(n.fd, &v); err != nil {
		return nil, err
	}

	return &Version{
		Major: v.major,
		Minor: v.minor,
		Patch: v.patch,
		Name:  string(name),
		Date:  string(date),
		Desc:  string(desc),
	}, nil
}

type PCIDevice struct {
	Vendor, Device       uint32
	SubVendor, SubDevice uint32
}

type Device interface{}

func (n *Node) GetDevice() (Device, error) {
	return n.getDevice()
}

func (n *Node) GetCap(cap Cap) (uint64, error) {
	return getCap(n.fd, uint64(cap))
}

func (n *Node) SetClientCap(cap ClientCap, val uint64) error {
	return setClientCap(n.fd, uint64(cap), val)
}

type ModeCard struct {
	FBs                                      []FBID
	CRTCs                                    []CRTCID
	Connectors                               []ConnectorID
	Encoders                                 []EncoderID
	MinWidth, MaxWidth, MinHeight, MaxHeight uint32
}

func (n *Node) ModeGetResources() (*ModeCard, error) {
	for {
		var r modeCardResp
		if err := modeGetResources(n.fd, &r); err != nil {
			return nil, err
		}
		count := r

		var fbs []FBID
		var crtcs []CRTCID
		var connectors []ConnectorID
		var encoders []EncoderID
		if r.fbsLen > 0 {
			fbs = make([]FBID, r.fbsLen)
			r.fbs = (*uint32)(unsafe.Pointer(&fbs[0]))
		}
		if r.crtcsLen > 0 {
			crtcs = make([]CRTCID, r.crtcsLen)
			r.crtcs = (*uint32)(unsafe.Pointer(&crtcs[0]))
		}
		if r.connectorsLen > 0 {
			connectors = make([]ConnectorID, r.connectorsLen)
			r.connectors = (*uint32)(unsafe.Pointer(&connectors[0]))
		}
		if r.encodersLen > 0 {
			encoders = make([]EncoderID, r.encodersLen)
			r.encoders = (*uint32)(unsafe.Pointer(&encoders[0]))
		}

		if err := modeGetResources(n.fd, &r); err != nil {
			return nil, err
		}

		if r.fbsLen != count.fbsLen || r.crtcsLen != count.crtcsLen || r.connectorsLen != count.connectorsLen || r.encodersLen != count.encodersLen {
			continue
		}

		return &ModeCard{
			FBs:        fbs,
			CRTCs:      crtcs,
			Connectors: connectors,
			Encoders:   encoders,
			MinWidth:   r.minWidth,
			MaxWidth:   r.maxWidth,
			MinHeight:  r.minHeight,
			MaxHeight:  r.maxHeight,
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
	Clock                                         uint32
	HDisplay, HSyncStart, HSyncEnd, HTotal, HSkew uint16
	VDisplay, VSyncStart, VSyncEnd, VTotal, VScan uint16

	VRefresh uint32

	Flags uint32
	Type  uint32
	Name  string
}

func newModeModeInfo(info *modeModeInfo) *ModeModeInfo {
	return &ModeModeInfo{
		Clock:      info.clock,
		HDisplay:   info.hDisplay,
		HSyncStart: info.hSyncStart,
		HSyncEnd:   info.hSyncEnd,
		HTotal:     info.hTotal,
		HSkew:      info.hSkew,
		VDisplay:   info.vDisplay,
		VSyncStart: info.vSyncStart,
		VSyncEnd:   info.vSyncEnd,
		VTotal:     info.vTotal,
		VScan:      info.vScan,
		VRefresh:   info.vRefresh,
		Flags:      info.flags,
		Type:       info.typ,
		Name:       newString(info.name[:]),
	}
}

func newModeModeInfoList(infos []modeModeInfo) []ModeModeInfo {
	l := make([]ModeModeInfo, len(infos))
	for i, info := range infos {
		l[i] = *newModeModeInfo(&info)
	}
	return l
}

type ModeCRTC struct {
	ID        CRTCID
	FB        FBID
	X, Y      uint32
	GammaSize uint32
	Mode      *ModeModeInfo
}

func (n *Node) ModeGetCRTC(id CRTCID) (*ModeCRTC, error) {
	r := modeCRTCResp{id: uint32(id)}
	if err := modeGetCRTC(n.fd, &r); err != nil {
		return nil, err
	}

	var mode *ModeModeInfo
	if r.modeValid != 0 {
		mode = newModeModeInfo(&r.mode)
	}

	return &ModeCRTC{
		ID:        CRTCID(r.id),
		FB:        FBID(r.fb),
		X:         r.x,
		Y:         r.y,
		GammaSize: r.gammaSize,
		Mode:      mode,
	}, nil
}

type ModeEncoder struct {
	ID                            EncoderID
	Type                          EncoderType
	CRTC                          CRTCID
	PossibleCRTCs, PossibleClones uint32
}

func (n *Node) ModeGetEncoder(id EncoderID) (*ModeEncoder, error) {
	r := modeEncoderResp{id: uint32(id)}
	if err := modeGetEncoder(n.fd, &r); err != nil {
		return nil, err
	}

	return &ModeEncoder{
		ID:             EncoderID(r.id),
		Type:           EncoderType(r.typ),
		CRTC:           CRTCID(r.crtc),
		PossibleCRTCs:  r.possibleCRTCs,
		PossibleClones: r.possibleClones,
	}, nil
}

type ModeConnector struct {
	PossibleEncoders []EncoderID
	Modes            []ModeModeInfo
	PropIDs          []PropertyID
	PropValues       []uint64

	Encoder ObjectID
	ID      ObjectID
	Type    ConnectorType

	Status              ConnectorStatus
	PhyWidth, PhyHeight uint32 // mm
	Subpixel            Subpixel
}

func (n *Node) ModeGetConnector(id ConnectorID) (*ModeConnector, error) {
	for {
		r := modeConnectorResp{id: uint32(id)}
		if err := modeGetConnector(n.fd, &r); err != nil {
			return nil, err
		}
		count := r

		var encoders []EncoderID
		var propIDs []PropertyID
		var modes []modeModeInfo
		var propValues []uint64
		if r.propsLen > 0 {
			propIDs = make([]PropertyID, r.propsLen)
			r.propIDs = (*uint32)(unsafe.Pointer(&propIDs[0]))
			propValues = make([]uint64, r.propsLen)
			r.propValues = (*uint64)(unsafe.Pointer(&propValues[0]))
		}
		if r.modesLen > 0 {
			modes = make([]modeModeInfo, r.modesLen)
			r.modes = (*modeModeInfo)(unsafe.Pointer(&modes[0]))
		}
		if r.encodersLen > 0 {
			encoders = make([]EncoderID, r.encodersLen)
			r.encoders = (*uint32)(unsafe.Pointer(&encoders[0]))
		}

		if err := modeGetConnector(n.fd, &r); err != nil {
			return nil, err
		}

		if r.propsLen != count.propsLen || r.modesLen != count.modesLen || r.encodersLen != count.encodersLen {
			continue
		}

		return &ModeConnector{
			PossibleEncoders: encoders,
			Modes:            newModeModeInfoList(modes),
			PropIDs:          propIDs,
			PropValues:       propValues,
			Encoder:          ObjectID(r.encoder),
			ID:               ObjectID(r.id),
			Type:             ConnectorType(r.typ),
			Status:           ConnectorStatus(r.status),
			PhyWidth:         r.phyWidth,
			PhyHeight:        r.phyHeight,
			Subpixel:         Subpixel(r.subpixel),
		}, nil
	}
}

func (n *Node) ModeGetPlaneResources() ([]PlaneID, error) {
	for {
		var r modePlaneResourcesResp
		if err := modeGetPlaneResources(n.fd, &r); err != nil {
			return nil, err
		}
		count := r

		var planes []PlaneID
		if r.planesLen > 0 {
			planes = make([]PlaneID, r.planesLen)
			r.planes = (*uint32)(unsafe.Pointer(&planes[0]))
		}

		if err := modeGetPlaneResources(n.fd, &r); err != nil {
			return nil, err
		}

		if r.planesLen != count.planesLen {
			continue
		}

		return planes, nil
	}
}

type ModePlane struct {
	ID PlaneID

	CRTC CRTCID
	FB   FBID

	PossibleCRTCs uint32
	GammaSize     uint32

	Formats []Format
}

func (n *Node) ModeGetPlane(id PlaneID) (*ModePlane, error) {
	for {
		r := modePlaneResp{id: uint32(id)}
		if err := modeGetPlane(n.fd, &r); err != nil {
			return nil, err
		}
		count := r

		var formats []Format
		if r.formatsLen > 0 {
			formats = make([]Format, r.formatsLen)
			r.formats = (*uint32)(unsafe.Pointer(&formats[0]))
		}

		if err := modeGetPlane(n.fd, &r); err != nil {
			return nil, err
		}

		if r.formatsLen != count.formatsLen {
			continue
		}

		return &ModePlane{
			ID:            PlaneID(r.id),
			CRTC:          CRTCID(r.crtc),
			FB:            FBID(r.fb),
			PossibleCRTCs: r.possibleCRTCs,
			GammaSize:     r.gammaSize,
			Formats:       formats,
		}, nil
	}
}
