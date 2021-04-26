package project

import "gin-demo/pkg/util/structutil"

const (
	TableProject = "project_test"
)

type Project struct {
	ProjectId  string `json:"project_id"`
	Name       string `json:"name"`
	UserId     string `json:"user_id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

var ProjectColumns = structutil.GetColumnsFromStruct(Project{})

// project input struct
type APIGetProjectsInput struct {
	UserId    string `json:"user_id" validate:"required"`
	ProjectId string `json:"project_id"`
	Name      string `json:"name"`
	Limit     int    `json:"limit" default:"10"`
	Offset    int    `json:"offset" default:"0"`
}

type APIGetProjectsOutput struct {
	Data       []Project `json:"data"`
	TotalCount uint32    `json:"total_count"`
	RetCode    int       `json:"ret_code"`
}
