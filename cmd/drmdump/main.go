package main

import (
	"log"
	"os"
	"path/filepath"

	"git.sr.ht/~emersion/go-drm"
)

func node(nodePath string) {
	f, err := os.Open(nodePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n := drm.NewNode(f.Fd())
	v, err := n.Version()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Version", v)

	dev, err := n.GetDevice()
	log.Println("GetDevice", dev, err)

	val, err := n.GetCap(drm.CapDumbBuffer)
	log.Println("CapDumbBuffer", val, err)

	err = n.SetClientCap(drm.ClientCapUniversalPlanes, 1)
	log.Println("ClientCapUniversalPlanes", err)

	r, err := n.ModeGetResources()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ModeGetResources", r)

	for _, id := range r.CRTCs {
		crtc, err := n.ModeGetCRTC(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("ModeGetCRTC", crtc, crtc.Mode)
	}

	for _, id := range r.Encoders {
		enc, err := n.ModeGetEncoder(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("ModeGetEncoder", enc)
	}

	for _, id := range r.Connectors {
		conn, err := n.ModeGetConnector(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("ModeGetConnector", conn)
	}

	planes, err := n.ModeGetPlaneResources()
	if err != nil {
		log.Fatal(err)
	}

	for _, id := range planes {
		plane, err := n.ModeGetPlane(id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("ModeGetPlane", plane)
	}

	if len(planes) > 0 {
		props, err := n.ModeObjectGetProperties(planes[0])
		if err != nil {
			log.Fatal(err)
		}

		for id := range props {
			prop, err := n.ModeGetProperty(id)
			if err != nil {
				log.Fatal(err)
			}

			log.Println("ModeGetProperty", prop)
		}
	}
}

func main() {
	paths, err := filepath.Glob(drm.NodePatternPrimary)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range paths {
		log.Println("Node", p)
		node(p)
	}
}
