package idrsservice

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

// GetStatistics invokes the idrsservice.GetStatistics API synchronously
func (client *Client) GetStatistics(request *GetStatisticsRequest) (response *GetStatisticsResponse, err error) {
	response = CreateGetStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// GetStatisticsWithChan invokes the idrsservice.GetStatistics API asynchronously
func (client *Client) GetStatisticsWithChan(request *GetStatisticsRequest) (<-chan *GetStatisticsResponse, <-chan error) {
	responseChan := make(chan *GetStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetStatistics(request)
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

// GetStatisticsWithCallback invokes the idrsservice.GetStatistics API asynchronously
func (client *Client) GetStatisticsWithCallback(request *GetStatisticsRequest, callback func(response *GetStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetStatisticsResponse
		var err error
		defer close(result)
		response, err = client.GetStatistics(request)
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

// GetStatisticsRequest is the request struct for api GetStatistics
type GetStatisticsRequest struct {
	*requests.RpcRequest
	DepartmentId *[]string `position:"Query" name:"DepartmentId"  type:"Repeated"`
	DateTo       string    `position:"Query" name:"DateTo"`
	DateFrom     string    `position:"Query" name:"DateFrom"`
}

// GetStatisticsResponse is the response struct for api GetStatistics
type GetStatisticsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetStatisticsRequest creates a request to invoke GetStatistics API
func CreateGetStatisticsRequest() (request *GetStatisticsRequest) {
	request = &GetStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "GetStatistics", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetStatisticsResponse creates a response to parse from GetStatistics response
func CreateGetStatisticsResponse() (response *GetStatisticsResponse) {
	response = &GetStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
