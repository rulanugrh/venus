package dto

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