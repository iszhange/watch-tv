package models

import (
	"context"
	"time"
	"watchtv/requests"

	"gorm.io/gorm"
)

type Stream struct {
	inst         *gorm.DB
	Channel      string     `json:"channel"`
	URL          string     `json:"url"`
	HttpReferrer string     `json:"http_referrer"`
	UserAgent    string     `json:"user_agent"`
	UpdatedAt    *time.Time `json:"-"`
	CreatedAt    *time.Time `json:"-"`
}

func (o Stream) TableName() string {
	return "wt_streams"
}

func (o *Stream) SetDB(ctx context.Context) {
	o.inst, _ = ctx.Value(GormContextKey).(*gorm.DB)
}

func (o *Stream) GetList(query requests.Stream) (int64, []Stream) {
	var streams []Stream
	stream := Stream{}
	if query.Channel != "" {
		stream.Channel = query.Channel
	}
	result := o.inst.Where(&stream).Order("sort asc").Find(&streams)

	return result.RowsAffected, streams
}
