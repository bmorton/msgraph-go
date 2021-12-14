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

// UsersTeamApiService UsersTeamApi service
type UsersTeamApiService service

type ApiUsersCreateJoinedTeamsRequest struct {
	ctx _context.Context
	ApiService *UsersTeamApiService
	userId string
	microsoftGraphTeam *MicrosoftGraphTeam
}

// New navigation property
func (r ApiUsersCreateJoinedTeamsRequest) MicrosoftGraphTeam(microsoftGraphTeam MicrosoftGraphTeam) ApiUsersCreateJoinedTeamsRequest {
	r.microsoftGraphTeam = &microsoftGraphTeam
	return r
}

func (r ApiUsersCreateJoinedTeamsRequest) Execute() (MicrosoftGraphTeam, *_nethttp.Response, error) {
	return r.ApiService.UsersCreateJoinedTeamsExecute(r)
}

/*
UsersCreateJoinedTeams Create new navigation property to joinedTeams for users

The Microsoft Teams teams that the user is a member of. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersCreateJoinedTeamsRequest
*/
func (a *UsersTeamApiService) UsersCreateJoinedTeams(ctx _context.Context, userId string) ApiUsersCreateJoinedTeamsRequest {
	return ApiUsersCreateJoinedTeamsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeam
func (a *UsersTeamApiService) UsersCreateJoinedTeamsExecute(r ApiUsersCreateJoinedTeamsRequest) (MicrosoftGraphTeam, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersTeamApiService.UsersCreateJoinedTeams")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/joinedTeams"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphTeam == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphTeam is required and must be specified")
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
	localVarPostBody = r.microsoftGraphTeam
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

type ApiUsersDeleteJoinedTeamsRequest struct {
	ctx _context.Context
	ApiService *UsersTeamApiService
	userId string
	teamId string
	ifMatch *string
}

// ETag
func (r ApiUsersDeleteJoinedTeamsRequest) IfMatch(ifMatch string) ApiUsersDeleteJoinedTeamsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUsersDeleteJoinedTeamsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersDeleteJoinedTeamsExecute(r)
}

/*
UsersDeleteJoinedTeams Delete navigation property joinedTeams for users

The Microsoft Teams teams that the user is a member of. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param teamId key: id of team
 @return ApiUsersDeleteJoinedTeamsRequest
*/
func (a *UsersTeamApiService) UsersDeleteJoinedTeams(ctx _context.Context, userId string, teamId string) ApiUsersDeleteJoinedTeamsRequest {
	return ApiUsersDeleteJoinedTeamsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		teamId: teamId,
	}
}

// Execute executes the request
func (a *UsersTeamApiService) UsersDeleteJoinedTeamsExecute(r ApiUsersDeleteJoinedTeamsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersTeamApiService.UsersDeleteJoinedTeams")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/joinedTeams/{team-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

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

type ApiUsersGetJoinedTeamsRequest struct {
	ctx _context.Context
	ApiService *UsersTeamApiService
	userId string
	teamId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiUsersGetJoinedTeamsRequest) Select_(select_ []string) ApiUsersGetJoinedTeamsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersGetJoinedTeamsRequest) Expand(expand []string) ApiUsersGetJoinedTeamsRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersGetJoinedTeamsRequest) Execute() (MicrosoftGraphTeam, *_nethttp.Response, error) {
	return r.ApiService.UsersGetJoinedTeamsExecute(r)
}

/*
UsersGetJoinedTeams Get joinedTeams from users

The Microsoft Teams teams that the user is a member of. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param teamId key: id of team
 @return ApiUsersGetJoinedTeamsRequest
*/
func (a *UsersTeamApiService) UsersGetJoinedTeams(ctx _context.Context, userId string, teamId string) ApiUsersGetJoinedTeamsRequest {
	return ApiUsersGetJoinedTeamsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeam
func (a *UsersTeamApiService) UsersGetJoinedTeamsExecute(r ApiUsersGetJoinedTeamsRequest) (MicrosoftGraphTeam, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersTeamApiService.UsersGetJoinedTeams")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/joinedTeams/{team-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

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

type ApiUsersListJoinedTeamsRequest struct {
	ctx _context.Context
	ApiService *UsersTeamApiService
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
func (r ApiUsersListJoinedTeamsRequest) Top(top int32) ApiUsersListJoinedTeamsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiUsersListJoinedTeamsRequest) Skip(skip int32) ApiUsersListJoinedTeamsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiUsersListJoinedTeamsRequest) Search(search string) ApiUsersListJoinedTeamsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiUsersListJoinedTeamsRequest) Filter(filter string) ApiUsersListJoinedTeamsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiUsersListJoinedTeamsRequest) Count(count bool) ApiUsersListJoinedTeamsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiUsersListJoinedTeamsRequest) Orderby(orderby []string) ApiUsersListJoinedTeamsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiUsersListJoinedTeamsRequest) Select_(select_ []string) ApiUsersListJoinedTeamsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiUsersListJoinedTeamsRequest) Expand(expand []string) ApiUsersListJoinedTeamsRequest {
	r.expand = &expand
	return r
}

func (r ApiUsersListJoinedTeamsRequest) Execute() (CollectionOfTeam, *_nethttp.Response, error) {
	return r.ApiService.UsersListJoinedTeamsExecute(r)
}

/*
UsersListJoinedTeams Get joinedTeams from users

The Microsoft Teams teams that the user is a member of. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUsersListJoinedTeamsRequest
*/
func (a *UsersTeamApiService) UsersListJoinedTeams(ctx _context.Context, userId string) ApiUsersListJoinedTeamsRequest {
	return ApiUsersListJoinedTeamsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfTeam
func (a *UsersTeamApiService) UsersListJoinedTeamsExecute(r ApiUsersListJoinedTeamsRequest) (CollectionOfTeam, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersTeamApiService.UsersListJoinedTeams")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/joinedTeams"
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

type ApiUsersUpdateJoinedTeamsRequest struct {
	ctx _context.Context
	ApiService *UsersTeamApiService
	userId string
	teamId string
	microsoftGraphTeam *MicrosoftGraphTeam
}

// New navigation property values
func (r ApiUsersUpdateJoinedTeamsRequest) MicrosoftGraphTeam(microsoftGraphTeam MicrosoftGraphTeam) ApiUsersUpdateJoinedTeamsRequest {
	r.microsoftGraphTeam = &microsoftGraphTeam
	return r
}

func (r ApiUsersUpdateJoinedTeamsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UsersUpdateJoinedTeamsExecute(r)
}

/*
UsersUpdateJoinedTeams Update the navigation property joinedTeams in users

The Microsoft Teams teams that the user is a member of. Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param teamId key: id of team
 @return ApiUsersUpdateJoinedTeamsRequest
*/
func (a *UsersTeamApiService) UsersUpdateJoinedTeams(ctx _context.Context, userId string, teamId string) ApiUsersUpdateJoinedTeamsRequest {
	return ApiUsersUpdateJoinedTeamsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		teamId: teamId,
	}
}

// Execute executes the request
func (a *UsersTeamApiService) UsersUpdateJoinedTeamsExecute(r ApiUsersUpdateJoinedTeamsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersTeamApiService.UsersUpdateJoinedTeams")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/joinedTeams/{team-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphTeam == nil {
		return nil, reportError("microsoftGraphTeam is required and must be specified")
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
	localVarPostBody = r.microsoftGraphTeam
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