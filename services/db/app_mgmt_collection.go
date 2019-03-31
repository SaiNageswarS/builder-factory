package db

type AppDetails struct {
	Name         string
	Cloud        string
	OrgId        int64
	AwsAccessKey string
	AwsSecret    string
	Description  string
}
