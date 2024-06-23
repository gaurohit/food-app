package health

type HealthRequest struct{}

type HealthResponse struct {
	Status string `json:"status"`
}
