package vo

// 修改用户密码
type ChangeAdminPasswordReq struct {
	Code        string
	NewPassword string `validate:"required"`
	OldPassword string `validate:"required"`
}

// 新增用户
type CreateAdminReq struct {
	Code     string
	Name     string
	RoleCode string
	Password string
}

// 更新用户
type UpdateAdminReq struct {
	Code     string
	Name     string
	RoleCode string
	Password string
}

// 新增角色
type CreateRoleReq struct {
	Code string
	Name string
}

// 更新角色
type UpdateRoleReq struct {
	Code string
	Name string
}
