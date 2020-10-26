package alerts

import (
	"time"

	"github.com/echoturing/alert/alerts/rules"
)

type AlertStatus uint

const (
	AlertStatusUndefined AlertStatus = 0
	AlertStatusOK        AlertStatus = 1
	AlertStatusPending   AlertStatus = 2
	AlertStatusAlerting  AlertStatus = 3
)

type Status uint

const (
	StatusUndefined Status = 0
	StatusOpen      Status = 1
	StatusClose     Status = 2
)

type Alert struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	Channels    []int64     `json:"channels"`
	Rule        *rules.Rule `json:"rule"`
	AlertStatus AlertStatus `json:"alertStatus"`
	Status      Status      `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
