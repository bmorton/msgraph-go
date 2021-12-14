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
	"os"
)

// Linger please
var (
	_ _context.Context
)

// CommunicationsOnlineMeetingApiService CommunicationsOnlineMeetingApi service
type CommunicationsOnlineMeetingApiService service

type ApiCommunicationsCreateOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *CommunicationsOnlineMeetingApiService
	microsoftGraphOnlineMeeting *MicrosoftGraphOnlineMeeting
}

// New navigation property
func (r ApiCommunicationsCreateOnlineMeetingsRequest) MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting MicrosoftGraphOnlineMeeting) ApiCommunicationsCreateOnlineMeetingsRequest {
	r.microsoftGraphOnlineMeeting = &microsoftGraphOnlineMeeting
	return r
}

func (r ApiCommunicationsCreateOnlineMeetingsRequest) Execute() (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	return r.ApiService.CommunicationsCreateOnlineMeetingsExecute(r)
}

/*
CommunicationsCreateOnlineMeetings Create new navigation property to onlineMeetings for communications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCommunicationsCreateOnlineMeetingsRequest
*/
func (a *CommunicationsOnlineMeetingApiService) CommunicationsCreateOnlineMeetings(ctx _context.Context) ApiCommunicationsCreateOnlineMeetingsRequest {
	return ApiCommunicationsCreateOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphOnlineMeeting
func (a *CommunicationsOnlineMeetingApiService) CommunicationsCreateOnlineMeetingsExecute(r ApiCommunicationsCreateOnlineMeetingsRequest) (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphOnlineMeeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationsOnlineMeetingApiService.CommunicationsCreateOnlineMeetings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communications/onlineMeetings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphOnlineMeeting == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphOnlineMeeting is required and must be specified")
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
	localVarPostBody = r.microsoftGraphOnlineMeeting
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

type ApiCommunicationsDeleteOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *CommunicationsOnlineMeetingApiService
	onlineMeetingId string
	ifMatch *string
}

// ETag
func (r ApiCommunicationsDeleteOnlineMeetingsRequest) IfMatch(ifMatch string) ApiCommunicationsDeleteOnlineMeetingsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiCommunicationsDeleteOnlineMeetingsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CommunicationsDeleteOnlineMeetingsExecute(r)
}

/*
CommunicationsDeleteOnlineMeetings Delete navigation property onlineMeetings for communications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiCommunicationsDeleteOnlineMeetingsRequest
*/
func (a *CommunicationsOnlineMeetingApiService) CommunicationsDeleteOnlineMeetings(ctx _context.Context, onlineMeetingId string) ApiCommunicationsDeleteOnlineMeetingsRequest {
	return ApiCommunicationsDeleteOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
func (a *CommunicationsOnlineMeetingApiService) CommunicationsDeleteOnlineMeetingsExecute(r ApiCommunicationsDeleteOnlineMeetingsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationsOnlineMeetingApiService.CommunicationsDeleteOnlineMeetings")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communications/onlineMeetings/{onlineMeeting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"onlineMeeting-id"+"}", _neturl.PathEscape(parameterToString(r.onlineMeetingId, "")), -1)

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

type ApiCommunicationsGetOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *CommunicationsOnlineMeetingApiService
	onlineMeetingId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiCommunicationsGetOnlineMeetingsRequest) Select_(select_ []string) ApiCommunicationsGetOnlineMeetingsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiCommunicationsGetOnlineMeetingsRequest) Expand(expand []string) ApiCommunicationsGetOnlineMeetingsRequest {
	r.expand = &expand
	return r
}

func (r ApiCommunicationsGetOnlineMeetingsRequest) Execute() (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	return r.ApiService.CommunicationsGetOnlineMeetingsExecute(r)
}

/*
CommunicationsGetOnlineMeetings Get onlineMeetings from communications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiCommunicationsGetOnlineMeetingsRequest
*/
func (a *CommunicationsOnlineMeetingApiService) CommunicationsGetOnlineMeetings(ctx _context.Context, onlineMeetingId string) ApiCommunicationsGetOnlineMeetingsRequest {
	return ApiCommunicationsGetOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphOnlineMeeting
func (a *CommunicationsOnlineMeetingApiService) CommunicationsGetOnlineMeetingsExecute(r ApiCommunicationsGetOnlineMeetingsRequest) (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphOnlineMeeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationsOnlineMeetingApiService.CommunicationsGetOnlineMeetings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communications/onlineMeetings/{onlineMeeting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"onlineMeeting-id"+"}", _neturl.PathEscape(parameterToString(r.onlineMeetingId, "")), -1)

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

type ApiCommunicationsGetOnlineMeetingsAttendeeReportRequest struct {
	ctx _context.Context
	ApiService *CommunicationsOnlineMeetingApiService
	onlineMeetingId string
}


func (r ApiCommunicationsGetOnlineMeetingsAttendeeReportRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.CommunicationsGetOnlineMeetingsAttendeeReportExecute(r)
}

/*
CommunicationsGetOnlineMeetingsAttendeeReport Get media content for the navigation property onlineMeetings from communications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiCommunicationsGetOnlineMeetingsAttendeeReportRequest
*/
func (a *CommunicationsOnlineMeetingApiService) CommunicationsGetOnlineMeetingsAttendeeReport(ctx _context.Context, onlineMeetingId string) ApiCommunicationsGetOnlineMeetingsAttendeeReportRequest {
	return ApiCommunicationsGetOnlineMeetingsAttendeeReportRequest{
		ApiService: a,
		ctx: ctx,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *CommunicationsOnlineMeetingApiService) CommunicationsGetOnlineMeetingsAttendeeReportExecute(r ApiCommunicationsGetOnlineMeetingsAttendeeReportRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationsOnlineMeetingApiService.CommunicationsGetOnlineMeetingsAttendeeReport")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communications/onlineMeetings/{onlineMeeting-id}/attendeeReport"
	localVarPath = strings.Replace(localVarPath, "{"+"onlineMeeting-id"+"}", _neturl.PathEscape(parameterToString(r.onlineMeetingId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

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

type ApiCommunicationsListOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *CommunicationsOnlineMeetingApiService
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
func (r ApiCommunicationsListOnlineMeetingsRequest) Top(top int32) ApiCommunicationsListOnlineMeetingsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiCommunicationsListOnlineMeetingsRequest) Skip(skip int32) ApiCommunicationsListOnlineMeetingsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiCommunicationsListOnlineMeetingsRequest) Search(search string) ApiCommunicationsListOnlineMeetingsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiCommunicationsListOnlineMeetingsRequest) Filter(filter string) ApiCommunicationsListOnlineMeetingsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiCommunicationsListOnlineMeetingsRequest) Count(count bool) ApiCommunicationsListOnlineMeetingsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiCommunicationsListOnlineMeetingsRequest) Orderby(orderby []string) ApiCommunicationsListOnlineMeetingsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiCommunicationsListOnlineMeetingsRequest) Select_(select_ []string) ApiCommunicationsListOnlineMeetingsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiCommunicationsListOnlineMeetingsRequest) Expand(expand []string) ApiCommunicationsListOnlineMeetingsRequest {
	r.expand = &expand
	return r
}

