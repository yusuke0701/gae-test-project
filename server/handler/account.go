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

// Accounts is handler bundle
func Accounts(g *gin.RouterGroup) {
	g.POST("", insertAccount)
	g.GET("/:account_id", getAccount)
	g.GET("", listAccount)
	g.PUT("/:account_id", updateAccount)
	g.DELETE("/:account_id", deleteAccount)

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
		switch err.(type) {
		case *errs.ErrConflict:
			ctx.String(http.StatusConflict, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, a)
}

func getAccount(ctx *gin.Context) {
	accountID := paramParser.accountID(ctx)

	account, err := (&store.Account{}).Get(ctx.Request.Context(), accountID)
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

	accountID := paramParser.accountID(ctx)
	if accountID != account.ID {
		err := fmt.Errorf("invalid id. paramID = %s, bodyId = %s", accountID, account.ID)
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := (&store.Account{}).Update(ctx.Request.Context(), account); err != nil {
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
	ctx.JSON(http.StatusOK, account)
}

func deleteAccount(ctx *gin.Context) {
	accountID := paramParser.accountID(ctx)

	if err := (&store.Account{}).Delete(ctx.Request.Context(), accountID); err != nil {
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

func login(ctx *gin.Context) {
	reqLA := new(model.LoginAccount)
	if err := ctx.Bind(reqLA); err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	a, err := (&store.Account{}).Get(ctx, reqLA.ID)
	if err != nil {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, "ID か パスワード が間違っています。")
		return
	}

	if reqLA.Password != a.Password {
		util.LogError(ctx.Request.Context(), err.Error)
		ctx.String(http.StatusInternalServerError, "ID か パスワード が間違っています。")
		return
	}

	ctx.JSON(http.StatusOK, a)
}

func logout(ctx *gin.Context) {
	// TODO: セッション管理してないので特にやることがない。同時アクセス不可ぐらいはやってもいいか
	ctx.String(http.StatusOK, "OK")
}
