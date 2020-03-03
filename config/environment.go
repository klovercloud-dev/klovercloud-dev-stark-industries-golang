package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)


var DatasourceUrl string
var DatabaseName string
var Port string
var ObjectStorageEndpoint string
var ObjectStorageSecretKey string
var ObjectStorageAccessKey string
var RunMode string
var S3Bucket string

func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = DEVELOP
	}

	log.Println("RUN MODE:", RunMode)

	if RunMode != PRODUCTION  && RunMode != TEST{
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}
	DatasourceUrl = os.Getenv("DATASOURCE_URL")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Port=os.Getenv("SERVER_PORT")
	ObjectStorageEndpoint=os.Getenv("OBJECT_STORAGE_ENDPOINT")
	ObjectStorageSecretKey=os.Getenv("OBJECT_STORAGE_USER_SECRET_KEY")
	ObjectStorageAccessKey=os.Getenv("OBJECT_STORAGE_USER_ACCESS_KEY")
	S3Bucket=os.Getenv("S3_BUCKET")
}

