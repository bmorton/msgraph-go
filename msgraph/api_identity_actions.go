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

// IdentityActionsApiService IdentityActionsApi service
type IdentityActionsApiService service

type ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest struct {
	ctx _context.Context
	ApiService *IdentityActionsApiService
	identityApiConnectorId string
	inlineObject356 *InlineObject356
}

func (r ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest) InlineObject356(inlineObject356 InlineObject356) ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest {
	r.inlineObject356 = &inlineObject356
	return r
}

func (r ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest) Execute() (AnyOfmicrosoftGraphIdentityApiConnector, *_nethttp.Response, error) {
	return r.ApiService.IdentityApiConnectorsIdentityApiConnectorUploadClientCertificateExecute(r)
}

/*
IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate Invoke action uploadClientCertificate

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param identityApiConnectorId key: id of identityApiConnector
 @return ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest
*/
func (a *IdentityActionsApiService) IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate(ctx _context.Context, identityApiConnectorId string) ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest {
	return ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest{
		ApiService: a,
		ctx: ctx,
		identityApiConnectorId: identityApiConnectorId,
	}
}

// Execute executes the request
//  @return AnyOfmicrosoftGraphIdentityApiConnector
func (a *IdentityActionsApiService) IdentityApiConnectorsIdentityApiConnectorUploadClientCertificateExecute(r ApiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest) (AnyOfmicrosoftGraphIdentityApiConnector, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AnyOfmicrosoftGraphIdentityApiConnector
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityActionsApiService.IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/apiConnectors/{identityApiConnector-id}/microsoft.graph.uploadClientCertificate"
	localVarPath = strings.Replace(localVarPath, "{"+"identityApiConnector-id"+"}", _neturl.PathEscape(parameterToString(r.identityApiConnectorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject356 == nil {
		return localVarReturnValue, nil, reportError("inlineObject356 is required and must be specified")
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
	localVarPostBody = r.inlineObject356
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

type ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest struct {
	ctx _context.Context
	ApiService *IdentityActionsApiService
	b2xIdentityUserFlowId string
	inlineObject357 *InlineObject357
}

func (r ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest) InlineObject357(inlineObject357 InlineObject357) ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest {
	r.inlineObject357 = &inlineObject357
	return r
}

func (r ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderExecute(r)
}

/*
IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder Invoke action setOrder

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param b2xIdentityUserFlowId key: id of b2xIdentityUserFlow
 @return ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest
*/
func (a *IdentityActionsApiService) IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder(ctx _context.Context, b2xIdentityUserFlowId string) ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest {
	return ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest{
		ApiService: a,
		ctx: ctx,
		b2xIdentityUserFlowId: b2xIdentityUserFlowId,
	}
}

// Execute executes the request
func (a *IdentityActionsApiService) IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderExecute(r ApiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityActionsApiService.IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identity/b2xUserFlows/{b2xIdentityUserFlow-id}/userAttributeAssignments/microsoft.graph.setOrder"
	localVarPath = strings.Replace(localVarPath, "{"+"b2xIdentityUserFlow-id"+"}", _neturl.PathEscape(parameterToString(r.b2xIdentityUserFlowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject357 == nil {
		return nil, reportError("inlineObject357 is required and must be specified")
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
	localVarPostBody = r.inlineObject357
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
