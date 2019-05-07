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
	log.Println("ModeGetResources", r, err)
}
