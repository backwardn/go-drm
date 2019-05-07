package drm

type Cap uint64

const (
	CapDumbBuffer Cap = 0x1
	CapVblankHighCRTC Cap = 0x2
	CapDumbPreferredDepth Cap = 0x3
	CapDumbPreferredShadow Cap = 0x4
)
