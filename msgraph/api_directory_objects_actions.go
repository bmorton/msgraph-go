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

// DirectoryObjectsActionsApiService DirectoryObjectsActionsApi service
type DirectoryObjectsActionsApiService service

type ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	directoryObjectId string
	inlineObject113 *InlineObject113
}

func (r ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest) InlineObject113(inlineObject113 InlineObject113) ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest {
	r.inlineObject113 = &inlineObject113
	return r
}

func (r ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsDirectoryObjectCheckMemberGroupsExecute(r)
}

/*
DirectoryObjectsDirectoryObjectCheckMemberGroups Invoke action checkMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryObjectId key: id of directoryObject
 @return ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectCheckMemberGroups(ctx _context.Context, directoryObjectId string) ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest {
	return ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		directoryObjectId: directoryObjectId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectCheckMemberGroupsExecute(r ApiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsDirectoryObjectCheckMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/{directoryObject-id}/microsoft.graph.checkMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryObject-id"+"}", _neturl.PathEscape(parameterToString(r.directoryObjectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject113 == nil {
		return localVarReturnValue, nil, reportError("inlineObject113 is required and must be specified")
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
	localVarPostBody = r.inlineObject113
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

type ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	directoryObjectId string
	inlineObject114 *InlineObject114
}

func (r ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest) InlineObject114(inlineObject114 InlineObject114) ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest {
	r.inlineObject114 = &inlineObject114
	return r
}

func (r ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsDirectoryObjectCheckMemberObjectsExecute(r)
}

/*
DirectoryObjectsDirectoryObjectCheckMemberObjects Invoke action checkMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryObjectId key: id of directoryObject
 @return ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectCheckMemberObjects(ctx _context.Context, directoryObjectId string) ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest {
	return ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		directoryObjectId: directoryObjectId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectCheckMemberObjectsExecute(r ApiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsDirectoryObjectCheckMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/{directoryObject-id}/microsoft.graph.checkMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryObject-id"+"}", _neturl.PathEscape(parameterToString(r.directoryObjectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject114 == nil {
		return localVarReturnValue, nil, reportError("inlineObject114 is required and must be specified")
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
	localVarPostBody = r.inlineObject114
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

type ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	directoryObjectId string
	inlineObject115 *InlineObject115
}

func (r ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest) InlineObject115(inlineObject115 InlineObject115) ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest {
	r.inlineObject115 = &inlineObject115
	return r
}

func (r ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsDirectoryObjectGetMemberGroupsExecute(r)
}

/*
DirectoryObjectsDirectoryObjectGetMemberGroups Invoke action getMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryObjectId key: id of directoryObject
 @return ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectGetMemberGroups(ctx _context.Context, directoryObjectId string) ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest {
	return ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		directoryObjectId: directoryObjectId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectGetMemberGroupsExecute(r ApiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsDirectoryObjectGetMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/{directoryObject-id}/microsoft.graph.getMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryObject-id"+"}", _neturl.PathEscape(parameterToString(r.directoryObjectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject115 == nil {
		return localVarReturnValue, nil, reportError("inlineObject115 is required and must be specified")
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
	localVarPostBody = r.inlineObject115
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

type ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	directoryObjectId string
	inlineObject116 *InlineObject116
}

func (r ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest) InlineObject116(inlineObject116 InlineObject116) ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest {
	r.inlineObject116 = &inlineObject116
	return r
}

func (r ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsDirectoryObjectGetMemberObjectsExecute(r)
}

/*
DirectoryObjectsDirectoryObjectGetMemberObjects Invoke action getMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryObjectId key: id of directoryObject
 @return ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectGetMemberObjects(ctx _context.Context, directoryObjectId string) ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest {
	return ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		directoryObjectId: directoryObjectId,
	}
}

// Execute executes the request
//  @return []string
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectGetMemberObjectsExecute(r ApiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsDirectoryObjectGetMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/{directoryObject-id}/microsoft.graph.getMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryObject-id"+"}", _neturl.PathEscape(parameterToString(r.directoryObjectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject116 == nil {
		return localVarReturnValue, nil, reportError("inlineObject116 is required and must be specified")
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
	localVarPostBody = r.inlineObject116
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

type ApiDirectoryObjectsDirectoryObjectRestoreRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	directoryObjectId string
}


func (r ApiDirectoryObjectsDirectoryObjectRestoreRequest) Execute() (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsDirectoryObjectRestoreExecute(r)
}

/*
DirectoryObjectsDirectoryObjectRestore Invoke action restore

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param directoryObjectId key: id of directoryObject
 @return ApiDirectoryObjectsDirectoryObjectRestoreRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectRestore(ctx _context.Context, directoryObjectId string) ApiDirectoryObjectsDirectoryObjectRestoreRequest {
	return ApiDirectoryObjectsDirectoryObjectRestoreRequest{
		ApiService: a,
		ctx: ctx,
		directoryObjectId: directoryObjectId,
	}
}

// Execute executes the request
//  @return AnyOfmicrosoftGraphDirectoryObject
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsDirectoryObjectRestoreExecute(r ApiDirectoryObjectsDirectoryObjectRestoreRequest) (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnyOfmicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsDirectoryObjectRestore")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/{directoryObject-id}/microsoft.graph.restore"
	localVarPath = strings.Replace(localVarPath, "{"+"directoryObject-id"+"}", _neturl.PathEscape(parameterToString(r.directoryObjectId, "")), -1)

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

type ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	inlineObject117 *InlineObject117
}

func (r ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest) InlineObject117(inlineObject117 InlineObject117) ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest {
	r.inlineObject117 = &inlineObject117
	return r
}

func (r ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest) Execute() ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsGetAvailableExtensionPropertiesExecute(r)
}

/*
DirectoryObjectsGetAvailableExtensionProperties Invoke action getAvailableExtensionProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsGetAvailableExtensionProperties(ctx _context.Context) ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest {
	return ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphExtensionProperty
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsGetAvailableExtensionPropertiesExecute(r ApiDirectoryObjectsGetAvailableExtensionPropertiesRequest) ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphExtensionProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsGetAvailableExtensionProperties")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/microsoft.graph.getAvailableExtensionProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject117 == nil {
		return localVarReturnValue, nil, reportError("inlineObject117 is required and must be specified")
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
	localVarPostBody = r.inlineObject117
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

type ApiDirectoryObjectsGetByIdsRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	inlineObject118 *InlineObject118
}

func (r ApiDirectoryObjectsGetByIdsRequest) InlineObject118(inlineObject118 InlineObject118) ApiDirectoryObjectsGetByIdsRequest {
	r.inlineObject118 = &inlineObject118
	return r
}

func (r ApiDirectoryObjectsGetByIdsRequest) Execute() ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsGetByIdsExecute(r)
}

/*
DirectoryObjectsGetByIds Invoke action getByIds

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryObjectsGetByIdsRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsGetByIds(ctx _context.Context) ApiDirectoryObjectsGetByIdsRequest {
	return ApiDirectoryObjectsGetByIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphDirectoryObject
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsGetByIdsExecute(r ApiDirectoryObjectsGetByIdsRequest) ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsGetByIds")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/microsoft.graph.getByIds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject118 == nil {
		return localVarReturnValue, nil, reportError("inlineObject118 is required and must be specified")
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
	localVarPostBody = r.inlineObject118
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

type ApiDirectoryObjectsValidatePropertiesRequest struct {
	ctx _context.Context
	ApiService *DirectoryObjectsActionsApiService
	inlineObject119 *InlineObject119
}

func (r ApiDirectoryObjectsValidatePropertiesRequest) InlineObject119(inlineObject119 InlineObject119) ApiDirectoryObjectsValidatePropertiesRequest {
	r.inlineObject119 = &inlineObject119
	return r
}

func (r ApiDirectoryObjectsValidatePropertiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DirectoryObjectsValidatePropertiesExecute(r)
}

/*
DirectoryObjectsValidateProperties Invoke action validateProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryObjectsValidatePropertiesRequest
*/
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsValidateProperties(ctx _context.Context) ApiDirectoryObjectsValidatePropertiesRequest {
	return ApiDirectoryObjectsValidatePropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DirectoryObjectsActionsApiService) DirectoryObjectsValidatePropertiesExecute(r ApiDirectoryObjectsValidatePropertiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryObjectsActionsApiService.DirectoryObjectsValidateProperties")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directoryObjects/microsoft.graph.validateProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject119 == nil {
		return nil, reportError("inlineObject119 is required and must be specified")
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
	localVarPostBody = r.inlineObject119
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
