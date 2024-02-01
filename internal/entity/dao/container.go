package dao

type Port struct {
	PrivatePort int64  `json:"private_port"`
	PublicPort  int64  `json:"public_port"`
	Type        string `json:"type"`
	IP          string `json:"ip"`
}

type Config struct {
	Hostname   string              `json:"hostname,omitempty"`
	Domainname string              `json:"domain_name,omitempty"`
	Image      string              `json:"image,omitempty"`
	Tty        bool                `json:"tty"`
	OpenStdin  bool                `json:"bool"`
	Env        []string            `json:"env"`
	Port       map[string]struct{} `json:"port"`
}

type Container struct {
	Name     string  `json:"name"`
	Platform string  `json:"platform"`
	Config   *Config `json:"config"`
}

type ListContainer struct {
	ID      string `json:"id"`
	Image   string `json:"image"`
	Command string `json:"command"`
	Status  string `json:"status"`
	State   string `json:"state"`
	Created int64  `json:"created"`
	Ports   []Port `json:"ports"`
}

type InspectContainer struct {
	ID           string              `json:"id"`
	Image        string              `json:"image"`
	HostnamePath string              `json:"hostname_path"`
	HostsPath    string              `json:"host_path"`
	Name         string              `json:"name"`
	Tty          bool                `json:"tty"`
	OpenStdin    bool                `json:"bool"`
	Env          []string            `json:"env"`
	Port         map[string]struct{} `json:"port"`
}