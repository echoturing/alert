package dals

import (
	"context"

	"github.com/echoturing/alert/ent"
)

func (i *impl) GetChannelByID(ctx context.Context, id int64) (*ent.Channel, error) {
	return i.client.Channel.Get(ctx, id)
}

func (i *impl) ListChannels(ctx context.Context) ([]*ent.Channel, error) {
	return i.client.Channel.Query().All(ctx)
}

func (i *impl) UpdateChannel(ctx context.Context, id int64, channel *ent.Channel) (*ent.Channel, error) {
	return i.client.Channel.UpdateOneID(id).
		SetName(channel.Name).
		SetType(channel.Type).
		SetDetail(channel.Detail).Save(ctx)
}

func (i *impl) CreateChannel(ctx context.Context, channel *ent.Channel) (*ent.Channel, error) {
	return i.client.Channel.Create().
		SetName(channel.Name).
		SetType(channel.Type).
		SetDetail(channel.Detail).Save(ctx)
}
