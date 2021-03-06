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

// DeviceManagementRemoteAssistancePartnerApiService DeviceManagementRemoteAssistancePartnerApi service
type DeviceManagementRemoteAssistancePartnerApiService service

type ApiDeviceManagementCreateRemoteAssistancePartnersRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementRemoteAssistancePartnerApiService
	microsoftGraphRemoteAssistancePartner *MicrosoftGraphRemoteAssistancePartner
}

// New navigation property
func (r ApiDeviceManagementCreateRemoteAssistancePartnersRequest) MicrosoftGraphRemoteAssistancePartner(microsoftGraphRemoteAssistancePartner MicrosoftGraphRemoteAssistancePartner) ApiDeviceManagementCreateRemoteAssistancePartnersRequest {
	r.microsoftGraphRemoteAssistancePartner = &microsoftGraphRemoteAssistancePartner
	return r
}

func (r ApiDeviceManagementCreateRemoteAssistancePartnersRequest) Execute() (MicrosoftGraphRemoteAssistancePartner, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementCreateRemoteAssistancePartnersExecute(r)
}

/*
DeviceManagementCreateRemoteAssistancePartners Create new navigation property to remoteAssistancePartners for deviceManagement

The remote assist partners.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementCreateRemoteAssistancePartnersRequest
*/
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementCreateRemoteAssistancePartners(ctx _context.Context) ApiDeviceManagementCreateRemoteAssistancePartnersRequest {
	return ApiDeviceManagementCreateRemoteAssistancePartnersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphRemoteAssistancePartner
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementCreateRemoteAssistancePartnersExecute(r ApiDeviceManagementCreateRemoteAssistancePartnersRequest) (MicrosoftGraphRemoteAssistancePartner, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphRemoteAssistancePartner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementRemoteAssistancePartnerApiService.DeviceManagementCreateRemoteAssistancePartners")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/remoteAssistancePartners"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphRemoteAssistancePartner == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphRemoteAssistancePartner is required and must be specified")
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
	localVarPostBody = r.microsoftGraphRemoteAssistancePartner
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

type ApiDeviceManagementDeleteRemoteAssistancePartnersRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementRemoteAssistancePartnerApiService
	remoteAssistancePartnerId string
	ifMatch *string
}

// ETag
func (r ApiDeviceManagementDeleteRemoteAssistancePartnersRequest) IfMatch(ifMatch string) ApiDeviceManagementDeleteRemoteAssistancePartnersRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeviceManagementDeleteRemoteAssistancePartnersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementDeleteRemoteAssistancePartnersExecute(r)
}

/*
DeviceManagementDeleteRemoteAssistancePartners Delete navigation property remoteAssistancePartners for deviceManagement

The remote assist partners.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param remoteAssistancePartnerId key: id of remoteAssistancePartner
 @return ApiDeviceManagementDeleteRemoteAssistancePartnersRequest
*/
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementDeleteRemoteAssistancePartners(ctx _context.Context, remoteAssistancePartnerId string) ApiDeviceManagementDeleteRemoteAssistancePartnersRequest {
	return ApiDeviceManagementDeleteRemoteAssistancePartnersRequest{
		ApiService: a,
		ctx: ctx,
		remoteAssistancePartnerId: remoteAssistancePartnerId,
	}
}

// Execute executes the request
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementDeleteRemoteAssistancePartnersExecute(r ApiDeviceManagementDeleteRemoteAssistancePartnersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementRemoteAssistancePartnerApiService.DeviceManagementDeleteRemoteAssistancePartners")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"remoteAssistancePartner-id"+"}", _neturl.PathEscape(parameterToString(r.remoteAssistancePartnerId, "")), -1)

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

type ApiDeviceManagementGetRemoteAssistancePartnersRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementRemoteAssistancePartnerApiService
	remoteAssistancePartnerId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiDeviceManagementGetRemoteAssistancePartnersRequest) Select_(select_ []string) ApiDeviceManagementGetRemoteAssistancePartnersRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementGetRemoteAssistancePartnersRequest) Expand(expand []string) ApiDeviceManagementGetRemoteAssistancePartnersRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementGetRemoteAssistancePartnersRequest) Execute() (MicrosoftGraphRemoteAssistancePartner, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementGetRemoteAssistancePartnersExecute(r)
}

