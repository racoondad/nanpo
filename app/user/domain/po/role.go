package po

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"
	"gorm.io/gorm"
)

type Role struct {
	changes map[string]interface{}
	ormtypes.Model
	Name        string `gorm:"column:name;unique"`
	Description string `gorm:"column:description"`
}

// TableName .
func (obj *Role) TableName() string {
	return "base_role"
}

func (obj *Role) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(obj)
}

// GetChanges .
func (obj *Role) GetChanges() map[string]interface{} {
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
func (obj *Role) Update(key string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[key] = value
}

func (obj *Role) SetCode(code string) {
	obj.Code = code
	obj.Update("code", code)
}

func (obj *Role) SetName(name string) {
	obj.Name = name
	obj.Update("Name", name)
}
