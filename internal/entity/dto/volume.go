package dto

type Volume struct {
	Name       string            `json:"name"`
	Driver     string            `json:"driver"`
	Labels     map[string]string `json:"label"`
	DriverOpts map[string]string `json:"driver_ops"`
}