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

// UpdateSlrConfiguration invokes the idrsservice.UpdateSlrConfiguration API synchronously
func (client *Client) UpdateSlrConfiguration(request *UpdateSlrConfigurationRequest) (response *UpdateSlrConfigurationResponse, err error) {
	response = CreateUpdateSlrConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSlrConfigurationWithChan invokes the idrsservice.UpdateSlrConfiguration API asynchronously
func (client *Client) UpdateSlrConfigurationWithChan(request *UpdateSlrConfigurationRequest) (<-chan *UpdateSlrConfigurationResponse, <-chan error) {
	responseChan := make(chan *UpdateSlrConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSlrConfiguration(request)
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

// UpdateSlrConfigurationWithCallback invokes the idrsservice.UpdateSlrConfiguration API asynchronously
func (client *Client) UpdateSlrConfigurationWithCallback(request *UpdateSlrConfigurationRequest, callback func(response *UpdateSlrConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSlrConfigurationResponse
		var err error
		defer close(result)
		response, err = client.UpdateSlrConfiguration(request)
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

// UpdateSlrConfigurationRequest is the request struct for api UpdateSlrConfiguration
type UpdateSlrConfigurationRequest struct {
	*requests.RpcRequest
	MqInstanceId string           `position:"Query" name:"MqInstanceId"`
	MqGroupId    string           `position:"Query" name:"MqGroupId"`
	MqEvent      *[]string        `position:"Query" name:"MqEvent"  type:"Repeated"`
	MqEndpoint   string           `position:"Query" name:"MqEndpoint"`
	MqTopic      string           `position:"Query" name:"MqTopic"`
	MqSubscribe  requests.Boolean `position:"Query" name:"MqSubscribe"`
}

// UpdateSlrConfigurationResponse is the response struct for api UpdateSlrConfiguration
type UpdateSlrConfigurationResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateSlrConfigurationRequest creates a request to invoke UpdateSlrConfiguration API
func CreateUpdateSlrConfigurationRequest() (request *UpdateSlrConfigurationRequest) {
	request = &UpdateSlrConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "UpdateSlrConfiguration", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateSlrConfigurationResponse creates a response to parse from UpdateSlrConfiguration response
func CreateUpdateSlrConfigurationResponse() (response *UpdateSlrConfigurationResponse) {
	response = &UpdateSlrConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
