package grf

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// 排序器
func Ordering(c *gin.Context, ordering_fields []string) string {
	order := string(c.DefaultQuery("order", ""))
	if order != "" {
		if strings.HasPrefix(order, "-") {
			tmp := strings.Split(order, "-")
			if len(tmp) == 2 {
				order = tmp[1] + " desc"
			}
		}
	} else {
		for _, o := range ordering_fields {
			if order != "" {
				order = order + ", "
			}
			if strings.HasPrefix(o, "-") {
				tmp := strings.Split(o, "-")
				if len(tmp) == 2 {
					order = order + tmp[1] + " desc"
				}
			} else {
				order = order + o
			}
		}
	}

	return order
}
