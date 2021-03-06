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

// DirectoryRolesActionsApiService DirectoryRolesActionsApi service
type DirectoryRolesActionsApiService service

type ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	directoryRoleId string
	inlineObject120 *InlineObject120
}

func (r ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest) InlineObject120(inlineObject120 InlineObject120) ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest {
	r.inlineObject120 = &inlineObject120
	return r
}

func (r ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleCheckMemberGroupsExecute(r)
}

/*
DirectoryRolesDirectoryRoleCheckMemberGroups Invoke action checkMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleCheckMemberGroups(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest {
	return ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleCheckMemberGroupsExecute(r ApiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesDirectoryRoleCheckMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/microsoft.graph.checkMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject120 == nil {
		return localVarReturnValue, nil, reportError("inlineObject120 is required and must be specified")
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
	localVarPostBody = r.inlineObject120
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

type ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	directoryRoleId string
	inlineObject121 *InlineObject121
}

func (r ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest) InlineObject121(inlineObject121 InlineObject121) ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest {
	r.inlineObject121 = &inlineObject121
	return r
}

func (r ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleCheckMemberObjectsExecute(r)
}

/*
DirectoryRolesDirectoryRoleCheckMemberObjects Invoke action checkMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleCheckMemberObjects(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest {
	return ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleCheckMemberObjectsExecute(r ApiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesDirectoryRoleCheckMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/microsoft.graph.checkMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject121 == nil {
		return localVarReturnValue, nil, reportError("inlineObject121 is required and must be specified")
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
	localVarPostBody = r.inlineObject121
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

type ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	directoryRoleId string
	inlineObject122 *InlineObject122
}

func (r ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest) InlineObject122(inlineObject122 InlineObject122) ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest {
	r.inlineObject122 = &inlineObject122
	return r
}

func (r ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleGetMemberGroupsExecute(r)
}

/*
DirectoryRolesDirectoryRoleGetMemberGroups Invoke action getMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleGetMemberGroups(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest {
	return ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleGetMemberGroupsExecute(r ApiDirectoryRolesDirectoryRoleGetMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesDirectoryRoleGetMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/microsoft.graph.getMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject122 == nil {
		return localVarReturnValue, nil, reportError("inlineObject122 is required and must be specified")
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
	localVarPostBody = r.inlineObject122
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

type ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	directoryRoleId string
	inlineObject123 *InlineObject123
}

func (r ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest) InlineObject123(inlineObject123 InlineObject123) ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest {
	r.inlineObject123 = &inlineObject123
	return r
}

func (r ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleGetMemberObjectsExecute(r)
}

/*
DirectoryRolesDirectoryRoleGetMemberObjects Invoke action getMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleGetMemberObjects(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest {
	return ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleGetMemberObjectsExecute(r ApiDirectoryRolesDirectoryRoleGetMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesDirectoryRoleGetMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/microsoft.graph.getMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRole-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject123 == nil {
		return localVarReturnValue, nil, reportError("inlineObject123 is required and must be specified")
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
	localVarPostBody = r.inlineObject123
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

type ApiDirectoryRolesDirectoryRoleRestoreRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	directoryRoleId string
}


func (r ApiDirectoryRolesDirectoryRoleRestoreRequest) Execute() (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesDirectoryRoleRestoreExecute(r)
}

/*
DirectoryRolesDirectoryRoleRestore Invoke action restore

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleId key: id of directoryRole
 @return ApiDirectoryRolesDirectoryRoleRestoreRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleRestore(ctx _context.Context, directoryRoleId string) ApiDirectoryRolesDirectoryRoleRestoreRequest {
	return ApiDirectoryRolesDirectoryRoleRestoreRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleId: directoryRoleId,
	}
}

// Execute executes the request
//  @return AnyOfmicrosoftGraphDirectoryObject
func (a *DirectoryRolesActionsApiService) DirectoryRolesDirectoryRoleRestoreExecute(r ApiDirectoryRolesDirectoryRoleRestoreRequest) (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnyOfmicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesDirectoryRoleRestore")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/{directoryRole-id}/microsoft.graph.restore"
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

type ApiDirectoryRolesGetAvailableExtensionPropertiesRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	inlineObject124 *InlineObject124
}

func (r ApiDirectoryRolesGetAvailableExtensionPropertiesRequest) InlineObject124(inlineObject124 InlineObject124) ApiDirectoryRolesGetAvailableExtensionPropertiesRequest {
	r.inlineObject124 = &inlineObject124
	return r
}

func (r ApiDirectoryRolesGetAvailableExtensionPropertiesRequest) Execute() ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesGetAvailableExtensionPropertiesExecute(r)
}

/*
DirectoryRolesGetAvailableExtensionProperties Invoke action getAvailableExtensionProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRolesGetAvailableExtensionPropertiesRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesGetAvailableExtensionProperties(ctx _context.Context) ApiDirectoryRolesGetAvailableExtensionPropertiesRequest {
	return ApiDirectoryRolesGetAvailableExtensionPropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphExtensionProperty
func (a *DirectoryRolesActionsApiService) DirectoryRolesGetAvailableExtensionPropertiesExecute(r ApiDirectoryRolesGetAvailableExtensionPropertiesRequest) ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphExtensionProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesGetAvailableExtensionProperties")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/microsoft.graph.getAvailableExtensionProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject124 == nil {
		return localVarReturnValue, nil, reportError("inlineObject124 is required and must be specified")
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
	localVarPostBody = r.inlineObject124
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

type ApiDirectoryRolesGetByIdsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	inlineObject125 *InlineObject125
}

func (r ApiDirectoryRolesGetByIdsRequest) InlineObject125(inlineObject125 InlineObject125) ApiDirectoryRolesGetByIdsRequest {
	r.inlineObject125 = &inlineObject125
	return r
}

func (r ApiDirectoryRolesGetByIdsRequest) Execute() ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesGetByIdsExecute(r)
}

/*
DirectoryRolesGetByIds Invoke action getByIds

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRolesGetByIdsRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesGetByIds(ctx _context.Context) ApiDirectoryRolesGetByIdsRequest {
	return ApiDirectoryRolesGetByIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphDirectoryObject
func (a *DirectoryRolesActionsApiService) DirectoryRolesGetByIdsExecute(r ApiDirectoryRolesGetByIdsRequest) ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesGetByIds")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/microsoft.graph.getByIds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject125 == nil {
		return localVarReturnValue, nil, reportError("inlineObject125 is required and must be specified")
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
	localVarPostBody = r.inlineObject125
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

type ApiDirectoryRolesValidatePropertiesRequest struct {
	ctx _context.Context
	ApiService *DirectoryRolesActionsApiService
	inlineObject126 *InlineObject126
}

func (r ApiDirectoryRolesValidatePropertiesRequest) InlineObject126(inlineObject126 InlineObject126) ApiDirectoryRolesValidatePropertiesRequest {
	r.inlineObject126 = &inlineObject126
	return r
}

func (r ApiDirectoryRolesValidatePropertiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DirectoryRolesValidatePropertiesExecute(r)
}

/*
DirectoryRolesValidateProperties Invoke action validateProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRolesValidatePropertiesRequest
*/
func (a *DirectoryRolesActionsApiService) DirectoryRolesValidateProperties(ctx _context.Context) ApiDirectoryRolesValidatePropertiesRequest {
	return ApiDirectoryRolesValidatePropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DirectoryRolesActionsApiService) DirectoryRolesValidatePropertiesExecute(r ApiDirectoryRolesValidatePropertiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRolesActionsApiService.DirectoryRolesValidateProperties")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoles/microsoft.graph.validateProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject126 == nil {
		return nil, reportError("inlineObject126 is required and must be specified")
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
	localVarPostBody = r.inlineObject126
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
