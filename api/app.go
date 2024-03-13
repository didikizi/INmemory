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

	server.GET("/users", delivery.List)               //Ошибка на уровне базы с логикой вывода
	server.POST("/users", delivery.Create)            //Поставить валидатор значений!
	server.DELETE("/users/:Account", delivery.Delete) //Удаление из MAP
	server.PUT("/users/:Account", delivery.Update)    //Обновить систему поиска

	log.Fatal(server.Start(":8081"))
}
