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

// ServicePrincipalsServicePrincipalApiService ServicePrincipalsServicePrincipalApi service
type ServicePrincipalsServicePrincipalApiService service

type ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest struct {
	ctx _context.Context
	ApiService *ServicePrincipalsServicePrincipalApiService
	microsoftGraphServicePrincipal *MicrosoftGraphServicePrincipal
}

// New entity
func (r ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest) MicrosoftGraphServicePrincipal(microsoftGraphServicePrincipal MicrosoftGraphServicePrincipal) ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest {
	r.microsoftGraphServicePrincipal = &microsoftGraphServicePrincipal
	return r
}

func (r ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest) Execute() (MicrosoftGraphServicePrincipal, *_nethttp.Response, error) {
	return r.ApiService.ServicePrincipalsServicePrincipalCreateServicePrincipalExecute(r)
}

/*
ServicePrincipalsServicePrincipalCreateServicePrincipal Add new entity to servicePrincipals

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest
*/
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalCreateServicePrincipal(ctx _context.Context) ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest {
	return ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphServicePrincipal
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalCreateServicePrincipalExecute(r ApiServicePrincipalsServicePrincipalCreateServicePrincipalRequest) (MicrosoftGraphServicePrincipal, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphServicePrincipal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServicePrincipalsServicePrincipalApiService.ServicePrincipalsServicePrincipalCreateServicePrincipal")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/servicePrincipals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphServicePrincipal == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphServicePrincipal is required and must be specified")
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
	localVarPostBody = r.microsoftGraphServicePrincipal
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

type ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest struct {
	ctx _context.Context
	ApiService *ServicePrincipalsServicePrincipalApiService
	servicePrincipalId string
	ifMatch *string
}

// ETag
func (r ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest) IfMatch(ifMatch string) ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ServicePrincipalsServicePrincipalDeleteServicePrincipalExecute(r)
}

/*
ServicePrincipalsServicePrincipalDeleteServicePrincipal Delete entity from servicePrincipals

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param servicePrincipalId key: id of servicePrincipal
 @return ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest
*/
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalDeleteServicePrincipal(ctx _context.Context, servicePrincipalId string) ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest {
	return ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest{
		ApiService: a,
		ctx: ctx,
		servicePrincipalId: servicePrincipalId,
	}
}

// Execute executes the request
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalDeleteServicePrincipalExecute(r ApiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServicePrincipalsServicePrincipalApiService.ServicePrincipalsServicePrincipalDeleteServicePrincipal")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/servicePrincipals/{servicePrincipal-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"servicePrincipal-id"+"}", _neturl.PathEscape(parameterToString(r.servicePrincipalId, "")), -1)

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

type ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest struct {
	ctx _context.Context
	ApiService *ServicePrincipalsServicePrincipalApiService
	servicePrincipalId string
	consistencyLevel *string
	select_ *[]string
	expand *[]string
}

// Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/
func (r ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest) ConsistencyLevel(consistencyLevel string) ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest {
	r.consistencyLevel = &consistencyLevel
	return r
}
// Select properties to be returned
func (r ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest) Select_(select_ []string) ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest) Expand(expand []string) ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest {
	r.expand = &expand
	return r
}

func (r ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest) Execute() (MicrosoftGraphServicePrincipal, *_nethttp.Response, error) {
	return r.ApiService.ServicePrincipalsServicePrincipalGetServicePrincipalExecute(r)
}

