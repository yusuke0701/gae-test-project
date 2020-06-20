package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yusuke0701/gae-test-project/model"
	"github.com/yusuke0701/gae-test-project/store"
	"github.com/yusuke0701/gae-test-project/util"
	errs "github.com/yusuke0701/goutils/error"
)

// Tags is handler bundle
func Tags(g *gin.RouterGroup) {
	g.POST("", insertTag)
	g.GET("/:tag_id", getTag)
	g.GET("", listTag)
	g.PUT("/:tag_id", updateTag)
	g.DELETE("/:tag_id", deleteTag)
}

func insertTag(ctx *gin.Context) {
	tag := new(model.Tag)
	if err := ctx.Bind(tag); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Tag{}).Insert(ctx.Request.Context(), tag); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *errs.ErrConflict:
			ctx.String(http.StatusConflict, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, tag)
}

func getTag(ctx *gin.Context) {
	tagID, err := paramParser.tagID(ctx)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	tag, err := (&store.Tag{}).Get(ctx.Request.Context(), tagID)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *errs.ErrNotFound:
			ctx.String(http.StatusNotFound, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, tag)
}

func listTag(ctx *gin.Context) {
	tags, err := (&store.Tag{}).List(ctx.Request.Context())
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, tags)
}

func updateTag(ctx *gin.Context) {
	tag := new(model.Tag)
	if err := ctx.Bind(tag); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	tagID, err := paramParser.tagID(ctx)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if tagID != tag.ID {
		err := fmt.Errorf("invalid id. paramID = %d, bodyId = %d", tagID, tag.ID)
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Tag{}).Update(ctx.Request.Context(), tag); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *errs.ErrNotFound:
			ctx.String(http.StatusNotFound, err.Error())
		case *errs.ErrConflict:
			ctx.String(http.StatusConflict, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, tag)
}

func deleteTag(ctx *gin.Context) {
	tagID, err := paramParser.tagID(ctx)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Tag{}).Delete(ctx.Request.Context(), tagID); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *errs.ErrNotFound:
			ctx.String(http.StatusNotFound, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.Status(http.StatusOK)
}
