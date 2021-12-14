# \MeMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateMailFolders**](MeMailFolderApi.md#MeCreateMailFolders) | **Post** /me/mailFolders | Create new navigation property to mailFolders for me
[**MeDeleteMailFolders**](MeMailFolderApi.md#MeDeleteMailFolders) | **Delete** /me/mailFolders/{mailFolder-id} | Delete navigation property mailFolders for me
[**MeGetMailFolders**](MeMailFolderApi.md#MeGetMailFolders) | **Get** /me/mailFolders/{mailFolder-id} | Get mailFolders from me
[**MeListMailFolders**](MeMailFolderApi.md#MeListMailFolders) | **Get** /me/mailFolders | Get mailFolders from me
[**MeMailFoldersCreateChildFolders**](MeMailFolderApi.md#MeMailFoldersCreateChildFolders) | **Post** /me/mailFolders/{mailFolder-id}/childFolders | Create new navigation property to childFolders for me
[**MeMailFoldersCreateMessageRules**](MeMailFolderApi.md#MeMailFoldersCreateMessageRules) | **Post** /me/mailFolders/{mailFolder-id}/messageRules | Create new navigation property to messageRules for me
[**MeMailFoldersCreateMessages**](MeMailFolderApi.md#MeMailFoldersCreateMessages) | **Post** /me/mailFolders/{mailFolder-id}/messages | Create new navigation property to messages for me
[**MeMailFoldersCreateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersCreateMultiValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeMailFoldersCreateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersCreateSingleValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for me
[**MeMailFoldersDeleteChildFolders**](MeMailFolderApi.md#MeMailFoldersDeleteChildFolders) | **Delete** /me/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Delete navigation property childFolders for me
[**MeMailFoldersDeleteMessageRules**](MeMailFolderApi.md#MeMailFoldersDeleteMessageRules) | **Delete** /me/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Delete navigation property messageRules for me
[**MeMailFoldersDeleteMessages**](MeMailFolderApi.md#MeMailFoldersDeleteMessages) | **Delete** /me/mailFolders/{mailFolder-id}/messages/{message-id} | Delete navigation property messages for me
[**MeMailFoldersDeleteMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersDeleteMultiValueExtendedProperties) | **Delete** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for me
[**MeMailFoldersDeleteSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersDeleteSingleValueExtendedProperties) | **Delete** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for me
[**MeMailFoldersGetChildFolders**](MeMailFolderApi.md#MeMailFoldersGetChildFolders) | **Get** /me/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Get childFolders from me
[**MeMailFoldersGetMessageRules**](MeMailFolderApi.md#MeMailFoldersGetMessageRules) | **Get** /me/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Get messageRules from me
[**MeMailFoldersGetMessages**](MeMailFolderApi.md#MeMailFoldersGetMessages) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id} | Get messages from me
[**MeMailFoldersGetMessagesContent**](MeMailFolderApi.md#MeMailFoldersGetMessagesContent) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/$value | Get media content for the navigation property messages from me
[**MeMailFoldersGetMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersGetMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from me
[**MeMailFoldersGetSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersGetSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from me
[**MeMailFoldersListChildFolders**](MeMailFolderApi.md#MeMailFoldersListChildFolders) | **Get** /me/mailFolders/{mailFolder-id}/childFolders | Get childFolders from me
[**MeMailFoldersListMessageRules**](MeMailFolderApi.md#MeMailFoldersListMessageRules) | **Get** /me/mailFolders/{mailFolder-id}/messageRules | Get messageRules from me
[**MeMailFoldersListMessages**](MeMailFolderApi.md#MeMailFoldersListMessages) | **Get** /me/mailFolders/{mailFolder-id}/messages | Get messages from me
[**MeMailFoldersListMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersListMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeMailFoldersListSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersListSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from me
[**MeMailFoldersMessagesCreateAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesCreateAttachments) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Create new navigation property to attachments for me
[**MeMailFoldersMessagesCreateExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesCreateExtensions) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Create new navigation property to extensions for me
[**MeMailFoldersMessagesCreateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesCreateMultiValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeMailFoldersMessagesCreateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesCreateSingleValueExtendedProperties) | **Post** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for me
[**MeMailFoldersMessagesDeleteAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesDeleteAttachments) | **Delete** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Delete navigation property attachments for me
[**MeMailFoldersMessagesDeleteExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesDeleteExtensions) | **Delete** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Delete navigation property extensions for me
[**MeMailFoldersMessagesDeleteMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesDeleteMultiValueExtendedProperties) | **Delete** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for me
[**MeMailFoldersMessagesDeleteSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesDeleteSingleValueExtendedProperties) | **Delete** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for me
[**MeMailFoldersMessagesGetAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesGetAttachments) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Get attachments from me
[**MeMailFoldersMessagesGetExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesGetExtensions) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Get extensions from me
[**MeMailFoldersMessagesGetMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesGetMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from me
[**MeMailFoldersMessagesGetSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesGetSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from me
[**MeMailFoldersMessagesListAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesListAttachments) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Get attachments from me
[**MeMailFoldersMessagesListExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesListExtensions) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Get extensions from me
[**MeMailFoldersMessagesListMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesListMultiValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeMailFoldersMessagesListSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesListSingleValueExtendedProperties) | **Get** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from me
[**MeMailFoldersMessagesUpdateAttachments**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateAttachments) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Update the navigation property attachments in me
[**MeMailFoldersMessagesUpdateExtensions**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateExtensions) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeMailFoldersMessagesUpdateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateMultiValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in me
[**MeMailFoldersMessagesUpdateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersMessagesUpdateSingleValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in me
[**MeMailFoldersUpdateChildFolders**](MeMailFolderApi.md#MeMailFoldersUpdateChildFolders) | **Patch** /me/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Update the navigation property childFolders in me
[**MeMailFoldersUpdateMessageRules**](MeMailFolderApi.md#MeMailFoldersUpdateMessageRules) | **Patch** /me/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Update the navigation property messageRules in me
[**MeMailFoldersUpdateMessages**](MeMailFolderApi.md#MeMailFoldersUpdateMessages) | **Patch** /me/mailFolders/{mailFolder-id}/messages/{message-id} | Update the navigation property messages in me
[**MeMailFoldersUpdateMessagesContent**](MeMailFolderApi.md#MeMailFoldersUpdateMessagesContent) | **Put** /me/mailFolders/{mailFolder-id}/messages/{message-id}/$value | Update media content for the navigation property messages in me
[**MeMailFoldersUpdateMultiValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersUpdateMultiValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in me
[**MeMailFoldersUpdateSingleValueExtendedProperties**](MeMailFolderApi.md#MeMailFoldersUpdateSingleValueExtendedProperties) | **Patch** /me/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in me
[**MeUpdateMailFolders**](MeMailFolderApi.md#MeUpdateMailFolders) | **Patch** /me/mailFolders/{mailFolder-id} | Update the navigation property mailFolders in me



## MeCreateMailFolders

> MicrosoftGraphMailFolder MeCreateMailFolders(ctx).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Create new navigation property to mailFolders for me



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
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeCreateMailFolders(context.Background()).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeCreateMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateMailFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeCreateMailFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateMailFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md) | New navigation property | 

### Return type

[**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDeleteMailFolders

> MeDeleteMailFolders(ctx, mailFolderId).IfMatch(ifMatch).Execute()

Delete navigation property mailFolders for me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeDeleteMailFolders(context.Background(), mailFolderId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeDeleteMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteMailFoldersRequest struct via the builder pattern


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


## MeGetMailFolders

> MicrosoftGraphMailFolder MeGetMailFolders(ctx, mailFolderId).Select_(select_).Execute()

Get mailFolders from me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeGetMailFolders(context.Background(), mailFolderId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeGetMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetMailFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeGetMailFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetMailFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListMailFolders

> CollectionOfMailFolder MeListMailFolders(ctx).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get mailFolders from me



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
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeListMailFolders(context.Background()).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeListMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListMailFolders`: CollectionOfMailFolder
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeListMailFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListMailFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfMailFolder**](CollectionOfMailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersCreateChildFolders

> MicrosoftGraphMailFolder MeMailFoldersCreateChildFolders(ctx, mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Create new navigation property to childFolders for me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersCreateChildFolders(context.Background(), mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersCreateChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersCreateChildFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersCreateChildFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersCreateChildFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md) | New navigation property | 

### Return type

[**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersCreateMessageRules

> MicrosoftGraphMessageRule MeMailFoldersCreateMessageRules(ctx, mailFolderId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()

Create new navigation property to messageRules for me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMessageRule := *openapiclient.NewMicrosoftGraphMessageRule() // MicrosoftGraphMessageRule | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersCreateMessageRules(context.Background(), mailFolderId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersCreateMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersCreateMessageRules`: MicrosoftGraphMessageRule
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersCreateMessageRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersCreateMessageRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMessageRule** | [**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md) | New navigation property | 

### Return type

[**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersCreateMessages

> MicrosoftGraphMessage MeMailFoldersCreateMessages(ctx, mailFolderId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersCreateMessages(context.Background(), mailFolderId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersCreateMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersCreateMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersCreateMessagesRequest struct via the builder pattern


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


## MeMailFoldersCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersCreateMultiValueExtendedProperties(ctx, mailFolderId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersCreateMultiValueExtendedProperties(context.Background(), mailFolderId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersCreateSingleValueExtendedProperties(ctx, mailFolderId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersCreateSingleValueExtendedProperties(context.Background(), mailFolderId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersDeleteChildFolders

> MeMailFoldersDeleteChildFolders(ctx, mailFolderId, mailFolderId1).IfMatch(ifMatch).Execute()

Delete navigation property childFolders for me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    mailFolderId1 := "mailFolderId1_example" // string | key: id of mailFolder
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersDeleteChildFolders(context.Background(), mailFolderId, mailFolderId1).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersDeleteChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**mailFolderId1** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersDeleteChildFoldersRequest struct via the builder pattern


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


## MeMailFoldersDeleteMessageRules

> MeMailFoldersDeleteMessageRules(ctx, mailFolderId, messageRuleId).IfMatch(ifMatch).Execute()

Delete navigation property messageRules for me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageRuleId := "messageRuleId_example" // string | key: id of messageRule
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersDeleteMessageRules(context.Background(), mailFolderId, messageRuleId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersDeleteMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageRuleId** | **string** | key: id of messageRule | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersDeleteMessageRulesRequest struct via the builder pattern


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


## MeMailFoldersDeleteMessages

> MeMailFoldersDeleteMessages(ctx, mailFolderId, messageId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersDeleteMessages(context.Background(), mailFolderId, messageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersDeleteMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersDeleteMessagesRequest struct via the builder pattern


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


## MeMailFoldersDeleteMultiValueExtendedProperties

> MeMailFoldersDeleteMultiValueExtendedProperties(ctx, mailFolderId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersDeleteMultiValueExtendedProperties(context.Background(), mailFolderId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersDeleteSingleValueExtendedProperties

> MeMailFoldersDeleteSingleValueExtendedProperties(ctx, mailFolderId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersDeleteSingleValueExtendedProperties(context.Background(), mailFolderId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersGetChildFolders

> MicrosoftGraphMailFolder MeMailFoldersGetChildFolders(ctx, mailFolderId, mailFolderId1).Select_(select_).Expand(expand).Execute()

Get childFolders from me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    mailFolderId1 := "mailFolderId1_example" // string | key: id of mailFolder
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersGetChildFolders(context.Background(), mailFolderId, mailFolderId1).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersGetChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersGetChildFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersGetChildFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**mailFolderId1** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersGetChildFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersGetMessageRules

> MicrosoftGraphMessageRule MeMailFoldersGetMessageRules(ctx, mailFolderId, messageRuleId).Select_(select_).Execute()

Get messageRules from me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageRuleId := "messageRuleId_example" // string | key: id of messageRule
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersGetMessageRules(context.Background(), mailFolderId, messageRuleId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersGetMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersGetMessageRules`: MicrosoftGraphMessageRule
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersGetMessageRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageRuleId** | **string** | key: id of messageRule | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersGetMessageRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersGetMessages

> MicrosoftGraphMessage MeMailFoldersGetMessages(ctx, mailFolderId, messageId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersGetMessages(context.Background(), mailFolderId, messageId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersGetMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

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


## MeMailFoldersGetMessagesContent

> *os.File MeMailFoldersGetMessagesContent(ctx, mailFolderId, messageId).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersGetMessagesContent(context.Background(), mailFolderId, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersGetMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersGetMessagesContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersGetMessagesContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersGetMessagesContentRequest struct via the builder pattern


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


## MeMailFoldersGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersGetMultiValueExtendedProperties(ctx, mailFolderId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersGetMultiValueExtendedProperties(context.Background(), mailFolderId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersGetSingleValueExtendedProperties(ctx, mailFolderId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersGetSingleValueExtendedProperties(context.Background(), mailFolderId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersListChildFolders

> CollectionOfMailFolder MeMailFoldersListChildFolders(ctx, mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get childFolders from me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersListChildFolders(context.Background(), mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersListChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersListChildFolders`: CollectionOfMailFolder
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersListChildFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersListChildFoldersRequest struct via the builder pattern


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

[**CollectionOfMailFolder**](CollectionOfMailFolder.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersListMessageRules

> CollectionOfMessageRule MeMailFoldersListMessageRules(ctx, mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get messageRules from me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersListMessageRules(context.Background(), mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersListMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersListMessageRules`: CollectionOfMessageRule
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersListMessageRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersListMessageRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfMessageRule**](CollectionOfMessageRule.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersListMessages

> CollectionOfMessage MeMailFoldersListMessages(ctx, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
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
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersListMessages(context.Background(), mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersListMessages`: CollectionOfMessage
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersListMessagesRequest struct via the builder pattern


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

[**CollectionOfMessage**](CollectionOfMessage.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeMailFoldersListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeMailFoldersListMultiValueExtendedProperties(ctx, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
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
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersListMultiValueExtendedProperties(context.Background(), mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty MeMailFoldersListSingleValueExtendedProperties(ctx, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
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
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersListSingleValueExtendedProperties(context.Background(), mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesCreateAttachments

> MicrosoftGraphAttachment MeMailFoldersMessagesCreateAttachments(ctx, mailFolderId, messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesCreateAttachments(context.Background(), mailFolderId, messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesCreateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesCreateAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesCreateAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesCreateAttachmentsRequest struct via the builder pattern


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


## MeMailFoldersMessagesCreateExtensions

> MicrosoftGraphExtension MeMailFoldersMessagesCreateExtensions(ctx, mailFolderId, messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesCreateExtensions(context.Background(), mailFolderId, messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesCreateExtensionsRequest struct via the builder pattern


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


## MeMailFoldersMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersMessagesCreateMultiValueExtendedProperties(ctx, mailFolderId, messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesCreateMultiValueExtendedProperties(context.Background(), mailFolderId, messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersMessagesCreateSingleValueExtendedProperties(ctx, mailFolderId, messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesCreateSingleValueExtendedProperties(context.Background(), mailFolderId, messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesDeleteAttachments

> MeMailFoldersMessagesDeleteAttachments(ctx, mailFolderId, messageId, attachmentId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesDeleteAttachments(context.Background(), mailFolderId, messageId, attachmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesDeleteAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesDeleteAttachmentsRequest struct via the builder pattern


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


## MeMailFoldersMessagesDeleteExtensions

> MeMailFoldersMessagesDeleteExtensions(ctx, mailFolderId, messageId, extensionId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesDeleteExtensions(context.Background(), mailFolderId, messageId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesDeleteExtensionsRequest struct via the builder pattern


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


## MeMailFoldersMessagesDeleteMultiValueExtendedProperties

> MeMailFoldersMessagesDeleteMultiValueExtendedProperties(ctx, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesDeleteMultiValueExtendedProperties(context.Background(), mailFolderId, messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesDeleteSingleValueExtendedProperties

> MeMailFoldersMessagesDeleteSingleValueExtendedProperties(ctx, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesDeleteSingleValueExtendedProperties(context.Background(), mailFolderId, messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesGetAttachments

> MicrosoftGraphAttachment MeMailFoldersMessagesGetAttachments(ctx, mailFolderId, messageId, attachmentId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesGetAttachments(context.Background(), mailFolderId, messageId, attachmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesGetAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesGetAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesGetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesGetAttachmentsRequest struct via the builder pattern


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


## MeMailFoldersMessagesGetExtensions

> MicrosoftGraphExtension MeMailFoldersMessagesGetExtensions(ctx, mailFolderId, messageId, extensionId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesGetExtensions(context.Background(), mailFolderId, messageId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesGetExtensionsRequest struct via the builder pattern


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


## MeMailFoldersMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeMailFoldersMessagesGetMultiValueExtendedProperties(ctx, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesGetMultiValueExtendedProperties(context.Background(), mailFolderId, messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeMailFoldersMessagesGetSingleValueExtendedProperties(ctx, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesGetSingleValueExtendedProperties(context.Background(), mailFolderId, messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesListAttachments

> CollectionOfAttachment MeMailFoldersMessagesListAttachments(ctx, mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
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
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesListAttachments(context.Background(), mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesListAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesListAttachments`: CollectionOfAttachment
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesListAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesListAttachmentsRequest struct via the builder pattern


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


## MeMailFoldersMessagesListExtensions

> CollectionOfExtension MeMailFoldersMessagesListExtensions(ctx, mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
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
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesListExtensions(context.Background(), mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesListExtensionsRequest struct via the builder pattern


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


## MeMailFoldersMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeMailFoldersMessagesListMultiValueExtendedProperties(ctx, mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
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
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesListMultiValueExtendedProperties(context.Background(), mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty MeMailFoldersMessagesListSingleValueExtendedProperties(ctx, mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
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
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesListSingleValueExtendedProperties(context.Background(), mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeMailFoldersMessagesListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeMailFolderApi.MeMailFoldersMessagesListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesUpdateAttachments

> MeMailFoldersMessagesUpdateAttachments(ctx, mailFolderId, messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesUpdateAttachments(context.Background(), mailFolderId, messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesUpdateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesUpdateAttachmentsRequest struct via the builder pattern


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


## MeMailFoldersMessagesUpdateExtensions

> MeMailFoldersMessagesUpdateExtensions(ctx, mailFolderId, messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesUpdateExtensions(context.Background(), mailFolderId, messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesUpdateExtensionsRequest struct via the builder pattern


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


## MeMailFoldersMessagesUpdateMultiValueExtendedProperties

> MeMailFoldersMessagesUpdateMultiValueExtendedProperties(ctx, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesUpdateMultiValueExtendedProperties(context.Background(), mailFolderId, messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersMessagesUpdateSingleValueExtendedProperties

> MeMailFoldersMessagesUpdateSingleValueExtendedProperties(ctx, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersMessagesUpdateSingleValueExtendedProperties(context.Background(), mailFolderId, messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersMessagesUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersMessagesUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersUpdateChildFolders

> MeMailFoldersUpdateChildFolders(ctx, mailFolderId, mailFolderId1).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Update the navigation property childFolders in me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    mailFolderId1 := "mailFolderId1_example" // string | key: id of mailFolder
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersUpdateChildFolders(context.Background(), mailFolderId, mailFolderId1).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersUpdateChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**mailFolderId1** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersUpdateChildFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md) | New navigation property values | 

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


## MeMailFoldersUpdateMessageRules

> MeMailFoldersUpdateMessageRules(ctx, mailFolderId, messageRuleId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()

Update the navigation property messageRules in me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageRuleId := "messageRuleId_example" // string | key: id of messageRule
    microsoftGraphMessageRule := *openapiclient.NewMicrosoftGraphMessageRule() // MicrosoftGraphMessageRule | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersUpdateMessageRules(context.Background(), mailFolderId, messageRuleId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersUpdateMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageRuleId** | **string** | key: id of messageRule | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersUpdateMessageRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphMessageRule** | [**MicrosoftGraphMessageRule**](MicrosoftGraphMessageRule.md) | New navigation property values | 

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


## MeMailFoldersUpdateMessages

> MeMailFoldersUpdateMessages(ctx, mailFolderId, messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersUpdateMessages(context.Background(), mailFolderId, messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersUpdateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersUpdateMessagesRequest struct via the builder pattern


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


## MeMailFoldersUpdateMessagesContent

> MeMailFoldersUpdateMessagesContent(ctx, mailFolderId, messageId).Body(body).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersUpdateMessagesContent(context.Background(), mailFolderId, messageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersUpdateMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersUpdateMessagesContentRequest struct via the builder pattern


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


## MeMailFoldersUpdateMultiValueExtendedProperties

> MeMailFoldersUpdateMultiValueExtendedProperties(ctx, mailFolderId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersUpdateMultiValueExtendedProperties(context.Background(), mailFolderId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeMailFoldersUpdateSingleValueExtendedProperties

> MeMailFoldersUpdateSingleValueExtendedProperties(ctx, mailFolderId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeMailFoldersUpdateSingleValueExtendedProperties(context.Background(), mailFolderId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeMailFoldersUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeMailFoldersUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeUpdateMailFolders

> MeUpdateMailFolders(ctx, mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Update the navigation property mailFolders in me



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeMailFolderApi.MeUpdateMailFolders(context.Background(), mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeMailFolderApi.MeUpdateMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateMailFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphMailFolder** | [**MicrosoftGraphMailFolder**](MicrosoftGraphMailFolder.md) | New navigation property values | 

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

