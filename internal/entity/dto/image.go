package dto

import "io"

type Image struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Platform   string `json:"platform"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}

type BuildImage struct {
	Dockerfile   string   `json:"file" form:"file"`
	Name         string   `qs:"t" json:"name" form:"name"`
	Labels       map[string]string `json:"labels" form:"labels"`
	InputStream  io.Reader `qs:"-" json:"input_stream"`
	OutputStream io.Writer `qs:"-" json:"output_stream"`
	Remote       string `json:"remote" form:"remote"`
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password"`
	Email        string `json:"email" form:"email"`
}
