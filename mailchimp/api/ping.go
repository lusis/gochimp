package api

// PingResponse is a ping response
/*
{
  "health_status": "Everything's Chimpy!"
}
*/
type PingResponse struct {
	HealthStatus string `json:"health_status"`
}
