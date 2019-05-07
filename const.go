package drm

type Cap uint64

const (
	CapDumbBuffer Cap = 0x1
	CapVblankHighCRTC Cap = 0x2
	CapDumbPreferredDepth Cap = 0x3
	CapDumbPreferredShadow Cap = 0x4
	CapPrime Cap = 0x5
	CapTimestampMonotonic Cap =0x6
	CapAsyncPageFlip Cap = 0x7
	CapCursorWidth Cap = 0x8
	CapCursorHeight Cap = 0x9
	CapAddFB2Modifiers Cap = 0x10
	CapPageFlipTarget Cap = 0x11
	CapCRTCInVBlankEvent Cap = 0x12
	CapSyncObj Cap = 0x13
)

const (
	CapPrimeImport = 0x1
	CapPrimeExport = 0x2
)

type Connector uint32

const (
	ConnectorUnknown Connector = 0
	ConnectorVGA Connector = 1
	ConnectorDVII Connector = 2
	ConnectorDVID Connector = 3
	ConnectorDVIA Connector = 4
	ConnectorComposite Connector = 5
	ConnectorSVideo Connector = 6
	ConnectorLVDS Connector = 7
	ConnectorComponent Connector = 8
	Connector9PinDIN Connector = 9
	ConnectorDisplayPort Connector = 10
	ConnectorHDMIA Connector = 11
	ConnectorHDMIB Connector = 12
	ConnectorTV Connector = 13
	ConnectorEDP Connector = 14
	ConnectorVirtual Connector = 15
	ConnectorDSI Connector = 16
	ConnectorDPI Connector = 17
	ConnectorWriteback Connector = 18
)
