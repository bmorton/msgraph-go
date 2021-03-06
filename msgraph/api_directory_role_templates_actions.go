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

// DirectoryRoleTemplatesActionsApiService DirectoryRoleTemplatesActionsApi service
type DirectoryRoleTemplatesActionsApiService service

type ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	directoryRoleTemplateId string
	inlineObject127 *InlineObject127
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest) InlineObject127(inlineObject127 InlineObject127) ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest {
	r.inlineObject127 = &inlineObject127
	return r
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsExecute(r)
}

/*
DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups Invoke action checkMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleTemplateId key: id of directoryRoleTemplate
 @return ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups(ctx _context.Context, directoryRoleTemplateId string) ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest {
	return ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleTemplateId: directoryRoleTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsExecute(r ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.checkMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRoleTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject127 == nil {
		return localVarReturnValue, nil, reportError("inlineObject127 is required and must be specified")
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
	localVarPostBody = r.inlineObject127
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

type ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	directoryRoleTemplateId string
	inlineObject128 *InlineObject128
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest) InlineObject128(inlineObject128 InlineObject128) ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest {
	r.inlineObject128 = &inlineObject128
	return r
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsExecute(r)
}

/*
DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects Invoke action checkMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleTemplateId key: id of directoryRoleTemplate
 @return ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects(ctx _context.Context, directoryRoleTemplateId string) ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest {
	return ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleTemplateId: directoryRoleTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsExecute(r ApiDirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesDirectoryRoleTemplateCheckMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.checkMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRoleTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject128 == nil {
		return localVarReturnValue, nil, reportError("inlineObject128 is required and must be specified")
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
	localVarPostBody = r.inlineObject128
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

type ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	directoryRoleTemplateId string
	inlineObject129 *InlineObject129
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest) InlineObject129(inlineObject129 InlineObject129) ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest {
	r.inlineObject129 = &inlineObject129
	return r
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsExecute(r)
}

/*
DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups Invoke action getMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleTemplateId key: id of directoryRoleTemplate
 @return ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups(ctx _context.Context, directoryRoleTemplateId string) ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest {
	return ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleTemplateId: directoryRoleTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsExecute(r ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.getMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRoleTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject129 == nil {
		return localVarReturnValue, nil, reportError("inlineObject129 is required and must be specified")
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
	localVarPostBody = r.inlineObject129
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

type ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	directoryRoleTemplateId string
	inlineObject130 *InlineObject130
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest) InlineObject130(inlineObject130 InlineObject130) ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest {
	r.inlineObject130 = &inlineObject130
	return r
}

func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsExecute(r)
}

/*
DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects Invoke action getMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleTemplateId key: id of directoryRoleTemplate
 @return ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects(ctx _context.Context, directoryRoleTemplateId string) ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest {
	return ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleTemplateId: directoryRoleTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsExecute(r ApiDirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesDirectoryRoleTemplateGetMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.getMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRoleTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject130 == nil {
		return localVarReturnValue, nil, reportError("inlineObject130 is required and must be specified")
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
	localVarPostBody = r.inlineObject130
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

type ApiDirectoryRoleTemplatesDirectoryRoleTemplateRestoreRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	directoryRoleTemplateId string
}


func (r ApiDirectoryRoleTemplatesDirectoryRoleTemplateRestoreRequest) Execute() (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesDirectoryRoleTemplateRestoreExecute(r)
}

/*
DirectoryRoleTemplatesDirectoryRoleTemplateRestore Invoke action restore

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryRoleTemplateId key: id of directoryRoleTemplate
 @return ApiDirectoryRoleTemplatesDirectoryRoleTemplateRestoreRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateRestore(ctx _context.Context, directoryRoleTemplateId string) ApiDirectoryRoleTemplatesDirectoryRoleTemplateRestoreRequest {
	return ApiDirectoryRoleTemplatesDirectoryRoleTemplateRestoreRequest{
		ApiService: a,
		ctx: ctx,
		directoryRoleTemplateId: directoryRoleTemplateId,
	}
}

// Execute executes the request
//  @return AnyOfmicrosoftGraphDirectoryObject
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesDirectoryRoleTemplateRestoreExecute(r ApiDirectoryRoleTemplatesDirectoryRoleTemplateRestoreRequest) (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnyOfmicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesDirectoryRoleTemplateRestore")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/{directoryRoleTemplate-id}/microsoft.graph.restore"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryRoleTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.directoryRoleTemplateId, "")), -1)

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

type ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	inlineObject131 *InlineObject131
}

func (r ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest) InlineObject131(inlineObject131 InlineObject131) ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest {
	r.inlineObject131 = &inlineObject131
	return r
}

func (r ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest) Execute() ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesGetAvailableExtensionPropertiesExecute(r)
}

/*
DirectoryRoleTemplatesGetAvailableExtensionProperties Invoke action getAvailableExtensionProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesGetAvailableExtensionProperties(ctx _context.Context) ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest {
	return ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphExtensionProperty
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesGetAvailableExtensionPropertiesExecute(r ApiDirectoryRoleTemplatesGetAvailableExtensionPropertiesRequest) ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphExtensionProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesGetAvailableExtensionProperties")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/microsoft.graph.getAvailableExtensionProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject131 == nil {
		return localVarReturnValue, nil, reportError("inlineObject131 is required and must be specified")
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
	localVarPostBody = r.inlineObject131
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

type ApiDirectoryRoleTemplatesGetByIdsRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	inlineObject132 *InlineObject132
}

func (r ApiDirectoryRoleTemplatesGetByIdsRequest) InlineObject132(inlineObject132 InlineObject132) ApiDirectoryRoleTemplatesGetByIdsRequest {
	r.inlineObject132 = &inlineObject132
	return r
}

func (r ApiDirectoryRoleTemplatesGetByIdsRequest) Execute() ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesGetByIdsExecute(r)
}

/*
DirectoryRoleTemplatesGetByIds Invoke action getByIds

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRoleTemplatesGetByIdsRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesGetByIds(ctx _context.Context) ApiDirectoryRoleTemplatesGetByIdsRequest {
	return ApiDirectoryRoleTemplatesGetByIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphDirectoryObject
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesGetByIdsExecute(r ApiDirectoryRoleTemplatesGetByIdsRequest) ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesGetByIds")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/microsoft.graph.getByIds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject132 == nil {
		return localVarReturnValue, nil, reportError("inlineObject132 is required and must be specified")
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
	localVarPostBody = r.inlineObject132
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

type ApiDirectoryRoleTemplatesValidatePropertiesRequest struct {
	ctx _context.Context
	ApiService *DirectoryRoleTemplatesActionsApiService
	inlineObject133 *InlineObject133
}

func (r ApiDirectoryRoleTemplatesValidatePropertiesRequest) InlineObject133(inlineObject133 InlineObject133) ApiDirectoryRoleTemplatesValidatePropertiesRequest {
	r.inlineObject133 = &inlineObject133
	return r
}

func (r ApiDirectoryRoleTemplatesValidatePropertiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DirectoryRoleTemplatesValidatePropertiesExecute(r)
}

/*
DirectoryRoleTemplatesValidateProperties Invoke action validateProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryRoleTemplatesValidatePropertiesRequest
*/
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesValidateProperties(ctx _context.Context) ApiDirectoryRoleTemplatesValidatePropertiesRequest {
	return ApiDirectoryRoleTemplatesValidatePropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DirectoryRoleTemplatesActionsApiService) DirectoryRoleTemplatesValidatePropertiesExecute(r ApiDirectoryRoleTemplatesValidatePropertiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryRoleTemplatesActionsApiService.DirectoryRoleTemplatesValidateProperties")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryRoleTemplates/microsoft.graph.validateProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject133 == nil {
		return nil, reportError("inlineObject133 is required and must be specified")
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
	localVarPostBody = r.inlineObject133
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
