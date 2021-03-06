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

// GroupsGroupApiService GroupsGroupApi service
type GroupsGroupApiService service

type ApiGroupsGroupCreateGroupRequest struct {
	ctx _context.Context
	ApiService *GroupsGroupApiService
	microsoftGraphGroup *MicrosoftGraphGroup
}

// New entity
func (r ApiGroupsGroupCreateGroupRequest) MicrosoftGraphGroup(microsoftGraphGroup MicrosoftGraphGroup) ApiGroupsGroupCreateGroupRequest {
	r.microsoftGraphGroup = &microsoftGraphGroup
	return r
}

func (r ApiGroupsGroupCreateGroupRequest) Execute() (MicrosoftGraphGroup, *_nethttp.Response, error) {
	return r.ApiService.GroupsGroupCreateGroupExecute(r)
}

/*
GroupsGroupCreateGroup Add new entity to groups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGroupsGroupCreateGroupRequest
*/
func (a *GroupsGroupApiService) GroupsGroupCreateGroup(ctx _context.Context) ApiGroupsGroupCreateGroupRequest {
	return ApiGroupsGroupCreateGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphGroup
func (a *GroupsGroupApiService) GroupsGroupCreateGroupExecute(r ApiGroupsGroupCreateGroupRequest) (MicrosoftGraphGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsGroupApiService.GroupsGroupCreateGroup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphGroup == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphGroup is required and must be specified")
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
	localVarPostBody = r.microsoftGraphGroup
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

type ApiGroupsGroupDeleteGroupRequest struct {
	ctx _context.Context
	ApiService *GroupsGroupApiService
	groupId string
	ifMatch *string
}

// ETag
func (r ApiGroupsGroupDeleteGroupRequest) IfMatch(ifMatch string) ApiGroupsGroupDeleteGroupRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiGroupsGroupDeleteGroupRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GroupsGroupDeleteGroupExecute(r)
}

/*
GroupsGroupDeleteGroup Delete entity from groups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @return ApiGroupsGroupDeleteGroupRequest
*/
func (a *GroupsGroupApiService) GroupsGroupDeleteGroup(ctx _context.Context, groupId string) ApiGroupsGroupDeleteGroupRequest {
	return ApiGroupsGroupDeleteGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GroupsGroupApiService) GroupsGroupDeleteGroupExecute(r ApiGroupsGroupDeleteGroupRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsGroupApiService.GroupsGroupDeleteGroup")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiGroupsGroupGetGroupRequest struct {
	ctx _context.Context
	ApiService *GroupsGroupApiService
	groupId string
	consistencyLevel *string
	select_ *[]string
	expand *[]string
}

// Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/
func (r ApiGroupsGroupGetGroupRequest) ConsistencyLevel(consistencyLevel string) ApiGroupsGroupGetGroupRequest {
	r.consistencyLevel = &consistencyLevel
	return r
}
// Select properties to be returned
func (r ApiGroupsGroupGetGroupRequest) Select_(select_ []string) ApiGroupsGroupGetGroupRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiGroupsGroupGetGroupRequest) Expand(expand []string) ApiGroupsGroupGetGroupRequest {
	r.expand = &expand
	return r
}

func (r ApiGroupsGroupGetGroupRequest) Execute() (MicrosoftGraphGroup, *_nethttp.Response, error) {
	return r.ApiService.GroupsGroupGetGroupExecute(r)
}

/*
GroupsGroupGetGroup Get entity from groups by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @return ApiGroupsGroupGetGroupRequest
*/
func (a *GroupsGroupApiService) GroupsGroupGetGroup(ctx _context.Context, groupId string) ApiGroupsGroupGetGroupRequest {
	return ApiGroupsGroupGetGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphGroup
func (a *GroupsGroupApiService) GroupsGroupGetGroupExecute(r ApiGroupsGroupGetGroupRequest) (MicrosoftGraphGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsGroupApiService.GroupsGroupGetGroup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)

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
	if r.consistencyLevel != nil {
		localVarHeaderParams["ConsistencyLevel"] = parameterToString(*r.consistencyLevel, "")
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

type ApiGroupsGroupListGroupRequest struct {
	ctx _context.Context
	ApiService *GroupsGroupApiService
	consistencyLevel *string
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
	select_ *[]string
	expand *[]string
}

// Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/
func (r ApiGroupsGroupListGroupRequest) ConsistencyLevel(consistencyLevel string) ApiGroupsGroupListGroupRequest {
	r.consistencyLevel = &consistencyLevel
	return r
}
// Show only the first n items
func (r ApiGroupsGroupListGroupRequest) Top(top int32) ApiGroupsGroupListGroupRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiGroupsGroupListGroupRequest) Skip(skip int32) ApiGroupsGroupListGroupRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiGroupsGroupListGroupRequest) Search(search string) ApiGroupsGroupListGroupRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiGroupsGroupListGroupRequest) Filter(filter string) ApiGroupsGroupListGroupRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiGroupsGroupListGroupRequest) Count(count bool) ApiGroupsGroupListGroupRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiGroupsGroupListGroupRequest) Orderby(orderby []string) ApiGroupsGroupListGroupRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiGroupsGroupListGroupRequest) Select_(select_ []string) ApiGroupsGroupListGroupRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiGroupsGroupListGroupRequest) Expand(expand []string) ApiGroupsGroupListGroupRequest {
	r.expand = &expand
	return r
}

func (r ApiGroupsGroupListGroupRequest) Execute() (CollectionOfGroup, *_nethttp.Response, error) {
	return r.ApiService.GroupsGroupListGroupExecute(r)
}

/*
GroupsGroupListGroup Get entities from groups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGroupsGroupListGroupRequest
*/
func (a *GroupsGroupApiService) GroupsGroupListGroup(ctx _context.Context) ApiGroupsGroupListGroupRequest {
	return ApiGroupsGroupListGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfGroup
func (a *GroupsGroupApiService) GroupsGroupListGroupExecute(r ApiGroupsGroupListGroupRequest) (CollectionOfGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsGroupApiService.GroupsGroupListGroup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

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
	if r.consistencyLevel != nil {
		localVarHeaderParams["ConsistencyLevel"] = parameterToString(*r.consistencyLevel, "")
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

type ApiGroupsGroupUpdateGroupRequest struct {
	ctx _context.Context
	ApiService *GroupsGroupApiService
	groupId string
	microsoftGraphGroup *MicrosoftGraphGroup
}

// New property values
func (r ApiGroupsGroupUpdateGroupRequest) MicrosoftGraphGroup(microsoftGraphGroup MicrosoftGraphGroup) ApiGroupsGroupUpdateGroupRequest {
	r.microsoftGraphGroup = &microsoftGraphGroup
	return r
}

func (r ApiGroupsGroupUpdateGroupRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GroupsGroupUpdateGroupExecute(r)
}

/*
GroupsGroupUpdateGroup Update entity in groups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @return ApiGroupsGroupUpdateGroupRequest
*/
func (a *GroupsGroupApiService) GroupsGroupUpdateGroup(ctx _context.Context, groupId string) ApiGroupsGroupUpdateGroupRequest {
	return ApiGroupsGroupUpdateGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
func (a *GroupsGroupApiService) GroupsGroupUpdateGroupExecute(r ApiGroupsGroupUpdateGroupRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsGroupApiService.GroupsGroupUpdateGroup")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphGroup == nil {
		return nil, reportError("microsoftGraphGroup is required and must be specified")
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
	localVarPostBody = r.microsoftGraphGroup
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
