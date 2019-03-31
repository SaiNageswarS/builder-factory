package db

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

var OrgCollectionName = "org"

type Org struct {
	Name string //unique
}

type OrgCollection struct {
	C *mongo.Collection
}

// NewOrgCollection creates appropriate indexes and returns
// reference to collection object for Org
func NewOrgCollection(database *mongo.Database) OrgCollection {
	c := database.Collection(OrgCollectionName)

	indexView := c.Indexes()
	_, err := indexView.CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bsonx.Doc{{"name", bsonx.Int32(1)}},
		Options: &options.IndexOptions{Unique: &[]bool{true}[0]},
	})
	if err != nil {
		log.Fatal("Error creating unique index", err)
	}

	return OrgCollection{
		C: c,
	}
}

func (col *OrgCollection) save(org Org) chan interface{} {
	insertedID := make(chan interface{})

	go func() {
		insertResult, err := col.C.InsertOne(context.Background(), org)
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
