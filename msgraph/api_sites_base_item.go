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

// SitesBaseItemApiService SitesBaseItemApi service
type SitesBaseItemApiService service

type ApiSitesCreateItemsRequest struct {
	ctx _context.Context
	ApiService *SitesBaseItemApiService
	siteId string
	microsoftGraphBaseItem *MicrosoftGraphBaseItem
}

// New navigation property
func (r ApiSitesCreateItemsRequest) MicrosoftGraphBaseItem(microsoftGraphBaseItem MicrosoftGraphBaseItem) ApiSitesCreateItemsRequest {
	r.microsoftGraphBaseItem = &microsoftGraphBaseItem
	return r
}

func (r ApiSitesCreateItemsRequest) Execute() (MicrosoftGraphBaseItem, *_nethttp.Response, error) {
	return r.ApiService.SitesCreateItemsExecute(r)
}

/*
SitesCreateItems Create new navigation property to items for sites

Used to address any item contained in this site. This collection can't be enumerated.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId key: id of site
 @return ApiSitesCreateItemsRequest
*/
func (a *SitesBaseItemApiService) SitesCreateItems(ctx _context.Context, siteId string) ApiSitesCreateItemsRequest {
	return ApiSitesCreateItemsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphBaseItem
func (a *SitesBaseItemApiService) SitesCreateItemsExecute(r ApiSitesCreateItemsRequest) (MicrosoftGraphBaseItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphBaseItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesBaseItemApiService.SitesCreateItems")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites/{site-id}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", _neturl.PathEscape(parameterToString(r.siteId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphBaseItem == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphBaseItem is required and must be specified")
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
	localVarPostBody = r.microsoftGraphBaseItem
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

type ApiSitesDeleteItemsRequest struct {
	ctx _context.Context
	ApiService *SitesBaseItemApiService
	siteId string
	baseItemId string
	ifMatch *string
}

// ETag
func (r ApiSitesDeleteItemsRequest) IfMatch(ifMatch string) ApiSitesDeleteItemsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiSitesDeleteItemsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SitesDeleteItemsExecute(r)
}

/*
SitesDeleteItems Delete navigation property items for sites

Used to address any item contained in this site. This collection can't be enumerated.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId key: id of site
 @param baseItemId key: id of baseItem
 @return ApiSitesDeleteItemsRequest
*/
func (a *SitesBaseItemApiService) SitesDeleteItems(ctx _context.Context, siteId string, baseItemId string) ApiSitesDeleteItemsRequest {
	return ApiSitesDeleteItemsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		baseItemId: baseItemId,
	}
}

// Execute executes the request
func (a *SitesBaseItemApiService) SitesDeleteItemsExecute(r ApiSitesDeleteItemsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesBaseItemApiService.SitesDeleteItems")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites/{site-id}/items/{baseItem-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", _neturl.PathEscape(parameterToString(r.siteId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"baseItem-id"+"}", _neturl.PathEscape(parameterToString(r.baseItemId, "")), -1)

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

type ApiSitesGetItemsRequest struct {
	ctx _context.Context
	ApiService *SitesBaseItemApiService
	siteId string
	baseItemId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiSitesGetItemsRequest) Select_(select_ []string) ApiSitesGetItemsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiSitesGetItemsRequest) Expand(expand []string) ApiSitesGetItemsRequest {
	r.expand = &expand
	return r
}

func (r ApiSitesGetItemsRequest) Execute() (MicrosoftGraphBaseItem, *_nethttp.Response, error) {
	return r.ApiService.SitesGetItemsExecute(r)
}

/*
SitesGetItems Get items from sites

Used to address any item contained in this site. This collection can't be enumerated.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId key: id of site
 @param baseItemId key: id of baseItem
 @return ApiSitesGetItemsRequest
*/
func (a *SitesBaseItemApiService) SitesGetItems(ctx _context.Context, siteId string, baseItemId string) ApiSitesGetItemsRequest {
	return ApiSitesGetItemsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		baseItemId: baseItemId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphBaseItem
func (a *SitesBaseItemApiService) SitesGetItemsExecute(r ApiSitesGetItemsRequest) (MicrosoftGraphBaseItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphBaseItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesBaseItemApiService.SitesGetItems")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites/{site-id}/items/{baseItem-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", _neturl.PathEscape(parameterToString(r.siteId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"baseItem-id"+"}", _neturl.PathEscape(parameterToString(r.baseItemId, "")), -1)

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

type ApiSitesListItemsRequest struct {
	ctx _context.Context
	ApiService *SitesBaseItemApiService
	siteId string
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
func (r ApiSitesListItemsRequest) Top(top int32) ApiSitesListItemsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiSitesListItemsRequest) Skip(skip int32) ApiSitesListItemsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiSitesListItemsRequest) Search(search string) ApiSitesListItemsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiSitesListItemsRequest) Filter(filter string) ApiSitesListItemsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiSitesListItemsRequest) Count(count bool) ApiSitesListItemsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiSitesListItemsRequest) Orderby(orderby []string) ApiSitesListItemsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiSitesListItemsRequest) Select_(select_ []string) ApiSitesListItemsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiSitesListItemsRequest) Expand(expand []string) ApiSitesListItemsRequest {
	r.expand = &expand
	return r
}

