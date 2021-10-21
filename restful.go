package drf

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
)

type ModelViewSet struct {
	QuerySet        *gorm.DB
	Serializer      interface{}
	FiltersetFields []string
	OrderingFields  []string
	Unscoped        bool
}

func (self *ModelViewSet) List(c *gin.Context) {
	catchException(c)

	results := reflect.New(reflect.SliceOf(reflect.TypeOf(self.Serializer))).Interface()

	page, size, count := Pagination(c)
	filterMap := Filter(c, self.FiltersetFields)
	ordering := Ordering(c, self.OrderingFields)

	if err := self.QuerySet.Model(self.Serializer).Where(filterMap).Count(&count).Limit(size).Offset(page).Order(ordering).Find(results).Error; err != nil {
		ErrorData(c, err)
		return
	}

	SuccessData(c, gin.H{"results": results, "count": count})
}

func (self *ModelViewSet) Retrieve(c *gin.Context) {
	catchException(c)
	result := reflect.New(reflect.TypeOf(self.Serializer).Elem()).Interface()

	id := c.Param("id")
	if err := self.QuerySet.Model(self.Serializer).First(result, id).Error; err != nil {
		NotFound(c)
		return
	}

	SuccessData(c, result)
	return
}

func (self *ModelViewSet) Put(c *gin.Context) {
	catchException(c)
	result := reflect.New(reflect.TypeOf(self.Serializer).Elem()).Interface()

	id := c.Param("id")
	if err := c.ShouldBind(result); err != nil {
		ErrorData(c, err)
		return
	}
	if err := self.QuerySet.Model(self.Serializer).Where("id = ?", id).Save(result).Error; err != nil {
		ErrorData(c, err)
		return
	}

	SuccessData(c, result)
	return
}

func (self *ModelViewSet) Delete(c *gin.Context) {
	catchException(c)

	var err error
	result := reflect.New(reflect.TypeOf(self.Serializer).Elem()).Interface()

	id := c.Param("id")
	if self.Unscoped {
		err = self.QuerySet.Model(self.Serializer).First(result, id).Unscoped().Delete(result).Error
	} else {
		err = self.QuerySet.Model(self.Serializer).First(result, id).Delete(result).Error
	}

	if err != nil {
		ErrorData(c, err)
		return
	}

	Success(c)
	return
}

func (self *ModelViewSet) Create(c *gin.Context) {
	catchException(c)
	result := reflect.New(reflect.TypeOf(self.Serializer).Elem()).Interface()

	if err := c.ShouldBind(result); err != nil {
		ErrorData(c, err)
		return
	}

	if err := self.QuerySet.Model(self.Serializer).Create(result).Error; err != nil {
		ErrorData(c, err)
		return
	}

	SuccessData(c, result)
}
