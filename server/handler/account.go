package handler

import (
	"fmt"
	"gae-test-project/model"
	"gae-test-project/store"
	"gae-test-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Accounts is handler bundle
func Accounts(g *gin.RouterGroup) {
	g.POST("", insertAccount)
	g.GET("/:id", getAccount)
	g.GET("", listAccount)
	g.PUT("/:id", updateAccount)

	g.POST("/login", login)
	g.POST("/logout", logout)
}

func insertAccount(ctx *gin.Context) {
	a := new(model.Account)
	if err := ctx.Bind(a); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Account{}).Insert(ctx.Request.Context(), a); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, a)
}

func getAccount(ctx *gin.Context) {
	accountID := ctx.Param("id")

	account, err := (&store.Account{}).Get(ctx.Request.Context(), accountID)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func listAccount(ctx *gin.Context) {
	accounts, err := (&store.Account{}).List(ctx.Request.Context())
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

func updateAccount(ctx *gin.Context) {
	account := new(model.Account)
	if err := ctx.Bind(account); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	accountID := ctx.Param("id")
	if accountID != account.ID {
		err := fmt.Errorf("invalid id. paramID = %s, bodyId = %s", accountID, account.ID)
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Account{}).Update(ctx.Request.Context(), account); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func login(ctx *gin.Context) {
	reqAccount := new(model.Account)
	if err := ctx.Bind(reqAccount); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	a, err := (&store.Account{}).Get(ctx, reqAccount.ID)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, "ID か パスワード が間違っています。")
		return
	}

	if reqAccount.Password != a.Password {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, "ID か パスワード が間違っています。")
		return
	}

	ctx.String(http.StatusOK, "OK")
}

func logout(ctx *gin.Context) {
	// TODO: セッション管理してないので特にやることがない。同時アクセス不可ぐらいはやってもいいか
	ctx.String(http.StatusOK, "OK")
}
