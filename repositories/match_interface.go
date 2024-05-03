package repositories

import "psprint/cat/models"

type MatchRepositoryInterface interface {
 InsertMatch(post models.InsertMatch) bool
 GetMatch() []models.Match
 UpdateMatch(post models.UpdateMatch) bool
 DeleteMatch(matchId uint) bool
}