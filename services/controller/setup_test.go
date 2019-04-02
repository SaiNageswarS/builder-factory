package controller

import (
	"context"

	"github.com/SaiNageswarS/builder-factory/services/db"
	config "github.com/spf13/viper"
)

var appMgmtService ApplicationMgmtController
var mgo db.Db

func setUp() {
	config.Set("dbURL", "mongodb://localhost:27017")
	config.Set("database", "testDB")

	mgo = db.NewDb()

	appMgmtService = ApplicationMgmtController{mgo}
}

func tearDown() {
	mgo.Org.Col().Drop(context.Background())
	mgo.App.Col().Drop(context.Background())
}
