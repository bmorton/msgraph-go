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

// ConnectionsExternalItemApiService ConnectionsExternalItemApi service
type ConnectionsExternalItemApiService service

type ApiConnectionsCreateItemsRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalItemApiService
	externalConnectionId string
	microsoftGraphExternalConnectorsExternalItem *MicrosoftGraphExternalConnectorsExternalItem
}

// New navigation property
func (r ApiConnectionsCreateItemsRequest) MicrosoftGraphExternalConnectorsExternalItem(microsoftGraphExternalConnectorsExternalItem MicrosoftGraphExternalConnectorsExternalItem) ApiConnectionsCreateItemsRequest {
	r.microsoftGraphExternalConnectorsExternalItem = &microsoftGraphExternalConnectorsExternalItem
	return r
}

func (r ApiConnectionsCreateItemsRequest) Execute() (MicrosoftGraphExternalConnectorsExternalItem, *_nethttp.Response, error) {
	return r.ApiService.ConnectionsCreateItemsExecute(r)
}

/*
ConnectionsCreateItems Create new navigation property to items for connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsCreateItemsRequest
*/
func (a *ConnectionsExternalItemApiService) ConnectionsCreateItems(ctx _context.Context, externalConnectionId string) ApiConnectionsCreateItemsRequest {
	return ApiConnectionsCreateItemsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExternalConnectorsExternalItem
func (a *ConnectionsExternalItemApiService) ConnectionsCreateItemsExecute(r ApiConnectionsCreateItemsRequest) (MicrosoftGraphExternalConnectorsExternalItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExternalConnectorsExternalItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalItemApiService.ConnectionsCreateItems")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphExternalConnectorsExternalItem == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphExternalConnectorsExternalItem is required and must be specified")
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
	localVarPostBody = r.microsoftGraphExternalConnectorsExternalItem
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

type ApiConnectionsDeleteItemsRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalItemApiService
	externalConnectionId string
	externalItemId string
	ifMatch *string
}

// ETag
func (r ApiConnectionsDeleteItemsRequest) IfMatch(ifMatch string) ApiConnectionsDeleteItemsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiConnectionsDeleteItemsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectionsDeleteItemsExecute(r)
}

/*
ConnectionsDeleteItems Delete navigation property items for connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @param externalItemId key: id of externalItem
 @return ApiConnectionsDeleteItemsRequest
*/
func (a *ConnectionsExternalItemApiService) ConnectionsDeleteItems(ctx _context.Context, externalConnectionId string, externalItemId string) ApiConnectionsDeleteItemsRequest {
	return ApiConnectionsDeleteItemsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
		externalItemId: externalItemId,
	}
}

// Execute executes the request
func (a *ConnectionsExternalItemApiService) ConnectionsDeleteItemsExecute(r ApiConnectionsDeleteItemsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalItemApiService.ConnectionsDeleteItems")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/items/{externalItem-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalItem-id"+"}", _neturl.PathEscape(parameterToString(r.externalItemId, "")), -1)

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

type ApiConnectionsGetItemsRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalItemApiService
	externalConnectionId string
	externalItemId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiConnectionsGetItemsRequest) Select_(select_ []string) ApiConnectionsGetItemsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiConnectionsGetItemsRequest) Expand(expand []string) ApiConnectionsGetItemsRequest {
	r.expand = &expand
	return r
}

func (r ApiConnectionsGetItemsRequest) Execute() (MicrosoftGraphExternalConnectorsExternalItem, *_nethttp.Response, error) {
	return r.ApiService.ConnectionsGetItemsExecute(r)
}

