package drm

type NodeType int

const (
	NodePrimary NodeType = 0
	NodeControl NodeType = 1
	NodeRender  NodeType = 2
)

func (t NodeType) String() string {
	switch t {
	case NodePrimary:
		return "primary"
	case NodeControl:
		return "control"
	case NodeRender:
		return "render"
	default:
		return "unknown"
	}
}

type BusType int

const (
	BusPCI      BusType = 0
	BusUSB      BusType = 1
	BusPlatform BusType = 2
	BusHost1x   BusType = 3
)

func (t BusType) String() string {
	switch t {
	case BusPCI:
		return "PCI"
	case BusUSB:
		return "USB"
	case BusPlatform:
		return "platform"
	case BusHost1x:
		return "host1x"
	default:
		return "unknown"
	}
}

type Cap uint64

const (
	CapDumbBuffer          Cap = 0x1
	CapVblankHighCRTC      Cap = 0x2
	CapDumbPreferredDepth  Cap = 0x3
	CapDumbPreferredShadow Cap = 0x4
	CapPrime               Cap = 0x5
	CapTimestampMonotonic  Cap = 0x6
	CapAsyncPageFlip       Cap = 0x7
	CapCursorWidth         Cap = 0x8
	CapCursorHeight        Cap = 0x9
	CapAddFB2Modifiers     Cap = 0x10
	CapPageFlipTarget      Cap = 0x11
	CapCRTCInVBlankEvent   Cap = 0x12
	CapSyncObj             Cap = 0x13
	CapSyncObjTimeline     Cap = 0x14
)

func (c Cap) String() string {
	switch c {
	case CapDumbBuffer:
		return "DUMB_BUFFER"
	case CapVblankHighCRTC:
		return "VBLANK_HIGH_CRTC"
	case CapDumbPreferredDepth:
		return "DUMB_PREFERRED_DEPTH"
	case CapDumbPreferredShadow:
		return "DUMB_PREFER_SHADOW"
	case CapPrime:
		return "PRIME"
	case CapTimestampMonotonic:
		return "TIMESTAMP_MONOTONIC"
	case CapAsyncPageFlip:
		return "ASYNC_PAGE_FLIP"
	case CapCursorWidth:
		return "CURSOR_WIDTH"
	case CapCursorHeight:
		return "CURSOR_HEIGHT"
	case CapAddFB2Modifiers:
		return "ADDFB2_MODIFIERS"
	case CapPageFlipTarget:
		return "PAGE_FLIP_TARGET"
	case CapCRTCInVBlankEvent:
		return "CRTC_IN_VBLANK_EVENT"
	case CapSyncObj:
		return "SYNCOBJ"
	case CapSyncObjTimeline:
		return "SYNCOBJ_TIMELINE"
	}
	return "unknown"
}

type ClientCap uint64

const (
	ClientCapStereo3D            ClientCap = 1
	ClientCapUniversalPlanes     ClientCap = 2
	ClientCapAtomic              ClientCap = 3
	ClientCapAspectRatio         ClientCap = 4
	ClientCapWritebackConnectors ClientCap = 5
)

func (c ClientCap) String() string {
	switch c {
	case ClientCapStereo3D:
		return "STEREO_3D"
	case ClientCapUniversalPlanes:
		return "UNIVERSAL_PLANES"
	case ClientCapAtomic:
		return "ATOMIC"
	case ClientCapAspectRatio:
		return "APSECT_RATIO"
	case ClientCapWritebackConnectors:
		return "WRITEBACK_CONNECTORS"
	}
	return "unknown"
}

const (
	CapPrimeImport = 0x1
	CapPrimeExport = 0x2
)

type ConnectorType uint32

const (
	ConnectorUnknown     ConnectorType = 0
	ConnectorVGA         ConnectorType = 1
	ConnectorDVII        ConnectorType = 2
	ConnectorDVID        ConnectorType = 3
	ConnectorDVIA        ConnectorType = 4
	ConnectorComposite   ConnectorType = 5
	ConnectorSVideo      ConnectorType = 6
	ConnectorLVDS        ConnectorType = 7
	ConnectorComponent   ConnectorType = 8
	Connector9PinDIN     ConnectorType = 9
	ConnectorDisplayPort ConnectorType = 10
	ConnectorHDMIA       ConnectorType = 11
	ConnectorHDMIB       ConnectorType = 12
	ConnectorTV          ConnectorType = 13
	ConnectorEDP         ConnectorType = 14
	ConnectorVirtual     ConnectorType = 15
	ConnectorDSI         ConnectorType = 16
	ConnectorDPI         ConnectorType = 17
	ConnectorWriteback   ConnectorType = 18
)

