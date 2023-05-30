package models

import (
	"context"
	"time"
	"watchtv/requests"

	"gorm.io/gorm"
)

type Channel struct {
	inst          *gorm.DB
	ID            string     `json:"id"`
	Name          string     `gorm:"name" json:"name"`
	AltNames      []string   `gorm:"serializer:json" json:"alt_names"`
	Network       string     `json:"network"`
	Owners        []string   `gorm:"serializer:json" json:"owners"`
	Country       string     `json:"country"`
	Subdivision   string     `json:"subdivision"`
	City          string     `json:"city"`
	BroadcastArea []string   `gorm:"serializer:json" json:"broadcast_area"`
	Languages     []string   `gorm:"serializer:json" json:"languages"`
	Categories    []string   `gorm:"serializer:json" json:"categories"`
	IsNsfw        int        `json:"is_nsfw"`
	Launched      *time.Time `json:"launched"`
	Closed        *time.Time `json:"closed"`
	ReplacedBy    string     `json:"replaced_by"`
	Website       string     `json:"website"`
	Logo          string     `json:"logo"`
	Sort          int        `json:"sort"`
	IsShow        int        `json:"is_show"`
	UpdatedAt     *time.Time `json:"-"`
	CreatedAt     *time.Time `json:"-"`
}

func (o Channel) TableName() string {
	return "wt_channels"
}

func (o *Channel) SetDB(ctx context.Context) {
	o.inst, _ = ctx.Value(GormContextKey).(*gorm.DB)
}

func (o *Channel) GetList(query requests.Channel) (int64, []Channel) {
	var channels []Channel
	channel := Channel{}
	if query.Country != "" {
		channel.Country = query.Country
	}
	if query.IsNsfw > 0 {
		channel.IsNsfw = query.IsNsfw
	}
	if query.IsShow > 0 {
		channel.IsShow = query.IsShow
	}
	result := o.inst.Where(&channel).Order("sort asc").Find(&channels)

	return result.RowsAffected, channels
}
