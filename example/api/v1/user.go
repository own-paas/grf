package v1

import (
	"github.com/gin-gonic/gin"
	restful "github.com/sestack/grf"
	"github.com/sestack/grf/example/global"
	"github.com/sestack/grf/example/model"
)

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
