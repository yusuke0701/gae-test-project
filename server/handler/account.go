package handler

import (
	"gae-test-project/model"
	"gae-test-project/store"
	"gae-test-project/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Accounts is handler bundle
func Accounts(g *gin.RouterGroup) {
	g.POST("/login", login)
	g.POST("/logout", logout)
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
