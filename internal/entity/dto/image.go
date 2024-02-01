package dto

type Image struct {
	Repository string `qs:"fromImage" json:"repository"`
	Tag        string `json:"string"`
	Platform   string `ver:"1.32" json:"platform"`
}