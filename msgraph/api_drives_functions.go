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

// DrivesFunctionsApiService DrivesFunctionsApi service
type DrivesFunctionsApiService service

type ApiDrivesDriveListContentTypesContentTypeBaseIsPublishedRequest struct {
	ctx _context.Context
	ApiService *DrivesFunctionsApiService
	driveId string
	contentTypeId string
}


func (r ApiDrivesDriveListContentTypesContentTypeBaseIsPublishedRequest) Execute() (bool, *_nethttp.Response, error) {
	return r.ApiService.DrivesDriveListContentTypesContentTypeBaseIsPublishedExecute(r)
}

/*
DrivesDriveListContentTypesContentTypeBaseIsPublished Invoke function isPublished

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param contentTypeId key: id of contentType
 @return ApiDrivesDriveListContentTypesContentTypeBaseIsPublishedRequest
*/
func (a *DrivesFunctionsApiService) DrivesDriveListContentTypesContentTypeBaseIsPublished(ctx _context.Context, driveId string, contentTypeId string) ApiDrivesDriveListContentTypesContentTypeBaseIsPublishedRequest {
	return ApiDrivesDriveListContentTypesContentTypeBaseIsPublishedRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
		contentTypeId: contentTypeId,
	}
}

// Execute executes the request
//  @return bool
func (a *DrivesFunctionsApiService) DrivesDriveListContentTypesContentTypeBaseIsPublishedExecute(r ApiDrivesDriveListContentTypesContentTypeBaseIsPublishedRequest) (bool, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesFunctionsApiService.DrivesDriveListContentTypesContentTypeBaseIsPublished")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives/{drive-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.isPublished()"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contentType-id"+"}", _neturl.PathEscape(parameterToString(r.contentTypeId, "")), -1)

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

type ApiDrivesDriveListContentTypesContentTypeIsPublishedRequest struct {
	ctx _context.Context
	ApiService *DrivesFunctionsApiService
	driveId string
	contentTypeId string
}


func (r ApiDrivesDriveListContentTypesContentTypeIsPublishedRequest) Execute() (bool, *_nethttp.Response, error) {
	return r.ApiService.DrivesDriveListContentTypesContentTypeIsPublishedExecute(r)
}

/*
DrivesDriveListContentTypesContentTypeIsPublished Invoke function isPublished

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param contentTypeId key: id of contentType
 @return ApiDrivesDriveListContentTypesContentTypeIsPublishedRequest
*/
func (a *DrivesFunctionsApiService) DrivesDriveListContentTypesContentTypeIsPublished(ctx _context.Context, driveId string, contentTypeId string) ApiDrivesDriveListContentTypesContentTypeIsPublishedRequest {
	return ApiDrivesDriveListContentTypesContentTypeIsPublishedRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
		contentTypeId: contentTypeId,
	}
}

// Execute executes the request
//  @return bool
func (a *DrivesFunctionsApiService) DrivesDriveListContentTypesContentTypeIsPublishedExecute(r ApiDrivesDriveListContentTypesContentTypeIsPublishedRequest) (bool, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesFunctionsApiService.DrivesDriveListContentTypesContentTypeIsPublished")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives/{drive-id}/list/contentTypes/{contentType-id}/microsoft.graph.isPublished()"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contentType-id"+"}", _neturl.PathEscape(parameterToString(r.contentTypeId, "")), -1)

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

type ApiDrivesDriveListItemsListItemGetActivitiesByInterval53eeRequest struct {
	ctx _context.Context
	ApiService *DrivesFunctionsApiService
	driveId string
	listItemId string
	startDateTime string
	endDateTime string
	interval string
}


func (r ApiDrivesDriveListItemsListItemGetActivitiesByInterval53eeRequest) Execute() ([]*AnyOfmicrosoftGraphItemActivityStat, *_nethttp.Response, error) {
	return r.ApiService.DrivesDriveListItemsListItemGetActivitiesByInterval53eeExecute(r)
}

/*
DrivesDriveListItemsListItemGetActivitiesByInterval53ee Invoke function getActivitiesByInterval

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param listItemId key: id of listItem
 @param startDateTime Usage: startDateTime={startDateTime}
 @param endDateTime Usage: endDateTime={endDateTime}
 @param interval Usage: interval={interval}
 @return ApiDrivesDriveListItemsListItemGetActivitiesByInterval53eeRequest
*/
func (a *DrivesFunctionsApiService) DrivesDriveListItemsListItemGetActivitiesByInterval53ee(ctx _context.Context, driveId string, listItemId string, startDateTime string, endDateTime string, interval string) ApiDrivesDriveListItemsListItemGetActivitiesByInterval53eeRequest {
	return ApiDrivesDriveListItemsListItemGetActivitiesByInterval53eeRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
		listItemId: listItemId,
		startDateTime: startDateTime,
		endDateTime: endDateTime,
		interval: interval,
	}
}

