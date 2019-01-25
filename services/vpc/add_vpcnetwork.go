//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api VPC AddVPCNetwork

package vpc

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// AddVPCNetworkRequest is request schema for AddVPCNetwork action
type AddVPCNetworkRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 源VPC短ID
	VPCId *string `required:"true"`

	// 增加网段
	Network []string `required:"true"`
}

// AddVPCNetworkResponse is response schema for AddVPCNetwork action
type AddVPCNetworkResponse struct {
	response.CommonBase
}

// NewAddVPCNetworkRequest will create request of AddVPCNetwork action.
func (c *VPCClient) NewAddVPCNetworkRequest() *AddVPCNetworkRequest {
	req := &AddVPCNetworkRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// AddVPCNetwork - 添加VPC网段
func (c *VPCClient) AddVPCNetwork(req *AddVPCNetworkRequest) (*AddVPCNetworkResponse, error) {
	var err error
	var res AddVPCNetworkResponse

	err = c.client.InvokeAction("AddVPCNetwork", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
