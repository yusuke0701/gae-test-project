package handler

import (
	"gae-test-project/model"
	"gae-test-project/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Comments is handler bundle
func Comments(g *gin.RouterGroup) {
	g.POST("", insertComment)
	g.GET("/:id", getComment)
	g.GET("", listComment)
	g.PUT("/:id", updateComment)
}

func insertComment(ctx *gin.Context) {
	var comment *model.Comment
	if err := ctx.Bind(comment); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Comment{}).InsertOrUpadte(ctx.Request.Context(), comment); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func getComment(ctx *gin.Context) {
	commentID := ctx.Param("id")

	comment, err := (&store.Comment{}).Get(ctx.Request.Context(), commentID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func listComment(ctx *gin.Context) {
	comments, err := (&store.Comment{}).List(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func updateComment(ctx *gin.Context) {
	var comment *model.Comment
	if err := ctx.Bind(comment); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	// commentID := ctx.Param("id")
	// TODO: 現状だとInsertと処理の違いがない。取得して値を使いまわす部分があってもよいかな

	if err := (&store.Comment{}).InsertOrUpadte(ctx.Request.Context(), comment); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
