package controller

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"

	pb "github.com/SaiNageswarS/builder-factory/model/services"
	"github.com/SaiNageswarS/builder-factory/services/db"
	odm "github.com/SaiNageswarS/mongo-odm"
)

//Provides api to manage an application.
//An application is a collection of components
type ApplicationMgmtController struct {
	Mgo db.Db
}

/**
* Creates a new application in DB
 */
func (s *ApplicationMgmtController) Create(ctx context.Context,
	in *pb.ApplicationDetails) (*pb.ApplicationCreateResponse, error) {
	log.Printf("ApplicationMgmtService::Create received %v", in)
	var org db.Org
	org := s.Mgo.Org.Col().FindOne(context.Background(),
		bson.D{{db.OrgFields.NAME, in.OrgName}}).Decode(&org)

	res, err := s.Mgo.App.Col().InsertOne(context.Background(),
		db.App{DocBase: *new(odm.DocBase).New(), Name: in.ApplicationName, Cloud: in.Cloud, OrgID: org.ID, AwsAccessKey: in.AwsAccessKey, AwsSecret: in.AwsSecret, Description: in.Description})

	if err != nil {
		return nil, err
	} 
	return &pb.ApplicationCreateResponse{ApplicationName: in.ApplicationName 
		ApplicationId: res.InsertedID}, nil
}
