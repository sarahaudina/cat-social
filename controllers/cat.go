package controllers

import (
	"database/sql"

	"psprint/cat/models"
	"psprint/cat/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)


type CatController struct {
	DB *sql.DB
}
   
func NewCatController(db *sql.DB) CatControllerInterface {
	return &CatController{DB: db}
}

func (m *CatController) GetCats(g *gin.Context) {
	db := m.DB
	repo_cat := repositories.NewCatRepository(db)
	get_cats := repo_cat.GetCats()
	if get_cats != nil {
	 g.JSON(200, gin.H{"status": "success", "data": get_cats, "msg": "get manga successfully"})
	} else {
	 g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get manga successfully"})
	}
}

func (m *CatController) CreateCat(c *gin.Context) {

	var catInput models.InsertCat_

	if err := c.ShouldBindJSON(&catInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract user id
	userId := c.Value("userId").(float64)
	// if (userId==nil) {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthenticated"})
	// }

	cat := models.InsertCat{
		Name: catInput.Name,
		Gender: catInput.Gender,
		UserId: uint(userId),
	}

	db := m.DB
	repo_cat := repositories.NewCatRepository(db)

	insert := repo_cat.InsertCat(cat)

	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
	}
}