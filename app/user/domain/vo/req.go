package vo

type ChangeAdminPasswordReq struct {
	Code        string
	NewPassword string `validate:"required"`
	OldPassword string `validate:"required"`
}

type CreateAdminReq struct {
	Code     string
	Name     string
	RoleCode string
	Password string
}
type UpdateAdminReq struct {
	Code     string
	Name     string
	RoleCode string
	Password string
}
type CreateRoleReq struct {
	Code string
	Name string
}
type UpdateRoleReq struct {
	Code string
	Name string
}
