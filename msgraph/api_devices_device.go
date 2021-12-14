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

// DevicesDeviceApiService DevicesDeviceApi service
type DevicesDeviceApiService service

type ApiDevicesDeviceCreateDeviceRequest struct {
	ctx _context.Context
	ApiService *DevicesDeviceApiService
	microsoftGraphDevice *MicrosoftGraphDevice
}

// New entity
func (r ApiDevicesDeviceCreateDeviceRequest) MicrosoftGraphDevice(microsoftGraphDevice MicrosoftGraphDevice) ApiDevicesDeviceCreateDeviceRequest {
	r.microsoftGraphDevice = &microsoftGraphDevice
	return r
}

func (r ApiDevicesDeviceCreateDeviceRequest) Execute() (MicrosoftGraphDevice, *_nethttp.Response, error) {
	return r.ApiService.DevicesDeviceCreateDeviceExecute(r)
}

/*
DevicesDeviceCreateDevice Add new entity to devices

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDevicesDeviceCreateDeviceRequest
*/
func (a *DevicesDeviceApiService) DevicesDeviceCreateDevice(ctx _context.Context) ApiDevicesDeviceCreateDeviceRequest {
	return ApiDevicesDeviceCreateDeviceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDevice
func (a *DevicesDeviceApiService) DevicesDeviceCreateDeviceExecute(r ApiDevicesDeviceCreateDeviceRequest) (MicrosoftGraphDevice, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDevice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevicesDeviceApiService.DevicesDeviceCreateDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDevice == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphDevice is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDevice
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

type ApiDevicesDeviceDeleteDeviceRequest struct {
	ctx _context.Context
	ApiService *DevicesDeviceApiService
	deviceId string
	ifMatch *string
}

// ETag
func (r ApiDevicesDeviceDeleteDeviceRequest) IfMatch(ifMatch string) ApiDevicesDeviceDeleteDeviceRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDevicesDeviceDeleteDeviceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DevicesDeviceDeleteDeviceExecute(r)
}

/*
DevicesDeviceDeleteDevice Delete entity from devices

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId key: id of device
 @return ApiDevicesDeviceDeleteDeviceRequest
*/
func (a *DevicesDeviceApiService) DevicesDeviceDeleteDevice(ctx _context.Context, deviceId string) ApiDevicesDeviceDeleteDeviceRequest {
	return ApiDevicesDeviceDeleteDeviceRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
func (a *DevicesDeviceApiService) DevicesDeviceDeleteDeviceExecute(r ApiDevicesDeviceDeleteDeviceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevicesDeviceApiService.DevicesDeviceDeleteDevice")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{device-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"device-id"+"}", _neturl.PathEscape(parameterToString(r.deviceId, "")), -1)

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

type ApiDevicesDeviceGetDeviceRequest struct {
	ctx _context.Context
	ApiService *DevicesDeviceApiService
	deviceId string
	consistencyLevel *string
	select_ *[]string
	expand *[]string
}

// Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/
func (r ApiDevicesDeviceGetDeviceRequest) ConsistencyLevel(consistencyLevel string) ApiDevicesDeviceGetDeviceRequest {
	r.consistencyLevel = &consistencyLevel
	return r
}
// Select properties to be returned
func (r ApiDevicesDeviceGetDeviceRequest) Select_(select_ []string) ApiDevicesDeviceGetDeviceRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDevicesDeviceGetDeviceRequest) Expand(expand []string) ApiDevicesDeviceGetDeviceRequest {
	r.expand = &expand
	return r
}

func (r ApiDevicesDeviceGetDeviceRequest) Execute() (MicrosoftGraphDevice, *_nethttp.Response, error) {
	return r.ApiService.DevicesDeviceGetDeviceExecute(r)
}

/*
DevicesDeviceGetDevice Get entity from devices by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId key: id of device
 @return ApiDevicesDeviceGetDeviceRequest
*/
func (a *DevicesDeviceApiService) DevicesDeviceGetDevice(ctx _context.Context, deviceId string) ApiDevicesDeviceGetDeviceRequest {
	return ApiDevicesDeviceGetDeviceRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDevice
func (a *DevicesDeviceApiService) DevicesDeviceGetDeviceExecute(r ApiDevicesDeviceGetDeviceRequest) (MicrosoftGraphDevice, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDevice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevicesDeviceApiService.DevicesDeviceGetDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{device-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"device-id"+"}", _neturl.PathEscape(parameterToString(r.deviceId, "")), -1)

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

type ApiDevicesDeviceListDeviceRequest struct {
	ctx _context.Context
	ApiService *DevicesDeviceApiService
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
func (r ApiDevicesDeviceListDeviceRequest) ConsistencyLevel(consistencyLevel string) ApiDevicesDeviceListDeviceRequest {
	r.consistencyLevel = &consistencyLevel
	return r
}
// Show only the first n items
func (r ApiDevicesDeviceListDeviceRequest) Top(top int32) ApiDevicesDeviceListDeviceRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDevicesDeviceListDeviceRequest) Skip(skip int32) ApiDevicesDeviceListDeviceRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDevicesDeviceListDeviceRequest) Search(search string) ApiDevicesDeviceListDeviceRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDevicesDeviceListDeviceRequest) Filter(filter string) ApiDevicesDeviceListDeviceRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDevicesDeviceListDeviceRequest) Count(count bool) ApiDevicesDeviceListDeviceRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDevicesDeviceListDeviceRequest) Orderby(orderby []string) ApiDevicesDeviceListDeviceRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDevicesDeviceListDeviceRequest) Select_(select_ []string) ApiDevicesDeviceListDeviceRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDevicesDeviceListDeviceRequest) Expand(expand []string) ApiDevicesDeviceListDeviceRequest {
	r.expand = &expand
	return r
}

