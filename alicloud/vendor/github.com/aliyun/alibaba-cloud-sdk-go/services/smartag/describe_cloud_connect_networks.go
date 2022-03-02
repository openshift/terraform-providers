package smartag

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeCloudConnectNetworks invokes the smartag.DescribeCloudConnectNetworks API synchronously
func (client *Client) DescribeCloudConnectNetworks(request *DescribeCloudConnectNetworksRequest) (response *DescribeCloudConnectNetworksResponse, err error) {
	response = CreateDescribeCloudConnectNetworksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudConnectNetworksWithChan invokes the smartag.DescribeCloudConnectNetworks API asynchronously
func (client *Client) DescribeCloudConnectNetworksWithChan(request *DescribeCloudConnectNetworksRequest) (<-chan *DescribeCloudConnectNetworksResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudConnectNetworksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudConnectNetworks(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeCloudConnectNetworksWithCallback invokes the smartag.DescribeCloudConnectNetworks API asynchronously
func (client *Client) DescribeCloudConnectNetworksWithCallback(request *DescribeCloudConnectNetworksRequest, callback func(response *DescribeCloudConnectNetworksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudConnectNetworksResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudConnectNetworks(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeCloudConnectNetworksRequest is the request struct for api DescribeCloudConnectNetworks
type DescribeCloudConnectNetworksRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                   `position:"Query" name:"ResourceOwnerId"`
	CcnId                string                             `position:"Query" name:"CcnId"`
	PageNumber           requests.Integer                   `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                             `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer                   `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeCloudConnectNetworksTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                             `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                   `position:"Query" name:"OwnerId"`
	Name                 string                             `position:"Query" name:"Name"`
}

// DescribeCloudConnectNetworksTag is a repeated param struct in DescribeCloudConnectNetworksRequest
type DescribeCloudConnectNetworksTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeCloudConnectNetworksResponse is the response struct for api DescribeCloudConnectNetworks
type DescribeCloudConnectNetworksResponse struct {
	*responses.BaseResponse
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	CloudConnectNetworks CloudConnectNetworks `json:"CloudConnectNetworks" xml:"CloudConnectNetworks"`
}

// CreateDescribeCloudConnectNetworksRequest creates a request to invoke DescribeCloudConnectNetworks API
func CreateDescribeCloudConnectNetworksRequest() (request *DescribeCloudConnectNetworksRequest) {
	request = &DescribeCloudConnectNetworksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeCloudConnectNetworks", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudConnectNetworksResponse creates a response to parse from DescribeCloudConnectNetworks response
func CreateDescribeCloudConnectNetworksResponse() (response *DescribeCloudConnectNetworksResponse) {
	response = &DescribeCloudConnectNetworksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
