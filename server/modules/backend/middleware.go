package main

import (
	"gae-test-project/util"

	"github.com/gin-gonic/gin"
)

func setEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		util.SetTraceID(r)

		c.Next()
	}
}
