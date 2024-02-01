package dto

type Image struct {
	Repository string `qs:"fromImage" json:"repository"`
	Tag        string `json:"string"`
	Platform   string `ver:"1.32" json:"platform"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}