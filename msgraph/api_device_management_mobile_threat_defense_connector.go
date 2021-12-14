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

// DeviceManagementMobileThreatDefenseConnectorApiService DeviceManagementMobileThreatDefenseConnectorApi service
type DeviceManagementMobileThreatDefenseConnectorApiService service

type ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementMobileThreatDefenseConnectorApiService
	microsoftGraphMobileThreatDefenseConnector *MicrosoftGraphMobileThreatDefenseConnector
}

// New navigation property
func (r ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest) MicrosoftGraphMobileThreatDefenseConnector(microsoftGraphMobileThreatDefenseConnector MicrosoftGraphMobileThreatDefenseConnector) ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest {
	r.microsoftGraphMobileThreatDefenseConnector = &microsoftGraphMobileThreatDefenseConnector
	return r
}

func (r ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest) Execute() (MicrosoftGraphMobileThreatDefenseConnector, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementCreateMobileThreatDefenseConnectorsExecute(r)
}

/*
DeviceManagementCreateMobileThreatDefenseConnectors Create new navigation property to mobileThreatDefenseConnectors for deviceManagement

The list of Mobile threat Defense connectors configured by the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest
*/
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementCreateMobileThreatDefenseConnectors(ctx _context.Context) ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest {
	return ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphMobileThreatDefenseConnector
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementCreateMobileThreatDefenseConnectorsExecute(r ApiDeviceManagementCreateMobileThreatDefenseConnectorsRequest) (MicrosoftGraphMobileThreatDefenseConnector, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphMobileThreatDefenseConnector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementMobileThreatDefenseConnectorApiService.DeviceManagementCreateMobileThreatDefenseConnectors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/mobileThreatDefenseConnectors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphMobileThreatDefenseConnector == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphMobileThreatDefenseConnector is required and must be specified")
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
	localVarPostBody = r.microsoftGraphMobileThreatDefenseConnector
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

type ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementMobileThreatDefenseConnectorApiService
	mobileThreatDefenseConnectorId string
	ifMatch *string
}

// ETag
func (r ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest) IfMatch(ifMatch string) ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementDeleteMobileThreatDefenseConnectorsExecute(r)
}

/*
DeviceManagementDeleteMobileThreatDefenseConnectors Delete navigation property mobileThreatDefenseConnectors for deviceManagement

The list of Mobile threat Defense connectors configured by the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mobileThreatDefenseConnectorId key: id of mobileThreatDefenseConnector
 @return ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest
*/
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementDeleteMobileThreatDefenseConnectors(ctx _context.Context, mobileThreatDefenseConnectorId string) ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest {
	return ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest{
		ApiService: a,
		ctx: ctx,
		mobileThreatDefenseConnectorId: mobileThreatDefenseConnectorId,
	}
}

// Execute executes the request
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementDeleteMobileThreatDefenseConnectorsExecute(r ApiDeviceManagementDeleteMobileThreatDefenseConnectorsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementMobileThreatDefenseConnectorApiService.DeviceManagementDeleteMobileThreatDefenseConnectors")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"mobileThreatDefenseConnector-id"+"}", _neturl.PathEscape(parameterToString(r.mobileThreatDefenseConnectorId, "")), -1)

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

type ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementMobileThreatDefenseConnectorApiService
	mobileThreatDefenseConnectorId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest) Select_(select_ []string) ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest) Expand(expand []string) ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest) Execute() (MicrosoftGraphMobileThreatDefenseConnector, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementGetMobileThreatDefenseConnectorsExecute(r)
}

