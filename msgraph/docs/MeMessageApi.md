# \MeMessageApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateMessages**](MeMessageApi.md#MeCreateMessages) | **Post** /me/messages | Create new navigation property to messages for me
[**MeDeleteMessages**](MeMessageApi.md#MeDeleteMessages) | **Delete** /me/messages/{message-id} | Delete navigation property messages for me
[**MeGetMessages**](MeMessageApi.md#MeGetMessages) | **Get** /me/messages/{message-id} | Get messages from me
[**MeGetMessagesContent**](MeMessageApi.md#MeGetMessagesContent) | **Get** /me/messages/{message-id}/$value | Get media content for the navigation property messages from me
[**MeListMessages**](MeMessageApi.md#MeListMessages) | **Get** /me/messages | Get messages from me
[**MeMessagesCreateAttachments**](MeMessageApi.md#MeMessagesCreateAttachments) | **Post** /me/messages/{message-id}/attachments | Create new navigation property to attachments for me
[**MeMessagesCreateExtensions**](MeMessageApi.md#MeMessagesCreateExtensions) | **Post** /me/messages/{message-id}/extensions | Create new navigation property to extensions for me
[**MeMessagesCreateMultiValueExtendedProperties**](MeMessageApi.md#MeMessagesCreateMultiValueExtendedProperties) | **Post** /me/messages/{message-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeMessagesCreateSingleValueExtendedProperties**](MeMessageApi.md#MeMessagesCreateSingleValueExtendedProperties) | **Post** /me/messages/{message-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for me
[**MeMessagesDeleteAttachments**](MeMessageApi.md#MeMessagesDeleteAttachments) | **Delete** /me/messages/{message-id}/attachments/{attachment-id} | Delete navigation property attachments for me
[**MeMessagesDeleteExtensions**](MeMessageApi.md#MeMessagesDeleteExtensions) | **Delete** /me/messages/{message-id}/extensions/{extension-id} | Delete navigation property extensions for me
[**MeMessagesDeleteMultiValueExtendedProperties**](MeMessageApi.md#MeMessagesDeleteMultiValueExtendedProperties) | **Delete** /me/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for me
[**MeMessagesDeleteSingleValueExtendedProperties**](MeMessageApi.md#MeMessagesDeleteSingleValueExtendedProperties) | **Delete** /me/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for me
[**MeMessagesGetAttachments**](MeMessageApi.md#MeMessagesGetAttachments) | **Get** /me/messages/{message-id}/attachments/{attachment-id} | Get attachments from me
[**MeMessagesGetExtensions**](MeMessageApi.md#MeMessagesGetExtensions) | **Get** /me/messages/{message-id}/extensions/{extension-id} | Get extensions from me
[**MeMessagesGetMultiValueExtendedProperties**](MeMessageApi.md#MeMessagesGetMultiValueExtendedProperties) | **Get** /me/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from me
[**MeMessagesGetSingleValueExtendedProperties**](MeMessageApi.md#MeMessagesGetSingleValueExtendedProperties) | **Get** /me/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from me
[**MeMessagesListAttachments**](MeMessageApi.md#MeMessagesListAttachments) | **Get** /me/messages/{message-id}/attachments | Get attachments from me
[**MeMessagesListExtensions**](MeMessageApi.md#MeMessagesListExtensions) | **Get** /me/messages/{message-id}/extensions | Get extensions from me
[**MeMessagesListMultiValueExtendedProperties**](MeMessageApi.md#MeMessagesListMultiValueExtendedProperties) | **Get** /me/messages/{message-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeMessagesListSingleValueExtendedProperties**](MeMessageApi.md#MeMessagesListSingleValueExtendedProperties) | **Get** /me/messages/{message-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from me
[**MeMessagesUpdateAttachments**](MeMessageApi.md#MeMessagesUpdateAttachments) | **Patch** /me/messages/{message-id}/attachments/{attachment-id} | Update the navigation property attachments in me
[**MeMessagesUpdateExtensions**](MeMessageApi.md#MeMessagesUpdateExtensions) | **Patch** /me/messages/{message-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeMessagesUpdateMultiValueExtendedProperties**](MeMessageApi.md#MeMessagesUpdateMultiValueExtendedProperties) | **Patch** /me/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in me
[**MeMessagesUpdateSingleValueExtendedProperties**](MeMessageApi.md#MeMessagesUpdateSingleValueExtendedProperties) | **Patch** /me/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in me
[**MeUpdateMessages**](MeMessageApi.md#MeUpdateMessages) | **Patch** /me/messages/{message-id} | Update the navigation property messages in me
[**MeUpdateMessagesContent**](MeMessageApi.md#MeUpdateMessagesContent) | **Put** /me/messages/{message-id}/$value | Update media content for the navigation property messages in me



## MeCreateMessages

> MicrosoftGraphMessage MeCreateMessages(ctx).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

Create new navigation property to messages for me



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
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeCreateMessages(context.Background()).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeCreateMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateMessagesRequest struct via the builder pattern


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


## MeDeleteMessages

> MeDeleteMessages(ctx, messageId).IfMatch(ifMatch).Execute()

Delete navigation property messages for me



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
    messageId := "messageId_example" // string | key: id of message
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeDeleteMessages(context.Background(), messageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeDeleteMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteMessagesRequest struct via the builder pattern


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


## MeGetMessages

> MicrosoftGraphMessage MeGetMessages(ctx, messageId).Select_(select_).Execute()

Get messages from me



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
    messageId := "messageId_example" // string | key: id of message
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeGetMessages(context.Background(), messageId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetMessagesRequest struct via the builder pattern


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


## MeGetMessagesContent

> *os.File MeGetMessagesContent(ctx, messageId).Execute()

Get media content for the navigation property messages from me

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
    messageId := "messageId_example" // string | key: id of message

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeGetMessagesContent(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeGetMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetMessagesContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeGetMessagesContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetMessagesContentRequest struct via the builder pattern


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


## MeListMessages

> CollectionOfMessage MeListMessages(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get messages from me



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
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeListMessages(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListMessages`: CollectionOfMessage
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeListMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListMessagesRequest struct via the builder pattern


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


## MeMessagesCreateAttachments

> MicrosoftGraphAttachment MeMessagesCreateAttachments(ctx, messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

Create new navigation property to attachments for me



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
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesCreateAttachments(context.Background(), messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesCreateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesCreateAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesCreateAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesCreateAttachmentsRequest struct via the builder pattern


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


## MeMessagesCreateExtensions

> MicrosoftGraphExtension MeMessagesCreateExtensions(ctx, messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for me



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
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesCreateExtensions(context.Background(), messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesCreateExtensionsRequest struct via the builder pattern


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


## MeMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMessagesCreateMultiValueExtendedProperties(ctx, messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

Create new navigation property to multiValueExtendedProperties for me



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
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesCreateMultiValueExtendedProperties(context.Background(), messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMessagesCreateSingleValueExtendedProperties(ctx, messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

Create new navigation property to singleValueExtendedProperties for me



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
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesCreateSingleValueExtendedProperties(context.Background(), messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesDeleteAttachments

> MeMessagesDeleteAttachments(ctx, messageId, attachmentId).IfMatch(ifMatch).Execute()

Delete navigation property attachments for me



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
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesDeleteAttachments(context.Background(), messageId, attachmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesDeleteAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesDeleteAttachmentsRequest struct via the builder pattern


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


## MeMessagesDeleteExtensions

> MeMessagesDeleteExtensions(ctx, messageId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for me



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
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesDeleteExtensions(context.Background(), messageId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesDeleteExtensionsRequest struct via the builder pattern


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


## MeMessagesDeleteMultiValueExtendedProperties

> MeMessagesDeleteMultiValueExtendedProperties(ctx, messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

Delete navigation property multiValueExtendedProperties for me



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
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesDeleteMultiValueExtendedProperties(context.Background(), messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesDeleteSingleValueExtendedProperties

> MeMessagesDeleteSingleValueExtendedProperties(ctx, messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

Delete navigation property singleValueExtendedProperties for me



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
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesDeleteSingleValueExtendedProperties(context.Background(), messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesGetAttachments

> MicrosoftGraphAttachment MeMessagesGetAttachments(ctx, messageId, attachmentId).Select_(select_).Expand(expand).Execute()

Get attachments from me



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
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesGetAttachments(context.Background(), messageId, attachmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesGetAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesGetAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesGetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesGetAttachmentsRequest struct via the builder pattern


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


## MeMessagesGetExtensions

> MicrosoftGraphExtension MeMessagesGetExtensions(ctx, messageId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesGetExtensions(context.Background(), messageId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesGetExtensionsRequest struct via the builder pattern


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


## MeMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMessagesGetMultiValueExtendedProperties(ctx, messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

Get multiValueExtendedProperties from me



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
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesGetMultiValueExtendedProperties(context.Background(), messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMessagesGetSingleValueExtendedProperties(ctx, messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

Get singleValueExtendedProperties from me



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
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesGetSingleValueExtendedProperties(context.Background(), messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesListAttachments

> CollectionOfAttachment MeMessagesListAttachments(ctx, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get attachments from me



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
    resp, r, err := api_client.MeMessageApi.MeMessagesListAttachments(context.Background(), messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesListAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesListAttachments`: CollectionOfAttachment
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesListAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesListAttachmentsRequest struct via the builder pattern


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


## MeMessagesListExtensions

> CollectionOfExtension MeMessagesListExtensions(ctx, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    resp, r, err := api_client.MeMessageApi.MeMessagesListExtensions(context.Background(), messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesListExtensionsRequest struct via the builder pattern


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


## MeMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeMessagesListMultiValueExtendedProperties(ctx, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get multiValueExtendedProperties from me



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
    resp, r, err := api_client.MeMessageApi.MeMessagesListMultiValueExtendedProperties(context.Background(), messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty MeMessagesListSingleValueExtendedProperties(ctx, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get singleValueExtendedProperties from me



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
    resp, r, err := api_client.MeMessageApi.MeMessagesListSingleValueExtendedProperties(context.Background(), messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMessagesListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMessageApi.MeMessagesListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesUpdateAttachments

> MeMessagesUpdateAttachments(ctx, messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

Update the navigation property attachments in me



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
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesUpdateAttachments(context.Background(), messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesUpdateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesUpdateAttachmentsRequest struct via the builder pattern


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


## MeMessagesUpdateExtensions

> MeMessagesUpdateExtensions(ctx, messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in me



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
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesUpdateExtensions(context.Background(), messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesUpdateExtensionsRequest struct via the builder pattern


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


## MeMessagesUpdateMultiValueExtendedProperties

> MeMessagesUpdateMultiValueExtendedProperties(ctx, messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

Update the navigation property multiValueExtendedProperties in me



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
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesUpdateMultiValueExtendedProperties(context.Background(), messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMessagesUpdateSingleValueExtendedProperties

> MeMessagesUpdateSingleValueExtendedProperties(ctx, messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

Update the navigation property singleValueExtendedProperties in me



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
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeMessagesUpdateSingleValueExtendedProperties(context.Background(), messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeMessagesUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMessagesUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeUpdateMessages

> MeUpdateMessages(ctx, messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

Update the navigation property messages in me



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
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeUpdateMessages(context.Background(), messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeUpdateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateMessagesRequest struct via the builder pattern


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


## MeUpdateMessagesContent

> MeUpdateMessagesContent(ctx, messageId).Body(body).Execute()

Update media content for the navigation property messages in me

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
    messageId := "messageId_example" // string | key: id of message
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMessageApi.MeUpdateMessagesContent(context.Background(), messageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMessageApi.MeUpdateMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateMessagesContentRequest struct via the builder pattern


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

