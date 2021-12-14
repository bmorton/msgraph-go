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

// AuditLogsSignInApiService AuditLogsSignInApi service
type AuditLogsSignInApiService service

type ApiAuditLogsCreateSignInsRequest struct {
	ctx _context.Context
	ApiService *AuditLogsSignInApiService
	microsoftGraphSignIn *MicrosoftGraphSignIn
}

// New navigation property
func (r ApiAuditLogsCreateSignInsRequest) MicrosoftGraphSignIn(microsoftGraphSignIn MicrosoftGraphSignIn) ApiAuditLogsCreateSignInsRequest {
	r.microsoftGraphSignIn = &microsoftGraphSignIn
	return r
}

func (r ApiAuditLogsCreateSignInsRequest) Execute() (MicrosoftGraphSignIn, *_nethttp.Response, error) {
	return r.ApiService.AuditLogsCreateSignInsExecute(r)
}

/*
AuditLogsCreateSignIns Create new navigation property to signIns for auditLogs

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditLogsCreateSignInsRequest
*/
func (a *AuditLogsSignInApiService) AuditLogsCreateSignIns(ctx _context.Context) ApiAuditLogsCreateSignInsRequest {
	return ApiAuditLogsCreateSignInsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSignIn
func (a *AuditLogsSignInApiService) AuditLogsCreateSignInsExecute(r ApiAuditLogsCreateSignInsRequest) (MicrosoftGraphSignIn, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSignIn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsSignInApiService.AuditLogsCreateSignIns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditLogs/signIns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSignIn == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphSignIn is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSignIn
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

type ApiAuditLogsDeleteSignInsRequest struct {
	ctx _context.Context
	ApiService *AuditLogsSignInApiService
	signInId string
	ifMatch *string
}

// ETag
func (r ApiAuditLogsDeleteSignInsRequest) IfMatch(ifMatch string) ApiAuditLogsDeleteSignInsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiAuditLogsDeleteSignInsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AuditLogsDeleteSignInsExecute(r)
}

/*
AuditLogsDeleteSignIns Delete navigation property signIns for auditLogs

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param signInId key: id of signIn
 @return ApiAuditLogsDeleteSignInsRequest
*/
func (a *AuditLogsSignInApiService) AuditLogsDeleteSignIns(ctx _context.Context, signInId string) ApiAuditLogsDeleteSignInsRequest {
	return ApiAuditLogsDeleteSignInsRequest{
		ApiService: a,
		ctx: ctx,
		signInId: signInId,
	}
}

// Execute executes the request
func (a *AuditLogsSignInApiService) AuditLogsDeleteSignInsExecute(r ApiAuditLogsDeleteSignInsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsSignInApiService.AuditLogsDeleteSignIns")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditLogs/signIns/{signIn-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"signIn-id"+"}", _neturl.PathEscape(parameterToString(r.signInId, "")), -1)

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

type ApiAuditLogsGetSignInsRequest struct {
	ctx _context.Context
	ApiService *AuditLogsSignInApiService
	signInId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiAuditLogsGetSignInsRequest) Select_(select_ []string) ApiAuditLogsGetSignInsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiAuditLogsGetSignInsRequest) Expand(expand []string) ApiAuditLogsGetSignInsRequest {
	r.expand = &expand
	return r
}

func (r ApiAuditLogsGetSignInsRequest) Execute() (MicrosoftGraphSignIn, *_nethttp.Response, error) {
	return r.ApiService.AuditLogsGetSignInsExecute(r)
}

/*
AuditLogsGetSignIns Get signIns from auditLogs

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param signInId key: id of signIn
 @return ApiAuditLogsGetSignInsRequest
*/
func (a *AuditLogsSignInApiService) AuditLogsGetSignIns(ctx _context.Context, signInId string) ApiAuditLogsGetSignInsRequest {
	return ApiAuditLogsGetSignInsRequest{
		ApiService: a,
		ctx: ctx,
		signInId: signInId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSignIn
func (a *AuditLogsSignInApiService) AuditLogsGetSignInsExecute(r ApiAuditLogsGetSignInsRequest) (MicrosoftGraphSignIn, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSignIn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsSignInApiService.AuditLogsGetSignIns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditLogs/signIns/{signIn-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"signIn-id"+"}", _neturl.PathEscape(parameterToString(r.signInId, "")), -1)

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

type ApiAuditLogsListSignInsRequest struct {
	ctx _context.Context
	ApiService *AuditLogsSignInApiService
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
func (r ApiAuditLogsListSignInsRequest) Top(top int32) ApiAuditLogsListSignInsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiAuditLogsListSignInsRequest) Skip(skip int32) ApiAuditLogsListSignInsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiAuditLogsListSignInsRequest) Search(search string) ApiAuditLogsListSignInsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiAuditLogsListSignInsRequest) Filter(filter string) ApiAuditLogsListSignInsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiAuditLogsListSignInsRequest) Count(count bool) ApiAuditLogsListSignInsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiAuditLogsListSignInsRequest) Orderby(orderby []string) ApiAuditLogsListSignInsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiAuditLogsListSignInsRequest) Select_(select_ []string) ApiAuditLogsListSignInsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiAuditLogsListSignInsRequest) Expand(expand []string) ApiAuditLogsListSignInsRequest {
	r.expand = &expand
	return r
}

func (r ApiAuditLogsListSignInsRequest) Execute() (CollectionOfSignIn, *_nethttp.Response, error) {
	return r.ApiService.AuditLogsListSignInsExecute(r)
}

/*
AuditLogsListSignIns Get signIns from auditLogs

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditLogsListSignInsRequest
*/
func (a *AuditLogsSignInApiService) AuditLogsListSignIns(ctx _context.Context) ApiAuditLogsListSignInsRequest {
	return ApiAuditLogsListSignInsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfSignIn
func (a *AuditLogsSignInApiService) AuditLogsListSignInsExecute(r ApiAuditLogsListSignInsRequest) (CollectionOfSignIn, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfSignIn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsSignInApiService.AuditLogsListSignIns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditLogs/signIns"

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

type ApiAuditLogsUpdateSignInsRequest struct {
	ctx _context.Context
	ApiService *AuditLogsSignInApiService
	signInId string
	microsoftGraphSignIn *MicrosoftGraphSignIn
}

// New navigation property values
func (r ApiAuditLogsUpdateSignInsRequest) MicrosoftGraphSignIn(microsoftGraphSignIn MicrosoftGraphSignIn) ApiAuditLogsUpdateSignInsRequest {
	r.microsoftGraphSignIn = &microsoftGraphSignIn
	return r
}

func (r ApiAuditLogsUpdateSignInsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AuditLogsUpdateSignInsExecute(r)
}

/*
AuditLogsUpdateSignIns Update the navigation property signIns in auditLogs

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param signInId key: id of signIn
 @return ApiAuditLogsUpdateSignInsRequest
*/
func (a *AuditLogsSignInApiService) AuditLogsUpdateSignIns(ctx _context.Context, signInId string) ApiAuditLogsUpdateSignInsRequest {
	return ApiAuditLogsUpdateSignInsRequest{
		ApiService: a,
		ctx: ctx,
		signInId: signInId,
	}
}

// Execute executes the request
func (a *AuditLogsSignInApiService) AuditLogsUpdateSignInsExecute(r ApiAuditLogsUpdateSignInsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsSignInApiService.AuditLogsUpdateSignIns")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditLogs/signIns/{signIn-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"signIn-id"+"}", _neturl.PathEscape(parameterToString(r.signInId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSignIn == nil {
		return nil, reportError("microsoftGraphSignIn is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSignIn
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
