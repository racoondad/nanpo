package vo

// 用户信息
type AdminInfoRes struct {
	Id   uint64
	Code string
	Name string
	Role struct {
		Code string
		Name string
	}
}

// 角色信息
type RoleInfoRes struct {
	Id   uint64
	Code string
	Name string
}
