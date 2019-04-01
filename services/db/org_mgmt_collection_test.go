package db

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrg(t *testing.T) {
	mgoDb := GetDatabaseWithParams("mongodb://localhost:27017", "testDB")
	mgoDb.Collection("org").Drop(context.Background())
	orgCollection := NewOrgCollection(mgoDb)

	Convey("Org save should save with unique name", t, func() {
		orgID := <-orgCollection.Save(Org{Name: "testOrg"})
		So(orgID, ShouldNotBeNil)

		// insert same name again should return error
		newOrgID := orgCollection.Save(Org{Name: "testOrg"})
		So(<-newOrgID, ShouldBeNil)
	})

	Convey("GetOrgByName should return return org given name", t, func() {
		fakeOrg := "testingGet"
		orgID := orgCollection.Save(Org{Name: fakeOrg})
		So(<-orgID, ShouldNotBeNil)

		// insert same name again should return error
		orgGet := <-orgCollection.GetOrgByName(fakeOrg)
		fmt.Printf("\nGot object as %v\n", orgGet)
		So(orgGet.Name, ShouldEqual, fakeOrg)
	})
}