/*
DeviceManagementGetRemoteAssistancePartners Get remoteAssistancePartners from deviceManagement

The remote assist partners.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param remoteAssistancePartnerId key: id of remoteAssistancePartner
 @return ApiDeviceManagementGetRemoteAssistancePartnersRequest
*/
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementGetRemoteAssistancePartners(ctx _context.Context, remoteAssistancePartnerId string) ApiDeviceManagementGetRemoteAssistancePartnersRequest {
	return ApiDeviceManagementGetRemoteAssistancePartnersRequest{
		ApiService: a,
		ctx: ctx,
		remoteAssistancePartnerId: remoteAssistancePartnerId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphRemoteAssistancePartner
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementGetRemoteAssistancePartnersExecute(r ApiDeviceManagementGetRemoteAssistancePartnersRequest) (MicrosoftGraphRemoteAssistancePartner, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphRemoteAssistancePartner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementRemoteAssistancePartnerApiService.DeviceManagementGetRemoteAssistancePartners")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"remoteAssistancePartner-id"+"}", _neturl.PathEscape(parameterToString(r.remoteAssistancePartnerId, "")), -1)

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

type ApiDeviceManagementListRemoteAssistancePartnersRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementRemoteAssistancePartnerApiService
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
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Top(top int32) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Skip(skip int32) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Search(search string) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Filter(filter string) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Count(count bool) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Orderby(orderby []string) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Select_(select_ []string) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Expand(expand []string) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementListRemoteAssistancePartnersRequest) Execute() (CollectionOfRemoteAssistancePartner, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementListRemoteAssistancePartnersExecute(r)
}

/*
DeviceManagementListRemoteAssistancePartners Get remoteAssistancePartners from deviceManagement

The remote assist partners.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementListRemoteAssistancePartnersRequest
*/
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementListRemoteAssistancePartners(ctx _context.Context) ApiDeviceManagementListRemoteAssistancePartnersRequest {
	return ApiDeviceManagementListRemoteAssistancePartnersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfRemoteAssistancePartner
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementListRemoteAssistancePartnersExecute(r ApiDeviceManagementListRemoteAssistancePartnersRequest) (CollectionOfRemoteAssistancePartner, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfRemoteAssistancePartner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementRemoteAssistancePartnerApiService.DeviceManagementListRemoteAssistancePartners")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/remoteAssistancePartners"

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

type ApiDeviceManagementUpdateRemoteAssistancePartnersRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementRemoteAssistancePartnerApiService
	remoteAssistancePartnerId string
	microsoftGraphRemoteAssistancePartner *MicrosoftGraphRemoteAssistancePartner
}

// New navigation property values
func (r ApiDeviceManagementUpdateRemoteAssistancePartnersRequest) MicrosoftGraphRemoteAssistancePartner(microsoftGraphRemoteAssistancePartner MicrosoftGraphRemoteAssistancePartner) ApiDeviceManagementUpdateRemoteAssistancePartnersRequest {
	r.microsoftGraphRemoteAssistancePartner = &microsoftGraphRemoteAssistancePartner
	return r
}

func (r ApiDeviceManagementUpdateRemoteAssistancePartnersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementUpdateRemoteAssistancePartnersExecute(r)
}

/*
DeviceManagementUpdateRemoteAssistancePartners Update the navigation property remoteAssistancePartners in deviceManagement

The remote assist partners.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param remoteAssistancePartnerId key: id of remoteAssistancePartner
 @return ApiDeviceManagementUpdateRemoteAssistancePartnersRequest
*/
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementUpdateRemoteAssistancePartners(ctx _context.Context, remoteAssistancePartnerId string) ApiDeviceManagementUpdateRemoteAssistancePartnersRequest {
	return ApiDeviceManagementUpdateRemoteAssistancePartnersRequest{
		ApiService: a,
		ctx: ctx,
		remoteAssistancePartnerId: remoteAssistancePartnerId,
	}
}

// Execute executes the request
func (a *DeviceManagementRemoteAssistancePartnerApiService) DeviceManagementUpdateRemoteAssistancePartnersExecute(r ApiDeviceManagementUpdateRemoteAssistancePartnersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementRemoteAssistancePartnerApiService.DeviceManagementUpdateRemoteAssistancePartners")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"remoteAssistancePartner-id"+"}", _neturl.PathEscape(parameterToString(r.remoteAssistancePartnerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphRemoteAssistancePartner == nil {
		return nil, reportError("microsoftGraphRemoteAssistancePartner is required and must be specified")
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
	localVarPostBody = r.microsoftGraphRemoteAssistancePartner
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
