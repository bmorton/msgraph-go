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
)

// Linger please
var (
	_ _context.Context
)

// EducationEducationRootApiService EducationEducationRootApi service
type EducationEducationRootApiService service

type ApiEducationEducationRootGetEducationRootRequest struct {
	ctx _context.Context
	ApiService *EducationEducationRootApiService
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiEducationEducationRootGetEducationRootRequest) Select_(select_ []string) ApiEducationEducationRootGetEducationRootRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiEducationEducationRootGetEducationRootRequest) Expand(expand []string) ApiEducationEducationRootGetEducationRootRequest {
	r.expand = &expand
	return r
}

func (r ApiEducationEducationRootGetEducationRootRequest) Execute() (MicrosoftGraphEducationRoot, *_nethttp.Response, error) {
	return r.ApiService.EducationEducationRootGetEducationRootExecute(r)
}

/*
EducationEducationRootGetEducationRoot Get education

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEducationEducationRootGetEducationRootRequest
*/
func (a *EducationEducationRootApiService) EducationEducationRootGetEducationRoot(ctx _context.Context) ApiEducationEducationRootGetEducationRootRequest {
	return ApiEducationEducationRootGetEducationRootRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphEducationRoot
func (a *EducationEducationRootApiService) EducationEducationRootGetEducationRootExecute(r ApiEducationEducationRootGetEducationRootRequest) (MicrosoftGraphEducationRoot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphEducationRoot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EducationEducationRootApiService.EducationEducationRootGetEducationRoot")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/education"

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

type ApiEducationEducationRootUpdateEducationRootRequest struct {
	ctx _context.Context
	ApiService *EducationEducationRootApiService
	microsoftGraphEducationRoot *MicrosoftGraphEducationRoot
}

// New property values
func (r ApiEducationEducationRootUpdateEducationRootRequest) MicrosoftGraphEducationRoot(microsoftGraphEducationRoot MicrosoftGraphEducationRoot) ApiEducationEducationRootUpdateEducationRootRequest {
	r.microsoftGraphEducationRoot = &microsoftGraphEducationRoot
	return r
}

func (r ApiEducationEducationRootUpdateEducationRootRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.EducationEducationRootUpdateEducationRootExecute(r)
}

/*
EducationEducationRootUpdateEducationRoot Update education

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEducationEducationRootUpdateEducationRootRequest
*/
func (a *EducationEducationRootApiService) EducationEducationRootUpdateEducationRoot(ctx _context.Context) ApiEducationEducationRootUpdateEducationRootRequest {
	return ApiEducationEducationRootUpdateEducationRootRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *EducationEducationRootApiService) EducationEducationRootUpdateEducationRootExecute(r ApiEducationEducationRootUpdateEducationRootRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EducationEducationRootApiService.EducationEducationRootUpdateEducationRoot")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/education"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphEducationRoot == nil {
		return nil, reportError("microsoftGraphEducationRoot is required and must be specified")
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
	localVarPostBody = r.microsoftGraphEducationRoot
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