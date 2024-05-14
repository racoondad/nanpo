package po

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"
	"gorm.io/gorm"
)

type Customer struct {
	changes map[string]interface{}
	ormtypes.Model
	OrgCode string `json:"orgCode" form:"orgCode"`
	OrgName string `json:"orgName" form:"orgName"`
}

// TableName .
func (obj *Customer) TableName() string {
	return "base_customer"
}

func (obj *Customer) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(obj)
}

// GetChanges .
func (obj *Customer) GetChanges() map[string]interface{} {
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
func (obj *Customer) Update(key string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[key] = value
}
