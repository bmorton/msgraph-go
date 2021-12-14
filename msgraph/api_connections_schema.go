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

// ConnectionsSchemaApiService ConnectionsSchemaApi service
type ConnectionsSchemaApiService service

type ApiConnectionsDeleteSchemaRequest struct {
	ctx _context.Context
	ApiService *ConnectionsSchemaApiService
	externalConnectionId string
	ifMatch *string
}

// ETag
func (r ApiConnectionsDeleteSchemaRequest) IfMatch(ifMatch string) ApiConnectionsDeleteSchemaRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiConnectionsDeleteSchemaRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectionsDeleteSchemaExecute(r)
}

/*
ConnectionsDeleteSchema Delete navigation property schema for connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsDeleteSchemaRequest
*/
func (a *ConnectionsSchemaApiService) ConnectionsDeleteSchema(ctx _context.Context, externalConnectionId string) ApiConnectionsDeleteSchemaRequest {
	return ApiConnectionsDeleteSchemaRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
func (a *ConnectionsSchemaApiService) ConnectionsDeleteSchemaExecute(r ApiConnectionsDeleteSchemaRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSchemaApiService.ConnectionsDeleteSchema")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/schema"
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

type ApiConnectionsGetSchemaRequest struct {
	ctx _context.Context
	ApiService *ConnectionsSchemaApiService
	externalConnectionId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiConnectionsGetSchemaRequest) Select_(select_ []string) ApiConnectionsGetSchemaRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiConnectionsGetSchemaRequest) Expand(expand []string) ApiConnectionsGetSchemaRequest {
	r.expand = &expand
	return r
}

func (r ApiConnectionsGetSchemaRequest) Execute() (MicrosoftGraphExternalConnectorsSchema, *_nethttp.Response, error) {
	return r.ApiService.ConnectionsGetSchemaExecute(r)
}

/*
ConnectionsGetSchema Get schema from connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsGetSchemaRequest
*/
func (a *ConnectionsSchemaApiService) ConnectionsGetSchema(ctx _context.Context, externalConnectionId string) ApiConnectionsGetSchemaRequest {
	return ApiConnectionsGetSchemaRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphExternalConnectorsSchema
func (a *ConnectionsSchemaApiService) ConnectionsGetSchemaExecute(r ApiConnectionsGetSchemaRequest) (MicrosoftGraphExternalConnectorsSchema, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphExternalConnectorsSchema
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSchemaApiService.ConnectionsGetSchema")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/schema"
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

type ApiConnectionsUpdateSchemaRequest struct {
	ctx _context.Context
	ApiService *ConnectionsSchemaApiService
	externalConnectionId string
	microsoftGraphExternalConnectorsSchema *MicrosoftGraphExternalConnectorsSchema
}

// New navigation property values
func (r ApiConnectionsUpdateSchemaRequest) MicrosoftGraphExternalConnectorsSchema(microsoftGraphExternalConnectorsSchema MicrosoftGraphExternalConnectorsSchema) ApiConnectionsUpdateSchemaRequest {
	r.microsoftGraphExternalConnectorsSchema = &microsoftGraphExternalConnectorsSchema
	return r
}

func (r ApiConnectionsUpdateSchemaRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectionsUpdateSchemaExecute(r)
}

/*
ConnectionsUpdateSchema Update the navigation property schema in connections

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalConnectionId key: id of externalConnection
 @return ApiConnectionsUpdateSchemaRequest
*/
func (a *ConnectionsSchemaApiService) ConnectionsUpdateSchema(ctx _context.Context, externalConnectionId string) ApiConnectionsUpdateSchemaRequest {
	return ApiConnectionsUpdateSchemaRequest{
		ApiService: a,
		ctx: ctx,
		externalConnectionId: externalConnectionId,
	}
}

// Execute executes the request
func (a *ConnectionsSchemaApiService) ConnectionsUpdateSchemaExecute(r ApiConnectionsUpdateSchemaRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSchemaApiService.ConnectionsUpdateSchema")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{externalConnection-id}/schema"
	localVarPath = strings.Replace(localVarPath, "{"+"externalConnection-id"+"}", _neturl.PathEscape(parameterToString(r.externalConnectionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphExternalConnectorsSchema == nil {
		return nil, reportError("microsoftGraphExternalConnectorsSchema is required and must be specified")
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
	localVarPostBody = r.microsoftGraphExternalConnectorsSchema
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