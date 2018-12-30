package gae

import (
	"gae-test-project/src/api"
	"net/http"

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

	apiGroup := g.Group("/api")
	api.InitUserAPI(apiGroup.Group("/users"))
	api.InitCommentAPI(apiGroup.Group("/comments"))
}
