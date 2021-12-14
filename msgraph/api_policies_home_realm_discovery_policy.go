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

// PoliciesHomeRealmDiscoveryPolicyApiService PoliciesHomeRealmDiscoveryPolicyApi service
type PoliciesHomeRealmDiscoveryPolicyApiService service

type ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest struct {
	ctx _context.Context
	ApiService *PoliciesHomeRealmDiscoveryPolicyApiService
	microsoftGraphHomeRealmDiscoveryPolicy *MicrosoftGraphHomeRealmDiscoveryPolicy
}

// New navigation property
func (r ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest) MicrosoftGraphHomeRealmDiscoveryPolicy(microsoftGraphHomeRealmDiscoveryPolicy MicrosoftGraphHomeRealmDiscoveryPolicy) ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest {
	r.microsoftGraphHomeRealmDiscoveryPolicy = &microsoftGraphHomeRealmDiscoveryPolicy
	return r
}

func (r ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest) Execute() (MicrosoftGraphHomeRealmDiscoveryPolicy, *_nethttp.Response, error) {
	return r.ApiService.PoliciesCreateHomeRealmDiscoveryPoliciesExecute(r)
}

/*
PoliciesCreateHomeRealmDiscoveryPolicies Create new navigation property to homeRealmDiscoveryPolicies for policies

The policy to control Azure AD authentication behavior for federated users.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest
*/
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesCreateHomeRealmDiscoveryPolicies(ctx _context.Context) ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest {
	return ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphHomeRealmDiscoveryPolicy
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesCreateHomeRealmDiscoveryPoliciesExecute(r ApiPoliciesCreateHomeRealmDiscoveryPoliciesRequest) (MicrosoftGraphHomeRealmDiscoveryPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphHomeRealmDiscoveryPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesHomeRealmDiscoveryPolicyApiService.PoliciesCreateHomeRealmDiscoveryPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/homeRealmDiscoveryPolicies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphHomeRealmDiscoveryPolicy == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphHomeRealmDiscoveryPolicy is required and must be specified")
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
	localVarPostBody = r.microsoftGraphHomeRealmDiscoveryPolicy
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

type ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest struct {
	ctx _context.Context
	ApiService *PoliciesHomeRealmDiscoveryPolicyApiService
	homeRealmDiscoveryPolicyId string
	ifMatch *string
}

// ETag
func (r ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest) IfMatch(ifMatch string) ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PoliciesDeleteHomeRealmDiscoveryPoliciesExecute(r)
}

/*
PoliciesDeleteHomeRealmDiscoveryPolicies Delete navigation property homeRealmDiscoveryPolicies for policies

The policy to control Azure AD authentication behavior for federated users.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param homeRealmDiscoveryPolicyId key: id of homeRealmDiscoveryPolicy
 @return ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest
*/
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesDeleteHomeRealmDiscoveryPolicies(ctx _context.Context, homeRealmDiscoveryPolicyId string) ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest {
	return ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		homeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// Execute executes the request
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesDeleteHomeRealmDiscoveryPoliciesExecute(r ApiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesHomeRealmDiscoveryPolicyApiService.PoliciesDeleteHomeRealmDiscoveryPolicies")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"homeRealmDiscoveryPolicy-id"+"}", _neturl.PathEscape(parameterToString(r.homeRealmDiscoveryPolicyId, "")), -1)

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

type ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest struct {
	ctx _context.Context
	ApiService *PoliciesHomeRealmDiscoveryPolicyApiService
	homeRealmDiscoveryPolicyId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest) Select_(select_ []string) ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest) Expand(expand []string) ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest {
	r.expand = &expand
	return r
}

func (r ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest) Execute() (MicrosoftGraphHomeRealmDiscoveryPolicy, *_nethttp.Response, error) {
	return r.ApiService.PoliciesGetHomeRealmDiscoveryPoliciesExecute(r)
}

/*
PoliciesGetHomeRealmDiscoveryPolicies Get homeRealmDiscoveryPolicies from policies

The policy to control Azure AD authentication behavior for federated users.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param homeRealmDiscoveryPolicyId key: id of homeRealmDiscoveryPolicy
 @return ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest
*/
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesGetHomeRealmDiscoveryPolicies(ctx _context.Context, homeRealmDiscoveryPolicyId string) ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest {
	return ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		homeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphHomeRealmDiscoveryPolicy
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesGetHomeRealmDiscoveryPoliciesExecute(r ApiPoliciesGetHomeRealmDiscoveryPoliciesRequest) (MicrosoftGraphHomeRealmDiscoveryPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphHomeRealmDiscoveryPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesHomeRealmDiscoveryPolicyApiService.PoliciesGetHomeRealmDiscoveryPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"homeRealmDiscoveryPolicy-id"+"}", _neturl.PathEscape(parameterToString(r.homeRealmDiscoveryPolicyId, "")), -1)

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

type ApiPoliciesListHomeRealmDiscoveryPoliciesRequest struct {
	ctx _context.Context
	ApiService *PoliciesHomeRealmDiscoveryPolicyApiService
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
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Top(top int32) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Skip(skip int32) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Search(search string) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Filter(filter string) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Count(count bool) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Orderby(orderby []string) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Select_(select_ []string) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Expand(expand []string) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	r.expand = &expand
	return r
}

func (r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) Execute() (CollectionOfHomeRealmDiscoveryPolicy, *_nethttp.Response, error) {
	return r.ApiService.PoliciesListHomeRealmDiscoveryPoliciesExecute(r)
}

/*
PoliciesListHomeRealmDiscoveryPolicies Get homeRealmDiscoveryPolicies from policies

The policy to control Azure AD authentication behavior for federated users.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPoliciesListHomeRealmDiscoveryPoliciesRequest
*/
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesListHomeRealmDiscoveryPolicies(ctx _context.Context) ApiPoliciesListHomeRealmDiscoveryPoliciesRequest {
	return ApiPoliciesListHomeRealmDiscoveryPoliciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfHomeRealmDiscoveryPolicy
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesListHomeRealmDiscoveryPoliciesExecute(r ApiPoliciesListHomeRealmDiscoveryPoliciesRequest) (CollectionOfHomeRealmDiscoveryPolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfHomeRealmDiscoveryPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesHomeRealmDiscoveryPolicyApiService.PoliciesListHomeRealmDiscoveryPolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/homeRealmDiscoveryPolicies"

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

type ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest struct {
	ctx _context.Context
	ApiService *PoliciesHomeRealmDiscoveryPolicyApiService
	homeRealmDiscoveryPolicyId string
	microsoftGraphHomeRealmDiscoveryPolicy *MicrosoftGraphHomeRealmDiscoveryPolicy
}

// New navigation property values
func (r ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest) MicrosoftGraphHomeRealmDiscoveryPolicy(microsoftGraphHomeRealmDiscoveryPolicy MicrosoftGraphHomeRealmDiscoveryPolicy) ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest {
	r.microsoftGraphHomeRealmDiscoveryPolicy = &microsoftGraphHomeRealmDiscoveryPolicy
	return r
}

func (r ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PoliciesUpdateHomeRealmDiscoveryPoliciesExecute(r)
}

/*
PoliciesUpdateHomeRealmDiscoveryPolicies Update the navigation property homeRealmDiscoveryPolicies in policies

The policy to control Azure AD authentication behavior for federated users.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param homeRealmDiscoveryPolicyId key: id of homeRealmDiscoveryPolicy
 @return ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest
*/
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesUpdateHomeRealmDiscoveryPolicies(ctx _context.Context, homeRealmDiscoveryPolicyId string) ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest {
	return ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest{
		ApiService: a,
		ctx: ctx,
		homeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// Execute executes the request
func (a *PoliciesHomeRealmDiscoveryPolicyApiService) PoliciesUpdateHomeRealmDiscoveryPoliciesExecute(r ApiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesHomeRealmDiscoveryPolicyApiService.PoliciesUpdateHomeRealmDiscoveryPolicies")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"homeRealmDiscoveryPolicy-id"+"}", _neturl.PathEscape(parameterToString(r.homeRealmDiscoveryPolicyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphHomeRealmDiscoveryPolicy == nil {
		return nil, reportError("microsoftGraphHomeRealmDiscoveryPolicy is required and must be specified")
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
	localVarPostBody = r.microsoftGraphHomeRealmDiscoveryPolicy
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
