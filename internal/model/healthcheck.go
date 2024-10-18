// Package model provides the data models for the learning-golang-api API.
package model

// PingResponse represents the response for the ping endpoint.
type PingResponse struct {
	Message string `json:"message"`
}
