package user

import (
	"fmt"

	"github.com/taget/testMock/doer"
	"github.com/taget/testMock/lib"
)

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	r, _ := lib.What()
	if r == 0 {
		return u.Doer.DoSomething(123, "Hello GoMock")
	} else {
		return fmt.Errorf("I am fail")
	}
}
