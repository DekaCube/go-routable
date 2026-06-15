package routableclient

import (
	_ "embed"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/DekaCube/go-routable/internal/types"
)

//go:embed testdata/fixture.json
var companiesFixture string

var _ = Describe("SimpleRoutableClient", func() {
	var (
		client *SimpleRoutableClient
		server *ghttp.Server
		resp   *types.RoutableListCompaniesResponse
		err    error
	)

	BeforeEach(func() {
		server = ghttp.NewServer()

		client, err = NewSimpleRoutableClient(
			"test-token",
			server.URL(),
			30*time.Second,
		)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Context("ListCompanies with no arguments", func() {
		BeforeEach(func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/v1/companies"),
				ghttp.VerifyHeaderKV("Authorization", "Bearer test-token"),
				ghttp.RespondWith(200, companiesFixture),
			))

			resp, err = client.ListCompanies()
		})

		It("should not return an error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a response with one company", func() {
			Expect(resp.Results).To(HaveLen(1))
		})

		It("should return the correct company ID", func() {
			Expect(resp.Results[0].ID).To(Equal("ec60d364-9100-4728-a3cb-0f200a8a0f5d"))
		})

		It("should return the correct business name", func() {
			Expect(resp.Results[0].BusinessName).To(Equal("Vance Refrigeration"))
		})
	})
})
