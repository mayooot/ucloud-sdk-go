//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDPN GetUDPNPrice

package udpn

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetUDPNPriceRequest is request schema for GetUDPNPrice action
type GetUDPNPriceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"false"`

	// 专线可用区1，支持地域：北京二：cn-bj2, 上海二：cn-sh2, 广东：cn-gd, 亚太： hk, 上海一：cn-sh1, 法兰克福：ge-fra, 新加坡：sg, 洛杉矶：us-la， 华盛顿：us-ws， 东京：jpn-tky
	Peer1 *string `required:"true"`

	// 专线可用区2，支持地域：北京二：cn-bj2, 上海二：cn-sh2, 广东：cn-gd, 亚太： hk, 上海一：cn-sh1, 法兰克福：ge-fra, 新加坡：sg, 洛杉矶：us-la， 华盛顿：us-ws， 东京：jpn-tky
	Peer2 *string `required:"true"`

	// 带宽信息
	Bandwidth *int `required:"true"`

	// 计费类型
	ChargeType *string `required:"false"`

	// 购买时长
	Quantity *int `required:"false"`
}

// GetUDPNPriceResponse is response schema for GetUDPNPrice action
type GetUDPNPriceResponse struct {
	response.CommonBase

	// 资源有效期 unix 时间戳
	PurchaseValue int

	// 专线价格
	Price float64
}

// NewGetUDPNPriceRequest will create request of GetUDPNPrice action.
func (c *UDPNClient) NewGetUDPNPriceRequest() *GetUDPNPriceRequest {
	req := &GetUDPNPriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetUDPNPrice - 获取 UDPN 价格
func (c *UDPNClient) GetUDPNPrice(req *GetUDPNPriceRequest) (*GetUDPNPriceResponse, error) {
	var err error
	var res GetUDPNPriceResponse

	err = c.client.InvokeAction("GetUDPNPrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
