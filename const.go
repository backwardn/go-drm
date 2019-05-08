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

type Connector uint32

const (
	ConnectorUnknown     Connector = 0
	ConnectorVGA         Connector = 1
	ConnectorDVII        Connector = 2
	ConnectorDVID        Connector = 3
	ConnectorDVIA        Connector = 4
	ConnectorComposite   Connector = 5
	ConnectorSVideo      Connector = 6
	ConnectorLVDS        Connector = 7
	ConnectorComponent   Connector = 8
	Connector9PinDIN     Connector = 9
	ConnectorDisplayPort Connector = 10
	ConnectorHDMIA       Connector = 11
	ConnectorHDMIB       Connector = 12
	ConnectorTV          Connector = 13
	ConnectorEDP         Connector = 14
	ConnectorVirtual     Connector = 15
	ConnectorDSI         Connector = 16
	ConnectorDPI         Connector = 17
	ConnectorWriteback   Connector = 18
)

func (c Connector) String() string {
	switch c {
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
