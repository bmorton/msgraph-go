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

// GroupSettingTemplatesActionsApiService GroupSettingTemplatesActionsApi service
type GroupSettingTemplatesActionsApiService service

type ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	inlineObject353 *InlineObject353
}

func (r ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest) InlineObject353(inlineObject353 InlineObject353) ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest {
	r.inlineObject353 = &inlineObject353
	return r
}

func (r ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest) Execute() ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesGetAvailableExtensionPropertiesExecute(r)
}

/*
GroupSettingTemplatesGetAvailableExtensionProperties Invoke action getAvailableExtensionProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGetAvailableExtensionProperties(ctx _context.Context) ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest {
	return ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphExtensionProperty
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGetAvailableExtensionPropertiesExecute(r ApiGroupSettingTemplatesGetAvailableExtensionPropertiesRequest) ([]MicrosoftGraphExtensionProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphExtensionProperty
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesGetAvailableExtensionProperties")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/microsoft.graph.getAvailableExtensionProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject353 == nil {
		return localVarReturnValue, nil, reportError("inlineObject353 is required and must be specified")
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
	localVarPostBody = r.inlineObject353
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

type ApiGroupSettingTemplatesGetByIdsRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	inlineObject354 *InlineObject354
}

func (r ApiGroupSettingTemplatesGetByIdsRequest) InlineObject354(inlineObject354 InlineObject354) ApiGroupSettingTemplatesGetByIdsRequest {
	r.inlineObject354 = &inlineObject354
	return r
}

func (r ApiGroupSettingTemplatesGetByIdsRequest) Execute() ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesGetByIdsExecute(r)
}

/*
GroupSettingTemplatesGetByIds Invoke action getByIds

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGroupSettingTemplatesGetByIdsRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGetByIds(ctx _context.Context) ApiGroupSettingTemplatesGetByIdsRequest {
	return ApiGroupSettingTemplatesGetByIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MicrosoftGraphDirectoryObject
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGetByIdsExecute(r ApiGroupSettingTemplatesGetByIdsRequest) ([]MicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []MicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesGetByIds")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/microsoft.graph.getByIds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject354 == nil {
		return localVarReturnValue, nil, reportError("inlineObject354 is required and must be specified")
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
	localVarPostBody = r.inlineObject354
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

type ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	groupSettingTemplateId string
	inlineObject349 *InlineObject349
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest) InlineObject349(inlineObject349 InlineObject349) ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest {
	r.inlineObject349 = &inlineObject349
	return r
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsExecute(r)
}

/*
GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups Invoke action checkMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingTemplateId key: id of groupSettingTemplate
 @return ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups(ctx _context.Context, groupSettingTemplateId string) ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest {
	return ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingTemplateId: groupSettingTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsExecute(r ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesGroupSettingTemplateCheckMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.checkMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSettingTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject349 == nil {
		return localVarReturnValue, nil, reportError("inlineObject349 is required and must be specified")
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
	localVarPostBody = r.inlineObject349
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

type ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	groupSettingTemplateId string
	inlineObject350 *InlineObject350
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest) InlineObject350(inlineObject350 InlineObject350) ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest {
	r.inlineObject350 = &inlineObject350
	return r
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsExecute(r)
}

/*
GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects Invoke action checkMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingTemplateId key: id of groupSettingTemplate
 @return ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects(ctx _context.Context, groupSettingTemplateId string) ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest {
	return ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingTemplateId: groupSettingTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsExecute(r ApiGroupSettingTemplatesGroupSettingTemplateCheckMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesGroupSettingTemplateCheckMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.checkMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSettingTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject350 == nil {
		return localVarReturnValue, nil, reportError("inlineObject350 is required and must be specified")
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
	localVarPostBody = r.inlineObject350
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

type ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	groupSettingTemplateId string
	inlineObject351 *InlineObject351
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest) InlineObject351(inlineObject351 InlineObject351) ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest {
	r.inlineObject351 = &inlineObject351
	return r
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesGroupSettingTemplateGetMemberGroupsExecute(r)
}

/*
GroupSettingTemplatesGroupSettingTemplateGetMemberGroups Invoke action getMemberGroups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingTemplateId key: id of groupSettingTemplate
 @return ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateGetMemberGroups(ctx _context.Context, groupSettingTemplateId string) ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest {
	return ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingTemplateId: groupSettingTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateGetMemberGroupsExecute(r ApiGroupSettingTemplatesGroupSettingTemplateGetMemberGroupsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesGroupSettingTemplateGetMemberGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.getMemberGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSettingTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject351 == nil {
		return localVarReturnValue, nil, reportError("inlineObject351 is required and must be specified")
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
	localVarPostBody = r.inlineObject351
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

type ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	groupSettingTemplateId string
	inlineObject352 *InlineObject352
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest) InlineObject352(inlineObject352 InlineObject352) ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest {
	r.inlineObject352 = &inlineObject352
	return r
}

func (r ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesGroupSettingTemplateGetMemberObjectsExecute(r)
}

/*
GroupSettingTemplatesGroupSettingTemplateGetMemberObjects Invoke action getMemberObjects

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingTemplateId key: id of groupSettingTemplate
 @return ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateGetMemberObjects(ctx _context.Context, groupSettingTemplateId string) ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest {
	return ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingTemplateId: groupSettingTemplateId,
	}
}

// Execute executes the request
//  @return []string
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateGetMemberObjectsExecute(r ApiGroupSettingTemplatesGroupSettingTemplateGetMemberObjectsRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesGroupSettingTemplateGetMemberObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.getMemberObjects"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSettingTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject352 == nil {
		return localVarReturnValue, nil, reportError("inlineObject352 is required and must be specified")
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
	localVarPostBody = r.inlineObject352
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

type ApiGroupSettingTemplatesGroupSettingTemplateRestoreRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	groupSettingTemplateId string
}


func (r ApiGroupSettingTemplatesGroupSettingTemplateRestoreRequest) Execute() (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesGroupSettingTemplateRestoreExecute(r)
}

/*
GroupSettingTemplatesGroupSettingTemplateRestore Invoke action restore

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupSettingTemplateId key: id of groupSettingTemplate
 @return ApiGroupSettingTemplatesGroupSettingTemplateRestoreRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateRestore(ctx _context.Context, groupSettingTemplateId string) ApiGroupSettingTemplatesGroupSettingTemplateRestoreRequest {
	return ApiGroupSettingTemplatesGroupSettingTemplateRestoreRequest{
		ApiService: a,
		ctx: ctx,
		groupSettingTemplateId: groupSettingTemplateId,
	}
}

// Execute executes the request
//  @return AnyOfmicrosoftGraphDirectoryObject
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesGroupSettingTemplateRestoreExecute(r ApiGroupSettingTemplatesGroupSettingTemplateRestoreRequest) (AnyOfmicrosoftGraphDirectoryObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnyOfmicrosoftGraphDirectoryObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesGroupSettingTemplateRestore")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/{groupSettingTemplate-id}/microsoft.graph.restore"
	localVarPath = strings.Replace(localVarPath, "{"+"groupSettingTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.groupSettingTemplateId, "")), -1)

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

type ApiGroupSettingTemplatesValidatePropertiesRequest struct {
	ctx _context.Context
	ApiService *GroupSettingTemplatesActionsApiService
	inlineObject355 *InlineObject355
}

func (r ApiGroupSettingTemplatesValidatePropertiesRequest) InlineObject355(inlineObject355 InlineObject355) ApiGroupSettingTemplatesValidatePropertiesRequest {
	r.inlineObject355 = &inlineObject355
	return r
}

func (r ApiGroupSettingTemplatesValidatePropertiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GroupSettingTemplatesValidatePropertiesExecute(r)
}

/*
GroupSettingTemplatesValidateProperties Invoke action validateProperties

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGroupSettingTemplatesValidatePropertiesRequest
*/
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesValidateProperties(ctx _context.Context) ApiGroupSettingTemplatesValidatePropertiesRequest {
	return ApiGroupSettingTemplatesValidatePropertiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *GroupSettingTemplatesActionsApiService) GroupSettingTemplatesValidatePropertiesExecute(r ApiGroupSettingTemplatesValidatePropertiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSettingTemplatesActionsApiService.GroupSettingTemplatesValidateProperties")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groupSettingTemplates/microsoft.graph.validateProperties"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject355 == nil {
		return nil, reportError("inlineObject355 is required and must be specified")
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
	localVarPostBody = r.inlineObject355
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
