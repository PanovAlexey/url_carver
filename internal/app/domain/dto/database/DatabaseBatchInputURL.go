package database

type DatabaseBatchInputURL struct {
	Correlation_id string `json:"correlation_id"`
	Original_url   string `json:"original_url"`
}

func NewDatabaseBatchInputURL(correlation_id, original_url string) DatabaseBatchInputURL {
	return DatabaseBatchInputURL{
		Correlation_id: correlation_id,
		Original_url:   original_url,
	}
}
