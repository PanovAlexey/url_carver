package database

type DatabaseBatchInputURL struct {
	CorrelationID string `json:"correlation_id"`
	Original_url  string `json:"original_url"`
}

func NewDatabaseBatchInputURL(correlationID, original_url string) DatabaseBatchInputURL {
	return DatabaseBatchInputURL{
		CorrelationID: correlationID,
		Original_url:  original_url,
	}
}
