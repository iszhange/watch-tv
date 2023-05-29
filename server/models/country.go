package models

import (
	"watchtv/requests"

	"gorm.io/gorm"
)

type Country struct {
	inst *gorm.DB
	gorm.Model
	Name      string `json:"name"`
	Code      string `json:"code"`
	Languages string `json:"languages"`
	Flag      string `json:"flag"`
	IsShow    int    `json:"is_show"`
}

func (o Country) TableName() string {
	return "wt_countries"
}

func (o *Country) SetDB(db *gorm.DB) {
	o.inst = db
}

func (o *Country) GetList(query requests.Country) (int, []Country) {
	var countries []Country
	country := Country{}
	if query.Code != "" {
		country.Code = query.Code
	}
	if query.IsShow > 0 {
		country.IsShow = query.IsShow
	}
	result := o.inst.Where(&country).Find(&countries)

	return int(result.RowsAffected), countries
}
