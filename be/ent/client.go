// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/echoturing/alert/ent/migrate"

	"github.com/echoturing/alert/ent/alert"
	"github.com/echoturing/alert/ent/alerthistory"
	"github.com/echoturing/alert/ent/channel"
	"github.com/echoturing/alert/ent/datasource"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Alert is the client for interacting with the Alert builders.
	Alert *AlertClient
	// AlertHistory is the client for interacting with the AlertHistory builders.
	AlertHistory *AlertHistoryClient
	// Channel is the client for interacting with the Channel builders.
	Channel *ChannelClient
	// Datasource is the client for interacting with the Datasource builders.
	Datasource *DatasourceClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Alert = NewAlertClient(c.config)
	c.AlertHistory = NewAlertHistoryClient(c.config)
	c.Channel = NewChannelClient(c.config)
	c.Datasource = NewDatasourceClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Alert:        NewAlertClient(cfg),
		AlertHistory: NewAlertHistoryClient(cfg),
		Channel:      NewChannelClient(cfg),
		Datasource:   NewDatasourceClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:       cfg,
		Alert:        NewAlertClient(cfg),
		AlertHistory: NewAlertHistoryClient(cfg),
		Channel:      NewChannelClient(cfg),
		Datasource:   NewDatasourceClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Alert.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Alert.Use(hooks...)
	c.AlertHistory.Use(hooks...)
	c.Channel.Use(hooks...)
	c.Datasource.Use(hooks...)
}

// AlertClient is a client for the Alert schema.
type AlertClient struct {
	config
}

