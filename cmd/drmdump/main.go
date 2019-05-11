package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"git.sr.ht/~emersion/go-drm"
)

type DriverVersion struct {
	Major int32  `json:"major"`
	Minor int32  `json:"minor"`
	Patch int32  `json:"patch"`
	Date  string `json:"date"`
}

type Driver struct {
	Name       string             `json:"name"`
	Desc       string             `json:"desc"`
	Version    DriverVersion      `json:"version"`
	Caps       map[string]*uint64 `json:"caps"`
	ClientCaps map[string]bool    `json:"client_caps"`
}

var capNames = map[drm.Cap]string{
	drm.CapDumbBuffer:          "DUMB_BUFFER",
	drm.CapVblankHighCRTC:      "VBLANK_HIGH_CRTC",
	drm.CapDumbPreferredDepth:  "DUMB_PREFERRED_DEPTH",
	drm.CapDumbPreferredShadow: "DUMB_PREFER_SHADOW",
	drm.CapPrime:               "PRIME",
	drm.CapTimestampMonotonic:  "TIMESTAMP_MONOTONIC",
	drm.CapAsyncPageFlip:       "ASYNC_PAGE_FLIP",
	drm.CapCursorWidth:         "CURSOR_WIDTH",
	drm.CapCursorHeight:        "CURSOR_HEIGHT",
	drm.CapAddFB2Modifiers:     "ADDFB2_MODIFIERS",
	drm.CapPageFlipTarget:      "PAGE_FLIP_TARGET",
	drm.CapCRTCInVBlankEvent:   "CRTC_IN_VBLANK_EVENT",
	drm.CapSyncObj:             "SYNCOBJ",
}

var clientCapNames = map[drm.ClientCap]string{
	drm.ClientCapStereo3D:            "STEREO_3D",
	drm.ClientCapUniversalPlanes:     "UNIVERSAL_PLANES",
	drm.ClientCapAtomic:              "ATOMIC",
	drm.ClientCapAspectRatio:         "APSECT_RATIO",
	drm.ClientCapWritebackConnectors: "WRITEBACK_CONNECTORS",
}

func driver(n *drm.Node) (*Driver, error) {
	v, err := n.Version()
	if err != nil {
		return nil, fmt.Errorf("cannot get version: %v", err)
	}

	caps := make(map[string]*uint64)
	for c, s := range capNames {
		var ptr *uint64
		val, err := n.GetCap(c)
		if err == nil {
			ptr = &val
		} else if err != syscall.EINVAL {
			return nil, fmt.Errorf("failed to get cap: %v", err)
		}
		caps[s] = ptr
	}

	clientCaps := make(map[string]bool)
	for c, s := range clientCapNames {
		ok := false
		if err := n.SetClientCap(c, 1); err == nil {
			ok = true
		} else if err != syscall.EINVAL {
			return nil, fmt.Errorf("failed to set client cap: %v", err)
		}
		clientCaps[s] = ok
	}

	return &Driver{
		Name: v.Name,
		Desc: v.Desc,
		Version: DriverVersion{
			Major: v.Major,
			Minor: v.Minor,
			Patch: v.Patch,
			Date:  v.Date,
		},
		Caps:       caps,
		ClientCaps: clientCaps,
	}, nil
}

type DevicePCI struct {
	Vendor    uint32 `json:"vendor"`
	Device    uint32 `json:"device"`
	SubVendor uint32 `json:"sub_vendor"`
	SubDevice uint32 `json:"sub_device"`
}

type Device struct {
	BusType drm.BusType `json:"bus_type"`
	PCI     *DevicePCI  `json:"pci,omitempty"`
}

func device(n *drm.Node) (*Device, error) {
	dev, err := n.GetDevice()
	if err != nil {
		return nil, fmt.Errorf("failed to get device: %v", err)
	}

	switch dev := dev.(type) {
	case *drm.PCIDevice:
		return &Device{
			BusType: dev.BusType(),
			PCI: &DevicePCI{
				Vendor:    dev.Vendor,
				Device:    dev.Device,
				SubVendor: dev.SubVendor,
				SubDevice: dev.SubDevice,
			},
		}, nil
	default:
		return &Device{BusType: dev.BusType()}, nil
	}
}

