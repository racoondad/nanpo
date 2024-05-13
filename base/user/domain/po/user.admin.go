package po

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"
	"gorm.io/gorm"
)

type Admin struct {
	changes map[string]interface{}
	ormtypes.Model
}

// TableName .
func (obj *Admin) TableName() string {
	return "base_Admin"
}

func (obj *Admin) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(obj)
}

// GetChanges .
func (obj *Admin) GetChanges() map[string]interface{} {
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
func (obj *Admin) Update(key string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[key] = value
}
