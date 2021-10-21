package drf

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// 排序器
func Ordering(c *gin.Context, ordering_fields []string) (order string) {
	for _, o := range ordering_fields {
		if order != "" {
			order = order + ", "
		}
		if strings.HasPrefix(o, "-") {
			order = order + strings.Split(o, "-")[1] + " desc"
		} else {
			order = order + o
		}
	}

	return
}
