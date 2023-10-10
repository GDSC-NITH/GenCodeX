package db

import (
	config "GenCodeX/handler"
	"context"
	"log"
	"sync"

	"github.com/adhocore/chin"
	"github.com/fatih/color"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {
	uri := config.DotEnvVariable("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	var wg sync.WaitGroup
	s := chin.New().WithWait(&wg)

	color.Blue("🌏 Connecting to Database")
	go s.Start()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		s.Stop()
		wg.Wait()
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		s.Stop()
		wg.Wait()
		color.Red("⛒ Connection Failed to Database \n")
		color.Yellow("⛒ Check your connection please")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	// Send a ping to confirm a successful connection
	if err := client.Database("boolerplates").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	// fmt.Println("Pinged your Database.")
	s.Stop()
	wg.Wait()
	return client
}
