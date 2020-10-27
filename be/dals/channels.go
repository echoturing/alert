package dals

import (
	"context"

	"github.com/echoturing/alert/ent"
)

func (i *impl) GetChannelByID(ctx context.Context, id int64) (*ent.Channel, error) {
	return i.client.Channel.Get(ctx, id)
}
