package po

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"
	"gorm.io/gorm"
)

type Position struct {
	changes map[string]interface{}
	ormtypes.Model
	Description string `gorm:"column:description;autoUpdateTime:milli"`
}

// TableName .
func (obj *Position) TableName() string {
	return "base_position"
}

func (obj *Position) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(obj)
}

// GetChanges .
func (obj *Position) GetChanges() map[string]interface{} {
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
func (obj *Position) Update(key string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[key] = value
}
