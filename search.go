package grf

import (
	"github.com/gin-gonic/gin"
)

// 搜索器
func Search(c *gin.Context, search_fields []string) map[string]string {
	conditions := map[string]string{}
	search := string(c.DefaultQuery("search", ""))

	if search != "" && search_fields != nil {
		for _, field := range search_fields {
			conditions[field+" like ?"] = "%" + search + "%"
		}
	}

	return conditions
}
