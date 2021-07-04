package main

import (
	"github.com/DenrianWeiss/catball/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", handler.MainPageHandler)
	r.GET("index.html", handler.MainPageHandler)

	r.GET("/:target", handler.Redirect)
	r.GET("/show/:target", handler.ShowRedirect)
	r.POST("/add/:token/:path", handler.AddRedirect)

	r.Run()
}