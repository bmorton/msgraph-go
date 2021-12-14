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

// UsersOnlineMeetingApiService UsersOnlineMeetingApi service
type UsersOnlineMeetingApiService service

type ApiUsersCreateOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *UsersOnlineMeetingApiService
	userId string
	microsoftGraphOnlineMeeting *MicrosoftGraphOnlineMeeting
}

// New navigation property
func (r ApiUsersCreateOnlineMeetingsRequest) MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting MicrosoftGraphOnlineMeeting) ApiUsersCreateOnlineMeetingsRequest {
	r.microsoftGraphOnlineMeeting = &microsoftGraphOnlineMeeting
	return r
}

func (r ApiUsersCreateOnlineMeetingsRequest) Execute() (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	return r.ApiService.UsersCreateOnlineMeetingsExecute(r)
}

/*
UsersCreateOnlineMeetings Create new navigation property to onlineMeetings for users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersCreateOnlineMeetingsRequest
*/
func (a *UsersOnlineMeetingApiService) UsersCreateOnlineMeetings(ctx _context.Context, userId string) ApiUsersCreateOnlineMeetingsRequest {
	return ApiUsersCreateOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphOnlineMeeting
func (a *UsersOnlineMeetingApiService) UsersCreateOnlineMeetingsExecute(r ApiUsersCreateOnlineMeetingsRequest) (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphOnlineMeeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOnlineMeetingApiService.UsersCreateOnlineMeetings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/onlineMeetings"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiUsersDeleteOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *UsersOnlineMeetingApiService
	userId string
	onlineMeetingId string
	ifMatch *string
}

// ETag
func (r ApiUsersDeleteOnlineMeetingsRequest) IfMatch(ifMatch string) ApiUsersDeleteOnlineMeetingsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUsersDeleteOnlineMeetingsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersDeleteOnlineMeetingsExecute(r)
}

/*
UsersDeleteOnlineMeetings Delete navigation property onlineMeetings for users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiUsersDeleteOnlineMeetingsRequest
*/
func (a *UsersOnlineMeetingApiService) UsersDeleteOnlineMeetings(ctx _context.Context, userId string, onlineMeetingId string) ApiUsersDeleteOnlineMeetingsRequest {
	return ApiUsersDeleteOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
func (a *UsersOnlineMeetingApiService) UsersDeleteOnlineMeetingsExecute(r ApiUsersDeleteOnlineMeetingsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOnlineMeetingApiService.UsersDeleteOnlineMeetings")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/onlineMeetings/{onlineMeeting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiUsersGetOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *UsersOnlineMeetingApiService
	userId string
	onlineMeetingId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiUsersGetOnlineMeetingsRequest) Select_(select_ []string) ApiUsersGetOnlineMeetingsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersGetOnlineMeetingsRequest) Expand(expand []string) ApiUsersGetOnlineMeetingsRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersGetOnlineMeetingsRequest) Execute() (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	return r.ApiService.UsersGetOnlineMeetingsExecute(r)
}

/*
UsersGetOnlineMeetings Get onlineMeetings from users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiUsersGetOnlineMeetingsRequest
*/
func (a *UsersOnlineMeetingApiService) UsersGetOnlineMeetings(ctx _context.Context, userId string, onlineMeetingId string) ApiUsersGetOnlineMeetingsRequest {
	return ApiUsersGetOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphOnlineMeeting
func (a *UsersOnlineMeetingApiService) UsersGetOnlineMeetingsExecute(r ApiUsersGetOnlineMeetingsRequest) (MicrosoftGraphOnlineMeeting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphOnlineMeeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOnlineMeetingApiService.UsersGetOnlineMeetings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/onlineMeetings/{onlineMeeting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiUsersGetOnlineMeetingsAttendeeReportRequest struct {
	ctx _context.Context
	ApiService *UsersOnlineMeetingApiService
	userId string
	onlineMeetingId string
}


func (r ApiUsersGetOnlineMeetingsAttendeeReportRequest) Execute() (*os.File, *_nethttp.Response, error) {
	return r.ApiService.UsersGetOnlineMeetingsAttendeeReportExecute(r)
}

/*
UsersGetOnlineMeetingsAttendeeReport Get media content for the navigation property onlineMeetings from users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiUsersGetOnlineMeetingsAttendeeReportRequest
*/
func (a *UsersOnlineMeetingApiService) UsersGetOnlineMeetingsAttendeeReport(ctx _context.Context, userId string, onlineMeetingId string) ApiUsersGetOnlineMeetingsAttendeeReportRequest {
	return ApiUsersGetOnlineMeetingsAttendeeReportRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *UsersOnlineMeetingApiService) UsersGetOnlineMeetingsAttendeeReportExecute(r ApiUsersGetOnlineMeetingsAttendeeReportRequest) (*os.File, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOnlineMeetingApiService.UsersGetOnlineMeetingsAttendeeReport")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/onlineMeetings/{onlineMeeting-id}/attendeeReport"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiUsersListOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *UsersOnlineMeetingApiService
	userId string
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
func (r ApiUsersListOnlineMeetingsRequest) Top(top int32) ApiUsersListOnlineMeetingsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiUsersListOnlineMeetingsRequest) Skip(skip int32) ApiUsersListOnlineMeetingsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiUsersListOnlineMeetingsRequest) Search(search string) ApiUsersListOnlineMeetingsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiUsersListOnlineMeetingsRequest) Filter(filter string) ApiUsersListOnlineMeetingsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiUsersListOnlineMeetingsRequest) Count(count bool) ApiUsersListOnlineMeetingsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiUsersListOnlineMeetingsRequest) Orderby(orderby []string) ApiUsersListOnlineMeetingsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiUsersListOnlineMeetingsRequest) Select_(select_ []string) ApiUsersListOnlineMeetingsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersListOnlineMeetingsRequest) Expand(expand []string) ApiUsersListOnlineMeetingsRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersListOnlineMeetingsRequest) Execute() (CollectionOfOnlineMeeting, *_nethttp.Response, error) {
	return r.ApiService.UsersListOnlineMeetingsExecute(r)
}

