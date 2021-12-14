# \GroupsConversationThreadApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateThreads**](GroupsConversationThreadApi.md#GroupsCreateThreads) | **Post** /groups/{group-id}/threads | Create new navigation property to threads for groups
[**GroupsDeleteThreads**](GroupsConversationThreadApi.md#GroupsDeleteThreads) | **Delete** /groups/{group-id}/threads/{conversationThread-id} | Delete navigation property threads for groups
[**GroupsGetThreads**](GroupsConversationThreadApi.md#GroupsGetThreads) | **Get** /groups/{group-id}/threads/{conversationThread-id} | Get threads from groups
[**GroupsListThreads**](GroupsConversationThreadApi.md#GroupsListThreads) | **Get** /groups/{group-id}/threads | Get threads from groups
[**GroupsThreadsCreatePosts**](GroupsConversationThreadApi.md#GroupsThreadsCreatePosts) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts | Create new navigation property to posts for groups
[**GroupsThreadsDeletePosts**](GroupsConversationThreadApi.md#GroupsThreadsDeletePosts) | **Delete** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id} | Delete navigation property posts for groups
[**GroupsThreadsGetPosts**](GroupsConversationThreadApi.md#GroupsThreadsGetPosts) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id} | Get posts from groups
[**GroupsThreadsListPosts**](GroupsConversationThreadApi.md#GroupsThreadsListPosts) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts | Get posts from groups
[**GroupsThreadsPostsCreateAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateAttachments) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Create new navigation property to attachments for groups
[**GroupsThreadsPostsCreateExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateExtensions) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Create new navigation property to extensions for groups
[**GroupsThreadsPostsCreateMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateMultiValueExtendedProperties) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for groups
[**GroupsThreadsPostsCreateSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsCreateSingleValueExtendedProperties) | **Post** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for groups
[**GroupsThreadsPostsDeleteAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsDeleteAttachments) | **Delete** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Delete navigation property attachments for groups
[**GroupsThreadsPostsDeleteExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsDeleteExtensions) | **Delete** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Delete navigation property extensions for groups
[**GroupsThreadsPostsDeleteInReplyTo**](GroupsConversationThreadApi.md#GroupsThreadsPostsDeleteInReplyTo) | **Delete** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Delete navigation property inReplyTo for groups
[**GroupsThreadsPostsDeleteMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsDeleteMultiValueExtendedProperties) | **Delete** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for groups
[**GroupsThreadsPostsDeleteSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsDeleteSingleValueExtendedProperties) | **Delete** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for groups
[**GroupsThreadsPostsGetAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetAttachments) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Get attachments from groups
[**GroupsThreadsPostsGetExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetExtensions) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Get extensions from groups
[**GroupsThreadsPostsGetInReplyTo**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetInReplyTo) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Get inReplyTo from groups
[**GroupsThreadsPostsGetMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetMultiValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from groups
[**GroupsThreadsPostsGetSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsGetSingleValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from groups
[**GroupsThreadsPostsListAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsListAttachments) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments | Get attachments from groups
[**GroupsThreadsPostsListExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsListExtensions) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions | Get extensions from groups
[**GroupsThreadsPostsListMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsListMultiValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from groups
[**GroupsThreadsPostsListSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsListSingleValueExtendedProperties) | **Get** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from groups
[**GroupsThreadsPostsUpdateAttachments**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateAttachments) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/attachments/{attachment-id} | Update the navigation property attachments in groups
[**GroupsThreadsPostsUpdateExtensions**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateExtensions) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/extensions/{extension-id} | Update the navigation property extensions in groups
[**GroupsThreadsPostsUpdateInReplyTo**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateInReplyTo) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/inReplyTo | Update the navigation property inReplyTo in groups
[**GroupsThreadsPostsUpdateMultiValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateMultiValueExtendedProperties) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in groups
[**GroupsThreadsPostsUpdateSingleValueExtendedProperties**](GroupsConversationThreadApi.md#GroupsThreadsPostsUpdateSingleValueExtendedProperties) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in groups
[**GroupsThreadsUpdatePosts**](GroupsConversationThreadApi.md#GroupsThreadsUpdatePosts) | **Patch** /groups/{group-id}/threads/{conversationThread-id}/posts/{post-id} | Update the navigation property posts in groups
[**GroupsUpdateThreads**](GroupsConversationThreadApi.md#GroupsUpdateThreads) | **Patch** /groups/{group-id}/threads/{conversationThread-id} | Update the navigation property threads in groups



## GroupsCreateThreads

> MicrosoftGraphConversationThread GroupsCreateThreads(ctx, groupId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()

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
    microsoftGraphConversationThread := *openapiclient.NewMicrosoftGraphConversationThread() // MicrosoftGraphConversationThread | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsCreateThreads(context.Background(), groupId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsCreateThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateThreads`: MicrosoftGraphConversationThread
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsCreateThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateThreadsRequest struct via the builder pattern


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


## GroupsDeleteThreads

> GroupsDeleteThreads(ctx, groupId, conversationThreadId).IfMatch(ifMatch).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsDeleteThreads(context.Background(), groupId, conversationThreadId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsDeleteThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteThreadsRequest struct via the builder pattern


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


## GroupsGetThreads

> MicrosoftGraphConversationThread GroupsGetThreads(ctx, groupId, conversationThreadId).Select_(select_).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsGetThreads(context.Background(), groupId, conversationThreadId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsGetThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetThreads`: MicrosoftGraphConversationThread
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsGetThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

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


## GroupsListThreads

> CollectionOfConversationThread GroupsListThreads(ctx, groupId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

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
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsListThreads(context.Background(), groupId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsListThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsListThreads`: CollectionOfConversationThread
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsListThreads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

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


## GroupsThreadsCreatePosts

> MicrosoftGraphPost GroupsThreadsCreatePosts(ctx, groupId, conversationThreadId).MicrosoftGraphPost(microsoftGraphPost).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    microsoftGraphPost := *openapiclient.NewMicrosoftGraphPost() // MicrosoftGraphPost | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsCreatePosts(context.Background(), groupId, conversationThreadId).MicrosoftGraphPost(microsoftGraphPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsCreatePosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsCreatePosts`: MicrosoftGraphPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsCreatePosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsCreatePostsRequest struct via the builder pattern


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


## GroupsThreadsDeletePosts

> GroupsThreadsDeletePosts(ctx, groupId, conversationThreadId, postId).IfMatch(ifMatch).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsDeletePosts(context.Background(), groupId, conversationThreadId, postId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsDeletePosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsDeletePostsRequest struct via the builder pattern


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


## GroupsThreadsGetPosts

> MicrosoftGraphPost GroupsThreadsGetPosts(ctx, groupId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsGetPosts(context.Background(), groupId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsGetPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsGetPosts`: MicrosoftGraphPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsGetPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsGetPostsRequest struct via the builder pattern


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


## GroupsThreadsListPosts

> CollectionOfPost GroupsThreadsListPosts(ctx, groupId, conversationThreadId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsListPosts(context.Background(), groupId, conversationThreadId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsListPosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsListPosts`: CollectionOfPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsListPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsListPostsRequest struct via the builder pattern


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


## GroupsThreadsPostsCreateAttachments

> MicrosoftGraphAttachment GroupsThreadsPostsCreateAttachments(ctx, groupId, conversationThreadId, postId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsCreateAttachments(context.Background(), groupId, conversationThreadId, postId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsCreateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsCreateAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsCreateAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsCreateAttachmentsRequest struct via the builder pattern


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


## GroupsThreadsPostsCreateExtensions

> MicrosoftGraphExtension GroupsThreadsPostsCreateExtensions(ctx, groupId, conversationThreadId, postId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsCreateExtensions(context.Background(), groupId, conversationThreadId, postId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsCreateExtensionsRequest struct via the builder pattern


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


## GroupsThreadsPostsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsThreadsPostsCreateMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsCreateMultiValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsThreadsPostsCreateSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsCreateSingleValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsDeleteAttachments

> GroupsThreadsPostsDeleteAttachments(ctx, groupId, conversationThreadId, postId, attachmentId).IfMatch(ifMatch).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    attachmentId := "attachmentId_example" // string | key: id of attachment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsDeleteAttachments(context.Background(), groupId, conversationThreadId, postId, attachmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsDeleteAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsDeleteAttachmentsRequest struct via the builder pattern


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


## GroupsThreadsPostsDeleteExtensions

> GroupsThreadsPostsDeleteExtensions(ctx, groupId, conversationThreadId, postId, extensionId).IfMatch(ifMatch).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsDeleteExtensions(context.Background(), groupId, conversationThreadId, postId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsDeleteExtensionsRequest struct via the builder pattern


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


## GroupsThreadsPostsDeleteInReplyTo

> GroupsThreadsPostsDeleteInReplyTo(ctx, groupId, conversationThreadId, postId).IfMatch(ifMatch).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsDeleteInReplyTo(context.Background(), groupId, conversationThreadId, postId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsDeleteInReplyTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsDeleteInReplyToRequest struct via the builder pattern


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


## GroupsThreadsPostsDeleteMultiValueExtendedProperties

> GroupsThreadsPostsDeleteMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsDeleteMultiValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsDeleteSingleValueExtendedProperties

> GroupsThreadsPostsDeleteSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsDeleteSingleValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsGetAttachments

> MicrosoftGraphAttachment GroupsThreadsPostsGetAttachments(ctx, groupId, conversationThreadId, postId, attachmentId).Select_(select_).Expand(expand).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    attachmentId := "attachmentId_example" // string | key: id of attachment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsGetAttachments(context.Background(), groupId, conversationThreadId, postId, attachmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsGetAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsGetAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsGetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsGetAttachmentsRequest struct via the builder pattern


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


## GroupsThreadsPostsGetExtensions

> MicrosoftGraphExtension GroupsThreadsPostsGetExtensions(ctx, groupId, conversationThreadId, postId, extensionId).Select_(select_).Expand(expand).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsGetExtensions(context.Background(), groupId, conversationThreadId, postId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsGetExtensionsRequest struct via the builder pattern


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


## GroupsThreadsPostsGetInReplyTo

> MicrosoftGraphPost GroupsThreadsPostsGetInReplyTo(ctx, groupId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsGetInReplyTo(context.Background(), groupId, conversationThreadId, postId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsGetInReplyTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsGetInReplyTo`: MicrosoftGraphPost
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsGetInReplyTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsGetInReplyToRequest struct via the builder pattern


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


## GroupsThreadsPostsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty GroupsThreadsPostsGetMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsGetMultiValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty GroupsThreadsPostsGetSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsGetSingleValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsListAttachments

> CollectionOfAttachment GroupsThreadsPostsListAttachments(ctx, groupId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsListAttachments(context.Background(), groupId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsListAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsListAttachments`: CollectionOfAttachment
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsListAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsListAttachmentsRequest struct via the builder pattern


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


## GroupsThreadsPostsListExtensions

> CollectionOfExtension GroupsThreadsPostsListExtensions(ctx, groupId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsListExtensions(context.Background(), groupId, conversationThreadId, postId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsListExtensionsRequest struct via the builder pattern


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


## GroupsThreadsPostsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty GroupsThreadsPostsListMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsListMultiValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty GroupsThreadsPostsListSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsListSingleValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsThreadsPostsListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `GroupsConversationThreadApi.GroupsThreadsPostsListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsUpdateAttachments

> GroupsThreadsPostsUpdateAttachments(ctx, groupId, conversationThreadId, postId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    attachmentId := "attachmentId_example" // string | key: id of attachment
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsUpdateAttachments(context.Background(), groupId, conversationThreadId, postId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsUpdateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsUpdateAttachmentsRequest struct via the builder pattern


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


## GroupsThreadsPostsUpdateExtensions

> GroupsThreadsPostsUpdateExtensions(ctx, groupId, conversationThreadId, postId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsUpdateExtensions(context.Background(), groupId, conversationThreadId, postId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsUpdateExtensionsRequest struct via the builder pattern


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


## GroupsThreadsPostsUpdateInReplyTo

> GroupsThreadsPostsUpdateInReplyTo(ctx, groupId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphPost := *openapiclient.NewMicrosoftGraphPost() // MicrosoftGraphPost | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsUpdateInReplyTo(context.Background(), groupId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsUpdateInReplyTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsUpdateInReplyToRequest struct via the builder pattern


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


## GroupsThreadsPostsUpdateMultiValueExtendedProperties

> GroupsThreadsPostsUpdateMultiValueExtendedProperties(ctx, groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsUpdateMultiValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsPostsUpdateSingleValueExtendedProperties

> GroupsThreadsPostsUpdateSingleValueExtendedProperties(ctx, groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsPostsUpdateSingleValueExtendedProperties(context.Background(), groupId, conversationThreadId, postId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsPostsUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsPostsUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## GroupsThreadsUpdatePosts

> GroupsThreadsUpdatePosts(ctx, groupId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    postId := "postId_example" // string | key: id of post
    microsoftGraphPost := *openapiclient.NewMicrosoftGraphPost() // MicrosoftGraphPost | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsThreadsUpdatePosts(context.Background(), groupId, conversationThreadId, postId).MicrosoftGraphPost(microsoftGraphPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsThreadsUpdatePosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 
**postId** | **string** | key: id of post | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsThreadsUpdatePostsRequest struct via the builder pattern


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


## GroupsUpdateThreads

> GroupsUpdateThreads(ctx, groupId, conversationThreadId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()

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
    conversationThreadId := "conversationThreadId_example" // string | key: id of conversationThread
    microsoftGraphConversationThread := *openapiclient.NewMicrosoftGraphConversationThread() // MicrosoftGraphConversationThread | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsConversationThreadApi.GroupsUpdateThreads(context.Background(), groupId, conversationThreadId).MicrosoftGraphConversationThread(microsoftGraphConversationThread).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsConversationThreadApi.GroupsUpdateThreads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**conversationThreadId** | **string** | key: id of conversationThread | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateThreadsRequest struct via the builder pattern


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

