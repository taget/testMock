package functional_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"

	"github.com/parnurzeal/gorequest"
)

var _ = Describe("Hospitality", func() {

	var (
		url    string
		resp   gorequest.Response
		body   string
		errors []error
	)

	BeforeEach(func() {
		By("set url")
		// read from config file
		url = "http://127.0.0.1:8888"

	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
	})

	Describe("Get Hospitality score", func() {
		Context("Call DoSomething", func() {
			BeforeEach(func() {
				By("set hospitality score url")
				// read from config file
				url = "http://127.0.0.1:8888/v1/hospitality"
				resp, body, errors = gorequest.New().Post(url).
					Send(`{"max_cache":2, "min_cache": 2}`).
					Set("Content-Type", "application/json").End()
			})
			It("should have no error", func() {
				By("Get request  ... ")
				fmt.Println(resp)
				fmt.Println(body)
				fmt.Println(errors)
			})
		})
	})
})
