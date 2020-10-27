package services

import (
	"context"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/ent/schema"
)

func (i *impl) ListChannels(ctx context.Context) ([]*ent.Channel, error) {
	return i.dal.ListChannels(ctx)
}

func (i *impl) CreateChannel(ctx context.Context, channel *ent.Channel) (*ent.Channel, error) {
	return i.dal.CreateChannel(ctx, channel)
}

type UpdateChannelRequest struct {
	Name   string               `json:"name,omitempty"`
	Type   schema.ChannelType   `json:"type,omitempty"`
	Detail schema.ChannelDetail `json:"detail,omitempty"`
}

func (i *impl) UpdateChannel(ctx context.Context, id int64, update *UpdateChannelRequest) (*ent.Channel, error) {
	channel := &ent.Channel{
		Name:   update.Name,
		Type:   update.Type,
		Detail: update.Detail,
	}
	return i.dal.UpdateChannel(ctx, id, channel)
}
