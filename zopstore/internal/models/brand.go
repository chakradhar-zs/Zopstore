package models

// Brand structure Definition
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}
