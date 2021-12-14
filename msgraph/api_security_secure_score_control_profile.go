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

// SecuritySecureScoreControlProfileApiService SecuritySecureScoreControlProfileApi service
type SecuritySecureScoreControlProfileApiService service

type ApiSecurityCreateSecureScoreControlProfilesRequest struct {
	ctx _context.Context
	ApiService *SecuritySecureScoreControlProfileApiService
	microsoftGraphSecureScoreControlProfile *MicrosoftGraphSecureScoreControlProfile
}

// New navigation property
func (r ApiSecurityCreateSecureScoreControlProfilesRequest) MicrosoftGraphSecureScoreControlProfile(microsoftGraphSecureScoreControlProfile MicrosoftGraphSecureScoreControlProfile) ApiSecurityCreateSecureScoreControlProfilesRequest {
	r.microsoftGraphSecureScoreControlProfile = &microsoftGraphSecureScoreControlProfile
	return r
}

func (r ApiSecurityCreateSecureScoreControlProfilesRequest) Execute() (MicrosoftGraphSecureScoreControlProfile, *_nethttp.Response, error) {
	return r.ApiService.SecurityCreateSecureScoreControlProfilesExecute(r)
}

/*
SecurityCreateSecureScoreControlProfiles Create new navigation property to secureScoreControlProfiles for security

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSecurityCreateSecureScoreControlProfilesRequest
*/
func (a *SecuritySecureScoreControlProfileApiService) SecurityCreateSecureScoreControlProfiles(ctx _context.Context) ApiSecurityCreateSecureScoreControlProfilesRequest {
	return ApiSecurityCreateSecureScoreControlProfilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSecureScoreControlProfile
func (a *SecuritySecureScoreControlProfileApiService) SecurityCreateSecureScoreControlProfilesExecute(r ApiSecurityCreateSecureScoreControlProfilesRequest) (MicrosoftGraphSecureScoreControlProfile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSecureScoreControlProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecuritySecureScoreControlProfileApiService.SecurityCreateSecureScoreControlProfiles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/secureScoreControlProfiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSecureScoreControlProfile == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphSecureScoreControlProfile is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSecureScoreControlProfile
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

type ApiSecurityDeleteSecureScoreControlProfilesRequest struct {
	ctx _context.Context
	ApiService *SecuritySecureScoreControlProfileApiService
	secureScoreControlProfileId string
	ifMatch *string
}

// ETag
func (r ApiSecurityDeleteSecureScoreControlProfilesRequest) IfMatch(ifMatch string) ApiSecurityDeleteSecureScoreControlProfilesRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiSecurityDeleteSecureScoreControlProfilesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SecurityDeleteSecureScoreControlProfilesExecute(r)
}

/*
SecurityDeleteSecureScoreControlProfiles Delete navigation property secureScoreControlProfiles for security

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secureScoreControlProfileId key: id of secureScoreControlProfile
 @return ApiSecurityDeleteSecureScoreControlProfilesRequest
*/
func (a *SecuritySecureScoreControlProfileApiService) SecurityDeleteSecureScoreControlProfiles(ctx _context.Context, secureScoreControlProfileId string) ApiSecurityDeleteSecureScoreControlProfilesRequest {
	return ApiSecurityDeleteSecureScoreControlProfilesRequest{
		ApiService: a,
		ctx: ctx,
		secureScoreControlProfileId: secureScoreControlProfileId,
	}
}

// Execute executes the request
func (a *SecuritySecureScoreControlProfileApiService) SecurityDeleteSecureScoreControlProfilesExecute(r ApiSecurityDeleteSecureScoreControlProfilesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecuritySecureScoreControlProfileApiService.SecurityDeleteSecureScoreControlProfiles")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/secureScoreControlProfiles/{secureScoreControlProfile-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"secureScoreControlProfile-id"+"}", _neturl.PathEscape(parameterToString(r.secureScoreControlProfileId, "")), -1)

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

type ApiSecurityGetSecureScoreControlProfilesRequest struct {
	ctx _context.Context
	ApiService *SecuritySecureScoreControlProfileApiService
	secureScoreControlProfileId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiSecurityGetSecureScoreControlProfilesRequest) Select_(select_ []string) ApiSecurityGetSecureScoreControlProfilesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiSecurityGetSecureScoreControlProfilesRequest) Expand(expand []string) ApiSecurityGetSecureScoreControlProfilesRequest {
	r.expand = &expand
	return r
}

func (r ApiSecurityGetSecureScoreControlProfilesRequest) Execute() (MicrosoftGraphSecureScoreControlProfile, *_nethttp.Response, error) {
	return r.ApiService.SecurityGetSecureScoreControlProfilesExecute(r)
}

