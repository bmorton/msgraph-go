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

// TeamsTeamsAsyncOperationApiService TeamsTeamsAsyncOperationApi service
type TeamsTeamsAsyncOperationApiService service

type ApiTeamsCreateOperationsRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamsAsyncOperationApiService
	teamId string
	microsoftGraphTeamsAsyncOperation *MicrosoftGraphTeamsAsyncOperation
}

// New navigation property
func (r ApiTeamsCreateOperationsRequest) MicrosoftGraphTeamsAsyncOperation(microsoftGraphTeamsAsyncOperation MicrosoftGraphTeamsAsyncOperation) ApiTeamsCreateOperationsRequest {
	r.microsoftGraphTeamsAsyncOperation = &microsoftGraphTeamsAsyncOperation
	return r
}

func (r ApiTeamsCreateOperationsRequest) Execute() (MicrosoftGraphTeamsAsyncOperation, *_nethttp.Response, error) {
	return r.ApiService.TeamsCreateOperationsExecute(r)
}

/*
TeamsCreateOperations Create new navigation property to operations for teams

The async operations that ran or are running on this team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsCreateOperationsRequest
*/
func (a *TeamsTeamsAsyncOperationApiService) TeamsCreateOperations(ctx _context.Context, teamId string) ApiTeamsCreateOperationsRequest {
	return ApiTeamsCreateOperationsRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeamsAsyncOperation
func (a *TeamsTeamsAsyncOperationApiService) TeamsCreateOperationsExecute(r ApiTeamsCreateOperationsRequest) (MicrosoftGraphTeamsAsyncOperation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeamsAsyncOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamsAsyncOperationApiService.TeamsCreateOperations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/operations"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphTeamsAsyncOperation == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphTeamsAsyncOperation is required and must be specified")
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
	localVarPostBody = r.microsoftGraphTeamsAsyncOperation
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

type ApiTeamsDeleteOperationsRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamsAsyncOperationApiService
	teamId string
	teamsAsyncOperationId string
	ifMatch *string
}

// ETag
func (r ApiTeamsDeleteOperationsRequest) IfMatch(ifMatch string) ApiTeamsDeleteOperationsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiTeamsDeleteOperationsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsDeleteOperationsExecute(r)
}

/*
TeamsDeleteOperations Delete navigation property operations for teams

The async operations that ran or are running on this team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @param teamsAsyncOperationId key: id of teamsAsyncOperation
 @return ApiTeamsDeleteOperationsRequest
*/
func (a *TeamsTeamsAsyncOperationApiService) TeamsDeleteOperations(ctx _context.Context, teamId string, teamsAsyncOperationId string) ApiTeamsDeleteOperationsRequest {
	return ApiTeamsDeleteOperationsRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		teamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// Execute executes the request
func (a *TeamsTeamsAsyncOperationApiService) TeamsDeleteOperationsExecute(r ApiTeamsDeleteOperationsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamsAsyncOperationApiService.TeamsDeleteOperations")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/operations/{teamsAsyncOperation-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamsAsyncOperation-id"+"}", _neturl.PathEscape(parameterToString(r.teamsAsyncOperationId, "")), -1)

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

type ApiTeamsGetOperationsRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamsAsyncOperationApiService
	teamId string
	teamsAsyncOperationId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiTeamsGetOperationsRequest) Select_(select_ []string) ApiTeamsGetOperationsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsGetOperationsRequest) Expand(expand []string) ApiTeamsGetOperationsRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsGetOperationsRequest) Execute() (MicrosoftGraphTeamsAsyncOperation, *_nethttp.Response, error) {
	return r.ApiService.TeamsGetOperationsExecute(r)
}

