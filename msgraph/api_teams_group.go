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

// TeamsGroupApiService TeamsGroupApi service
type TeamsGroupApiService service

type ApiTeamsDeleteRefGroupRequest struct {
	ctx _context.Context
	ApiService *TeamsGroupApiService
	teamId string
	ifMatch *string
}

// ETag
func (r ApiTeamsDeleteRefGroupRequest) IfMatch(ifMatch string) ApiTeamsDeleteRefGroupRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiTeamsDeleteRefGroupRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsDeleteRefGroupExecute(r)
}

/*
TeamsDeleteRefGroup Delete ref of navigation property group for teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsDeleteRefGroupRequest
*/
func (a *TeamsGroupApiService) TeamsDeleteRefGroup(ctx _context.Context, teamId string) ApiTeamsDeleteRefGroupRequest {
	return ApiTeamsDeleteRefGroupRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
func (a *TeamsGroupApiService) TeamsDeleteRefGroupExecute(r ApiTeamsDeleteRefGroupRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsGroupApiService.TeamsDeleteRefGroup")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/group/$ref"
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

type ApiTeamsGetGroupRequest struct {
	ctx _context.Context
	ApiService *TeamsGroupApiService
	teamId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiTeamsGetGroupRequest) Select_(select_ []string) ApiTeamsGetGroupRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsGetGroupRequest) Expand(expand []string) ApiTeamsGetGroupRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsGetGroupRequest) Execute() (MicrosoftGraphGroup, *_nethttp.Response, error) {
	return r.ApiService.TeamsGetGroupExecute(r)
}

/*
TeamsGetGroup Get group from teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsGetGroupRequest
*/
func (a *TeamsGroupApiService) TeamsGetGroup(ctx _context.Context, teamId string) ApiTeamsGetGroupRequest {
	return ApiTeamsGetGroupRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphGroup
func (a *TeamsGroupApiService) TeamsGetGroupExecute(r ApiTeamsGetGroupRequest) (MicrosoftGraphGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsGroupApiService.TeamsGetGroup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/group"
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

type ApiTeamsGetRefGroupRequest struct {
	ctx _context.Context
	ApiService *TeamsGroupApiService
	teamId string
}


func (r ApiTeamsGetRefGroupRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.TeamsGetRefGroupExecute(r)
}

/*
TeamsGetRefGroup Get ref of group from teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsGetRefGroupRequest
*/
func (a *TeamsGroupApiService) TeamsGetRefGroup(ctx _context.Context, teamId string) ApiTeamsGetRefGroupRequest {
	return ApiTeamsGetRefGroupRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return string
func (a *TeamsGroupApiService) TeamsGetRefGroupExecute(r ApiTeamsGetRefGroupRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsGroupApiService.TeamsGetRefGroup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/group/$ref"
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

type ApiTeamsUpdateRefGroupRequest struct {
	ctx _context.Context
	ApiService *TeamsGroupApiService
	teamId string
	requestBody *map[string]map[string]interface{}
}

// New navigation property ref values
func (r ApiTeamsUpdateRefGroupRequest) RequestBody(requestBody map[string]map[string]interface{}) ApiTeamsUpdateRefGroupRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiTeamsUpdateRefGroupRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsUpdateRefGroupExecute(r)
}

/*
TeamsUpdateRefGroup Update the ref of navigation property group in teams

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsUpdateRefGroupRequest
*/
func (a *TeamsGroupApiService) TeamsUpdateRefGroup(ctx _context.Context, teamId string) ApiTeamsUpdateRefGroupRequest {
	return ApiTeamsUpdateRefGroupRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
func (a *TeamsGroupApiService) TeamsUpdateRefGroupExecute(r ApiTeamsUpdateRefGroupRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsGroupApiService.TeamsUpdateRefGroup")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/group/$ref"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.requestBody == nil {
		return nil, reportError("requestBody is required and must be specified")
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
