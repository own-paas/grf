# GRF

Gin REST Framework是一个golang gin版RESTful API框架，gin+gorm，简单易上手。

## 获取源码

git clone https://github.com/sestack/grf.git

## 项目结构

```
|-restful.go 模型restful视图结构体
|-respone.go 状态返回
|-filter.go 过滤器
|-ordering.go 排序器
|-pagination.go 分页器
|-exception.go 捕获异常
|-example 示例
```

## REST URL

```
GET http://localhost:8080/v1/user/
GET http://localhost:8080/v1/user/:id/
POST http://localhost:8080/v1/user/
PUT http://localhost:8080/v1/user/:id/
DELETE http://localhost:8080/v1/user/:id/
```

## JSON RESULT

```
{
    "status": 0,
    "msg": "请求成功",
    "data": {}
}
```

## 测试user模型

```
type User struct {
	global.Model
	Name     string `json:"name" gorm:"uniqueIndex;not null;type:varchar(64);comment:用户名"` // 用户名
	Password string `json:"password" gorm:"not null;type:varchar(64);comment:密码"`          // 密码
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:电话号码"`                    // 电话号码
	Email    string `json:"email" gorm:"type:varchar(32);comment:邮箱"`                      // 邮箱
}
```

## API 示例

```
type user struct {
	restful.ModelViewSet
}

func UserView() user {
	return user{
		restful.ModelViewSet{
			QuerySet:       global.DB,
			Serializer:     &model.User{},
			FilterFields:   []string{"name", "phone", "email"},
			SearchFields:   []string{"name", "phone"},
			DisplayFields:  []string{"id", "name", "password", "phone", "email"},
			OrderingFields: []string{"id"},
			Unscoped:       true,
		},
	}
}
```

## ROUTER 示例

```
func Register(r *gin.RouterGroup) {
	SwaggerRouter(r, "/swagger")
	UserRouter(r, "/user")
}

func SwaggerRouter(r *gin.RouterGroup, location string) {
	api := r.Group(location)
	{
		api.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

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
```

## SERVICE 示例

```
func InitHttpServer() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())         // 日志服务
	app.Use(gin.Recovery())       // 日志服务
	app.NoRoute(restful.NotFound) // 处理 404

	appV1 := app.Group("v1")
	router.Register(appV1)

	return app
}
```

## 重写API接口与生成swagger文档

```
type user struct {
	restful.ModelViewSet
}

func UserView() user {
	return user{
		restful.ModelViewSet{
			QuerySet:       global.DB,
			Serializer:     &model.User{},
			FilterFields:   []string{"name", "phone", "email"},
			SearchFields:   []string{"name", "phone"},
			DisplayFields:  []string{"id", "name", "password", "phone", "email"},
			OrderingFields: []string{"id"},
			Unscoped:       true,
		},
	}
}

// @Tags 用户管理
// @Summary 用户列表
// @accept application/json
// @Produce application/json
// @Param page query integer false "分页页面"
// @Param size query integer false "分页大小"
// @Param order query string false "排序"
// @Param search query string false "搜索"
// @Success 200 {object} object "{"code": 0, "data": {"count": 0, "results": []}}"
// @Router /v1/user/ [get]
func (self *user) List(c *gin.Context) {
	self.ModelViewSet.List(c)
}

// @Tags 用户管理
// @Summary 查看用户
// @accept application/json
// @Produce application/json
// @Param id path integer true "用户id"
// @Success 200 {object} object "{"code": 0, "data":{}}"
// @Router /v1/user/{id}/ [get]
func (self *user) Retrieve(c *gin.Context) {
	self.ModelViewSet.Retrieve(c)
}

// @Tags 用户管理
// @Summary 添加用户
// @accept application/json
// @Produce application/json
// @Param user body model.User true "传入参数是struct"
// @Success 200 {object} object "{"code": 0, "message": "操作成功！"}"
// @Router /v1/user/ [post]
func (self *user) Create(c *gin.Context) {
	self.ModelViewSet.Create(c)
}

// @Tags 用户管理
// @Summary 修改用户
// @accept application/json
// @Produce application/json
// @Param id path integer true "用户id"
// @Param user body model.User true "传入参数是struct"
// @Success 200 {object} object "{"code": 0, "message": "操作成功！"}"
// @Router /v1/user/{id}/ [put]
func (self *user) Update(c *gin.Context) {
	self.ModelViewSet.Update(c)
}

// @Tags 用户管理
// @Summary 删除用户
// @accept application/json
// @Produce application/json
// @Param id path integer true "用户id"
// @Success 200 {object} object "{"code": 0, "message": "操作成功！"}"
// @Router /v1/user/{id}/ [delete]
func (self *user) Delete(c *gin.Context) {
	self.ModelViewSet.Delete(c)
}
```