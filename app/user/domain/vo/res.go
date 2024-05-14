package vo

type AdminInfoRes struct {
	Id   uint64
	Code string
	Name string
	Role struct {
		Code string
		Name string
	}
}

type RoleInfoRes struct {
	Id   uint64
	Code string
	Name string
}
