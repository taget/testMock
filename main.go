package main

import (
	"github.com/taget/testMock/doer"
	"github.com/taget/testMock/user"
)

func main() {
	u := &user.User{&doer.Foo{}}
	u.Use()
}
