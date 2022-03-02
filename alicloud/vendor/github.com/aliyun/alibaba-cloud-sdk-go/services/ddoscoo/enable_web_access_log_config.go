package ddoscoo

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

// EnableWebAccessLogConfig invokes the ddoscoo.EnableWebAccessLogConfig API synchronously
func (client *Client) EnableWebAccessLogConfig(request *EnableWebAccessLogConfigRequest) (response *EnableWebAccessLogConfigResponse, err error) {
	response = CreateEnableWebAccessLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// EnableWebAccessLogConfigWithChan invokes the ddoscoo.EnableWebAccessLogConfig API asynchronously
func (client *Client) EnableWebAccessLogConfigWithChan(request *EnableWebAccessLogConfigRequest) (<-chan *EnableWebAccessLogConfigResponse, <-chan error) {
	responseChan := make(chan *EnableWebAccessLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableWebAccessLogConfig(request)
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

// EnableWebAccessLogConfigWithCallback invokes the ddoscoo.EnableWebAccessLogConfig API asynchronously
func (client *Client) EnableWebAccessLogConfigWithCallback(request *EnableWebAccessLogConfigRequest, callback func(response *EnableWebAccessLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableWebAccessLogConfigResponse
		var err error
		defer close(result)
		response, err = client.EnableWebAccessLogConfig(request)
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

// EnableWebAccessLogConfigRequest is the request struct for api EnableWebAccessLogConfig
type EnableWebAccessLogConfigRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
	Lang            string `position:"Query" name:"Lang"`
}

// EnableWebAccessLogConfigResponse is the response struct for api EnableWebAccessLogConfig
type EnableWebAccessLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableWebAccessLogConfigRequest creates a request to invoke EnableWebAccessLogConfig API
func CreateEnableWebAccessLogConfigRequest() (request *EnableWebAccessLogConfigRequest) {
	request = &EnableWebAccessLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "EnableWebAccessLogConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableWebAccessLogConfigResponse creates a response to parse from EnableWebAccessLogConfig response
func CreateEnableWebAccessLogConfigResponse() (response *EnableWebAccessLogConfigResponse) {
	response = &EnableWebAccessLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