func (r ApiCommunicationsListOnlineMeetingsRequest) Execute() (CollectionOfOnlineMeeting, *_nethttp.Response, error) {
	return r.ApiService.CommunicationsListOnlineMeetingsExecute(r)
}

/*
CommunicationsListOnlineMeetings Get onlineMeetings from communications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCommunicationsListOnlineMeetingsRequest
*/
func (a *CommunicationsOnlineMeetingApiService) CommunicationsListOnlineMeetings(ctx _context.Context) ApiCommunicationsListOnlineMeetingsRequest {
	return ApiCommunicationsListOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfOnlineMeeting
func (a *CommunicationsOnlineMeetingApiService) CommunicationsListOnlineMeetingsExecute(r ApiCommunicationsListOnlineMeetingsRequest) (CollectionOfOnlineMeeting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfOnlineMeeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationsOnlineMeetingApiService.CommunicationsListOnlineMeetings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communications/onlineMeetings"

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

type ApiCommunicationsUpdateOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *CommunicationsOnlineMeetingApiService
	onlineMeetingId string
	microsoftGraphOnlineMeeting *MicrosoftGraphOnlineMeeting
}

// New navigation property values
func (r ApiCommunicationsUpdateOnlineMeetingsRequest) MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting MicrosoftGraphOnlineMeeting) ApiCommunicationsUpdateOnlineMeetingsRequest {
	r.microsoftGraphOnlineMeeting = &microsoftGraphOnlineMeeting
	return r
}

func (r ApiCommunicationsUpdateOnlineMeetingsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CommunicationsUpdateOnlineMeetingsExecute(r)
}

/*
CommunicationsUpdateOnlineMeetings Update the navigation property onlineMeetings in communications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiCommunicationsUpdateOnlineMeetingsRequest
*/
func (a *CommunicationsOnlineMeetingApiService) CommunicationsUpdateOnlineMeetings(ctx _context.Context, onlineMeetingId string) ApiCommunicationsUpdateOnlineMeetingsRequest {
	return ApiCommunicationsUpdateOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
func (a *CommunicationsOnlineMeetingApiService) CommunicationsUpdateOnlineMeetingsExecute(r ApiCommunicationsUpdateOnlineMeetingsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationsOnlineMeetingApiService.CommunicationsUpdateOnlineMeetings")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communications/onlineMeetings/{onlineMeeting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"onlineMeeting-id"+"}", _neturl.PathEscape(parameterToString(r.onlineMeetingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphOnlineMeeting == nil {
		return nil, reportError("microsoftGraphOnlineMeeting is required and must be specified")
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
	localVarPostBody = r.microsoftGraphOnlineMeeting
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

type ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest struct {
	ctx _context.Context
	ApiService *CommunicationsOnlineMeetingApiService
	onlineMeetingId string
	body **os.File
}

// New media content.
func (r ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest) Body(body *os.File) ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest {
	r.body = &body
	return r
}

func (r ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CommunicationsUpdateOnlineMeetingsAttendeeReportExecute(r)
}

/*
CommunicationsUpdateOnlineMeetingsAttendeeReport Update media content for the navigation property onlineMeetings in communications

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest
*/
func (a *CommunicationsOnlineMeetingApiService) CommunicationsUpdateOnlineMeetingsAttendeeReport(ctx _context.Context, onlineMeetingId string) ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest {
	return ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest{
		ApiService: a,
		ctx: ctx,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
func (a *CommunicationsOnlineMeetingApiService) CommunicationsUpdateOnlineMeetingsAttendeeReportExecute(r ApiCommunicationsUpdateOnlineMeetingsAttendeeReportRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommunicationsOnlineMeetingApiService.CommunicationsUpdateOnlineMeetingsAttendeeReport")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/communications/onlineMeetings/{onlineMeeting-id}/attendeeReport"
	localVarPath = strings.Replace(localVarPath, "{"+"onlineMeeting-id"+"}", _neturl.PathEscape(parameterToString(r.onlineMeetingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

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
	localVarPostBody = r.body
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