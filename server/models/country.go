package models

import (
	"context"
	"time"
	"watchtv/requests"

	"gorm.io/gorm"
)

type Country struct {
	inst      *gorm.DB
	Name      string     `json:"name"`
	Code      string     `gorm:"primaryKey" json:"code"`
	Languages []string   `gorm:"serializer:json" json:"languages"`
	Flag      string     `json:"flag"`
	IsShow    int        `json:"is_show"`
	UpdatedAt *time.Time `json:"-"`
	CreatedAt *time.Time `json:"-"`
}

func (o Country) TableName() string {
	return "wt_countries"
}

func (o *Country) SetDB(ctx context.Context) {
	o.inst, _ = ctx.Value(GormContextKey).(*gorm.DB)
}

func (o *Country) GetList(query requests.Country) (int64, []Country) {
	var countries []Country
	country := Country{}
	if query.Code != "" {
		country.Code = query.Code
	}
	if query.IsShow > 0 {
		country.IsShow = query.IsShow
	}
	result := o.inst.Where(&country).Find(&countries)

	return result.RowsAffected, countries
}
