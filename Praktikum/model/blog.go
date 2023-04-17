package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	UserId uint   `json:"user_id"  form:"user_id"`
	Judul  string `json:"judul"    form:"judul"`
	Konten string `json:"konten"   form:"konten"`
	User   User
}
