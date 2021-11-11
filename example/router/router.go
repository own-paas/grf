package router

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	SwaggerRouter(r, "/swagger")
	UserRouter(r, "/user")
}
