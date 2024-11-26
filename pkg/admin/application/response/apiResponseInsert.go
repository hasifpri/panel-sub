package adminapplicationresponse

import "time"

type ApiResponseInsert struct {
	CorrelationID string               `json:"correlationid"`
	Success       bool                 `json:"success"`
	Error         string               `json:"error"`
	Tin           time.Time            `json:"tin"`
	Tout          time.Time            `json:"tout"`
	Latency       string               `json:"latency"`
	Data          *CreateAdminResponse `json:"data"`
}
