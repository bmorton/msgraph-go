# \UsersMessageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateMessages**](UsersMessageApi.md#UsersCreateMessages) | **Post** /users/{user-id}/messages | Create new navigation property to messages for users
[**UsersDeleteMessages**](UsersMessageApi.md#UsersDeleteMessages) | **Delete** /users/{user-id}/messages/{message-id} | Delete navigation property messages for users
[**UsersGetMessages**](UsersMessageApi.md#UsersGetMessages) | **Get** /users/{user-id}/messages/{message-id} | Get messages from users
[**UsersGetMessagesContent**](UsersMessageApi.md#UsersGetMessagesContent) | **Get** /users/{user-id}/messages/{message-id}/$value | Get media content for the navigation property messages from users
[**UsersListMessages**](UsersMessageApi.md#UsersListMessages) | **Get** /users/{user-id}/messages | Get messages from users
[**UsersMessagesCreateAttachments**](UsersMessageApi.md#UsersMessagesCreateAttachments) | **Post** /users/{user-id}/messages/{message-id}/attachments | Create new navigation property to attachments for users
[**UsersMessagesCreateExtensions**](UsersMessageApi.md#UsersMessagesCreateExtensions) | **Post** /users/{user-id}/messages/{message-id}/extensions | Create new navigation property to extensions for users
[**UsersMessagesCreateMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersMessagesCreateSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersMessagesDeleteAttachments**](UsersMessageApi.md#UsersMessagesDeleteAttachments) | **Delete** /users/{user-id}/messages/{message-id}/attachments/{attachment-id} | Delete navigation property attachments for users
[**UsersMessagesDeleteExtensions**](UsersMessageApi.md#UsersMessagesDeleteExtensions) | **Delete** /users/{user-id}/messages/{message-id}/extensions/{extension-id} | Delete navigation property extensions for users
[**UsersMessagesDeleteMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesDeleteMultiValueExtendedProperties) | **Delete** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for users
[**UsersMessagesDeleteSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesDeleteSingleValueExtendedProperties) | **Delete** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for users
[**UsersMessagesGetAttachments**](UsersMessageApi.md#UsersMessagesGetAttachments) | **Get** /users/{user-id}/messages/{message-id}/attachments/{attachment-id} | Get attachments from users
[**UsersMessagesGetExtensions**](UsersMessageApi.md#UsersMessagesGetExtensions) | **Get** /users/{user-id}/messages/{message-id}/extensions/{extension-id} | Get extensions from users
[**UsersMessagesGetMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesGetMultiValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersMessagesGetSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesGetSingleValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersMessagesListAttachments**](UsersMessageApi.md#UsersMessagesListAttachments) | **Get** /users/{user-id}/messages/{message-id}/attachments | Get attachments from users
[**UsersMessagesListExtensions**](UsersMessageApi.md#UsersMessagesListExtensions) | **Get** /users/{user-id}/messages/{message-id}/extensions | Get extensions from users
[**UsersMessagesListMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesListMultiValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersMessagesListSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesListSingleValueExtendedProperties) | **Get** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersMessagesUpdateAttachments**](UsersMessageApi.md#UsersMessagesUpdateAttachments) | **Patch** /users/{user-id}/messages/{message-id}/attachments/{attachment-id} | Update the navigation property attachments in users
[**UsersMessagesUpdateExtensions**](UsersMessageApi.md#UsersMessagesUpdateExtensions) | **Patch** /users/{user-id}/messages/{message-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersMessagesUpdateMultiValueExtendedProperties**](UsersMessageApi.md#UsersMessagesUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersMessagesUpdateSingleValueExtendedProperties**](UsersMessageApi.md#UsersMessagesUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersUpdateMessages**](UsersMessageApi.md#UsersUpdateMessages) | **Patch** /users/{user-id}/messages/{message-id} | Update the navigation property messages in users
[**UsersUpdateMessagesContent**](UsersMessageApi.md#UsersUpdateMessagesContent) | **Put** /users/{user-id}/messages/{message-id}/$value | Update media content for the navigation property messages in users



## UsersCreateMessages

> MicrosoftGraphMessage UsersCreateMessages(ctx, userId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

Create new navigation property to messages for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersCreateMessages(context.Background(), userId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersCreateMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMessage** | [**MicrosoftGraphMessage**](MicrosoftGraphMessage.md) | New navigation property | 

### Return type

[**MicrosoftGraphMessage**](MicrosoftGraphMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteMessages

> UsersDeleteMessages(ctx, userId, messageId).IfMatch(ifMatch).Execute()

Delete navigation property messages for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersDeleteMessages(context.Background(), userId, messageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersDeleteMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteMessagesRequest struct via the builder pattern


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


## UsersGetMessages

> MicrosoftGraphMessage UsersGetMessages(ctx, userId, messageId).Select_(select_).Execute()

Get messages from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersGetMessages(context.Background(), userId, messageId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphMessage**](MicrosoftGraphMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetMessagesContent

> *os.File UsersGetMessagesContent(ctx, userId, messageId).Execute()

Get media content for the navigation property messages from users

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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersGetMessagesContent(context.Background(), userId, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersGetMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetMessagesContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersGetMessagesContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetMessagesContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListMessages

> CollectionOfMessage UsersListMessages(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get messages from users



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
    userId := "userId_example" // string | key: id of user
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersListMessages(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListMessages`: CollectionOfMessage
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListMessagesRequest struct via the builder pattern


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

[**CollectionOfMessage**](CollectionOfMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersMessagesCreateAttachments

> MicrosoftGraphAttachment UsersMessagesCreateAttachments(ctx, userId, messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

Create new navigation property to attachments for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesCreateAttachments(context.Background(), userId, messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesCreateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesCreateAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesCreateAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesCreateAttachmentsRequest struct via the builder pattern


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


## UsersMessagesCreateExtensions

> MicrosoftGraphExtension UsersMessagesCreateExtensions(ctx, userId, messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesCreateExtensions(context.Background(), userId, messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesCreateExtensionsRequest struct via the builder pattern


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


## UsersMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMessagesCreateMultiValueExtendedProperties(ctx, userId, messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

Create new navigation property to multiValueExtendedProperties for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesCreateMultiValueExtendedProperties(context.Background(), userId, messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMessagesCreateSingleValueExtendedProperties(ctx, userId, messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

Create new navigation property to singleValueExtendedProperties for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesCreateSingleValueExtendedProperties(context.Background(), userId, messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesDeleteAttachments

> UsersMessagesDeleteAttachments(ctx, userId, messageId, attachmentId).IfMatch(ifMatch).Execute()

Delete navigation property attachments for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesDeleteAttachments(context.Background(), userId, messageId, attachmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesDeleteAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesDeleteAttachmentsRequest struct via the builder pattern


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


## UsersMessagesDeleteExtensions

> UsersMessagesDeleteExtensions(ctx, userId, messageId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesDeleteExtensions(context.Background(), userId, messageId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesDeleteExtensionsRequest struct via the builder pattern


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


## UsersMessagesDeleteMultiValueExtendedProperties

> UsersMessagesDeleteMultiValueExtendedProperties(ctx, userId, messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

Delete navigation property multiValueExtendedProperties for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesDeleteMultiValueExtendedProperties(context.Background(), userId, messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesDeleteSingleValueExtendedProperties

> UsersMessagesDeleteSingleValueExtendedProperties(ctx, userId, messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

Delete navigation property singleValueExtendedProperties for users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesDeleteSingleValueExtendedProperties(context.Background(), userId, messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesGetAttachments

> MicrosoftGraphAttachment UsersMessagesGetAttachments(ctx, userId, messageId, attachmentId).Select_(select_).Expand(expand).Execute()

Get attachments from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesGetAttachments(context.Background(), userId, messageId, attachmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesGetAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesGetAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesGetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesGetAttachmentsRequest struct via the builder pattern


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


## UsersMessagesGetExtensions

> MicrosoftGraphExtension UsersMessagesGetExtensions(ctx, userId, messageId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesGetExtensions(context.Background(), userId, messageId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesGetExtensionsRequest struct via the builder pattern


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


## UsersMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMessagesGetMultiValueExtendedProperties(ctx, userId, messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

Get multiValueExtendedProperties from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesGetMultiValueExtendedProperties(context.Background(), userId, messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMessagesGetSingleValueExtendedProperties(ctx, userId, messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

Get singleValueExtendedProperties from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesGetSingleValueExtendedProperties(context.Background(), userId, messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesListAttachments

> CollectionOfAttachment UsersMessagesListAttachments(ctx, userId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get attachments from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesListAttachments(context.Background(), userId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesListAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesListAttachments`: CollectionOfAttachment
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesListAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesListAttachmentsRequest struct via the builder pattern


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


## UsersMessagesListExtensions

> CollectionOfExtension UsersMessagesListExtensions(ctx, userId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesListExtensions(context.Background(), userId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesListExtensionsRequest struct via the builder pattern


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


## UsersMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersMessagesListMultiValueExtendedProperties(ctx, userId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get multiValueExtendedProperties from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
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
    resp, r, err := api_client.UsersMessageApi.UsersMessagesListMultiValueExtendedProperties(context.Background(), userId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersMessagesListSingleValueExtendedProperties(ctx, userId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get singleValueExtendedProperties from users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
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
    resp, r, err := api_client.UsersMessageApi.UsersMessagesListSingleValueExtendedProperties(context.Background(), userId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMessagesListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMessageApi.UsersMessagesListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesUpdateAttachments

> UsersMessagesUpdateAttachments(ctx, userId, messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

Update the navigation property attachments in users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesUpdateAttachments(context.Background(), userId, messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesUpdateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesUpdateAttachmentsRequest struct via the builder pattern


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


## UsersMessagesUpdateExtensions

> UsersMessagesUpdateExtensions(ctx, userId, messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesUpdateExtensions(context.Background(), userId, messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesUpdateExtensionsRequest struct via the builder pattern


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


## UsersMessagesUpdateMultiValueExtendedProperties

> UsersMessagesUpdateMultiValueExtendedProperties(ctx, userId, messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

Update the navigation property multiValueExtendedProperties in users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesUpdateMultiValueExtendedProperties(context.Background(), userId, messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMessagesUpdateSingleValueExtendedProperties

> UsersMessagesUpdateSingleValueExtendedProperties(ctx, userId, messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

Update the navigation property singleValueExtendedProperties in users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersMessagesUpdateSingleValueExtendedProperties(context.Background(), userId, messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersMessagesUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMessagesUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersUpdateMessages

> UsersUpdateMessages(ctx, userId, messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

Update the navigation property messages in users



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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersUpdateMessages(context.Background(), userId, messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersUpdateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphMessage** | [**MicrosoftGraphMessage**](MicrosoftGraphMessage.md) | New navigation property values | 

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


## UsersUpdateMessagesContent

> UsersUpdateMessagesContent(ctx, userId, messageId).Body(body).Execute()

Update media content for the navigation property messages in users

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
    userId := "userId_example" // string | key: id of user
    messageId := "messageId_example" // string | key: id of message
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMessageApi.UsersUpdateMessagesContent(context.Background(), userId, messageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMessageApi.UsersUpdateMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateMessagesContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | New media content. | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

