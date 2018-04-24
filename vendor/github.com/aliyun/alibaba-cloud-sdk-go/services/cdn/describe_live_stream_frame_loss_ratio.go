package cdn

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

// DescribeLiveStreamFrameLossRatio invokes the cdn.DescribeLiveStreamFrameLossRatio API synchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamframelossratio.html
func (client *Client) DescribeLiveStreamFrameLossRatio(request *DescribeLiveStreamFrameLossRatioRequest) (response *DescribeLiveStreamFrameLossRatioResponse, err error) {
	response = CreateDescribeLiveStreamFrameLossRatioResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamFrameLossRatioWithChan invokes the cdn.DescribeLiveStreamFrameLossRatio API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamframelossratio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamFrameLossRatioWithChan(request *DescribeLiveStreamFrameLossRatioRequest) (<-chan *DescribeLiveStreamFrameLossRatioResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamFrameLossRatioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamFrameLossRatio(request)
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

// DescribeLiveStreamFrameLossRatioWithCallback invokes the cdn.DescribeLiveStreamFrameLossRatio API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamframelossratio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamFrameLossRatioWithCallback(request *DescribeLiveStreamFrameLossRatioRequest, callback func(response *DescribeLiveStreamFrameLossRatioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamFrameLossRatioResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamFrameLossRatio(request)
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

// DescribeLiveStreamFrameLossRatioRequest is the request struct for api DescribeLiveStreamFrameLossRatio
type DescribeLiveStreamFrameLossRatioRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
}

// DescribeLiveStreamFrameLossRatioResponse is the response struct for api DescribeLiveStreamFrameLossRatio
type DescribeLiveStreamFrameLossRatioResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	FrameLossRatioInfos FrameLossRatioInfos `json:"FrameLossRatioInfos" xml:"FrameLossRatioInfos"`
}

// CreateDescribeLiveStreamFrameLossRatioRequest creates a request to invoke DescribeLiveStreamFrameLossRatio API
func CreateDescribeLiveStreamFrameLossRatioRequest() (request *DescribeLiveStreamFrameLossRatioRequest) {
	request = &DescribeLiveStreamFrameLossRatioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameLossRatio", "", "")
	return
}

// CreateDescribeLiveStreamFrameLossRatioResponse creates a response to parse from DescribeLiveStreamFrameLossRatio response
func CreateDescribeLiveStreamFrameLossRatioResponse() (response *DescribeLiveStreamFrameLossRatioResponse) {
	response = &DescribeLiveStreamFrameLossRatioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
