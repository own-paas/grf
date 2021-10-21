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
    "data": null
}
```

## 测试user模型

```
type User struct {
	gorm.Model
	Name          string         `json:"name" gorm:"uniqueIndex;not null;type:varchar(64);comment:用户名"`
	Password      string         `json:"password" gorm:"not null;type:varchar(64);comment:密码"`
	Phone         string         `json:"phone" gorm:"type:varchar(11);comment:电话号码"`
	Email         string         `json:"email" gorm:"type:varchar(32);comment:邮箱"`
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
			QuerySet:        global.DB,
			Serializer:      &model.User{},
			FiltersetFields: []string{"name", "phone", "email"},
			OrderingFields:  []string{"id"},
		},
	}
}
```

## ROUTER 示例

```
func UserRouter(r *gin.RouterGroup, location string) {
	view := api.UserView()
	api := r.Group(location)
	{
		api.GET("/", view.List)          // 用户列表
		api.POST("/", view.Create)       // 添加用户
		api.GET("/:id/", view.Retrieve)  // 查看用户
		api.PUT("/:id/", view.Put)       // 修改用户
		api.DELETE("/:id/", view.Delete) // 删除用户
	}
}
```

## SERVICE 示例

```
	app := gin.New()
	appV1 := app.Group("v1")
	router.UserRouter(appV1,"/user")

	app.run("127.0.0.1:8080")
```

## 重写API接口与生成swagger文档

```
type user struct {
	restful.ModelViewSet
}

func UserView() user {
	return user{
		restful.ModelViewSet{
			QuerySet:        global.DB,
			Serializer:      &model.User{},
			FiltersetFields: []string{"name", "phone", "email"},
			OrderingFields:  []string{"id"},
		},
	}
}

/ @Tags 用户管理
// @Summary 用户列表
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code": 0, "data":[]}"
// @Router /v1/user/ [get]
func (self *user) List(c *gin.Context) {
	self.ModelViewSet.List(c)
}

// @Tags 用户管理
// @Summary 查看用户
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code": 0, "data":{}}"
// @Router /v1/user/:id/ [get]
func (self *user) Retrieve(c *gin.Context) {
	self.ModelViewSet.Retrieve(c)
}

// @Tags 用户管理
// @Summary 添加用户
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code": 0, "data":{}}"
// @Router /v1/user/ [post]
func (self *user) Create(c *gin.Context) {
	self.ModelViewSet.Create(c)
}

// @Tags 用户管理
// @Summary 修改用户
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code": 0, "data":{}}"
// @Router /v1/user/:id/ [put]
func (self *user) Put(c *gin.Context) {
	self.ModelViewSet.Put(c)
}

// @Tags 用户管理
// @Summary 删除用户
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code": 0, "data":{}}"
// @Router /v1/user/:id/ [delete]
func (self *user) Delete(c *gin.Context) {
	self.ModelViewSet.Delete(c)
}
```