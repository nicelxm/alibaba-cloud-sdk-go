package nas

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

// CancelRecycleBinJob invokes the nas.CancelRecycleBinJob API synchronously
func (client *Client) CancelRecycleBinJob(request *CancelRecycleBinJobRequest) (response *CancelRecycleBinJobResponse, err error) {
	response = CreateCancelRecycleBinJobResponse()
	err = client.DoAction(request, response)
	return
}

// CancelRecycleBinJobWithChan invokes the nas.CancelRecycleBinJob API asynchronously
func (client *Client) CancelRecycleBinJobWithChan(request *CancelRecycleBinJobRequest) (<-chan *CancelRecycleBinJobResponse, <-chan error) {
	responseChan := make(chan *CancelRecycleBinJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelRecycleBinJob(request)
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

// CancelRecycleBinJobWithCallback invokes the nas.CancelRecycleBinJob API asynchronously
func (client *Client) CancelRecycleBinJobWithCallback(request *CancelRecycleBinJobRequest, callback func(response *CancelRecycleBinJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelRecycleBinJobResponse
		var err error
		defer close(result)
		response, err = client.CancelRecycleBinJob(request)
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

// CancelRecycleBinJobRequest is the request struct for api CancelRecycleBinJob
type CancelRecycleBinJobRequest struct {
	*requests.RpcRequest
	JobId string `position:"Query" name:"JobId"`
}

// CancelRecycleBinJobResponse is the response struct for api CancelRecycleBinJob
type CancelRecycleBinJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelRecycleBinJobRequest creates a request to invoke CancelRecycleBinJob API
func CreateCancelRecycleBinJobRequest() (request *CancelRecycleBinJobRequest) {
	request = &CancelRecycleBinJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CancelRecycleBinJob", "", "")
	request.Method = requests.GET
	return
}

// CreateCancelRecycleBinJobResponse creates a response to parse from CancelRecycleBinJob response
func CreateCancelRecycleBinJobResponse() (response *CancelRecycleBinJobResponse) {
	response = &CancelRecycleBinJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
