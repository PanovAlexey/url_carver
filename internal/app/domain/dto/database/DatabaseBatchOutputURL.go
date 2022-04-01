package database

type DatabaseBatchOutputURL struct {
	Correlation_id string `json:"correlation_id"`
	Short_url      string `json:"short_url"`
}

func NewDatabaseBatchOutputURL(correlation_id, short_url string) DatabaseBatchOutputURL {
	return DatabaseBatchOutputURL{
		Correlation_id: correlation_id,
		Short_url:      short_url,
	}
}
