package main

import (
	"syl-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/user/:user", routes.Get_User)
	r.GET("/", routes.Home)

	r.Run()
}