/*
DeviceManagementGetMobileThreatDefenseConnectors Get mobileThreatDefenseConnectors from deviceManagement

The list of Mobile threat Defense connectors configured by the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mobileThreatDefenseConnectorId key: id of mobileThreatDefenseConnector
 @return ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest
*/
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementGetMobileThreatDefenseConnectors(ctx _context.Context, mobileThreatDefenseConnectorId string) ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest {
	return ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest{
		ApiService: a,
		ctx: ctx,
		mobileThreatDefenseConnectorId: mobileThreatDefenseConnectorId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphMobileThreatDefenseConnector
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementGetMobileThreatDefenseConnectorsExecute(r ApiDeviceManagementGetMobileThreatDefenseConnectorsRequest) (MicrosoftGraphMobileThreatDefenseConnector, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphMobileThreatDefenseConnector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementMobileThreatDefenseConnectorApiService.DeviceManagementGetMobileThreatDefenseConnectors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"mobileThreatDefenseConnector-id"+"}", _neturl.PathEscape(parameterToString(r.mobileThreatDefenseConnectorId, "")), -1)

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

type ApiDeviceManagementListMobileThreatDefenseConnectorsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementMobileThreatDefenseConnectorApiService
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
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Top(top int32) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Skip(skip int32) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Search(search string) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Filter(filter string) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Count(count bool) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Orderby(orderby []string) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Select_(select_ []string) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Expand(expand []string) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	r.expand = &expand
	return r
}

func (r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) Execute() (CollectionOfMobileThreatDefenseConnector, *_nethttp.Response, error) {
	return r.ApiService.DeviceManagementListMobileThreatDefenseConnectorsExecute(r)
}

/*
DeviceManagementListMobileThreatDefenseConnectors Get mobileThreatDefenseConnectors from deviceManagement

The list of Mobile threat Defense connectors configured by the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeviceManagementListMobileThreatDefenseConnectorsRequest
*/
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementListMobileThreatDefenseConnectors(ctx _context.Context) ApiDeviceManagementListMobileThreatDefenseConnectorsRequest {
	return ApiDeviceManagementListMobileThreatDefenseConnectorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfMobileThreatDefenseConnector
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementListMobileThreatDefenseConnectorsExecute(r ApiDeviceManagementListMobileThreatDefenseConnectorsRequest) (CollectionOfMobileThreatDefenseConnector, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfMobileThreatDefenseConnector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementMobileThreatDefenseConnectorApiService.DeviceManagementListMobileThreatDefenseConnectors")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/mobileThreatDefenseConnectors"

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

type ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementMobileThreatDefenseConnectorApiService
	mobileThreatDefenseConnectorId string
	microsoftGraphMobileThreatDefenseConnector *MicrosoftGraphMobileThreatDefenseConnector
}

// New navigation property values
func (r ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest) MicrosoftGraphMobileThreatDefenseConnector(microsoftGraphMobileThreatDefenseConnector MicrosoftGraphMobileThreatDefenseConnector) ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest {
	r.microsoftGraphMobileThreatDefenseConnector = &microsoftGraphMobileThreatDefenseConnector
	return r
}

func (r ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeviceManagementUpdateMobileThreatDefenseConnectorsExecute(r)
}

/*
DeviceManagementUpdateMobileThreatDefenseConnectors Update the navigation property mobileThreatDefenseConnectors in deviceManagement

The list of Mobile threat Defense connectors configured by the tenant.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mobileThreatDefenseConnectorId key: id of mobileThreatDefenseConnector
 @return ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest
*/
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementUpdateMobileThreatDefenseConnectors(ctx _context.Context, mobileThreatDefenseConnectorId string) ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest {
	return ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest{
		ApiService: a,
		ctx: ctx,
		mobileThreatDefenseConnectorId: mobileThreatDefenseConnectorId,
	}
}

// Execute executes the request
func (a *DeviceManagementMobileThreatDefenseConnectorApiService) DeviceManagementUpdateMobileThreatDefenseConnectorsExecute(r ApiDeviceManagementUpdateMobileThreatDefenseConnectorsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementMobileThreatDefenseConnectorApiService.DeviceManagementUpdateMobileThreatDefenseConnectors")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deviceManagement/mobileThreatDefenseConnectors/{mobileThreatDefenseConnector-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"mobileThreatDefenseConnector-id"+"}", _neturl.PathEscape(parameterToString(r.mobileThreatDefenseConnectorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphMobileThreatDefenseConnector == nil {
		return nil, reportError("microsoftGraphMobileThreatDefenseConnector is required and must be specified")
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
	localVarPostBody = r.microsoftGraphMobileThreatDefenseConnector
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
