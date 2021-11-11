package grf

import (
	"github.com/gin-gonic/gin"
)

func catchException(c *gin.Context) {
	// 捕获异常
	defer func() {
		err := recover()
		if err != nil {
			InternalServerError(c)
		}
	}()
}
