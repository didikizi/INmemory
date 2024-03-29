package delivery

import (
	memory "inmemory/local/cmd/memory"
	"inmemory/local/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoDeliveries interface {
	Create(ectx echo.Context) error
	Update(ectx echo.Context) error
	List(ectx echo.Context) error
	Delete(ectx echo.Context) error
	Logon(b memory.Base) echo.MiddlewareFunc
}

type EchoDelivery struct {
	base *memory.Base
}

func New(b memory.Base) EchoDelivery {
	return EchoDelivery{
		base: &b,
	}
}

func (ech *EchoDelivery) Create(ectx echo.Context) error {
	newUser := &models.Filter{}
	err := ectx.Bind(newUser)
	if err != nil {
		return err
	}
	if newUser.Account == 0 || newUser.Name == "" || newUser.Value == 0 {
		return ectx.JSON(http.StatusBadRequest, "Были переданы не все параметры")
	}
	return ectx.JSON(http.StatusOK, ech.base.Create(newUser))
}

func (ech *EchoDelivery) List(ectx echo.Context) error {
	cnt := 0
	tmp := ectx.QueryParam("account")
	if tmp == "" {
		tmp = "0"
		cnt++
	}
	account, err := strconv.Atoi(tmp)
	if err != nil {
		return ectx.JSON(http.StatusBadRequest, "Был передано некоректное значение Account")
	}

	tmp2 := ectx.QueryParam("value")
	if tmp2 == "" {
		tmp2 = "0"
		cnt++
	}
	value, err := strconv.ParseFloat(tmp2, 64)
	if err != nil {
		return ectx.JSON(http.StatusBadRequest, "Был передано некоректное значение Value")
	}

	name := ectx.QueryParam("name")
	if name == "" {
		cnt++
	}

	filter := &models.Filter{
		Name:    name,
		Account: account,
		Value:   value,
	}

	if cnt < 2 {
		return ectx.JSON(http.StatusBadRequest, "Было передано более одного параметра фильтрации")
	}

	return ectx.JSON(http.StatusOK, ech.base.List(filter))
}

func (ech *EchoDelivery) Delete(ectx echo.Context) error {
	account, err := strconv.Atoi(ectx.Param("Account"))
	if err != nil {
		return ectx.JSON(http.StatusBadRequest, nil)
	}
	return ectx.JSON(http.StatusOK, ech.base.Delete(account))
}

func (ech *EchoDelivery) Update(ectx echo.Context) error {
	newUser := &models.Filter{}
	err := ectx.Bind(newUser)
	if err != nil {
		return err
	}
	if newUser.Account == 0 || newUser.Name == "" || newUser.Value == 0 {
		return ectx.JSON(http.StatusBadRequest, "Были переданы не все параметры")
	}
	account, err := strconv.Atoi(ectx.Param("Account"))
	if err != nil {
		return ectx.JSON(http.StatusBadRequest, "Имя аккаунта не коректно")
	}
	return ectx.JSON(http.StatusOK, ech.base.Update(newUser, account))
}
