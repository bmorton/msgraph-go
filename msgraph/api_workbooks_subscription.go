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

// WorkbooksSubscriptionApiService WorkbooksSubscriptionApi service
type WorkbooksSubscriptionApiService service

type ApiWorkbooksCreateSubscriptionsRequest struct {
	ctx _context.Context
	ApiService *WorkbooksSubscriptionApiService
	driveItemId string
	microsoftGraphSubscription *MicrosoftGraphSubscription
}

// New navigation property
func (r ApiWorkbooksCreateSubscriptionsRequest) MicrosoftGraphSubscription(microsoftGraphSubscription MicrosoftGraphSubscription) ApiWorkbooksCreateSubscriptionsRequest {
	r.microsoftGraphSubscription = &microsoftGraphSubscription
	return r
}

func (r ApiWorkbooksCreateSubscriptionsRequest) Execute() (MicrosoftGraphSubscription, *_nethttp.Response, error) {
	return r.ApiService.WorkbooksCreateSubscriptionsExecute(r)
}

/*
WorkbooksCreateSubscriptions Create new navigation property to subscriptions for workbooks

The set of subscriptions on the item. Only supported on the root of a drive.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveItemId key: id of driveItem
 @return ApiWorkbooksCreateSubscriptionsRequest
*/
func (a *WorkbooksSubscriptionApiService) WorkbooksCreateSubscriptions(ctx _context.Context, driveItemId string) ApiWorkbooksCreateSubscriptionsRequest {
	return ApiWorkbooksCreateSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		driveItemId: driveItemId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSubscription
func (a *WorkbooksSubscriptionApiService) WorkbooksCreateSubscriptionsExecute(r ApiWorkbooksCreateSubscriptionsRequest) (MicrosoftGraphSubscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkbooksSubscriptionApiService.WorkbooksCreateSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workbooks/{driveItem-id}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"driveItem-id"+"}", _neturl.PathEscape(parameterToString(r.driveItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSubscription == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphSubscription is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSubscription
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

type ApiWorkbooksDeleteSubscriptionsRequest struct {
	ctx _context.Context
	ApiService *WorkbooksSubscriptionApiService
	driveItemId string
	subscriptionId string
	ifMatch *string
}

// ETag
func (r ApiWorkbooksDeleteSubscriptionsRequest) IfMatch(ifMatch string) ApiWorkbooksDeleteSubscriptionsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiWorkbooksDeleteSubscriptionsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WorkbooksDeleteSubscriptionsExecute(r)
}

/*
WorkbooksDeleteSubscriptions Delete navigation property subscriptions for workbooks

The set of subscriptions on the item. Only supported on the root of a drive.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveItemId key: id of driveItem
 @param subscriptionId key: id of subscription
 @return ApiWorkbooksDeleteSubscriptionsRequest
*/
func (a *WorkbooksSubscriptionApiService) WorkbooksDeleteSubscriptions(ctx _context.Context, driveItemId string, subscriptionId string) ApiWorkbooksDeleteSubscriptionsRequest {
	return ApiWorkbooksDeleteSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		driveItemId: driveItemId,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *WorkbooksSubscriptionApiService) WorkbooksDeleteSubscriptionsExecute(r ApiWorkbooksDeleteSubscriptionsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkbooksSubscriptionApiService.WorkbooksDeleteSubscriptions")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workbooks/{driveItem-id}/subscriptions/{subscription-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"driveItem-id"+"}", _neturl.PathEscape(parameterToString(r.driveItemId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscription-id"+"}", _neturl.PathEscape(parameterToString(r.subscriptionId, "")), -1)

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

type ApiWorkbooksGetSubscriptionsRequest struct {
	ctx _context.Context
	ApiService *WorkbooksSubscriptionApiService
	driveItemId string
	subscriptionId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiWorkbooksGetSubscriptionsRequest) Select_(select_ []string) ApiWorkbooksGetSubscriptionsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiWorkbooksGetSubscriptionsRequest) Expand(expand []string) ApiWorkbooksGetSubscriptionsRequest {
	r.expand = &expand
	return r
}

func (r ApiWorkbooksGetSubscriptionsRequest) Execute() (MicrosoftGraphSubscription, *_nethttp.Response, error) {
	return r.ApiService.WorkbooksGetSubscriptionsExecute(r)
}

/*
WorkbooksGetSubscriptions Get subscriptions from workbooks

The set of subscriptions on the item. Only supported on the root of a drive.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveItemId key: id of driveItem
 @param subscriptionId key: id of subscription
 @return ApiWorkbooksGetSubscriptionsRequest
*/
func (a *WorkbooksSubscriptionApiService) WorkbooksGetSubscriptions(ctx _context.Context, driveItemId string, subscriptionId string) ApiWorkbooksGetSubscriptionsRequest {
	return ApiWorkbooksGetSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		driveItemId: driveItemId,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphSubscription
func (a *WorkbooksSubscriptionApiService) WorkbooksGetSubscriptionsExecute(r ApiWorkbooksGetSubscriptionsRequest) (MicrosoftGraphSubscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkbooksSubscriptionApiService.WorkbooksGetSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workbooks/{driveItem-id}/subscriptions/{subscription-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"driveItem-id"+"}", _neturl.PathEscape(parameterToString(r.driveItemId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscription-id"+"}", _neturl.PathEscape(parameterToString(r.subscriptionId, "")), -1)

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

type ApiWorkbooksListSubscriptionsRequest struct {
	ctx _context.Context
	ApiService *WorkbooksSubscriptionApiService
	driveItemId string
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
func (r ApiWorkbooksListSubscriptionsRequest) Top(top int32) ApiWorkbooksListSubscriptionsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiWorkbooksListSubscriptionsRequest) Skip(skip int32) ApiWorkbooksListSubscriptionsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiWorkbooksListSubscriptionsRequest) Search(search string) ApiWorkbooksListSubscriptionsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiWorkbooksListSubscriptionsRequest) Filter(filter string) ApiWorkbooksListSubscriptionsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiWorkbooksListSubscriptionsRequest) Count(count bool) ApiWorkbooksListSubscriptionsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiWorkbooksListSubscriptionsRequest) Orderby(orderby []string) ApiWorkbooksListSubscriptionsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiWorkbooksListSubscriptionsRequest) Select_(select_ []string) ApiWorkbooksListSubscriptionsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiWorkbooksListSubscriptionsRequest) Expand(expand []string) ApiWorkbooksListSubscriptionsRequest {
	r.expand = &expand
	return r
}

