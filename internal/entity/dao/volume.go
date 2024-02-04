package dao

type Volume struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Driver     string            `json:"driver"`
	Labels     map[string]string `json:"label"`
	DriverOpts map[string]string `json:"driver_ops"`
}