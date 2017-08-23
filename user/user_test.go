package user_test

import (
	"errors"

	"github.com/taget/testMock/lib"
	"github.com/taget/testMock/mocks"
	. "github.com/taget/testMock/user"

	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/prashantv/gostub"
)

var _ = Describe("User", func() {
	var (
		mockCtrl *gomock.Controller
		mockDoer *mocks.MockDoer
		testUser *User
	)

	BeforeEach(func() {
		By("Create mock Doer")
		mockCtrl = gomock.NewController(GinkgoT())
		mockDoer = mocks.NewMockDoer(mockCtrl)
		testUser = &User{Doer: mockDoer}

	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Test User Dosometing action", func() {
		Context("Call DoSomething", func() {
			// Focus on this spec
			FIt("should have no error", func() {
				By("Stub func ... ")
				subs := StubFunc(&lib.What, 0, nil)
				defer subs.Reset()
				mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)
				testUser.Use()

			})
			// Mark this spec as pending
			PIt("should returns errors", func() {

				subs := StubFunc(&lib.What, 0, nil)
				defer subs.Reset()

				dummyError := errors.New("dummy error")
				// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
				mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)
				err := testUser.Use()
				Expect(err).To(Equal(dummyError))
			})

		})
	})

})
