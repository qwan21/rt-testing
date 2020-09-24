package rt

import (
	"rt/models"
)

type Repository interface {
	GetProduct() (*models.Product, error)
}
