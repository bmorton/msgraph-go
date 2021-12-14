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

// InvitationsInvitationApiService InvitationsInvitationApi service
type InvitationsInvitationApiService service

type ApiInvitationsInvitationCreateInvitationRequest struct {
	ctx _context.Context
	ApiService *InvitationsInvitationApiService
	microsoftGraphInvitation *MicrosoftGraphInvitation
}

// New entity
func (r ApiInvitationsInvitationCreateInvitationRequest) MicrosoftGraphInvitation(microsoftGraphInvitation MicrosoftGraphInvitation) ApiInvitationsInvitationCreateInvitationRequest {
	r.microsoftGraphInvitation = &microsoftGraphInvitation
	return r
}

func (r ApiInvitationsInvitationCreateInvitationRequest) Execute() (MicrosoftGraphInvitation, *_nethttp.Response, error) {
	return r.ApiService.InvitationsInvitationCreateInvitationExecute(r)
}

/*
InvitationsInvitationCreateInvitation Add new entity to invitations

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInvitationsInvitationCreateInvitationRequest
*/
func (a *InvitationsInvitationApiService) InvitationsInvitationCreateInvitation(ctx _context.Context) ApiInvitationsInvitationCreateInvitationRequest {
	return ApiInvitationsInvitationCreateInvitationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphInvitation
func (a *InvitationsInvitationApiService) InvitationsInvitationCreateInvitationExecute(r ApiInvitationsInvitationCreateInvitationRequest) (MicrosoftGraphInvitation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvitationsInvitationApiService.InvitationsInvitationCreateInvitation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invitations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphInvitation == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphInvitation is required and must be specified")
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
	localVarPostBody = r.microsoftGraphInvitation
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

type ApiInvitationsInvitationDeleteInvitationRequest struct {
	ctx _context.Context
	ApiService *InvitationsInvitationApiService
	invitationId string
	ifMatch *string
}

// ETag
func (r ApiInvitationsInvitationDeleteInvitationRequest) IfMatch(ifMatch string) ApiInvitationsInvitationDeleteInvitationRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiInvitationsInvitationDeleteInvitationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InvitationsInvitationDeleteInvitationExecute(r)
}

/*
InvitationsInvitationDeleteInvitation Delete entity from invitations

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invitationId key: id of invitation
 @return ApiInvitationsInvitationDeleteInvitationRequest
*/
func (a *InvitationsInvitationApiService) InvitationsInvitationDeleteInvitation(ctx _context.Context, invitationId string) ApiInvitationsInvitationDeleteInvitationRequest {
	return ApiInvitationsInvitationDeleteInvitationRequest{
		ApiService: a,
		ctx: ctx,
		invitationId: invitationId,
	}
}

// Execute executes the request
func (a *InvitationsInvitationApiService) InvitationsInvitationDeleteInvitationExecute(r ApiInvitationsInvitationDeleteInvitationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvitationsInvitationApiService.InvitationsInvitationDeleteInvitation")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invitations/{invitation-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"invitation-id"+"}", _neturl.PathEscape(parameterToString(r.invitationId, "")), -1)

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

type ApiInvitationsInvitationGetInvitationRequest struct {
	ctx _context.Context
	ApiService *InvitationsInvitationApiService
	invitationId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiInvitationsInvitationGetInvitationRequest) Select_(select_ []string) ApiInvitationsInvitationGetInvitationRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiInvitationsInvitationGetInvitationRequest) Expand(expand []string) ApiInvitationsInvitationGetInvitationRequest {
	r.expand = &expand
	return r
}

func (r ApiInvitationsInvitationGetInvitationRequest) Execute() (MicrosoftGraphInvitation, *_nethttp.Response, error) {
	return r.ApiService.InvitationsInvitationGetInvitationExecute(r)
}

