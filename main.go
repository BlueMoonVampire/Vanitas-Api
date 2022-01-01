package main

import (
	"syl-api/bot"
	"syl-api/db"
	"syl-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	db.Database()

	r.GET("/user/:user", routes.Get_User)
	r.GET("/", routes.Home)
	bot.BotInit()
	r.Run()
}
