package batch

type BatchInputURL struct {
	CorrelationID string `json:"correlation_id"`
	Original_url  string `json:"original_url"`
}

func NewBatchInputURL(correlationID, original_url string) BatchInputURL {
	return BatchInputURL{
		CorrelationID: correlationID,
		Original_url:  original_url,
	}
}
