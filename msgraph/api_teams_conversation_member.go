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

// TeamsConversationMemberApiService TeamsConversationMemberApi service
type TeamsConversationMemberApiService service

type ApiTeamsCreateMembersRequest struct {
	ctx _context.Context
	ApiService *TeamsConversationMemberApiService
	teamId string
	microsoftGraphConversationMember *MicrosoftGraphConversationMember
}

// New navigation property
func (r ApiTeamsCreateMembersRequest) MicrosoftGraphConversationMember(microsoftGraphConversationMember MicrosoftGraphConversationMember) ApiTeamsCreateMembersRequest {
	r.microsoftGraphConversationMember = &microsoftGraphConversationMember
	return r
}

func (r ApiTeamsCreateMembersRequest) Execute() (MicrosoftGraphConversationMember, *_nethttp.Response, error) {
	return r.ApiService.TeamsCreateMembersExecute(r)
}

/*
TeamsCreateMembers Create new navigation property to members for teams

Members and owners of the team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsCreateMembersRequest
*/
func (a *TeamsConversationMemberApiService) TeamsCreateMembers(ctx _context.Context, teamId string) ApiTeamsCreateMembersRequest {
	return ApiTeamsCreateMembersRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphConversationMember
func (a *TeamsConversationMemberApiService) TeamsCreateMembersExecute(r ApiTeamsCreateMembersRequest) (MicrosoftGraphConversationMember, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphConversationMember
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsConversationMemberApiService.TeamsCreateMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphConversationMember == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphConversationMember is required and must be specified")
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
	localVarPostBody = r.microsoftGraphConversationMember
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

type ApiTeamsDeleteMembersRequest struct {
	ctx _context.Context
	ApiService *TeamsConversationMemberApiService
	teamId string
	conversationMemberId string
	ifMatch *string
}

// ETag
func (r ApiTeamsDeleteMembersRequest) IfMatch(ifMatch string) ApiTeamsDeleteMembersRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiTeamsDeleteMembersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsDeleteMembersExecute(r)
}

/*
TeamsDeleteMembers Delete navigation property members for teams

Members and owners of the team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @param conversationMemberId key: id of conversationMember
 @return ApiTeamsDeleteMembersRequest
*/
func (a *TeamsConversationMemberApiService) TeamsDeleteMembers(ctx _context.Context, teamId string, conversationMemberId string) ApiTeamsDeleteMembersRequest {
	return ApiTeamsDeleteMembersRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		conversationMemberId: conversationMemberId,
	}
}

// Execute executes the request
func (a *TeamsConversationMemberApiService) TeamsDeleteMembersExecute(r ApiTeamsDeleteMembersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsConversationMemberApiService.TeamsDeleteMembers")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/members/{conversationMember-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"conversationMember-id"+"}", _neturl.PathEscape(parameterToString(r.conversationMemberId, "")), -1)

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

type ApiTeamsGetMembersRequest struct {
	ctx _context.Context
	ApiService *TeamsConversationMemberApiService
	teamId string
	conversationMemberId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiTeamsGetMembersRequest) Select_(select_ []string) ApiTeamsGetMembersRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsGetMembersRequest) Expand(expand []string) ApiTeamsGetMembersRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsGetMembersRequest) Execute() (MicrosoftGraphConversationMember, *_nethttp.Response, error) {
	return r.ApiService.TeamsGetMembersExecute(r)
}

/*
TeamsGetMembers Get members from teams

Members and owners of the team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @param conversationMemberId key: id of conversationMember
 @return ApiTeamsGetMembersRequest
*/
func (a *TeamsConversationMemberApiService) TeamsGetMembers(ctx _context.Context, teamId string, conversationMemberId string) ApiTeamsGetMembersRequest {
	return ApiTeamsGetMembersRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		conversationMemberId: conversationMemberId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphConversationMember
func (a *TeamsConversationMemberApiService) TeamsGetMembersExecute(r ApiTeamsGetMembersRequest) (MicrosoftGraphConversationMember, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphConversationMember
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsConversationMemberApiService.TeamsGetMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/members/{conversationMember-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"conversationMember-id"+"}", _neturl.PathEscape(parameterToString(r.conversationMemberId, "")), -1)

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

type ApiTeamsListMembersRequest struct {
	ctx _context.Context
	ApiService *TeamsConversationMemberApiService
	teamId string
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
func (r ApiTeamsListMembersRequest) Top(top int32) ApiTeamsListMembersRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiTeamsListMembersRequest) Skip(skip int32) ApiTeamsListMembersRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiTeamsListMembersRequest) Search(search string) ApiTeamsListMembersRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiTeamsListMembersRequest) Filter(filter string) ApiTeamsListMembersRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiTeamsListMembersRequest) Count(count bool) ApiTeamsListMembersRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiTeamsListMembersRequest) Orderby(orderby []string) ApiTeamsListMembersRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiTeamsListMembersRequest) Select_(select_ []string) ApiTeamsListMembersRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsListMembersRequest) Expand(expand []string) ApiTeamsListMembersRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsListMembersRequest) Execute() (CollectionOfConversationMember, *_nethttp.Response, error) {
	return r.ApiService.TeamsListMembersExecute(r)
}

/*
TeamsListMembers Get members from teams

Members and owners of the team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsListMembersRequest
*/
func (a *TeamsConversationMemberApiService) TeamsListMembers(ctx _context.Context, teamId string) ApiTeamsListMembersRequest {
	return ApiTeamsListMembersRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return CollectionOfConversationMember
func (a *TeamsConversationMemberApiService) TeamsListMembersExecute(r ApiTeamsListMembersRequest) (CollectionOfConversationMember, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfConversationMember
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsConversationMemberApiService.TeamsListMembers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

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

type ApiTeamsUpdateMembersRequest struct {
	ctx _context.Context
	ApiService *TeamsConversationMemberApiService
	teamId string
	conversationMemberId string
	microsoftGraphConversationMember *MicrosoftGraphConversationMember
}

// New navigation property values
func (r ApiTeamsUpdateMembersRequest) MicrosoftGraphConversationMember(microsoftGraphConversationMember MicrosoftGraphConversationMember) ApiTeamsUpdateMembersRequest {
	r.microsoftGraphConversationMember = &microsoftGraphConversationMember
	return r
}

func (r ApiTeamsUpdateMembersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsUpdateMembersExecute(r)
}

/*
TeamsUpdateMembers Update the navigation property members in teams

Members and owners of the team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @param conversationMemberId key: id of conversationMember
 @return ApiTeamsUpdateMembersRequest
*/
func (a *TeamsConversationMemberApiService) TeamsUpdateMembers(ctx _context.Context, teamId string, conversationMemberId string) ApiTeamsUpdateMembersRequest {
	return ApiTeamsUpdateMembersRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		conversationMemberId: conversationMemberId,
	}
}

// Execute executes the request
func (a *TeamsConversationMemberApiService) TeamsUpdateMembersExecute(r ApiTeamsUpdateMembersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsConversationMemberApiService.TeamsUpdateMembers")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/members/{conversationMember-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"conversationMember-id"+"}", _neturl.PathEscape(parameterToString(r.conversationMemberId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphConversationMember == nil {
		return nil, reportError("microsoftGraphConversationMember is required and must be specified")
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
	localVarPostBody = r.microsoftGraphConversationMember
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
