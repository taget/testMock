package functional_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"

	"gopkg.in/gavv/httpexpect.v1"
)

var _ = Describe("Hospitality", func() {

	var (
		url string
		he  *httpexpect.Expect
	)

	BeforeEach(func() {
		By("set url")
		// read from config file
		url = "http://127.0.0.1:8888"

	})

	JustBeforeEach(func() {
		he = httpexpect.New(GinkgoT(), url)
	})

	AfterEach(func() {
	})

	Describe("Get Hospitality score", func() {
		Context("Call Hospitality", func() {
			data := map[string]interface{}{
				"max_cache": 2,
				"min_cache": 2,
			}
			BeforeEach(func() {
				By("set hospitality score url")
			})

			FIt("should have no error", func() {
				By("Get request  ... ")
				schema := `{
					"type": "object",
					"properties": {
						"score": {
							"type": "object",
							"properties": {
								"l3": {
									"type": "object",
									"properties": {
										"0": {
											"type": "integer",
											"minimum": 0
										},
										"1": {
											"type": "integer",
											"minimum": 0
										}
									}
								}
							}
						}
					},
					"required": ["score"]
				}`

				repos := he.POST("/v1/hospitality").
					WithHeader("Content-Type", "application/json").
					WithJSON(data).
					Expect().
					Status(http.StatusOK).
					JSON()

				//obj.Keys().ContainsOnly("score")
				repos.Schema(schema)

			})
		})
	})
})
