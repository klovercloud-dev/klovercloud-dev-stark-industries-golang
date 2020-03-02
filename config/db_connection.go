package config

import (
	"context"
	"fmt"
	"github.com/go-bongo/bongo"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var connectDB *bongo.Connection
var Staff *bongo.Collection

// Connect Database
func InitDBConnection() {
	// DB Connect
	connection, err := CreateConnectionDB()
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return
	}
	connectDB = connection
}

// Initialize Database Collections
func InitDBCollections() {
	Staff = connectDB.Collection("staffs")
}

func CreateConnectionDB() (*bongo.Connection, error) {
	config := &bongo.Config{
		ConnectionString: DatabaseConnectionString,
		Database:         DatabaseName,
	}
	log.Println(DatabaseConnectionString)
	connection, err := bongo.Connect(config)
	return connection, err
}

func CloseConnectionDB(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}
