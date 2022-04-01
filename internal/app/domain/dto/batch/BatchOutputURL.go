package batch

type BatchOutputURL struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func NewBatchOutputURL(correlationID, shortURL string) BatchOutputURL {
	return BatchOutputURL{
		CorrelationID: correlationID,
		ShortURL:      shortURL,
	}
}