func (r ApiWorkbooksListSubscriptionsRequest) Execute() (CollectionOfSubscription, *_nethttp.Response, error) {
	return r.ApiService.WorkbooksListSubscriptionsExecute(r)
}

/*
WorkbooksListSubscriptions Get subscriptions from workbooks

The set of subscriptions on the item. Only supported on the root of a drive.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveItemId key: id of driveItem
 @return ApiWorkbooksListSubscriptionsRequest
*/
func (a *WorkbooksSubscriptionApiService) WorkbooksListSubscriptions(ctx _context.Context, driveItemId string) ApiWorkbooksListSubscriptionsRequest {
	return ApiWorkbooksListSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		driveItemId: driveItemId,
	}
}

// Execute executes the request
//  @return CollectionOfSubscription
func (a *WorkbooksSubscriptionApiService) WorkbooksListSubscriptionsExecute(r ApiWorkbooksListSubscriptionsRequest) (CollectionOfSubscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkbooksSubscriptionApiService.WorkbooksListSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workbooks/{driveItem-id}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"driveItem-id"+"}", _neturl.PathEscape(parameterToString(r.driveItemId, "")), -1)

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

type ApiWorkbooksUpdateSubscriptionsRequest struct {
	ctx _context.Context
	ApiService *WorkbooksSubscriptionApiService
	driveItemId string
	subscriptionId string
	microsoftGraphSubscription *MicrosoftGraphSubscription
}

// New navigation property values
func (r ApiWorkbooksUpdateSubscriptionsRequest) MicrosoftGraphSubscription(microsoftGraphSubscription MicrosoftGraphSubscription) ApiWorkbooksUpdateSubscriptionsRequest {
	r.microsoftGraphSubscription = &microsoftGraphSubscription
	return r
}

func (r ApiWorkbooksUpdateSubscriptionsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WorkbooksUpdateSubscriptionsExecute(r)
}

/*
WorkbooksUpdateSubscriptions Update the navigation property subscriptions in workbooks

The set of subscriptions on the item. Only supported on the root of a drive.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveItemId key: id of driveItem
 @param subscriptionId key: id of subscription
 @return ApiWorkbooksUpdateSubscriptionsRequest
*/
func (a *WorkbooksSubscriptionApiService) WorkbooksUpdateSubscriptions(ctx _context.Context, driveItemId string, subscriptionId string) ApiWorkbooksUpdateSubscriptionsRequest {
	return ApiWorkbooksUpdateSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		driveItemId: driveItemId,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *WorkbooksSubscriptionApiService) WorkbooksUpdateSubscriptionsExecute(r ApiWorkbooksUpdateSubscriptionsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkbooksSubscriptionApiService.WorkbooksUpdateSubscriptions")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workbooks/{driveItem-id}/subscriptions/{subscription-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"driveItem-id"+"}", _neturl.PathEscape(parameterToString(r.driveItemId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscription-id"+"}", _neturl.PathEscape(parameterToString(r.subscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphSubscription == nil {
		return nil, reportError("microsoftGraphSubscription is required and must be specified")
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
	localVarPostBody = r.microsoftGraphSubscription
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
