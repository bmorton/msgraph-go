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

// ExternalExternalConnectionApiService ExternalExternalConnectionApi service
type ExternalExternalConnectionApiService service

type ApiExternalCreateConnectionsRequest struct {
	ctx _context.Context
	ApiService *ExternalExternalConnectionApiService
	microsoftGraphExternalConnectorsExternalConnection *MicrosoftGraphExternalConnectorsExternalConnection
}

// New navigation property
func (r ApiExternalCreateConnectionsRequest) MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection MicrosoftGraphExternalConnectorsExternalConnection) ApiExternalCreateConnectionsRequest {
	r.microsoftGraphExternalConnectorsExternalConnection = &microsoftGraphExternalConnectorsExternalConnection
	return r
}

func (r ApiExternalCreateConnectionsRequest) Execute() (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	return r.ApiService.ExternalCreateConnectionsExecute(r)
}

/*
ExternalCreateConnections Create new navigation property to connections for external

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalCreateConnectionsRequest
*/
func (a *ExternalExternalConnectionApiService) ExternalCreateConnections(ctx _context.Context) ApiExternalCreateConnectionsRequest {
	return ApiExternalCreateConnectionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExternalConnectorsExternalConnection
func (a *ExternalExternalConnectionApiService) ExternalCreateConnectionsExecute(r ApiExternalCreateConnectionsRequest) (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExternalConnectorsExternalConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalExternalConnectionApiService.ExternalCreateConnections")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/connections"

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

type ApiExternalDeleteConnectionsRequest struct {
	ctx _context.Context
	ApiService *ExternalExternalConnectionApiService
	externalConnectionId string
	ifMatch *string
}

// ETag
func (r ApiExternalDeleteConnectionsRequest) IfMatch(ifMatch string) ApiExternalDeleteConnectionsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiExternalDeleteConnectionsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ExternalDeleteConnectionsExecute(r)
}

/*
ExternalDeleteConnections Delete navigation property connections for external

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiExternalDeleteConnectionsRequest
*/
func (a *ExternalExternalConnectionApiService) ExternalDeleteConnections(ctx _context.Context, externalConnectionId string) ApiExternalDeleteConnectionsRequest {
	return ApiExternalDeleteConnectionsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
func (a *ExternalExternalConnectionApiService) ExternalDeleteConnectionsExecute(r ApiExternalDeleteConnectionsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalExternalConnectionApiService.ExternalDeleteConnections")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/connections/{externalConnection-id}"
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

type ApiExternalGetConnectionsRequest struct {
	ctx _context.Context
	ApiService *ExternalExternalConnectionApiService
	externalConnectionId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiExternalGetConnectionsRequest) Select_(select_ []string) ApiExternalGetConnectionsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiExternalGetConnectionsRequest) Expand(expand []string) ApiExternalGetConnectionsRequest {
	r.expand = &expand
	return r
}

func (r ApiExternalGetConnectionsRequest) Execute() (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	return r.ApiService.ExternalGetConnectionsExecute(r)
}

/*
ExternalGetConnections Get connections from external

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiExternalGetConnectionsRequest
*/
func (a *ExternalExternalConnectionApiService) ExternalGetConnections(ctx _context.Context, externalConnectionId string) ApiExternalGetConnectionsRequest {
	return ApiExternalGetConnectionsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExternalConnectorsExternalConnection
func (a *ExternalExternalConnectionApiService) ExternalGetConnectionsExecute(r ApiExternalGetConnectionsRequest) (MicrosoftGraphExternalConnectorsExternalConnection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExternalConnectorsExternalConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalExternalConnectionApiService.ExternalGetConnections")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/connections/{externalConnection-id}"
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

type ApiExternalListConnectionsRequest struct {
	ctx _context.Context
	ApiService *ExternalExternalConnectionApiService
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
func (r ApiExternalListConnectionsRequest) Top(top int32) ApiExternalListConnectionsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiExternalListConnectionsRequest) Skip(skip int32) ApiExternalListConnectionsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiExternalListConnectionsRequest) Search(search string) ApiExternalListConnectionsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiExternalListConnectionsRequest) Filter(filter string) ApiExternalListConnectionsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiExternalListConnectionsRequest) Count(count bool) ApiExternalListConnectionsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiExternalListConnectionsRequest) Orderby(orderby []string) ApiExternalListConnectionsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiExternalListConnectionsRequest) Select_(select_ []string) ApiExternalListConnectionsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiExternalListConnectionsRequest) Expand(expand []string) ApiExternalListConnectionsRequest {
	r.expand = &expand
	return r
}

func (r ApiExternalListConnectionsRequest) Execute() (CollectionOfExternalConnection, *_nethttp.Response, error) {
	return r.ApiService.ExternalListConnectionsExecute(r)
}

/*
ExternalListConnections Get connections from external

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalListConnectionsRequest
*/
func (a *ExternalExternalConnectionApiService) ExternalListConnections(ctx _context.Context) ApiExternalListConnectionsRequest {
	return ApiExternalListConnectionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfExternalConnection
func (a *ExternalExternalConnectionApiService) ExternalListConnectionsExecute(r ApiExternalListConnectionsRequest) (CollectionOfExternalConnection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfExternalConnection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalExternalConnectionApiService.ExternalListConnections")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/connections"

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

type ApiExternalUpdateConnectionsRequest struct {
	ctx _context.Context
	ApiService *ExternalExternalConnectionApiService
	externalConnectionId string
	microsoftGraphExternalConnectorsExternalConnection *MicrosoftGraphExternalConnectorsExternalConnection
}

// New navigation property values
func (r ApiExternalUpdateConnectionsRequest) MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection MicrosoftGraphExternalConnectorsExternalConnection) ApiExternalUpdateConnectionsRequest {
	r.microsoftGraphExternalConnectorsExternalConnection = &microsoftGraphExternalConnectorsExternalConnection
	return r
}

func (r ApiExternalUpdateConnectionsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ExternalUpdateConnectionsExecute(r)
}

/*
ExternalUpdateConnections Update the navigation property connections in external

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiExternalUpdateConnectionsRequest
*/
func (a *ExternalExternalConnectionApiService) ExternalUpdateConnections(ctx _context.Context, externalConnectionId string) ApiExternalUpdateConnectionsRequest {
	return ApiExternalUpdateConnectionsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
func (a *ExternalExternalConnectionApiService) ExternalUpdateConnectionsExecute(r ApiExternalUpdateConnectionsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalExternalConnectionApiService.ExternalUpdateConnections")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/connections/{externalConnection-id}"
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