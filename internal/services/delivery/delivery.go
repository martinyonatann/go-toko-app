package delivery

import "github.com/labstack/echo/v4"

type Handlers interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc

	GetProductByID() echo.HandlerFunc
}
