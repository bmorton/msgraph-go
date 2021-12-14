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

// DeviceManagementDeviceCategoryApiService DeviceManagementDeviceCategoryApi service
type DeviceManagementDeviceCategoryApiService service

type ApiDeviceManagementCreateDeviceCategoriesRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCategoryApiService
	microsoftGraphDeviceCategory *MicrosoftGraphDeviceCategory
}

// New navigation property
func (r ApiDeviceManagementCreateDeviceCategoriesRequest) MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory MicrosoftGraphDeviceCategory) ApiDeviceManagementCreateDeviceCategoriesRequest {
	r.microsoftGraphDeviceCategory = &microsoftGraphDeviceCategory
	return r
}

func (r ApiDeviceManagementCreateDeviceCategoriesRequest) Execute() (MicrosoftGraphDeviceCategory, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementCreateDeviceCategoriesExecute(r)
}

/*
DeviceManagementCreateDeviceCategories Create new navigation property to deviceCategories for deviceManagement

The list of device categories with the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementCreateDeviceCategoriesRequest
*/
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementCreateDeviceCategories(ctx _context.Context) ApiDeviceManagementCreateDeviceCategoriesRequest {
	return ApiDeviceManagementCreateDeviceCategoriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDeviceCategory
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementCreateDeviceCategoriesExecute(r ApiDeviceManagementCreateDeviceCategoriesRequest) (MicrosoftGraphDeviceCategory, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDeviceCategory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCategoryApiService.DeviceManagementCreateDeviceCategories")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCategories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDeviceCategory == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphDeviceCategory is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDeviceCategory
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

type ApiDeviceManagementDeleteDeviceCategoriesRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCategoryApiService
	deviceCategoryId string
	ifMatch *string
}

// ETag
func (r ApiDeviceManagementDeleteDeviceCategoriesRequest) IfMatch(ifMatch string) ApiDeviceManagementDeleteDeviceCategoriesRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeviceManagementDeleteDeviceCategoriesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementDeleteDeviceCategoriesExecute(r)
}

/*
DeviceManagementDeleteDeviceCategories Delete navigation property deviceCategories for deviceManagement

The list of device categories with the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceCategoryId key: id of deviceCategory
 @return ApiDeviceManagementDeleteDeviceCategoriesRequest
*/
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementDeleteDeviceCategories(ctx _context.Context, deviceCategoryId string) ApiDeviceManagementDeleteDeviceCategoriesRequest {
	return ApiDeviceManagementDeleteDeviceCategoriesRequest{
		ApiService: a,
		ctx: ctx,
		deviceCategoryId: deviceCategoryId,
	}
}

// Execute executes the request
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementDeleteDeviceCategoriesExecute(r ApiDeviceManagementDeleteDeviceCategoriesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCategoryApiService.DeviceManagementDeleteDeviceCategories")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCategories/{deviceCategory-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceCategory-id"+"}", _neturl.PathEscape(parameterToString(r.deviceCategoryId, "")), -1)

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

type ApiDeviceManagementGetDeviceCategoriesRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCategoryApiService
	deviceCategoryId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiDeviceManagementGetDeviceCategoriesRequest) Select_(select_ []string) ApiDeviceManagementGetDeviceCategoriesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementGetDeviceCategoriesRequest) Expand(expand []string) ApiDeviceManagementGetDeviceCategoriesRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementGetDeviceCategoriesRequest) Execute() (MicrosoftGraphDeviceCategory, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementGetDeviceCategoriesExecute(r)
}

