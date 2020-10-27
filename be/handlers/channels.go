package handlers

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/echoturing/alert/ent"
	"github.com/echoturing/alert/services"
)

type ListChannelsReply struct {
	List []*ent.Channel `json:"list"`
}

func (i *impl) ListChannels(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	channels, err := i.service.ListChannels(ctx)
	if err != nil {
		return nil, err
	}
	return &ListChannelsReply{
		List: channels,
	}, nil
}

func (i *impl) CreateChannel(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	channel := &ent.Channel{}
	if err := c.Bind(channel); err != nil {
		return nil, err
	}

	channel, err := i.service.CreateChannel(ctx, channel)
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func (i *impl) UpdateChannel(c echo.Context) (interface{}, error) {
	ctx := c.Request().Context()
	update := &services.UpdateChannelRequest{}
	if err := c.Bind(update); err != nil {
		return nil, err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		return nil, err
	}
	channel, err := i.service.UpdateChannel(ctx, id, update)
	if err != nil {
		return nil, err
	}
	return channel, nil
}
