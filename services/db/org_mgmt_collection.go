package db

import (
	"context"
	"log"

	odm "github.com/SaiNageswarS/mongo-odm"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Org struct {
	odm.DocBase `bson:",inline"`
	Name        string `odm:unique`
}

type OrgCollection struct {
	*mongo.Collection
}

// NewOrgCollection creates appropriate indexes and returns
// reference to collection object for Org
func NewOrgCollection(database *mongo.Database) *OrgCollection {
	orgCollection := new(OrgCollection)
	orgCollection.Collection = odm.CreateCollection(database, (*Org)(nil))

	return orgCollection
}

func (col *OrgCollection) Save(org Org) chan interface{} {
	insertedID := make(chan interface{})

	go func() {
		org.PreSave()
		insertResult, err := col.InsertOne(context.Background(), org)
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

//GetOrgByName returns org given name
func (col *OrgCollection) GetOrgByName(name string) chan Org {
	orgChan := make(chan Org)

	go func() {
		var org Org
		//bson.D{{"name", name}}
		err := col.FindOne(Ctx(), bson.D{{"name", name}}).Decode(&org)
		if err != nil {
			log.Println(err)
		}
		orgChan <- org
		close(orgChan)
	}()
	return orgChan
}
