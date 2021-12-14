# \GroupsConversationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsConversationsCreateThreads**](GroupsConversationApi.md#GroupsConversationsCreateThreads) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads | Create new navigation property to threads for groups
[**GroupsConversationsDeleteThreads**](GroupsConversationApi.md#GroupsConversationsDeleteThreads) | **Delete** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id} | Delete navigation property threads for groups
[**GroupsConversationsGetThreads**](GroupsConversationApi.md#GroupsConversationsGetThreads) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id} | Get threads from groups
[**GroupsConversationsListThreads**](GroupsConversationApi.md#GroupsConversationsListThreads) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads | Get threads from groups
[**GroupsConversationsThreadsCreatePosts**](GroupsConversationApi.md#GroupsConversationsThreadsCreatePosts) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts | Create new navigation property to posts for groups
[**GroupsConversationsThreadsDeletePosts**](GroupsConversationApi.md#GroupsConversationsThreadsDeletePosts) | **Delete** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id} | Delete navigation property posts for groups
[**GroupsConversationsThreadsGetPosts**](GroupsConversationApi.md#GroupsConversationsThreadsGetPosts) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id} | Get posts from groups
[**GroupsConversationsThreadsListPosts**](GroupsConversationApi.md#GroupsConversationsThreadsListPosts) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts | Get posts from groups
[**GroupsConversationsThreadsPostsCreateAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateAttachments) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Create new navigation property to attachments for groups
[**GroupsConversationsThreadsPostsCreateExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateExtensions) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Create new navigation property to extensions for groups
[**GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for groups
[**GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties) | **Post** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for groups
[**GroupsConversationsThreadsPostsDeleteAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsDeleteAttachments) | **Delete** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Delete navigation property attachments for groups
[**GroupsConversationsThreadsPostsDeleteExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsDeleteExtensions) | **Delete** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Delete navigation property extensions for groups
[**GroupsConversationsThreadsPostsDeleteInReplyTo**](GroupsConversationApi.md#GroupsConversationsThreadsPostsDeleteInReplyTo) | **Delete** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Delete navigation property inReplyTo for groups
[**GroupsConversationsThreadsPostsDeleteMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsDeleteMultiValueExtendedProperties) | **Delete** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for groups
[**GroupsConversationsThreadsPostsDeleteSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsDeleteSingleValueExtendedProperties) | **Delete** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for groups
[**GroupsConversationsThreadsPostsGetAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetAttachments) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Get attachments from groups
[**GroupsConversationsThreadsPostsGetExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetExtensions) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Get extensions from groups
[**GroupsConversationsThreadsPostsGetInReplyTo**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetInReplyTo) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Get inReplyTo from groups
[**GroupsConversationsThreadsPostsGetMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetMultiValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsGetSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsGetSingleValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsListAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListAttachments) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Get attachments from groups
[**GroupsConversationsThreadsPostsListExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListExtensions) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Get extensions from groups
[**GroupsConversationsThreadsPostsListMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListMultiValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsListSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsListSingleValueExtendedProperties) | **Get** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from groups
[**GroupsConversationsThreadsPostsUpdateAttachments**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateAttachments) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Update the navigation property attachments in groups
[**GroupsConversationsThreadsPostsUpdateExtensions**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateExtensions) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Update the navigation property extensions in groups
[**GroupsConversationsThreadsPostsUpdateInReplyTo**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateInReplyTo) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Update the navigation property inReplyTo in groups
[**GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in groups
[**GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties**](GroupsConversationApi.md#GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in groups
[**GroupsConversationsThreadsUpdatePosts**](GroupsConversationApi.md#GroupsConversationsThreadsUpdatePosts) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id}/posts/{post-id} | Update the navigation property posts in groups
[**GroupsConversationsUpdateThreads**](GroupsConversationApi.md#GroupsConversationsUpdateThreads) | **Patch** /groups/{group-id}/conversations/{conversation-id}/threads/{conversationThread-id} | Update the navigation property threads in groups
[**GroupsCreateConversations**](GroupsConversationApi.md#GroupsCreateConversations) | **Post** /groups/{group-id}/conversations | Create new navigation property to conversations for groups
[**GroupsDeleteConversations**](GroupsConversationApi.md#GroupsDeleteConversations) | **Delete** /groups/{group-id}/conversations/{conversation-id} | Delete navigation property conversations for groups
[**GroupsGetConversations**](GroupsConversationApi.md#GroupsGetConversations) | **Get** /groups/{group-id}/conversations/{conversation-id} | Get conversations from groups
[**GroupsListConversations**](GroupsConversationApi.md#GroupsListConversations) | **Get** /groups/{group-id}/conversations | Get conversations from groups
[**GroupsUpdateConversations**](GroupsConversationApi.md#GroupsUpdateConversations) | **Patch** /groups/{group-id}/conversations/{conversation-id} | Update the navigation property conversations in groups



## GroupsConversationsCreateThreads

> MicrosoftGraphConversationThread GroupsConversationsCreateThreads(ctx, groupId, conversationId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()

Create new navigation property to threads for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    microsoftGraphConversationThread := *openapiclient.NewMicrosoftGraphConversationThread() // MicrosoftGraphConversationThread | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsCreateThreads(context.Background(), groupId, conversationId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsCreateThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsCreateThreads`: MicrosoftGraphConversationThread
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsCreateThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsCreateThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphConversationThread** | [**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md) | New navigation property | 

### Return type

[**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsDeleteThreads

> GroupsConversationsDeleteThreads(ctx, groupId, conversationId, conversationThreadId).IfMatch(ifMatch).Execute()

Delete navigation property threads for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsDeleteThreads(context.Background(), groupId, conversationId, conversationThreadId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsDeleteThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsDeleteThreadsRequest struct via the builder pattern


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


## GroupsConversationsGetThreads

> MicrosoftGraphConversationThread GroupsConversationsGetThreads(ctx, groupId, conversationId, conversationThreadId).Select_(select_).Expand(expand).Execute()

Get threads from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsGetThreads(context.Background(), groupId, conversationId, conversationThreadId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsGetThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsGetThreads`: MicrosoftGraphConversationThread
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsGetThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsGetThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsListThreads

> CollectionOfConversationThread GroupsConversationsListThreads(ctx, groupId, conversationId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get threads from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsListThreads(context.Background(), groupId, conversationId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsListThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsListThreads`: CollectionOfConversationThread
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsListThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsListThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfConversationThread**](CollectionOfConversationThread.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsCreatePosts

> MicrosoftGraphPost GroupsConversationsThreadsCreatePosts(ctx, groupId, conversationId, conversationThreadId).MicrosoftGraphPost(microsoftGraphPost).Execute()

Create new navigation property to posts for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    microsoftGraphPost := *openapiclient.NewMicrosoftGraphPost() // MicrosoftGraphPost | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsCreatePosts(context.Background(), groupId, conversationId, conversationThreadId).MicrosoftGraphPost(microsoftGraphPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsCreatePosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsCreatePosts`: MicrosoftGraphPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsCreatePosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsCreatePostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md) | New navigation property | 

### Return type

[**MicrosoftGraphPost**](MicrosoftGraphPost.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsDeletePosts

> GroupsConversationsThreadsDeletePosts(ctx, groupId, conversationId, conversationThreadId, postId).IfMatch(ifMatch).Execute()

Delete navigation property posts for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsDeletePosts(context.Background(), groupId, conversationId, conversationThreadId, postId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsDeletePosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsDeletePostsRequest struct via the builder pattern


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


## GroupsConversationsThreadsGetPosts

> MicrosoftGraphPost GroupsConversationsThreadsGetPosts(ctx, groupId, conversationId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()

Get posts from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsGetPosts(context.Background(), groupId, conversationId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsGetPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsGetPosts`: MicrosoftGraphPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsGetPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsGetPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPost**](MicrosoftGraphPost.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsListPosts

> CollectionOfPost GroupsConversationsThreadsListPosts(ctx, groupId, conversationId, conversationThreadId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get posts from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsListPosts(context.Background(), groupId, conversationId, conversationThreadId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsListPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsListPosts`: CollectionOfPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsListPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsListPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfPost**](CollectionOfPost.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsCreateAttachments

> MicrosoftGraphAttachment GroupsConversationsThreadsPostsCreateAttachments(ctx, groupId, conversationId, conversationThreadId, postId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

Create new navigation property to attachments for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsCreateAttachments(context.Background(), groupId, conversationId, conversationThreadId, postId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsCreateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsCreateAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsCreateAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsCreateAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphAttachment** | [**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md) | New navigation property | 

### Return type

[**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsCreateExtensions

> MicrosoftGraphExtension GroupsConversationsThreadsPostsCreateExtensions(ctx, groupId, conversationId, conversationThreadId, postId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsCreateExtensions(context.Background(), groupId, conversationId, conversationThreadId, postId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsCreateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

Create new navigation property to multiValueExtendedProperties for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | New navigation property | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

Create new navigation property to singleValueExtendedProperties for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | New navigation property | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsDeleteAttachments

> GroupsConversationsThreadsPostsDeleteAttachments(ctx, groupId, conversationId, conversationThreadId, postId, attachmentId).IfMatch(ifMatch).Execute()

Delete navigation property attachments for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    attachmentId := "attachmentId_example" // string | key: id of attachment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsDeleteAttachments(context.Background(), groupId, conversationId, conversationThreadId, postId, attachmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsDeleteAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsDeleteAttachmentsRequest struct via the builder pattern


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


## GroupsConversationsThreadsPostsDeleteExtensions

> GroupsConversationsThreadsPostsDeleteExtensions(ctx, groupId, conversationId, conversationThreadId, postId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsDeleteExtensions(context.Background(), groupId, conversationId, conversationThreadId, postId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsDeleteExtensionsRequest struct via the builder pattern


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


## GroupsConversationsThreadsPostsDeleteInReplyTo

> GroupsConversationsThreadsPostsDeleteInReplyTo(ctx, groupId, conversationId, conversationThreadId, postId).IfMatch(ifMatch).Execute()

Delete navigation property inReplyTo for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsDeleteInReplyTo(context.Background(), groupId, conversationId, conversationThreadId, postId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsDeleteInReplyTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsDeleteInReplyToRequest struct via the builder pattern


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


## GroupsConversationsThreadsPostsDeleteMultiValueExtendedProperties

> GroupsConversationsThreadsPostsDeleteMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

Delete navigation property multiValueExtendedProperties for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsDeleteMultiValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsConversationsThreadsPostsDeleteSingleValueExtendedProperties

> GroupsConversationsThreadsPostsDeleteSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

Delete navigation property singleValueExtendedProperties for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsDeleteSingleValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsConversationsThreadsPostsGetAttachments

> MicrosoftGraphAttachment GroupsConversationsThreadsPostsGetAttachments(ctx, groupId, conversationId, conversationThreadId, postId, attachmentId).Select_(select_).Expand(expand).Execute()

Get attachments from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    attachmentId := "attachmentId_example" // string | key: id of attachment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsGetAttachments(context.Background(), groupId, conversationId, conversationThreadId, postId, attachmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsGetAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsGetAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsGetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsGetAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsGetExtensions

> MicrosoftGraphExtension GroupsConversationsThreadsPostsGetExtensions(ctx, groupId, conversationId, conversationThreadId, postId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsGetExtensions(context.Background(), groupId, conversationId, conversationThreadId, postId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsGetInReplyTo

> MicrosoftGraphPost GroupsConversationsThreadsPostsGetInReplyTo(ctx, groupId, conversationId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()

Get inReplyTo from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsGetInReplyTo(context.Background(), groupId, conversationId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsGetInReplyTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsGetInReplyTo`: MicrosoftGraphPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsGetInReplyTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsGetInReplyToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPost**](MicrosoftGraphPost.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsConversationsThreadsPostsGetMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

Get multiValueExtendedProperties from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsGetMultiValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsGetMultiValueExtendedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsConversationsThreadsPostsGetSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

Get singleValueExtendedProperties from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsGetSingleValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsGetSingleValueExtendedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsListAttachments

> CollectionOfAttachment GroupsConversationsThreadsPostsListAttachments(ctx, groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get attachments from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsListAttachments(context.Background(), groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsListAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsListAttachments`: CollectionOfAttachment
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsListAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsListAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfAttachment**](CollectionOfAttachment.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsListExtensions

> CollectionOfExtension GroupsConversationsThreadsPostsListExtensions(ctx, groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsListExtensions(context.Background(), groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsListExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfExtension**](CollectionOfExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty GroupsConversationsThreadsPostsListMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get multiValueExtendedProperties from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
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
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsListMultiValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsListMultiValueExtendedPropertiesRequest struct via the builder pattern


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

[**CollectionOfMultiValueLegacyExtendedProperty**](CollectionOfMultiValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty GroupsConversationsThreadsPostsListSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get singleValueExtendedProperties from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
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
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsListSingleValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsConversationsThreadsPostsListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsConversationsThreadsPostsListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsListSingleValueExtendedPropertiesRequest struct via the builder pattern


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

[**CollectionOfSingleValueLegacyExtendedProperty**](CollectionOfSingleValueLegacyExtendedProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsConversationsThreadsPostsUpdateAttachments

> GroupsConversationsThreadsPostsUpdateAttachments(ctx, groupId, conversationId, conversationThreadId, postId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

Update the navigation property attachments in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    attachmentId := "attachmentId_example" // string | key: id of attachment
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsUpdateAttachments(context.Background(), groupId, conversationId, conversationThreadId, postId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsUpdateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsUpdateAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **microsoftGraphAttachment** | [**MicrosoftGraphAttachment**](MicrosoftGraphAttachment.md) | New navigation property values | 

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


## GroupsConversationsThreadsPostsUpdateExtensions

> GroupsConversationsThreadsPostsUpdateExtensions(ctx, groupId, conversationId, conversationThreadId, postId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsUpdateExtensions(context.Background(), groupId, conversationId, conversationThreadId, postId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsUpdateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property values | 

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


## GroupsConversationsThreadsPostsUpdateInReplyTo

> GroupsConversationsThreadsPostsUpdateInReplyTo(ctx, groupId, conversationId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()

Update the navigation property inReplyTo in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphPost := *openapiclient.NewMicrosoftGraphPost() // MicrosoftGraphPost | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsUpdateInReplyTo(context.Background(), groupId, conversationId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsUpdateInReplyTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsUpdateInReplyToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md) | New navigation property values | 

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


## GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties

> GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

Update the navigation property multiValueExtendedProperties in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **microsoftGraphMultiValueLegacyExtendedProperty** | [**MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | New navigation property values | 

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


## GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties

> GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties(ctx, groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

Update the navigation property singleValueExtendedProperties in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties(context.Background(), groupId, conversationId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsPostsUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsPostsUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **microsoftGraphSingleValueLegacyExtendedProperty** | [**MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | New navigation property values | 

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


## GroupsConversationsThreadsUpdatePosts

> GroupsConversationsThreadsUpdatePosts(ctx, groupId, conversationId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()

Update the navigation property posts in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphPost := *openapiclient.NewMicrosoftGraphPost() // MicrosoftGraphPost | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsThreadsUpdatePosts(context.Background(), groupId, conversationId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsThreadsUpdatePosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsThreadsUpdatePostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **microsoftGraphPost** | [**MicrosoftGraphPost**](MicrosoftGraphPost.md) | New navigation property values | 

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


## GroupsConversationsUpdateThreads

> GroupsConversationsUpdateThreads(ctx, groupId, conversationId, conversationThreadId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()

Update the navigation property threads in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    microsoftGraphConversationThread := *openapiclient.NewMicrosoftGraphConversationThread() // MicrosoftGraphConversationThread | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsConversationsUpdateThreads(context.Background(), groupId, conversationId, conversationThreadId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsConversationsUpdateThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsConversationsUpdateThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **microsoftGraphConversationThread** | [**MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md) | New navigation property values | 

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


## GroupsCreateConversations

> MicrosoftGraphConversation GroupsCreateConversations(ctx, groupId).MicrosoftGraphConversation(microsoftGraphConversation).Execute()

Create new navigation property to conversations for groups



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
    groupId := "groupId_example" // string | key: id of group
    microsoftGraphConversation := *openapiclient.NewMicrosoftGraphConversation() // MicrosoftGraphConversation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsCreateConversations(context.Background(), groupId).MicrosoftGraphConversation(microsoftGraphConversation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsCreateConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateConversations`: MicrosoftGraphConversation
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsCreateConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphConversation** | [**MicrosoftGraphConversation**](MicrosoftGraphConversation.md) | New navigation property | 

### Return type

[**MicrosoftGraphConversation**](MicrosoftGraphConversation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsDeleteConversations

> GroupsDeleteConversations(ctx, groupId, conversationId).IfMatch(ifMatch).Execute()

Delete navigation property conversations for groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsDeleteConversations(context.Background(), groupId, conversationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsDeleteConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteConversationsRequest struct via the builder pattern


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


## GroupsGetConversations

> MicrosoftGraphConversation GroupsGetConversations(ctx, groupId, conversationId).Select_(select_).Execute()

Get conversations from groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsGetConversations(context.Background(), groupId, conversationId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsGetConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetConversations`: MicrosoftGraphConversation
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsGetConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphConversation**](MicrosoftGraphConversation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsListConversations

> CollectionOfConversation GroupsListConversations(ctx, groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get conversations from groups



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
    groupId := "groupId_example" // string | key: id of group
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsListConversations(context.Background(), groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsListConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsListConversations`: CollectionOfConversation
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationApi.GroupsListConversations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfConversation**](CollectionOfConversation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsUpdateConversations

> GroupsUpdateConversations(ctx, groupId, conversationId).MicrosoftGraphConversation(microsoftGraphConversation).Execute()

Update the navigation property conversations in groups



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
    groupId := "groupId_example" // string | key: id of group
    conversationId := "conversationId_example" // string | key: id of conversation
    microsoftGraphConversation := *openapiclient.NewMicrosoftGraphConversation() // MicrosoftGraphConversation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationApi.GroupsUpdateConversations(context.Background(), groupId, conversationId).MicrosoftGraphConversation(microsoftGraphConversation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationApi.GroupsUpdateConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationId** | **string** | key: id of conversation | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphConversation** | [**MicrosoftGraphConversation**](MicrosoftGraphConversation.md) | New navigation property values | 

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

