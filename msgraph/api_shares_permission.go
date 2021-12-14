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

// SharesPermissionApiService SharesPermissionApi service
type SharesPermissionApiService service

type ApiSharesDeletePermissionRequest struct {
	ctx _context.Context
	ApiService *SharesPermissionApiService
	sharedDriveItemId string
	ifMatch *string
}

// ETag
func (r ApiSharesDeletePermissionRequest) IfMatch(ifMatch string) ApiSharesDeletePermissionRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiSharesDeletePermissionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SharesDeletePermissionExecute(r)
}

/*
SharesDeletePermission Delete navigation property permission for shares

Used to access the permission representing the underlying sharing link

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sharedDriveItemId key: id of sharedDriveItem
 @return ApiSharesDeletePermissionRequest
*/
func (a *SharesPermissionApiService) SharesDeletePermission(ctx _context.Context, sharedDriveItemId string) ApiSharesDeletePermissionRequest {
	return ApiSharesDeletePermissionRequest{
		ApiService: a,
		ctx: ctx,
		sharedDriveItemId: sharedDriveItemId,
	}
}

// Execute executes the request
func (a *SharesPermissionApiService) SharesDeletePermissionExecute(r ApiSharesDeletePermissionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharesPermissionApiService.SharesDeletePermission")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shares/{sharedDriveItem-id}/permission"
	localVarPath = strings.Replace(localVarPath, "{"+"sharedDriveItem-id"+"}", _neturl.PathEscape(parameterToString(r.sharedDriveItemId, "")), -1)

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

type ApiSharesGetPermissionRequest struct {
	ctx _context.Context
	ApiService *SharesPermissionApiService
	sharedDriveItemId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiSharesGetPermissionRequest) Select_(select_ []string) ApiSharesGetPermissionRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiSharesGetPermissionRequest) Expand(expand []string) ApiSharesGetPermissionRequest {
	r.expand = &expand
	return r
}

func (r ApiSharesGetPermissionRequest) Execute() (MicrosoftGraphPermission, *_nethttp.Response, error) {
	return r.ApiService.SharesGetPermissionExecute(r)
}

/*
SharesGetPermission Get permission from shares

Used to access the permission representing the underlying sharing link

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sharedDriveItemId key: id of sharedDriveItem
 @return ApiSharesGetPermissionRequest
*/
func (a *SharesPermissionApiService) SharesGetPermission(ctx _context.Context, sharedDriveItemId string) ApiSharesGetPermissionRequest {
	return ApiSharesGetPermissionRequest{
		ApiService: a,
		ctx: ctx,
		sharedDriveItemId: sharedDriveItemId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphPermission
func (a *SharesPermissionApiService) SharesGetPermissionExecute(r ApiSharesGetPermissionRequest) (MicrosoftGraphPermission, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphPermission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharesPermissionApiService.SharesGetPermission")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shares/{sharedDriveItem-id}/permission"
	localVarPath = strings.Replace(localVarPath, "{"+"sharedDriveItem-id"+"}", _neturl.PathEscape(parameterToString(r.sharedDriveItemId, "")), -1)

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

type ApiSharesUpdatePermissionRequest struct {
	ctx _context.Context
	ApiService *SharesPermissionApiService
	sharedDriveItemId string
	microsoftGraphPermission *MicrosoftGraphPermission
}

// New navigation property values
func (r ApiSharesUpdatePermissionRequest) MicrosoftGraphPermission(microsoftGraphPermission MicrosoftGraphPermission) ApiSharesUpdatePermissionRequest {
	r.microsoftGraphPermission = &microsoftGraphPermission
	return r
}

func (r ApiSharesUpdatePermissionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SharesUpdatePermissionExecute(r)
}

/*
SharesUpdatePermission Update the navigation property permission in shares

Used to access the permission representing the underlying sharing link

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sharedDriveItemId key: id of sharedDriveItem
 @return ApiSharesUpdatePermissionRequest
*/
func (a *SharesPermissionApiService) SharesUpdatePermission(ctx _context.Context, sharedDriveItemId string) ApiSharesUpdatePermissionRequest {
	return ApiSharesUpdatePermissionRequest{
		ApiService: a,
		ctx: ctx,
		sharedDriveItemId: sharedDriveItemId,
	}
}

// Execute executes the request
func (a *SharesPermissionApiService) SharesUpdatePermissionExecute(r ApiSharesUpdatePermissionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharesPermissionApiService.SharesUpdatePermission")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shares/{sharedDriveItem-id}/permission"
	localVarPath = strings.Replace(localVarPath, "{"+"sharedDriveItem-id"+"}", _neturl.PathEscape(parameterToString(r.sharedDriveItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphPermission == nil {
		return nil, reportError("microsoftGraphPermission is required and must be specified")
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
	localVarPostBody = r.microsoftGraphPermission
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