type Mode struct {
	Clock      uint32 `json:"clock"`
	HDisplay   uint16 `json:"hdisplay"`
	HSyncStart uint16 `json:"hsync_start"`
	HSyncEnd   uint16 `json:"hsync_end"`
	HTotal     uint16 `json:"htotal"`
	HSkew      uint16 `json:"hskew"`
	VDisplay   uint16 `json:"vdisplay"`
	VSyncStart uint16 `json:"vsync_total"`
	VSyncEnd   uint16 `json:"vsync_end"`
	VTotal     uint16 `json:"vtotal"`
	VScan      uint16 `json:"vscan"`
	VRefresh   uint32 `json:"vrefresh"`
	Flags      uint32 `json:"flags"`
	Type       uint32 `json:"type"`
	Name       string `json:"name"`
}

func modeList(modes []drm.ModeModeInfo) []Mode {
	l := make([]Mode, len(modes))
	for i, m := range modes {
		l[i] = Mode(m)
	}
	return l
}

func parseFixedPoint16(v interface{}) (interface{}, error) {
	return v.(uint64) >> 16, nil
}

func parseInFormats(v interface{}) (interface{}, error) {
	b := v.([]byte)
	set, err := drm.ParseFormatModifierSet(b)
	if err != nil {
		return nil, err
	}
	return set.Map(), nil
}

func parseModeID(v interface{}) (interface{}, error) {
	b := v.([]byte)
	mode, err := drm.ParseModeModeInfo(b)
	if err != nil {
		return nil, err
	}
	return (*Mode)(mode), nil
}

func parseWritebackPixelFormats(v interface{}) (interface{}, error) {
	b := v.([]byte)
	formats, err := drm.ParseFormats(b)
	if err != nil {
		return nil, err
	}
	return formats, nil
}

func parsePath(v interface{}) (interface{}, error) {
	b := v.([]byte)
	return drm.ParsePath(b)
}

var propertyParsers = map[string]struct {
	objectType   drm.ObjectType
	propertyType drm.PropertyType
	f            func(v interface{}) (interface{}, error)
}{
	"SRC_X":                   {drm.ObjectPlane, drm.PropertyRange, parseFixedPoint16},
	"SRC_Y":                   {drm.ObjectPlane, drm.PropertyRange, parseFixedPoint16},
	"SRC_W":                   {drm.ObjectPlane, drm.PropertyRange, parseFixedPoint16},
	"SRC_H":                   {drm.ObjectPlane, drm.PropertyRange, parseFixedPoint16},
	"IN_FORMATS":              {drm.ObjectPlane, drm.PropertyBlob, parseInFormats},
	"MODE_ID":                 {drm.ObjectCRTC, drm.PropertyBlob, parseModeID},
	"WRITEBACK_PIXEL_FORMATS": {drm.ObjectConnector, drm.PropertyBlob, parseWritebackPixelFormats},
	"PATH":                    {drm.ObjectConnector, drm.PropertyBlob, parsePath},
}

type Property struct {
	ID        drm.PropertyID   `json:"id"`
	Type      drm.PropertyType `json:"type"`
	Immutable bool             `json:"immutable"`
	Atomic    bool             `json:"atomic"`
	RawValue  uint64           `json:"raw_value"`
	// Value interpreted with the property type
	Value interface{} `json:"value"`
	// Value interpreted with the property name
	Data interface{} `json:"data,omitempty"`
}

