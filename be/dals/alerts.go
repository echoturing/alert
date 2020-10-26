package dals

import (
	"context"
	"fmt"
	"time"

	"github.com/echoturing/alert/alerts"
)

const (
	tableAlert             = "alert"
	columnID               = "id"
	AlertColumnName        = "name"
	AlertColumnChannels    = "channels"
	AlertColumnRule        = "rule"
	AlertColumnAlertStatus = "alertStatus"
	AlertColumnStatus      = "status"
	AlertColumnCreatedAt   = "createdAt"
	AlertColumnUpdatedAt   = "updatedAt"
)

var insertColumns = []string{
	AlertColumnName,
	AlertColumnChannels,
	AlertColumnRule,
	AlertColumnAlertStatus,
	AlertColumnStatus,
}

var allColumns = append(
	append([]string{
		columnID,
	}, insertColumns...,
	), AlertColumnCreatedAt,
	AlertColumnUpdatedAt)

func listToStrWithQuotes(sList []string) string {
	res := ""
	for _, s := range sList {
		res += "`" + s + "`,"
	}
	if len(res) > 0 {
		res = res[:len(res)-1] // remove last comma

	}
	return res
}

func generatePlaceholders(count int) string {
	placeholders := ""
	for i := 0; i < count; i++ {
		placeholders += "?,"
	}
	if len(placeholders) > 0 {
		placeholders = placeholders[:len(placeholders)-1] // remove last comma
	}
	return placeholders
}

func (i *impl) ListAlerts(ctx context.Context, status alerts.Status, alertStatus alerts.AlertStatus) ([]*alerts.Alert, error) {
	query := fmt.Sprintf("select %s from %s where 1=1", listToStrWithQuotes(allColumns), tableAlert)
	var values []interface{}
	if status != alerts.StatusUndefined {
		query += fmt.Sprintf(" and %d=?", status)
		values = append(values, status)
	}
	if alertStatus != alerts.AlertStatusUndefined {
		query += fmt.Sprintf(" and %s=?", AlertColumnAlertStatus)
		values = append(values, alertStatus)
	}

	rows, err := i.db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, query, values)
	}
	results := make([]*alerts.Alert, 0)
	for rows.Next() {
		a := &alerts.Alert{}
		err := rows.Scan(&a.ID, &a.Name, &a.Channels, &a.Rule, &a.AlertStatus, &a.Status, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%w:%s	%#v", err, query, values)
		}
		results = append(results, a)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, query, values)
	}
	return results, nil
}

func (i *impl) CreateAlert(ctx context.Context, alert *alerts.Alert) (*alerts.Alert, error) {
	statement := fmt.Sprintf("insert into %s (%s) values (%s)", tableAlert, listToStrWithQuotes(insertColumns), generatePlaceholders(len(insertColumns)))
	res, err := i.db.ExecContext(ctx, statement, alert.Name, alert.Channels, alert.Rule, alert.AlertStatus, alert.Status)
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, statement, alert)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, statement, alert)
	}
	alert.ID = id
	alert.CreatedAt = time.Now()
	alert.UpdatedAt = time.Now()
	return alert, nil
}

func mapKeysToUpdateValues(kvs map[string]interface{}) (string, []interface{}) {
	keysRes := ""
	valueRes := make([]interface{}, len(kvs))
	for key := range kvs {
		keysRes = "`" + key + "`=?,"
		valueRes = append(valueRes, kvs[key])
	}
	if len(keysRes) > 0 {
		keysRes = keysRes[:len(keysRes)-1] // remove last comma
	}
	return keysRes, valueRes
}

func (i *impl) UpdateAlert(ctx context.Context, id int64, kvs map[string]interface{}) (int64, error) {
	sets, values := mapKeysToUpdateValues(kvs)
	statement := fmt.Sprintf("update %s set %s	where id=?", tableAlert, sets)

	result, err := i.db.ExecContext(ctx, statement, values...)
	if err != nil {
		return 0, fmt.Errorf("%w:%s	%#v", err, statement, kvs)
	}
	return result.RowsAffected()
}

func (i *impl) GetAlertByID(ctx context.Context, id int64) (*alerts.Alert, error) {
	query := fmt.Sprintf("select %s from %s where id=?", listToStrWithQuotes(allColumns), tableAlert)

	row := i.db.QueryRowContext(ctx, query, id)
	a := &alerts.Alert{}
	err := row.Scan(&a.ID, &a.Name, &a.Channels, &a.Rule, &a.AlertStatus, &a.Status, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%w:%s	%#v", err, query, id)
	}
	return a, nil
}
