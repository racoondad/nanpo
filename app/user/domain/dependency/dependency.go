package dependency

import (
	"github.com/racoondad/nanpo/app/user/domain/entity"
	"github.com/racoondad/nanpo/app/user/domain/vo"
)

type AdminRepo interface {
	Get(Code string) (obj *entity.Admin, err error)
	Save(entity.Admin) (err error)
	New(userDto vo.CreateAdminReq) (obj *entity.Admin, err error)
}
type RoleRepo interface {
	Get(Code string) (obj *entity.Role, err error)
	Save(entity.Role) (err error)
	New(userDto vo.CreateRoleReq) (obj *entity.Role, err error)
}
