package User

import (
	"sample/mocktest/Test_User"
)

type User struct {
	Test_User Test_User.TestUserInterface
}

func (u *User) Use() {
	u.Test_User.AddUser(1, "sample test")
}
