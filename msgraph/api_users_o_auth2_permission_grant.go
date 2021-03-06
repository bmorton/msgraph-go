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

// UsersOAuth2PermissionGrantApiService UsersOAuth2PermissionGrantApi service
type UsersOAuth2PermissionGrantApiService service

type ApiUsersCreateRefOauth2PermissionGrantsRequest struct {
	ctx _context.Context
	ApiService *UsersOAuth2PermissionGrantApiService
	userId string
	requestBody *map[string]map[string]interface{}
}

// New navigation property ref value
func (r ApiUsersCreateRefOauth2PermissionGrantsRequest) RequestBody(requestBody map[string]map[string]interface{}) ApiUsersCreateRefOauth2PermissionGrantsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiUsersCreateRefOauth2PermissionGrantsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UsersCreateRefOauth2PermissionGrantsExecute(r)
}

/*
UsersCreateRefOauth2PermissionGrants Create new navigation property ref to oauth2PermissionGrants for users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersCreateRefOauth2PermissionGrantsRequest
*/
func (a *UsersOAuth2PermissionGrantApiService) UsersCreateRefOauth2PermissionGrants(ctx _context.Context, userId string) ApiUsersCreateRefOauth2PermissionGrantsRequest {
	return ApiUsersCreateRefOauth2PermissionGrantsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *UsersOAuth2PermissionGrantApiService) UsersCreateRefOauth2PermissionGrantsExecute(r ApiUsersCreateRefOauth2PermissionGrantsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOAuth2PermissionGrantApiService.UsersCreateRefOauth2PermissionGrants")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/oauth2PermissionGrants/$ref"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiUsersListOauth2PermissionGrantsRequest struct {
	ctx _context.Context
	ApiService *UsersOAuth2PermissionGrantApiService
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
func (r ApiUsersListOauth2PermissionGrantsRequest) Top(top int32) ApiUsersListOauth2PermissionGrantsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiUsersListOauth2PermissionGrantsRequest) Skip(skip int32) ApiUsersListOauth2PermissionGrantsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiUsersListOauth2PermissionGrantsRequest) Search(search string) ApiUsersListOauth2PermissionGrantsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiUsersListOauth2PermissionGrantsRequest) Filter(filter string) ApiUsersListOauth2PermissionGrantsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiUsersListOauth2PermissionGrantsRequest) Count(count bool) ApiUsersListOauth2PermissionGrantsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiUsersListOauth2PermissionGrantsRequest) Orderby(orderby []string) ApiUsersListOauth2PermissionGrantsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiUsersListOauth2PermissionGrantsRequest) Select_(select_ []string) ApiUsersListOauth2PermissionGrantsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersListOauth2PermissionGrantsRequest) Expand(expand []string) ApiUsersListOauth2PermissionGrantsRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersListOauth2PermissionGrantsRequest) Execute() (CollectionOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	return r.ApiService.UsersListOauth2PermissionGrantsExecute(r)
}

/*
UsersListOauth2PermissionGrants Get oauth2PermissionGrants from users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersListOauth2PermissionGrantsRequest
*/
func (a *UsersOAuth2PermissionGrantApiService) UsersListOauth2PermissionGrants(ctx _context.Context, userId string) ApiUsersListOauth2PermissionGrantsRequest {
	return ApiUsersListOauth2PermissionGrantsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfOAuth2PermissionGrant
func (a *UsersOAuth2PermissionGrantApiService) UsersListOauth2PermissionGrantsExecute(r ApiUsersListOauth2PermissionGrantsRequest) (CollectionOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfOAuth2PermissionGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOAuth2PermissionGrantApiService.UsersListOauth2PermissionGrants")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/oauth2PermissionGrants"
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

type ApiUsersListRefOauth2PermissionGrantsRequest struct {
	ctx _context.Context
	ApiService *UsersOAuth2PermissionGrantApiService
	userId string
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
}

// Show only the first n items
func (r ApiUsersListRefOauth2PermissionGrantsRequest) Top(top int32) ApiUsersListRefOauth2PermissionGrantsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiUsersListRefOauth2PermissionGrantsRequest) Skip(skip int32) ApiUsersListRefOauth2PermissionGrantsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiUsersListRefOauth2PermissionGrantsRequest) Search(search string) ApiUsersListRefOauth2PermissionGrantsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiUsersListRefOauth2PermissionGrantsRequest) Filter(filter string) ApiUsersListRefOauth2PermissionGrantsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiUsersListRefOauth2PermissionGrantsRequest) Count(count bool) ApiUsersListRefOauth2PermissionGrantsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiUsersListRefOauth2PermissionGrantsRequest) Orderby(orderby []string) ApiUsersListRefOauth2PermissionGrantsRequest {
	r.orderby = &orderby
	return r
}

func (r ApiUsersListRefOauth2PermissionGrantsRequest) Execute() (CollectionOfLinksOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	return r.ApiService.UsersListRefOauth2PermissionGrantsExecute(r)
}

/*
UsersListRefOauth2PermissionGrants Get ref of oauth2PermissionGrants from users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersListRefOauth2PermissionGrantsRequest
*/
func (a *UsersOAuth2PermissionGrantApiService) UsersListRefOauth2PermissionGrants(ctx _context.Context, userId string) ApiUsersListRefOauth2PermissionGrantsRequest {
	return ApiUsersListRefOauth2PermissionGrantsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfLinksOfOAuth2PermissionGrant
func (a *UsersOAuth2PermissionGrantApiService) UsersListRefOauth2PermissionGrantsExecute(r ApiUsersListRefOauth2PermissionGrantsRequest) (CollectionOfLinksOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfLinksOfOAuth2PermissionGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersOAuth2PermissionGrantApiService.UsersListRefOauth2PermissionGrants")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/oauth2PermissionGrants/$ref"
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
