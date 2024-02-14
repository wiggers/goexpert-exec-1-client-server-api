package main

import (
	"net/http"

	"github.com/wiggers/goexpert/exec/1-client-server-api/configs"
	"github.com/wiggers/goexpert/exec/1-client-server-api/internal/entity"
	"github.com/wiggers/goexpert/exec/1-client-server-api/internal/infra/controller"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Exchange{})
	controller := controller.NewExchangeController(db)

	http.HandleFunc("/cotacao", controller.CheckExchangeRate)
	http.ListenAndServe(":"+config.WebServerPort, nil)
}
