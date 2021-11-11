package router

import (
	"github.com/gin-gonic/gin"
	api "github.com/sestack/grf/example/api/v1"
)

func UserRouter(r *gin.RouterGroup, location string) {
	view := api.UserView()
	userApi := r.Group(location)
	{
		userApi.GET("/", view.List)          // 用户列表
		userApi.POST("/", view.Create)       // 添加用户
		userApi.GET("/:id/", view.Retrieve)  // 查看用户
		userApi.PUT("/:id/", view.Update)    // 修改用户
		userApi.DELETE("/:id/", view.Delete) // 删除用户
	}
}
