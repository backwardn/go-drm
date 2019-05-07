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

func main() {
	paths, err := filepath.Glob(drm.NodePrimaryPattern)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range paths {
		log.Println("Node", p)
		node(p)
	}
}
