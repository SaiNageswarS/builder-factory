package controller

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	pb "github.com/SaiNageswarS/builder-factory/model/services"
	"github.com/SaiNageswarS/builder-factory/services/dao"
	"github.com/SaiNageswarS/builder-factory/services/db"
	odm "github.com/SaiNageswarS/mongo-odm"
)

//Provides api to manage an application.
//An application is a collection of components
type ApplicationMgmtController struct {
	Mgo db.Db
}

// Create creates a new application in DB
func (s *ApplicationMgmtController) Create(ctx context.Context,
	in *pb.ApplicationDetails) (*pb.ApplicationCreateResponse, error) {
	log.Printf("ApplicationMgmtService::Create received %v", in)
	var org db.Org
	err := dao.GetOrgByName(s.Mgo, &org, in.OrgName)

	if err != nil {
		log.Printf("Got err fetching org.\n %v", err)
		return nil, err
	}

	res, err := dao.InsertOneApp(s.Mgo,
		db.App{
			DocBase:      *new(odm.DocBase).New(),
			Name:         in.ApplicationName,
			Cloud:        in.Cloud,
			OrgID:        org.ID,
			AwsAccessKey: in.AwsAccessKey,
			AwsSecret:    in.AwsSecret,
			Description:  in.Description})

	if err != nil {
		log.Printf("Got err fetching org.\n %v", err)
		return nil, err
	}
	applicationID := res.InsertedID.(primitive.ObjectID).Hex()
	return &pb.ApplicationCreateResponse{ApplicationName: in.ApplicationName,
		ApplicationId: applicationID}, nil
}

//GetAllApplications returns all applications for an org
func (s *ApplicationMgmtController) GetAllApplications(ctx context.Context,
	in *pb.GetApplicationsRequest) (*pb.GetApplicationsResponse, error) {

	log.Printf("ApplicationMgmtService::Get All Applications for %v", in)

	var org db.Org
	err := dao.GetOrgByName(s.Mgo, &org, in.OrgName)

	if err != nil {
		log.Printf("Got err fetching org.\n %v", err)
		return nil, err
	}

	appsCursor, err := dao.GetAllAppsForAnOrg(s.Mgo, org.ID)

	if err != nil {
		log.Printf("Got err fetching apps.\n %v", err)
		return nil, err
	}

	var applicationDetails []*pb.ApplicationDetails
	for appsCursor.Next(context.Background()) {
		var a db.App
		appsCursor.Decode(&a)

		applicationDetails = append(applicationDetails, &pb.ApplicationDetails{
			ApplicationName: a.Name,
			Cloud:           a.Cloud,
			AwsAccessKey:    a.AwsAccessKey,
			AwsSecret:       a.AwsSecret,
			Description:     a.Description,
		})
	}

	log.Printf("Returning %v", applicationDetails)
	return &pb.GetApplicationsResponse{Applications: applicationDetails}, nil
}
