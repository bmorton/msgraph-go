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

// ChatsChatApiService ChatsChatApi service
type ChatsChatApiService service

type ApiChatsChatCreateChatRequest struct {
	ctx _context.Context
	ApiService *ChatsChatApiService
	microsoftGraphChat *MicrosoftGraphChat
}

// New entity
func (r ApiChatsChatCreateChatRequest) MicrosoftGraphChat(microsoftGraphChat MicrosoftGraphChat) ApiChatsChatCreateChatRequest {
	r.microsoftGraphChat = &microsoftGraphChat
	return r
}

func (r ApiChatsChatCreateChatRequest) Execute() (MicrosoftGraphChat, *_nethttp.Response, error) {
	return r.ApiService.ChatsChatCreateChatExecute(r)
}

/*
ChatsChatCreateChat Add new entity to chats

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiChatsChatCreateChatRequest
*/
func (a *ChatsChatApiService) ChatsChatCreateChat(ctx _context.Context) ApiChatsChatCreateChatRequest {
	return ApiChatsChatCreateChatRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MicrosoftGraphChat
func (a *ChatsChatApiService) ChatsChatCreateChatExecute(r ApiChatsChatCreateChatRequest) (MicrosoftGraphChat, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphChat
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsChatApiService.ChatsChatCreateChat")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphChat == nil {
		return localVarReturnValue, nil, reportError("microsoftGraphChat is required and must be specified")
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
	localVarPostBody = r.microsoftGraphChat
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

type ApiChatsChatDeleteChatRequest struct {
	ctx _context.Context
	ApiService *ChatsChatApiService
	chatId string
	ifMatch *string
}

// ETag
func (r ApiChatsChatDeleteChatRequest) IfMatch(ifMatch string) ApiChatsChatDeleteChatRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiChatsChatDeleteChatRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ChatsChatDeleteChatExecute(r)
}

/*
ChatsChatDeleteChat Delete entity from chats

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chatId key: id of chat
 @return ApiChatsChatDeleteChatRequest
*/
func (a *ChatsChatApiService) ChatsChatDeleteChat(ctx _context.Context, chatId string) ApiChatsChatDeleteChatRequest {
	return ApiChatsChatDeleteChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
func (a *ChatsChatApiService) ChatsChatDeleteChatExecute(r ApiChatsChatDeleteChatRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsChatApiService.ChatsChatDeleteChat")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats/{chat-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"chat-id"+"}", _neturl.PathEscape(parameterToString(r.chatId, "")), -1)

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

type ApiChatsChatGetChatRequest struct {
	ctx _context.Context
	ApiService *ChatsChatApiService
	chatId string
	select_ *[]string
	expand *[]string
}

// Select properties to be returned
func (r ApiChatsChatGetChatRequest) Select_(select_ []string) ApiChatsChatGetChatRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiChatsChatGetChatRequest) Expand(expand []string) ApiChatsChatGetChatRequest {
	r.expand = &expand
	return r
}

func (r ApiChatsChatGetChatRequest) Execute() (MicrosoftGraphChat, *_nethttp.Response, error) {
	return r.ApiService.ChatsChatGetChatExecute(r)
}

/*
ChatsChatGetChat Get entity from chats by key

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chatId key: id of chat
 @return ApiChatsChatGetChatRequest
*/
func (a *ChatsChatApiService) ChatsChatGetChat(ctx _context.Context, chatId string) ApiChatsChatGetChatRequest {
	return ApiChatsChatGetChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
//  @return MicrosoftGraphChat
func (a *ChatsChatApiService) ChatsChatGetChatExecute(r ApiChatsChatGetChatRequest) (MicrosoftGraphChat, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphChat
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsChatApiService.ChatsChatGetChat")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats/{chat-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"chat-id"+"}", _neturl.PathEscape(parameterToString(r.chatId, "")), -1)

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

type ApiChatsChatListChatRequest struct {
	ctx _context.Context
	ApiService *ChatsChatApiService
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
func (r ApiChatsChatListChatRequest) Top(top int32) ApiChatsChatListChatRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiChatsChatListChatRequest) Skip(skip int32) ApiChatsChatListChatRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiChatsChatListChatRequest) Search(search string) ApiChatsChatListChatRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiChatsChatListChatRequest) Filter(filter string) ApiChatsChatListChatRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiChatsChatListChatRequest) Count(count bool) ApiChatsChatListChatRequest {
	r.count = &count
	return r
}
// Order items by property values
func (r ApiChatsChatListChatRequest) Orderby(orderby []string) ApiChatsChatListChatRequest {
	r.orderby = &orderby
	return r
}
// Select properties to be returned
func (r ApiChatsChatListChatRequest) Select_(select_ []string) ApiChatsChatListChatRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiChatsChatListChatRequest) Expand(expand []string) ApiChatsChatListChatRequest {
	r.expand = &expand
	return r
}

func (r ApiChatsChatListChatRequest) Execute() (CollectionOfChat, *_nethttp.Response, error) {
	return r.ApiService.ChatsChatListChatExecute(r)
}

/*
ChatsChatListChat Get entities from chats

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiChatsChatListChatRequest
*/
func (a *ChatsChatApiService) ChatsChatListChat(ctx _context.Context) ApiChatsChatListChatRequest {
	return ApiChatsChatListChatRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfChat
func (a *ChatsChatApiService) ChatsChatListChatExecute(r ApiChatsChatListChatRequest) (CollectionOfChat, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfChat
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsChatApiService.ChatsChatListChat")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats"

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

type ApiChatsChatUpdateChatRequest struct {
	ctx _context.Context
	ApiService *ChatsChatApiService
	chatId string
	microsoftGraphChat *MicrosoftGraphChat
}

// New property values
func (r ApiChatsChatUpdateChatRequest) MicrosoftGraphChat(microsoftGraphChat MicrosoftGraphChat) ApiChatsChatUpdateChatRequest {
	r.microsoftGraphChat = &microsoftGraphChat
	return r
}

func (r ApiChatsChatUpdateChatRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ChatsChatUpdateChatExecute(r)
}

/*
ChatsChatUpdateChat Update entity in chats

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param chatId key: id of chat
 @return ApiChatsChatUpdateChatRequest
*/
func (a *ChatsChatApiService) ChatsChatUpdateChat(ctx _context.Context, chatId string) ApiChatsChatUpdateChatRequest {
	return ApiChatsChatUpdateChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
func (a *ChatsChatApiService) ChatsChatUpdateChatExecute(r ApiChatsChatUpdateChatRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChatsChatApiService.ChatsChatUpdateChat")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/chats/{chat-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"chat-id"+"}", _neturl.PathEscape(parameterToString(r.chatId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.microsoftGraphChat == nil {
		return nil, reportError("microsoftGraphChat is required and must be specified")
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
	localVarPostBody = r.microsoftGraphChat
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
