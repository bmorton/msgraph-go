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

// DirectoryRolesDirectoryRoleApiService DirectoryRolesDirectoryRoleApi service
type DirectoryRolesDirectoryRoleApiService service

type ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryRoleApiService
	microsoftGraphDirectoryRole *MicrosoftGraphDirectoryRole
}

// New entity
func (r ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest) MicrosoftGraphDirectoryRole(microsoftGraphDirectoryRole MicrosoftGraphDirectoryRole) ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest {
	r.microsoftGraphDirectoryRole = &microsoftGraphDirectoryRole
	return r
}

func (r ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest) Execute() (MicrosoftGraphDirectoryRole, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleCreateDirectoryRoleExecute(r)
}

/*
DirectoryRolesDirectoryRoleCreateDirectoryRole Add new entity to directoryRoles

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest
*/
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleCreateDirectoryRole(ctx _context.Context) ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest {
	return ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDirectoryRole
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleCreateDirectoryRoleExecute(r ApiDirectoryRolesDirectoryRoleCreateDirectoryRoleRequest) (MicrosoftGraphDirectoryRole, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDirectoryRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryRoleApiService.DirectoryRolesDirectoryRoleCreateDirectoryRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDirectoryRole == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphDirectoryRole is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDirectoryRole
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

type ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryRoleApiService
	directoryRoleId string
	ifMatch *string
}

// ETag
func (r ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest) IfMatch(ifMatch string) ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleDeleteDirectoryRoleExecute(r)
}

/*
DirectoryRolesDirectoryRoleDeleteDirectoryRole Delete entity from directoryRoles

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest
*/
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleDeleteDirectoryRole(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest {
	return ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleDeleteDirectoryRoleExecute(r ApiDirectoryRolesDirectoryRoleDeleteDirectoryRoleRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryRoleApiService.DirectoryRolesDirectoryRoleDeleteDirectoryRole")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

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

type ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryRoleApiService
	directoryRoleId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest) Select_(select_ []string) ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest) Expand(expand []string) ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest {
	r.expand = &expand
	return r
}

func (r ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest) Execute() (MicrosoftGraphDirectoryRole, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleGetDirectoryRoleExecute(r)
}

/*
DirectoryRolesDirectoryRoleGetDirectoryRole Get entity from directoryRoles by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest
*/
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleGetDirectoryRole(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest {
	return ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDirectoryRole
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleGetDirectoryRoleExecute(r ApiDirectoryRolesDirectoryRoleGetDirectoryRoleRequest) (MicrosoftGraphDirectoryRole, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDirectoryRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryRoleApiService.DirectoryRolesDirectoryRoleGetDirectoryRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

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

type ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryRoleApiService
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
	select_ *[]string
	expand *[]string
}

// Skip the first n items
func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Skip(skip int32) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Search(search string) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Filter(filter string) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Count(count bool) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Orderby(orderby []string) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Select_(select_ []string) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Expand(expand []string) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	r.expand = &expand
	return r
}

func (r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) Execute() (CollectionOfDirectoryRole, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleListDirectoryRoleExecute(r)
}

/*
DirectoryRolesDirectoryRoleListDirectoryRole Get entities from directoryRoles

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest
*/
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleListDirectoryRole(ctx _context.Context) ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest {
	return ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfDirectoryRole
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleListDirectoryRoleExecute(r ApiDirectoryRolesDirectoryRoleListDirectoryRoleRequest) (CollectionOfDirectoryRole, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfDirectoryRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryRoleApiService.DirectoryRolesDirectoryRoleListDirectoryRole")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesDirectoryRoleApiService
	directoryRoleId string
	microsoftGraphDirectoryRole *MicrosoftGraphDirectoryRole
}

// New property values
func (r ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest) MicrosoftGraphDirectoryRole(microsoftGraphDirectoryRole MicrosoftGraphDirectoryRole) ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest {
	r.microsoftGraphDirectoryRole = &microsoftGraphDirectoryRole
	return r
}

func (r ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleUpdateDirectoryRoleExecute(r)
}

/*
DirectoryRolesDirectoryRoleUpdateDirectoryRole Update entity in directoryRoles

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest
*/
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleUpdateDirectoryRole(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest {
	return ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
func (a *DirectoryRolesDirectoryRoleApiService) DirectoryRolesDirectoryRoleUpdateDirectoryRoleExecute(r ApiDirectoryRolesDirectoryRoleUpdateDirectoryRoleRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesDirectoryRoleApiService.DirectoryRolesDirectoryRoleUpdateDirectoryRole")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDirectoryRole == nil {
		return nil, reportError("microsoftGraphDirectoryRole is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDirectoryRole
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