func properties(n *drm.Node, id drm.AnyID) (map[string]Property, error) {
	props, err := n.ModeObjectGetProperties(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get properties for object %v: %v", id, err)
	}

	m := make(map[string]Property, len(props))
	for propID, propValue := range props {
		prop, err := n.ModeGetProperty(propID)
		if err != nil {
			return nil, fmt.Errorf("failed to get property %v: %v", propID, err)
		}

		var val interface{}
		switch prop.Type() {
		case drm.PropertyBlob:
			blobID := drm.BlobID(propValue)
			if blobID == 0 {
				// NULL blob
				val = []byte(nil)
				break
			}
			b, err := n.ModeGetBlob(blobID)
			if err != nil {
				return nil, fmt.Errorf("failed to get blob %v: %v", blobID, err)
			}
			val = b
		case drm.PropertyObject:
			val = drm.ObjectID(propValue)
		case drm.PropertySignedRange:
			val = int64(propValue)
		default:
			val = propValue
		}

		var data interface{}
		if parser, ok := propertyParsers[prop.Name]; ok {
			if parser.objectType != id.Type() {
				log.Printf("Cannot parse property %v: expected object type %v, got %v", prop.Name, parser.objectType, id.Type())
			}
			if parser.propertyType != prop.Type() {
				log.Printf("Cannot parse property %v: expected property type %v, got %v", prop.Name, parser.propertyType, prop.Type())
			}
			data, err = parser.f(val)
			if err != nil {
				log.Printf("Cannot parse property %v: %v", prop.Name, err)
			}
		}

		m[prop.Name] = Property{
			ID:        prop.ID,
			Type:      prop.Type(),
			Immutable: prop.Immutable(),
			Atomic:    prop.Atomic(),
			RawValue:  propValue,
			Value:     val,
			Data:      data,
		}
	}

	return m, nil
}

type Connector struct {
	ID         drm.ConnectorID     `json:"id"`
	Type       drm.ConnectorType   `json:"type"`
	Status     drm.ConnectorStatus `json:"status"`
	PhyWidth   uint32              `json:"phy_width"`
	PhyHeight  uint32              `json:"phy_height"`
	Subpixel   drm.Subpixel        `json:"subpixel"`
	Encoders   []drm.EncoderID     `json:"encoders"`
	Modes      []Mode              `json:"modes"`
	Properties map[string]Property `json:"properties"`
}

func connectors(n *drm.Node, card *drm.ModeCard) ([]Connector, error) {
	l := make([]Connector, len(card.Connectors))
	for i, id := range card.Connectors {
		conn, err := n.ModeGetConnector(id)
		if err != nil {
			return nil, fmt.Errorf("failed to get connector: %v", err)
		}

		props, err := properties(n, conn.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get connector properties: %v", err)
		}

		l[i] = Connector{
			ID:         conn.ID,
			Type:       conn.Type,
			Status:     conn.Status,
			PhyWidth:   conn.PhyWidth,
			PhyHeight:  conn.PhyHeight,
			Subpixel:   conn.Subpixel,
			Encoders:   conn.PossibleEncoders,
			Modes:      modeList(conn.Modes),
			Properties: props,
		}
	}
	return l, nil
}

type Encoder struct {
	ID             drm.EncoderID   `json:"id"`
	Type           drm.EncoderType `json:"type"`
	CRTC           drm.CRTCID      `json:"crtc"`
	PossibleCRTCs  uint32          `json:"possible_crtcs"`
	PossibleClones uint32          `json:"possible_clones"`
}

func encoders(n *drm.Node, card *drm.ModeCard) ([]Encoder, error) {
	l := make([]Encoder, len(card.Encoders))
	for i, id := range card.Encoders {
		enc, err := n.ModeGetEncoder(id)
		if err != nil {
			return nil, fmt.Errorf("failed to get encoder: %v", err)
		}

		l[i] = Encoder(*enc)
	}
	return l, nil
}

type CRTC struct {
	ID         drm.CRTCID          `json:"id"`
	FB         drm.FBID            `json:"fb"`
	X          uint32              `json:"x"`
	Y          uint32              `json:"y"`
	GammaSize  uint32              `json:"gamma_size"`
	Mode       *Mode               `json:"mode"`
	Properties map[string]Property `json:"properties"`
}

