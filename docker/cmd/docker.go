package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

const (
	cmdDockerImages       = "docker image ls"
	cmdDockerInspect      = "docker inspect -f '{{.Created}}' %s"
	cmdDockerImagesRemove = "docker image rm %s"
)

func ListDockerImages(imageName string) DockerImages {
	out := RunCmd(cmdDockerImages)

	var images DockerImages

	split := strings.Split(out, "\n")
	for i, s := range split {
		if i == 0 {
			continue
		}

		fields := strings.Fields(s)
		if len(fields) == 0 {
			continue
		}

		if fields[0] == imageName {
			image := &DockerImage{
				ID:   fields[2],
				Name: fields[0],
				Tag:  fields[1],
			}
			images = append(images, image)
		}
	}
	return images
}

func GetRemoveList(images DockerImages, maxCount int) DockerImages {
	if images == nil || maxCount == 0 {
		return nil
	}

	if len(images) > maxCount {
		for _, image := range images {
			out := RunCmd(fmt.Sprintf(cmdDockerInspect, image.ID))
			parse, err := time.Parse("2006-01-02T15:04:05Z", strings.ReplaceAll(out, "\x0a", ""))
			if err != nil {
				log.Fatalln("failed to parse created time:", err)
			}
			image.CreatedAt = parse
		}
	}

	sort.Sort(sort.Reverse(images))

	removeList := images[maxCount:]
	return removeList
}

func RemoveDockerImages(removeList DockerImages) {
	for _, image := range removeList {
		out := RunCmd(fmt.Sprintf(cmdDockerImagesRemove, image.ID))
		if out != "" {
			fmt.Println("command output:")
			fmt.Println(out)
		}
	}
}
