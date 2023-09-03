/*
package routes

import (

	"GO/controllers"

	"github.com/labstack/echo/v4"

)

	func New() *echo.Echo {
		e := echo.New()

		e.GET("/users", controllers.GetUserControllers)

		return e
	}
*/
package routes

import (
	"os"
	"prakerja/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("PRIVATE_KEY_JWT"))))

	eLog := eAuth.Group("")
	eLog.Use(middleware.Logger())
	eLog.GET("/book", controllers.GetBookController)

	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.POST("/book", controllers.AddBookController)
	eAuth.GET("/book/:id", controllers.GetDetailBookController)
}
