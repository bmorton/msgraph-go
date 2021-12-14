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

// GroupsExtensionApiService GroupsExtensionApi service
type GroupsExtensionApiService service

type ApiGroupsCreateExtensionsRequest struct {
	ctx _context.Context
	ApiService *GroupsExtensionApiService
	groupId string
	microsoftGraphExtension *MicrosoftGraphExtension
}

// New navigation property
func (r ApiGroupsCreateExtensionsRequest) MicrosoftGraphExtension(microsoftGraphExtension MicrosoftGraphExtension) ApiGroupsCreateExtensionsRequest {
	r.microsoftGraphExtension = &microsoftGraphExtension
	return r
}

func (r ApiGroupsCreateExtensionsRequest) Execute() (MicrosoftGraphExtension, *_nethttp.Response, error) {
	return r.ApiService.GroupsCreateExtensionsExecute(r)
}

/*
GroupsCreateExtensions Create new navigation property to extensions for groups

The collection of open extensions defined for the group. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @return ApiGroupsCreateExtensionsRequest
*/
func (a *GroupsExtensionApiService) GroupsCreateExtensions(ctx _context.Context, groupId string) ApiGroupsCreateExtensionsRequest {
	return ApiGroupsCreateExtensionsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExtension
func (a *GroupsExtensionApiService) GroupsCreateExtensionsExecute(r ApiGroupsCreateExtensionsRequest) (MicrosoftGraphExtension, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExtension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsExtensionApiService.GroupsCreateExtensions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}/extensions"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphExtension == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphExtension is required and must be specified")
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
	localVarPostBody = r.microsoftGraphExtension
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

type ApiGroupsDeleteExtensionsRequest struct {
	ctx _context.Context
	ApiService *GroupsExtensionApiService
	groupId string
	extensionId string
	ifMatch *string
}

// ETag
func (r ApiGroupsDeleteExtensionsRequest) IfMatch(ifMatch string) ApiGroupsDeleteExtensionsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiGroupsDeleteExtensionsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GroupsDeleteExtensionsExecute(r)
}

/*
GroupsDeleteExtensions Delete navigation property extensions for groups

The collection of open extensions defined for the group. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @param extensionId key: id of extension
 @return ApiGroupsDeleteExtensionsRequest
*/
func (a *GroupsExtensionApiService) GroupsDeleteExtensions(ctx _context.Context, groupId string, extensionId string) ApiGroupsDeleteExtensionsRequest {
	return ApiGroupsDeleteExtensionsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		extensionId: extensionId,
	}
}

// Execute executes the request
func (a *GroupsExtensionApiService) GroupsDeleteExtensionsExecute(r ApiGroupsDeleteExtensionsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsExtensionApiService.GroupsDeleteExtensions")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}/extensions/{extension-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extension-id"+"}", _neturl.PathEscape(parameterToString(r.extensionId, "")), -1)

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

type ApiGroupsGetExtensionsRequest struct {
	ctx _context.Context
	ApiService *GroupsExtensionApiService
	groupId string
	extensionId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiGroupsGetExtensionsRequest) Select_(select_ []string) ApiGroupsGetExtensionsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiGroupsGetExtensionsRequest) Expand(expand []string) ApiGroupsGetExtensionsRequest {
	r.expand = &expand
	return r
}

func (r ApiGroupsGetExtensionsRequest) Execute() (MicrosoftGraphExtension, *_nethttp.Response, error) {
	return r.ApiService.GroupsGetExtensionsExecute(r)
}

/*
GroupsGetExtensions Get extensions from groups

The collection of open extensions defined for the group. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @param extensionId key: id of extension
 @return ApiGroupsGetExtensionsRequest
*/
func (a *GroupsExtensionApiService) GroupsGetExtensions(ctx _context.Context, groupId string, extensionId string) ApiGroupsGetExtensionsRequest {
	return ApiGroupsGetExtensionsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		extensionId: extensionId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExtension
func (a *GroupsExtensionApiService) GroupsGetExtensionsExecute(r ApiGroupsGetExtensionsRequest) (MicrosoftGraphExtension, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExtension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsExtensionApiService.GroupsGetExtensions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}/extensions/{extension-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extension-id"+"}", _neturl.PathEscape(parameterToString(r.extensionId, "")), -1)

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

type ApiGroupsListExtensionsRequest struct {
	ctx _context.Context
	ApiService *GroupsExtensionApiService
	groupId string
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
func (r ApiGroupsListExtensionsRequest) Top(top int32) ApiGroupsListExtensionsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiGroupsListExtensionsRequest) Skip(skip int32) ApiGroupsListExtensionsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiGroupsListExtensionsRequest) Search(search string) ApiGroupsListExtensionsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiGroupsListExtensionsRequest) Filter(filter string) ApiGroupsListExtensionsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiGroupsListExtensionsRequest) Count(count bool) ApiGroupsListExtensionsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiGroupsListExtensionsRequest) Orderby(orderby []string) ApiGroupsListExtensionsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiGroupsListExtensionsRequest) Select_(select_ []string) ApiGroupsListExtensionsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiGroupsListExtensionsRequest) Expand(expand []string) ApiGroupsListExtensionsRequest {
	r.expand = &expand
	return r
}

func (r ApiGroupsListExtensionsRequest) Execute() (CollectionOfExtension, *_nethttp.Response, error) {
	return r.ApiService.GroupsListExtensionsExecute(r)
}

/*
GroupsListExtensions Get extensions from groups

The collection of open extensions defined for the group. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @return ApiGroupsListExtensionsRequest
*/
func (a *GroupsExtensionApiService) GroupsListExtensions(ctx _context.Context, groupId string) ApiGroupsListExtensionsRequest {
	return ApiGroupsListExtensionsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return CollectionOfExtension
func (a *GroupsExtensionApiService) GroupsListExtensionsExecute(r ApiGroupsListExtensionsRequest) (CollectionOfExtension, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfExtension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsExtensionApiService.GroupsListExtensions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}/extensions"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)

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

type ApiGroupsUpdateExtensionsRequest struct {
	ctx _context.Context
	ApiService *GroupsExtensionApiService
	groupId string
	extensionId string
	microsoftGraphExtension *MicrosoftGraphExtension
}

// New navigation property values
func (r ApiGroupsUpdateExtensionsRequest) MicrosoftGraphExtension(microsoftGraphExtension MicrosoftGraphExtension) ApiGroupsUpdateExtensionsRequest {
	r.microsoftGraphExtension = &microsoftGraphExtension
	return r
}

func (r ApiGroupsUpdateExtensionsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GroupsUpdateExtensionsExecute(r)
}

/*
GroupsUpdateExtensions Update the navigation property extensions in groups

The collection of open extensions defined for the group. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId key: id of group
 @param extensionId key: id of extension
 @return ApiGroupsUpdateExtensionsRequest
*/
func (a *GroupsExtensionApiService) GroupsUpdateExtensions(ctx _context.Context, groupId string, extensionId string) ApiGroupsUpdateExtensionsRequest {
	return ApiGroupsUpdateExtensionsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		extensionId: extensionId,
	}
}

// Execute executes the request
func (a *GroupsExtensionApiService) GroupsUpdateExtensionsExecute(r ApiGroupsUpdateExtensionsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsExtensionApiService.GroupsUpdateExtensions")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{group-id}/extensions/{extension-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"group-id"+"}", _neturl.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extension-id"+"}", _neturl.PathEscape(parameterToString(r.extensionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphExtension == nil {
		return nil, reportError("microsoftGraphExtension is required and must be specified")
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
	localVarPostBody = r.microsoftGraphExtension
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
