package slb

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

// DescribeLoadBalancerHTTPSListenerAttribute invokes the slb.DescribeLoadBalancerHTTPSListenerAttribute API synchronously
func (client *Client) DescribeLoadBalancerHTTPSListenerAttribute(request *DescribeLoadBalancerHTTPSListenerAttributeRequest) (response *DescribeLoadBalancerHTTPSListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerHTTPSListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerHTTPSListenerAttributeWithChan invokes the slb.DescribeLoadBalancerHTTPSListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerHTTPSListenerAttributeWithChan(request *DescribeLoadBalancerHTTPSListenerAttributeRequest) (<-chan *DescribeLoadBalancerHTTPSListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerHTTPSListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerHTTPSListenerAttribute(request)
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

// DescribeLoadBalancerHTTPSListenerAttributeWithCallback invokes the slb.DescribeLoadBalancerHTTPSListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerHTTPSListenerAttributeWithCallback(request *DescribeLoadBalancerHTTPSListenerAttributeRequest, callback func(response *DescribeLoadBalancerHTTPSListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerHTTPSListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerHTTPSListenerAttribute(request)
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

// DescribeLoadBalancerHTTPSListenerAttributeRequest is the request struct for api DescribeLoadBalancerHTTPSListenerAttribute
type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerHTTPSListenerAttributeResponse is the response struct for api DescribeLoadBalancerHTTPSListenerAttribute
type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                                string                                                         `json:"RequestId" xml:"RequestId"`
	ListenerPort                             int                                                            `json:"ListenerPort" xml:"ListenerPort"`
	BackendServerPort                        int                                                            `json:"BackendServerPort" xml:"BackendServerPort"`
	BackendProtocol                          string                                                         `json:"BackendProtocol" xml:"BackendProtocol"`
	Bandwidth                                int                                                            `json:"Bandwidth" xml:"Bandwidth"`
	Status                                   string                                                         `json:"Status" xml:"Status"`
	SecurityStatus                           string                                                         `json:"SecurityStatus" xml:"SecurityStatus"`
	XForwardedFor                            string                                                         `json:"XForwardedFor" xml:"XForwardedFor"`
	Scheduler                                string                                                         `json:"Scheduler" xml:"Scheduler"`
	StickySession                            string                                                         `json:"StickySession" xml:"StickySession"`
	StickySessionType                        string                                                         `json:"StickySessionType" xml:"StickySessionType"`
	CookieTimeout                            int                                                            `json:"CookieTimeout" xml:"CookieTimeout"`
	Cookie                                   string                                                         `json:"Cookie" xml:"Cookie"`
	HealthCheck                              string                                                         `json:"HealthCheck" xml:"HealthCheck"`
	HealthCheckType                          string                                                         `json:"HealthCheckType" xml:"HealthCheckType"`
	HealthCheckDomain                        string                                                         `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	HealthCheckURI                           string                                                         `json:"HealthCheckURI" xml:"HealthCheckURI"`
	HealthyThreshold                         int                                                            `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold                       int                                                            `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckTimeout                       int                                                            `json:"HealthCheckTimeout" xml:"HealthCheckTimeout"`
	HealthCheckInterval                      int                                                            `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthCheckConnectPort                   int                                                            `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckHttpCode                      string                                                         `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	ServerCertificateId                      string                                                         `json:"ServerCertificateId" xml:"ServerCertificateId"`
	CACertificateId                          string                                                         `json:"CACertificateId" xml:"CACertificateId"`
	HealthCheckMethod                        string                                                         `json:"HealthCheckMethod" xml:"HealthCheckMethod"`
	HealthCheckHttpVersion                   string                                                         `json:"HealthCheckHttpVersion" xml:"HealthCheckHttpVersion"`
	MaxConnection                            int                                                            `json:"MaxConnection" xml:"MaxConnection"`
	VServerGroupId                           string                                                         `json:"VServerGroupId" xml:"VServerGroupId"`
	Gzip                                     string                                                         `json:"Gzip" xml:"Gzip"`
	XForwardedForSLBIP                       string                                                         `json:"XForwardedFor_SLBIP" xml:"XForwardedFor_SLBIP"`
	XForwardedForSLBID                       string                                                         `json:"XForwardedFor_SLBID" xml:"XForwardedFor_SLBID"`
	XForwardedForProto                       string                                                         `json:"XForwardedFor_proto" xml:"XForwardedFor_proto"`
	AclId                                    string                                                         `json:"AclId" xml:"AclId"`
	AclType                                  string                                                         `json:"AclType" xml:"AclType"`
	AclStatus                                string                                                         `json:"AclStatus" xml:"AclStatus"`
	VpcIds                                   string                                                         `json:"VpcIds" xml:"VpcIds"`
	RequestTimeout                           int                                                            `json:"RequestTimeout" xml:"RequestTimeout"`
	IdleTimeout                              int                                                            `json:"IdleTimeout" xml:"IdleTimeout"`
	EnableHttp2                              string                                                         `json:"EnableHttp2" xml:"EnableHttp2"`
	TLSCipherPolicy                          string                                                         `json:"TLSCipherPolicy" xml:"TLSCipherPolicy"`
	Description                              string                                                         `json:"Description" xml:"Description"`
	XForwardedForSLBPORT                     string                                                         `json:"XForwardedFor_SLBPORT" xml:"XForwardedFor_SLBPORT"`
	XForwardedForClientSrcPort               string                                                         `json:"XForwardedFor_ClientSrcPort" xml:"XForwardedFor_ClientSrcPort"`
	XForwardedForClientCertSubjectDN         string                                                         `json:"XForwardedFor_ClientCertSubjectDN" xml:"XForwardedFor_ClientCertSubjectDN"`
	XForwardedForClientCertIssuerDN          string                                                         `json:"XForwardedFor_ClientCertIssuerDN" xml:"XForwardedFor_ClientCertIssuerDN"`
	XForwardedForClientCertFingerprint       string                                                         `json:"XForwardedFor_ClientCertFingerprint" xml:"XForwardedFor_ClientCertFingerprint"`
	XForwardedForClientCertClientVerify      string                                                         `json:"XForwardedFor_ClientCertClientVerify" xml:"XForwardedFor_ClientCertClientVerify"`
	XForwardedForClientCertSubjectDNAlias    string                                                         `json:"XForwardedFor_ClientCertSubjectDNAlias" xml:"XForwardedFor_ClientCertSubjectDNAlias"`
	XForwardedForClientCertIssuerDNAlias     string                                                         `json:"XForwardedFor_ClientCertIssuerDNAlias" xml:"XForwardedFor_ClientCertIssuerDNAlias"`
	XForwardedForClientCertFingerprintAlias  string                                                         `json:"XForwardedFor_ClientCertFingerprintAlias" xml:"XForwardedFor_ClientCertFingerprintAlias"`
	XForwardedForClientCertClientVerifyAlias string                                                         `json:"XForwardedFor_ClientCertClientVerifyAlias" xml:"XForwardedFor_ClientCertClientVerifyAlias"`
	AclIds                                   AclIdsInDescribeLoadBalancerHTTPSListenerAttribute             `json:"AclIds" xml:"AclIds"`
	Rules                                    RulesInDescribeLoadBalancerHTTPSListenerAttribute              `json:"Rules" xml:"Rules"`
	DomainExtensions                         DomainExtensionsInDescribeLoadBalancerHTTPSListenerAttribute   `json:"DomainExtensions" xml:"DomainExtensions"`
	ServerCertificates                       ServerCertificatesInDescribeLoadBalancerHTTPSListenerAttribute `json:"ServerCertificates" xml:"ServerCertificates"`
}

// CreateDescribeLoadBalancerHTTPSListenerAttributeRequest creates a request to invoke DescribeLoadBalancerHTTPSListenerAttribute API
func CreateDescribeLoadBalancerHTTPSListenerAttributeRequest() (request *DescribeLoadBalancerHTTPSListenerAttributeRequest) {
	request = &DescribeLoadBalancerHTTPSListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeLoadBalancerHTTPSListenerAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerHTTPSListenerAttributeResponse creates a response to parse from DescribeLoadBalancerHTTPSListenerAttribute response
func CreateDescribeLoadBalancerHTTPSListenerAttributeResponse() (response *DescribeLoadBalancerHTTPSListenerAttributeResponse) {
	response = &DescribeLoadBalancerHTTPSListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
