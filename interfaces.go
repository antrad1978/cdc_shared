package cdc_shared

import (
	"context"
	"time"
)

type ConnectorProvider interface {
	Name() string
	Modes() []string
	InsertRows(connector Connector, rows []map[string]interface{}) int
	MoveData(sync Sync, ctx context.Context)
}

type DatabaseConnectorProvider interface {
	Name() string
	InsertRows(connector Connector, rows []map[string]interface{}) int
	GetRowsById(connector Connector, lastId int64) ([]map[string]interface{}, int64)
	GetMaxTableId(connector Connector) int64
	GetMaxTimestamp(connector Connector) (time.Time, error)
	GetRecordsByTimestamp(connector Connector, lastTimestamp time.Time) ([]map[string]interface{}, time.Time)
}
