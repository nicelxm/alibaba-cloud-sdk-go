package retailcloud

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

// BatchAddServers invokes the retailcloud.BatchAddServers API synchronously
func (client *Client) BatchAddServers(request *BatchAddServersRequest) (response *BatchAddServersResponse, err error) {
	response = CreateBatchAddServersResponse()
	err = client.DoAction(request, response)
	return
}

// BatchAddServersWithChan invokes the retailcloud.BatchAddServers API asynchronously
func (client *Client) BatchAddServersWithChan(request *BatchAddServersRequest) (<-chan *BatchAddServersResponse, <-chan error) {
	responseChan := make(chan *BatchAddServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchAddServers(request)
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

// BatchAddServersWithCallback invokes the retailcloud.BatchAddServers API asynchronously
func (client *Client) BatchAddServersWithCallback(request *BatchAddServersRequest, callback func(response *BatchAddServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchAddServersResponse
		var err error
		defer close(result)
		response, err = client.BatchAddServers(request)
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

// BatchAddServersRequest is the request struct for api BatchAddServers
type BatchAddServersRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	VpcId      string `position:"Query" name:"VpcId"`
	Sign       string `position:"Query" name:"Sign"`
}

// BatchAddServersResponse is the response struct for api BatchAddServers
type BatchAddServersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateBatchAddServersRequest creates a request to invoke BatchAddServers API
func CreateBatchAddServersRequest() (request *BatchAddServersRequest) {
	request = &BatchAddServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "BatchAddServers", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchAddServersResponse creates a response to parse from BatchAddServers response
func CreateBatchAddServersResponse() (response *BatchAddServersResponse) {
	response = &BatchAddServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
