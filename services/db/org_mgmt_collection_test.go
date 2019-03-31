package db

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrg(t *testing.T) {
	mgoDb := GetDatabaseWithParams("mongodb://localhost:27017", "testDB")
	mgoDb.Collection(OrgCollectionName).Drop(context.Background())

	Convey("Org save should save with unique name", t, func() {
		orgCollection := NewOrgCollection(mgoDb)
		orgID := orgCollection.save(Org{Name: "testOrg"})
		So(<-orgID, ShouldNotBeNil)

		// insert same name again should return error
		newOrgID := orgCollection.save(Org{Name: "testOrg"})
		So(<-newOrgID, ShouldBeNil)
	})
}
