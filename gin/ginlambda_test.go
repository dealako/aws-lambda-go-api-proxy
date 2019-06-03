package ginadapter_test

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/dealako/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GinLambda tests", func() {
	Context("Simple ping request", func() {
		It("Proxies the event correctly", func() {
			log.Info("Starting test")
			r := gin.Default()
			r.GET("/ping", func(c *gin.Context) {
				log.Info("Handler!!")
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})

			adapter := ginadapter.New(r)

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
