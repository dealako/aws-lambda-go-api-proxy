package httpadapter_test

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dealako/aws-lambda-go-api-proxy/httpadapter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("unfortunately-required-header", "")
	fmt.Fprintf(w, "Go Lambda!!")
}

var _ = Describe("HTTPAdapter tests", func() {
	Context("Simple ping request", func() {
		It("Proxies the event correctly", func() {
			log.Println("Starting test")

			var httpHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Add("unfortunately-required-header", "")
				fmt.Fprintf(w, "Go Lambda!!")
			})

			adapter := httpadapter.New(httpHandler)

			req := events.APIGatewayProxyRequest{
				Path:       "/ping",
				HTTPMethod: "GET",
			}

			resp, err := adapter.ProxyWithContext(context.Background(), req)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))

			resp, err = adapter.Proxy(req)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
