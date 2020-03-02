package config

import (
	"github.com/joho/godotenv"
	"os"
)


var DatabaseConnectionString string
var DatabaseName string
var Port string
var ObjectStorageEndpoint string
var ObjectStorageSecretKey string
var ObjectStorageAccessKey string
var RunMode string
var S3Bucket string
func InitEnvironmentVariables() {
	godotenv.Load()

	DatabaseConnectionString = os.Getenv("DATABASE_CONNECTION_STR")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Port=os.Getenv("SERVER_PORT")
	ObjectStorageEndpoint=os.Getenv("OBJECT_STORAGE_ENDPOINT")
	ObjectStorageSecretKey=os.Getenv("OBJECT_STORAGE_USER_SECRET_KEY")
	ObjectStorageAccessKey=os.Getenv("OBJECT_STORAGE_USER_ACCESS_KEY")
	S3Bucket=os.Getenv("S3_BUCKET")
}

