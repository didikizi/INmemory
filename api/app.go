package app

import (
	"inmemory/local/cmd/delivery"
	memory "inmemory/local/cmd/memory"
	"log"

	"github.com/labstack/echo/v4"
)

func App() {
	base := memory.NewBase()
	delivery := delivery.New(*base)
	server := echo.New()

	server.GET("/users", delivery.List)               //+
	server.POST("/users", delivery.Create)            //+
	server.DELETE("/users/:Account", delivery.Delete) //+
	server.PUT("/users/:Account", delivery.Update)    //+

	log.Fatal(server.Start(":8081"))
}
