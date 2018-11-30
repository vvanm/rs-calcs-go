package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/vvanm/rs-calcs-go/lookup"
	"github.com/vvanm/rs-calcs-go/prices"
	"github.com/vvanm/rs-calcs-go/raven"
	"github.com/vvanm/rs-calcs-go/transcript"
	"github.com/vvanm/rs-calcs-go/user"
	"os"
)

func main() {
	raven.SetupStore()

	//set default port for local work
	port := "4000"

	//if prod, get port from env
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	//Set up gin
	r := gin.New()

	//setup logging
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//Add cors
	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET,PUT,POST,DELETE,PATCH",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		Credentials:    true,
	}))

	//No route handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"errorMsg": "routeNotFound"})
	})

	r.GET("/auth", user.Auth)
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.GET("/logout", user.Logout)

	//get prices endpoint
	r.GET("/prices/:category", prices.Controller)

	//lookup
	r.GET("/lookup/:rsn", lookup.ByName)

	r.POST("/transcripts", transcript.Create)
	r.GET("/transcripts", transcript.Search)
	r.GET("/transcripts/:id", transcript.Load)
	r.PATCH("/transcripts/:id", transcript.Update)

	r.Run(":" + port)

}
