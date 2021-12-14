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

// ConnectionsExternalConnectionApiService ConnectionsExternalConnectionApi service
type ConnectionsExternalConnectionApiService service

type ApiConnectionsExternalConnectionCreateExternalConnectionRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalConnectionApiService
	microsoftGraphExternalConnectorsExternalConnection *MicrosoftGraphExternalConnectorsExternalConnection
}

// New entity
func (r ApiConnectionsExternalConnectionCreateExternalConnectionRequest) MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection MicrosoftGraphExternalConnectorsExternalConnection) ApiConnectionsExternalConnectionCreateExternalConnectionRequest {
	r.microsoftGraphExternalConnectorsExternalConnection = &microsoftGraphExternalConnectorsExternalConnection
	return r
}

func (r ApiConnectionsExternalConnectionCreateExternalConnectionRequest) Execute() (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	return r.ApiService.ConnectionsExternalConnectionCreateExternalConnectionExecute(r)
}

/*
ConnectionsExternalConnectionCreateExternalConnection Add new entity to connections

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConnectionsExternalConnectionCreateExternalConnectionRequest
*/
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionCreateExternalConnection(ctx _context.Context) ApiConnectionsExternalConnectionCreateExternalConnectionRequest {
	return ApiConnectionsExternalConnectionCreateExternalConnectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExternalConnectorsExternalConnection
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionCreateExternalConnectionExecute(r ApiConnectionsExternalConnectionCreateExternalConnectionRequest) (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExternalConnectorsExternalConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalConnectionApiService.ConnectionsExternalConnectionCreateExternalConnection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphExternalConnectorsExternalConnection == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphExternalConnectorsExternalConnection is required and must be specified")
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
	localVarPostBody = r.microsoftGraphExternalConnectorsExternalConnection
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

type ApiConnectionsExternalConnectionDeleteExternalConnectionRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalConnectionApiService
	externalConnectionId string
	ifMatch *string
}

// ETag
func (r ApiConnectionsExternalConnectionDeleteExternalConnectionRequest) IfMatch(ifMatch string) ApiConnectionsExternalConnectionDeleteExternalConnectionRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiConnectionsExternalConnectionDeleteExternalConnectionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectionsExternalConnectionDeleteExternalConnectionExecute(r)
}

/*
ConnectionsExternalConnectionDeleteExternalConnection Delete entity from connections

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsExternalConnectionDeleteExternalConnectionRequest
*/
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionDeleteExternalConnection(ctx _context.Context, externalConnectionId string) ApiConnectionsExternalConnectionDeleteExternalConnectionRequest {
	return ApiConnectionsExternalConnectionDeleteExternalConnectionRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionDeleteExternalConnectionExecute(r ApiConnectionsExternalConnectionDeleteExternalConnectionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalConnectionApiService.ConnectionsExternalConnectionDeleteExternalConnection")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)

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

type ApiConnectionsExternalConnectionGetExternalConnectionRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalConnectionApiService
	externalConnectionId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiConnectionsExternalConnectionGetExternalConnectionRequest) Select_(select_ []string) ApiConnectionsExternalConnectionGetExternalConnectionRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiConnectionsExternalConnectionGetExternalConnectionRequest) Expand(expand []string) ApiConnectionsExternalConnectionGetExternalConnectionRequest {
	r.expand = &expand
	return r
}

func (r ApiConnectionsExternalConnectionGetExternalConnectionRequest) Execute() (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	return r.ApiService.ConnectionsExternalConnectionGetExternalConnectionExecute(r)
}

/*
ConnectionsExternalConnectionGetExternalConnection Get entity from connections by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsExternalConnectionGetExternalConnectionRequest
*/
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionGetExternalConnection(ctx _context.Context, externalConnectionId string) ApiConnectionsExternalConnectionGetExternalConnectionRequest {
	return ApiConnectionsExternalConnectionGetExternalConnectionRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExternalConnectorsExternalConnection
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionGetExternalConnectionExecute(r ApiConnectionsExternalConnectionGetExternalConnectionRequest) (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExternalConnectorsExternalConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalConnectionApiService.ConnectionsExternalConnectionGetExternalConnection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)

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

type ApiConnectionsExternalConnectionListExternalConnectionRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalConnectionApiService
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
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Top(top int32) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Skip(skip int32) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Search(search string) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Filter(filter string) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Count(count bool) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Orderby(orderby []string) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Select_(select_ []string) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Expand(expand []string) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	r.expand = &expand
	return r
}

func (r ApiConnectionsExternalConnectionListExternalConnectionRequest) Execute() (CollectionOfExternalConnection, *_nethttp.Response, error) {
	return r.ApiService.ConnectionsExternalConnectionListExternalConnectionExecute(r)
}

/*
ConnectionsExternalConnectionListExternalConnection Get entities from connections

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConnectionsExternalConnectionListExternalConnectionRequest
*/
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionListExternalConnection(ctx _context.Context) ApiConnectionsExternalConnectionListExternalConnectionRequest {
	return ApiConnectionsExternalConnectionListExternalConnectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfExternalConnection
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionListExternalConnectionExecute(r ApiConnectionsExternalConnectionListExternalConnectionRequest) (CollectionOfExternalConnection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfExternalConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalConnectionApiService.ConnectionsExternalConnectionListExternalConnection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections"

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

type ApiConnectionsExternalConnectionUpdateExternalConnectionRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalConnectionApiService
	externalConnectionId string
	microsoftGraphExternalConnectorsExternalConnection *MicrosoftGraphExternalConnectorsExternalConnection
}

// New property values
func (r ApiConnectionsExternalConnectionUpdateExternalConnectionRequest) MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection MicrosoftGraphExternalConnectorsExternalConnection) ApiConnectionsExternalConnectionUpdateExternalConnectionRequest {
	r.microsoftGraphExternalConnectorsExternalConnection = &microsoftGraphExternalConnectorsExternalConnection
	return r
}

func (r ApiConnectionsExternalConnectionUpdateExternalConnectionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectionsExternalConnectionUpdateExternalConnectionExecute(r)
}

/*
ConnectionsExternalConnectionUpdateExternalConnection Update entity in connections

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsExternalConnectionUpdateExternalConnectionRequest
*/
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionUpdateExternalConnection(ctx _context.Context, externalConnectionId string) ApiConnectionsExternalConnectionUpdateExternalConnectionRequest {
	return ApiConnectionsExternalConnectionUpdateExternalConnectionRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
func (a *ConnectionsExternalConnectionApiService) ConnectionsExternalConnectionUpdateExternalConnectionExecute(r ApiConnectionsExternalConnectionUpdateExternalConnectionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalConnectionApiService.ConnectionsExternalConnectionUpdateExternalConnection")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphExternalConnectorsExternalConnection == nil {
		return nil, reportError("microsoftGraphExternalConnectorsExternalConnection is required and must be specified")
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
	localVarPostBody = r.microsoftGraphExternalConnectorsExternalConnection
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
