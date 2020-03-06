package apis

import (
	"github.com/labstack/echo"
	"github.com/stark-industries/pkg/avenger"
	"net/http"
)

func SaveAvenger(context echo.Context) error {
	formData := new(avenger.Avenger)
	if err := context.Bind(formData); err != nil {
		return err
	}
	data := formData

	savingError:=data.Save()
	if savingError!=nil{
		return context.JSON(http.StatusBadRequest,"Operation Failed!")
	}

return context.JSON(http.StatusAccepted,"Operation Successful")

}


func FindById(context echo.Context) error {
	id:= context.Param("id")
	data := new(avenger.Avenger)
	data.Id=id
	staff:=data.FindById()
	return context.JSON(http.StatusAccepted,staff)

}

func FindAll(context echo.Context) error {
	data := new(avenger.Avenger)
	staff:=data.FindAll()
	return context.JSON(http.StatusAccepted,staff)

}