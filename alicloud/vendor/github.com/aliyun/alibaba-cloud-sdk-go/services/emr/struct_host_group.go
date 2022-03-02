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

// HostGroup is a nested struct in emr response
type HostGroup struct {
	HostGroupType            string  `json:"HostGroupType" xml:"HostGroupType"`
	MemoryCapacity           int     `json:"MemoryCapacity" xml:"MemoryCapacity"`
	LockReason               string  `json:"LockReason" xml:"LockReason"`
	HostGroupSubType         string  `json:"HostGroupSubType" xml:"HostGroupSubType"`
	CpuCore                  int     `json:"CpuCore" xml:"CpuCore"`
	SecurityGroupId          string  `json:"SecurityGroupId" xml:"SecurityGroupId"`
	SystemDiskCount          int     `json:"SystemDiskCount" xml:"SystemDiskCount"`
	ChargeType               string  `json:"ChargeType" xml:"ChargeType"`
	SysDiskCapacity          int     `json:"SysDiskCapacity" xml:"SysDiskCapacity"`
	SystemDiskSize           int     `json:"SystemDiskSize" xml:"SystemDiskSize"`
	InstanceType             string  `json:"InstanceType" xml:"InstanceType"`
	DataDiskSize             int     `json:"DataDiskSize" xml:"DataDiskSize"`
	HostGroupChangeStatus    string  `json:"HostGroupChangeStatus" xml:"HostGroupChangeStatus"`
	NodeCount                int     `json:"NodeCount" xml:"NodeCount"`
	GmtModified              string  `json:"gmtModified" xml:"gmtModified"`
	DiskCount                int     `json:"DiskCount" xml:"DiskCount"`
	Comment                  string  `json:"Comment" xml:"Comment"`
	VswitchId                string  `json:"VswitchId" xml:"VswitchId"`
	HostGroupName            string  `json:"HostGroupName" xml:"HostGroupName"`
	ScalingGroupMinNode      int     `json:"ScalingGroupMinNode" xml:"ScalingGroupMinNode"`
	DiskCapacity             int     `json:"DiskCapacity" xml:"DiskCapacity"`
	ScalingGroupBizId        string  `json:"ScalingGroupBizId" xml:"ScalingGroupBizId"`
	ScalingInMode            string  `json:"ScalingInMode" xml:"ScalingInMode"`
	BandWidth                string  `json:"BandWidth" xml:"BandWidth"`
	Memory                   int     `json:"Memory" xml:"Memory"`
	DiskType                 string  `json:"DiskType" xml:"DiskType"`
	PayType                  string  `json:"PayType" xml:"PayType"`
	LockType                 string  `json:"LockType" xml:"LockType"`
	MultiInstanceTypes       string  `json:"MultiInstanceTypes" xml:"MultiInstanceTypes"`
	Period                   string  `json:"Period" xml:"Period"`
	GmtCreate                string  `json:"gmtCreate" xml:"gmtCreate"`
	CostSavingPercent        float64 `json:"CostSavingPercent" xml:"CostSavingPercent"`
	DataDiskType             string  `json:"DataDiskType" xml:"DataDiskType"`
	ScalingGroupActiveStatus string  `json:"ScalingGroupActiveStatus" xml:"ScalingGroupActiveStatus"`
	ScalingGroupConfigState  string  `json:"ScalingGroupConfigState" xml:"ScalingGroupConfigState"`
	SysDiskType              string  `json:"SysDiskType" xml:"SysDiskType"`
	HostGroupChangeType      string  `json:"HostGroupChangeType" xml:"HostGroupChangeType"`
	Cpu                      int     `json:"Cpu" xml:"Cpu"`
	Status                   string  `json:"Status" xml:"Status"`
	HostGroupId              string  `json:"HostGroupId" xml:"HostGroupId"`
	DataDiskCount            int     `json:"DataDiskCount" xml:"DataDiskCount"`
	ScalingGroupMaxNode      int     `json:"ScalingGroupMaxNode" xml:"ScalingGroupMaxNode"`
	SystemDiskType           string  `json:"SystemDiskType" xml:"SystemDiskType"`
	Nodes                    Nodes   `json:"Nodes" xml:"Nodes"`
}
