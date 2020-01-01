package handler

import (
	"fmt"
	"gae-test-project/model"
	"gae-test-project/store"
	"gae-test-project/util"
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
	comment := new(model.Comment)
	if err := ctx.Bind(comment); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Comment{}).Insert(ctx.Request.Context(), comment); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *util.ErrConflict:
			ctx.String(http.StatusConflict, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func getComment(ctx *gin.Context) {
	commentID := ctx.Param("id")

	comment, err := (&store.Comment{}).Get(ctx.Request.Context(), commentID)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *util.ErrNotFound:
			ctx.String(http.StatusNotFound, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func listComment(ctx *gin.Context) {
	comments, err := (&store.Comment{}).List(ctx.Request.Context())
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func updateComment(ctx *gin.Context) {
	comment := new(model.Comment)
	if err := ctx.Bind(comment); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	commentID := ctx.Param("id")
	if commentID != comment.ID {
		err := fmt.Errorf("invalid id. paramID = %s, bodyId = %s", commentID, comment.ID)
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Comment{}).Update(ctx.Request.Context(), comment); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *util.ErrNotFound:
			ctx.String(http.StatusNotFound, err.Error())
		case *util.ErrConflict:
			ctx.String(http.StatusConflict, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
