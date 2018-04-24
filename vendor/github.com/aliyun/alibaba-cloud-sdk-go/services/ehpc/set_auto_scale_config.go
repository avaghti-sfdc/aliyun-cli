package ehpc

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

// SetAutoScaleConfig invokes the ehpc.SetAutoScaleConfig API synchronously
// api document: https://help.aliyun.com/api/ehpc/setautoscaleconfig.html
func (client *Client) SetAutoScaleConfig(request *SetAutoScaleConfigRequest) (response *SetAutoScaleConfigResponse, err error) {
	response = CreateSetAutoScaleConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetAutoScaleConfigWithChan invokes the ehpc.SetAutoScaleConfig API asynchronously
// api document: https://help.aliyun.com/api/ehpc/setautoscaleconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetAutoScaleConfigWithChan(request *SetAutoScaleConfigRequest) (<-chan *SetAutoScaleConfigResponse, <-chan error) {
	responseChan := make(chan *SetAutoScaleConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetAutoScaleConfig(request)
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

// SetAutoScaleConfigWithCallback invokes the ehpc.SetAutoScaleConfig API asynchronously
// api document: https://help.aliyun.com/api/ehpc/setautoscaleconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetAutoScaleConfigWithCallback(request *SetAutoScaleConfigRequest, callback func(response *SetAutoScaleConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetAutoScaleConfigResponse
		var err error
		defer close(result)
		response, err = client.SetAutoScaleConfig(request)
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

// SetAutoScaleConfigRequest is the request struct for api SetAutoScaleConfig
type SetAutoScaleConfigRequest struct {
	*requests.RpcRequest
	ClusterId               string           `position:"Query" name:"ClusterId"`
	EnableAutoGrow          requests.Boolean `position:"Query" name:"EnableAutoGrow"`
	EnableAutoShrink        requests.Boolean `position:"Query" name:"EnableAutoShrink"`
	GrowIntervalInMinutes   requests.Integer `position:"Query" name:"GrowIntervalInMinutes"`
	ShrinkIntervalInMinutes requests.Integer `position:"Query" name:"ShrinkIntervalInMinutes"`
	ShrinkIdleTimes         requests.Integer `position:"Query" name:"ShrinkIdleTimes"`
	GrowTimeoutInMinutes    requests.Integer `position:"Query" name:"GrowTimeoutInMinutes"`
	ExtraNodesGrowRatio     requests.Integer `position:"Query" name:"ExtraNodesGrowRatio"`
	GrowRatio               requests.Integer `position:"Query" name:"GrowRatio"`
	MaxNodesInCluster       requests.Integer `position:"Query" name:"MaxNodesInCluster"`
	ExcludeNodes            string           `position:"Query" name:"ExcludeNodes"`
}

// SetAutoScaleConfigResponse is the response struct for api SetAutoScaleConfig
type SetAutoScaleConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetAutoScaleConfigRequest creates a request to invoke SetAutoScaleConfig API
func CreateSetAutoScaleConfigRequest() (request *SetAutoScaleConfigRequest) {
	request = &SetAutoScaleConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "SetAutoScaleConfig", "ehs", "openAPI")
	return
}

// CreateSetAutoScaleConfigResponse creates a response to parse from SetAutoScaleConfig response
func CreateSetAutoScaleConfigResponse() (response *SetAutoScaleConfigResponse) {
	response = &SetAutoScaleConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
