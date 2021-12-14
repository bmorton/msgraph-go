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

// BrandingOrganizationalBrandingLocalizationApiService BrandingOrganizationalBrandingLocalizationApi service
type BrandingOrganizationalBrandingLocalizationApiService service

type ApiBrandingCreateLocalizationsRequest struct {
	ctx _context.Context
	ApiService *BrandingOrganizationalBrandingLocalizationApiService
	microsoftGraphOrganizationalBrandingLocalization *MicrosoftGraphOrganizationalBrandingLocalization
}

// New navigation property
func (r ApiBrandingCreateLocalizationsRequest) MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization MicrosoftGraphOrganizationalBrandingLocalization) ApiBrandingCreateLocalizationsRequest {
	r.microsoftGraphOrganizationalBrandingLocalization = &microsoftGraphOrganizationalBrandingLocalization
	return r
}

func (r ApiBrandingCreateLocalizationsRequest) Execute() (MicrosoftGraphOrganizationalBrandingLocalization, *_nethttp.Response, error) {
	return r.ApiService.BrandingCreateLocalizationsExecute(r)
}

/*
BrandingCreateLocalizations Create new navigation property to localizations for branding

Add different branding based on a locale.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBrandingCreateLocalizationsRequest
*/
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingCreateLocalizations(ctx _context.Context) ApiBrandingCreateLocalizationsRequest {
	return ApiBrandingCreateLocalizationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphOrganizationalBrandingLocalization
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingCreateLocalizationsExecute(r ApiBrandingCreateLocalizationsRequest) (MicrosoftGraphOrganizationalBrandingLocalization, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphOrganizationalBrandingLocalization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingOrganizationalBrandingLocalizationApiService.BrandingCreateLocalizations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/branding/localizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphOrganizationalBrandingLocalization == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphOrganizationalBrandingLocalization is required and must be specified")
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
	localVarPostBody = r.microsoftGraphOrganizationalBrandingLocalization
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

type ApiBrandingDeleteLocalizationsRequest struct {
	ctx _context.Context
	ApiService *BrandingOrganizationalBrandingLocalizationApiService
	organizationalBrandingLocalizationId string
	ifMatch *string
}

// ETag
func (r ApiBrandingDeleteLocalizationsRequest) IfMatch(ifMatch string) ApiBrandingDeleteLocalizationsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiBrandingDeleteLocalizationsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.BrandingDeleteLocalizationsExecute(r)
}

/*
BrandingDeleteLocalizations Delete navigation property localizations for branding

Add different branding based on a locale.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationalBrandingLocalizationId key: id of organizationalBrandingLocalization
 @return ApiBrandingDeleteLocalizationsRequest
*/
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingDeleteLocalizations(ctx _context.Context, organizationalBrandingLocalizationId string) ApiBrandingDeleteLocalizationsRequest {
	return ApiBrandingDeleteLocalizationsRequest{
		ApiService: a,
		ctx: ctx,
		organizationalBrandingLocalizationId: organizationalBrandingLocalizationId,
	}
}

// Execute executes the request
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingDeleteLocalizationsExecute(r ApiBrandingDeleteLocalizationsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingOrganizationalBrandingLocalizationApiService.BrandingDeleteLocalizations")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/branding/localizations/{organizationalBrandingLocalization-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationalBrandingLocalization-id"+"}", _neturl.PathEscape(parameterToString(r.organizationalBrandingLocalizationId, "")), -1)

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

type ApiBrandingGetLocalizationsRequest struct {
	ctx _context.Context
	ApiService *BrandingOrganizationalBrandingLocalizationApiService
	organizationalBrandingLocalizationId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiBrandingGetLocalizationsRequest) Select_(select_ []string) ApiBrandingGetLocalizationsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiBrandingGetLocalizationsRequest) Expand(expand []string) ApiBrandingGetLocalizationsRequest {
	r.expand = &expand
	return r
}

func (r ApiBrandingGetLocalizationsRequest) Execute() (MicrosoftGraphOrganizationalBrandingLocalization, *_nethttp.Response, error) {
	return r.ApiService.BrandingGetLocalizationsExecute(r)
}

