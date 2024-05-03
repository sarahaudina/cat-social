package controllers

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
 CreateUser(g *gin.Context)
 Login(g *gin.Context)
}