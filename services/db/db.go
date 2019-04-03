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

//AppFields provides namespace for filters
var AppFields = struct {
	ID             string
	NAME           string
	CLOUD          string
	ORG_ID         string
	AWS_ACCESS_KEY string
	AWS_SECRET     string
	DESCRIPTION    string
}{"_id", "name", "cloud", "orgid", "awsaccesskey", "awssecret", "description"}

type Org struct {
	odm.DocBase `bson:",inline"`
	Name        string `odmIndex:"unique"`
}

//OrgFields provides namespace for filters
var OrgFields = struct {
	ID   string
	NAME string
}{"_id", "name"}

type Db struct {
	OrgCol odm.Collection
	AppCol odm.Collection
}

func NewDb() Db {
	dbInstance := odm.GetDatabase(config.GetString("dbURL"),
		config.GetString("database"))
	return Db{
		OrgCol: odm.Collection{Db: dbInstance, Doc: (*Org)(nil)},
		AppCol: odm.Collection{Db: dbInstance, Doc: (*App)(nil)},
	}
}
