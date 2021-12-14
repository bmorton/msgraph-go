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

// IdentityIdentityProviderBaseApiService IdentityIdentityProviderBaseApi service
type IdentityIdentityProviderBaseApiService service

type ApiIdentityCreateIdentityProvidersRequest struct {
	ctx _context.Context
	ApiService *IdentityIdentityProviderBaseApiService
	microsoftGraphIdentityProviderBase *MicrosoftGraphIdentityProviderBase
}

// New navigation property
func (r ApiIdentityCreateIdentityProvidersRequest) MicrosoftGraphIdentityProviderBase(microsoftGraphIdentityProviderBase MicrosoftGraphIdentityProviderBase) ApiIdentityCreateIdentityProvidersRequest {
	r.microsoftGraphIdentityProviderBase = &microsoftGraphIdentityProviderBase
	return r
}

func (r ApiIdentityCreateIdentityProvidersRequest) Execute() (MicrosoftGraphIdentityProviderBase, *_nethttp.Response, error) {
	return r.ApiService.IdentityCreateIdentityProvidersExecute(r)
}

/*
IdentityCreateIdentityProviders Create new navigation property to identityProviders for identity

Represents entry point for identity provider base.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIdentityCreateIdentityProvidersRequest
*/
func (a *IdentityIdentityProviderBaseApiService) IdentityCreateIdentityProviders(ctx _context.Context) ApiIdentityCreateIdentityProvidersRequest {
	return ApiIdentityCreateIdentityProvidersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphIdentityProviderBase
func (a *IdentityIdentityProviderBaseApiService) IdentityCreateIdentityProvidersExecute(r ApiIdentityCreateIdentityProvidersRequest) (MicrosoftGraphIdentityProviderBase, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphIdentityProviderBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityIdentityProviderBaseApiService.IdentityCreateIdentityProviders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/identityProviders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphIdentityProviderBase == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphIdentityProviderBase is required and must be specified")
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
	localVarPostBody = r.microsoftGraphIdentityProviderBase
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

type ApiIdentityDeleteIdentityProvidersRequest struct {
	ctx _context.Context
	ApiService *IdentityIdentityProviderBaseApiService
	identityProviderBaseId string
	ifMatch *string
}

// ETag
func (r ApiIdentityDeleteIdentityProvidersRequest) IfMatch(ifMatch string) ApiIdentityDeleteIdentityProvidersRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiIdentityDeleteIdentityProvidersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdentityDeleteIdentityProvidersExecute(r)
}

/*
IdentityDeleteIdentityProviders Delete navigation property identityProviders for identity

Represents entry point for identity provider base.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identityProviderBaseId key: id of identityProviderBase
 @return ApiIdentityDeleteIdentityProvidersRequest
*/
func (a *IdentityIdentityProviderBaseApiService) IdentityDeleteIdentityProviders(ctx _context.Context, identityProviderBaseId string) ApiIdentityDeleteIdentityProvidersRequest {
	return ApiIdentityDeleteIdentityProvidersRequest{
		ApiService: a,
		ctx: ctx,
		identityProviderBaseId: identityProviderBaseId,
	}
}

// Execute executes the request
func (a *IdentityIdentityProviderBaseApiService) IdentityDeleteIdentityProvidersExecute(r ApiIdentityDeleteIdentityProvidersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityIdentityProviderBaseApiService.IdentityDeleteIdentityProviders")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/identityProviders/{identityProviderBase-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"identityProviderBase-id"+"}", _neturl.PathEscape(parameterToString(r.identityProviderBaseId, "")), -1)

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

type ApiIdentityGetIdentityProvidersRequest struct {
	ctx _context.Context
	ApiService *IdentityIdentityProviderBaseApiService
	identityProviderBaseId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiIdentityGetIdentityProvidersRequest) Select_(select_ []string) ApiIdentityGetIdentityProvidersRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiIdentityGetIdentityProvidersRequest) Expand(expand []string) ApiIdentityGetIdentityProvidersRequest {
	r.expand = &expand
	return r
}

func (r ApiIdentityGetIdentityProvidersRequest) Execute() (MicrosoftGraphIdentityProviderBase, *_nethttp.Response, error) {
	return r.ApiService.IdentityGetIdentityProvidersExecute(r)
}

