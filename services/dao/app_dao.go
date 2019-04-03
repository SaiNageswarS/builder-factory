package dao

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/SaiNageswarS/builder-factory/services/db"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func InsertOneApp(mgo db.Db, app db.App) (*mongo.InsertOneResult, error) {
	return mgo.AppCol.Col().InsertOne(context.Background(), app)
}

func GetAllAppsForAnOrg(mgo db.Db, orgId primitive.ObjectID) (*mongo.Cursor, error) {
	return mgo.AppCol.Col().Find(context.Background(),
		bson.D{{db.AppFields.ORG_ID, orgId}})
}
