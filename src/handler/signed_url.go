package handler

import (
	"gae-test-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const bucketName = "asada-test"

// SignedURLs is handler bundle
func SignedURLs(g *gin.RouterGroup) {
	g.GET("/:file_name", getDownloadURL)
}

func getDownloadURL(ctx *gin.Context) {
	fileName := ctx.Param("file_name")

	url, err := util.GetSignedURL(ctx, util.SignedURLTypeGET, bucketName, fileName, "text/csv")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(http.StatusOK, url)
}
