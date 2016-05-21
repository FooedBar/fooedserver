package models

type CreateSessionRequest struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
