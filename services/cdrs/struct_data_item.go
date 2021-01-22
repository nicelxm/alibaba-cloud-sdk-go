package cdrs

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

// DataItem is a nested struct in cdrs response
type DataItem struct {
	TagMetrics    string `json:"TagMetrics" xml:"TagMetrics"`
	DateId        string `json:"DateId" xml:"DateId"`
	Coordinates   string `json:"Coordinates" xml:"Coordinates"`
	UserGroupId   string `json:"UserGroupId" xml:"UserGroupId"`
	PersonId      string `json:"PersonId" xml:"PersonId"`
	Times         string `json:"Times" xml:"Times"`
	PoiName       string `json:"PoiName" xml:"PoiName"`
	Frequency     string `json:"Frequency" xml:"Frequency"`
	IntervalTime  string `json:"IntervalTime" xml:"IntervalTime"`
	DeviceGroupId string `json:"DeviceGroupId" xml:"DeviceGroupId"`
	TagCode       string `json:"TagCode" xml:"TagCode"`
	TagMetric     string `json:"TagMetric" xml:"TagMetric"`
	CorpId        string `json:"CorpId" xml:"CorpId"`
	DeviceId      string `json:"DeviceId" xml:"DeviceId"`
	TagValue      string `json:"TagValue" xml:"TagValue"`
	DateTime      string `json:"DateTime" xml:"DateTime"`
	PoiId         string `json:"PoiId" xml:"PoiId"`
	PassHour      string `json:"PassHour" xml:"PassHour"`
}
