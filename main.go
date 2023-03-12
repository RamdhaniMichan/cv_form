package main

import (
	"fmt"
	"log"
	"net/http"
	router "template/http"
	"template/routes"

	"github.com/joho/godotenv"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	httpRouter.BASE("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Healthy")
	})

	migrateRoutes := routes.MigrateRoutes{}
	migrateRoutes.Routing(httpRouter)

	bookRoutes := routes.BookRoute{}
	bookRoutes.Routing(httpRouter)

	httpRouter.SERVE()

}
