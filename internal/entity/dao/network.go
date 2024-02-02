package dao

import docker "github.com/fsouza/go-dockerclient"

type Network struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Driver     string `json:"driver"`
	EnableIPV6 bool   `json:"ipv6"`
	Internal   bool   `json:"internal"`
	Scope      string `json:"scope"`
	Container  map[string]docker.Endpoint `json:"container"`
}