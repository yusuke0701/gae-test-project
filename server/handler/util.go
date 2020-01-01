package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var paramParser *param

type param struct{}

func (p *param) accountID(ctx *gin.Context) string {
	return ctx.Param("account_id")
}

func (p *param) commentID(ctx *gin.Context) (int64, error) {
	idStr := ctx.Param("comment_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *param) tagID(ctx *gin.Context) (int64, error) {
	idStr := ctx.Param("tag_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *param) threadID(ctx *gin.Context) string {
	return ctx.Param("thread_id")
}
