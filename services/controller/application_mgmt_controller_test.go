package controller

import (
	"context"
	"log"
	"testing"

	odm "github.com/SaiNageswarS/mongo-odm"

	pb "github.com/SaiNageswarS/builder-factory/model/services"
	"github.com/SaiNageswarS/builder-factory/services/db"
	config "github.com/spf13/viper"
)

func getApplicationDetailProto(orgName string, appName string) *pb.ApplicationDetails {
	return &pb.ApplicationDetails{ApplicationName: appName,
		OrgName:      orgName,
		Cloud:        "testAws",
		AwsAccessKey: "tst",
		AwsSecret:    "test",
		Description:  "just work",
	}
}

func SetUp() db.Db {
	config.Set("dbURL", "mongodb://localhost:27017")
	config.Set("database", "testDB")
	return db.NewDb()
}

func TestApplicationMgmtService(t *testing.T) {
	db := SetUp()

	// insert org
	org := db.Org{
		DocBase: *(new(odm.DocBase).New()),
		Name:    "testOrg",
	}

	insertedOrg, _ := db.Org.Col().InsertOne(context.Background(), org)
	log.Printf("Org is %v", insertedOrg.InsertedID)

	Convey("CreateApplication", t, func() {
		app1, _ := appMgmtService.Create(nil, getApplicationDetailProto(org.Name, "testApp1"))
		app2, _ := appMgmtService.Create(nil, getApplicationDetailProto(org.Name, "testApp2"))

		if app1.ApplicationId <= 0 || app2.ApplicationId <= 0 {
			t.Errorf("Expected app1.ApplicationId>0 && app2.ApplicationId>0. Actual values: %v, %v\n", app1, app2)
		}
	})
}
