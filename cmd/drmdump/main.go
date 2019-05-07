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

	d := drm.NewDevice(f.Fd())
	v, err := d.Version()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Version", v)

	val, err := d.GetCap(drm.CapDumbBuffer)
	log.Println("CapDumbBuffer", val, err)
}
