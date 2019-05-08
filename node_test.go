package drm_test

import (
	"git.sr.ht/~emersion/go-drm"
)

var (
	_ drm.AnyID = drm.ObjectID(0)
	_ drm.AnyID = drm.CRTCID(0)
	_ drm.AnyID = drm.ConnectorID(0)
	_ drm.AnyID = drm.EncoderID(0)
	_ drm.AnyID = drm.ModeID(0)
	_ drm.AnyID = drm.PropertyID(0)
	_ drm.AnyID = drm.FBID(0)
	_ drm.AnyID = drm.BlobID(0)
	_ drm.AnyID = drm.PlaneID(0)
)
