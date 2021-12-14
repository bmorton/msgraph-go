/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DeviceAppManagementWindowsInformationProtectionPolicyApiService DeviceAppManagementWindowsInformationProtectionPolicyApi service
type DeviceAppManagementWindowsInformationProtectionPolicyApiService service

type ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest struct {
	ctx _context.Context
	ApiService *DeviceAppManagementWindowsInformationProtectionPolicyApiService
	microsoftGraphWindowsInformationProtectionPolicy *MicrosoftGraphWindowsInformationProtectionPolicy
}

// New navigation property
func (r ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest) MicrosoftGraphWindowsInformationProtectionPolicy(microsoftGraphWindowsInformationProtectionPolicy MicrosoftGraphWindowsInformationProtectionPolicy) ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest {
	r.microsoftGraphWindowsInformationProtectionPolicy = &microsoftGraphWindowsInformationProtectionPolicy
	return r
}

func (r ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest) Execute() (MicrosoftGraphWindowsInformationProtectionPolicy, *_nethttp.Response, error) {
	return r.ApiService.DeviceAppManagementCreateWindowsInformationProtectionPoliciesExecute(r)
}

/*
DeviceAppManagementCreateWindowsInformationProtectionPolicies Create new navigation property to windowsInformationProtectionPolicies for deviceAppManagement

Windows information protection for apps running on devices which are not MDM enrolled.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest
*/
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementCreateWindowsInformationProtectionPolicies(ctx _context.Context) ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest {
	return ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphWindowsInformationProtectionPolicy
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementCreateWindowsInformationProtectionPoliciesExecute(r ApiDeviceAppManagementCreateWindowsInformationProtectionPoliciesRequest) (MicrosoftGraphWindowsInformationProtectionPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphWindowsInformationProtectionPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAppManagementWindowsInformationProtectionPolicyApiService.DeviceAppManagementCreateWindowsInformationProtectionPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceAppManagement/windowsInformationProtectionPolicies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphWindowsInformationProtectionPolicy == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphWindowsInformationProtectionPolicy is required and must be specified")
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
	localVarPostBody = r.microsoftGraphWindowsInformationProtectionPolicy
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

type ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest struct {
	ctx _context.Context
	ApiService *DeviceAppManagementWindowsInformationProtectionPolicyApiService
	windowsInformationProtectionPolicyId string
	ifMatch *string
}

// ETag
func (r ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest) IfMatch(ifMatch string) ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceAppManagementDeleteWindowsInformationProtectionPoliciesExecute(r)
}

/*
DeviceAppManagementDeleteWindowsInformationProtectionPolicies Delete navigation property windowsInformationProtectionPolicies for deviceAppManagement

Windows information protection for apps running on devices which are not MDM enrolled.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param windowsInformationProtectionPolicyId key: id of windowsInformationProtectionPolicy
 @return ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest
*/
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementDeleteWindowsInformationProtectionPolicies(ctx _context.Context, windowsInformationProtectionPolicyId string) ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest {
	return ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		windowsInformationProtectionPolicyId: windowsInformationProtectionPolicyId,
	}
}

// Execute executes the request
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementDeleteWindowsInformationProtectionPoliciesExecute(r ApiDeviceAppManagementDeleteWindowsInformationProtectionPoliciesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAppManagementWindowsInformationProtectionPolicyApiService.DeviceAppManagementDeleteWindowsInformationProtectionPolicies")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"windowsInformationProtectionPolicy-id"+"}", _neturl.PathEscape(parameterToString(r.windowsInformationProtectionPolicyId, "")), -1)

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

type ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest struct {
	ctx _context.Context
	ApiService *DeviceAppManagementWindowsInformationProtectionPolicyApiService
	windowsInformationProtectionPolicyId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest) Select_(select_ []string) ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest) Expand(expand []string) ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest) Execute() (MicrosoftGraphWindowsInformationProtectionPolicy, *_nethttp.Response, error) {
	return r.ApiService.DeviceAppManagementGetWindowsInformationProtectionPoliciesExecute(r)
}

