package database

type DatabaseBatchOutputURL struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func NewDatabaseBatchOutputURL(correlationID, shortURL string) DatabaseBatchOutputURL {
	return DatabaseBatchOutputURL{
		CorrelationID: correlationID,
		ShortURL:      shortURL,
	}
}
