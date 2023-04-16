package models

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Repo struct {
	sku []Sku
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(sku Sku) {
	r.sku = append(r.sku, sku)
}

func (r *Repo) GetAll() []Sku {
	return r.sku
}

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *MockDB {
	m.Called(query, args)
	return m
}

func (m *MockDB) First(out interface{}, where ...interface{}) *MockDB {
	m.Called(out, where)
	return m
}

func (m *MockDB) Save(value interface{}) *MockDB {
	m.Called(value)
	return m
}

type DB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type SkuMock interface {
	GetAllProduct(c *gin.Context)
	UpdateSku(c *gin.Context)
	GetByFeild(c *gin.Context)
	First(dest interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}
