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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// GroupSettingsGroupSettingApiService GroupSettingsGroupSettingApi service
type GroupSettingsGroupSettingApiService service

type ApiGroupSettingsGroupSettingCreateGroupSettingRequest struct {
	ctx _context.Context
	ApiService *GroupSettingsGroupSettingApiService
	microsoftGraphGroupSetting *MicrosoftGraphGroupSetting
}

// New entity
func (r ApiGroupSettingsGroupSettingCreateGroupSettingRequest) MicrosoftGraphGroupSetting(microsoftGraphGroupSetting MicrosoftGraphGroupSetting) ApiGroupSettingsGroupSettingCreateGroupSettingRequest {
	r.microsoftGraphGroupSetting = &microsoftGraphGroupSetting
	return r
}

func (r ApiGroupSettingsGroupSettingCreateGroupSettingRequest) Execute() (MicrosoftGraphGroupSetting, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingsGroupSettingCreateGroupSettingExecute(r)
}

/*
GroupSettingsGroupSettingCreateGroupSetting Add new entity to groupSettings

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGroupSettingsGroupSettingCreateGroupSettingRequest
*/
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingCreateGroupSetting(ctx _context.Context) ApiGroupSettingsGroupSettingCreateGroupSettingRequest {
	return ApiGroupSettingsGroupSettingCreateGroupSettingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphGroupSetting
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingCreateGroupSettingExecute(r ApiGroupSettingsGroupSettingCreateGroupSettingRequest) (MicrosoftGraphGroupSetting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphGroupSetting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingsGroupSettingApiService.GroupSettingsGroupSettingCreateGroupSetting")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphGroupSetting == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphGroupSetting is required and must be specified")
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
	localVarPostBody = r.microsoftGraphGroupSetting
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

type ApiGroupSettingsGroupSettingDeleteGroupSettingRequest struct {
	ctx _context.Context
	ApiService *GroupSettingsGroupSettingApiService
	groupSettingId string
	ifMatch *string
}

// ETag
func (r ApiGroupSettingsGroupSettingDeleteGroupSettingRequest) IfMatch(ifMatch string) ApiGroupSettingsGroupSettingDeleteGroupSettingRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiGroupSettingsGroupSettingDeleteGroupSettingRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GroupSettingsGroupSettingDeleteGroupSettingExecute(r)
}

/*
GroupSettingsGroupSettingDeleteGroupSetting Delete entity from groupSettings

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingId key: id of groupSetting
 @return ApiGroupSettingsGroupSettingDeleteGroupSettingRequest
*/
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingDeleteGroupSetting(ctx _context.Context, groupSettingId string) ApiGroupSettingsGroupSettingDeleteGroupSettingRequest {
	return ApiGroupSettingsGroupSettingDeleteGroupSettingRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingId: groupSettingId,
	}
}

// Execute executes the request
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingDeleteGroupSettingExecute(r ApiGroupSettingsGroupSettingDeleteGroupSettingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingsGroupSettingApiService.GroupSettingsGroupSettingDeleteGroupSetting")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettings/{groupSetting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSetting-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingId, "")), -1)

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

type ApiGroupSettingsGroupSettingGetGroupSettingRequest struct {
	ctx _context.Context
	ApiService *GroupSettingsGroupSettingApiService
	groupSettingId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiGroupSettingsGroupSettingGetGroupSettingRequest) Select_(select_ []string) ApiGroupSettingsGroupSettingGetGroupSettingRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiGroupSettingsGroupSettingGetGroupSettingRequest) Expand(expand []string) ApiGroupSettingsGroupSettingGetGroupSettingRequest {
	r.expand = &expand
	return r
}

func (r ApiGroupSettingsGroupSettingGetGroupSettingRequest) Execute() (MicrosoftGraphGroupSetting, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingsGroupSettingGetGroupSettingExecute(r)
}

/*
GroupSettingsGroupSettingGetGroupSetting Get entity from groupSettings by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingId key: id of groupSetting
 @return ApiGroupSettingsGroupSettingGetGroupSettingRequest
*/
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingGetGroupSetting(ctx _context.Context, groupSettingId string) ApiGroupSettingsGroupSettingGetGroupSettingRequest {
	return ApiGroupSettingsGroupSettingGetGroupSettingRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingId: groupSettingId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphGroupSetting
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingGetGroupSettingExecute(r ApiGroupSettingsGroupSettingGetGroupSettingRequest) (MicrosoftGraphGroupSetting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphGroupSetting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingsGroupSettingApiService.GroupSettingsGroupSettingGetGroupSetting")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettings/{groupSetting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSetting-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingId, "")), -1)

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

type ApiGroupSettingsGroupSettingListGroupSettingRequest struct {
	ctx _context.Context
	ApiService *GroupSettingsGroupSettingApiService
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
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Top(top int32) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Skip(skip int32) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Search(search string) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Filter(filter string) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Count(count bool) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Orderby(orderby []string) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Select_(select_ []string) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Expand(expand []string) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	r.expand = &expand
	return r
}

func (r ApiGroupSettingsGroupSettingListGroupSettingRequest) Execute() (CollectionOfGroupSetting, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingsGroupSettingListGroupSettingExecute(r)
}

/*
GroupSettingsGroupSettingListGroupSetting Get entities from groupSettings

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGroupSettingsGroupSettingListGroupSettingRequest
*/
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingListGroupSetting(ctx _context.Context) ApiGroupSettingsGroupSettingListGroupSettingRequest {
	return ApiGroupSettingsGroupSettingListGroupSettingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfGroupSetting
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingListGroupSettingExecute(r ApiGroupSettingsGroupSettingListGroupSettingRequest) (CollectionOfGroupSetting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfGroupSetting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingsGroupSettingApiService.GroupSettingsGroupSettingListGroupSetting")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettings"

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

type ApiGroupSettingsGroupSettingUpdateGroupSettingRequest struct {
	ctx _context.Context
	ApiService *GroupSettingsGroupSettingApiService
	groupSettingId string
	microsoftGraphGroupSetting *MicrosoftGraphGroupSetting
}

// New property values
func (r ApiGroupSettingsGroupSettingUpdateGroupSettingRequest) MicrosoftGraphGroupSetting(microsoftGraphGroupSetting MicrosoftGraphGroupSetting) ApiGroupSettingsGroupSettingUpdateGroupSettingRequest {
	r.microsoftGraphGroupSetting = &microsoftGraphGroupSetting
	return r
}

func (r ApiGroupSettingsGroupSettingUpdateGroupSettingRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GroupSettingsGroupSettingUpdateGroupSettingExecute(r)
}

/*
GroupSettingsGroupSettingUpdateGroupSetting Update entity in groupSettings

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingId key: id of groupSetting
 @return ApiGroupSettingsGroupSettingUpdateGroupSettingRequest
*/
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingUpdateGroupSetting(ctx _context.Context, groupSettingId string) ApiGroupSettingsGroupSettingUpdateGroupSettingRequest {
	return ApiGroupSettingsGroupSettingUpdateGroupSettingRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingId: groupSettingId,
	}
}

// Execute executes the request
func (a *GroupSettingsGroupSettingApiService) GroupSettingsGroupSettingUpdateGroupSettingExecute(r ApiGroupSettingsGroupSettingUpdateGroupSettingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingsGroupSettingApiService.GroupSettingsGroupSettingUpdateGroupSetting")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettings/{groupSetting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSetting-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphGroupSetting == nil {
		return nil, reportError("microsoftGraphGroupSetting is required and must be specified")
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
	localVarPostBody = r.microsoftGraphGroupSetting
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
