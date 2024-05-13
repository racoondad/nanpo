package ormtypes

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement"`               // id
	UUID      string    `gorm:"column:uuid;unique;<-:create;not null"`            // 雪花id
	Code      string    `gorm:"column:code;unique;<-:create;not null"`            // 编码
	CreatedAt time.Time `gorm:"column:created_at;<-:create;autoUpdateTime:milli"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli"`           // 更新时间
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if m.Code == "" {
		m.Code = uuid.NewString()
	}
	if m.UUID == "" {
		m.UUID = uuid.NewString()
	}
	return nil
}

/*
entity *Entity common.IModelInterface
*/
type IModelInterface interface {
	// # return: name string 表名称
	// # description: [实例]表名称
	TableName() string
	// # param: db *gorm.DB 引擎
	// # return: err error 异常
	// # description: [实例]表创建
	Migrate(db *gorm.DB) error
	// # param: db *gorm.DB 引擎
	// # return: error 异常
	// # description: [实例]新增
	Create(db *gorm.DB) error
	// # param: db *gorm.DB 引擎
	// # return: error 异常
	// # description: [实例]更新
	Update(db *gorm.DB) error
	// # param: db *gorm.DB 引擎
	// # return: error 异常
	// # description: 通过编号查询[实例]
	Query(db *gorm.DB, code string) error
	// # param: db *gorm.DB 引擎
	// # return: err error 异常
	// # description: 删除[实例]
	Delete(db *gorm.DB, codeList []string) (err error)
	// # param: db *gorm.DB 引擎; offset int 页码; limit int 尺寸; whereScope ...*gorm.DB 条件
	// # return: resultList interface{} 结果; err error 异常
	// # description: 分页获取[实例]列表
	List(db *gorm.DB, offset, limit int, whereScope ...*gorm.DB) (resultList interface{}, err error)
	// # param: db *gorm.DB 引擎; whereScope ...*gorm.DB 条件
	// # return: total int64 数量; err error 异常
	// # description: 分页获取[实例]列表
	Count(db *gorm.DB, whereScope ...*gorm.DB) (total int64)
}
