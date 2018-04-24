package mts

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

// ReportTerrorismJobResult invokes the mts.ReportTerrorismJobResult API synchronously
// api document: https://help.aliyun.com/api/mts/reportterrorismjobresult.html
func (client *Client) ReportTerrorismJobResult(request *ReportTerrorismJobResultRequest) (response *ReportTerrorismJobResultResponse, err error) {
	response = CreateReportTerrorismJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportTerrorismJobResultWithChan invokes the mts.ReportTerrorismJobResult API asynchronously
// api document: https://help.aliyun.com/api/mts/reportterrorismjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportTerrorismJobResultWithChan(request *ReportTerrorismJobResultRequest) (<-chan *ReportTerrorismJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportTerrorismJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportTerrorismJobResult(request)
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

// ReportTerrorismJobResultWithCallback invokes the mts.ReportTerrorismJobResult API asynchronously
// api document: https://help.aliyun.com/api/mts/reportterrorismjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportTerrorismJobResultWithCallback(request *ReportTerrorismJobResultRequest, callback func(response *ReportTerrorismJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportTerrorismJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportTerrorismJobResult(request)
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

// ReportTerrorismJobResultRequest is the request struct for api ReportTerrorismJobResult
type ReportTerrorismJobResultRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
	Label                string           `position:"Query" name:"Label"`
	Detail               string           `position:"Query" name:"Detail"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// ReportTerrorismJobResultResponse is the response struct for api ReportTerrorismJobResult
type ReportTerrorismJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportTerrorismJobResultRequest creates a request to invoke ReportTerrorismJobResult API
func CreateReportTerrorismJobResultRequest() (request *ReportTerrorismJobResultRequest) {
	request = &ReportTerrorismJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportTerrorismJobResult", "mts", "openAPI")
	return
}

// CreateReportTerrorismJobResultResponse creates a response to parse from ReportTerrorismJobResult response
func CreateReportTerrorismJobResultResponse() (response *ReportTerrorismJobResultResponse) {
	response = &ReportTerrorismJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
