package entity

import (
	"errors"

	"github.com/racoondad/nanpo"
	"github.com/racoondad/nanpo/app/user/domain/event"
	"github.com/racoondad/nanpo/app/user/domain/po"
)

type Admin struct {
	nanpo.Entity
	po.Admin
}

// Identity 唯一
func (ety *Admin) Identity() string {
	return ety.UUID
}

func (u *Admin) ChangePassword(newPassword, oldPassword string) error {
	if u.Password != oldPassword {
		return errors.New("password error")
	}
	u.SetPassword(newPassword)

	// 用户实体加入修改密码事件
	u.AddPubEvent(&event.ChangePassword{})
	return nil
}
