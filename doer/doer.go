package doer

import (
	"fmt"
)

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks github.com/taget/testMock/doer Doer

type Doer interface {
	DoSomething(int, string) error
}

type Foo struct {
}

func (f *Foo) DoSomething(i int, name string) error {

	fmt.Printf("Foo %d, %s\n", i, name)
	return nil

}
