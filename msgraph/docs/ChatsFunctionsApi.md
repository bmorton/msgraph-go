# \ChatsFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatsChatMessagesChatMessageRepliesDelta**](ChatsFunctionsApi.md#ChatsChatMessagesChatMessageRepliesDelta) | **Get** /chats/{chat-id}/messages/{chatMessage-id}/replies/microsoft.graph.delta() | Invoke function delta
[**ChatsChatMessagesDelta**](ChatsFunctionsApi.md#ChatsChatMessagesDelta) | **Get** /chats/{chat-id}/messages/microsoft.graph.delta() | Invoke function delta
[**ChatsGetAllMessages**](ChatsFunctionsApi.md#ChatsGetAllMessages) | **Get** /chats/microsoft.graph.getAllMessages() | Invoke function getAllMessages



## ChatsChatMessagesChatMessageRepliesDelta

> []*AnyOfmicrosoftGraphChatMessage ChatsChatMessagesChatMessageRepliesDelta(ctx, chatId, chatMessageId).Execute()

Invoke function delta

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chatId := "chatId_example" // string | key: id of chat
    chatMessageId := "chatMessageId_example" // string | key: id of chatMessage

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsFunctionsApi.ChatsChatMessagesChatMessageRepliesDelta(context.Background(), chatId, chatMessageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsFunctionsApi.ChatsChatMessagesChatMessageRepliesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsChatMessagesChatMessageRepliesDelta`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsFunctionsApi.ChatsChatMessagesChatMessageRepliesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsChatMessagesChatMessageRepliesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsChatMessagesDelta

> []*AnyOfmicrosoftGraphChatMessage ChatsChatMessagesDelta(ctx, chatId).Execute()

Invoke function delta

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    chatId := "chatId_example" // string | key: id of chat

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsFunctionsApi.ChatsChatMessagesDelta(context.Background(), chatId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsFunctionsApi.ChatsChatMessagesDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsChatMessagesDelta`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsFunctionsApi.ChatsChatMessagesDelta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsChatMessagesDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsGetAllMessages

> []*AnyOfmicrosoftGraphChatMessage ChatsGetAllMessages(ctx).Execute()

Invoke function getAllMessages

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsFunctionsApi.ChatsGetAllMessages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsFunctionsApi.ChatsGetAllMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsGetAllMessages`: []*AnyOfmicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsFunctionsApi.ChatsGetAllMessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatsGetAllMessagesRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphChatMessage**](anyOf&lt;microsoft.graph.chatMessage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

