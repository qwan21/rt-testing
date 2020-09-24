package models

type Offer struct {
	Product
	TotalCost Price `json:"totalCost"`
}
