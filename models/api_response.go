package models

type Response struct {
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Debug   string      `json:"debug,omitempty"`
}