/*
UsersListOnlineMeetings Get onlineMeetings from users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersListOnlineMeetingsRequest
*/
func (a *UsersOnlineMeetingApiService) UsersListOnlineMeetings(ctx _context.Context, userId string) ApiUsersListOnlineMeetingsRequest {
	return ApiUsersListOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfOnlineMeeting
func (a *UsersOnlineMeetingApiService) UsersListOnlineMeetingsExecute(r ApiUsersListOnlineMeetingsRequest) (CollectionOfOnlineMeeting, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfOnlineMeeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOnlineMeetingApiService.UsersListOnlineMeetings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/onlineMeetings"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiUsersUpdateOnlineMeetingsRequest struct {
	ctx _context.Context
	ApiService *UsersOnlineMeetingApiService
	userId string
	onlineMeetingId string
	microsoftGraphOnlineMeeting *MicrosoftGraphOnlineMeeting
}

// New navigation property values
func (r ApiUsersUpdateOnlineMeetingsRequest) MicrosoftGraphOnlineMeeting(microsoftGraphOnlineMeeting MicrosoftGraphOnlineMeeting) ApiUsersUpdateOnlineMeetingsRequest {
	r.microsoftGraphOnlineMeeting = &microsoftGraphOnlineMeeting
	return r
}

func (r ApiUsersUpdateOnlineMeetingsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersUpdateOnlineMeetingsExecute(r)
}

/*
UsersUpdateOnlineMeetings Update the navigation property onlineMeetings in users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiUsersUpdateOnlineMeetingsRequest
*/
func (a *UsersOnlineMeetingApiService) UsersUpdateOnlineMeetings(ctx _context.Context, userId string, onlineMeetingId string) ApiUsersUpdateOnlineMeetingsRequest {
	return ApiUsersUpdateOnlineMeetingsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
func (a *UsersOnlineMeetingApiService) UsersUpdateOnlineMeetingsExecute(r ApiUsersUpdateOnlineMeetingsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOnlineMeetingApiService.UsersUpdateOnlineMeetings")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/onlineMeetings/{onlineMeeting-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
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

type ApiUsersUpdateOnlineMeetingsAttendeeReportRequest struct {
	ctx _context.Context
	ApiService *UsersOnlineMeetingApiService
	userId string
	onlineMeetingId string
	body **os.File
}

// New media content.
func (r ApiUsersUpdateOnlineMeetingsAttendeeReportRequest) Body(body *os.File) ApiUsersUpdateOnlineMeetingsAttendeeReportRequest {
	r.body = &body
	return r
}

func (r ApiUsersUpdateOnlineMeetingsAttendeeReportRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersUpdateOnlineMeetingsAttendeeReportExecute(r)
}

/*
UsersUpdateOnlineMeetingsAttendeeReport Update media content for the navigation property onlineMeetings in users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param onlineMeetingId key: id of onlineMeeting
 @return ApiUsersUpdateOnlineMeetingsAttendeeReportRequest
*/
func (a *UsersOnlineMeetingApiService) UsersUpdateOnlineMeetingsAttendeeReport(ctx _context.Context, userId string, onlineMeetingId string) ApiUsersUpdateOnlineMeetingsAttendeeReportRequest {
	return ApiUsersUpdateOnlineMeetingsAttendeeReportRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		onlineMeetingId: onlineMeetingId,
	}
}

// Execute executes the request
func (a *UsersOnlineMeetingApiService) UsersUpdateOnlineMeetingsAttendeeReportExecute(r ApiUsersUpdateOnlineMeetingsAttendeeReportRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOnlineMeetingApiService.UsersUpdateOnlineMeetingsAttendeeReport")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/onlineMeetings/{onlineMeeting-id}/attendeeReport"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
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
