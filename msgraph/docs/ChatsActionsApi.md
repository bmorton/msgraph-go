# \ChatsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatsChatInstalledAppsTeamsAppInstallationUpgrade**](ChatsActionsApi.md#ChatsChatInstalledAppsTeamsAppInstallationUpgrade) | **Post** /chats/{chat-id}/installedApps/{teamsAppInstallation-id}/microsoft.graph.upgrade | Invoke action upgrade
[**ChatsChatMembersAdd**](ChatsActionsApi.md#ChatsChatMembersAdd) | **Post** /chats/{chat-id}/members/microsoft.graph.add | Invoke action add
[**ChatsChatSendActivityNotification**](ChatsActionsApi.md#ChatsChatSendActivityNotification) | **Post** /chats/{chat-id}/microsoft.graph.sendActivityNotification | Invoke action sendActivityNotification



## ChatsChatInstalledAppsTeamsAppInstallationUpgrade

> ChatsChatInstalledAppsTeamsAppInstallationUpgrade(ctx, chatId, teamsAppInstallationId).Execute()

Invoke action upgrade

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
    teamsAppInstallationId := "teamsAppInstallationId_example" // string | key: id of teamsAppInstallation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsActionsApi.ChatsChatInstalledAppsTeamsAppInstallationUpgrade(context.Background(), chatId, teamsAppInstallationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsActionsApi.ChatsChatInstalledAppsTeamsAppInstallationUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**teamsAppInstallationId** | **string** | key: id of teamsAppInstallation | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsChatInstalledAppsTeamsAppInstallationUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsChatMembersAdd

> []*AnyOfmicrosoftGraphActionResultPart ChatsChatMembersAdd(ctx, chatId).InlineObject19(inlineObject19).Execute()

Invoke action add

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
    inlineObject19 := *openapiclient.NewInlineObject19() // InlineObject19 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsActionsApi.ChatsChatMembersAdd(context.Background(), chatId).InlineObject19(inlineObject19).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsActionsApi.ChatsChatMembersAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsChatMembersAdd`: []*AnyOfmicrosoftGraphActionResultPart
    fmt.Fprintf(os.Stdout, "Response from `ChatsActionsApi.ChatsChatMembersAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsChatMembersAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject19** | [**InlineObject19**](InlineObject19.md) |  | 

### Return type

[**[]*AnyOfmicrosoftGraphActionResultPart**](anyOf&lt;microsoft.graph.actionResultPart&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsChatSendActivityNotification

> ChatsChatSendActivityNotification(ctx, chatId).InlineObject20(inlineObject20).Execute()

Invoke action sendActivityNotification

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
    inlineObject20 := *openapiclient.NewInlineObject20() // InlineObject20 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsActionsApi.ChatsChatSendActivityNotification(context.Background(), chatId).InlineObject20(inlineObject20).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsActionsApi.ChatsChatSendActivityNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsChatSendActivityNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject20** | [**InlineObject20**](InlineObject20.md) |  | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

