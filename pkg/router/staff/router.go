package staff

import (
	"github.com/labstack/echo"
	"github.com/stark-industries/pkg/apis"
)

func Router(g *echo.Group) {
	g.GET("",apis.FindAll)
	g.GET("/:id",apis.FindById)
	g.POST("",apis.SaveStaff)
}



