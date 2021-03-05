package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/yasuoza/clean-architecture-go/server/config"
	"github.com/yasuoza/clean-architecture-go/server/infrastructure/datastore"
	"github.com/yasuoza/clean-architecture-go/server/infrastructure/router"
	"github.com/yasuoza/clean-architecture-go/server/registry"
	"gorm.io/gorm"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.Debug()
	defer closeDB(db)

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
}

func closeDB(db *gorm.DB) error {
	d, _ := db.DB()
	return d.Close()
}
