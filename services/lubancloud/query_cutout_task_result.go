package lubancloud

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

// QueryCutoutTaskResult invokes the lubancloud.QueryCutoutTaskResult API synchronously
// api document: https://help.aliyun.com/api/lubancloud/querycutouttaskresult.html
func (client *Client) QueryCutoutTaskResult(request *QueryCutoutTaskResultRequest) (response *QueryCutoutTaskResultResponse, err error) {
	response = CreateQueryCutoutTaskResultResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCutoutTaskResultWithChan invokes the lubancloud.QueryCutoutTaskResult API asynchronously
// api document: https://help.aliyun.com/api/lubancloud/querycutouttaskresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCutoutTaskResultWithChan(request *QueryCutoutTaskResultRequest) (<-chan *QueryCutoutTaskResultResponse, <-chan error) {
	responseChan := make(chan *QueryCutoutTaskResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCutoutTaskResult(request)
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

// QueryCutoutTaskResultWithCallback invokes the lubancloud.QueryCutoutTaskResult API asynchronously
// api document: https://help.aliyun.com/api/lubancloud/querycutouttaskresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCutoutTaskResultWithCallback(request *QueryCutoutTaskResultRequest, callback func(response *QueryCutoutTaskResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCutoutTaskResultResponse
		var err error
		defer close(result)
		response, err = client.QueryCutoutTaskResult(request)
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

// QueryCutoutTaskResultRequest is the request struct for api QueryCutoutTaskResult
type QueryCutoutTaskResultRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// QueryCutoutTaskResultResponse is the response struct for api QueryCutoutTaskResult
type QueryCutoutTaskResultResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	TotalSize   int       `json:"TotalSize" xml:"TotalSize"`
	WaitSize    int       `json:"WaitSize" xml:"WaitSize"`
	SuccessSize int       `json:"SuccessSize" xml:"SuccessSize"`
	FailSize    int       `json:"FailSize" xml:"FailSize"`
	Status      int       `json:"Status" xml:"Status"`
	Pictures    []Picture `json:"Pictures" xml:"Pictures"`
}

// CreateQueryCutoutTaskResultRequest creates a request to invoke QueryCutoutTaskResult API
func CreateQueryCutoutTaskResultRequest() (request *QueryCutoutTaskResultRequest) {
	request = &QueryCutoutTaskResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lubancloud", "2018-05-09", "QueryCutoutTaskResult", "luban", "openAPI")
	return
}

// CreateQueryCutoutTaskResultResponse creates a response to parse from QueryCutoutTaskResult response
func CreateQueryCutoutTaskResultResponse() (response *QueryCutoutTaskResultResponse) {
	response = &QueryCutoutTaskResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}