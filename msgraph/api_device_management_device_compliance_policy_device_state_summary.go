/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi service
type DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService service

type ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService
	ifMatch *string
}

// ETag
func (r ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest) IfMatch(ifMatch string) ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryExecute(r)
}

/*
DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary Delete navigation property deviceCompliancePolicyDeviceStateSummary for deviceManagement

The device compliance state summary for this account.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest
*/
func (a *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService) DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary(ctx _context.Context) ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest {
	return ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService) DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryExecute(r ApiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService.DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCompliancePolicyDeviceStateSummary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifMatch != nil {
		localVarHeaderParams["If-Match"] = parameterToString(*r.ifMatch, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest) Select_(select_ []string) ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest) Expand(expand []string) ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest) Execute() (MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryExecute(r)
}

/*
DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary Get deviceCompliancePolicyDeviceStateSummary from deviceManagement

The device compliance state summary for this account.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest
*/
func (a *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService) DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary(ctx _context.Context) ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest {
	return ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary
func (a *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService) DeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryExecute(r ApiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest) (MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService.DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCompliancePolicyDeviceStateSummary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("$select", parameterToString(*r.select_, "csv"))
	}
	if r.expand != nil {
		localVarQueryParams.Add("$expand", parameterToString(*r.expand, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService
	microsoftGraphDeviceCompliancePolicyDeviceStateSummary *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary
}

// New navigation property values
func (r ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest) MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary(microsoftGraphDeviceCompliancePolicyDeviceStateSummary MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest {
	r.microsoftGraphDeviceCompliancePolicyDeviceStateSummary = &microsoftGraphDeviceCompliancePolicyDeviceStateSummary
	return r
}

func (r ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryExecute(r)
}

/*
DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary Update the navigation property deviceCompliancePolicyDeviceStateSummary in deviceManagement

The device compliance state summary for this account.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest
*/
func (a *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService) DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary(ctx _context.Context) ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest {
	return ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService) DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryExecute(r ApiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApiService.DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCompliancePolicyDeviceStateSummary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDeviceCompliancePolicyDeviceStateSummary == nil {
		return nil, reportError("microsoftGraphDeviceCompliancePolicyDeviceStateSummary is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.microsoftGraphDeviceCompliancePolicyDeviceStateSummary
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
