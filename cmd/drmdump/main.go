package main

import (
	"log"
	"os"

	"git.sr.ht/~emersion/go-drm"
)

func main() {
	devPath := "/dev/dri/card0"

	f, err := os.Open(devPath)
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

	val, err := n.GetCap(drm.CapDumbBuffer)
	log.Println("CapDumbBuffer", val, err)

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
}
