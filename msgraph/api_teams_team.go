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

// TeamsTeamApiService TeamsTeamApi service
type TeamsTeamApiService service

type ApiTeamsTeamCreateTeamRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamApiService
	microsoftGraphTeam *MicrosoftGraphTeam
}

// New entity
func (r ApiTeamsTeamCreateTeamRequest) MicrosoftGraphTeam(microsoftGraphTeam MicrosoftGraphTeam) ApiTeamsTeamCreateTeamRequest {
	r.microsoftGraphTeam = &microsoftGraphTeam
	return r
}

func (r ApiTeamsTeamCreateTeamRequest) Execute() (MicrosoftGraphTeam, *_nethttp.Response, error) {
	return r.ApiService.TeamsTeamCreateTeamExecute(r)
}

/*
TeamsTeamCreateTeam Add new entity to teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTeamsTeamCreateTeamRequest
*/
func (a *TeamsTeamApiService) TeamsTeamCreateTeam(ctx _context.Context) ApiTeamsTeamCreateTeamRequest {
	return ApiTeamsTeamCreateTeamRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeam
func (a *TeamsTeamApiService) TeamsTeamCreateTeamExecute(r ApiTeamsTeamCreateTeamRequest) (MicrosoftGraphTeam, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamApiService.TeamsTeamCreateTeam")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams"

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

type ApiTeamsTeamDeleteTeamRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamApiService
	teamId string
	ifMatch *string
}

// ETag
func (r ApiTeamsTeamDeleteTeamRequest) IfMatch(ifMatch string) ApiTeamsTeamDeleteTeamRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiTeamsTeamDeleteTeamRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsTeamDeleteTeamExecute(r)
}

/*
TeamsTeamDeleteTeam Delete entity from teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsTeamDeleteTeamRequest
*/
func (a *TeamsTeamApiService) TeamsTeamDeleteTeam(ctx _context.Context, teamId string) ApiTeamsTeamDeleteTeamRequest {
	return ApiTeamsTeamDeleteTeamRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
func (a *TeamsTeamApiService) TeamsTeamDeleteTeamExecute(r ApiTeamsTeamDeleteTeamRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamApiService.TeamsTeamDeleteTeam")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}"
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

type ApiTeamsTeamGetTeamRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamApiService
	teamId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiTeamsTeamGetTeamRequest) Select_(select_ []string) ApiTeamsTeamGetTeamRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsTeamGetTeamRequest) Expand(expand []string) ApiTeamsTeamGetTeamRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsTeamGetTeamRequest) Execute() (MicrosoftGraphTeam, *_nethttp.Response, error) {
	return r.ApiService.TeamsTeamGetTeamExecute(r)
}

/*
TeamsTeamGetTeam Get entity from teams by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsTeamGetTeamRequest
*/
func (a *TeamsTeamApiService) TeamsTeamGetTeam(ctx _context.Context, teamId string) ApiTeamsTeamGetTeamRequest {
	return ApiTeamsTeamGetTeamRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeam
func (a *TeamsTeamApiService) TeamsTeamGetTeamExecute(r ApiTeamsTeamGetTeamRequest) (MicrosoftGraphTeam, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamApiService.TeamsTeamGetTeam")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}"
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

type ApiTeamsTeamListTeamRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamApiService
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
func (r ApiTeamsTeamListTeamRequest) Top(top int32) ApiTeamsTeamListTeamRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiTeamsTeamListTeamRequest) Skip(skip int32) ApiTeamsTeamListTeamRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiTeamsTeamListTeamRequest) Search(search string) ApiTeamsTeamListTeamRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiTeamsTeamListTeamRequest) Filter(filter string) ApiTeamsTeamListTeamRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiTeamsTeamListTeamRequest) Count(count bool) ApiTeamsTeamListTeamRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiTeamsTeamListTeamRequest) Orderby(orderby []string) ApiTeamsTeamListTeamRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiTeamsTeamListTeamRequest) Select_(select_ []string) ApiTeamsTeamListTeamRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsTeamListTeamRequest) Expand(expand []string) ApiTeamsTeamListTeamRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsTeamListTeamRequest) Execute() (CollectionOfTeam, *_nethttp.Response, error) {
	return r.ApiService.TeamsTeamListTeamExecute(r)
}

/*
TeamsTeamListTeam Get entities from teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTeamsTeamListTeamRequest
*/
func (a *TeamsTeamApiService) TeamsTeamListTeam(ctx _context.Context) ApiTeamsTeamListTeamRequest {
	return ApiTeamsTeamListTeamRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfTeam
func (a *TeamsTeamApiService) TeamsTeamListTeamExecute(r ApiTeamsTeamListTeamRequest) (CollectionOfTeam, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfTeam
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamApiService.TeamsTeamListTeam")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams"

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

type ApiTeamsTeamUpdateTeamRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamApiService
	teamId string
	microsoftGraphTeam *MicrosoftGraphTeam
}

// New property values
func (r ApiTeamsTeamUpdateTeamRequest) MicrosoftGraphTeam(microsoftGraphTeam MicrosoftGraphTeam) ApiTeamsTeamUpdateTeamRequest {
	r.microsoftGraphTeam = &microsoftGraphTeam
	return r
}

func (r ApiTeamsTeamUpdateTeamRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsTeamUpdateTeamExecute(r)
}

/*
TeamsTeamUpdateTeam Update entity in teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsTeamUpdateTeamRequest
*/
func (a *TeamsTeamApiService) TeamsTeamUpdateTeam(ctx _context.Context, teamId string) ApiTeamsTeamUpdateTeamRequest {
	return ApiTeamsTeamUpdateTeamRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
func (a *TeamsTeamApiService) TeamsTeamUpdateTeamExecute(r ApiTeamsTeamUpdateTeamRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamApiService.TeamsTeamUpdateTeam")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}"
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
