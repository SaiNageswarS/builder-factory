package db

import (
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

var AppCollectionName = "app"

// app name is unique in an org
// Unique index on combination of Name and OrgId
type App struct {
	DocBase      `bson:",inline"`
	Name         string
	Cloud        string
	OrgID        interface{}
	AwsAccessKey string
	AwsSecret    string
	Description  string
}

type AppCollection struct {
	*mongo.Collection
}

// NewAppCollection creates appropriate indexes and returns
// reference to collection object for App
// Should be called once during initialization of server and instance should
// be injected into other classes
func NewAppCollection(database *mongo.Database) *AppCollection {
	appCollection := new(AppCollection)
	appCollection.Collection = database.Collection(AppCollectionName)

	indexView := appCollection.Indexes()
	// app name is unique in an org
	// Unique index on combination of Name and OrgId
	_, err := indexView.CreateOne(Ctx(), mongo.IndexModel{
		Keys:    bsonx.Doc{{"name", bsonx.Int32(1)}, {"orgid", bsonx.Int32(1)}},
		Options: &options.IndexOptions{Unique: &[]bool{true}[0]},
	})
	if err != nil {
		log.Fatal("Error creating unique index", err)
	}

	return appCollection
}

//Save saves new app
func (col *AppCollection) Save(app App) chan interface{} {
	insertedID := make(chan interface{})

	go func() {
		app.PreSave()
		insertResult, err := col.InsertOne(Ctx(), app)
		if err != nil {
			log.Println(err)
			insertedID <- nil
		} else {
			insertedID <- insertResult.InsertedID
		}
		close(insertedID)
	}()
	return insertedID
}
