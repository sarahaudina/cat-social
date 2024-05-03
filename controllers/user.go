package controllers

import (
	"database/sql"

	"psprint/cat/models"
	"psprint/cat/repositories"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

)


type UserController struct {
	DB *sql.DB
}
   
func NewUserController(db *sql.DB) UserControllerInterface {
	return &UserController{DB: db}
}

func (m *UserController) CreateUser(c *gin.Context) {

	var authInput models.InsertUser

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.InsertUser{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	db := m.DB
	repo_user := repositories.NewUserRepository(db)

	userFound := repo_user.FindUser(user.Username)

	if userFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	insert := repo_user.InsertUser(user)

	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
	}
}

func (m *UserController) Login(c *gin.Context) {

	var authInput models.InsertUser

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := m.DB
	repo_user := repositories.NewUserRepository(db)
	get_user := repo_user.FindUser(authInput.Username)

	if get_user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(get_user.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  get_user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}