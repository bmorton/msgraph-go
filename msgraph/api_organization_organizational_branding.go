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

// OrganizationOrganizationalBrandingApiService OrganizationOrganizationalBrandingApi service
type OrganizationOrganizationalBrandingApiService service

type ApiOrganizationDeleteBrandingRequest struct {
	ctx _context.Context
	ApiService *OrganizationOrganizationalBrandingApiService
	organizationId string
	ifMatch *string
}

// ETag
func (r ApiOrganizationDeleteBrandingRequest) IfMatch(ifMatch string) ApiOrganizationDeleteBrandingRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiOrganizationDeleteBrandingRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OrganizationDeleteBrandingExecute(r)
}

/*
OrganizationDeleteBranding Delete navigation property branding for organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId key: id of organization
 @return ApiOrganizationDeleteBrandingRequest
*/
func (a *OrganizationOrganizationalBrandingApiService) OrganizationDeleteBranding(ctx _context.Context, organizationId string) ApiOrganizationDeleteBrandingRequest {
	return ApiOrganizationDeleteBrandingRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
func (a *OrganizationOrganizationalBrandingApiService) OrganizationDeleteBrandingExecute(r ApiOrganizationDeleteBrandingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationOrganizationalBrandingApiService.OrganizationDeleteBranding")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organization-id}/branding"
	localVarPath = strings.Replace(localVarPath, "{"+"organization-id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type ApiOrganizationGetBrandingRequest struct {
	ctx _context.Context
	ApiService *OrganizationOrganizationalBrandingApiService
	organizationId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiOrganizationGetBrandingRequest) Select_(select_ []string) ApiOrganizationGetBrandingRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiOrganizationGetBrandingRequest) Expand(expand []string) ApiOrganizationGetBrandingRequest {
	r.expand = &expand
	return r
}

func (r ApiOrganizationGetBrandingRequest) Execute() (MicrosoftGraphOrganizationalBranding, *_nethttp.Response, error) {
	return r.ApiService.OrganizationGetBrandingExecute(r)
}

/*
OrganizationGetBranding Get branding from organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId key: id of organization
 @return ApiOrganizationGetBrandingRequest
*/
func (a *OrganizationOrganizationalBrandingApiService) OrganizationGetBranding(ctx _context.Context, organizationId string) ApiOrganizationGetBrandingRequest {
	return ApiOrganizationGetBrandingRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphOrganizationalBranding
func (a *OrganizationOrganizationalBrandingApiService) OrganizationGetBrandingExecute(r ApiOrganizationGetBrandingRequest) (MicrosoftGraphOrganizationalBranding, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphOrganizationalBranding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationOrganizationalBrandingApiService.OrganizationGetBranding")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organization-id}/branding"
	localVarPath = strings.Replace(localVarPath, "{"+"organization-id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

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

type ApiOrganizationUpdateBrandingRequest struct {
	ctx _context.Context
	ApiService *OrganizationOrganizationalBrandingApiService
	organizationId string
	microsoftGraphOrganizationalBranding *MicrosoftGraphOrganizationalBranding
}

// New navigation property values
func (r ApiOrganizationUpdateBrandingRequest) MicrosoftGraphOrganizationalBranding(microsoftGraphOrganizationalBranding MicrosoftGraphOrganizationalBranding) ApiOrganizationUpdateBrandingRequest {
	r.microsoftGraphOrganizationalBranding = &microsoftGraphOrganizationalBranding
	return r
}

func (r ApiOrganizationUpdateBrandingRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OrganizationUpdateBrandingExecute(r)
}

/*
OrganizationUpdateBranding Update the navigation property branding in organization

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId key: id of organization
 @return ApiOrganizationUpdateBrandingRequest
*/
func (a *OrganizationOrganizationalBrandingApiService) OrganizationUpdateBranding(ctx _context.Context, organizationId string) ApiOrganizationUpdateBrandingRequest {
	return ApiOrganizationUpdateBrandingRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
func (a *OrganizationOrganizationalBrandingApiService) OrganizationUpdateBrandingExecute(r ApiOrganizationUpdateBrandingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationOrganizationalBrandingApiService.OrganizationUpdateBranding")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organization-id}/branding"
	localVarPath = strings.Replace(localVarPath, "{"+"organization-id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphOrganizationalBranding == nil {
		return nil, reportError("microsoftGraphOrganizationalBranding is required and must be specified")
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
	localVarPostBody = r.microsoftGraphOrganizationalBranding
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
