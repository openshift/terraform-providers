package alidns

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

// DescribeDnsGtmInstanceAddressPools invokes the alidns.DescribeDnsGtmInstanceAddressPools API synchronously
func (client *Client) DescribeDnsGtmInstanceAddressPools(request *DescribeDnsGtmInstanceAddressPoolsRequest) (response *DescribeDnsGtmInstanceAddressPoolsResponse, err error) {
	response = CreateDescribeDnsGtmInstanceAddressPoolsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDnsGtmInstanceAddressPoolsWithChan invokes the alidns.DescribeDnsGtmInstanceAddressPools API asynchronously
func (client *Client) DescribeDnsGtmInstanceAddressPoolsWithChan(request *DescribeDnsGtmInstanceAddressPoolsRequest) (<-chan *DescribeDnsGtmInstanceAddressPoolsResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsGtmInstanceAddressPoolsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsGtmInstanceAddressPools(request)
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

// DescribeDnsGtmInstanceAddressPoolsWithCallback invokes the alidns.DescribeDnsGtmInstanceAddressPools API asynchronously
func (client *Client) DescribeDnsGtmInstanceAddressPoolsWithCallback(request *DescribeDnsGtmInstanceAddressPoolsRequest, callback func(response *DescribeDnsGtmInstanceAddressPoolsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsGtmInstanceAddressPoolsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsGtmInstanceAddressPools(request)
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

// DescribeDnsGtmInstanceAddressPoolsRequest is the request struct for api DescribeDnsGtmInstanceAddressPools
type DescribeDnsGtmInstanceAddressPoolsRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
}

// DescribeDnsGtmInstanceAddressPoolsResponse is the response struct for api DescribeDnsGtmInstanceAddressPools
type DescribeDnsGtmInstanceAddressPoolsResponse struct {
	*responses.BaseResponse
	RequestId  string                                        `json:"RequestId" xml:"RequestId"`
	TotalItems int                                           `json:"TotalItems" xml:"TotalItems"`
	TotalPages int                                           `json:"TotalPages" xml:"TotalPages"`
	PageNumber int                                           `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                                           `json:"PageSize" xml:"PageSize"`
	AddrPools  AddrPoolsInDescribeDnsGtmInstanceAddressPools `json:"AddrPools" xml:"AddrPools"`
}

// CreateDescribeDnsGtmInstanceAddressPoolsRequest creates a request to invoke DescribeDnsGtmInstanceAddressPools API
func CreateDescribeDnsGtmInstanceAddressPoolsRequest() (request *DescribeDnsGtmInstanceAddressPoolsRequest) {
	request = &DescribeDnsGtmInstanceAddressPoolsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsGtmInstanceAddressPools", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDnsGtmInstanceAddressPoolsResponse creates a response to parse from DescribeDnsGtmInstanceAddressPools response
func CreateDescribeDnsGtmInstanceAddressPoolsResponse() (response *DescribeDnsGtmInstanceAddressPoolsResponse) {
	response = &DescribeDnsGtmInstanceAddressPoolsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
