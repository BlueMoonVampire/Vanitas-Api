package routes

import (
	"strconv"
	"syl-api/db"

	"github.com/gin-gonic/gin"
)

func Get_User(c *gin.Context) {
	user := string(c.Param("user"))
	ok, _ := strconv.Atoi(user)
	result, err := db.GetUser(ok)

	if err != nil {
		c.JSON(200, gin.H{
			"blacklisted": false,
		})
	} else {

		c.JSON(
			200,
			gin.H{
				"user":     result.User,
				"reason":   result.Reason,
				"enforcer": result.Enforcer,
				"message":  result.Message,
				"blacklisted":   true,
			},
		)
	}

}

func Home(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"message": "Welcome to the Syl API",
			"version": "1.0.0",
			"routes":  []string{"/user/:user", "/"},
		},
	)
}

// func Home(c *gin.Context) {

// 	c.HTML(
// 		200,
// 		"index.html",
// 		nil,
// 	)
// }
