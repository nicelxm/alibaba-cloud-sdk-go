package vcs

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

// ListMotorAlgorithmResults invokes the vcs.ListMotorAlgorithmResults API synchronously
func (client *Client) ListMotorAlgorithmResults(request *ListMotorAlgorithmResultsRequest) (response *ListMotorAlgorithmResultsResponse, err error) {
	response = CreateListMotorAlgorithmResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMotorAlgorithmResultsWithChan invokes the vcs.ListMotorAlgorithmResults API asynchronously
func (client *Client) ListMotorAlgorithmResultsWithChan(request *ListMotorAlgorithmResultsRequest) (<-chan *ListMotorAlgorithmResultsResponse, <-chan error) {
	responseChan := make(chan *ListMotorAlgorithmResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMotorAlgorithmResults(request)
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

// ListMotorAlgorithmResultsWithCallback invokes the vcs.ListMotorAlgorithmResults API asynchronously
func (client *Client) ListMotorAlgorithmResultsWithCallback(request *ListMotorAlgorithmResultsRequest, callback func(response *ListMotorAlgorithmResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMotorAlgorithmResultsResponse
		var err error
		defer close(result)
		response, err = client.ListMotorAlgorithmResults(request)
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

// ListMotorAlgorithmResultsRequest is the request struct for api ListMotorAlgorithmResults
type ListMotorAlgorithmResultsRequest struct {
	*requests.RpcRequest
	AlgorithmType string `position:"Body" name:"AlgorithmType"`
	CorpId        string `position:"Body" name:"CorpId"`
	EndTime       string `position:"Body" name:"EndTime"`
	StartTime     string `position:"Body" name:"StartTime"`
	PageNumber    string `position:"Body" name:"PageNumber"`
	PlateNumber   string `position:"Body" name:"PlateNumber"`
	DataSourceId  string `position:"Body" name:"DataSourceId"`
	PageSize      string `position:"Body" name:"PageSize"`
}

// ListMotorAlgorithmResultsResponse is the response struct for api ListMotorAlgorithmResults
type ListMotorAlgorithmResultsResponse struct {
	*responses.BaseResponse
	Code      string                          `json:"Code" xml:"Code"`
	Message   string                          `json:"Message" xml:"Message"`
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Data      DataInListMotorAlgorithmResults `json:"Data" xml:"Data"`
}

// CreateListMotorAlgorithmResultsRequest creates a request to invoke ListMotorAlgorithmResults API
func CreateListMotorAlgorithmResultsRequest() (request *ListMotorAlgorithmResultsRequest) {
	request = &ListMotorAlgorithmResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListMotorAlgorithmResults", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListMotorAlgorithmResultsResponse creates a response to parse from ListMotorAlgorithmResults response
func CreateListMotorAlgorithmResultsResponse() (response *ListMotorAlgorithmResultsResponse) {
	response = &ListMotorAlgorithmResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
