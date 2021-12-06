package main

import "time"

type DockerImage struct {
	ID        string
	Name      string
	Tag       string
	CreatedAt time.Time
}

type DockerImages []*DockerImage

func (d DockerImages) Len() int {
	return len(d)
}

func (d DockerImages) Less(i, j int) bool {
	return d[i].CreatedAt.Before(d[j].CreatedAt)
}

func (d DockerImages) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
