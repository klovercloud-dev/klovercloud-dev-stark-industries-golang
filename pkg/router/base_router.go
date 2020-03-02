package router

import (
	"github.com/labstack/echo"
	"github.com/stark-industries/pkg/router/s3"
	"github.com/stark-industries/pkg/router/staff"
	"net/http"
)

func Routes(e *echo.Echo) {
	// Index Page
	e.GET("/", index)

	// Health Page
	e.GET("/health", health)

	s3Monitor := e.Group("/api/v1/s3s")
	s3.Router(s3Monitor)

	staffMonitor := e.Group("/api/v1/staffs")
	staff.Router(staffMonitor)

}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "This is KloverCloud Notification Service")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "I am live!")
}
