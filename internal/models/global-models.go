package models

type Response[T any] struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    T              `json:"data"`
	Error   map[string]any `json:"error"`
	Meta    map[string]any `json:"meta"`
}
