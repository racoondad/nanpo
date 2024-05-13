package po

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"
	"gorm.io/gorm"
)

type Employee struct {
	changes map[string]interface{}
	ormtypes.Model
	OrgCode string `json:"orgCode" form:"orgCode"`
	OrgName string `json:"orgName" form:"orgName"`
}

// TableName .
func (obj *Employee) TableName() string {
	return "base_employee"
}

func (obj *Employee) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(obj)
}

// GetChanges .
func (obj *Employee) GetChanges() map[string]interface{} {
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
func (obj *Employee) Update(key string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[key] = value
}
