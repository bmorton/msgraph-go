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
)

// Linger please
var (
	_ _context.Context
)

// MeUserSettingsApiService MeUserSettingsApi service
type MeUserSettingsApiService service

type ApiMeDeleteSettingsRequest struct {
	ctx _context.Context
	ApiService *MeUserSettingsApiService
	ifMatch *string
}

// ETag
func (r ApiMeDeleteSettingsRequest) IfMatch(ifMatch string) ApiMeDeleteSettingsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiMeDeleteSettingsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeDeleteSettingsExecute(r)
}

/*
MeDeleteSettings Delete navigation property settings for me

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeDeleteSettingsRequest
*/
func (a *MeUserSettingsApiService) MeDeleteSettings(ctx _context.Context) ApiMeDeleteSettingsRequest {
	return ApiMeDeleteSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MeUserSettingsApiService) MeDeleteSettingsExecute(r ApiMeDeleteSettingsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeUserSettingsApiService.MeDeleteSettings")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/settings"

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

type ApiMeGetSettingsRequest struct {
	ctx _context.Context
	ApiService *MeUserSettingsApiService
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiMeGetSettingsRequest) Select_(select_ []string) ApiMeGetSettingsRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiMeGetSettingsRequest) Expand(expand []string) ApiMeGetSettingsRequest {
	r.expand = &expand
	return r
}

func (r ApiMeGetSettingsRequest) Execute() (MicrosoftGraphUserSettings, *_nethttp.Response, error) {
	return r.ApiService.MeGetSettingsExecute(r)
}

/*
MeGetSettings Get settings from me

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeGetSettingsRequest
*/
func (a *MeUserSettingsApiService) MeGetSettings(ctx _context.Context) ApiMeGetSettingsRequest {
	return ApiMeGetSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphUserSettings
func (a *MeUserSettingsApiService) MeGetSettingsExecute(r ApiMeGetSettingsRequest) (MicrosoftGraphUserSettings, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphUserSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeUserSettingsApiService.MeGetSettings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/settings"

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

type ApiMeSettingsDeleteShiftPreferencesRequest struct {
	ctx _context.Context
	ApiService *MeUserSettingsApiService
	ifMatch *string
}

// ETag
func (r ApiMeSettingsDeleteShiftPreferencesRequest) IfMatch(ifMatch string) ApiMeSettingsDeleteShiftPreferencesRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiMeSettingsDeleteShiftPreferencesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeSettingsDeleteShiftPreferencesExecute(r)
}

/*
MeSettingsDeleteShiftPreferences Delete navigation property shiftPreferences for me

The shift preferences for the user.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeSettingsDeleteShiftPreferencesRequest
*/
func (a *MeUserSettingsApiService) MeSettingsDeleteShiftPreferences(ctx _context.Context) ApiMeSettingsDeleteShiftPreferencesRequest {
	return ApiMeSettingsDeleteShiftPreferencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MeUserSettingsApiService) MeSettingsDeleteShiftPreferencesExecute(r ApiMeSettingsDeleteShiftPreferencesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeUserSettingsApiService.MeSettingsDeleteShiftPreferences")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/settings/shiftPreferences"

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

type ApiMeSettingsGetShiftPreferencesRequest struct {
	ctx _context.Context
	ApiService *MeUserSettingsApiService
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiMeSettingsGetShiftPreferencesRequest) Select_(select_ []string) ApiMeSettingsGetShiftPreferencesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiMeSettingsGetShiftPreferencesRequest) Expand(expand []string) ApiMeSettingsGetShiftPreferencesRequest {
	r.expand = &expand
	return r
}

func (r ApiMeSettingsGetShiftPreferencesRequest) Execute() (MicrosoftGraphShiftPreferences, *_nethttp.Response, error) {
	return r.ApiService.MeSettingsGetShiftPreferencesExecute(r)
}

/*
MeSettingsGetShiftPreferences Get shiftPreferences from me

The shift preferences for the user.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeSettingsGetShiftPreferencesRequest
*/
func (a *MeUserSettingsApiService) MeSettingsGetShiftPreferences(ctx _context.Context) ApiMeSettingsGetShiftPreferencesRequest {
	return ApiMeSettingsGetShiftPreferencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphShiftPreferences
func (a *MeUserSettingsApiService) MeSettingsGetShiftPreferencesExecute(r ApiMeSettingsGetShiftPreferencesRequest) (MicrosoftGraphShiftPreferences, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphShiftPreferences
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeUserSettingsApiService.MeSettingsGetShiftPreferences")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/settings/shiftPreferences"

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

type ApiMeSettingsUpdateShiftPreferencesRequest struct {
	ctx _context.Context
	ApiService *MeUserSettingsApiService
	microsoftGraphShiftPreferences *MicrosoftGraphShiftPreferences
}

// New navigation property values
func (r ApiMeSettingsUpdateShiftPreferencesRequest) MicrosoftGraphShiftPreferences(microsoftGraphShiftPreferences MicrosoftGraphShiftPreferences) ApiMeSettingsUpdateShiftPreferencesRequest {
	r.microsoftGraphShiftPreferences = &microsoftGraphShiftPreferences
	return r
}

func (r ApiMeSettingsUpdateShiftPreferencesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeSettingsUpdateShiftPreferencesExecute(r)
}

/*
MeSettingsUpdateShiftPreferences Update the navigation property shiftPreferences in me

The shift preferences for the user.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeSettingsUpdateShiftPreferencesRequest
*/
func (a *MeUserSettingsApiService) MeSettingsUpdateShiftPreferences(ctx _context.Context) ApiMeSettingsUpdateShiftPreferencesRequest {
	return ApiMeSettingsUpdateShiftPreferencesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MeUserSettingsApiService) MeSettingsUpdateShiftPreferencesExecute(r ApiMeSettingsUpdateShiftPreferencesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeUserSettingsApiService.MeSettingsUpdateShiftPreferences")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/settings/shiftPreferences"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphShiftPreferences == nil {
		return nil, reportError("microsoftGraphShiftPreferences is required and must be specified")
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
	localVarPostBody = r.microsoftGraphShiftPreferences
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

type ApiMeUpdateSettingsRequest struct {
	ctx _context.Context
	ApiService *MeUserSettingsApiService
	microsoftGraphUserSettings *MicrosoftGraphUserSettings
}

// New navigation property values
func (r ApiMeUpdateSettingsRequest) MicrosoftGraphUserSettings(microsoftGraphUserSettings MicrosoftGraphUserSettings) ApiMeUpdateSettingsRequest {
	r.microsoftGraphUserSettings = &microsoftGraphUserSettings
	return r
}

func (r ApiMeUpdateSettingsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.MeUpdateSettingsExecute(r)
}

/*
MeUpdateSettings Update the navigation property settings in me

Read-only. Nullable.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMeUpdateSettingsRequest
*/
func (a *MeUserSettingsApiService) MeUpdateSettings(ctx _context.Context) ApiMeUpdateSettingsRequest {
	return ApiMeUpdateSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *MeUserSettingsApiService) MeUpdateSettingsExecute(r ApiMeUpdateSettingsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeUserSettingsApiService.MeUpdateSettings")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphUserSettings == nil {
		return nil, reportError("microsoftGraphUserSettings is required and must be specified")
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
	localVarPostBody = r.microsoftGraphUserSettings
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
