package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-toko-app/internal/usecase/product_usecase"
)

func (x *Handler) GetProductByID() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(500, responseBody{StatusCode: 500, Message: http.StatusText(500), Error: err.Error()})
		}

		product, err := x.ProductUC.GetProductByID(c.Request().Context(), product_usecase.GetProductByIDRequest{ProductID: productID})
		if err != nil {
			x.log.Err(err).Msg("[GetProductByID][productUC.GetProductByID]")

			return c.JSON(500, responseBody{StatusCode: 500, Message: http.StatusText(500), Error: err.Error()})
		}
		return c.JSON(http.StatusOK, responseBody{http.StatusOK, http.StatusText(http.StatusOK), "", product})
	}
}
