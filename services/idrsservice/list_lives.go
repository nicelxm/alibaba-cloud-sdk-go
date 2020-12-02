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

// ListLives invokes the idrsservice.ListLives API synchronously
func (client *Client) ListLives(request *ListLivesRequest) (response *ListLivesResponse, err error) {
	response = CreateListLivesResponse()
	err = client.DoAction(request, response)
	return
}

// ListLivesWithChan invokes the idrsservice.ListLives API asynchronously
func (client *Client) ListLivesWithChan(request *ListLivesRequest) (<-chan *ListLivesResponse, <-chan error) {
	responseChan := make(chan *ListLivesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLives(request)
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

// ListLivesWithCallback invokes the idrsservice.ListLives API asynchronously
func (client *Client) ListLivesWithCallback(request *ListLivesRequest, callback func(response *ListLivesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLivesResponse
		var err error
		defer close(result)
		response, err = client.ListLives(request)
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

// ListLivesRequest is the request struct for api ListLives
type ListLivesRequest struct {
	*requests.RpcRequest
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	PageIndex requests.Integer `position:"Query" name:"PageIndex"`
}

// ListLivesResponse is the response struct for api ListLives
type ListLivesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListLivesRequest creates a request to invoke ListLives API
func CreateListLivesRequest() (request *ListLivesRequest) {
	request = &ListLivesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "ListLives", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLivesResponse creates a response to parse from ListLives response
func CreateListLivesResponse() (response *ListLivesResponse) {
	response = &ListLivesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
