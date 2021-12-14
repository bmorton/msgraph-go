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

// DirectoryFunctionsApiService DirectoryFunctionsApi service
type DirectoryFunctionsApiService service

type ApiDirectoryAdministrativeUnitsDeltaRequest struct {
	ctx _context.Context
	ApiService *DirectoryFunctionsApiService
}


func (r ApiDirectoryAdministrativeUnitsDeltaRequest) Execute() ([]*AnyOfmicrosoftGraphAdministrativeUnit, *_nethttp.Response, error) {
	return r.ApiService.DirectoryAdministrativeUnitsDeltaExecute(r)
}

/*
DirectoryAdministrativeUnitsDelta Invoke function delta

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDirectoryAdministrativeUnitsDeltaRequest
*/
func (a *DirectoryFunctionsApiService) DirectoryAdministrativeUnitsDelta(ctx _context.Context) ApiDirectoryAdministrativeUnitsDeltaRequest {
	return ApiDirectoryAdministrativeUnitsDeltaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []*AnyOfmicrosoftGraphAdministrativeUnit
func (a *DirectoryFunctionsApiService) DirectoryAdministrativeUnitsDeltaExecute(r ApiDirectoryAdministrativeUnitsDeltaRequest) ([]*AnyOfmicrosoftGraphAdministrativeUnit, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []*AnyOfmicrosoftGraphAdministrativeUnit
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DirectoryFunctionsApiService.DirectoryAdministrativeUnitsDelta")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/directory/administrativeUnits/microsoft.graph.delta()"

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
