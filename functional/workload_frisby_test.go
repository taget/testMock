package functional_test

import (
	. "github.com/onsi/ginkgo"

	"github.com/verdverm/frisby"
	//"github.com/bitly/go-simplejson"
)

var _ = Describe("Workload frisby", func() {

	var (
		url string
		F   *frisby.Frisby
	)

	BeforeEach(func() {
		By("set url and get request")
		// read from config file
		url = "http://127.0.0.1:8888/v1/workloads"
		F = frisby.Create("Test Workload")

	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
	})

	Describe("Get Workload", func() {
		Context("Get all workloads", func() {
			It("should have no error", func() {
				By("Get request  ... ")
				F.Get(url).
					SetHeader("Content-Type", "application/json").
					Send().
					ExpectStatus(200).
					ExpectJsonLength("json", 0)
			})
		})
	})
})
