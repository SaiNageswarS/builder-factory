package db

import (
	"context"
	"testing"

	odm "github.com/SaiNageswarS/mongo-odm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestApp(t *testing.T) {
	mgoDb := odm.GetDatabaseWithParams("mongodb://localhost:27017", "testDB")
	mgoDb.Collection("app").Drop(context.Background())
	appCollection := NewAppCollection(mgoDb)
	parentOrgId := <-NewOrgCollection(mgoDb).Save(Org{Name: "parentOrg"})

	Convey("Org save should save with unique name", t, func() {
		appId := <-appCollection.Save(App{Name: "testApp", OrgID: parentOrgId, Cloud: "aws"})
		So(appId, ShouldNotBeNil)
	})
}
