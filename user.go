package root

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}

// taken from https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e