/*
ServicePrincipalsServicePrincipalGetServicePrincipal Get entity from servicePrincipals by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param servicePrincipalId key: id of servicePrincipal
 @return ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest
*/
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalGetServicePrincipal(ctx _context.Context, servicePrincipalId string) ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest {
	return ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest{
		ApiService: a,
		ctx: ctx,
		servicePrincipalId: servicePrincipalId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphServicePrincipal
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalGetServicePrincipalExecute(r ApiServicePrincipalsServicePrincipalGetServicePrincipalRequest) (MicrosoftGraphServicePrincipal, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphServicePrincipal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServicePrincipalsServicePrincipalApiService.ServicePrincipalsServicePrincipalGetServicePrincipal")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/servicePrincipals/{servicePrincipal-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"servicePrincipal-id"+"}", _neturl.PathEscape(parameterToString(r.servicePrincipalId, "")), -1)

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
	if r.consistencyLevel != nil {
		localVarHeaderParams["ConsistencyLevel"] = parameterToString(*r.consistencyLevel, "")
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

type ApiServicePrincipalsServicePrincipalListServicePrincipalRequest struct {
	ctx _context.Context
	ApiService *ServicePrincipalsServicePrincipalApiService
	consistencyLevel *string
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
	select_ *[]string
	expand *[]string
}

// Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) ConsistencyLevel(consistencyLevel string) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.consistencyLevel = &consistencyLevel
	return r
}
// Show only the first n items
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Top(top int32) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Skip(skip int32) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Search(search string) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Filter(filter string) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Count(count bool) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Orderby(orderby []string) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Select_(select_ []string) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Expand(expand []string) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	r.expand = &expand
	return r
}

func (r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) Execute() (CollectionOfServicePrincipal, *_nethttp.Response, error) {
	return r.ApiService.ServicePrincipalsServicePrincipalListServicePrincipalExecute(r)
}

/*
ServicePrincipalsServicePrincipalListServicePrincipal Get entities from servicePrincipals

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServicePrincipalsServicePrincipalListServicePrincipalRequest
*/
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalListServicePrincipal(ctx _context.Context) ApiServicePrincipalsServicePrincipalListServicePrincipalRequest {
	return ApiServicePrincipalsServicePrincipalListServicePrincipalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfServicePrincipal
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalListServicePrincipalExecute(r ApiServicePrincipalsServicePrincipalListServicePrincipalRequest) (CollectionOfServicePrincipal, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfServicePrincipal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServicePrincipalsServicePrincipalApiService.ServicePrincipalsServicePrincipalListServicePrincipal")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/servicePrincipals"

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
	if r.consistencyLevel != nil {
		localVarHeaderParams["ConsistencyLevel"] = parameterToString(*r.consistencyLevel, "")
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

type ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest struct {
	ctx _context.Context
	ApiService *ServicePrincipalsServicePrincipalApiService
	servicePrincipalId string
	microsoftGraphServicePrincipal *MicrosoftGraphServicePrincipal
}

// New property values
func (r ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest) MicrosoftGraphServicePrincipal(microsoftGraphServicePrincipal MicrosoftGraphServicePrincipal) ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest {
	r.microsoftGraphServicePrincipal = &microsoftGraphServicePrincipal
	return r
}

func (r ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ServicePrincipalsServicePrincipalUpdateServicePrincipalExecute(r)
}

/*
ServicePrincipalsServicePrincipalUpdateServicePrincipal Update entity in servicePrincipals

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param servicePrincipalId key: id of servicePrincipal
 @return ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest
*/
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalUpdateServicePrincipal(ctx _context.Context, servicePrincipalId string) ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest {
	return ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest{
		ApiService: a,
		ctx: ctx,
		servicePrincipalId: servicePrincipalId,
	}
}

// Execute executes the request
func (a *ServicePrincipalsServicePrincipalApiService) ServicePrincipalsServicePrincipalUpdateServicePrincipalExecute(r ApiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServicePrincipalsServicePrincipalApiService.ServicePrincipalsServicePrincipalUpdateServicePrincipal")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/servicePrincipals/{servicePrincipal-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"servicePrincipal-id"+"}", _neturl.PathEscape(parameterToString(r.servicePrincipalId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphServicePrincipal == nil {
		return nil, reportError("microsoftGraphServicePrincipal is required and must be specified")
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
	localVarPostBody = r.microsoftGraphServicePrincipal
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
