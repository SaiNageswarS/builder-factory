package dao

import (
	"context"

	"github.com/SaiNageswarS/builder-factory/services/db"
	"github.com/mongodb/mongo-go-driver/bson"
)

func GetOrgByName(mgo db.Db, orgInstance *db.Org, orgName string) error {
	return mgo.OrgCol.Col().FindOne(context.Background(),
		bson.D{{db.OrgFields.NAME, orgName}}).Decode(orgInstance)
}