func crtcs(n *drm.Node, card *drm.ModeCard) ([]CRTC, error) {
	l := make([]CRTC, len(card.CRTCs))
	for i, id := range card.CRTCs {
		crtc, err := n.ModeGetCRTC(id)
		if err != nil {
			return nil, fmt.Errorf("failed to get CRTC: %v", err)
		}

		props, err := properties(n, crtc.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get CRTC properties: %v", err)
		}

		l[i] = CRTC{
			ID:         crtc.ID,
			FB:         crtc.FB,
			X:          crtc.X,
			Y:          crtc.Y,
			GammaSize:  crtc.GammaSize,
			Mode:       (*Mode)(crtc.Mode),
			Properties: props,
		}
	}
	return l, nil
}

type Plane struct {
	ID            drm.PlaneID         `json:"id"`
	CRTC          drm.CRTCID          `json:"crtc"`
	FB            drm.FBID            `json:"fb"`
	PossibleCRTCs uint32              `json:"possible_crtcs"`
	GammaSize     uint32              `json:"gamma_size"`
	Formats       []drm.Format        `json:"formats"`
	Properties    map[string]Property `json:"properties"`
}

func planes(n *drm.Node) ([]Plane, error) {
	planes, err := n.ModeGetPlaneResources()
	if err != nil {
		log.Fatal(err)
	}

	l := make([]Plane, len(planes))
	for i, id := range planes {
		plane, err := n.ModeGetPlane(id)
		if err != nil {
			return nil, fmt.Errorf("failed to get CRTC: %v", err)
		}

		props, err := properties(n, plane.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get plane properties: %v", err)
		}

		l[i] = Plane{
			ID:            plane.ID,
			CRTC:          plane.CRTC,
			FB:            plane.FB,
			PossibleCRTCs: plane.PossibleCRTCs,
			GammaSize:     plane.GammaSize,
			Formats:       plane.Formats,
			Properties:    props,
		}
	}

	return l, nil
}

type Node struct {
	Driver     *Driver     `json:"driver"`
	Device     *Device     `json:"device"`
	Connectors []Connector `json:"connectors"`
	Encoders   []Encoder   `json:"encoders"`
	CRTCs      []CRTC      `json:"crtcs"`
	Planes     []Plane     `json:"planes"`
}

func node(nodePath string) (*Node, error) {
	f, err := os.Open(nodePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open DRM node: %v", err)
	}
	defer f.Close()

	n := drm.NewNode(f.Fd())

	drv, err := driver(n)
	if err != nil {
		return nil, err
	}

	dev, err := device(n)
	if err != nil {
		return nil, err
	}

	r, err := n.ModeGetResources()
	if err != nil {
		return nil, fmt.Errorf("failed to get DRM resources: %v", err)
	}

	conns, err := connectors(n, r)
	if err != nil {
		return nil, err
	}

	encs, err := encoders(n, r)
	if err != nil {
		return nil, err
	}

	crtcs, err := crtcs(n, r)
	if err != nil {
		return nil, err
	}

	planes, err := planes(n)
	if err != nil {
		return nil, err
	}

	return &Node{
		Driver:     drv,
		Device:     dev,
		Connectors: conns,
		Encoders:   encs,
		CRTCs:      crtcs,
		Planes:     planes,
	}, nil
}

func main() {
	paths, err := filepath.Glob(drm.NodePatternPrimary)
	if err != nil {
		log.Fatalf("Failed to list DRM nodes: %v", err)
	}

	m := make(map[string]*Node)
	for _, p := range paths {
		n, err := node(p)
		if err != nil {
			log.Fatal(err)
		}
		m[p] = n
	}

	err = json.NewEncoder(os.Stdout).Encode(m)
	if err != nil {
		log.Fatalf("Failed to write JSON: %v", err)
	}
}
