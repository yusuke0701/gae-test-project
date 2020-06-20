package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yusuke0701/gae-test-project/util"
)

const bucketName = "asada-test"

// SignedURLs is handler bundle
func SignedURLs(g *gin.RouterGroup) {
	g.GET("/csv-download/:file_name", getDownloadURL)
	g.GET("/csv-upload/:file_name", getCSVUploadURL)

	g.GET("/pdf-downlaod/:file_name", getDownloadURL)
	g.GET("/pdf-upload/:file_name", getPDFUploadURL)
}

func getDownloadURL(ctx *gin.Context) {
	fileName := ctx.Param("file_name")

	url, err := util.GetSignedURL(ctx, util.SignedURLTypeGET, bucketName, fileName, "")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(http.StatusOK, url)
}

func getCSVUploadURL(ctx *gin.Context) {
	fileName := ctx.Param("file_name")

	url, err := util.GetSignedURL(ctx, util.SignedURLTypePUT, bucketName, fileName, "text/csv")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(http.StatusOK, url)
}

func getPDFUploadURL(ctx *gin.Context) {
	fileName := ctx.Param("file_name")

	url, err := util.GetSignedURL(ctx, util.SignedURLTypePUT, bucketName, fileName, "application/pdf")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(http.StatusOK, url)
}
