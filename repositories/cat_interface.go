package repositories

import "psprint/cat/models"

type CatRepositoryInterface interface {
 InsertCat(post models.InsertCat) bool
 GetCats() []models.Cat
}