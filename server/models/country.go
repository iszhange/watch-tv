package models

import (
	"watchtv/requests"

	"gorm.io/gorm"
)

type Country struct {
	Inst *gorm.DB
}

type CountryModel struct {
	gorm.Model
	Name      string `json:"name"`
	Code      string `json:"code"`
	Languages string `json:"languages"`
	Flag      string `json:"flag"`
	IsShow    int    `json:"is_show"`
}

func (o *Country) GetList(query requests.Country) []CountryModel {
	if query.Code != "" {
		o.Inst.Where("code = ?", query.Code)
	}
	if query.IsShow > 0 {
		o.Inst.Where("is_show = ?", query.IsShow)
	}

}
