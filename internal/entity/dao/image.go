package dao

import "time"

type Image struct {
	ID          string   `json:"id"`
	Tag         []string `json:"tag"`
	Created     int64    `json:"created"`
	Size        int64    `json:"size"`
	VirtualSize int64    `json:"virtual_size"`
	Labels      map[string]string
}

type InspectImage struct {
	ID            string    `json:"id"`
	Tag           []string  `json:"tag"`
	Created       time.Time `json:"created"`
	Container     string    `json:"container"`
	OS            string    `json:"os"`
	Architecture  string    `json:"architecture"`
	Size          int64     `json:"size"`
	VirtualSize   int64     `json:"virtual_size"`
	Author        string    `json:"author"`
	DockerVersion string    `json:"docker_version"`
}
