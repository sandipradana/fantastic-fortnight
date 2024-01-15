package vercel

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

var router *fiber.App
var routerAdaptor http.HandlerFunc

func init() {
	router = fiber.New()

	routerGroup := router.Group("/api")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World !")
	})

	routerAdaptor = adaptor.FiberApp(router)
}

func Handler(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	routerAdaptor.ServeHTTP(httpResponse, httpRequest)
}
