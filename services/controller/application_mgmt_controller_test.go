package controller

import (
	"context"
	"log"
	"testing"

	odm "github.com/SaiNageswarS/mongo-odm"

	pb "github.com/SaiNageswarS/builder-factory/model/services"
	"github.com/SaiNageswarS/builder-factory/services/db"
	. "github.com/smartystreets/goconvey/convey"
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

func TestApplicationMgmtService(t *testing.T) {
	setUp()

	// insert org
	org := db.Org{
		DocBase: *(new(odm.DocBase).New()),
		Name:    "testOrg",
	}

	_, err := mgo.Org.Col().InsertOne(context.Background(), org)
	log.Printf("error is %v", err)

	Convey("CreateApplication", t, func() {
		app1, _ := appMgmtService.Create(context.TODO(), getApplicationDetailProto(org.Name, "testApp1"))
		app2, _ := appMgmtService.Create(context.TODO(), getApplicationDetailProto(org.Name, "testApp2"))

		So(app1, ShouldNotBeNil)
		So(app2, ShouldNotBeNil)
	})
	tearDown()
}
