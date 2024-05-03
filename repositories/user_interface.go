package repositories

import "psprint/cat/models"

type UserRepositoryInterface interface {
 InsertUser(post models.InsertUser) bool
 FindUser(username string) *models.User
}