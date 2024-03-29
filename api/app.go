package app

import (
	"context"
	"errors"
	"inmemory/local/cmd/delivery"
	jwt "inmemory/local/cmd/delivery/jwt"
	memory "inmemory/local/cmd/memory"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/labstack/gommon/log"
)

var key = "secretkey"

func App() {
	base := memory.NewBase()
	delivery := delivery.New(*base)
	jwtBase := jwt.New(*base, key)
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Logger())
	server.Logger.SetLevel(log.DEBUG)

	server.POST("/login", jwtBase.Login)
	server.GET("/users", delivery.List)

	v1Group := server.Group("/v1")
	v1Group.Use(jwt.JWTAutoMiddleware(key))
	v1Group.POST("/users", delivery.Create)
	v1Group.DELETE("/users/:Account", delivery.Delete)
	v1Group.PUT("/users/:Account", delivery.Update)

	go func() {
		if err := server.Start(":8081"); err != nil && errors.Is(err, http.ErrServerClosed) {
			server.Logger.Fatal(err)
		}
	}()

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	<-quite
	server.Logger.Info("shutdown inited")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
