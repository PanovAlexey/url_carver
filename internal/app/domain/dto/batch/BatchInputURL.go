package batch

type BatchInputURL struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

func NewBatchInputURL(correlationID, originalURL string) BatchInputURL {
	return BatchInputURL{
		CorrelationID: correlationID,
		OriginalURL:   originalURL,
	}
}
