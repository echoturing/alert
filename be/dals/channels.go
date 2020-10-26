package dals

import (
	"context"

	"github.com/echoturing/alert/channels"
)

func (i *impl) GetChannelByID(ctx context.Context, id int64) (*channels.Channel, error) {
	panic("implement me")
}
