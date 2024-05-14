package entity

import (
	"github.com/racoondad/nanpo/app/user/domain/po"
)

type Role struct {
	po.Role
}

// Identity 唯一
func (ety *Role) Identity() string {
	return ety.UUID
}
