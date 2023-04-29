package chapter2

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler interface {
	IsUserPaymentActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	e := echo.New()

	// /user/:userId/payment/active
	// can be used by other team to check a user
	// has an active subscription or not
	e.GET("/user/:userId/payment/active", func(c echo.Context) error {
		// check auth, etc

		userId := c.Param("userId")
		isActive := u.IsUserPaymentActive(c.Request().Context(), userId)

		return c.JSON(http.StatusOK, UserActiveResponse{IsActive: isActive})
	})
}
