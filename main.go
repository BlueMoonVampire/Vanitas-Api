package main

import (
	"syl-api/routes"
	"syl-api/db"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	db.Database()

	r.GET("/user/:user", routes.Get_User)
	r.GET("/", routes.Home)

	r.Run()
}
