package storage

import (
	"fastbingo/models"
)

// PersistanceContract
type PersistanceContract interface {
	Migrate()
	InsertProduct(*models.Product)
	FindProductsBy(map[string]interface{}) []*models.Product
	InsertTranslations(fieldsentities map[string]interface{})
}
