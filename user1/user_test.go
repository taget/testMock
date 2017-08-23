package user1

import (
	"errors"
	"testing"

	"github.com/taget/testMock/lib"
	"github.com/taget/testMock/mocks"

	gomock "github.com/golang/mock/gomock"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	Convey("Test Use return nil", t, func() {
		subs := StubFunc(&lib.What, 0, nil)
		defer subs.Reset()
		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)
		err := testUser.Use()
		So(err, ShouldEqual, nil)
	})

	Convey("Test Use returns errors", t, func() {

		subs := StubFunc(&lib.What, 0, nil)
		defer subs.Reset()

		dummyError := errors.New("dummy error")
		// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)
		err := testUser.Use()
		So(err, ShouldEqual, dummyError)
	})

	Convey("Test Use returns errors as lib returns non-zero", t, func() {

		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Times(0)
		err := testUser.Use()
		So(err, ShouldNotBeNil)
	})

}