func (r ApiDevicesDeviceListDeviceRequest) Execute() (CollectionOfDevice, *_nethttp.Response, error) {
	return r.ApiService.DevicesDeviceListDeviceExecute(r)
}

/*
DevicesDeviceListDevice Get entities from devices

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDevicesDeviceListDeviceRequest
*/
func (a *DevicesDeviceApiService) DevicesDeviceListDevice(ctx _context.Context) ApiDevicesDeviceListDeviceRequest {
	return ApiDevicesDeviceListDeviceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfDevice
func (a *DevicesDeviceApiService) DevicesDeviceListDeviceExecute(r ApiDevicesDeviceListDeviceRequest) (CollectionOfDevice, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfDevice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevicesDeviceApiService.DevicesDeviceListDevice")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices"

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

type ApiDevicesDeviceUpdateDeviceRequest struct {
	ctx _context.Context
	ApiService *DevicesDeviceApiService
	deviceId string
	microsoftGraphDevice *MicrosoftGraphDevice
}

// New property values
func (r ApiDevicesDeviceUpdateDeviceRequest) MicrosoftGraphDevice(microsoftGraphDevice MicrosoftGraphDevice) ApiDevicesDeviceUpdateDeviceRequest {
	r.microsoftGraphDevice = &microsoftGraphDevice
	return r
}

func (r ApiDevicesDeviceUpdateDeviceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DevicesDeviceUpdateDeviceExecute(r)
}

/*
DevicesDeviceUpdateDevice Update entity in devices

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId key: id of device
 @return ApiDevicesDeviceUpdateDeviceRequest
*/
func (a *DevicesDeviceApiService) DevicesDeviceUpdateDevice(ctx _context.Context, deviceId string) ApiDevicesDeviceUpdateDeviceRequest {
	return ApiDevicesDeviceUpdateDeviceRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
func (a *DevicesDeviceApiService) DevicesDeviceUpdateDeviceExecute(r ApiDevicesDeviceUpdateDeviceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevicesDeviceApiService.DevicesDeviceUpdateDevice")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{device-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"device-id"+"}", _neturl.PathEscape(parameterToString(r.deviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDevice == nil {
		return nil, reportError("microsoftGraphDevice is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDevice
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
