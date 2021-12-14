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

// AdminFunctionsApiService AdminFunctionsApi service
type AdminFunctionsApiService service

type ApiAdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportRequest struct {
	ctx _context.Context
	ApiService *AdminFunctionsApiService
	serviceHealthId string
	serviceHealthIssueId string
}


func (r ApiAdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportExecute(r)
}

/*
AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport Invoke function incidentReport

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceHealthId key: id of serviceHealth
 @param serviceHealthIssueId key: id of serviceHealthIssue
 @return ApiAdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportRequest
*/
func (a *AdminFunctionsApiService) AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport(ctx _context.Context, serviceHealthId string, serviceHealthIssueId string) ApiAdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportRequest {
	return ApiAdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportRequest{
		ApiService: a,
		ctx: ctx,
		serviceHealthId: serviceHealthId,
		serviceHealthIssueId: serviceHealthIssueId,
	}
}

// Execute executes the request
//  @return string
func (a *AdminFunctionsApiService) AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportExecute(r ApiAdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminFunctionsApiService.AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/serviceAnnouncement/healthOverviews/{serviceHealth-id}/issues/{serviceHealthIssue-id}/microsoft.graph.incidentReport()"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceHealth-id"+"}", _neturl.PathEscape(parameterToString(r.serviceHealthId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceHealthIssue-id"+"}", _neturl.PathEscape(parameterToString(r.serviceHealthIssueId, "")), -1)

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

type ApiAdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportRequest struct {
	ctx _context.Context
	ApiService *AdminFunctionsApiService
	serviceHealthIssueId string
}


func (r ApiAdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportExecute(r)
}

/*
AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport Invoke function incidentReport

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceHealthIssueId key: id of serviceHealthIssue
 @return ApiAdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportRequest
*/
func (a *AdminFunctionsApiService) AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport(ctx _context.Context, serviceHealthIssueId string) ApiAdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportRequest {
	return ApiAdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportRequest{
		ApiService: a,
		ctx: ctx,
		serviceHealthIssueId: serviceHealthIssueId,
	}
}

// Execute executes the request
//  @return string
func (a *AdminFunctionsApiService) AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportExecute(r ApiAdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminFunctionsApiService.AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/serviceAnnouncement/issues/{serviceHealthIssue-id}/microsoft.graph.incidentReport()"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceHealthIssue-id"+"}", _neturl.PathEscape(parameterToString(r.serviceHealthIssueId, "")), -1)

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
