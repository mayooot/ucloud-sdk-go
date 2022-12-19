// Code is generated by ucloud-model, DO NOT EDIT IT.

package udbproxy

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UDBProxy API Schema

// ListUDBProxyClientRequest is request schema for ListUDBProxyClient action
type ListUDBProxyClientRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](https://docs.ucloud.cn/api/summary/get_project_list)
	// ProjectId *string `required:"false"`

	// [公共参数] 地域。 参见 [地域和可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](https://docs.ucloud.cn/api/summary/regionlist)
	// Zone *string `required:"true"`

	// 代理ID
	UDBProxyID *string `required:"true"`
}

// ListUDBProxyClientResponse is response schema for ListUDBProxyClient action
type ListUDBProxyClientResponse struct {
	response.CommonBase

	// 代理节点客户端IP连接信息
	NodeClientInfos []NodeClientInfo

	// 代理ID
	UDBProxyID string
}

// NewListUDBProxyClientRequest will create request of ListUDBProxyClient action.
func (c *UDBProxyClient) NewListUDBProxyClientRequest() *ListUDBProxyClientRequest {
	req := &ListUDBProxyClientRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ListUDBProxyClient

查询代理客户端连接IP信息(实时）
*/
func (c *UDBProxyClient) ListUDBProxyClient(req *ListUDBProxyClientRequest) (*ListUDBProxyClientResponse, error) {
	var err error
	var res ListUDBProxyClientResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ListUDBProxyClient", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
