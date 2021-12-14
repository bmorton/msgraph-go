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

// MeDriveApiService MeDriveApi service
type MeDriveApiService service

type ApiMeCreateDrivesRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
	microsoftGraphDrive *MicrosoftGraphDrive
}

// New navigation property
func (r ApiMeCreateDrivesRequest) MicrosoftGraphDrive(microsoftGraphDrive MicrosoftGraphDrive) ApiMeCreateDrivesRequest {
	r.microsoftGraphDrive = &microsoftGraphDrive
	return r
}

func (r ApiMeCreateDrivesRequest) Execute() (MicrosoftGraphDrive, *_nethttp.Response, error) {
	return r.ApiService.MeCreateDrivesExecute(r)
}

/*
MeCreateDrives Create new navigation property to drives for me

A collection of drives available for this user. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeCreateDrivesRequest
*/
func (a *MeDriveApiService) MeCreateDrives(ctx _context.Context) ApiMeCreateDrivesRequest {
	return ApiMeCreateDrivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDrive
func (a *MeDriveApiService) MeCreateDrivesExecute(r ApiMeCreateDrivesRequest) (MicrosoftGraphDrive, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDrive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeCreateDrives")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drives"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDrive == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphDrive is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDrive
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

type ApiMeDeleteDriveRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
	ifMatch *string
}

// ETag
func (r ApiMeDeleteDriveRequest) IfMatch(ifMatch string) ApiMeDeleteDriveRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiMeDeleteDriveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeDeleteDriveExecute(r)
}

/*
MeDeleteDrive Delete navigation property drive for me

The user's OneDrive. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeDeleteDriveRequest
*/
func (a *MeDriveApiService) MeDeleteDrive(ctx _context.Context) ApiMeDeleteDriveRequest {
	return ApiMeDeleteDriveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MeDriveApiService) MeDeleteDriveExecute(r ApiMeDeleteDriveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeDeleteDrive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drive"

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

type ApiMeDeleteDrivesRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
	driveId string
	ifMatch *string
}

// ETag
func (r ApiMeDeleteDrivesRequest) IfMatch(ifMatch string) ApiMeDeleteDrivesRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiMeDeleteDrivesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeDeleteDrivesExecute(r)
}

/*
MeDeleteDrives Delete navigation property drives for me

A collection of drives available for this user. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @return ApiMeDeleteDrivesRequest
*/
func (a *MeDriveApiService) MeDeleteDrives(ctx _context.Context, driveId string) ApiMeDeleteDrivesRequest {
	return ApiMeDeleteDrivesRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
	}
}

// Execute executes the request
func (a *MeDriveApiService) MeDeleteDrivesExecute(r ApiMeDeleteDrivesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeDeleteDrives")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drives/{drive-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)

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

type ApiMeGetDriveRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiMeGetDriveRequest) Select_(select_ []string) ApiMeGetDriveRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiMeGetDriveRequest) Expand(expand []string) ApiMeGetDriveRequest {
	r.expand = &expand
	return r
}

func (r ApiMeGetDriveRequest) Execute() (MicrosoftGraphDrive, *_nethttp.Response, error) {
	return r.ApiService.MeGetDriveExecute(r)
}

/*
MeGetDrive Get drive from me

The user's OneDrive. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeGetDriveRequest
*/
func (a *MeDriveApiService) MeGetDrive(ctx _context.Context) ApiMeGetDriveRequest {
	return ApiMeGetDriveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDrive
func (a *MeDriveApiService) MeGetDriveExecute(r ApiMeGetDriveRequest) (MicrosoftGraphDrive, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDrive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeGetDrive")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drive"

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

type ApiMeGetDrivesRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
	driveId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiMeGetDrivesRequest) Select_(select_ []string) ApiMeGetDrivesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiMeGetDrivesRequest) Expand(expand []string) ApiMeGetDrivesRequest {
	r.expand = &expand
	return r
}

func (r ApiMeGetDrivesRequest) Execute() (MicrosoftGraphDrive, *_nethttp.Response, error) {
	return r.ApiService.MeGetDrivesExecute(r)
}

