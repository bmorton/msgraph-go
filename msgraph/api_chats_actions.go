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

// ChatsActionsApiService ChatsActionsApi service
type ChatsActionsApiService service

type ApiChatsChatInstalledAppsTeamsAppInstallationUpgradeRequest struct {
	ctx _context.Context
	ApiService *ChatsActionsApiService
	chatId string
	teamsAppInstallationId string
}


func (r ApiChatsChatInstalledAppsTeamsAppInstallationUpgradeRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ChatsChatInstalledAppsTeamsAppInstallationUpgradeExecute(r)
}

/*
ChatsChatInstalledAppsTeamsAppInstallationUpgrade Invoke action upgrade

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chatId key: id of chat
 @param teamsAppInstallationId key: id of teamsAppInstallation
 @return ApiChatsChatInstalledAppsTeamsAppInstallationUpgradeRequest
*/
func (a *ChatsActionsApiService) ChatsChatInstalledAppsTeamsAppInstallationUpgrade(ctx _context.Context, chatId string, teamsAppInstallationId string) ApiChatsChatInstalledAppsTeamsAppInstallationUpgradeRequest {
	return ApiChatsChatInstalledAppsTeamsAppInstallationUpgradeRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
		teamsAppInstallationId: teamsAppInstallationId,
	}
}

// Execute executes the request
func (a *ChatsActionsApiService) ChatsChatInstalledAppsTeamsAppInstallationUpgradeExecute(r ApiChatsChatInstalledAppsTeamsAppInstallationUpgradeRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsActionsApiService.ChatsChatInstalledAppsTeamsAppInstallationUpgrade")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats/{chat-id}/installedApps/{teamsAppInstallation-id}/microsoft.graph.upgrade"
	localVarPath = strings.Replace(localVarPath, "{"+"chat-id"+"}", _neturl.PathEscape(parameterToString(r.chatId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamsAppInstallation-id"+"}", _neturl.PathEscape(parameterToString(r.teamsAppInstallationId, "")), -1)

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

type ApiChatsChatMembersAddRequest struct {
	ctx _context.Context
	ApiService *ChatsActionsApiService
	chatId string
	inlineObject19 *InlineObject19
}

func (r ApiChatsChatMembersAddRequest) InlineObject19(inlineObject19 InlineObject19) ApiChatsChatMembersAddRequest {
	r.inlineObject19 = &inlineObject19
	return r
}

func (r ApiChatsChatMembersAddRequest) Execute() ([]*AnyOfmicrosoftGraphActionResultPart, *_nethttp.Response, error) {
	return r.ApiService.ChatsChatMembersAddExecute(r)
}

/*
ChatsChatMembersAdd Invoke action add

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chatId key: id of chat
 @return ApiChatsChatMembersAddRequest
*/
func (a *ChatsActionsApiService) ChatsChatMembersAdd(ctx _context.Context, chatId string) ApiChatsChatMembersAddRequest {
	return ApiChatsChatMembersAddRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
//  @return []*AnyOfmicrosoftGraphActionResultPart
func (a *ChatsActionsApiService) ChatsChatMembersAddExecute(r ApiChatsChatMembersAddRequest) ([]*AnyOfmicrosoftGraphActionResultPart, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []*AnyOfmicrosoftGraphActionResultPart
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsActionsApiService.ChatsChatMembersAdd")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats/{chat-id}/members/microsoft.graph.add"
	localVarPath = strings.Replace(localVarPath, "{"+"chat-id"+"}", _neturl.PathEscape(parameterToString(r.chatId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject19 == nil {
		return localVarReturnValue, nil, reportError("inlineObject19 is required and must be specified")
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
	localVarPostBody = r.inlineObject19
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

type ApiChatsChatSendActivityNotificationRequest struct {
	ctx _context.Context
	ApiService *ChatsActionsApiService
	chatId string
	inlineObject20 *InlineObject20
}

func (r ApiChatsChatSendActivityNotificationRequest) InlineObject20(inlineObject20 InlineObject20) ApiChatsChatSendActivityNotificationRequest {
	r.inlineObject20 = &inlineObject20
	return r
}

func (r ApiChatsChatSendActivityNotificationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ChatsChatSendActivityNotificationExecute(r)
}

/*
ChatsChatSendActivityNotification Invoke action sendActivityNotification

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chatId key: id of chat
 @return ApiChatsChatSendActivityNotificationRequest
*/
func (a *ChatsActionsApiService) ChatsChatSendActivityNotification(ctx _context.Context, chatId string) ApiChatsChatSendActivityNotificationRequest {
	return ApiChatsChatSendActivityNotificationRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
func (a *ChatsActionsApiService) ChatsChatSendActivityNotificationExecute(r ApiChatsChatSendActivityNotificationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsActionsApiService.ChatsChatSendActivityNotification")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats/{chat-id}/microsoft.graph.sendActivityNotification"
	localVarPath = strings.Replace(localVarPath, "{"+"chat-id"+"}", _neturl.PathEscape(parameterToString(r.chatId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject20 == nil {
		return nil, reportError("inlineObject20 is required and must be specified")
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
	localVarPostBody = r.inlineObject20
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
