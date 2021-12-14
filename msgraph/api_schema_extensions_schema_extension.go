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

// SchemaExtensionsSchemaExtensionApiService SchemaExtensionsSchemaExtensionApi service
type SchemaExtensionsSchemaExtensionApiService service

type ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest struct {
	ctx _context.Context
	ApiService *SchemaExtensionsSchemaExtensionApiService
	microsoftGraphSchemaExtension *MicrosoftGraphSchemaExtension
}

// New entity
func (r ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest) MicrosoftGraphSchemaExtension(microsoftGraphSchemaExtension MicrosoftGraphSchemaExtension) ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest {
	r.microsoftGraphSchemaExtension = &microsoftGraphSchemaExtension
	return r
}

func (r ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest) Execute() (MicrosoftGraphSchemaExtension, *_nethttp.Response, error) {
	return r.ApiService.SchemaExtensionsSchemaExtensionCreateSchemaExtensionExecute(r)
}

/*
SchemaExtensionsSchemaExtensionCreateSchemaExtension Add new entity to schemaExtensions

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest
*/
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionCreateSchemaExtension(ctx _context.Context) ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest {
	return ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSchemaExtension
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionCreateSchemaExtensionExecute(r ApiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest) (MicrosoftGraphSchemaExtension, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSchemaExtension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaExtensionsSchemaExtensionApiService.SchemaExtensionsSchemaExtensionCreateSchemaExtension")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schemaExtensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSchemaExtension == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphSchemaExtension is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSchemaExtension
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

type ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest struct {
	ctx _context.Context
	ApiService *SchemaExtensionsSchemaExtensionApiService
	schemaExtensionId string
	ifMatch *string
}

// ETag
func (r ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest) IfMatch(ifMatch string) ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SchemaExtensionsSchemaExtensionDeleteSchemaExtensionExecute(r)
}

/*
SchemaExtensionsSchemaExtensionDeleteSchemaExtension Delete entity from schemaExtensions

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param schemaExtensionId key: id of schemaExtension
 @return ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest
*/
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionDeleteSchemaExtension(ctx _context.Context, schemaExtensionId string) ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest {
	return ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest{
		ApiService: a,
		ctx: ctx,
		schemaExtensionId: schemaExtensionId,
	}
}

// Execute executes the request
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionDeleteSchemaExtensionExecute(r ApiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaExtensionsSchemaExtensionApiService.SchemaExtensionsSchemaExtensionDeleteSchemaExtension")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schemaExtensions/{schemaExtension-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"schemaExtension-id"+"}", _neturl.PathEscape(parameterToString(r.schemaExtensionId, "")), -1)

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

type ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest struct {
	ctx _context.Context
	ApiService *SchemaExtensionsSchemaExtensionApiService
	schemaExtensionId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest) Select_(select_ []string) ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest) Expand(expand []string) ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest {
	r.expand = &expand
	return r
}

func (r ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest) Execute() (MicrosoftGraphSchemaExtension, *_nethttp.Response, error) {
	return r.ApiService.SchemaExtensionsSchemaExtensionGetSchemaExtensionExecute(r)
}

/*
SchemaExtensionsSchemaExtensionGetSchemaExtension Get entity from schemaExtensions by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param schemaExtensionId key: id of schemaExtension
 @return ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest
*/
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionGetSchemaExtension(ctx _context.Context, schemaExtensionId string) ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest {
	return ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest{
		ApiService: a,
		ctx: ctx,
		schemaExtensionId: schemaExtensionId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSchemaExtension
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionGetSchemaExtensionExecute(r ApiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest) (MicrosoftGraphSchemaExtension, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSchemaExtension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaExtensionsSchemaExtensionApiService.SchemaExtensionsSchemaExtensionGetSchemaExtension")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schemaExtensions/{schemaExtension-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"schemaExtension-id"+"}", _neturl.PathEscape(parameterToString(r.schemaExtensionId, "")), -1)

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

type ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest struct {
	ctx _context.Context
	ApiService *SchemaExtensionsSchemaExtensionApiService
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
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Top(top int32) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Skip(skip int32) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Search(search string) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Filter(filter string) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Count(count bool) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Orderby(orderby []string) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Select_(select_ []string) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Expand(expand []string) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	r.expand = &expand
	return r
}

func (r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) Execute() (CollectionOfSchemaExtension, *_nethttp.Response, error) {
	return r.ApiService.SchemaExtensionsSchemaExtensionListSchemaExtensionExecute(r)
}

/*
SchemaExtensionsSchemaExtensionListSchemaExtension Get entities from schemaExtensions

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest
*/
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionListSchemaExtension(ctx _context.Context) ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest {
	return ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfSchemaExtension
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionListSchemaExtensionExecute(r ApiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest) (CollectionOfSchemaExtension, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfSchemaExtension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaExtensionsSchemaExtensionApiService.SchemaExtensionsSchemaExtensionListSchemaExtension")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schemaExtensions"

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

type ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest struct {
	ctx _context.Context
	ApiService *SchemaExtensionsSchemaExtensionApiService
	schemaExtensionId string
	microsoftGraphSchemaExtension *MicrosoftGraphSchemaExtension
}

// New property values
func (r ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest) MicrosoftGraphSchemaExtension(microsoftGraphSchemaExtension MicrosoftGraphSchemaExtension) ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest {
	r.microsoftGraphSchemaExtension = &microsoftGraphSchemaExtension
	return r
}

func (r ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SchemaExtensionsSchemaExtensionUpdateSchemaExtensionExecute(r)
}

/*
SchemaExtensionsSchemaExtensionUpdateSchemaExtension Update entity in schemaExtensions

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param schemaExtensionId key: id of schemaExtension
 @return ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest
*/
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionUpdateSchemaExtension(ctx _context.Context, schemaExtensionId string) ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest {
	return ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest{
		ApiService: a,
		ctx: ctx,
		schemaExtensionId: schemaExtensionId,
	}
}

// Execute executes the request
func (a *SchemaExtensionsSchemaExtensionApiService) SchemaExtensionsSchemaExtensionUpdateSchemaExtensionExecute(r ApiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaExtensionsSchemaExtensionApiService.SchemaExtensionsSchemaExtensionUpdateSchemaExtension")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schemaExtensions/{schemaExtension-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"schemaExtension-id"+"}", _neturl.PathEscape(parameterToString(r.schemaExtensionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSchemaExtension == nil {
		return nil, reportError("microsoftGraphSchemaExtension is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSchemaExtension
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
