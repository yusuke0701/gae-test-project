package main

import (
	"net/http"

	_ "gae-test-project/docs"
	"gae-test-project/src/api"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	g := gin.New()

	initAPI(g)

	http.Handle("/", g)
}

func initAPI(g *gin.Engine) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API
	apiGin := g.Group("/api")
	api.InitUserAPI(apiGin.Group("/users"))
	api.InitCommentAPI(apiGin.Group("/comments"))
}
