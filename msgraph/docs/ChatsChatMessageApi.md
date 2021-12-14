# \ChatsChatMessageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatsCreateMessages**](ChatsChatMessageApi.md#ChatsCreateMessages) | **Post** /chats/{chat-id}/messages | Create new navigation property to messages for chats
[**ChatsDeleteMessages**](ChatsChatMessageApi.md#ChatsDeleteMessages) | **Delete** /chats/{chat-id}/messages/{chatMessage-id} | Delete navigation property messages for chats
[**ChatsGetMessages**](ChatsChatMessageApi.md#ChatsGetMessages) | **Get** /chats/{chat-id}/messages/{chatMessage-id} | Get messages from chats
[**ChatsListMessages**](ChatsChatMessageApi.md#ChatsListMessages) | **Get** /chats/{chat-id}/messages | Get messages from chats
[**ChatsMessagesCreateHostedContents**](ChatsChatMessageApi.md#ChatsMessagesCreateHostedContents) | **Post** /chats/{chat-id}/messages/{chatMessage-id}/hostedContents | Create new navigation property to hostedContents for chats
[**ChatsMessagesCreateReplies**](ChatsChatMessageApi.md#ChatsMessagesCreateReplies) | **Post** /chats/{chat-id}/messages/{chatMessage-id}/replies | Create new navigation property to replies for chats
[**ChatsMessagesDeleteHostedContents**](ChatsChatMessageApi.md#ChatsMessagesDeleteHostedContents) | **Delete** /chats/{chat-id}/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Delete navigation property hostedContents for chats
[**ChatsMessagesDeleteReplies**](ChatsChatMessageApi.md#ChatsMessagesDeleteReplies) | **Delete** /chats/{chat-id}/messages/{chatMessage-id}/replies/{chatMessage-id1} | Delete navigation property replies for chats
[**ChatsMessagesGetHostedContents**](ChatsChatMessageApi.md#ChatsMessagesGetHostedContents) | **Get** /chats/{chat-id}/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Get hostedContents from chats
[**ChatsMessagesGetReplies**](ChatsChatMessageApi.md#ChatsMessagesGetReplies) | **Get** /chats/{chat-id}/messages/{chatMessage-id}/replies/{chatMessage-id1} | Get replies from chats
[**ChatsMessagesListHostedContents**](ChatsChatMessageApi.md#ChatsMessagesListHostedContents) | **Get** /chats/{chat-id}/messages/{chatMessage-id}/hostedContents | Get hostedContents from chats
[**ChatsMessagesListReplies**](ChatsChatMessageApi.md#ChatsMessagesListReplies) | **Get** /chats/{chat-id}/messages/{chatMessage-id}/replies | Get replies from chats
[**ChatsMessagesUpdateHostedContents**](ChatsChatMessageApi.md#ChatsMessagesUpdateHostedContents) | **Patch** /chats/{chat-id}/messages/{chatMessage-id}/hostedContents/{chatMessageHostedContent-id} | Update the navigation property hostedContents in chats
[**ChatsMessagesUpdateReplies**](ChatsChatMessageApi.md#ChatsMessagesUpdateReplies) | **Patch** /chats/{chat-id}/messages/{chatMessage-id}/replies/{chatMessage-id1} | Update the navigation property replies in chats
[**ChatsUpdateMessages**](ChatsChatMessageApi.md#ChatsUpdateMessages) | **Patch** /chats/{chat-id}/messages/{chatMessage-id} | Update the navigation property messages in chats



## ChatsCreateMessages

> MicrosoftGraphChatMessage ChatsCreateMessages(ctx, chatId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Create new navigation property to messages for chats



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
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsCreateMessages(context.Background(), chatId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsCreateMessages`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsCreateMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsCreateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsDeleteMessages

> ChatsDeleteMessages(ctx, chatId, chatMessageId).IfMatch(ifMatch).Execute()

Delete navigation property messages for chats



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsDeleteMessages(context.Background(), chatId, chatMessageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsDeleteMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsDeleteMessagesRequest struct via the builder pattern


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


## ChatsGetMessages

> MicrosoftGraphChatMessage ChatsGetMessages(ctx, chatId, chatMessageId).Select_(select_).Expand(expand).Execute()

Get messages from chats



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsGetMessages(context.Background(), chatId, chatMessageId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsGetMessages`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsListMessages

> CollectionOfChatMessage ChatsListMessages(ctx, chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get messages from chats



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
    resp, r, err := api_client.ChatsChatMessageApi.ChatsListMessages(context.Background(), chatId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsListMessages`: CollectionOfChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsListMessagesRequest struct via the builder pattern


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

[**CollectionOfChatMessage**](CollectionOfChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsMessagesCreateHostedContents

> MicrosoftGraphChatMessageHostedContent ChatsMessagesCreateHostedContents(ctx, chatId, chatMessageId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()

Create new navigation property to hostedContents for chats



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
    microsoftGraphChatMessageHostedContent := *openapiclient.NewMicrosoftGraphChatMessageHostedContent() // MicrosoftGraphChatMessageHostedContent | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesCreateHostedContents(context.Background(), chatId, chatMessageId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesCreateHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsMessagesCreateHostedContents`: MicrosoftGraphChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsMessagesCreateHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesCreateHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChatMessageHostedContent** | [**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsMessagesCreateReplies

> MicrosoftGraphChatMessage ChatsMessagesCreateReplies(ctx, chatId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Create new navigation property to replies for chats



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
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesCreateReplies(context.Background(), chatId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesCreateReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsMessagesCreateReplies`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsMessagesCreateReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesCreateRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsMessagesDeleteHostedContents

> ChatsMessagesDeleteHostedContents(ctx, chatId, chatMessageId, chatMessageHostedContentId).IfMatch(ifMatch).Execute()

Delete navigation property hostedContents for chats



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
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesDeleteHostedContents(context.Background(), chatId, chatMessageId, chatMessageHostedContentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesDeleteHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesDeleteHostedContentsRequest struct via the builder pattern


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


## ChatsMessagesDeleteReplies

> ChatsMessagesDeleteReplies(ctx, chatId, chatMessageId, chatMessageId1).IfMatch(ifMatch).Execute()

Delete navigation property replies for chats



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
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesDeleteReplies(context.Background(), chatId, chatMessageId, chatMessageId1).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesDeleteReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesDeleteRepliesRequest struct via the builder pattern


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


## ChatsMessagesGetHostedContents

> MicrosoftGraphChatMessageHostedContent ChatsMessagesGetHostedContents(ctx, chatId, chatMessageId, chatMessageHostedContentId).Select_(select_).Expand(expand).Execute()

Get hostedContents from chats



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
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesGetHostedContents(context.Background(), chatId, chatMessageId, chatMessageHostedContentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesGetHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsMessagesGetHostedContents`: MicrosoftGraphChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsMessagesGetHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesGetHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsMessagesGetReplies

> MicrosoftGraphChatMessage ChatsMessagesGetReplies(ctx, chatId, chatMessageId, chatMessageId1).Select_(select_).Expand(expand).Execute()

Get replies from chats



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
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesGetReplies(context.Background(), chatId, chatMessageId, chatMessageId1).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesGetReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsMessagesGetReplies`: MicrosoftGraphChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsMessagesGetReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesGetRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsMessagesListHostedContents

> CollectionOfChatMessageHostedContent ChatsMessagesListHostedContents(ctx, chatId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get hostedContents from chats



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
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesListHostedContents(context.Background(), chatId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesListHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsMessagesListHostedContents`: CollectionOfChatMessageHostedContent
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsMessagesListHostedContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesListHostedContentsRequest struct via the builder pattern


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

[**CollectionOfChatMessageHostedContent**](CollectionOfChatMessageHostedContent.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsMessagesListReplies

> CollectionOfChatMessage ChatsMessagesListReplies(ctx, chatId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get replies from chats



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
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesListReplies(context.Background(), chatId, chatMessageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesListReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChatsMessagesListReplies`: CollectionOfChatMessage
    fmt.Fprintf(os.Stdout, "Response from `ChatsChatMessageApi.ChatsMessagesListReplies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesListRepliesRequest struct via the builder pattern


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

[**CollectionOfChatMessage**](CollectionOfChatMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatsMessagesUpdateHostedContents

> ChatsMessagesUpdateHostedContents(ctx, chatId, chatMessageId, chatMessageHostedContentId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()

Update the navigation property hostedContents in chats



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
    chatMessageHostedContentId := "chatMessageHostedContentId_example" // string | key: id of chatMessageHostedContent
    microsoftGraphChatMessageHostedContent := *openapiclient.NewMicrosoftGraphChatMessageHostedContent() // MicrosoftGraphChatMessageHostedContent | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesUpdateHostedContents(context.Background(), chatId, chatMessageId, chatMessageHostedContentId).MicrosoftGraphChatMessageHostedContent(microsoftGraphChatMessageHostedContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesUpdateHostedContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageHostedContentId** | **string** | key: id of chatMessageHostedContent | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesUpdateHostedContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphChatMessageHostedContent** | [**MicrosoftGraphChatMessageHostedContent**](MicrosoftGraphChatMessageHostedContent.md) | New navigation property values | 

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


## ChatsMessagesUpdateReplies

> ChatsMessagesUpdateReplies(ctx, chatId, chatMessageId, chatMessageId1).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Update the navigation property replies in chats



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
    chatMessageId1 := "chatMessageId1_example" // string | key: id of chatMessage
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsMessagesUpdateReplies(context.Background(), chatId, chatMessageId, chatMessageId1).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsMessagesUpdateReplies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 
**chatMessageId1** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsMessagesUpdateRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property values | 

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


## ChatsUpdateMessages

> ChatsUpdateMessages(ctx, chatId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()

Update the navigation property messages in chats



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
    microsoftGraphChatMessage := *openapiclient.NewMicrosoftGraphChatMessage() // MicrosoftGraphChatMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ChatsChatMessageApi.ChatsUpdateMessages(context.Background(), chatId, chatMessageId).MicrosoftGraphChatMessage(microsoftGraphChatMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChatsChatMessageApi.ChatsUpdateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | key: id of chat | 
**chatMessageId** | **string** | key: id of chatMessage | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatsUpdateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphChatMessage** | [**MicrosoftGraphChatMessage**](MicrosoftGraphChatMessage.md) | New navigation property values | 

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

