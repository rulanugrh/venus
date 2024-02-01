package dao

type Image struct {
	ID          string   `json:"id"`
	Tag         []string `json:"tag"`
	Created     int64    `json:"created"`
	Size        int64    `json:"size"`
	VirtualSize int64    `json:"virtual_size"`
	Labels      map[string]string
}
