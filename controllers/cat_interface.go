package controllers

import "github.com/gin-gonic/gin"

type CatControllerInterface interface {
 CreateCat(g *gin.Context)
 GetCats(g *gin.Context)
}