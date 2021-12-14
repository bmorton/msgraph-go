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

// DeviceManagementDeviceAndAppManagementRoleAssignmentApiService DeviceManagementDeviceAndAppManagementRoleAssignmentApi service
type DeviceManagementDeviceAndAppManagementRoleAssignmentApiService service

type ApiDeviceManagementCreateRoleAssignmentsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService
	microsoftGraphDeviceAndAppManagementRoleAssignment *MicrosoftGraphDeviceAndAppManagementRoleAssignment
}

// New navigation property
func (r ApiDeviceManagementCreateRoleAssignmentsRequest) MicrosoftGraphDeviceAndAppManagementRoleAssignment(microsoftGraphDeviceAndAppManagementRoleAssignment MicrosoftGraphDeviceAndAppManagementRoleAssignment) ApiDeviceManagementCreateRoleAssignmentsRequest {
	r.microsoftGraphDeviceAndAppManagementRoleAssignment = &microsoftGraphDeviceAndAppManagementRoleAssignment
	return r
}

func (r ApiDeviceManagementCreateRoleAssignmentsRequest) Execute() (MicrosoftGraphDeviceAndAppManagementRoleAssignment, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementCreateRoleAssignmentsExecute(r)
}

/*
DeviceManagementCreateRoleAssignments Create new navigation property to roleAssignments for deviceManagement

The Role Assignments.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementCreateRoleAssignmentsRequest
*/
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementCreateRoleAssignments(ctx _context.Context) ApiDeviceManagementCreateRoleAssignmentsRequest {
	return ApiDeviceManagementCreateRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDeviceAndAppManagementRoleAssignment
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementCreateRoleAssignmentsExecute(r ApiDeviceManagementCreateRoleAssignmentsRequest) (MicrosoftGraphDeviceAndAppManagementRoleAssignment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDeviceAndAppManagementRoleAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceAndAppManagementRoleAssignmentApiService.DeviceManagementCreateRoleAssignments")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/roleAssignments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDeviceAndAppManagementRoleAssignment == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphDeviceAndAppManagementRoleAssignment is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDeviceAndAppManagementRoleAssignment
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

type ApiDeviceManagementDeleteRoleAssignmentsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService
	deviceAndAppManagementRoleAssignmentId string
	ifMatch *string
}

// ETag
func (r ApiDeviceManagementDeleteRoleAssignmentsRequest) IfMatch(ifMatch string) ApiDeviceManagementDeleteRoleAssignmentsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeviceManagementDeleteRoleAssignmentsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementDeleteRoleAssignmentsExecute(r)
}

/*
DeviceManagementDeleteRoleAssignments Delete navigation property roleAssignments for deviceManagement

The Role Assignments.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceAndAppManagementRoleAssignmentId key: id of deviceAndAppManagementRoleAssignment
 @return ApiDeviceManagementDeleteRoleAssignmentsRequest
*/
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementDeleteRoleAssignments(ctx _context.Context, deviceAndAppManagementRoleAssignmentId string) ApiDeviceManagementDeleteRoleAssignmentsRequest {
	return ApiDeviceManagementDeleteRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		deviceAndAppManagementRoleAssignmentId: deviceAndAppManagementRoleAssignmentId,
	}
}

// Execute executes the request
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementDeleteRoleAssignmentsExecute(r ApiDeviceManagementDeleteRoleAssignmentsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceAndAppManagementRoleAssignmentApiService.DeviceManagementDeleteRoleAssignments")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceAndAppManagementRoleAssignment-id"+"}", _neturl.PathEscape(parameterToString(r.deviceAndAppManagementRoleAssignmentId, "")), -1)

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

type ApiDeviceManagementGetRoleAssignmentsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService
	deviceAndAppManagementRoleAssignmentId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiDeviceManagementGetRoleAssignmentsRequest) Select_(select_ []string) ApiDeviceManagementGetRoleAssignmentsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementGetRoleAssignmentsRequest) Expand(expand []string) ApiDeviceManagementGetRoleAssignmentsRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementGetRoleAssignmentsRequest) Execute() (MicrosoftGraphDeviceAndAppManagementRoleAssignment, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementGetRoleAssignmentsExecute(r)
}

