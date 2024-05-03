package controllers

import (
	"database/sql"

	"psprint/cat/models"
	"psprint/cat/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)


type MatchController struct {
	DB *sql.DB
}
   
func NewMatchController(db *sql.DB) MatchControllerInterface {
	return &MatchController{DB: db}
}

func (m *MatchController) GetMatch(c *gin.Context) {
	db := m.DB
	repo_match := repositories.NewMatchRepository(db)
	get_matchs := repo_match.GetMatch()
	if get_matchs != nil {
	 c.JSON(200, gin.H{"status": "success", "data": get_matchs, "msg": "get manga successfully"})
	} else {
	 c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get manga successfully"})
	}
}

func (m *MatchController) CreateMatch(c *gin.Context) {

	var matchInput models.InsertMatch_

	if err := c.ShouldBindJSON(&matchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract user id
	// userId := c.Value("userId").(float64)
	// todo: use user id to validate if issuer id (cat) belong to userId

	match := models.InsertMatch{
		Status: "waiting",
		Issuer: matchInput.Issuer,
		Receiver: matchInput.Receiver,
	}

	db := m.DB
	repo_match := repositories.NewMatchRepository(db)

	insert := repo_match.InsertMatch(match)

	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
	}
}

func (m *MatchController) DeleteMatch(c *gin.Context) {

	var deleteMatchInput models.DeleteMatch

	if err := c.ShouldBindJSON(&deleteMatchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := m.DB
	repo_match := repositories.NewMatchRepository(db)

	delete := repo_match.DeleteMatch(deleteMatchInput.ID)

	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
	}
}

func (m *MatchController) ApproveMatch(c *gin.Context) {

	var input models.ResponseToMatchRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateMatch := models.UpdateMatch{
		Status: "approved",
		ID: input.ID,
	}

	db := m.DB
	repo_match := repositories.NewMatchRepository(db)
	update := repo_match.UpdateMatch(UpdateMatch)

	if update {
		c.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
	}
}

func (m *MatchController) RejectMatch(c *gin.Context) {

	var input models.ResponseToMatchRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateMatch := models.UpdateMatch{
		Status: "declined",
		ID: input.ID,
	}

	db := m.DB
	repo_match := repositories.NewMatchRepository(db)

	update := repo_match.UpdateMatch(UpdateMatch)

	if update {
		c.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
	}
}