// Execute executes the request
//  @return []*AnyOfmicrosoftGraphItemActivityStat
func (a *DrivesFunctionsApiService) DrivesDriveListItemsListItemGetActivitiesByInterval53eeExecute(r ApiDrivesDriveListItemsListItemGetActivitiesByInterval53eeRequest) ([]*AnyOfmicrosoftGraphItemActivityStat, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []*AnyOfmicrosoftGraphItemActivityStat
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesFunctionsApiService.DrivesDriveListItemsListItemGetActivitiesByInterval53ee")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives/{drive-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listItem-id"+"}", _neturl.PathEscape(parameterToString(r.listItemId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"startDateTime"+"}", _neturl.PathEscape(parameterToString(r.startDateTime, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"endDateTime"+"}", _neturl.PathEscape(parameterToString(r.endDateTime, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"interval"+"}", _neturl.PathEscape(parameterToString(r.interval, "")), -1)

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

type ApiDrivesDriveListItemsListItemGetActivitiesByInterval96b0Request struct {
	ctx _context.Context
	ApiService *DrivesFunctionsApiService
	driveId string
	listItemId string
}


func (r ApiDrivesDriveListItemsListItemGetActivitiesByInterval96b0Request) Execute() ([]*AnyOfmicrosoftGraphItemActivityStat, *_nethttp.Response, error) {
	return r.ApiService.DrivesDriveListItemsListItemGetActivitiesByInterval96b0Execute(r)
}

/*
DrivesDriveListItemsListItemGetActivitiesByInterval96b0 Invoke function getActivitiesByInterval

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param listItemId key: id of listItem
 @return ApiDrivesDriveListItemsListItemGetActivitiesByInterval96b0Request
*/
func (a *DrivesFunctionsApiService) DrivesDriveListItemsListItemGetActivitiesByInterval96b0(ctx _context.Context, driveId string, listItemId string) ApiDrivesDriveListItemsListItemGetActivitiesByInterval96b0Request {
	return ApiDrivesDriveListItemsListItemGetActivitiesByInterval96b0Request{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
		listItemId: listItemId,
	}
}

// Execute executes the request
//  @return []*AnyOfmicrosoftGraphItemActivityStat
func (a *DrivesFunctionsApiService) DrivesDriveListItemsListItemGetActivitiesByInterval96b0Execute(r ApiDrivesDriveListItemsListItemGetActivitiesByInterval96b0Request) ([]*AnyOfmicrosoftGraphItemActivityStat, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []*AnyOfmicrosoftGraphItemActivityStat
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesFunctionsApiService.DrivesDriveListItemsListItemGetActivitiesByInterval96b0")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives/{drive-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval()"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listItem-id"+"}", _neturl.PathEscape(parameterToString(r.listItemId, "")), -1)

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

type ApiDrivesDriveRecentRequest struct {
	ctx _context.Context
	ApiService *DrivesFunctionsApiService
	driveId string
}


func (r ApiDrivesDriveRecentRequest) Execute() ([]*AnyOfmicrosoftGraphDriveItem, *_nethttp.Response, error) {
	return r.ApiService.DrivesDriveRecentExecute(r)
}

/*
DrivesDriveRecent Invoke function recent

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @return ApiDrivesDriveRecentRequest
*/
func (a *DrivesFunctionsApiService) DrivesDriveRecent(ctx _context.Context, driveId string) ApiDrivesDriveRecentRequest {
	return ApiDrivesDriveRecentRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
	}
}

// Execute executes the request
//  @return []*AnyOfmicrosoftGraphDriveItem
func (a *DrivesFunctionsApiService) DrivesDriveRecentExecute(r ApiDrivesDriveRecentRequest) ([]*AnyOfmicrosoftGraphDriveItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []*AnyOfmicrosoftGraphDriveItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesFunctionsApiService.DrivesDriveRecent")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives/{drive-id}/microsoft.graph.recent()"
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

type ApiDrivesDriveSearchRequest struct {
	ctx _context.Context
	ApiService *DrivesFunctionsApiService
	driveId string
	q string
}


func (r ApiDrivesDriveSearchRequest) Execute() ([]*AnyOfmicrosoftGraphDriveItem, *_nethttp.Response, error) {
	return r.ApiService.DrivesDriveSearchExecute(r)
}

/*
DrivesDriveSearch Invoke function search

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param q Usage: q={q}
 @return ApiDrivesDriveSearchRequest
*/
func (a *DrivesFunctionsApiService) DrivesDriveSearch(ctx _context.Context, driveId string, q string) ApiDrivesDriveSearchRequest {
	return ApiDrivesDriveSearchRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
		q: q,
	}
}

// Execute executes the request
//  @return []*AnyOfmicrosoftGraphDriveItem
func (a *DrivesFunctionsApiService) DrivesDriveSearchExecute(r ApiDrivesDriveSearchRequest) ([]*AnyOfmicrosoftGraphDriveItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []*AnyOfmicrosoftGraphDriveItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesFunctionsApiService.DrivesDriveSearch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives/{drive-id}/microsoft.graph.search(q='{q}')"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", _neturl.PathEscape(parameterToString(r.driveId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"q"+"}", _neturl.PathEscape(parameterToString(r.q, "")), -1)

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

type ApiDrivesDriveSharedWithMeRequest struct {
	ctx _context.Context
	ApiService *DrivesFunctionsApiService
	driveId string
}


func (r ApiDrivesDriveSharedWithMeRequest) Execute() ([]*AnyOfmicrosoftGraphDriveItem, *_nethttp.Response, error) {
	return r.ApiService.DrivesDriveSharedWithMeExecute(r)
}

/*
DrivesDriveSharedWithMe Invoke function sharedWithMe

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @return ApiDrivesDriveSharedWithMeRequest
*/
func (a *DrivesFunctionsApiService) DrivesDriveSharedWithMe(ctx _context.Context, driveId string) ApiDrivesDriveSharedWithMeRequest {
	return ApiDrivesDriveSharedWithMeRequest{
		ApiService: a,
		ctx: ctx,
		driveId: driveId,
	}
}

// Execute executes the request
//  @return []*AnyOfmicrosoftGraphDriveItem
func (a *DrivesFunctionsApiService) DrivesDriveSharedWithMeExecute(r ApiDrivesDriveSharedWithMeRequest) ([]*AnyOfmicrosoftGraphDriveItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []*AnyOfmicrosoftGraphDriveItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesFunctionsApiService.DrivesDriveSharedWithMe")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drives/{drive-id}/microsoft.graph.sharedWithMe()"
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