/*
DeviceManagementGetRoleAssignments Get roleAssignments from deviceManagement

The Role Assignments.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceAndAppManagementRoleAssignmentId key: id of deviceAndAppManagementRoleAssignment
 @return ApiDeviceManagementGetRoleAssignmentsRequest
*/
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementGetRoleAssignments(ctx _context.Context, deviceAndAppManagementRoleAssignmentId string) ApiDeviceManagementGetRoleAssignmentsRequest {
	return ApiDeviceManagementGetRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		deviceAndAppManagementRoleAssignmentId: deviceAndAppManagementRoleAssignmentId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDeviceAndAppManagementRoleAssignment
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementGetRoleAssignmentsExecute(r ApiDeviceManagementGetRoleAssignmentsRequest) (MicrosoftGraphDeviceAndAppManagementRoleAssignment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDeviceAndAppManagementRoleAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceAndAppManagementRoleAssignmentApiService.DeviceManagementGetRoleAssignments")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceAndAppManagementRoleAssignment-id"+"}", _neturl.PathEscape(parameterToString(r.deviceAndAppManagementRoleAssignmentId, "")), -1)

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

type ApiDeviceManagementListRoleAssignmentsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService
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
func (r ApiDeviceManagementListRoleAssignmentsRequest) Top(top int32) ApiDeviceManagementListRoleAssignmentsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDeviceManagementListRoleAssignmentsRequest) Skip(skip int32) ApiDeviceManagementListRoleAssignmentsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDeviceManagementListRoleAssignmentsRequest) Search(search string) ApiDeviceManagementListRoleAssignmentsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDeviceManagementListRoleAssignmentsRequest) Filter(filter string) ApiDeviceManagementListRoleAssignmentsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDeviceManagementListRoleAssignmentsRequest) Count(count bool) ApiDeviceManagementListRoleAssignmentsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDeviceManagementListRoleAssignmentsRequest) Orderby(orderby []string) ApiDeviceManagementListRoleAssignmentsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDeviceManagementListRoleAssignmentsRequest) Select_(select_ []string) ApiDeviceManagementListRoleAssignmentsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementListRoleAssignmentsRequest) Expand(expand []string) ApiDeviceManagementListRoleAssignmentsRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementListRoleAssignmentsRequest) Execute() (CollectionOfDeviceAndAppManagementRoleAssignment, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementListRoleAssignmentsExecute(r)
}

/*
DeviceManagementListRoleAssignments Get roleAssignments from deviceManagement

The Role Assignments.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementListRoleAssignmentsRequest
*/
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementListRoleAssignments(ctx _context.Context) ApiDeviceManagementListRoleAssignmentsRequest {
	return ApiDeviceManagementListRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfDeviceAndAppManagementRoleAssignment
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementListRoleAssignmentsExecute(r ApiDeviceManagementListRoleAssignmentsRequest) (CollectionOfDeviceAndAppManagementRoleAssignment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfDeviceAndAppManagementRoleAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceAndAppManagementRoleAssignmentApiService.DeviceManagementListRoleAssignments")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/roleAssignments"

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

type ApiDeviceManagementUpdateRoleAssignmentsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService
	deviceAndAppManagementRoleAssignmentId string
	microsoftGraphDeviceAndAppManagementRoleAssignment *MicrosoftGraphDeviceAndAppManagementRoleAssignment
}

// New navigation property values
func (r ApiDeviceManagementUpdateRoleAssignmentsRequest) MicrosoftGraphDeviceAndAppManagementRoleAssignment(microsoftGraphDeviceAndAppManagementRoleAssignment MicrosoftGraphDeviceAndAppManagementRoleAssignment) ApiDeviceManagementUpdateRoleAssignmentsRequest {
	r.microsoftGraphDeviceAndAppManagementRoleAssignment = &microsoftGraphDeviceAndAppManagementRoleAssignment
	return r
}

func (r ApiDeviceManagementUpdateRoleAssignmentsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementUpdateRoleAssignmentsExecute(r)
}

/*
DeviceManagementUpdateRoleAssignments Update the navigation property roleAssignments in deviceManagement

The Role Assignments.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceAndAppManagementRoleAssignmentId key: id of deviceAndAppManagementRoleAssignment
 @return ApiDeviceManagementUpdateRoleAssignmentsRequest
*/
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementUpdateRoleAssignments(ctx _context.Context, deviceAndAppManagementRoleAssignmentId string) ApiDeviceManagementUpdateRoleAssignmentsRequest {
	return ApiDeviceManagementUpdateRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		deviceAndAppManagementRoleAssignmentId: deviceAndAppManagementRoleAssignmentId,
	}
}

// Execute executes the request
func (a *DeviceManagementDeviceAndAppManagementRoleAssignmentApiService) DeviceManagementUpdateRoleAssignmentsExecute(r ApiDeviceManagementUpdateRoleAssignmentsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceAndAppManagementRoleAssignmentApiService.DeviceManagementUpdateRoleAssignments")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceAndAppManagementRoleAssignment-id"+"}", _neturl.PathEscape(parameterToString(r.deviceAndAppManagementRoleAssignmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDeviceAndAppManagementRoleAssignment == nil {
		return nil, reportError("microsoftGraphDeviceAndAppManagementRoleAssignment is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDeviceAndAppManagementRoleAssignment
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
