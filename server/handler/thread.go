package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yusuke0701/gae-test-project/db/datastore/model"
	"github.com/yusuke0701/gae-test-project/db/datastore/store"
	errs "github.com/yusuke0701/goutils/error"
	"github.com/yusuke0701/goutils/gcp"
)

// Threads is handler bundle
func Threads(g *gin.RouterGroup) {
	g.POST("", insertThread)
	g.GET("/:thread_id", getThread)
	g.GET("", listThread)
	g.PUT("/:thread_id", updateThread)
	g.DELETE("/:thread_id", deleteThread)
}

func insertThread(ctx *gin.Context) {
	thread := new(model.Thread)
	if err := ctx.Bind(thread); err != nil {
		gcp.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Thread{}).Insert(ctx.Request.Context(), thread); err != nil {
		gcp.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *errs.ErrConflict:
			ctx.String(http.StatusConflict, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, thread)
}

func getThread(ctx *gin.Context) {
	threadID := paramParser.threadID(ctx)

	thread, err := (&store.Thread{}).Get(ctx.Request.Context(), threadID)
	if err != nil {
		gcp.LogError(ctx.Request.Context(), err.Error)
		switch err.(type) {
		case *errs.ErrNotFound:
			ctx.String(http.StatusNotFound, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, thread)
}

func listThread(ctx *gin.Context) {
	threads, err := (&store.Thread{}).List(ctx.Request.Context())
	if err != nil {
		gcp.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, threads)
}

func updateThread(ctx *gin.Context) {
	thread := new(model.Thread)
	if err := ctx.Bind(thread); err != nil {
		gcp.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	threadID := paramParser.threadID(ctx)
	if threadID != thread.ID {
		err := fmt.Errorf("invalid id. paramID = %s, bodyId = %s", threadID, thread.ID)
		gcp.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Thread{}).Update(ctx.Request.Context(), thread); err != nil {
		gcp.LogError(ctx.Request.Context(), err.Error)
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
	ctx.JSON(http.StatusOK, thread)
}

func deleteThread(ctx *gin.Context) {
	threadID := paramParser.threadID(ctx)

	if err := (&store.Thread{}).Delete(ctx.Request.Context(), threadID); err != nil {
		gcp.LogError(ctx.Request.Context(), err.Error)
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
