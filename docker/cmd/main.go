package main

import (
	"log"
)

func main() {
	opts, err := ParseFlags()
	if err != nil {
		log.Fatalln("failed to parse options:", err)
	}

	images := ListDockerImages(opts.ImageName)
	removeList := GetRemoveList(images, opts.MaxCount)
	RemoveDockerImages(removeList)
}
