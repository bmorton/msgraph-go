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

// UsersScopedRoleMembershipApiService UsersScopedRoleMembershipApi service
type UsersScopedRoleMembershipApiService service

type ApiUsersCreateScopedRoleMemberOfRequest struct {
	ctx _context.Context
	ApiService *UsersScopedRoleMembershipApiService
	userId string
	microsoftGraphScopedRoleMembership *MicrosoftGraphScopedRoleMembership
}

// New navigation property
func (r ApiUsersCreateScopedRoleMemberOfRequest) MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership MicrosoftGraphScopedRoleMembership) ApiUsersCreateScopedRoleMemberOfRequest {
	r.microsoftGraphScopedRoleMembership = &microsoftGraphScopedRoleMembership
	return r
}

func (r ApiUsersCreateScopedRoleMemberOfRequest) Execute() (MicrosoftGraphScopedRoleMembership, *_nethttp.Response, error) {
	return r.ApiService.UsersCreateScopedRoleMemberOfExecute(r)
}

/*
UsersCreateScopedRoleMemberOf Create new navigation property to scopedRoleMemberOf for users

The scoped-role administrative unit memberships for this user. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersCreateScopedRoleMemberOfRequest
*/
func (a *UsersScopedRoleMembershipApiService) UsersCreateScopedRoleMemberOf(ctx _context.Context, userId string) ApiUsersCreateScopedRoleMemberOfRequest {
	return ApiUsersCreateScopedRoleMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphScopedRoleMembership
func (a *UsersScopedRoleMembershipApiService) UsersCreateScopedRoleMemberOfExecute(r ApiUsersCreateScopedRoleMemberOfRequest) (MicrosoftGraphScopedRoleMembership, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphScopedRoleMembership
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersScopedRoleMembershipApiService.UsersCreateScopedRoleMemberOf")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/scopedRoleMemberOf"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphScopedRoleMembership == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphScopedRoleMembership is required and must be specified")
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
	localVarPostBody = r.microsoftGraphScopedRoleMembership
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

type ApiUsersDeleteScopedRoleMemberOfRequest struct {
	ctx _context.Context
	ApiService *UsersScopedRoleMembershipApiService
	userId string
	scopedRoleMembershipId string
	ifMatch *string
}

// ETag
func (r ApiUsersDeleteScopedRoleMemberOfRequest) IfMatch(ifMatch string) ApiUsersDeleteScopedRoleMemberOfRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUsersDeleteScopedRoleMemberOfRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersDeleteScopedRoleMemberOfExecute(r)
}

/*
UsersDeleteScopedRoleMemberOf Delete navigation property scopedRoleMemberOf for users

The scoped-role administrative unit memberships for this user. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param scopedRoleMembershipId key: id of scopedRoleMembership
 @return ApiUsersDeleteScopedRoleMemberOfRequest
*/
func (a *UsersScopedRoleMembershipApiService) UsersDeleteScopedRoleMemberOf(ctx _context.Context, userId string, scopedRoleMembershipId string) ApiUsersDeleteScopedRoleMemberOfRequest {
	return ApiUsersDeleteScopedRoleMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		scopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// Execute executes the request
func (a *UsersScopedRoleMembershipApiService) UsersDeleteScopedRoleMemberOfExecute(r ApiUsersDeleteScopedRoleMemberOfRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersScopedRoleMembershipApiService.UsersDeleteScopedRoleMemberOf")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopedRoleMembership-id"+"}", _neturl.PathEscape(parameterToString(r.scopedRoleMembershipId, "")), -1)

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

type ApiUsersGetScopedRoleMemberOfRequest struct {
	ctx _context.Context
	ApiService *UsersScopedRoleMembershipApiService
	userId string
	scopedRoleMembershipId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiUsersGetScopedRoleMemberOfRequest) Select_(select_ []string) ApiUsersGetScopedRoleMemberOfRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersGetScopedRoleMemberOfRequest) Expand(expand []string) ApiUsersGetScopedRoleMemberOfRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersGetScopedRoleMemberOfRequest) Execute() (MicrosoftGraphScopedRoleMembership, *_nethttp.Response, error) {
	return r.ApiService.UsersGetScopedRoleMemberOfExecute(r)
}

