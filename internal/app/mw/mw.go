package mw

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
	"time"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {

	currentTime := time.Now()

	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(val, roleAdmin) {
			log.Println("Admin has logged at " + currentTime.Format("2006-01-02 3:4:5 PM"))
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
