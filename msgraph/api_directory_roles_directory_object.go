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

// DirectoryRolesDirectoryObjectApiService DirectoryRolesDirectoryObjectApi service
type DirectoryRolesDirectoryObjectApiService service

type ApiDirectoryRolesCreateRefMembersRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryObjectApiService
	directoryRoleId string
	requestBody *map[string]map[string]interface{}
}

// New navigation property ref value
func (r ApiDirectoryRolesCreateRefMembersRequest) RequestBody(requestBody map[string]map[string]interface{}) ApiDirectoryRolesCreateRefMembersRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiDirectoryRolesCreateRefMembersRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesCreateRefMembersExecute(r)
}

/*
DirectoryRolesCreateRefMembers Create new navigation property ref to members for directoryRoles

Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesCreateRefMembersRequest
*/
func (a *DirectoryRolesDirectoryObjectApiService) DirectoryRolesCreateRefMembers(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesCreateRefMembersRequest {
	return ApiDirectoryRolesCreateRefMembersRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DirectoryRolesDirectoryObjectApiService) DirectoryRolesCreateRefMembersExecute(r ApiDirectoryRolesCreateRefMembersRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryObjectApiService.DirectoryRolesCreateRefMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/members/$ref"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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

type ApiDirectoryRolesListMembersRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryObjectApiService
	directoryRoleId string
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
func (r ApiDirectoryRolesListMembersRequest) Top(top int32) ApiDirectoryRolesListMembersRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDirectoryRolesListMembersRequest) Skip(skip int32) ApiDirectoryRolesListMembersRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDirectoryRolesListMembersRequest) Search(search string) ApiDirectoryRolesListMembersRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDirectoryRolesListMembersRequest) Filter(filter string) ApiDirectoryRolesListMembersRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDirectoryRolesListMembersRequest) Count(count bool) ApiDirectoryRolesListMembersRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDirectoryRolesListMembersRequest) Orderby(orderby []string) ApiDirectoryRolesListMembersRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDirectoryRolesListMembersRequest) Select_(select_ []string) ApiDirectoryRolesListMembersRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDirectoryRolesListMembersRequest) Expand(expand []string) ApiDirectoryRolesListMembersRequest {
	r.expand = &expand
	return r
}

func (r ApiDirectoryRolesListMembersRequest) Execute() (CollectionOfDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesListMembersExecute(r)
}

/*
DirectoryRolesListMembers Get members from directoryRoles

Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesListMembersRequest
*/
func (a *DirectoryRolesDirectoryObjectApiService) DirectoryRolesListMembers(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesListMembersRequest {
	return ApiDirectoryRolesListMembersRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return CollectionOfDirectoryObject
func (a *DirectoryRolesDirectoryObjectApiService) DirectoryRolesListMembersExecute(r ApiDirectoryRolesListMembersRequest) (CollectionOfDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryObjectApiService.DirectoryRolesListMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

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

type ApiDirectoryRolesListRefMembersRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryObjectApiService
	directoryRoleId string
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
}

// Show only the first n items
func (r ApiDirectoryRolesListRefMembersRequest) Top(top int32) ApiDirectoryRolesListRefMembersRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDirectoryRolesListRefMembersRequest) Skip(skip int32) ApiDirectoryRolesListRefMembersRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDirectoryRolesListRefMembersRequest) Search(search string) ApiDirectoryRolesListRefMembersRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDirectoryRolesListRefMembersRequest) Filter(filter string) ApiDirectoryRolesListRefMembersRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDirectoryRolesListRefMembersRequest) Count(count bool) ApiDirectoryRolesListRefMembersRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDirectoryRolesListRefMembersRequest) Orderby(orderby []string) ApiDirectoryRolesListRefMembersRequest {
	r.orderby = &orderby
	return r
}

func (r ApiDirectoryRolesListRefMembersRequest) Execute() (CollectionOfLinksOfDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesListRefMembersExecute(r)
}

/*
DirectoryRolesListRefMembers Get ref of members from directoryRoles

Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesListRefMembersRequest
*/
func (a *DirectoryRolesDirectoryObjectApiService) DirectoryRolesListRefMembers(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesListRefMembersRequest {
	return ApiDirectoryRolesListRefMembersRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return CollectionOfLinksOfDirectoryObject
func (a *DirectoryRolesDirectoryObjectApiService) DirectoryRolesListRefMembersExecute(r ApiDirectoryRolesListRefMembersRequest) (CollectionOfLinksOfDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfLinksOfDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryObjectApiService.DirectoryRolesListRefMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/members/$ref"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

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
