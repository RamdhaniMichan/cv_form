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

	profileRoutes := routes.ProfileRoute{}
	profileRoutes.Routing(httpRouter)

	skillRoutes := routes.SkillRoute{}
	skillRoutes.Routing(httpRouter)

	weRoutes := routes.WERoute{}
	weRoutes.Routing(httpRouter)

	educationRoutes := routes.EducationRoute{}
	educationRoutes.Routing(httpRouter)

	employmentRoutes := routes.EmploymentRoute{}
	employmentRoutes.Routing(httpRouter)

	uploadRoutes := routes.UploadRoute{}
	uploadRoutes.Routing(httpRouter)

	httpRouter.SERVE()

}