/*
TeamsGetOperations Get operations from teams

The async operations that ran or are running on this team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @param teamsAsyncOperationId key: id of teamsAsyncOperation
 @return ApiTeamsGetOperationsRequest
*/
func (a *TeamsTeamsAsyncOperationApiService) TeamsGetOperations(ctx _context.Context, teamId string, teamsAsyncOperationId string) ApiTeamsGetOperationsRequest {
	return ApiTeamsGetOperationsRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		teamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeamsAsyncOperation
func (a *TeamsTeamsAsyncOperationApiService) TeamsGetOperationsExecute(r ApiTeamsGetOperationsRequest) (MicrosoftGraphTeamsAsyncOperation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeamsAsyncOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamsAsyncOperationApiService.TeamsGetOperations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/operations/{teamsAsyncOperation-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamsAsyncOperation-id"+"}", _neturl.PathEscape(parameterToString(r.teamsAsyncOperationId, "")), -1)

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

type ApiTeamsListOperationsRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamsAsyncOperationApiService
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
func (r ApiTeamsListOperationsRequest) Top(top int32) ApiTeamsListOperationsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiTeamsListOperationsRequest) Skip(skip int32) ApiTeamsListOperationsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiTeamsListOperationsRequest) Search(search string) ApiTeamsListOperationsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiTeamsListOperationsRequest) Filter(filter string) ApiTeamsListOperationsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiTeamsListOperationsRequest) Count(count bool) ApiTeamsListOperationsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiTeamsListOperationsRequest) Orderby(orderby []string) ApiTeamsListOperationsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiTeamsListOperationsRequest) Select_(select_ []string) ApiTeamsListOperationsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsListOperationsRequest) Expand(expand []string) ApiTeamsListOperationsRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsListOperationsRequest) Execute() (CollectionOfTeamsAsyncOperation, *_nethttp.Response, error) {
	return r.ApiService.TeamsListOperationsExecute(r)
}

/*
TeamsListOperations Get operations from teams

The async operations that ran or are running on this team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @return ApiTeamsListOperationsRequest
*/
func (a *TeamsTeamsAsyncOperationApiService) TeamsListOperations(ctx _context.Context, teamId string) ApiTeamsListOperationsRequest {
	return ApiTeamsListOperationsRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return CollectionOfTeamsAsyncOperation
func (a *TeamsTeamsAsyncOperationApiService) TeamsListOperationsExecute(r ApiTeamsListOperationsRequest) (CollectionOfTeamsAsyncOperation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfTeamsAsyncOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamsAsyncOperationApiService.TeamsListOperations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/operations"
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

type ApiTeamsUpdateOperationsRequest struct {
	ctx _context.Context
	ApiService *TeamsTeamsAsyncOperationApiService
	teamId string
	teamsAsyncOperationId string
	microsoftGraphTeamsAsyncOperation *MicrosoftGraphTeamsAsyncOperation
}

// New navigation property values
func (r ApiTeamsUpdateOperationsRequest) MicrosoftGraphTeamsAsyncOperation(microsoftGraphTeamsAsyncOperation MicrosoftGraphTeamsAsyncOperation) ApiTeamsUpdateOperationsRequest {
	r.microsoftGraphTeamsAsyncOperation = &microsoftGraphTeamsAsyncOperation
	return r
}

func (r ApiTeamsUpdateOperationsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsUpdateOperationsExecute(r)
}

/*
TeamsUpdateOperations Update the navigation property operations in teams

The async operations that ran or are running on this team.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId key: id of team
 @param teamsAsyncOperationId key: id of teamsAsyncOperation
 @return ApiTeamsUpdateOperationsRequest
*/
func (a *TeamsTeamsAsyncOperationApiService) TeamsUpdateOperations(ctx _context.Context, teamId string, teamsAsyncOperationId string) ApiTeamsUpdateOperationsRequest {
	return ApiTeamsUpdateOperationsRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		teamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// Execute executes the request
func (a *TeamsTeamsAsyncOperationApiService) TeamsUpdateOperationsExecute(r ApiTeamsUpdateOperationsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTeamsAsyncOperationApiService.TeamsUpdateOperations")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team-id}/operations/{teamsAsyncOperation-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team-id"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamsAsyncOperation-id"+"}", _neturl.PathEscape(parameterToString(r.teamsAsyncOperationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphTeamsAsyncOperation == nil {
		return nil, reportError("microsoftGraphTeamsAsyncOperation is required and must be specified")
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
	localVarPostBody = r.microsoftGraphTeamsAsyncOperation
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
