package drm

type Cap uint64

const (
	CapDumbBuffer Cap = 0x1
	CapVblankHighCRTC Cap = 0x2
	CapDumbPreferredDepth Cap = 0x3
	CapDumbPreferredShadow Cap = 0x4
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
