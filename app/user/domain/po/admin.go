package po

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/ormtypes"
	"gorm.io/gorm"
)

type Admin struct {
	changes map[string]interface{}
	ormtypes.Model
	Name     string
	Password string
	RoleCode string
	Role     Role
}

// TableName .
func (obj *Admin) TableName() string {
	return "base_admin"
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

func (obj *Admin) SetName(name string) {
	obj.Name = name
	obj.Update("name", name)
}

func (obj *Admin) SetRole(roleCode string) {
	obj.RoleCode = roleCode
	obj.Update("role_code", roleCode)
}

func (obj *Admin) SetPassword(newPassword string) {
	obj.Password = newPassword
	obj.Update("password", newPassword)
}
