package grf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回成功
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": []string{},
		"message": "操作成功！",
	})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

// 返回失败
func Error(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": "操作失败！",
	})
}

func ErrorData(c *gin.Context, data interface{}) {
	switch v := data.(type) {
	case string:
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": v,
		})
	case error:
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": v.Error(),
		})
	}
}

// 404
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    404,
		"message": "未找到！",
	})
}

// 401
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    401,
		"message": "认证未授权！",
	})
}

// 403
func NotForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"code":    403,
		"message": "无权限访问！",
	})
}

// 500
func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": "服务器异常！",
	})
}
