package routes

import (
	"net/http"

	router "template/http"
	"template/migration"
)

var (
	httpRouterMigrate router.Router = router.NewMuxRouter()
)

type MigrateRoutes struct{}

func (c *MigrateRoutes) Routing(httpRouter router.Router) {
	httpRouter.GET("/migrate", func(w http.ResponseWriter, r *http.Request) {
		migration.Migrate(w, r)
		return
	})
}
