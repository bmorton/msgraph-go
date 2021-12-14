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

// UsersSiteApiService UsersSiteApi service
type UsersSiteApiService service

type ApiUsersCreateRefFollowedSitesRequest struct {
	ctx _context.Context
	ApiService *UsersSiteApiService
	userId string
	requestBody *map[string]map[string]interface{}
}

// New navigation property ref value
func (r ApiUsersCreateRefFollowedSitesRequest) RequestBody(requestBody map[string]map[string]interface{}) ApiUsersCreateRefFollowedSitesRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiUsersCreateRefFollowedSitesRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.UsersCreateRefFollowedSitesExecute(r)
}

/*
UsersCreateRefFollowedSites Create new navigation property ref to followedSites for users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersCreateRefFollowedSitesRequest
*/
func (a *UsersSiteApiService) UsersCreateRefFollowedSites(ctx _context.Context, userId string) ApiUsersCreateRefFollowedSitesRequest {
	return ApiUsersCreateRefFollowedSitesRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *UsersSiteApiService) UsersCreateRefFollowedSitesExecute(r ApiUsersCreateRefFollowedSitesRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersSiteApiService.UsersCreateRefFollowedSites")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/followedSites/$ref"
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

type ApiUsersListFollowedSitesRequest struct {
	ctx _context.Context
	ApiService *UsersSiteApiService
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
func (r ApiUsersListFollowedSitesRequest) Top(top int32) ApiUsersListFollowedSitesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiUsersListFollowedSitesRequest) Skip(skip int32) ApiUsersListFollowedSitesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiUsersListFollowedSitesRequest) Search(search string) ApiUsersListFollowedSitesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiUsersListFollowedSitesRequest) Filter(filter string) ApiUsersListFollowedSitesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiUsersListFollowedSitesRequest) Count(count bool) ApiUsersListFollowedSitesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiUsersListFollowedSitesRequest) Orderby(orderby []string) ApiUsersListFollowedSitesRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiUsersListFollowedSitesRequest) Select_(select_ []string) ApiUsersListFollowedSitesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersListFollowedSitesRequest) Expand(expand []string) ApiUsersListFollowedSitesRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersListFollowedSitesRequest) Execute() (CollectionOfSite, *_nethttp.Response, error) {
	return r.ApiService.UsersListFollowedSitesExecute(r)
}

/*
UsersListFollowedSites Get followedSites from users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersListFollowedSitesRequest
*/
func (a *UsersSiteApiService) UsersListFollowedSites(ctx _context.Context, userId string) ApiUsersListFollowedSitesRequest {
	return ApiUsersListFollowedSitesRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfSite
func (a *UsersSiteApiService) UsersListFollowedSitesExecute(r ApiUsersListFollowedSitesRequest) (CollectionOfSite, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfSite
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersSiteApiService.UsersListFollowedSites")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/followedSites"
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

type ApiUsersListRefFollowedSitesRequest struct {
	ctx _context.Context
	ApiService *UsersSiteApiService
	userId string
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
}

// Show only the first n items
func (r ApiUsersListRefFollowedSitesRequest) Top(top int32) ApiUsersListRefFollowedSitesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiUsersListRefFollowedSitesRequest) Skip(skip int32) ApiUsersListRefFollowedSitesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiUsersListRefFollowedSitesRequest) Search(search string) ApiUsersListRefFollowedSitesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiUsersListRefFollowedSitesRequest) Filter(filter string) ApiUsersListRefFollowedSitesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiUsersListRefFollowedSitesRequest) Count(count bool) ApiUsersListRefFollowedSitesRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiUsersListRefFollowedSitesRequest) Orderby(orderby []string) ApiUsersListRefFollowedSitesRequest {
	r.orderby = &orderby
	return r
}

func (r ApiUsersListRefFollowedSitesRequest) Execute() (CollectionOfLinksOfSite, *_nethttp.Response, error) {
	return r.ApiService.UsersListRefFollowedSitesExecute(r)
}

/*
UsersListRefFollowedSites Get ref of followedSites from users

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersListRefFollowedSitesRequest
*/
func (a *UsersSiteApiService) UsersListRefFollowedSites(ctx _context.Context, userId string) ApiUsersListRefFollowedSitesRequest {
	return ApiUsersListRefFollowedSitesRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfLinksOfSite
func (a *UsersSiteApiService) UsersListRefFollowedSitesExecute(r ApiUsersListRefFollowedSitesRequest) (CollectionOfLinksOfSite, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfLinksOfSite
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersSiteApiService.UsersListRefFollowedSites")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/followedSites/$ref"
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
