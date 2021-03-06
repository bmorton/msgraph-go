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

// AgreementAcceptancesAgreementAcceptanceApiService AgreementAcceptancesAgreementAcceptanceApi service
type AgreementAcceptancesAgreementAcceptanceApiService service

type ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest struct {
	ctx _context.Context
	ApiService *AgreementAcceptancesAgreementAcceptanceApiService
	microsoftGraphAgreementAcceptance *MicrosoftGraphAgreementAcceptance
}

// New entity
func (r ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest) MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance MicrosoftGraphAgreementAcceptance) ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest {
	r.microsoftGraphAgreementAcceptance = &microsoftGraphAgreementAcceptance
	return r
}

func (r ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest) Execute() (MicrosoftGraphAgreementAcceptance, *_nethttp.Response, error) {
	return r.ApiService.AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceExecute(r)
}

/*
AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance Add new entity to agreementAcceptances

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest
*/
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance(ctx _context.Context) ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest {
	return ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphAgreementAcceptance
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceExecute(r ApiAgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptanceRequest) (MicrosoftGraphAgreementAcceptance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphAgreementAcceptance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAcceptancesAgreementAcceptanceApiService.AgreementAcceptancesAgreementAcceptanceCreateAgreementAcceptance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agreementAcceptances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphAgreementAcceptance == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphAgreementAcceptance is required and must be specified")
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
	localVarPostBody = r.microsoftGraphAgreementAcceptance
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

type ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest struct {
	ctx _context.Context
	ApiService *AgreementAcceptancesAgreementAcceptanceApiService
	agreementAcceptanceId string
	ifMatch *string
}

// ETag
func (r ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest) IfMatch(ifMatch string) ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceExecute(r)
}

/*
AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance Delete entity from agreementAcceptances

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agreementAcceptanceId key: id of agreementAcceptance
 @return ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest
*/
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance(ctx _context.Context, agreementAcceptanceId string) ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest {
	return ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest{
		ApiService: a,
		ctx: ctx,
		agreementAcceptanceId: agreementAcceptanceId,
	}
}

// Execute executes the request
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceExecute(r ApiAgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptanceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAcceptancesAgreementAcceptanceApiService.AgreementAcceptancesAgreementAcceptanceDeleteAgreementAcceptance")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agreementAcceptances/{agreementAcceptance-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"agreementAcceptance-id"+"}", _neturl.PathEscape(parameterToString(r.agreementAcceptanceId, "")), -1)

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

type ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest struct {
	ctx _context.Context
	ApiService *AgreementAcceptancesAgreementAcceptanceApiService
	agreementAcceptanceId string
	select_ *[]string
}

// Select properties to be returned
func (r ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest) Select_(select_ []string) ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest {
	r.select_ = &select_
	return r
}

func (r ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest) Execute() (MicrosoftGraphAgreementAcceptance, *_nethttp.Response, error) {
	return r.ApiService.AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceExecute(r)
}

/*
AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance Get entity from agreementAcceptances by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agreementAcceptanceId key: id of agreementAcceptance
 @return ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest
*/
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance(ctx _context.Context, agreementAcceptanceId string) ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest {
	return ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest{
		ApiService: a,
		ctx: ctx,
		agreementAcceptanceId: agreementAcceptanceId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphAgreementAcceptance
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceExecute(r ApiAgreementAcceptancesAgreementAcceptanceGetAgreementAcceptanceRequest) (MicrosoftGraphAgreementAcceptance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphAgreementAcceptance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAcceptancesAgreementAcceptanceApiService.AgreementAcceptancesAgreementAcceptanceGetAgreementAcceptance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agreementAcceptances/{agreementAcceptance-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"agreementAcceptance-id"+"}", _neturl.PathEscape(parameterToString(r.agreementAcceptanceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("$select", parameterToString(*r.select_, "csv"))
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

type ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest struct {
	ctx _context.Context
	ApiService *AgreementAcceptancesAgreementAcceptanceApiService
	search *string
	select_ *[]string
}

// Search items by search phrases
func (r ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest) Search(search string) ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest {
	r.search = &search
	return r
}
// Select properties to be returned
func (r ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest) Select_(select_ []string) ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest {
	r.select_ = &select_
	return r
}

func (r ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest) Execute() (CollectionOfAgreementAcceptance, *_nethttp.Response, error) {
	return r.ApiService.AgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceExecute(r)
}

/*
AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance Get entities from agreementAcceptances

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest
*/
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance(ctx _context.Context) ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest {
	return ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfAgreementAcceptance
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceExecute(r ApiAgreementAcceptancesAgreementAcceptanceListAgreementAcceptanceRequest) (CollectionOfAgreementAcceptance, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfAgreementAcceptance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAcceptancesAgreementAcceptanceApiService.AgreementAcceptancesAgreementAcceptanceListAgreementAcceptance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agreementAcceptances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.search != nil {
		localVarQueryParams.Add("$search", parameterToString(*r.search, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("$select", parameterToString(*r.select_, "csv"))
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

type ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest struct {
	ctx _context.Context
	ApiService *AgreementAcceptancesAgreementAcceptanceApiService
	agreementAcceptanceId string
	microsoftGraphAgreementAcceptance *MicrosoftGraphAgreementAcceptance
}

// New property values
func (r ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest) MicrosoftGraphAgreementAcceptance(microsoftGraphAgreementAcceptance MicrosoftGraphAgreementAcceptance) ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest {
	r.microsoftGraphAgreementAcceptance = &microsoftGraphAgreementAcceptance
	return r
}

func (r ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceExecute(r)
}

/*
AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance Update entity in agreementAcceptances

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agreementAcceptanceId key: id of agreementAcceptance
 @return ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest
*/
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance(ctx _context.Context, agreementAcceptanceId string) ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest {
	return ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest{
		ApiService: a,
		ctx: ctx,
		agreementAcceptanceId: agreementAcceptanceId,
	}
}

// Execute executes the request
func (a *AgreementAcceptancesAgreementAcceptanceApiService) AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceExecute(r ApiAgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptanceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAcceptancesAgreementAcceptanceApiService.AgreementAcceptancesAgreementAcceptanceUpdateAgreementAcceptance")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agreementAcceptances/{agreementAcceptance-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"agreementAcceptance-id"+"}", _neturl.PathEscape(parameterToString(r.agreementAcceptanceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphAgreementAcceptance == nil {
		return nil, reportError("microsoftGraphAgreementAcceptance is required and must be specified")
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
	localVarPostBody = r.microsoftGraphAgreementAcceptance
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
