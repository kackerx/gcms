package domain

type UserRepo interface {
	FindByID(id int64) (*User, error)
	FindByName(name string) (*User, bool, error)
	SaveUser(user *User) (int, error)
}
