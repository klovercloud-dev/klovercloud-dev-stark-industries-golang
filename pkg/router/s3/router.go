package s3

import (
	"github.com/labstack/echo"
	"github.com/stark-industries/pkg/apis"
)

func Router(g *echo.Group) {
	g.GET("/:keyname", apis.GetFile)
	g.POST("",apis.Upload)
}


