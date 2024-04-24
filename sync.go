package cdc_shared

type Sync struct {
	Id                   string    `json:"id"`
	SyncName             string    `json:"sync_name" binding:"required"`
	SourceConnector      Connector `json:"source_connector"`
	DestinationConnector Connector `json:"destination_connector"`
	Mode                 string    `json:"mode"`
	Disabled             bool      `json:"disabled"`
}

func (mt Sync) Equals(other Sync) bool {
	return mt.Id == other.Id
}

func (mt Sync) NotEquals(other Sync) bool {
	return mt.Id != other.Id
}

type Connector struct {
	IdField              string            `json:"id"`
	ConnectorName        string            `json:"connector_name" binding:"required"`
	ConnectorType        string            `json:"connector_source_type" binding:"required"`
	ConnectionString     string            `json:"connection_string" binding:"required"`
	Database             string            `json:"database"`
	Query                string            `json:"query"`
	Table                string            `json:"table""`
	PollingTime          int               `json:"polling_time" default:"60"`
	TimestampField       string            `json:"timestamp_field"`
	TimestampFieldFormat string            `json:"timestamp_field_format"`
	MaxRecordBatchSize   int               `json:"max_record_batch_size" default:"50000"`
	SaveMode             string            `json:"save_mode"`
	StartOffset          int               `json:"start_offset"`
	Token                string            `json:"token"`
	Attributes           map[string]string `json:"attributes"`
}

type SyncStatus struct {
	SyncId                 string `json:"connector_id"`
	LastError              string `json:"last_error"`
	LastErrorTimestamp     string `json:"last_error_timestamp"`
	CurrentOffset          int64  `json:"current_offset"`
	LastTimestampProcessed string `json:"last_timestamp_processed"`
	MaxRowsBatchSize       int    `json:"max_rows_batch_size"`
}

type SyncLock struct {
	SyncId          string `json:"connector_id"`
	WorkerProcessId string `json:"worker_process_id"`
}
