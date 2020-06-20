package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gae-test-project/firebase"
	"gae-test-project/model"
	"gae-test-project/util"
)

// Users is handler bundle
func Users(g *gin.RouterGroup) {
	g.POST("", createUser)
	g.GET("/:uid", getUser)
}

func createUser(ctx *gin.Context) {
	req := new(model.User)
	if err := ctx.Bind(req); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	u, err := firebase.CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		err := fmt.Errorf("failed to create user: %v", err)
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func getUser(ctx *gin.Context) {
	uid := paramParser.userID(ctx)

	u, err := firebase.GetUser(ctx, uid)
	if err != nil {
		err := fmt.Errorf("failed to create user: %v", err)
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, u)
}
