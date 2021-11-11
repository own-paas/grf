package grf

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// 分页器
func Pagination(c *gin.Context) (page int, size int, count int64) {
	p := string(c.DefaultQuery("page", "1"))
	page, _ = strconv.Atoi(p)
	s := string(c.DefaultQuery("size", "0"))
	size, _ = strconv.Atoi(s)
	count = 0
	return (page - 1) * size, size, count
}