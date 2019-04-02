package db

import (
	odm "github.com/SaiNageswarS/mongo-odm"
	config "github.com/spf13/viper"
)

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

type Org struct {
	odm.DocBase `bson:",inline"`
	Name        string `odmIndex:"unique"`
}

//OrgFields provides namespace for filters
var OrgFields = struct {
	ID   string
	NAME string
}{"ID", "NAME"}

type Db struct {
	Org odm.Collection
	App odm.Collection
}

func NewDb() Db {
	dbInstance := odm.GetDatabase(config.GetString("dbURL"),
		config.GetString("database"))
	return Db{
		Org: odm.Collection{Db: dbInstance, Doc: (*Org)(nil)},
		App: odm.Collection{Db: dbInstance, Doc: (*App)(nil)},
	}
}