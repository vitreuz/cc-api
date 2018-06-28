package main_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"

	"github.com/vitreuz/hackday/api"
	. "github.com/vitreuz/hackday/examples/get_apps"
)

var _ = Describe("Run", func() {
	var (
		client API

		apps []string
		err  error

		server *Server
	)

	JustBeforeEach(func() {
		server = NewServer()
		client = api.NewClient(server.URL(), "")
		response := `{
			"resources": [
				{
					"name": "app-1"
				},
				{
					"name": "app-2"
				}
			]
		}`

		server.AppendHandlers(
			CombineHandlers(
				VerifyRequest(http.MethodGet, "/v3/apps"),
				RespondWith(http.StatusOK, response),
			),
		)

		apps, err = Run(client)
	})

	It("prints all the apps", func() {
		Expect(server.ReceivedRequests()).To(HaveLen(1))

		Expect(err).ToNot(HaveOccurred())
		Expect(apps).To(ConsistOf("app-1", "app-2"))

	})
})