/*
DeviceManagementGetDeviceCategories Get deviceCategories from deviceManagement

The list of device categories with the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceCategoryId key: id of deviceCategory
 @return ApiDeviceManagementGetDeviceCategoriesRequest
*/
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementGetDeviceCategories(ctx _context.Context, deviceCategoryId string) ApiDeviceManagementGetDeviceCategoriesRequest {
	return ApiDeviceManagementGetDeviceCategoriesRequest{
		ApiService: a,
		ctx: ctx,
		deviceCategoryId: deviceCategoryId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphDeviceCategory
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementGetDeviceCategoriesExecute(r ApiDeviceManagementGetDeviceCategoriesRequest) (MicrosoftGraphDeviceCategory, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphDeviceCategory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCategoryApiService.DeviceManagementGetDeviceCategories")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCategories/{deviceCategory-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceCategory-id"+"}", _neturl.PathEscape(parameterToString(r.deviceCategoryId, "")), -1)

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

type ApiDeviceManagementListDeviceCategoriesRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCategoryApiService
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
func (r ApiDeviceManagementListDeviceCategoriesRequest) Top(top int32) ApiDeviceManagementListDeviceCategoriesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDeviceManagementListDeviceCategoriesRequest) Skip(skip int32) ApiDeviceManagementListDeviceCategoriesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDeviceManagementListDeviceCategoriesRequest) Search(search string) ApiDeviceManagementListDeviceCategoriesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDeviceManagementListDeviceCategoriesRequest) Filter(filter string) ApiDeviceManagementListDeviceCategoriesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDeviceManagementListDeviceCategoriesRequest) Count(count bool) ApiDeviceManagementListDeviceCategoriesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDeviceManagementListDeviceCategoriesRequest) Orderby(orderby []string) ApiDeviceManagementListDeviceCategoriesRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDeviceManagementListDeviceCategoriesRequest) Select_(select_ []string) ApiDeviceManagementListDeviceCategoriesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementListDeviceCategoriesRequest) Expand(expand []string) ApiDeviceManagementListDeviceCategoriesRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementListDeviceCategoriesRequest) Execute() (CollectionOfDeviceCategory, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementListDeviceCategoriesExecute(r)
}

/*
DeviceManagementListDeviceCategories Get deviceCategories from deviceManagement

The list of device categories with the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementListDeviceCategoriesRequest
*/
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementListDeviceCategories(ctx _context.Context) ApiDeviceManagementListDeviceCategoriesRequest {
	return ApiDeviceManagementListDeviceCategoriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfDeviceCategory
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementListDeviceCategoriesExecute(r ApiDeviceManagementListDeviceCategoriesRequest) (CollectionOfDeviceCategory, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfDeviceCategory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCategoryApiService.DeviceManagementListDeviceCategories")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCategories"

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

type ApiDeviceManagementUpdateDeviceCategoriesRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementDeviceCategoryApiService
	deviceCategoryId string
	microsoftGraphDeviceCategory *MicrosoftGraphDeviceCategory
}

// New navigation property values
func (r ApiDeviceManagementUpdateDeviceCategoriesRequest) MicrosoftGraphDeviceCategory(microsoftGraphDeviceCategory MicrosoftGraphDeviceCategory) ApiDeviceManagementUpdateDeviceCategoriesRequest {
	r.microsoftGraphDeviceCategory = &microsoftGraphDeviceCategory
	return r
}

func (r ApiDeviceManagementUpdateDeviceCategoriesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementUpdateDeviceCategoriesExecute(r)
}

/*
DeviceManagementUpdateDeviceCategories Update the navigation property deviceCategories in deviceManagement

The list of device categories with the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceCategoryId key: id of deviceCategory
 @return ApiDeviceManagementUpdateDeviceCategoriesRequest
*/
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementUpdateDeviceCategories(ctx _context.Context, deviceCategoryId string) ApiDeviceManagementUpdateDeviceCategoriesRequest {
	return ApiDeviceManagementUpdateDeviceCategoriesRequest{
		ApiService: a,
		ctx: ctx,
		deviceCategoryId: deviceCategoryId,
	}
}

// Execute executes the request
func (a *DeviceManagementDeviceCategoryApiService) DeviceManagementUpdateDeviceCategoriesExecute(r ApiDeviceManagementUpdateDeviceCategoriesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementDeviceCategoryApiService.DeviceManagementUpdateDeviceCategories")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/deviceCategories/{deviceCategory-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceCategory-id"+"}", _neturl.PathEscape(parameterToString(r.deviceCategoryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphDeviceCategory == nil {
		return nil, reportError("microsoftGraphDeviceCategory is required and must be specified")
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
	localVarPostBody = r.microsoftGraphDeviceCategory
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