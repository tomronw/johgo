package coreModels

type ApiEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HealthCheckResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"Error"`
}