func (t ConnectorType) String() string {
	switch t {
	case ConnectorVGA:
		return "VGA"
	case ConnectorDVII:
		return "DVI-I"
	case ConnectorDVID:
		return "DVI-D"
	case ConnectorDVIA:
		return "DVI-A"
	case ConnectorComposite:
		return "composite"
	case ConnectorSVideo:
		return "S-Video"
	case ConnectorLVDS:
		return "LVDS"
	case ConnectorComponent:
		return "component"
	case Connector9PinDIN:
		return "9PinDIN"
	case ConnectorDisplayPort:
		return "DisplayPort"
	case ConnectorHDMIA:
		return "HDMI-A"
	case ConnectorHDMIB:
		return "HDMI-B"
	case ConnectorTV:
		return "TV"
	case ConnectorEDP:
		return "eDP"
	case ConnectorVirtual:
		return "virtual"
	case ConnectorDSI:
		return "DSI"
	case ConnectorDPI:
		return "DPI"
	case ConnectorWriteback:
		return "writeback"
	default:
		return "unknown"
	}
}

type EncoderType uint32

const (
	EncoderNone    EncoderType = 0
	EncoderDAC     EncoderType = 1
	EncoderTDMS    EncoderType = 2
	EncoderLVDS    EncoderType = 3
	EncoderTVDAC   EncoderType = 4
	EncoderVirtual EncoderType = 5
	EncoderDSI     EncoderType = 6
	EncoderDPMST   EncoderType = 7
	EncoderDPI     EncoderType = 8
)

func (t EncoderType) String() string {
	switch t {
	case EncoderNone:
		return "none"
	case EncoderDAC:
		return "DAC"
	case EncoderTDMS:
		return "TDMS"
	case EncoderLVDS:
		return "LVDS"
	case EncoderTVDAC:
		return "TV DAC"
	case EncoderVirtual:
		return "virtual"
	case EncoderDSI:
		return "DSI"
	case EncoderDPMST:
		return "DP MST"
	case EncoderDPI:
		return "DPI"
	default:
		return "unknown"
	}
}

type ConnectorStatus uint32

const (
	ConnectorStatusConnected    ConnectorStatus = 1
	ConnectorStatusDisconnected ConnectorStatus = 2
	ConnectorStatusUnknown      ConnectorStatus = 3
)

func (s ConnectorStatus) String() string {
	switch s {
	case ConnectorStatusConnected:
		return "connected"
	case ConnectorStatusDisconnected:
		return "disconnected"
	default:
		return "unknown"
	}
}

type Subpixel uint32

const (
	SubpixelUnknown       Subpixel = 0
	SubpixelHorizontalRGB Subpixel = 1
	SubpixelHorizontalBGR Subpixel = 3
	SubpixelVerticalRGB   Subpixel = 4
	SubpixelVerticalBGR   Subpixel = 5
	SubpixelNone          Subpixel = 6
)

func (s Subpixel) String() string {
	switch s {
	case SubpixelHorizontalRGB:
		return "horizontal RGB"
	case SubpixelHorizontalBGR:
		return "horizontal BGR"
	case SubpixelVerticalRGB:
		return "vertical RGB"
	case SubpixelVerticalBGR:
		return "vertical BGR"
	case SubpixelNone:
		return "none"
	default:
		return "unknown"
	}
}

type PlaneType uint32

const (
	PlaneOverlay PlaneType = 0
	PlanePrimary PlaneType = 1
	PlaneCursor PlaneType = 2
)

func (t PlaneType) String() string {
	switch t {
	case PlaneOverlay:
		return "overlay"
	case PlanePrimary:
		return "primary"
	case PlaneCursor:
		return "cursor"
	default:
		return "unknown"
	}
}
