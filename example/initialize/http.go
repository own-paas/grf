package initialize

import (
	"github.com/gin-gonic/gin"
	restful "github.com/sestack/grf"
	"github.com/sestack/grf/example/router"
)

func InitHttpServer() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())         // 日志服务
	app.Use(gin.Recovery())       // 日志服务
	app.NoRoute(restful.NotFound) // 处理 404

	appV1 := app.Group("v1")
	router.Register(appV1)

	return app
}
