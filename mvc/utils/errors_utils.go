package utils

type ApplicationError struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Code    string `json:"code,omitempty"`
}
