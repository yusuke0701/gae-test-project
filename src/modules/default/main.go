package gae

import (
	"gae-test-project/src/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	g := gin.New()
	g.AppEngine = true

	initAPI(g.Group("/api"))
	api.InitTQAPI(g.Group("/tq"))

	http.Handle("/", g)
}

func initAPI(g *gin.RouterGroup) {
	api.InitUserAPI(g.Group("/users"))
	api.InitCommentAPI(g.Group("/comments"))
}
