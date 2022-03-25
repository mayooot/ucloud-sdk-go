// Code is generated by ucloud-model, DO NOT EDIT IT.

package iam

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// IAM API Schema

// CreateProjectRequest is request schema for CreateProject action
type CreateProjectRequest struct {
	request.CommonBase

	// 【该字段已废弃，请谨慎使用】
	ParentId *string `required:"false" deprecated:"true"`

	// 项目名称，不得与现有项目重名
	ProjectName *string `required:"true"`
}

// CreateProjectResponse is response schema for CreateProject action
type CreateProjectResponse struct {
	response.CommonBase

	// 所创建项目的Id
	ProjectId string
}

// NewCreateProjectRequest will create request of CreateProject action.
func (c *IAMClient) NewCreateProjectRequest() *CreateProjectRequest {
	req := &CreateProjectRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateProject

创建项目
*/
func (c *IAMClient) CreateProject(req *CreateProjectRequest) (*CreateProjectResponse, error) {
	var err error
	var res CreateProjectResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateProject", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteProjectRequest is request schema for DeleteProject action
type DeleteProjectRequest struct {
	request.CommonBase

	// 项目ID
	ProjectID *string `required:"true"`
}

// DeleteProjectResponse is response schema for DeleteProject action
type DeleteProjectResponse struct {
	response.CommonBase

	// 错误消息
	Message string
}

// NewDeleteProjectRequest will create request of DeleteProject action.
func (c *IAMClient) NewDeleteProjectRequest() *DeleteProjectRequest {
	req := &DeleteProjectRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteProject

删除项目
*/
func (c *IAMClient) DeleteProject(req *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	var err error
	var res DeleteProjectResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteProject", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