func (r ApiSitesListItemsRequest) Execute() (CollectionOfBaseItem, *_nethttp.Response, error) {
	return r.ApiService.SitesListItemsExecute(r)
}

/*
SitesListItems Get items from sites

Used to address any item contained in this site. This collection can't be enumerated.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId key: id of site
 @return ApiSitesListItemsRequest
*/
func (a *SitesBaseItemApiService) SitesListItems(ctx _context.Context, siteId string) ApiSitesListItemsRequest {
	return ApiSitesListItemsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return CollectionOfBaseItem
func (a *SitesBaseItemApiService) SitesListItemsExecute(r ApiSitesListItemsRequest) (CollectionOfBaseItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfBaseItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesBaseItemApiService.SitesListItems")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites/{site-id}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", _neturl.PathEscape(parameterToString(r.siteId, "")), -1)

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

type ApiSitesUpdateItemsRequest struct {
	ctx _context.Context
	ApiService *SitesBaseItemApiService
	siteId string
	baseItemId string
	microsoftGraphBaseItem *MicrosoftGraphBaseItem
}

// New navigation property values
func (r ApiSitesUpdateItemsRequest) MicrosoftGraphBaseItem(microsoftGraphBaseItem MicrosoftGraphBaseItem) ApiSitesUpdateItemsRequest {
	r.microsoftGraphBaseItem = &microsoftGraphBaseItem
	return r
}

func (r ApiSitesUpdateItemsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.SitesUpdateItemsExecute(r)
}

/*
SitesUpdateItems Update the navigation property items in sites

Used to address any item contained in this site. This collection can't be enumerated.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId key: id of site
 @param baseItemId key: id of baseItem
 @return ApiSitesUpdateItemsRequest
*/
func (a *SitesBaseItemApiService) SitesUpdateItems(ctx _context.Context, siteId string, baseItemId string) ApiSitesUpdateItemsRequest {
	return ApiSitesUpdateItemsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		baseItemId: baseItemId,
	}
}

// Execute executes the request
func (a *SitesBaseItemApiService) SitesUpdateItemsExecute(r ApiSitesUpdateItemsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesBaseItemApiService.SitesUpdateItems")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sites/{site-id}/items/{baseItem-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"site-id"+"}", _neturl.PathEscape(parameterToString(r.siteId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"baseItem-id"+"}", _neturl.PathEscape(parameterToString(r.baseItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphBaseItem == nil {
		return nil, reportError("microsoftGraphBaseItem is required and must be specified")
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
	localVarPostBody = r.microsoftGraphBaseItem
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
