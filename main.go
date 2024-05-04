package main

import (
 "psprint/cat/app"
)

func main() {
 var a app.App
 a.CreateConnection()
 a.CreateRoutes()
 a.Run()
}