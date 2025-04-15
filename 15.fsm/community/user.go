package community

type User struct {
	id      string
	isAdmin bool
}

func NewUser(id string, isAdmin bool) *User {
	return &User{
		id:      id,
		isAdmin: isAdmin,
	}
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) IsAdmin() bool {
	return u.isAdmin
}
