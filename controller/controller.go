package options

import (
	models "GenCodeX/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Options struct {
	ProgLang string
	Topics   []string
}

func GetLangs(client *mongo.Client) []Options {
	//  get all languages from db and return
	collection := client.Database("boilerplates").Collection("proglangs")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	var allResults []Options

	for cursor.Next(context.Background()) {
		var result Options
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%+v\n", result)
		allResults = append(allResults, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.Background())
	return allResults
}

func GetTemplateSuggestions(client *mongo.Client, progLang string, topic string) []models.BoilerPlate {
	//  get all languages from db and return
	collection := client.Database("boilerplates").Collection("templates")
	cursor, err := collection.Find(context.Background(), bson.M{"topics": topic, "progLang": progLang})
	if err != nil {
		panic(err)
	}
	var allResults []models.BoilerPlate

	for cursor.Next(context.Background()) {
		var result models.BoilerPlate
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		allResults = append(allResults, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.Background())
	return allResults
}
