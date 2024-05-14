/*
 * @Author       : lptecodad lptecodad@sina.com
 * @Date         : 2023-01-07 08:48:06
 * @LastEditors  : lptecodad lptecodad@sina.com
 * Copyright (c) 2023 by lptecodad lptecodad@sina.com, All Rights Reserved.
 */
package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MODEL struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id" form:"id" default:"1" example:"1"`                                                                 // id
	UUID      string    `gorm:"column:uuid;unique;<-:create;not null" json:"uuid" form:"uuid" default:"字符编码" example:"字符编码"`                                                    // 雪花id
	Code      string    `gorm:"column:code;unique;<-:create;not null" json:"code" form:"code" default:"编号" example:"编号"`                                                        // 编码
	CreatedAt time.Time `gorm:"column:created_at;<-:create;autoUpdateTime:milli" json:"createdAt" form:"createdAt" default:"2006-01-02 15:03:04" example:"2006-01-02 15:03:04"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:milli" json:"updatedAt" form:"updatedAt" default:"2006-01-02 15:03:04" example:"2006-01-02 15:03:04"`           // 更新时间
}

// 新增
func (u *MODEL) BeforeSave(tx *gorm.DB) (err error) {
	return
}

func (u *MODEL) BeforeCreate(tx *gorm.DB) (err error) {
	// tx 是带有 `NewDB` 选项的新会话模式
	// 基于 tx 的操作会在同一个事务中，但不会带上任何当前的条件
	if u.UUID == "" {
		u.UUID = uuid.NewString()
	}
	if u.Code == "" {
		u.Code = uuid.NewString()
	}
	return
}

func (u *MODEL) AfterSave(tx *gorm.DB) (err error) {
	return
}

func (u *MODEL) AfterCreate(tx *gorm.DB) (err error) {
	return
}

// 修改
func (u *MODEL) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

func (u *MODEL) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// 删除
func (u *MODEL) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

func (u *MODEL) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// 查询
func (u *MODEL) AfterFind(tx *gorm.DB) (err error) {
	return
}
