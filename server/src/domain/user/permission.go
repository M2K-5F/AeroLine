package user

type Permission string

const (
	CustomerPermission = Permission("CUSTOMER")
	AdminPermission    = Permission("ADMIN")
)

func (ths Permission) String() string {
	return string(ths)
}
