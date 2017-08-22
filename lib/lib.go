package lib

import "fmt"

var What = func() (int, error) {
	// only define a variable in testcase we can stub it
	fmt.Println("I am a lib func, you need to stub me for return value in your test case")

	return 1, nil
}
