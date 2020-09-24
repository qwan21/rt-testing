package models

type Product struct {
	Name       string      `json:"name"`
	Components []Component `json:"components"`
}
