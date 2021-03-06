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
)

// Linger please
var (
	_ _context.Context
)

// MeOAuth2PermissionGrantApiService MeOAuth2PermissionGrantApi service
type MeOAuth2PermissionGrantApiService service

type ApiMeCreateRefOauth2PermissionGrantsRequest struct {
	ctx _context.Context
	ApiService *MeOAuth2PermissionGrantApiService
	requestBody *map[string]map[string]interface{}
}

// New navigation property ref value
func (r ApiMeCreateRefOauth2PermissionGrantsRequest) RequestBody(requestBody map[string]map[string]interface{}) ApiMeCreateRefOauth2PermissionGrantsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiMeCreateRefOauth2PermissionGrantsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.MeCreateRefOauth2PermissionGrantsExecute(r)
}

/*
MeCreateRefOauth2PermissionGrants Create new navigation property ref to oauth2PermissionGrants for me

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeCreateRefOauth2PermissionGrantsRequest
*/
func (a *MeOAuth2PermissionGrantApiService) MeCreateRefOauth2PermissionGrants(ctx _context.Context) ApiMeCreateRefOauth2PermissionGrantsRequest {
	return ApiMeCreateRefOauth2PermissionGrantsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *MeOAuth2PermissionGrantApiService) MeCreateRefOauth2PermissionGrantsExecute(r ApiMeCreateRefOauth2PermissionGrantsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeOAuth2PermissionGrantApiService.MeCreateRefOauth2PermissionGrants")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/oauth2PermissionGrants/$ref"

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

type ApiMeListOauth2PermissionGrantsRequest struct {
	ctx _context.Context
	ApiService *MeOAuth2PermissionGrantApiService
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
func (r ApiMeListOauth2PermissionGrantsRequest) Top(top int32) ApiMeListOauth2PermissionGrantsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiMeListOauth2PermissionGrantsRequest) Skip(skip int32) ApiMeListOauth2PermissionGrantsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiMeListOauth2PermissionGrantsRequest) Search(search string) ApiMeListOauth2PermissionGrantsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiMeListOauth2PermissionGrantsRequest) Filter(filter string) ApiMeListOauth2PermissionGrantsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiMeListOauth2PermissionGrantsRequest) Count(count bool) ApiMeListOauth2PermissionGrantsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiMeListOauth2PermissionGrantsRequest) Orderby(orderby []string) ApiMeListOauth2PermissionGrantsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiMeListOauth2PermissionGrantsRequest) Select_(select_ []string) ApiMeListOauth2PermissionGrantsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiMeListOauth2PermissionGrantsRequest) Expand(expand []string) ApiMeListOauth2PermissionGrantsRequest {
	r.expand = &expand
	return r
}

func (r ApiMeListOauth2PermissionGrantsRequest) Execute() (CollectionOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	return r.ApiService.MeListOauth2PermissionGrantsExecute(r)
}

/*
MeListOauth2PermissionGrants Get oauth2PermissionGrants from me

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeListOauth2PermissionGrantsRequest
*/
func (a *MeOAuth2PermissionGrantApiService) MeListOauth2PermissionGrants(ctx _context.Context) ApiMeListOauth2PermissionGrantsRequest {
	return ApiMeListOauth2PermissionGrantsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfOAuth2PermissionGrant
func (a *MeOAuth2PermissionGrantApiService) MeListOauth2PermissionGrantsExecute(r ApiMeListOauth2PermissionGrantsRequest) (CollectionOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfOAuth2PermissionGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeOAuth2PermissionGrantApiService.MeListOauth2PermissionGrants")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/oauth2PermissionGrants"

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

type ApiMeListRefOauth2PermissionGrantsRequest struct {
	ctx _context.Context
	ApiService *MeOAuth2PermissionGrantApiService
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
}

// Show only the first n items
func (r ApiMeListRefOauth2PermissionGrantsRequest) Top(top int32) ApiMeListRefOauth2PermissionGrantsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiMeListRefOauth2PermissionGrantsRequest) Skip(skip int32) ApiMeListRefOauth2PermissionGrantsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiMeListRefOauth2PermissionGrantsRequest) Search(search string) ApiMeListRefOauth2PermissionGrantsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiMeListRefOauth2PermissionGrantsRequest) Filter(filter string) ApiMeListRefOauth2PermissionGrantsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiMeListRefOauth2PermissionGrantsRequest) Count(count bool) ApiMeListRefOauth2PermissionGrantsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiMeListRefOauth2PermissionGrantsRequest) Orderby(orderby []string) ApiMeListRefOauth2PermissionGrantsRequest {
	r.orderby = &orderby
	return r
}

func (r ApiMeListRefOauth2PermissionGrantsRequest) Execute() (CollectionOfLinksOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	return r.ApiService.MeListRefOauth2PermissionGrantsExecute(r)
}

/*
MeListRefOauth2PermissionGrants Get ref of oauth2PermissionGrants from me

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeListRefOauth2PermissionGrantsRequest
*/
func (a *MeOAuth2PermissionGrantApiService) MeListRefOauth2PermissionGrants(ctx _context.Context) ApiMeListRefOauth2PermissionGrantsRequest {
	return ApiMeListRefOauth2PermissionGrantsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfLinksOfOAuth2PermissionGrant
func (a *MeOAuth2PermissionGrantApiService) MeListRefOauth2PermissionGrantsExecute(r ApiMeListRefOauth2PermissionGrantsRequest) (CollectionOfLinksOfOAuth2PermissionGrant, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfLinksOfOAuth2PermissionGrant
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeOAuth2PermissionGrantApiService.MeListRefOauth2PermissionGrants")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/oauth2PermissionGrants/$ref"

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
