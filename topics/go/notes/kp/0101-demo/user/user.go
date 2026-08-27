package user

type User struct {
	name string // 小写：仅本包可见
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Name() string {
	return u.name
}
