package rt

import (
	"rt/models"
)

type UseCase interface {
	GetProduct() (*models.Product, error)
	GetOffer(product *models.Product, conditions []models.Condition) (*models.Offer, error)
}
