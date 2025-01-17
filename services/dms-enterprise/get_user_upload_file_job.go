package dms_enterprise

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

// GetUserUploadFileJob invokes the dms_enterprise.GetUserUploadFileJob API synchronously
func (client *Client) GetUserUploadFileJob(request *GetUserUploadFileJobRequest) (response *GetUserUploadFileJobResponse, err error) {
	response = CreateGetUserUploadFileJobResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserUploadFileJobWithChan invokes the dms_enterprise.GetUserUploadFileJob API asynchronously
func (client *Client) GetUserUploadFileJobWithChan(request *GetUserUploadFileJobRequest) (<-chan *GetUserUploadFileJobResponse, <-chan error) {
	responseChan := make(chan *GetUserUploadFileJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserUploadFileJob(request)
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

// GetUserUploadFileJobWithCallback invokes the dms_enterprise.GetUserUploadFileJob API asynchronously
func (client *Client) GetUserUploadFileJobWithCallback(request *GetUserUploadFileJobRequest, callback func(response *GetUserUploadFileJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserUploadFileJobResponse
		var err error
		defer close(result)
		response, err = client.GetUserUploadFileJob(request)
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

// GetUserUploadFileJobRequest is the request struct for api GetUserUploadFileJob
type GetUserUploadFileJobRequest struct {
	*requests.RpcRequest
	JobKey string           `position:"Query" name:"JobKey"`
	Tid    requests.Integer `position:"Query" name:"Tid"`
}

// GetUserUploadFileJobResponse is the response struct for api GetUserUploadFileJob
type GetUserUploadFileJobResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	ErrorCode           string              `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage        string              `json:"ErrorMessage" xml:"ErrorMessage"`
	Success             bool                `json:"Success" xml:"Success"`
	UploadFileJobDetail UploadFileJobDetail `json:"UploadFileJobDetail" xml:"UploadFileJobDetail"`
}

// CreateGetUserUploadFileJobRequest creates a request to invoke GetUserUploadFileJob API
func CreateGetUserUploadFileJobRequest() (request *GetUserUploadFileJobRequest) {
	request = &GetUserUploadFileJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetUserUploadFileJob", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserUploadFileJobResponse creates a response to parse from GetUserUploadFileJob response
func CreateGetUserUploadFileJobResponse() (response *GetUserUploadFileJobResponse) {
	response = &GetUserUploadFileJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
