package dto

type Network struct {
	Name       string `json:"name"`
	Driver     string `json:"driver"`
	Internal   bool   `json:"internal"`
	EnableIPV6 bool   `json:"ipv6"`
	Scope      string `json:"scope"`
}