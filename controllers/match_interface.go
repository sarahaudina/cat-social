package controllers

import "github.com/gin-gonic/gin"

type MatchControllerInterface interface {
 CreateMatch(g *gin.Context)
 GetMatch(g *gin.Context)
 ApproveMatch(g *gin.Context)
 RejectMatch(g *gin.Context)
 DeleteMatch(g *gin.Context)
}