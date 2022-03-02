package emr

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

// ClusterInfoInListClusters is a nested struct in emr response
type ClusterInfoInListClusters struct {
	Id                  string             `json:"Id" xml:"Id"`
	Name                string             `json:"Name" xml:"Name"`
	MachineType         string             `json:"MachineType" xml:"MachineType"`
	Type                string             `json:"Type" xml:"Type"`
	CreateTime          int64              `json:"CreateTime" xml:"CreateTime"`
	RunningTime         int                `json:"RunningTime" xml:"RunningTime"`
	Status              string             `json:"Status" xml:"Status"`
	ChargeType          string             `json:"ChargeType" xml:"ChargeType"`
	ExpiredTime         int64              `json:"ExpiredTime" xml:"ExpiredTime"`
	Period              int                `json:"Period" xml:"Period"`
	HasUncompletedOrder bool               `json:"HasUncompletedOrder" xml:"HasUncompletedOrder"`
	OrderList           string             `json:"OrderList" xml:"OrderList"`
	CreateResource      string             `json:"CreateResource" xml:"CreateResource"`
	DepositType         string             `json:"DepositType" xml:"DepositType"`
	MetaStoreType       string             `json:"MetaStoreType" xml:"MetaStoreType"`
	K8sClusterId        string             `json:"K8sClusterId" xml:"K8sClusterId"`
	OperationId         int64              `json:"OperationId" xml:"OperationId"`
	OrderTaskInfo       OrderTaskInfo      `json:"OrderTaskInfo" xml:"OrderTaskInfo"`
	FailReason          FailReason         `json:"FailReason" xml:"FailReason"`
	Tags                TagsInListClusters `json:"Tags" xml:"Tags"`
}