// NewAlertClient returns a client for the Alert from the given config.
func NewAlertClient(c config) *AlertClient {
	return &AlertClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `alert.Hooks(f(g(h())))`.
func (c *AlertClient) Use(hooks ...Hook) {
	c.hooks.Alert = append(c.hooks.Alert, hooks...)
}

// Create returns a create builder for Alert.
func (c *AlertClient) Create() *AlertCreate {
	mutation := newAlertMutation(c.config, OpCreate)
	return &AlertCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Alert entities.
func (c *AlertClient) CreateBulk(builders ...*AlertCreate) *AlertCreateBulk {
	return &AlertCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Alert.
func (c *AlertClient) Update() *AlertUpdate {
	mutation := newAlertMutation(c.config, OpUpdate)
	return &AlertUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AlertClient) UpdateOne(a *Alert) *AlertUpdateOne {
	mutation := newAlertMutation(c.config, OpUpdateOne, withAlert(a))
	return &AlertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AlertClient) UpdateOneID(id int64) *AlertUpdateOne {
	mutation := newAlertMutation(c.config, OpUpdateOne, withAlertID(id))
	return &AlertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Alert.
func (c *AlertClient) Delete() *AlertDelete {
	mutation := newAlertMutation(c.config, OpDelete)
	return &AlertDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AlertClient) DeleteOne(a *Alert) *AlertDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AlertClient) DeleteOneID(id int64) *AlertDeleteOne {
	builder := c.Delete().Where(alert.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AlertDeleteOne{builder}
}

// Query returns a query builder for Alert.
func (c *AlertClient) Query() *AlertQuery {
	return &AlertQuery{config: c.config}
}

// Get returns a Alert entity by its id.
func (c *AlertClient) Get(ctx context.Context, id int64) (*Alert, error) {
	return c.Query().Where(alert.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AlertClient) GetX(ctx context.Context, id int64) *Alert {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AlertClient) Hooks() []Hook {
	return c.hooks.Alert
}

// AlertHistoryClient is a client for the AlertHistory schema.
type AlertHistoryClient struct {
	config
}

// NewAlertHistoryClient returns a client for the AlertHistory from the given config.
func NewAlertHistoryClient(c config) *AlertHistoryClient {
	return &AlertHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `alerthistory.Hooks(f(g(h())))`.
func (c *AlertHistoryClient) Use(hooks ...Hook) {
	c.hooks.AlertHistory = append(c.hooks.AlertHistory, hooks...)
}

// Create returns a create builder for AlertHistory.
func (c *AlertHistoryClient) Create() *AlertHistoryCreate {
	mutation := newAlertHistoryMutation(c.config, OpCreate)
	return &AlertHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of AlertHistory entities.
func (c *AlertHistoryClient) CreateBulk(builders ...*AlertHistoryCreate) *AlertHistoryCreateBulk {
	return &AlertHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AlertHistory.
func (c *AlertHistoryClient) Update() *AlertHistoryUpdate {
	mutation := newAlertHistoryMutation(c.config, OpUpdate)
	return &AlertHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AlertHistoryClient) UpdateOne(ah *AlertHistory) *AlertHistoryUpdateOne {
	mutation := newAlertHistoryMutation(c.config, OpUpdateOne, withAlertHistory(ah))
	return &AlertHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AlertHistoryClient) UpdateOneID(id int64) *AlertHistoryUpdateOne {
	mutation := newAlertHistoryMutation(c.config, OpUpdateOne, withAlertHistoryID(id))
	return &AlertHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AlertHistory.
func (c *AlertHistoryClient) Delete() *AlertHistoryDelete {
	mutation := newAlertHistoryMutation(c.config, OpDelete)
	return &AlertHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AlertHistoryClient) DeleteOne(ah *AlertHistory) *AlertHistoryDeleteOne {
	return c.DeleteOneID(ah.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AlertHistoryClient) DeleteOneID(id int64) *AlertHistoryDeleteOne {
	builder := c.Delete().Where(alerthistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AlertHistoryDeleteOne{builder}
}

// Query returns a query builder for AlertHistory.
func (c *AlertHistoryClient) Query() *AlertHistoryQuery {
	return &AlertHistoryQuery{config: c.config}
}

// Get returns a AlertHistory entity by its id.
func (c *AlertHistoryClient) Get(ctx context.Context, id int64) (*AlertHistory, error) {
	return c.Query().Where(alerthistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AlertHistoryClient) GetX(ctx context.Context, id int64) *AlertHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AlertHistoryClient) Hooks() []Hook {
	return c.hooks.AlertHistory
}

// ChannelClient is a client for the Channel schema.
type ChannelClient struct {
	config
}

// NewChannelClient returns a client for the Channel from the given config.
func NewChannelClient(c config) *ChannelClient {
	return &ChannelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `channel.Hooks(f(g(h())))`.
func (c *ChannelClient) Use(hooks ...Hook) {
	c.hooks.Channel = append(c.hooks.Channel, hooks...)
}

// Create returns a create builder for Channel.
func (c *ChannelClient) Create() *ChannelCreate {
	mutation := newChannelMutation(c.config, OpCreate)
	return &ChannelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Channel entities.
func (c *ChannelClient) CreateBulk(builders ...*ChannelCreate) *ChannelCreateBulk {
	return &ChannelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Channel.
func (c *ChannelClient) Update() *ChannelUpdate {
	mutation := newChannelMutation(c.config, OpUpdate)
	return &ChannelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChannelClient) UpdateOne(ch *Channel) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannel(ch))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChannelClient) UpdateOneID(id int64) *ChannelUpdateOne {
	mutation := newChannelMutation(c.config, OpUpdateOne, withChannelID(id))
	return &ChannelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Channel.
func (c *ChannelClient) Delete() *ChannelDelete {
	mutation := newChannelMutation(c.config, OpDelete)
	return &ChannelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ChannelClient) DeleteOne(ch *Channel) *ChannelDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ChannelClient) DeleteOneID(id int64) *ChannelDeleteOne {
	builder := c.Delete().Where(channel.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChannelDeleteOne{builder}
}

// Query returns a query builder for Channel.
func (c *ChannelClient) Query() *ChannelQuery {
	return &ChannelQuery{config: c.config}
}

// Get returns a Channel entity by its id.
func (c *ChannelClient) Get(ctx context.Context, id int64) (*Channel, error) {
	return c.Query().Where(channel.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChannelClient) GetX(ctx context.Context, id int64) *Channel {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChannelClient) Hooks() []Hook {
	return c.hooks.Channel
}

// DatasourceClient is a client for the Datasource schema.
type DatasourceClient struct {
	config
}

// NewDatasourceClient returns a client for the Datasource from the given config.
func NewDatasourceClient(c config) *DatasourceClient {
	return &DatasourceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `datasource.Hooks(f(g(h())))`.
func (c *DatasourceClient) Use(hooks ...Hook) {
	c.hooks.Datasource = append(c.hooks.Datasource, hooks...)
}

// Create returns a create builder for Datasource.
func (c *DatasourceClient) Create() *DatasourceCreate {
	mutation := newDatasourceMutation(c.config, OpCreate)
	return &DatasourceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Datasource entities.
func (c *DatasourceClient) CreateBulk(builders ...*DatasourceCreate) *DatasourceCreateBulk {
	return &DatasourceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Datasource.
func (c *DatasourceClient) Update() *DatasourceUpdate {
	mutation := newDatasourceMutation(c.config, OpUpdate)
	return &DatasourceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DatasourceClient) UpdateOne(d *Datasource) *DatasourceUpdateOne {
	mutation := newDatasourceMutation(c.config, OpUpdateOne, withDatasource(d))
	return &DatasourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DatasourceClient) UpdateOneID(id int64) *DatasourceUpdateOne {
	mutation := newDatasourceMutation(c.config, OpUpdateOne, withDatasourceID(id))
	return &DatasourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Datasource.
func (c *DatasourceClient) Delete() *DatasourceDelete {
	mutation := newDatasourceMutation(c.config, OpDelete)
	return &DatasourceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DatasourceClient) DeleteOne(d *Datasource) *DatasourceDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DatasourceClient) DeleteOneID(id int64) *DatasourceDeleteOne {
	builder := c.Delete().Where(datasource.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DatasourceDeleteOne{builder}
}

// Query returns a query builder for Datasource.
func (c *DatasourceClient) Query() *DatasourceQuery {
	return &DatasourceQuery{config: c.config}
}

// Get returns a Datasource entity by its id.
func (c *DatasourceClient) Get(ctx context.Context, id int64) (*Datasource, error) {
	return c.Query().Where(datasource.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DatasourceClient) GetX(ctx context.Context, id int64) *Datasource {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DatasourceClient) Hooks() []Hook {
	return c.hooks.Datasource
}
