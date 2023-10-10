package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// BoilerPlate -> boilerplate model
type BoilerPlate struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	ProgLang string             `json:"progLang,omitempty" bson:"progLang,omitempty"`
	Topics   []string           `json:"topics,omitempty" bson:"topics,omitempty"`
	Files    []File             `json:"files,omitempty" bson:"files,omitempty"`
}

type File struct {
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Content   string `json:"content,omitempty" bson:"content,omitempty"`
	Extension string `json:"extension,omitempty" bson:"extension,omitempty"`
	PathToGo  string `json:"pathToGo,omitempty" bson:"pathToGo,omitempty"`
}
