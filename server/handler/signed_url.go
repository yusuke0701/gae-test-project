package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yusuke0701/goutils/gcp/gcs"
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

	url, err := gcs.GetSignedURL(ctx, bucketName, fileName, "", gcs.SignedURLTypeGET, time.Now().Add(15*time.Minute))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(http.StatusOK, url)
}

func getCSVUploadURL(ctx *gin.Context) {
	fileName := ctx.Param("file_name")

	url, err := gcs.GetSignedURL(ctx, bucketName, fileName, "text/csv", gcs.SignedURLTypePUT, time.Now().Add(15*time.Minute))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(http.StatusOK, url)
}

func getPDFUploadURL(ctx *gin.Context) {
	fileName := ctx.Param("file_name")

	url, err := gcs.GetSignedURL(ctx, bucketName, fileName, "application/pdf", gcs.SignedURLTypePUT, time.Now().Add(15*time.Minute))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	ctx.String(http.StatusOK, url)
}
