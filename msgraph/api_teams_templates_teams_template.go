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

// TeamsTemplatesTeamsTemplateApiService TeamsTemplatesTeamsTemplateApi service
type TeamsTemplatesTeamsTemplateApiService service

type ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest struct {
	ctx _context.Context
	ApiService *TeamsTemplatesTeamsTemplateApiService
	microsoftGraphTeamsTemplate *MicrosoftGraphTeamsTemplate
}

// New entity
func (r ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest) MicrosoftGraphTeamsTemplate(microsoftGraphTeamsTemplate MicrosoftGraphTeamsTemplate) ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest {
	r.microsoftGraphTeamsTemplate = &microsoftGraphTeamsTemplate
	return r
}

func (r ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest) Execute() (MicrosoftGraphTeamsTemplate, *_nethttp.Response, error) {
	return r.ApiService.TeamsTemplatesTeamsTemplateCreateTeamsTemplateExecute(r)
}

/*
TeamsTemplatesTeamsTemplateCreateTeamsTemplate Add new entity to teamsTemplates

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest
*/
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateCreateTeamsTemplate(ctx _context.Context) ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest {
	return ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeamsTemplate
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateCreateTeamsTemplateExecute(r ApiTeamsTemplatesTeamsTemplateCreateTeamsTemplateRequest) (MicrosoftGraphTeamsTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeamsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTemplatesTeamsTemplateApiService.TeamsTemplatesTeamsTemplateCreateTeamsTemplate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teamsTemplates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphTeamsTemplate == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphTeamsTemplate is required and must be specified")
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
	localVarPostBody = r.microsoftGraphTeamsTemplate
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

type ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest struct {
	ctx _context.Context
	ApiService *TeamsTemplatesTeamsTemplateApiService
	teamsTemplateId string
	ifMatch *string
}

// ETag
func (r ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest) IfMatch(ifMatch string) ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsTemplatesTeamsTemplateDeleteTeamsTemplateExecute(r)
}

/*
TeamsTemplatesTeamsTemplateDeleteTeamsTemplate Delete entity from teamsTemplates

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamsTemplateId key: id of teamsTemplate
 @return ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest
*/
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateDeleteTeamsTemplate(ctx _context.Context, teamsTemplateId string) ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest {
	return ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest{
		ApiService: a,
		ctx: ctx,
		teamsTemplateId: teamsTemplateId,
	}
}

// Execute executes the request
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateDeleteTeamsTemplateExecute(r ApiTeamsTemplatesTeamsTemplateDeleteTeamsTemplateRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTemplatesTeamsTemplateApiService.TeamsTemplatesTeamsTemplateDeleteTeamsTemplate")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teamsTemplates/{teamsTemplate-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"teamsTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.teamsTemplateId, "")), -1)

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

type ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest struct {
	ctx _context.Context
	ApiService *TeamsTemplatesTeamsTemplateApiService
	teamsTemplateId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest) Select_(select_ []string) ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest) Expand(expand []string) ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest) Execute() (MicrosoftGraphTeamsTemplate, *_nethttp.Response, error) {
	return r.ApiService.TeamsTemplatesTeamsTemplateGetTeamsTemplateExecute(r)
}

/*
TeamsTemplatesTeamsTemplateGetTeamsTemplate Get entity from teamsTemplates by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamsTemplateId key: id of teamsTemplate
 @return ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest
*/
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateGetTeamsTemplate(ctx _context.Context, teamsTemplateId string) ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest {
	return ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest{
		ApiService: a,
		ctx: ctx,
		teamsTemplateId: teamsTemplateId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphTeamsTemplate
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateGetTeamsTemplateExecute(r ApiTeamsTemplatesTeamsTemplateGetTeamsTemplateRequest) (MicrosoftGraphTeamsTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTeamsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTemplatesTeamsTemplateApiService.TeamsTemplatesTeamsTemplateGetTeamsTemplate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teamsTemplates/{teamsTemplate-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"teamsTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.teamsTemplateId, "")), -1)

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

type ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest struct {
	ctx _context.Context
	ApiService *TeamsTemplatesTeamsTemplateApiService
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
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Top(top int32) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Skip(skip int32) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Search(search string) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Filter(filter string) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Count(count bool) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Orderby(orderby []string) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Select_(select_ []string) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Expand(expand []string) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	r.expand = &expand
	return r
}

func (r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) Execute() (CollectionOfTeamsTemplate, *_nethttp.Response, error) {
	return r.ApiService.TeamsTemplatesTeamsTemplateListTeamsTemplateExecute(r)
}

/*
TeamsTemplatesTeamsTemplateListTeamsTemplate Get entities from teamsTemplates

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest
*/
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateListTeamsTemplate(ctx _context.Context) ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest {
	return ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfTeamsTemplate
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateListTeamsTemplateExecute(r ApiTeamsTemplatesTeamsTemplateListTeamsTemplateRequest) (CollectionOfTeamsTemplate, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfTeamsTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTemplatesTeamsTemplateApiService.TeamsTemplatesTeamsTemplateListTeamsTemplate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teamsTemplates"

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

type ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest struct {
	ctx _context.Context
	ApiService *TeamsTemplatesTeamsTemplateApiService
	teamsTemplateId string
	microsoftGraphTeamsTemplate *MicrosoftGraphTeamsTemplate
}

// New property values
func (r ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest) MicrosoftGraphTeamsTemplate(microsoftGraphTeamsTemplate MicrosoftGraphTeamsTemplate) ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest {
	r.microsoftGraphTeamsTemplate = &microsoftGraphTeamsTemplate
	return r
}

func (r ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TeamsTemplatesTeamsTemplateUpdateTeamsTemplateExecute(r)
}

/*
TeamsTemplatesTeamsTemplateUpdateTeamsTemplate Update entity in teamsTemplates

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamsTemplateId key: id of teamsTemplate
 @return ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest
*/
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateUpdateTeamsTemplate(ctx _context.Context, teamsTemplateId string) ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest {
	return ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest{
		ApiService: a,
		ctx: ctx,
		teamsTemplateId: teamsTemplateId,
	}
}

// Execute executes the request
func (a *TeamsTemplatesTeamsTemplateApiService) TeamsTemplatesTeamsTemplateUpdateTeamsTemplateExecute(r ApiTeamsTemplatesTeamsTemplateUpdateTeamsTemplateRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsTemplatesTeamsTemplateApiService.TeamsTemplatesTeamsTemplateUpdateTeamsTemplate")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teamsTemplates/{teamsTemplate-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"teamsTemplate-id"+"}", _neturl.PathEscape(parameterToString(r.teamsTemplateId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphTeamsTemplate == nil {
		return nil, reportError("microsoftGraphTeamsTemplate is required and must be specified")
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
	localVarPostBody = r.microsoftGraphTeamsTemplate
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