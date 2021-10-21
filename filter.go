package drf

import (
	"github.com/gin-gonic/gin"
)

// 过滤器
func Filter(c *gin.Context, filterset_fields []string) (filterMap map[string]interface{}) {
	filterMap = map[string]interface{}{}

	for _, v := range filterset_fields {
		f := string(c.DefaultQuery(v, ""))
		if f != "" {
			filterMap[v] = f
		}
	}

	return
}
