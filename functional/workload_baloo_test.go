package functional_test

import (
	"testing"

	. "github.com/onsi/ginkgo"

	"gopkg.in/h2non/baloo.v1"
)

var _ = Describe("Workload", func() {

	var (
		url  string
		test *baloo.Client
		t    *testing.T
	)

	BeforeEach(func() {
		By("set url and get request")
		// read from config file
		url = "http://127.0.0.1:8888/v1/workloads"
		test = baloo.New(url)

	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
	})

	Describe("Get Workload", func() {
		Context("Get all workloads", func() {
			It("should have no error", func() {
				By("Get request  ... ")
				test.Get("/").
					SetHeader("Content-Type", "application/json").
					Expect(GinkgoT()).
					Status(200).
					Done()
			})
		})
	})
})
