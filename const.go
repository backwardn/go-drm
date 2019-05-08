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
		return "Primary"
	case NodeControl:
		return "Control"
	case NodeRender:
		return "Render"
	default:
		return "Unknown"
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
		return "Platform"
	case BusHost1x:
		return "host1x"
	default:
		return "Unknown"
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
)

type ClientCap uint64

const (
	ClientCapStereo3D            ClientCap = 1
	ClientCapUniversalPlanes     ClientCap = 2
	ClientCapAtomic              ClientCap = 3
	ClientCapAspectRatio         ClientCap = 4
	ClientCapWritebackConnectors ClientCap = 5
)

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
		return "DVII"
	case ConnectorDVID:
		return "DVID"
	case ConnectorDVIA:
		return "DVIA"
	case ConnectorComposite:
		return "Composite"
	case ConnectorSVideo:
		return "SVideo"
	case ConnectorLVDS:
		return "LVDS"
	case ConnectorComponent:
		return "Component"
	case Connector9PinDIN:
		return "9PinDIN"
	case ConnectorDisplayPort:
		return "DisplayPort"
	case ConnectorHDMIA:
		return "HDMIA"
	case ConnectorHDMIB:
		return "HDMIB"
	case ConnectorTV:
		return "TV"
	case ConnectorEDP:
		return "EDP"
	case ConnectorVirtual:
		return "Virtual"
	case ConnectorDSI:
		return "DSI"
	case ConnectorDPI:
		return "DPI"
	case ConnectorWriteback:
		return "Writeback"
	default:
		return "Unknown"
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
		return "None"
	case EncoderDAC:
		return "DAC"
	case EncoderTDMS:
		return "TDMS"
	case EncoderLVDS:
		return "LVDS"
	case EncoderTVDAC:
		return "TVDAC"
	case EncoderVirtual:
		return "Virtual"
	case EncoderDSI:
		return "DSI"
	case EncoderDPMST:
		return "DPMST"
	case EncoderDPI:
		return "DPI"
	default:
		return "Unknown"
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
		return "Connected"
	case ConnectorStatusDisconnected:
		return "Disconnected"
	default:
		return "Unknown"
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
		return "Horizontal RGB"
	case SubpixelHorizontalBGR:
		return "Horizontal BGR"
	case SubpixelVerticalRGB:
		return "Vertical RGB"
	case SubpixelVerticalBGR:
		return "Vertical BGR"
	case SubpixelNone:
		return "None"
	default:
		return "Unknown"
	}
}
