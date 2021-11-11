package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sestack/grf/example/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SwaggerRouter(r *gin.RouterGroup, location string) {
	api := r.Group(location)
	{
		api.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
