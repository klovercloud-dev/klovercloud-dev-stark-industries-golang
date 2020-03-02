package apis

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo"
	"github.com/stark-industries/config"
	"io/ioutil"
	"net/http"
)

func Upload(context echo.Context) error{
	client,_:=config.GetObjectStorageClientConnection()
	input:=s3.PutObjectInput{
		Body:               nil,
		CacheControl:       nil,
		ContentType:        nil,
		StorageClass:       nil,
	}
	_,err:=client.PutObject(&input)

	if(err!=nil){
		return context.JSON(http.StatusBadRequest,"Operation failed!")
	}
	return context.JSON(http.StatusAccepted,"Operation Successful")
}

func GetFile(context echo.Context)error{
	keyname:= context.Param("keyname")
	client,_:=config.GetObjectStorageClientConnection()
	output,err:=client.GetObject(&s3.GetObjectInput{
		Bucket: &config.S3Bucket,
		Key:    &keyname,
	})
	if err!=nil{
		return context.JSON(http.StatusBadRequest,"Operation failed! "+err.Error())
	}
	body, err := ioutil.ReadAll(output.Body)
	return context.JSON(http.StatusAccepted,body)
}