package user

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/taget/testMock/mocks"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	Convey("Test Use returns errors", t, func() {
		// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
		mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)
		err := testUser.Use()
		So(err, ShouldEqual, dummyError)
	})
}
