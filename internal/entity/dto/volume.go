package dto

type Volume struct {
	Name       string            `json:"name" form:"name"`
	Driver     string            `json:"driver" form:"driver"`
	Labels     map[string]string `json:"label" form:"labels"`
	DriverOpts map[string]string `json:"driver_ops" form:"driver_opts"`
}