/*
IdentityGetIdentityProviders Get identityProviders from identity

Represents entry point for identity provider base.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identityProviderBaseId key: id of identityProviderBase
 @return ApiIdentityGetIdentityProvidersRequest
*/
func (a *IdentityIdentityProviderBaseApiService) IdentityGetIdentityProviders(ctx _context.Context, identityProviderBaseId string) ApiIdentityGetIdentityProvidersRequest {
	return ApiIdentityGetIdentityProvidersRequest{
		ApiService: a,
		ctx: ctx,
		identityProviderBaseId: identityProviderBaseId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphIdentityProviderBase
func (a *IdentityIdentityProviderBaseApiService) IdentityGetIdentityProvidersExecute(r ApiIdentityGetIdentityProvidersRequest) (MicrosoftGraphIdentityProviderBase, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphIdentityProviderBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityIdentityProviderBaseApiService.IdentityGetIdentityProviders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/identityProviders/{identityProviderBase-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"identityProviderBase-id"+"}", _neturl.PathEscape(parameterToString(r.identityProviderBaseId, "")), -1)

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

type ApiIdentityListIdentityProvidersRequest struct {
	ctx _context.Context
	ApiService *IdentityIdentityProviderBaseApiService
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
func (r ApiIdentityListIdentityProvidersRequest) Top(top int32) ApiIdentityListIdentityProvidersRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiIdentityListIdentityProvidersRequest) Skip(skip int32) ApiIdentityListIdentityProvidersRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiIdentityListIdentityProvidersRequest) Search(search string) ApiIdentityListIdentityProvidersRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiIdentityListIdentityProvidersRequest) Filter(filter string) ApiIdentityListIdentityProvidersRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiIdentityListIdentityProvidersRequest) Count(count bool) ApiIdentityListIdentityProvidersRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiIdentityListIdentityProvidersRequest) Orderby(orderby []string) ApiIdentityListIdentityProvidersRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiIdentityListIdentityProvidersRequest) Select_(select_ []string) ApiIdentityListIdentityProvidersRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiIdentityListIdentityProvidersRequest) Expand(expand []string) ApiIdentityListIdentityProvidersRequest {
	r.expand = &expand
	return r
}

func (r ApiIdentityListIdentityProvidersRequest) Execute() (CollectionOfIdentityProviderBase, *_nethttp.Response, error) {
	return r.ApiService.IdentityListIdentityProvidersExecute(r)
}

/*
IdentityListIdentityProviders Get identityProviders from identity

Represents entry point for identity provider base.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIdentityListIdentityProvidersRequest
*/
func (a *IdentityIdentityProviderBaseApiService) IdentityListIdentityProviders(ctx _context.Context) ApiIdentityListIdentityProvidersRequest {
	return ApiIdentityListIdentityProvidersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfIdentityProviderBase
func (a *IdentityIdentityProviderBaseApiService) IdentityListIdentityProvidersExecute(r ApiIdentityListIdentityProvidersRequest) (CollectionOfIdentityProviderBase, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfIdentityProviderBase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityIdentityProviderBaseApiService.IdentityListIdentityProviders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/identityProviders"

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

type ApiIdentityUpdateIdentityProvidersRequest struct {
	ctx _context.Context
	ApiService *IdentityIdentityProviderBaseApiService
	identityProviderBaseId string
	microsoftGraphIdentityProviderBase *MicrosoftGraphIdentityProviderBase
}

// New navigation property values
func (r ApiIdentityUpdateIdentityProvidersRequest) MicrosoftGraphIdentityProviderBase(microsoftGraphIdentityProviderBase MicrosoftGraphIdentityProviderBase) ApiIdentityUpdateIdentityProvidersRequest {
	r.microsoftGraphIdentityProviderBase = &microsoftGraphIdentityProviderBase
	return r
}

func (r ApiIdentityUpdateIdentityProvidersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdentityUpdateIdentityProvidersExecute(r)
}

/*
IdentityUpdateIdentityProviders Update the navigation property identityProviders in identity

Represents entry point for identity provider base.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identityProviderBaseId key: id of identityProviderBase
 @return ApiIdentityUpdateIdentityProvidersRequest
*/
func (a *IdentityIdentityProviderBaseApiService) IdentityUpdateIdentityProviders(ctx _context.Context, identityProviderBaseId string) ApiIdentityUpdateIdentityProvidersRequest {
	return ApiIdentityUpdateIdentityProvidersRequest{
		ApiService: a,
		ctx: ctx,
		identityProviderBaseId: identityProviderBaseId,
	}
}

// Execute executes the request
func (a *IdentityIdentityProviderBaseApiService) IdentityUpdateIdentityProvidersExecute(r ApiIdentityUpdateIdentityProvidersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityIdentityProviderBaseApiService.IdentityUpdateIdentityProviders")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/identityProviders/{identityProviderBase-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"identityProviderBase-id"+"}", _neturl.PathEscape(parameterToString(r.identityProviderBaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphIdentityProviderBase == nil {
		return nil, reportError("microsoftGraphIdentityProviderBase is required and must be specified")
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
	localVarPostBody = r.microsoftGraphIdentityProviderBase
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
