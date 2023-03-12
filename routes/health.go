package routes

import (
	"fmt"
	"net/http"
	router "template/http"
)

var (
	httpRouterHealth router.Router = router.NewMuxRouter()
)

type HealthRoutes struct{}

func (c *HealthRoutes) Routing() {
	httpRouterHealth.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up n run")
	})
}
