package app

import (
 "database/sql"
 _ "github.com/lib/pq"
 "fmt"
 "log"

 "github.com/gin-gonic/gin"

 "psprint/cat/controllers"
 "psprint/cat/middlewares"
)

type App struct {
 DB     *sql.DB
 Routes *gin.Engine
}

func (a *App) CreateConnection(){
 connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
 db, err := sql.Open("postgres", connStr)
 if err != nil {
  log.Fatal(err)
 }
 a.DB = db
}

func (a *App) CreateRoutes() {
 routes := gin.Default()

 // auth
 userController := controllers.NewUserController(a.DB)
 routes.POST("/user/register", userController.CreateUser)
 routes.POST("/user/login", userController.Login)

 // cats
 catController := controllers.NewCatController(a.DB)
 routes.POST("/cat", middlewares.CheckAuth, catController.CreateCat)
 routes.GET("/cat", catController.GetCats)

 // match
 matchController := controllers.NewMatchController(a.DB)
 routes.POST("/match", middlewares.CheckAuth, matchController.CreateMatch)
 routes.GET("/match", middlewares.CheckAuth, matchController.GetMatch)
 routes.DELETE("/match", middlewares.CheckAuth, matchController.DeleteMatch)
 routes.PUT("/match/approve", middlewares.CheckAuth, matchController.ApproveMatch)
 routes.PUT("/match/decline", middlewares.CheckAuth, matchController.RejectMatch)

 a.Routes = routes
}

func (a *App) Run(){
 a.Routes.Run(":8080")
}