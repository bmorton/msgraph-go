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

// MeManagedAppRegistrationApiService MeManagedAppRegistrationApi service
type MeManagedAppRegistrationApiService service

type ApiMeCreateRefManagedAppRegistrationsRequest struct {
	ctx _context.Context
	ApiService *MeManagedAppRegistrationApiService
	requestBody *map[string]map[string]interface{}
}

// New navigation property ref value
func (r ApiMeCreateRefManagedAppRegistrationsRequest) RequestBody(requestBody map[string]map[string]interface{}) ApiMeCreateRefManagedAppRegistrationsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiMeCreateRefManagedAppRegistrationsRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.MeCreateRefManagedAppRegistrationsExecute(r)
}

/*
MeCreateRefManagedAppRegistrations Create new navigation property ref to managedAppRegistrations for me

Zero or more managed app registrations that belong to the user.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeCreateRefManagedAppRegistrationsRequest
*/
func (a *MeManagedAppRegistrationApiService) MeCreateRefManagedAppRegistrations(ctx _context.Context) ApiMeCreateRefManagedAppRegistrationsRequest {
	return ApiMeCreateRefManagedAppRegistrationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *MeManagedAppRegistrationApiService) MeCreateRefManagedAppRegistrationsExecute(r ApiMeCreateRefManagedAppRegistrationsRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeManagedAppRegistrationApiService.MeCreateRefManagedAppRegistrations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/managedAppRegistrations/$ref"

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

type ApiMeListManagedAppRegistrationsRequest struct {
	ctx _context.Context
	ApiService *MeManagedAppRegistrationApiService
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
func (r ApiMeListManagedAppRegistrationsRequest) Top(top int32) ApiMeListManagedAppRegistrationsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiMeListManagedAppRegistrationsRequest) Skip(skip int32) ApiMeListManagedAppRegistrationsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiMeListManagedAppRegistrationsRequest) Search(search string) ApiMeListManagedAppRegistrationsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiMeListManagedAppRegistrationsRequest) Filter(filter string) ApiMeListManagedAppRegistrationsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiMeListManagedAppRegistrationsRequest) Count(count bool) ApiMeListManagedAppRegistrationsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiMeListManagedAppRegistrationsRequest) Orderby(orderby []string) ApiMeListManagedAppRegistrationsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiMeListManagedAppRegistrationsRequest) Select_(select_ []string) ApiMeListManagedAppRegistrationsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiMeListManagedAppRegistrationsRequest) Expand(expand []string) ApiMeListManagedAppRegistrationsRequest {
	r.expand = &expand
	return r
}

func (r ApiMeListManagedAppRegistrationsRequest) Execute() (CollectionOfManagedAppRegistration, *_nethttp.Response, error) {
	return r.ApiService.MeListManagedAppRegistrationsExecute(r)
}

/*
MeListManagedAppRegistrations Get managedAppRegistrations from me

Zero or more managed app registrations that belong to the user.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeListManagedAppRegistrationsRequest
*/
func (a *MeManagedAppRegistrationApiService) MeListManagedAppRegistrations(ctx _context.Context) ApiMeListManagedAppRegistrationsRequest {
	return ApiMeListManagedAppRegistrationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfManagedAppRegistration
func (a *MeManagedAppRegistrationApiService) MeListManagedAppRegistrationsExecute(r ApiMeListManagedAppRegistrationsRequest) (CollectionOfManagedAppRegistration, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfManagedAppRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeManagedAppRegistrationApiService.MeListManagedAppRegistrations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/managedAppRegistrations"

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

type ApiMeListRefManagedAppRegistrationsRequest struct {
	ctx _context.Context
	ApiService *MeManagedAppRegistrationApiService
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	orderby *[]string
}

// Show only the first n items
func (r ApiMeListRefManagedAppRegistrationsRequest) Top(top int32) ApiMeListRefManagedAppRegistrationsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiMeListRefManagedAppRegistrationsRequest) Skip(skip int32) ApiMeListRefManagedAppRegistrationsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiMeListRefManagedAppRegistrationsRequest) Search(search string) ApiMeListRefManagedAppRegistrationsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiMeListRefManagedAppRegistrationsRequest) Filter(filter string) ApiMeListRefManagedAppRegistrationsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiMeListRefManagedAppRegistrationsRequest) Count(count bool) ApiMeListRefManagedAppRegistrationsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiMeListRefManagedAppRegistrationsRequest) Orderby(orderby []string) ApiMeListRefManagedAppRegistrationsRequest {
	r.orderby = &orderby
	return r
}

func (r ApiMeListRefManagedAppRegistrationsRequest) Execute() (CollectionOfLinksOfManagedAppRegistration, *_nethttp.Response, error) {
	return r.ApiService.MeListRefManagedAppRegistrationsExecute(r)
}

/*
MeListRefManagedAppRegistrations Get ref of managedAppRegistrations from me

Zero or more managed app registrations that belong to the user.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeListRefManagedAppRegistrationsRequest
*/
func (a *MeManagedAppRegistrationApiService) MeListRefManagedAppRegistrations(ctx _context.Context) ApiMeListRefManagedAppRegistrationsRequest {
	return ApiMeListRefManagedAppRegistrationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfLinksOfManagedAppRegistration
func (a *MeManagedAppRegistrationApiService) MeListRefManagedAppRegistrationsExecute(r ApiMeListRefManagedAppRegistrationsRequest) (CollectionOfLinksOfManagedAppRegistration, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfLinksOfManagedAppRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeManagedAppRegistrationApiService.MeListRefManagedAppRegistrations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/managedAppRegistrations/$ref"

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
