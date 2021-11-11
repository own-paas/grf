package grf

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"time"
)

const default_timeout = time.Second * 3

type ModelViewSet struct {
	QuerySet       *gorm.DB
	Serializer     interface{}
	DisplayFields  []string
	FilterFields   []string
	SearchFields   []string
	OrderingFields []string
	Preload        string
	Unscoped       bool
	Timeout        time.Duration
}

func (self *ModelViewSet) getQuerySet() *gorm.DB {
	if self.Timeout == 0 {
		self.Timeout = default_timeout
	}
	ctx, _ := context.WithTimeout(context.Background(), self.Timeout)
	return self.QuerySet.WithContext(ctx).Model(self.Serializer)
}

func (self *ModelViewSet) List(c *gin.Context) {
	catchException(c)

	var err error
	var results = []map[string]interface{}{}

	page, size, count := Pagination(c)
	filterMap := Filter(c, self.FilterFields)
	searchMap := Search(c, self.SearchFields)
	ordering := Ordering(c, self.OrderingFields)

	tx := self.getQuerySet()

	if self.Preload != "" {
		tx.Preload(self.Preload)
	}

	tx.Preload(clause.Associations)

	if self.DisplayFields != nil && len(self.DisplayFields) > 0 {
		tx.Select(self.DisplayFields)
	}

	if filterMap != nil && len(filterMap) > 0 {
		tx.Where(filterMap)
	}

	if searchMap != nil && len(searchMap) > 0 {
		for k, v := range searchMap {
			tx.Or(k, v)
		}
	}

	tx.Count(&count)

	if size != 0 {
		tx.Limit(size).Offset(page)
	}

	if ordering != "" {
		tx.Order(ordering)
	}

	if err = tx.Find(&results).Error; err != nil {
		ErrorData(c, err)
		return
	}

	SuccessData(c, gin.H{"results": results, "count": count})
}

func (self *ModelViewSet) Retrieve(c *gin.Context) {
	catchException(c)

	var err error
	var result = map[string]interface{}{}
	id := c.Param("id")

	tx := self.getQuerySet()

	if self.Preload != "" {
		tx.Preload(self.Preload)
	}

	tx.Preload(clause.Associations)

	if self.DisplayFields != nil && len(self.DisplayFields) > 0 {
		tx.Select(self.DisplayFields)
	}

	if err = tx.First(&result, id).Error; err != nil {
		NotFound(c)
		return
	}

	SuccessData(c, result)
}

func (self *ModelViewSet) Update(c *gin.Context) {
	catchException(c)

	var err error
	id := c.Param("id")
	result := map[string]interface{}{}

	tx := self.getQuerySet()

	if err = c.ShouldBind(&result); err != nil {
		ErrorData(c, err)
		return
	}

	if err = tx.Where("id = ?", id).Updates(result).Error; err != nil {
		ErrorData(c, err)
		return
	}

	Success(c)
}

func (self *ModelViewSet) Delete(c *gin.Context) {
	catchException(c)

	var err error
	id := c.Param("id")
	result := reflect.New(reflect.TypeOf(self.Serializer).Elem()).Interface()

	tx := self.getQuerySet()

	if err = tx.First(result, id).Error; err != nil {
		NotFound(c)
		return
	}

	if self.Unscoped {
		tx.Unscoped()
	}

	if err = tx.Delete(result).Error; err != nil {
		ErrorData(c, err)
		return
	}

	Success(c)
}

func (self *ModelViewSet) Create(c *gin.Context) {
	catchException(c)

	var err error
	result := reflect.New(reflect.TypeOf(self.Serializer).Elem()).Interface()

	if err = c.ShouldBind(result); err != nil {
		ErrorData(c, err)
		return
	}

	tx := self.getQuerySet()

	if err = tx.Create(result).Error; err != nil {
		ErrorData(c, err)
		return
	}

	Success(c)
}