/*
InvitationsInvitationGetInvitation Get entity from invitations by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invitationId key: id of invitation
 @return ApiInvitationsInvitationGetInvitationRequest
*/
func (a *InvitationsInvitationApiService) InvitationsInvitationGetInvitation(ctx _context.Context, invitationId string) ApiInvitationsInvitationGetInvitationRequest {
	return ApiInvitationsInvitationGetInvitationRequest{
		ApiService: a,
		ctx: ctx,
		invitationId: invitationId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphInvitation
func (a *InvitationsInvitationApiService) InvitationsInvitationGetInvitationExecute(r ApiInvitationsInvitationGetInvitationRequest) (MicrosoftGraphInvitation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvitationsInvitationApiService.InvitationsInvitationGetInvitation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invitations/{invitation-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"invitation-id"+"}", _neturl.PathEscape(parameterToString(r.invitationId, "")), -1)

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

type ApiInvitationsInvitationListInvitationRequest struct {
	ctx _context.Context
	ApiService *InvitationsInvitationApiService
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
func (r ApiInvitationsInvitationListInvitationRequest) Top(top int32) ApiInvitationsInvitationListInvitationRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiInvitationsInvitationListInvitationRequest) Skip(skip int32) ApiInvitationsInvitationListInvitationRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiInvitationsInvitationListInvitationRequest) Search(search string) ApiInvitationsInvitationListInvitationRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiInvitationsInvitationListInvitationRequest) Filter(filter string) ApiInvitationsInvitationListInvitationRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiInvitationsInvitationListInvitationRequest) Count(count bool) ApiInvitationsInvitationListInvitationRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiInvitationsInvitationListInvitationRequest) Orderby(orderby []string) ApiInvitationsInvitationListInvitationRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiInvitationsInvitationListInvitationRequest) Select_(select_ []string) ApiInvitationsInvitationListInvitationRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiInvitationsInvitationListInvitationRequest) Expand(expand []string) ApiInvitationsInvitationListInvitationRequest {
	r.expand = &expand
	return r
}

func (r ApiInvitationsInvitationListInvitationRequest) Execute() (CollectionOfInvitation, *_nethttp.Response, error) {
	return r.ApiService.InvitationsInvitationListInvitationExecute(r)
}

/*
InvitationsInvitationListInvitation Get entities from invitations

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInvitationsInvitationListInvitationRequest
*/
func (a *InvitationsInvitationApiService) InvitationsInvitationListInvitation(ctx _context.Context) ApiInvitationsInvitationListInvitationRequest {
	return ApiInvitationsInvitationListInvitationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfInvitation
func (a *InvitationsInvitationApiService) InvitationsInvitationListInvitationExecute(r ApiInvitationsInvitationListInvitationRequest) (CollectionOfInvitation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfInvitation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvitationsInvitationApiService.InvitationsInvitationListInvitation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invitations"

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

type ApiInvitationsInvitationUpdateInvitationRequest struct {
	ctx _context.Context
	ApiService *InvitationsInvitationApiService
	invitationId string
	microsoftGraphInvitation *MicrosoftGraphInvitation
}

// New property values
func (r ApiInvitationsInvitationUpdateInvitationRequest) MicrosoftGraphInvitation(microsoftGraphInvitation MicrosoftGraphInvitation) ApiInvitationsInvitationUpdateInvitationRequest {
	r.microsoftGraphInvitation = &microsoftGraphInvitation
	return r
}

func (r ApiInvitationsInvitationUpdateInvitationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.InvitationsInvitationUpdateInvitationExecute(r)
}

/*
InvitationsInvitationUpdateInvitation Update entity in invitations

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param invitationId key: id of invitation
 @return ApiInvitationsInvitationUpdateInvitationRequest
*/
func (a *InvitationsInvitationApiService) InvitationsInvitationUpdateInvitation(ctx _context.Context, invitationId string) ApiInvitationsInvitationUpdateInvitationRequest {
	return ApiInvitationsInvitationUpdateInvitationRequest{
		ApiService: a,
		ctx: ctx,
		invitationId: invitationId,
	}
}

// Execute executes the request
func (a *InvitationsInvitationApiService) InvitationsInvitationUpdateInvitationExecute(r ApiInvitationsInvitationUpdateInvitationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvitationsInvitationApiService.InvitationsInvitationUpdateInvitation")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invitations/{invitation-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"invitation-id"+"}", _neturl.PathEscape(parameterToString(r.invitationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphInvitation == nil {
		return nil, reportError("microsoftGraphInvitation is required and must be specified")
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
	localVarPostBody = r.microsoftGraphInvitation
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
