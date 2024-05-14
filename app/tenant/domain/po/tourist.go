package po

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"
	"gorm.io/gorm"
)

type Tourist struct {
	changes map[string]interface{}
	ormtypes.Model
	OrgCode string `json:"orgCode" form:"orgCode"`
	OrgName string `json:"orgName" form:"orgName"`
}

// TableName .
func (obj *Tourist) TableName() string {
	return "base_tourist"
}

func (obj *Tourist) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(obj)
}

// GetChanges .
func (obj *Tourist) GetChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// Update .
func (obj *Tourist) Update(key string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[key] = value
}
