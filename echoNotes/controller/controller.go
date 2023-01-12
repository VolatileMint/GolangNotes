package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func Get(c echo.Context) error {
	// GET通信で受け取る
	id := c.QueryParam("id")
	fmt.Println("id=", id)
	return c.String(http.StatusOK, "Getで受けとり"+id) // 文字列で返す
}
func Save(c echo.Context) error {
	// POST通信で受け取る
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"` // ここで設定しておくとformなどから受け取れたり、JSONに変換ができる
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// GET,POSTどちらでもOK
func Json(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u) // JSON形式で返す
}
