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

// PermissionGrantsActionsApiService PermissionGrantsActionsApi service
type PermissionGrantsActionsApiService service

type ApiPermissionGrantsGetAvailableExtensionPropertiesRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	inlineObject710 *InlineObject710
}

func (r ApiPermissionGrantsGetAvailableExtensionPropertiesRequest) InlineObject710(inlineObject710 InlineObject710) ApiPermissionGrantsGetAvailableExtensionPropertiesRequest {
	r.inlineObject710 = &inlineObject710
	return r
}

func (r ApiPermissionGrantsGetAvailableExtensionPropertiesRequest) Execute() ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsGetAvailableExtensionPropertiesExecute(r)
}

/*
PermissionGrantsGetAvailableExtensionProperties Invoke action getAvailableExtensionProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionGrantsGetAvailableExtensionPropertiesRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsGetAvailableExtensionProperties(ctx _context.Context) ApiPermissionGrantsGetAvailableExtensionPropertiesRequest {
	return ApiPermissionGrantsGetAvailableExtensionPropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphExtensionProperty
func (a *PermissionGrantsActionsApiService) PermissionGrantsGetAvailableExtensionPropertiesExecute(r ApiPermissionGrantsGetAvailableExtensionPropertiesRequest) ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphExtensionProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsGetAvailableExtensionProperties")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/microsoft.graph.getAvailableExtensionProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject710 == nil {
		return localVarReturnValue, nil, reportError("inlineObject710 is required and must be specified")
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
	localVarPostBody = r.inlineObject710
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

type ApiPermissionGrantsGetByIdsRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	inlineObject711 *InlineObject711
}

func (r ApiPermissionGrantsGetByIdsRequest) InlineObject711(inlineObject711 InlineObject711) ApiPermissionGrantsGetByIdsRequest {
	r.inlineObject711 = &inlineObject711
	return r
}

func (r ApiPermissionGrantsGetByIdsRequest) Execute() ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsGetByIdsExecute(r)
}

/*
PermissionGrantsGetByIds Invoke action getByIds

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionGrantsGetByIdsRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsGetByIds(ctx _context.Context) ApiPermissionGrantsGetByIdsRequest {
	return ApiPermissionGrantsGetByIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphDirectoryObject
func (a *PermissionGrantsActionsApiService) PermissionGrantsGetByIdsExecute(r ApiPermissionGrantsGetByIdsRequest) ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsGetByIds")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/microsoft.graph.getByIds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject711 == nil {
		return localVarReturnValue, nil, reportError("inlineObject711 is required and must be specified")
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
	localVarPostBody = r.inlineObject711
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

type ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	resourceSpecificPermissionGrantId string
	inlineObject706 *InlineObject706
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest) InlineObject706(inlineObject706 InlineObject706) ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest {
	r.inlineObject706 = &inlineObject706
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups Invoke action checkMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
//  @return []string
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.checkMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSpecificPermissionGrant-id"+"}", _neturl.PathEscape(parameterToString(r.resourceSpecificPermissionGrantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject706 == nil {
		return localVarReturnValue, nil, reportError("inlineObject706 is required and must be specified")
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
	localVarPostBody = r.inlineObject706
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

type ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	resourceSpecificPermissionGrantId string
	inlineObject707 *InlineObject707
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest) InlineObject707(inlineObject707 InlineObject707) ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest {
	r.inlineObject707 = &inlineObject707
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects Invoke action checkMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
//  @return []string
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.checkMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSpecificPermissionGrant-id"+"}", _neturl.PathEscape(parameterToString(r.resourceSpecificPermissionGrantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject707 == nil {
		return localVarReturnValue, nil, reportError("inlineObject707 is required and must be specified")
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
	localVarPostBody = r.inlineObject707
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

type ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	resourceSpecificPermissionGrantId string
	inlineObject708 *InlineObject708
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest) InlineObject708(inlineObject708 InlineObject708) ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest {
	r.inlineObject708 = &inlineObject708
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups Invoke action getMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
//  @return []string
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.getMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSpecificPermissionGrant-id"+"}", _neturl.PathEscape(parameterToString(r.resourceSpecificPermissionGrantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject708 == nil {
		return localVarReturnValue, nil, reportError("inlineObject708 is required and must be specified")
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
	localVarPostBody = r.inlineObject708
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

type ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	resourceSpecificPermissionGrantId string
	inlineObject709 *InlineObject709
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest) InlineObject709(inlineObject709 InlineObject709) ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest {
	r.inlineObject709 = &inlineObject709
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects Invoke action getMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
//  @return []string
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.getMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSpecificPermissionGrant-id"+"}", _neturl.PathEscape(parameterToString(r.resourceSpecificPermissionGrantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject709 == nil {
		return localVarReturnValue, nil, reportError("inlineObject709 is required and must be specified")
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
	localVarPostBody = r.inlineObject709
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

type ApiPermissionGrantsResourceSpecificPermissionGrantRestoreRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	resourceSpecificPermissionGrantId string
}


func (r ApiPermissionGrantsResourceSpecificPermissionGrantRestoreRequest) Execute() (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantRestoreExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantRestore Invoke action restore

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantRestoreRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantRestore(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantRestoreRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantRestoreRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
//  @return AnyOfmicrosoftGraphDirectoryObject
func (a *PermissionGrantsActionsApiService) PermissionGrantsResourceSpecificPermissionGrantRestoreExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantRestoreRequest) (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnyOfmicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsResourceSpecificPermissionGrantRestore")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.restore"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSpecificPermissionGrant-id"+"}", _neturl.PathEscape(parameterToString(r.resourceSpecificPermissionGrantId, "")), -1)

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

type ApiPermissionGrantsValidatePropertiesRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsActionsApiService
	inlineObject712 *InlineObject712
}

func (r ApiPermissionGrantsValidatePropertiesRequest) InlineObject712(inlineObject712 InlineObject712) ApiPermissionGrantsValidatePropertiesRequest {
	r.inlineObject712 = &inlineObject712
	return r
}

func (r ApiPermissionGrantsValidatePropertiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsValidatePropertiesExecute(r)
}

/*
PermissionGrantsValidateProperties Invoke action validateProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionGrantsValidatePropertiesRequest
*/
func (a *PermissionGrantsActionsApiService) PermissionGrantsValidateProperties(ctx _context.Context) ApiPermissionGrantsValidatePropertiesRequest {
	return ApiPermissionGrantsValidatePropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PermissionGrantsActionsApiService) PermissionGrantsValidatePropertiesExecute(r ApiPermissionGrantsValidatePropertiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsActionsApiService.PermissionGrantsValidateProperties")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/microsoft.graph.validateProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject712 == nil {
		return nil, reportError("inlineObject712 is required and must be specified")
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
	localVarPostBody = r.inlineObject712
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
