package repository

import (
	"encoding/json"
	"io/ioutil"
	"rt/models"
)

//RtRepository is ...
type RtRepository struct {
}

// NewRtRepository creates repository of tuner.
func NewRtRepository() *RtRepository {
	return &RtRepository{}
}

//GetProduct gets product from same database
func (r RtRepository) GetProduct() (*models.Product, error) {
	file, err := ioutil.ReadFile("product.json")

	if err != nil {
		return nil, err
	}
	product := models.Product{}

	_ = json.Unmarshal([]byte(file), &product)
	return &product, nil
}