/*
MeGetDrives Get drives from me

A collection of drives available for this user. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @return ApiMeGetDrivesRequest
*/
func (a *MeDriveApiService) MeGetDrives(ctx _context.Context, driveId string) ApiMeGetDrivesRequest {
	return ApiMeGetDrivesRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDrive
func (a *MeDriveApiService) MeGetDrivesExecute(r ApiMeGetDrivesRequest) (MicrosoftGraphDrive, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDrive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeGetDrives")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drives/{drive-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)

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

type ApiMeListDrivesRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
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
func (r ApiMeListDrivesRequest) Top(top int32) ApiMeListDrivesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiMeListDrivesRequest) Skip(skip int32) ApiMeListDrivesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiMeListDrivesRequest) Search(search string) ApiMeListDrivesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiMeListDrivesRequest) Filter(filter string) ApiMeListDrivesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiMeListDrivesRequest) Count(count bool) ApiMeListDrivesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiMeListDrivesRequest) Orderby(orderby []string) ApiMeListDrivesRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiMeListDrivesRequest) Select_(select_ []string) ApiMeListDrivesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiMeListDrivesRequest) Expand(expand []string) ApiMeListDrivesRequest {
	r.expand = &expand
	return r
}

func (r ApiMeListDrivesRequest) Execute() (CollectionOfDrive, *_nethttp.Response, error) {
	return r.ApiService.MeListDrivesExecute(r)
}

/*
MeListDrives Get drives from me

A collection of drives available for this user. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeListDrivesRequest
*/
func (a *MeDriveApiService) MeListDrives(ctx _context.Context) ApiMeListDrivesRequest {
	return ApiMeListDrivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfDrive
func (a *MeDriveApiService) MeListDrivesExecute(r ApiMeListDrivesRequest) (CollectionOfDrive, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfDrive
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeListDrives")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drives"

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

type ApiMeUpdateDriveRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
	microsoftGraphDrive *MicrosoftGraphDrive
}

// New navigation property values
func (r ApiMeUpdateDriveRequest) MicrosoftGraphDrive(microsoftGraphDrive MicrosoftGraphDrive) ApiMeUpdateDriveRequest {
	r.microsoftGraphDrive = &microsoftGraphDrive
	return r
}

func (r ApiMeUpdateDriveRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeUpdateDriveExecute(r)
}

/*
MeUpdateDrive Update the navigation property drive in me

The user's OneDrive. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeUpdateDriveRequest
*/
func (a *MeDriveApiService) MeUpdateDrive(ctx _context.Context) ApiMeUpdateDriveRequest {
	return ApiMeUpdateDriveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MeDriveApiService) MeUpdateDriveExecute(r ApiMeUpdateDriveRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeUpdateDrive")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDrive == nil {
		return nil, reportError("microsoftGraphDrive is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDrive
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

type ApiMeUpdateDrivesRequest struct {
	ctx _context.Context
	ApiService *MeDriveApiService
	driveId string
	microsoftGraphDrive *MicrosoftGraphDrive
}

// New navigation property values
func (r ApiMeUpdateDrivesRequest) MicrosoftGraphDrive(microsoftGraphDrive MicrosoftGraphDrive) ApiMeUpdateDrivesRequest {
	r.microsoftGraphDrive = &microsoftGraphDrive
	return r
}

func (r ApiMeUpdateDrivesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeUpdateDrivesExecute(r)
}

/*
MeUpdateDrives Update the navigation property drives in me

A collection of drives available for this user. Read-only.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @return ApiMeUpdateDrivesRequest
*/
func (a *MeDriveApiService) MeUpdateDrives(ctx _context.Context, driveId string) ApiMeUpdateDrivesRequest {
	return ApiMeUpdateDrivesRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
	}
}

// Execute executes the request
func (a *MeDriveApiService) MeUpdateDrivesExecute(r ApiMeUpdateDrivesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDriveApiService.MeUpdateDrives")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drives/{drive-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDrive == nil {
		return nil, reportError("microsoftGraphDrive is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDrive
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
