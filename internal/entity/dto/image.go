package dto

import "io"

type Image struct {
	Repository string `qs:"fromImage" json:"repository"`
	Tag        string `json:"string"`
	Platform   string `ver:"1.32" json:"platform"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}

type BuildImage struct {
	Dockerfile   string   `json:"file" form:"file"`
	Name         string   `qs:"t"`
	Labels       map[string]string
	InputStream  io.Reader `qs:"-"`
	OutputStream io.Writer `qs:"-"`
	Remote       string
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
}
