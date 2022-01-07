package main

import (
	"syl-api/db"
	"syl-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
// 	r.LoadHTMLGlob("templates/*")
	r.GET("/" , routes.Home)
	db.Database()

	r.GET("/user/:user", routes.Get_User)
// 	r.GET("/", routes.Home)
	r.Run()
}