/*
UsersGetScopedRoleMemberOf Get scopedRoleMemberOf from users

The scoped-role administrative unit memberships for this user. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param scopedRoleMembershipId key: id of scopedRoleMembership
 @return ApiUsersGetScopedRoleMemberOfRequest
*/
func (a *UsersScopedRoleMembershipApiService) UsersGetScopedRoleMemberOf(ctx _context.Context, userId string, scopedRoleMembershipId string) ApiUsersGetScopedRoleMemberOfRequest {
	return ApiUsersGetScopedRoleMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		scopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphScopedRoleMembership
func (a *UsersScopedRoleMembershipApiService) UsersGetScopedRoleMemberOfExecute(r ApiUsersGetScopedRoleMemberOfRequest) (MicrosoftGraphScopedRoleMembership, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphScopedRoleMembership
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersScopedRoleMembershipApiService.UsersGetScopedRoleMemberOf")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopedRoleMembership-id"+"}", _neturl.PathEscape(parameterToString(r.scopedRoleMembershipId, "")), -1)

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

type ApiUsersListScopedRoleMemberOfRequest struct {
	ctx _context.Context
	ApiService *UsersScopedRoleMembershipApiService
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
func (r ApiUsersListScopedRoleMemberOfRequest) Top(top int32) ApiUsersListScopedRoleMemberOfRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiUsersListScopedRoleMemberOfRequest) Skip(skip int32) ApiUsersListScopedRoleMemberOfRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiUsersListScopedRoleMemberOfRequest) Search(search string) ApiUsersListScopedRoleMemberOfRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiUsersListScopedRoleMemberOfRequest) Filter(filter string) ApiUsersListScopedRoleMemberOfRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiUsersListScopedRoleMemberOfRequest) Count(count bool) ApiUsersListScopedRoleMemberOfRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiUsersListScopedRoleMemberOfRequest) Orderby(orderby []string) ApiUsersListScopedRoleMemberOfRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiUsersListScopedRoleMemberOfRequest) Select_(select_ []string) ApiUsersListScopedRoleMemberOfRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersListScopedRoleMemberOfRequest) Expand(expand []string) ApiUsersListScopedRoleMemberOfRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersListScopedRoleMemberOfRequest) Execute() (CollectionOfScopedRoleMembership, *_nethttp.Response, error) {
	return r.ApiService.UsersListScopedRoleMemberOfExecute(r)
}

/*
UsersListScopedRoleMemberOf Get scopedRoleMemberOf from users

The scoped-role administrative unit memberships for this user. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersListScopedRoleMemberOfRequest
*/
func (a *UsersScopedRoleMembershipApiService) UsersListScopedRoleMemberOf(ctx _context.Context, userId string) ApiUsersListScopedRoleMemberOfRequest {
	return ApiUsersListScopedRoleMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfScopedRoleMembership
func (a *UsersScopedRoleMembershipApiService) UsersListScopedRoleMemberOfExecute(r ApiUsersListScopedRoleMemberOfRequest) (CollectionOfScopedRoleMembership, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfScopedRoleMembership
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersScopedRoleMembershipApiService.UsersListScopedRoleMemberOf")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/scopedRoleMemberOf"
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

type ApiUsersUpdateScopedRoleMemberOfRequest struct {
	ctx _context.Context
	ApiService *UsersScopedRoleMembershipApiService
	userId string
	scopedRoleMembershipId string
	microsoftGraphScopedRoleMembership *MicrosoftGraphScopedRoleMembership
}

// New navigation property values
func (r ApiUsersUpdateScopedRoleMemberOfRequest) MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership MicrosoftGraphScopedRoleMembership) ApiUsersUpdateScopedRoleMemberOfRequest {
	r.microsoftGraphScopedRoleMembership = &microsoftGraphScopedRoleMembership
	return r
}

func (r ApiUsersUpdateScopedRoleMemberOfRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersUpdateScopedRoleMemberOfExecute(r)
}

/*
UsersUpdateScopedRoleMemberOf Update the navigation property scopedRoleMemberOf in users

The scoped-role administrative unit memberships for this user. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param scopedRoleMembershipId key: id of scopedRoleMembership
 @return ApiUsersUpdateScopedRoleMemberOfRequest
*/
func (a *UsersScopedRoleMembershipApiService) UsersUpdateScopedRoleMemberOf(ctx _context.Context, userId string, scopedRoleMembershipId string) ApiUsersUpdateScopedRoleMemberOfRequest {
	return ApiUsersUpdateScopedRoleMemberOfRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		scopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// Execute executes the request
func (a *UsersScopedRoleMembershipApiService) UsersUpdateScopedRoleMemberOfExecute(r ApiUsersUpdateScopedRoleMemberOfRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersScopedRoleMembershipApiService.UsersUpdateScopedRoleMemberOf")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/scopedRoleMemberOf/{scopedRoleMembership-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopedRoleMembership-id"+"}", _neturl.PathEscape(parameterToString(r.scopedRoleMembershipId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphScopedRoleMembership == nil {
		return nil, reportError("microsoftGraphScopedRoleMembership is required and must be specified")
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
	localVarPostBody = r.microsoftGraphScopedRoleMembership
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