/*
BrandingGetLocalizations Get localizations from branding

Add different branding based on a locale.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationalBrandingLocalizationId key: id of organizationalBrandingLocalization
 @return ApiBrandingGetLocalizationsRequest
*/
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingGetLocalizations(ctx _context.Context, organizationalBrandingLocalizationId string) ApiBrandingGetLocalizationsRequest {
	return ApiBrandingGetLocalizationsRequest{
		ApiService: a,
		ctx: ctx,
		organizationalBrandingLocalizationId: organizationalBrandingLocalizationId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphOrganizationalBrandingLocalization
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingGetLocalizationsExecute(r ApiBrandingGetLocalizationsRequest) (MicrosoftGraphOrganizationalBrandingLocalization, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphOrganizationalBrandingLocalization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingOrganizationalBrandingLocalizationApiService.BrandingGetLocalizations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/branding/localizations/{organizationalBrandingLocalization-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationalBrandingLocalization-id"+"}", _neturl.PathEscape(parameterToString(r.organizationalBrandingLocalizationId, "")), -1)

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

type ApiBrandingListLocalizationsRequest struct {
	ctx _context.Context
	ApiService *BrandingOrganizationalBrandingLocalizationApiService
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
func (r ApiBrandingListLocalizationsRequest) Top(top int32) ApiBrandingListLocalizationsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiBrandingListLocalizationsRequest) Skip(skip int32) ApiBrandingListLocalizationsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiBrandingListLocalizationsRequest) Search(search string) ApiBrandingListLocalizationsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiBrandingListLocalizationsRequest) Filter(filter string) ApiBrandingListLocalizationsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiBrandingListLocalizationsRequest) Count(count bool) ApiBrandingListLocalizationsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiBrandingListLocalizationsRequest) Orderby(orderby []string) ApiBrandingListLocalizationsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiBrandingListLocalizationsRequest) Select_(select_ []string) ApiBrandingListLocalizationsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiBrandingListLocalizationsRequest) Expand(expand []string) ApiBrandingListLocalizationsRequest {
	r.expand = &expand
	return r
}

func (r ApiBrandingListLocalizationsRequest) Execute() (CollectionOfOrganizationalBrandingLocalization, *_nethttp.Response, error) {
	return r.ApiService.BrandingListLocalizationsExecute(r)
}

/*
BrandingListLocalizations Get localizations from branding

Add different branding based on a locale.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBrandingListLocalizationsRequest
*/
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingListLocalizations(ctx _context.Context) ApiBrandingListLocalizationsRequest {
	return ApiBrandingListLocalizationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfOrganizationalBrandingLocalization
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingListLocalizationsExecute(r ApiBrandingListLocalizationsRequest) (CollectionOfOrganizationalBrandingLocalization, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfOrganizationalBrandingLocalization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingOrganizationalBrandingLocalizationApiService.BrandingListLocalizations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/branding/localizations"

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

type ApiBrandingUpdateLocalizationsRequest struct {
	ctx _context.Context
	ApiService *BrandingOrganizationalBrandingLocalizationApiService
	organizationalBrandingLocalizationId string
	microsoftGraphOrganizationalBrandingLocalization *MicrosoftGraphOrganizationalBrandingLocalization
}

// New navigation property values
func (r ApiBrandingUpdateLocalizationsRequest) MicrosoftGraphOrganizationalBrandingLocalization(microsoftGraphOrganizationalBrandingLocalization MicrosoftGraphOrganizationalBrandingLocalization) ApiBrandingUpdateLocalizationsRequest {
	r.microsoftGraphOrganizationalBrandingLocalization = &microsoftGraphOrganizationalBrandingLocalization
	return r
}

func (r ApiBrandingUpdateLocalizationsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.BrandingUpdateLocalizationsExecute(r)
}

/*
BrandingUpdateLocalizations Update the navigation property localizations in branding

Add different branding based on a locale.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationalBrandingLocalizationId key: id of organizationalBrandingLocalization
 @return ApiBrandingUpdateLocalizationsRequest
*/
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingUpdateLocalizations(ctx _context.Context, organizationalBrandingLocalizationId string) ApiBrandingUpdateLocalizationsRequest {
	return ApiBrandingUpdateLocalizationsRequest{
		ApiService: a,
		ctx: ctx,
		organizationalBrandingLocalizationId: organizationalBrandingLocalizationId,
	}
}

// Execute executes the request
func (a *BrandingOrganizationalBrandingLocalizationApiService) BrandingUpdateLocalizationsExecute(r ApiBrandingUpdateLocalizationsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BrandingOrganizationalBrandingLocalizationApiService.BrandingUpdateLocalizations")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/branding/localizations/{organizationalBrandingLocalization-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationalBrandingLocalization-id"+"}", _neturl.PathEscape(parameterToString(r.organizationalBrandingLocalizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphOrganizationalBrandingLocalization == nil {
		return nil, reportError("microsoftGraphOrganizationalBrandingLocalization is required and must be specified")
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
	localVarPostBody = r.microsoftGraphOrganizationalBrandingLocalization
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
