package models

type Component struct {
	Name   string  `json:"name"`
	IsMain bool    `json:"isMain,omitempty"`
	Prices []Price `json:"prices"`
}
