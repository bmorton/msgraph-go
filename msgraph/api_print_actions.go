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

// PrintActionsApiService PrintActionsApi service
type PrintActionsApiService service

type ApiPrintPrintersCreateRequest struct {
	ctx _context.Context
	ApiService *PrintActionsApiService
	inlineObject713 *InlineObject713
}

func (r ApiPrintPrintersCreateRequest) InlineObject713(inlineObject713 InlineObject713) ApiPrintPrintersCreateRequest {
	r.inlineObject713 = &inlineObject713
	return r
}

func (r ApiPrintPrintersCreateRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PrintPrintersCreateExecute(r)
}

/*
PrintPrintersCreate Invoke action create

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPrintPrintersCreateRequest
*/
func (a *PrintActionsApiService) PrintPrintersCreate(ctx _context.Context) ApiPrintPrintersCreateRequest {
	return ApiPrintPrintersCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PrintActionsApiService) PrintPrintersCreateExecute(r ApiPrintPrintersCreateRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrintActionsApiService.PrintPrintersCreate")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/print/printers/microsoft.graph.create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject713 == nil {
		return nil, reportError("inlineObject713 is required and must be specified")
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
	localVarPostBody = r.inlineObject713
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

type ApiPrintPrintersPrinterRestoreFactoryDefaultsRequest struct {
	ctx _context.Context
	ApiService *PrintActionsApiService
	printerId string
}


func (r ApiPrintPrintersPrinterRestoreFactoryDefaultsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PrintPrintersPrinterRestoreFactoryDefaultsExecute(r)
}

/*
PrintPrintersPrinterRestoreFactoryDefaults Invoke action restoreFactoryDefaults

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param printerId key: id of printer
 @return ApiPrintPrintersPrinterRestoreFactoryDefaultsRequest
*/
func (a *PrintActionsApiService) PrintPrintersPrinterRestoreFactoryDefaults(ctx _context.Context, printerId string) ApiPrintPrintersPrinterRestoreFactoryDefaultsRequest {
	return ApiPrintPrintersPrinterRestoreFactoryDefaultsRequest{
		ApiService: a,
		ctx: ctx,
		printerId: printerId,
	}
}

// Execute executes the request
func (a *PrintActionsApiService) PrintPrintersPrinterRestoreFactoryDefaultsExecute(r ApiPrintPrintersPrinterRestoreFactoryDefaultsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrintActionsApiService.PrintPrintersPrinterRestoreFactoryDefaults")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/print/printers/{printer-id}/microsoft.graph.restoreFactoryDefaults"
	localVarPath = strings.Replace(localVarPath, "{"+"printer-id"+"}", _neturl.PathEscape(parameterToString(r.printerId, "")), -1)

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

type ApiPrintSharesPrinterSharePrinterRestoreFactoryDefaultsRequest struct {
	ctx _context.Context
	ApiService *PrintActionsApiService
	printerShareId string
}


func (r ApiPrintSharesPrinterSharePrinterRestoreFactoryDefaultsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PrintSharesPrinterSharePrinterRestoreFactoryDefaultsExecute(r)
}

/*
PrintSharesPrinterSharePrinterRestoreFactoryDefaults Invoke action restoreFactoryDefaults

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param printerShareId key: id of printerShare
 @return ApiPrintSharesPrinterSharePrinterRestoreFactoryDefaultsRequest
*/
func (a *PrintActionsApiService) PrintSharesPrinterSharePrinterRestoreFactoryDefaults(ctx _context.Context, printerShareId string) ApiPrintSharesPrinterSharePrinterRestoreFactoryDefaultsRequest {
	return ApiPrintSharesPrinterSharePrinterRestoreFactoryDefaultsRequest{
		ApiService: a,
		ctx: ctx,
		printerShareId: printerShareId,
	}
}

// Execute executes the request
func (a *PrintActionsApiService) PrintSharesPrinterSharePrinterRestoreFactoryDefaultsExecute(r ApiPrintSharesPrinterSharePrinterRestoreFactoryDefaultsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrintActionsApiService.PrintSharesPrinterSharePrinterRestoreFactoryDefaults")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/print/shares/{printerShare-id}/printer/microsoft.graph.restoreFactoryDefaults"
	localVarPath = strings.Replace(localVarPath, "{"+"printerShare-id"+"}", _neturl.PathEscape(parameterToString(r.printerShareId, "")), -1)

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