/*
SecurityGetSecureScoreControlProfiles Get secureScoreControlProfiles from security

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secureScoreControlProfileId key: id of secureScoreControlProfile
 @return ApiSecurityGetSecureScoreControlProfilesRequest
*/
func (a *SecuritySecureScoreControlProfileApiService) SecurityGetSecureScoreControlProfiles(ctx _context.Context, secureScoreControlProfileId string) ApiSecurityGetSecureScoreControlProfilesRequest {
	return ApiSecurityGetSecureScoreControlProfilesRequest{
		ApiService: a,
		ctx: ctx,
		secureScoreControlProfileId: secureScoreControlProfileId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSecureScoreControlProfile
func (a *SecuritySecureScoreControlProfileApiService) SecurityGetSecureScoreControlProfilesExecute(r ApiSecurityGetSecureScoreControlProfilesRequest) (MicrosoftGraphSecureScoreControlProfile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSecureScoreControlProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecuritySecureScoreControlProfileApiService.SecurityGetSecureScoreControlProfiles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/secureScoreControlProfiles/{secureScoreControlProfile-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"secureScoreControlProfile-id"+"}", _neturl.PathEscape(parameterToString(r.secureScoreControlProfileId, "")), -1)

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

type ApiSecurityListSecureScoreControlProfilesRequest struct {
	ctx _context.Context
	ApiService *SecuritySecureScoreControlProfileApiService
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
func (r ApiSecurityListSecureScoreControlProfilesRequest) Top(top int32) ApiSecurityListSecureScoreControlProfilesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiSecurityListSecureScoreControlProfilesRequest) Skip(skip int32) ApiSecurityListSecureScoreControlProfilesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiSecurityListSecureScoreControlProfilesRequest) Search(search string) ApiSecurityListSecureScoreControlProfilesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiSecurityListSecureScoreControlProfilesRequest) Filter(filter string) ApiSecurityListSecureScoreControlProfilesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiSecurityListSecureScoreControlProfilesRequest) Count(count bool) ApiSecurityListSecureScoreControlProfilesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiSecurityListSecureScoreControlProfilesRequest) Orderby(orderby []string) ApiSecurityListSecureScoreControlProfilesRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiSecurityListSecureScoreControlProfilesRequest) Select_(select_ []string) ApiSecurityListSecureScoreControlProfilesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiSecurityListSecureScoreControlProfilesRequest) Expand(expand []string) ApiSecurityListSecureScoreControlProfilesRequest {
	r.expand = &expand
	return r
}

func (r ApiSecurityListSecureScoreControlProfilesRequest) Execute() (CollectionOfSecureScoreControlProfile, *_nethttp.Response, error) {
	return r.ApiService.SecurityListSecureScoreControlProfilesExecute(r)
}

/*
SecurityListSecureScoreControlProfiles Get secureScoreControlProfiles from security

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSecurityListSecureScoreControlProfilesRequest
*/
func (a *SecuritySecureScoreControlProfileApiService) SecurityListSecureScoreControlProfiles(ctx _context.Context) ApiSecurityListSecureScoreControlProfilesRequest {
	return ApiSecurityListSecureScoreControlProfilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfSecureScoreControlProfile
func (a *SecuritySecureScoreControlProfileApiService) SecurityListSecureScoreControlProfilesExecute(r ApiSecurityListSecureScoreControlProfilesRequest) (CollectionOfSecureScoreControlProfile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfSecureScoreControlProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecuritySecureScoreControlProfileApiService.SecurityListSecureScoreControlProfiles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/secureScoreControlProfiles"

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

type ApiSecurityUpdateSecureScoreControlProfilesRequest struct {
	ctx _context.Context
	ApiService *SecuritySecureScoreControlProfileApiService
	secureScoreControlProfileId string
	microsoftGraphSecureScoreControlProfile *MicrosoftGraphSecureScoreControlProfile
}

// New navigation property values
func (r ApiSecurityUpdateSecureScoreControlProfilesRequest) MicrosoftGraphSecureScoreControlProfile(microsoftGraphSecureScoreControlProfile MicrosoftGraphSecureScoreControlProfile) ApiSecurityUpdateSecureScoreControlProfilesRequest {
	r.microsoftGraphSecureScoreControlProfile = &microsoftGraphSecureScoreControlProfile
	return r
}

func (r ApiSecurityUpdateSecureScoreControlProfilesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SecurityUpdateSecureScoreControlProfilesExecute(r)
}

/*
SecurityUpdateSecureScoreControlProfiles Update the navigation property secureScoreControlProfiles in security

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secureScoreControlProfileId key: id of secureScoreControlProfile
 @return ApiSecurityUpdateSecureScoreControlProfilesRequest
*/
func (a *SecuritySecureScoreControlProfileApiService) SecurityUpdateSecureScoreControlProfiles(ctx _context.Context, secureScoreControlProfileId string) ApiSecurityUpdateSecureScoreControlProfilesRequest {
	return ApiSecurityUpdateSecureScoreControlProfilesRequest{
		ApiService: a,
		ctx: ctx,
		secureScoreControlProfileId: secureScoreControlProfileId,
	}
}

// Execute executes the request
func (a *SecuritySecureScoreControlProfileApiService) SecurityUpdateSecureScoreControlProfilesExecute(r ApiSecurityUpdateSecureScoreControlProfilesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecuritySecureScoreControlProfileApiService.SecurityUpdateSecureScoreControlProfiles")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/security/secureScoreControlProfiles/{secureScoreControlProfile-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"secureScoreControlProfile-id"+"}", _neturl.PathEscape(parameterToString(r.secureScoreControlProfileId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSecureScoreControlProfile == nil {
		return nil, reportError("microsoftGraphSecureScoreControlProfile is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSecureScoreControlProfile
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