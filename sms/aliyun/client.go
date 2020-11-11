package aliyun

/**
 * 功能描述: 阿里云API Client
 * @Date: 2020/4/29
 * @author: lixiaoming
 */

import (
"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
"github.com/lexkong/log"
)

type AliAPIClient struct {
	Instance *dysmsapi.Client
}

var ALIClient *AliAPIClient

func NewAPIClient(regionId, ak, sk string) *AliAPIClient  {
	if regionId == "" {
		regionId = "cn-hangzhou"
	}
	client, err := dysmsapi.NewClientWithAccessKey(regionId, ak, sk)
	if err != nil {
		log.Errorf(err, "failed to init alibaba api client.")
	} else {
		return &AliAPIClient{Instance: client}
	}

	return nil
}

func (c *AliAPIClient) Init(regionId, ak, sk string) {
	if regionId == "" {
		regionId = "cn-hangzhou"
	}
	client, err := dysmsapi.NewClientWithAccessKey(regionId, ak, sk)
	if err != nil {
		log.Errorf(err, "failed to init alibaba api client.")
	} else {
		ALIClient = &AliAPIClient{Instance: client}
	}
}

