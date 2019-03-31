package controller

import (
	"context"
	"log"
)

//Provides api to manage an application.
//An application is a collection of components
type ApplicationMgmtController struct {
	AppMgmtDb db.AppMgmtDb
	OrgDb     db.OrgDb
}