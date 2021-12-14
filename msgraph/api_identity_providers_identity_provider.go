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

// IdentityProvidersIdentityProviderApiService IdentityProvidersIdentityProviderApi service
type IdentityProvidersIdentityProviderApiService service

type ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest struct {
	ctx _context.Context
	ApiService *IdentityProvidersIdentityProviderApiService
	microsoftGraphIdentityProvider *MicrosoftGraphIdentityProvider
}

// New entity
func (r ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest) MicrosoftGraphIdentityProvider(microsoftGraphIdentityProvider MicrosoftGraphIdentityProvider) ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest {
	r.microsoftGraphIdentityProvider = &microsoftGraphIdentityProvider
	return r
}

func (r ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest) Execute() (MicrosoftGraphIdentityProvider, *_nethttp.Response, error) {
	return r.ApiService.IdentityProvidersIdentityProviderCreateIdentityProviderExecute(r)
}

/*
IdentityProvidersIdentityProviderCreateIdentityProvider Add new entity to identityProviders

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest
*/
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderCreateIdentityProvider(ctx _context.Context) ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest {
	return ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphIdentityProvider
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderCreateIdentityProviderExecute(r ApiIdentityProvidersIdentityProviderCreateIdentityProviderRequest) (MicrosoftGraphIdentityProvider, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphIdentityProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProvidersIdentityProviderApiService.IdentityProvidersIdentityProviderCreateIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityProviders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphIdentityProvider == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphIdentityProvider is required and must be specified")
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
	localVarPostBody = r.microsoftGraphIdentityProvider
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

type ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest struct {
	ctx _context.Context
	ApiService *IdentityProvidersIdentityProviderApiService
	identityProviderId string
	ifMatch *string
}

// ETag
func (r ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest) IfMatch(ifMatch string) ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdentityProvidersIdentityProviderDeleteIdentityProviderExecute(r)
}

/*
IdentityProvidersIdentityProviderDeleteIdentityProvider Delete entity from identityProviders

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identityProviderId key: id of identityProvider
 @return ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest
*/
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderDeleteIdentityProvider(ctx _context.Context, identityProviderId string) ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest {
	return ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest{
		ApiService: a,
		ctx: ctx,
		identityProviderId: identityProviderId,
	}
}

// Execute executes the request
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderDeleteIdentityProviderExecute(r ApiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProvidersIdentityProviderApiService.IdentityProvidersIdentityProviderDeleteIdentityProvider")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityProviders/{identityProvider-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"identityProvider-id"+"}", _neturl.PathEscape(parameterToString(r.identityProviderId, "")), -1)

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

type ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest struct {
	ctx _context.Context
	ApiService *IdentityProvidersIdentityProviderApiService
	identityProviderId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest) Select_(select_ []string) ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest) Expand(expand []string) ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest {
	r.expand = &expand
	return r
}

func (r ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest) Execute() (MicrosoftGraphIdentityProvider, *_nethttp.Response, error) {
	return r.ApiService.IdentityProvidersIdentityProviderGetIdentityProviderExecute(r)
}

/*
IdentityProvidersIdentityProviderGetIdentityProvider Get entity from identityProviders by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identityProviderId key: id of identityProvider
 @return ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest
*/
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderGetIdentityProvider(ctx _context.Context, identityProviderId string) ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest {
	return ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest{
		ApiService: a,
		ctx: ctx,
		identityProviderId: identityProviderId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphIdentityProvider
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderGetIdentityProviderExecute(r ApiIdentityProvidersIdentityProviderGetIdentityProviderRequest) (MicrosoftGraphIdentityProvider, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphIdentityProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProvidersIdentityProviderApiService.IdentityProvidersIdentityProviderGetIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityProviders/{identityProvider-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"identityProvider-id"+"}", _neturl.PathEscape(parameterToString(r.identityProviderId, "")), -1)

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

type ApiIdentityProvidersIdentityProviderListIdentityProviderRequest struct {
	ctx _context.Context
	ApiService *IdentityProvidersIdentityProviderApiService
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
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Top(top int32) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Skip(skip int32) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Search(search string) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Filter(filter string) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Count(count bool) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Orderby(orderby []string) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Select_(select_ []string) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Expand(expand []string) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	r.expand = &expand
	return r
}

func (r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) Execute() (CollectionOfIdentityProvider, *_nethttp.Response, error) {
	return r.ApiService.IdentityProvidersIdentityProviderListIdentityProviderExecute(r)
}

/*
IdentityProvidersIdentityProviderListIdentityProvider Get entities from identityProviders

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIdentityProvidersIdentityProviderListIdentityProviderRequest
*/
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderListIdentityProvider(ctx _context.Context) ApiIdentityProvidersIdentityProviderListIdentityProviderRequest {
	return ApiIdentityProvidersIdentityProviderListIdentityProviderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfIdentityProvider
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderListIdentityProviderExecute(r ApiIdentityProvidersIdentityProviderListIdentityProviderRequest) (CollectionOfIdentityProvider, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfIdentityProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProvidersIdentityProviderApiService.IdentityProvidersIdentityProviderListIdentityProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityProviders"

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

type ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest struct {
	ctx _context.Context
	ApiService *IdentityProvidersIdentityProviderApiService
	identityProviderId string
	microsoftGraphIdentityProvider *MicrosoftGraphIdentityProvider
}

// New property values
func (r ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest) MicrosoftGraphIdentityProvider(microsoftGraphIdentityProvider MicrosoftGraphIdentityProvider) ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest {
	r.microsoftGraphIdentityProvider = &microsoftGraphIdentityProvider
	return r
}

func (r ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdentityProvidersIdentityProviderUpdateIdentityProviderExecute(r)
}

/*
IdentityProvidersIdentityProviderUpdateIdentityProvider Update entity in identityProviders

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identityProviderId key: id of identityProvider
 @return ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest
*/
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderUpdateIdentityProvider(ctx _context.Context, identityProviderId string) ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest {
	return ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest{
		ApiService: a,
		ctx: ctx,
		identityProviderId: identityProviderId,
	}
}

// Execute executes the request
func (a *IdentityProvidersIdentityProviderApiService) IdentityProvidersIdentityProviderUpdateIdentityProviderExecute(r ApiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityProvidersIdentityProviderApiService.IdentityProvidersIdentityProviderUpdateIdentityProvider")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityProviders/{identityProvider-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"identityProvider-id"+"}", _neturl.PathEscape(parameterToString(r.identityProviderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphIdentityProvider == nil {
		return nil, reportError("microsoftGraphIdentityProvider is required and must be specified")
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
	localVarPostBody = r.microsoftGraphIdentityProvider
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