/*
ConnectionsGetItems Get items from connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @param externalItemId key: id of externalItem
 @return ApiConnectionsGetItemsRequest
*/
func (a *ConnectionsExternalItemApiService) ConnectionsGetItems(ctx _context.Context, externalConnectionId string, externalItemId string) ApiConnectionsGetItemsRequest {
	return ApiConnectionsGetItemsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
		externalItemId: externalItemId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExternalConnectorsExternalItem
func (a *ConnectionsExternalItemApiService) ConnectionsGetItemsExecute(r ApiConnectionsGetItemsRequest) (MicrosoftGraphExternalConnectorsExternalItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExternalConnectorsExternalItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalItemApiService.ConnectionsGetItems")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/items/{externalItem-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalItem-id"+"}", _neturl.PathEscape(parameterToString(r.externalItemId, "")), -1)

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

type ApiConnectionsListItemsRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalItemApiService
	externalConnectionId string
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
func (r ApiConnectionsListItemsRequest) Top(top int32) ApiConnectionsListItemsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiConnectionsListItemsRequest) Skip(skip int32) ApiConnectionsListItemsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiConnectionsListItemsRequest) Search(search string) ApiConnectionsListItemsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiConnectionsListItemsRequest) Filter(filter string) ApiConnectionsListItemsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiConnectionsListItemsRequest) Count(count bool) ApiConnectionsListItemsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiConnectionsListItemsRequest) Orderby(orderby []string) ApiConnectionsListItemsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiConnectionsListItemsRequest) Select_(select_ []string) ApiConnectionsListItemsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiConnectionsListItemsRequest) Expand(expand []string) ApiConnectionsListItemsRequest {
	r.expand = &expand
	return r
}

func (r ApiConnectionsListItemsRequest) Execute() (CollectionOfExternalItem, *_nethttp.Response, error) {
	return r.ApiService.ConnectionsListItemsExecute(r)
}

/*
ConnectionsListItems Get items from connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsListItemsRequest
*/
func (a *ConnectionsExternalItemApiService) ConnectionsListItems(ctx _context.Context, externalConnectionId string) ApiConnectionsListItemsRequest {
	return ApiConnectionsListItemsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
//  @return CollectionOfExternalItem
func (a *ConnectionsExternalItemApiService) ConnectionsListItemsExecute(r ApiConnectionsListItemsRequest) (CollectionOfExternalItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfExternalItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalItemApiService.ConnectionsListItems")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)

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

type ApiConnectionsUpdateItemsRequest struct {
	ctx _context.Context
	ApiService *ConnectionsExternalItemApiService
	externalConnectionId string
	externalItemId string
	microsoftGraphExternalConnectorsExternalItem *MicrosoftGraphExternalConnectorsExternalItem
}

// New navigation property values
func (r ApiConnectionsUpdateItemsRequest) MicrosoftGraphExternalConnectorsExternalItem(microsoftGraphExternalConnectorsExternalItem MicrosoftGraphExternalConnectorsExternalItem) ApiConnectionsUpdateItemsRequest {
	r.microsoftGraphExternalConnectorsExternalItem = &microsoftGraphExternalConnectorsExternalItem
	return r
}

func (r ApiConnectionsUpdateItemsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectionsUpdateItemsExecute(r)
}

/*
ConnectionsUpdateItems Update the navigation property items in connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @param externalItemId key: id of externalItem
 @return ApiConnectionsUpdateItemsRequest
*/
func (a *ConnectionsExternalItemApiService) ConnectionsUpdateItems(ctx _context.Context, externalConnectionId string, externalItemId string) ApiConnectionsUpdateItemsRequest {
	return ApiConnectionsUpdateItemsRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
		externalItemId: externalItemId,
	}
}

// Execute executes the request
func (a *ConnectionsExternalItemApiService) ConnectionsUpdateItemsExecute(r ApiConnectionsUpdateItemsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsExternalItemApiService.ConnectionsUpdateItems")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/items/{externalItem-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"externalItem-id"+"}", _neturl.PathEscape(parameterToString(r.externalItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphExternalConnectorsExternalItem == nil {
		return nil, reportError("microsoftGraphExternalConnectorsExternalItem is required and must be specified")
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
	localVarPostBody = r.microsoftGraphExternalConnectorsExternalItem
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
