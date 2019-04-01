package db

import (
	"context"
	"log"

	odm "github.com/SaiNageswarS/mongo-odm"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// app name is unique in an org
// Unique index on combination of Name and OrgId
type App struct {
	odm.DocBase `bson:",inline"`
	Name        string
	Cloud       string
	//OrgID is indexed
	OrgID        interface{} `odmIndex:""`
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
	appCollection.Collection = odm.CreateCollection(database, (*App)(nil))
	return appCollection
}

//Save saves new app
func (col *AppCollection) Save(app App) chan interface{} {
	insertedID := make(chan interface{})

	go func() {
		app.PreSave()
		insertResult, err := col.InsertOne(context.Background(), app)
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
