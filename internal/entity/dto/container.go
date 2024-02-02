package dto

import docker "github.com/fsouza/go-dockerclient"

type HostConfig struct {
	Binds       []string                           `json:"binding"`
	NetworkMode string                             `json:"network_mode"`
	PortBinding map[docker.Port][]docker.PortBinding `json:"port_binding"`
}

type Config struct {
	Hostname   string                   `json:"hostname,omitempty"`
	Domainname string                   `json:"domain_name,omitempty"`
	Image      string                   `json:"image,omitempty"`
	Tty        bool                     `json:"tty"`
	OpenStdin  bool                     `json:"bool"`
	Env        []string                 `json:"env"`
	Port       map[docker.Port]struct{} `json:"port"`
}

type Container struct {
	Name       string      `json:"name"`
	Platform   string      `json:"platform"`
	Config     *Config     `json:"config"`
	HostConfig *HostConfig `json:"host_config"`
}
