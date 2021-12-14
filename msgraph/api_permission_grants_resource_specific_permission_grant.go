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

// PermissionGrantsResourceSpecificPermissionGrantApiService PermissionGrantsResourceSpecificPermissionGrantApi service
type PermissionGrantsResourceSpecificPermissionGrantApiService service

type ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsResourceSpecificPermissionGrantApiService
	microsoftGraphResourceSpecificPermissionGrant *MicrosoftGraphResourceSpecificPermissionGrant
}

// New entity
func (r ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest) MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant MicrosoftGraphResourceSpecificPermissionGrant) ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest {
	r.microsoftGraphResourceSpecificPermissionGrant = &microsoftGraphResourceSpecificPermissionGrant
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest) Execute() (MicrosoftGraphResourceSpecificPermissionGrant, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant Add new entity to permissionGrants

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest
*/
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant(ctx _context.Context) ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphResourceSpecificPermissionGrant
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrantRequest) (MicrosoftGraphResourceSpecificPermissionGrant, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphResourceSpecificPermissionGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsResourceSpecificPermissionGrantApiService.PermissionGrantsResourceSpecificPermissionGrantCreateResourceSpecificPermissionGrant")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphResourceSpecificPermissionGrant == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphResourceSpecificPermissionGrant is required and must be specified")
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
	localVarPostBody = r.microsoftGraphResourceSpecificPermissionGrant
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

type ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsResourceSpecificPermissionGrantApiService
	resourceSpecificPermissionGrantId string
	ifMatch *string
}

// ETag
func (r ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest) IfMatch(ifMatch string) ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant Delete entity from permissionGrants

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest
*/
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrantRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsResourceSpecificPermissionGrantApiService.PermissionGrantsResourceSpecificPermissionGrantDeleteResourceSpecificPermissionGrant")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}"
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

type ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsResourceSpecificPermissionGrantApiService
	resourceSpecificPermissionGrantId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest) Select_(select_ []string) ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest) Expand(expand []string) ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest {
	r.expand = &expand
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest) Execute() (MicrosoftGraphResourceSpecificPermissionGrant, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant Get entity from permissionGrants by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest
*/
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphResourceSpecificPermissionGrant
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrantRequest) (MicrosoftGraphResourceSpecificPermissionGrant, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphResourceSpecificPermissionGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsResourceSpecificPermissionGrantApiService.PermissionGrantsResourceSpecificPermissionGrantGetResourceSpecificPermissionGrant")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSpecificPermissionGrant-id"+"}", _neturl.PathEscape(parameterToString(r.resourceSpecificPermissionGrantId, "")), -1)

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

type ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsResourceSpecificPermissionGrantApiService
	search *string
	filter *string
	orderby *[]string
	select_ *[]string
	expand *[]string
}

// Search items by search phrases
func (r ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest) Search(search string) ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest) Filter(filter string) ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest {
	r.filter = &filter
	return r
}
// Order items by property values
func (r ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest) Orderby(orderby []string) ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest) Select_(select_ []string) ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest) Expand(expand []string) ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest {
	r.expand = &expand
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest) Execute() (CollectionOfResourceSpecificPermissionGrant, *_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant Get entities from permissionGrants

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest
*/
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant(ctx _context.Context) ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfResourceSpecificPermissionGrant
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrantRequest) (CollectionOfResourceSpecificPermissionGrant, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfResourceSpecificPermissionGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsResourceSpecificPermissionGrantApiService.PermissionGrantsResourceSpecificPermissionGrantListResourceSpecificPermissionGrant")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("$search", parameterToString(*r.search, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("$filter", parameterToString(*r.filter, ""))
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

type ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest struct {
	ctx _context.Context
	ApiService *PermissionGrantsResourceSpecificPermissionGrantApiService
	resourceSpecificPermissionGrantId string
	microsoftGraphResourceSpecificPermissionGrant *MicrosoftGraphResourceSpecificPermissionGrant
}

// New property values
func (r ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest) MicrosoftGraphResourceSpecificPermissionGrant(microsoftGraphResourceSpecificPermissionGrant MicrosoftGraphResourceSpecificPermissionGrant) ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest {
	r.microsoftGraphResourceSpecificPermissionGrant = &microsoftGraphResourceSpecificPermissionGrant
	return r
}

func (r ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantExecute(r)
}

/*
PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant Update entity in permissionGrants

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceSpecificPermissionGrantId key: id of resourceSpecificPermissionGrant
 @return ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest
*/
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant(ctx _context.Context, resourceSpecificPermissionGrantId string) ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest {
	return ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest{
		ApiService: a,
		ctx: ctx,
		resourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// Execute executes the request
func (a *PermissionGrantsResourceSpecificPermissionGrantApiService) PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantExecute(r ApiPermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrantRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionGrantsResourceSpecificPermissionGrantApiService.PermissionGrantsResourceSpecificPermissionGrantUpdateResourceSpecificPermissionGrant")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissionGrants/{resourceSpecificPermissionGrant-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceSpecificPermissionGrant-id"+"}", _neturl.PathEscape(parameterToString(r.resourceSpecificPermissionGrantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphResourceSpecificPermissionGrant == nil {
		return nil, reportError("microsoftGraphResourceSpecificPermissionGrant is required and must be specified")
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
	localVarPostBody = r.microsoftGraphResourceSpecificPermissionGrant
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
