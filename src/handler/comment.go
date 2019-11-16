package handler

import (
	"gae-test-project/model"
	"gae-test-project/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Comments is handler bundle
func Comments(g *gin.RouterGroup) {
	g.PUT("", insertComment)
	g.GET("/:id", getComment)
}

func insertComment(ctx *gin.Context) {
	body, ok := ctx.GetQuery("body")
	if !ok {
		ctx.String(http.StatusBadRequest, "body is required")
		return
	}
	comment := &model.Comment{ID: "10", Body: body}

	if err := store.Comment.Insert(ctx.Request.Context(), comment); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func getComment(ctx *gin.Context) {
	commentID := ctx.Param("id")

	comment, err := store.Comment.Get(ctx.Request.Context(), commentID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
