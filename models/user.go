package models

import (
	"errors"

	"github.com/kikiyou/agent/forms"
)

//user ...
type user struct {
	ID       string `db:"id, primarykey, autoincrement" json:"id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"-"`
	Role     string `db:"role" json:"role"`
}

//userModel ...
type UserModel struct{}

// 写死user
var userList = []user{
	user{ID: "0", Name: "admin", Password: "admin", Role: "0"},
	user{ID: "1", Name: "tom", Password: "cat", Role: "1"},
}

// Check if the username and password combination is valid
func isUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Name == username && u.Password == password {
			return true
		}
	}
	return false
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Name == username {
			return false
		}
	}
	return true
}

func getUser(username, password string) user {
	for _, u := range userList {
		if u.Name == username && u.Password == password {
			return u
		}
	}
	return user{}
}

//Signin ...
func (m UserModel) Login(form forms.LoginForm) (user user, err error) {

	user = getUser(form.Name, form.Password)
	if user.Password == "" {
		return user, errors.New("账户或密码错误")
	}
	return user, nil
}
