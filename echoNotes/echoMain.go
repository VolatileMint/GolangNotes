package echoNotes

import (
	"GolangNotes/echoNotes/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoMain() {
	e := echo.New()
	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルート設定
	e.GET("/", controller.Hello) // Helloコントローラーに飛ばす
	// curl.exe http://localhost:8080/get?id=1
	e.GET("/get", controller.Get)
	// curl.exe -d "name=Joe Smith" -d "email=joe@labstack.com" http://localhost:8080/save
	e.POST("/save", controller.Save) // post name="namae",email="meado"
	// curl.exe 'http://localhost:8080/jsonGet?name=aa&email=aa@aa'
	e.GET("/jsonGet", controller.Json)
	// curl.exe 'http://localhost:8080/json' -F "name=a" -F "email=a@a"
	e.POST("/jsonPost", controller.Json)

	//curl.exe 'http://localhost:8080/admin/'
	g := e.Group("/admin") // ルートをグループ化 /admin/~になる
	g.GET("/", controller.Admin)
	// curl.exe 'http://localhost:8080/admin/users/20'
	g.GET("/users/:id", controller.Find) // :idの文字列を idパラムで受け取れる
	e.Logger.Fatal(e.Start(":8080"))
}
