package http

import (
	rt2 "rt/internal/rt"
	"rt/models"
)

// RT is implementation of Rt http server.
type RT struct {
	UC rt2.UseCase
}

//GetProduct gets product throuth http
func (r RT) GetProduct() (*models.Product, error) {
	res, err := r.UC.GetProduct()

	if err != nil {
		return nil, err
	}
	return res, nil
}

//GetOffer gets offer throuth http
func (r RT) GetOffer(product *models.Product, conditions []models.Condition) (*models.Offer, error) {
	res, err := r.UC.GetOffer(product, conditions)

	if err != nil {
		return nil, err
	}
	return res, nil
}
