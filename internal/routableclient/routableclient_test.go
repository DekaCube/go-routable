package routableclient

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoutableClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routable Client Suite")
}

var _ = Describe("NewSimpleRoutableClient", func() {
	Context("when host is invalid", func() {
		It("should return an error when scheme is missing", func() {
			_, err := NewSimpleRoutableClient("token", "api.routable.com", 10*time.Second)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("scheme and host are required"))
		})

		It("should return an error when host is missing", func() {
			_, err := NewSimpleRoutableClient("token", "https://", 10*time.Second)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("scheme and host are required"))
		})

		It("should succeed when scheme and host are provided", func() {
			_, err := NewSimpleRoutableClient("token", "https://api.routable.com", 10*time.Second)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
