package jwt

import (
	"inmemory/local/cmd/memory"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt"
)

type Base struct {
	users memory.Base
	key   []byte
}

type CustomClaims struct {
	jwt.RegisteredClaims
}

type Credentials struct {
	Account int    `json:"account"`
	Pass    string `json:"password"`
}

func New(base memory.Base, key string) Base {
	return Base{users: base, key: []byte(key)}
}

func (b *Base) Login(ectx echo.Context) error {
	creds := &Credentials{}
	err := ectx.Bind(creds)
	log.Print(ectx.Request().Body)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if !b.users.Validate(creds.Account, creds.Pass) {
		return echo.ErrUnauthorized
	}

	customClaims := CustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Login func",
			Subject:   string(creds.Account),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	stoken, err := token.SignedString(b.key)
	if err != nil {
		return err
	}
	return ectx.JSON(http.StatusOK, echo.Map{"token": stoken})
}

func JWTAutoMiddleware(key string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(key),
	})
}
