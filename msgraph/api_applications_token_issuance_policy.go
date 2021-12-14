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

// ApplicationsTokenIssuancePolicyApiService ApplicationsTokenIssuancePolicyApi service
type ApplicationsTokenIssuancePolicyApiService service

type ApiApplicationsCreateRefTokenIssuancePoliciesRequest struct {
	ctx _context.Context
	ApiService *ApplicationsTokenIssuancePolicyApiService
	applicationId string
	requestBody *map[string]map[string]interface{}
}

// New navigation property ref value
func (r ApiApplicationsCreateRefTokenIssuancePoliciesRequest) RequestBody(requestBody map[string]map[string]interface{}) ApiApplicationsCreateRefTokenIssuancePoliciesRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiApplicationsCreateRefTokenIssuancePoliciesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ApplicationsCreateRefTokenIssuancePoliciesExecute(r)
}

/*
ApplicationsCreateRefTokenIssuancePolicies Create new navigation property ref to tokenIssuancePolicies for applications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId key: id of application
 @return ApiApplicationsCreateRefTokenIssuancePoliciesRequest
*/
func (a *ApplicationsTokenIssuancePolicyApiService) ApplicationsCreateRefTokenIssuancePolicies(ctx _context.Context, applicationId string) ApiApplicationsCreateRefTokenIssuancePoliciesRequest {
	return ApiApplicationsCreateRefTokenIssuancePoliciesRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ApplicationsTokenIssuancePolicyApiService) ApplicationsCreateRefTokenIssuancePoliciesExecute(r ApiApplicationsCreateRefTokenIssuancePoliciesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsTokenIssuancePolicyApiService.ApplicationsCreateRefTokenIssuancePolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application-id}/tokenIssuancePolicies/$ref"
	localVarPath = strings.Replace(localVarPath, "{"+"application-id"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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

type ApiApplicationsListRefTokenIssuancePoliciesRequest struct {
	ctx _context.Context
	ApiService *ApplicationsTokenIssuancePolicyApiService
	applicationId string
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
}

// Show only the first n items
func (r ApiApplicationsListRefTokenIssuancePoliciesRequest) Top(top int32) ApiApplicationsListRefTokenIssuancePoliciesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiApplicationsListRefTokenIssuancePoliciesRequest) Skip(skip int32) ApiApplicationsListRefTokenIssuancePoliciesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiApplicationsListRefTokenIssuancePoliciesRequest) Search(search string) ApiApplicationsListRefTokenIssuancePoliciesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiApplicationsListRefTokenIssuancePoliciesRequest) Filter(filter string) ApiApplicationsListRefTokenIssuancePoliciesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiApplicationsListRefTokenIssuancePoliciesRequest) Count(count bool) ApiApplicationsListRefTokenIssuancePoliciesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiApplicationsListRefTokenIssuancePoliciesRequest) Orderby(orderby []string) ApiApplicationsListRefTokenIssuancePoliciesRequest {
	r.orderby = &orderby
	return r
}

func (r ApiApplicationsListRefTokenIssuancePoliciesRequest) Execute() (CollectionOfLinksOfTokenIssuancePolicy, *_nethttp.Response, error) {
	return r.ApiService.ApplicationsListRefTokenIssuancePoliciesExecute(r)
}

/*
ApplicationsListRefTokenIssuancePolicies Get ref of tokenIssuancePolicies from applications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId key: id of application
 @return ApiApplicationsListRefTokenIssuancePoliciesRequest
*/
func (a *ApplicationsTokenIssuancePolicyApiService) ApplicationsListRefTokenIssuancePolicies(ctx _context.Context, applicationId string) ApiApplicationsListRefTokenIssuancePoliciesRequest {
	return ApiApplicationsListRefTokenIssuancePoliciesRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return CollectionOfLinksOfTokenIssuancePolicy
func (a *ApplicationsTokenIssuancePolicyApiService) ApplicationsListRefTokenIssuancePoliciesExecute(r ApiApplicationsListRefTokenIssuancePoliciesRequest) (CollectionOfLinksOfTokenIssuancePolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfLinksOfTokenIssuancePolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsTokenIssuancePolicyApiService.ApplicationsListRefTokenIssuancePolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application-id}/tokenIssuancePolicies/$ref"
	localVarPath = strings.Replace(localVarPath, "{"+"application-id"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

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

type ApiApplicationsListTokenIssuancePoliciesRequest struct {
	ctx _context.Context
	ApiService *ApplicationsTokenIssuancePolicyApiService
	applicationId string
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
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Top(top int32) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Skip(skip int32) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Search(search string) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Filter(filter string) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Count(count bool) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Orderby(orderby []string) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Select_(select_ []string) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiApplicationsListTokenIssuancePoliciesRequest) Expand(expand []string) ApiApplicationsListTokenIssuancePoliciesRequest {
	r.expand = &expand
	return r
}

func (r ApiApplicationsListTokenIssuancePoliciesRequest) Execute() (CollectionOfTokenIssuancePolicy, *_nethttp.Response, error) {
	return r.ApiService.ApplicationsListTokenIssuancePoliciesExecute(r)
}

/*
ApplicationsListTokenIssuancePolicies Get tokenIssuancePolicies from applications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId key: id of application
 @return ApiApplicationsListTokenIssuancePoliciesRequest
*/
func (a *ApplicationsTokenIssuancePolicyApiService) ApplicationsListTokenIssuancePolicies(ctx _context.Context, applicationId string) ApiApplicationsListTokenIssuancePoliciesRequest {
	return ApiApplicationsListTokenIssuancePoliciesRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return CollectionOfTokenIssuancePolicy
func (a *ApplicationsTokenIssuancePolicyApiService) ApplicationsListTokenIssuancePoliciesExecute(r ApiApplicationsListTokenIssuancePoliciesRequest) (CollectionOfTokenIssuancePolicy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfTokenIssuancePolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsTokenIssuancePolicyApiService.ApplicationsListTokenIssuancePolicies")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application-id}/tokenIssuancePolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"application-id"+"}", _neturl.PathEscape(parameterToString(r.applicationId, "")), -1)

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
