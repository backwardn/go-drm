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

type Format uint32

const (
	FormatInvalid Format = 0
)

func (f Format) String() string {
	r1 := rune(uint32(f) & 0xFF)
	r2 := rune(uint32(f>>8) & 0xFF)
	r3 := rune(uint32(f>>16) & 0xFF)
	r4 := rune(uint32(f>>24) & 0xFF)
	return string([]rune{r1, r2, r3, r4})
}

type Modifier uint64

const (
	ModifierLinear  Modifier = 0
	ModifierInvalid Modifier = (1 << 56) - 1
)

func (mod Modifier) Vendor() ModifierVendor {
	return ModifierVendor(mod >> 56)
}

type ModifierVendor uint8

const (
	ModifierVendorNone      ModifierVendor = 0
	ModifierVendorIntel     ModifierVendor = 0x01
	ModifierVendorAMD       ModifierVendor = 0x02
	ModifierVendorNVIDIA    ModifierVendor = 0x03
	ModifierVendorSamsung   ModifierVendor = 0x04
	ModifierVendorQcom      ModifierVendor = 0x05
	ModifierVendorVivante   ModifierVendor = 0x06
	ModifierVendorBroadcom  ModifierVendor = 0x07
	ModifierVendorARM       ModifierVendor = 0x08
	ModifierVendorAllwinner ModifierVendor = 0x09
)

func (vendor ModifierVendor) String() string {
	switch vendor {
	case ModifierVendorNone:
		return "none"
	case ModifierVendorIntel:
		return "Intel"
	case ModifierVendorAMD:
		return "AMD"
	case ModifierVendorNVIDIA:
		return "NVIDIA"
	case ModifierVendorSamsung:
		return "Samsung"
	case ModifierVendorQcom:
		return "Qcom"
	case ModifierVendorVivante:
		return "Vivante"
	case ModifierVendorBroadcom:
		return "Broadcom"
	case ModifierVendorARM:
		return "ARM"
	case ModifierVendorAllwinner:
		return "Allwinner"
	default:
		return "unknown"
	}
}
