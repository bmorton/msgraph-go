# \ChatsConversationMemberApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatsCreateMembers**](ChatsConversationMemberApi.md#ChatsCreateMembers) | **Post** /chats/{chat-id}/members | Create new navigation property to members for chats
[**ChatsDeleteMembers**](ChatsConversationMemberApi.md#ChatsDeleteMembers) | **Delete** /chats/{chat-id}/members/{conversationMember-id} | Delete navigation property members for chats
[**ChatsGetMembers**](ChatsConversationMemberApi.md#ChatsGetMembers) | **Get** /chats/{chat-id}/members/{conversationMember-id} | Get members from chats
[**ChatsListMembers**](ChatsConversationMemberApi.md#ChatsListMembers) | **Get** /chats/{chat-id}/members | Get members from chats
[**ChatsUpdateMembers**](ChatsConversationMemberApi.md#ChatsUpdateMembers) | **Patch** /chats/{chat-id}/members/{conversationMember-id} | Update the navigation property members in chats



## ChatsCreateMembers

> MicrosoftGraphConversationMember ChatsCreateMembers(ctx, chatId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()

Create new navigation property to members for chats



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
    microsoftGraphConversationMember := *openapiclient.NewMicrosoftGraphConversationMember() // MicrosoftGraphConversationMember | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsConversationMemberApi.ChatsCreateMembers(context.Background(), chatId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsConversationMemberApi.ChatsCreateMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsCreateMembers`: MicrosoftGraphConversationMember
    fmt.Fprintf(os.Stdout, "Response from `ChatsConversationMemberApi.ChatsCreateMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsCreateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | New navigation property | 

### Return type

[**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsDeleteMembers

> ChatsDeleteMembers(ctx, chatId, conversationMemberId).IfMatch(ifMatch).Execute()

Delete navigation property members for chats



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsConversationMemberApi.ChatsDeleteMembers(context.Background(), chatId, conversationMemberId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsConversationMemberApi.ChatsDeleteMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsDeleteMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | ETag | 

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


## ChatsGetMembers

> MicrosoftGraphConversationMember ChatsGetMembers(ctx, chatId, conversationMemberId).Select_(select_).Expand(expand).Execute()

Get members from chats



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsConversationMemberApi.ChatsGetMembers(context.Background(), chatId, conversationMemberId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsConversationMemberApi.ChatsGetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsGetMembers`: MicrosoftGraphConversationMember
    fmt.Fprintf(os.Stdout, "Response from `ChatsConversationMemberApi.ChatsGetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsListMembers

> CollectionOfConversationMember ChatsListMembers(ctx, chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get members from chats



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
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsConversationMemberApi.ChatsListMembers(context.Background(), chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsConversationMemberApi.ChatsListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsListMembers`: CollectionOfConversationMember
    fmt.Fprintf(os.Stdout, "Response from `ChatsConversationMemberApi.ChatsListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfConversationMember**](CollectionOfConversationMember.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsUpdateMembers

> ChatsUpdateMembers(ctx, chatId, conversationMemberId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()

Update the navigation property members in chats



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
    conversationMemberId := "conversationMemberId_example" // string | key: id of conversationMember
    microsoftGraphConversationMember := *openapiclient.NewMicrosoftGraphConversationMember() // MicrosoftGraphConversationMember | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsConversationMemberApi.ChatsUpdateMembers(context.Background(), chatId, conversationMemberId).MicrosoftGraphConversationMember(microsoftGraphConversationMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsConversationMemberApi.ChatsUpdateMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**conversationMemberId** | **string** | key: id of conversationMember | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsUpdateMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphConversationMember** | [**MicrosoftGraphConversationMember**](MicrosoftGraphConversationMember.md) | New navigation property values | 

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