/*
DeviceAppManagementGetWindowsInformationProtectionPolicies Get windowsInformationProtectionPolicies from deviceAppManagement

Windows information protection for apps running on devices which are not MDM enrolled.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param windowsInformationProtectionPolicyId key: id of windowsInformationProtectionPolicy
 @return ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest
*/
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementGetWindowsInformationProtectionPolicies(ctx _context.Context, windowsInformationProtectionPolicyId string) ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest {
	return ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		windowsInformationProtectionPolicyId: windowsInformationProtectionPolicyId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphWindowsInformationProtectionPolicy
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementGetWindowsInformationProtectionPoliciesExecute(r ApiDeviceAppManagementGetWindowsInformationProtectionPoliciesRequest) (MicrosoftGraphWindowsInformationProtectionPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphWindowsInformationProtectionPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAppManagementWindowsInformationProtectionPolicyApiService.DeviceAppManagementGetWindowsInformationProtectionPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"windowsInformationProtectionPolicy-id"+"}", _neturl.PathEscape(parameterToString(r.windowsInformationProtectionPolicyId, "")), -1)

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

type ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest struct {
	ctx _context.Context
	ApiService *DeviceAppManagementWindowsInformationProtectionPolicyApiService
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
	select_ *[]string
	expand *[]string
}

// Show only the first n items
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Top(top int32) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Skip(skip int32) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Search(search string) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Filter(filter string) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Count(count bool) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Orderby(orderby []string) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Select_(select_ []string) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Expand(expand []string) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) Execute() (CollectionOfWindowsInformationProtectionPolicy, *_nethttp.Response, error) {
	return r.ApiService.DeviceAppManagementListWindowsInformationProtectionPoliciesExecute(r)
}

/*
DeviceAppManagementListWindowsInformationProtectionPolicies Get windowsInformationProtectionPolicies from deviceAppManagement

Windows information protection for apps running on devices which are not MDM enrolled.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest
*/
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementListWindowsInformationProtectionPolicies(ctx _context.Context) ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest {
	return ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfWindowsInformationProtectionPolicy
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementListWindowsInformationProtectionPoliciesExecute(r ApiDeviceAppManagementListWindowsInformationProtectionPoliciesRequest) (CollectionOfWindowsInformationProtectionPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfWindowsInformationProtectionPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAppManagementWindowsInformationProtectionPolicyApiService.DeviceAppManagementListWindowsInformationProtectionPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceAppManagement/windowsInformationProtectionPolicies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.top != nil {
		localVarQueryParams.Add("$top", parameterToString(*r.top, ""))
	}
	if r.skip != nil {
		localVarQueryParams.Add("$skip", parameterToString(*r.skip, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("$search", parameterToString(*r.search, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("$filter", parameterToString(*r.filter, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("$count", parameterToString(*r.count, ""))
	}
	if r.orderby != nil {
		localVarQueryParams.Add("$orderby", parameterToString(*r.orderby, "csv"))
	}
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

type ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest struct {
	ctx _context.Context
	ApiService *DeviceAppManagementWindowsInformationProtectionPolicyApiService
	windowsInformationProtectionPolicyId string
	microsoftGraphWindowsInformationProtectionPolicy *MicrosoftGraphWindowsInformationProtectionPolicy
}

// New navigation property values
func (r ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest) MicrosoftGraphWindowsInformationProtectionPolicy(microsoftGraphWindowsInformationProtectionPolicy MicrosoftGraphWindowsInformationProtectionPolicy) ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest {
	r.microsoftGraphWindowsInformationProtectionPolicy = &microsoftGraphWindowsInformationProtectionPolicy
	return r
}

func (r ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceAppManagementUpdateWindowsInformationProtectionPoliciesExecute(r)
}

/*
DeviceAppManagementUpdateWindowsInformationProtectionPolicies Update the navigation property windowsInformationProtectionPolicies in deviceAppManagement

Windows information protection for apps running on devices which are not MDM enrolled.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param windowsInformationProtectionPolicyId key: id of windowsInformationProtectionPolicy
 @return ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest
*/
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementUpdateWindowsInformationProtectionPolicies(ctx _context.Context, windowsInformationProtectionPolicyId string) ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest {
	return ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		windowsInformationProtectionPolicyId: windowsInformationProtectionPolicyId,
	}
}

// Execute executes the request
func (a *DeviceAppManagementWindowsInformationProtectionPolicyApiService) DeviceAppManagementUpdateWindowsInformationProtectionPoliciesExecute(r ApiDeviceAppManagementUpdateWindowsInformationProtectionPoliciesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceAppManagementWindowsInformationProtectionPolicyApiService.DeviceAppManagementUpdateWindowsInformationProtectionPolicies")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceAppManagement/windowsInformationProtectionPolicies/{windowsInformationProtectionPolicy-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"windowsInformationProtectionPolicy-id"+"}", _neturl.PathEscape(parameterToString(r.windowsInformationProtectionPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphWindowsInformationProtectionPolicy == nil {
		return nil, reportError("microsoftGraphWindowsInformationProtectionPolicy is required and must be specified")
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
	localVarPostBody = r.microsoftGraphWindowsInformationProtectionPolicy
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
