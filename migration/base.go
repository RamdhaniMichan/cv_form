package migration

import (
	"net/http"
	"template/entity"

	"template/datasource"
	"template/function"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migrate(w http.ResponseWriter, r *http.Request) {
	autoMigrate()

	function.SendResponse(w, http.StatusOK, "Success", nil)
}

// autoMigrate ...
func autoMigrate() {
	db := datasource.OpenDB()
	db.AutoMigrate(
		&entity.Profile{},
		&entity.Skill{},
		&entity.WorkingExperience{},
		&entity.Education{},
		&entity.Employment{},
	)
}
