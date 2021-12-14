# \UsersMailFolderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateMailFolders**](UsersMailFolderApi.md#UsersCreateMailFolders) | **Post** /users/{user-id}/mailFolders | Create new navigation property to mailFolders for users
[**UsersDeleteMailFolders**](UsersMailFolderApi.md#UsersDeleteMailFolders) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id} | Delete navigation property mailFolders for users
[**UsersGetMailFolders**](UsersMailFolderApi.md#UsersGetMailFolders) | **Get** /users/{user-id}/mailFolders/{mailFolder-id} | Get mailFolders from users
[**UsersListMailFolders**](UsersMailFolderApi.md#UsersListMailFolders) | **Get** /users/{user-id}/mailFolders | Get mailFolders from users
[**UsersMailFoldersCreateChildFolders**](UsersMailFolderApi.md#UsersMailFoldersCreateChildFolders) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders | Create new navigation property to childFolders for users
[**UsersMailFoldersCreateMessageRules**](UsersMailFolderApi.md#UsersMailFoldersCreateMessageRules) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules | Create new navigation property to messageRules for users
[**UsersMailFoldersCreateMessages**](UsersMailFolderApi.md#UsersMailFoldersCreateMessages) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages | Create new navigation property to messages for users
[**UsersMailFoldersCreateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersMailFoldersCreateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersMailFoldersDeleteChildFolders**](UsersMailFolderApi.md#UsersMailFoldersDeleteChildFolders) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Delete navigation property childFolders for users
[**UsersMailFoldersDeleteMessageRules**](UsersMailFolderApi.md#UsersMailFoldersDeleteMessageRules) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Delete navigation property messageRules for users
[**UsersMailFoldersDeleteMessages**](UsersMailFolderApi.md#UsersMailFoldersDeleteMessages) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id} | Delete navigation property messages for users
[**UsersMailFoldersDeleteMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersDeleteMultiValueExtendedProperties) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for users
[**UsersMailFoldersDeleteSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersDeleteSingleValueExtendedProperties) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for users
[**UsersMailFoldersGetChildFolders**](UsersMailFolderApi.md#UsersMailFoldersGetChildFolders) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Get childFolders from users
[**UsersMailFoldersGetMessageRules**](UsersMailFolderApi.md#UsersMailFoldersGetMessageRules) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Get messageRules from users
[**UsersMailFoldersGetMessages**](UsersMailFolderApi.md#UsersMailFoldersGetMessages) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id} | Get messages from users
[**UsersMailFoldersGetMessagesContent**](UsersMailFolderApi.md#UsersMailFoldersGetMessagesContent) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/$value | Get media content for the navigation property messages from users
[**UsersMailFoldersGetMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersGetMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersMailFoldersGetSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersGetSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersMailFoldersListChildFolders**](UsersMailFolderApi.md#UsersMailFoldersListChildFolders) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders | Get childFolders from users
[**UsersMailFoldersListMessageRules**](UsersMailFolderApi.md#UsersMailFoldersListMessageRules) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules | Get messageRules from users
[**UsersMailFoldersListMessages**](UsersMailFolderApi.md#UsersMailFoldersListMessages) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages | Get messages from users
[**UsersMailFoldersListMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersListMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersMailFoldersListSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersListSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesCreateAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateAttachments) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Create new navigation property to attachments for users
[**UsersMailFoldersMessagesCreateExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateExtensions) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Create new navigation property to extensions for users
[**UsersMailFoldersMessagesCreateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersMailFoldersMessagesCreateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersMailFoldersMessagesDeleteAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesDeleteAttachments) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Delete navigation property attachments for users
[**UsersMailFoldersMessagesDeleteExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesDeleteExtensions) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Delete navigation property extensions for users
[**UsersMailFoldersMessagesDeleteMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesDeleteMultiValueExtendedProperties) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for users
[**UsersMailFoldersMessagesDeleteSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesDeleteSingleValueExtendedProperties) | **Delete** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for users
[**UsersMailFoldersMessagesGetAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetAttachments) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Get attachments from users
[**UsersMailFoldersMessagesGetExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetExtensions) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Get extensions from users
[**UsersMailFoldersMessagesGetMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersMailFoldersMessagesGetSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesGetSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesListAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesListAttachments) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments | Get attachments from users
[**UsersMailFoldersMessagesListExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesListExtensions) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions | Get extensions from users
[**UsersMailFoldersMessagesListMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesListMultiValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersMailFoldersMessagesListSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesListSingleValueExtendedProperties) | **Get** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersMailFoldersMessagesUpdateAttachments**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateAttachments) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/attachments/{attachment-id} | Update the navigation property attachments in users
[**UsersMailFoldersMessagesUpdateExtensions**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateExtensions) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersMailFoldersMessagesUpdateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersMailFoldersMessagesUpdateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersMessagesUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersMailFoldersUpdateChildFolders**](UsersMailFolderApi.md#UsersMailFoldersUpdateChildFolders) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1} | Update the navigation property childFolders in users
[**UsersMailFoldersUpdateMessageRules**](UsersMailFolderApi.md#UsersMailFoldersUpdateMessageRules) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messageRules/{messageRule-id} | Update the navigation property messageRules in users
[**UsersMailFoldersUpdateMessages**](UsersMailFolderApi.md#UsersMailFoldersUpdateMessages) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id} | Update the navigation property messages in users
[**UsersMailFoldersUpdateMessagesContent**](UsersMailFolderApi.md#UsersMailFoldersUpdateMessagesContent) | **Put** /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/$value | Update media content for the navigation property messages in users
[**UsersMailFoldersUpdateMultiValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersMailFoldersUpdateSingleValueExtendedProperties**](UsersMailFolderApi.md#UsersMailFoldersUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersUpdateMailFolders**](UsersMailFolderApi.md#UsersUpdateMailFolders) | **Patch** /users/{user-id}/mailFolders/{mailFolder-id} | Update the navigation property mailFolders in users



## UsersCreateMailFolders

> MicrosoftGraphMailFolder UsersCreateMailFolders(ctx, userId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Create new navigation property to mailFolders for users



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
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersCreateMailFolders(context.Background(), userId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersCreateMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateMailFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersCreateMailFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateMailFoldersRequest struct via the builder pattern


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


## UsersDeleteMailFolders

> UsersDeleteMailFolders(ctx, userId, mailFolderId).IfMatch(ifMatch).Execute()

Delete navigation property mailFolders for users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersDeleteMailFolders(context.Background(), userId, mailFolderId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersDeleteMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteMailFoldersRequest struct via the builder pattern


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


## UsersGetMailFolders

> MicrosoftGraphMailFolder UsersGetMailFolders(ctx, userId, mailFolderId).Select_(select_).Execute()

Get mailFolders from users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersGetMailFolders(context.Background(), userId, mailFolderId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersGetMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetMailFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersGetMailFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetMailFoldersRequest struct via the builder pattern


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


## UsersListMailFolders

> CollectionOfMailFolder UsersListMailFolders(ctx, userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get mailFolders from users



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
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersListMailFolders(context.Background(), userId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersListMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListMailFolders`: CollectionOfMailFolder
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersListMailFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListMailFoldersRequest struct via the builder pattern


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


## UsersMailFoldersCreateChildFolders

> MicrosoftGraphMailFolder UsersMailFoldersCreateChildFolders(ctx, userId, mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Create new navigation property to childFolders for users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersCreateChildFolders(context.Background(), userId, mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersCreateChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersCreateChildFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersCreateChildFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersCreateChildFoldersRequest struct via the builder pattern


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


## UsersMailFoldersCreateMessageRules

> MicrosoftGraphMessageRule UsersMailFoldersCreateMessageRules(ctx, userId, mailFolderId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()

Create new navigation property to messageRules for users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMessageRule := *openapiclient.NewMicrosoftGraphMessageRule() // MicrosoftGraphMessageRule | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersCreateMessageRules(context.Background(), userId, mailFolderId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersCreateMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersCreateMessageRules`: MicrosoftGraphMessageRule
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersCreateMessageRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersCreateMessageRulesRequest struct via the builder pattern


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


## UsersMailFoldersCreateMessages

> MicrosoftGraphMessage UsersMailFoldersCreateMessages(ctx, userId, mailFolderId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersCreateMessages(context.Background(), userId, mailFolderId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersCreateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersCreateMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersCreateMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersCreateMessagesRequest struct via the builder pattern


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


## UsersMailFoldersCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersCreateMultiValueExtendedProperties(ctx, userId, mailFolderId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersCreateMultiValueExtendedProperties(context.Background(), userId, mailFolderId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersCreateSingleValueExtendedProperties(ctx, userId, mailFolderId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersCreateSingleValueExtendedProperties(context.Background(), userId, mailFolderId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersDeleteChildFolders

> UsersMailFoldersDeleteChildFolders(ctx, userId, mailFolderId, mailFolderId1).IfMatch(ifMatch).Execute()

Delete navigation property childFolders for users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    mailFolderId1 := "mailFolderId1_example" // string | key: id of mailFolder
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersDeleteChildFolders(context.Background(), userId, mailFolderId, mailFolderId1).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersDeleteChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**mailFolderId1** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersDeleteChildFoldersRequest struct via the builder pattern


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


## UsersMailFoldersDeleteMessageRules

> UsersMailFoldersDeleteMessageRules(ctx, userId, mailFolderId, messageRuleId).IfMatch(ifMatch).Execute()

Delete navigation property messageRules for users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageRuleId := "messageRuleId_example" // string | key: id of messageRule
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersDeleteMessageRules(context.Background(), userId, mailFolderId, messageRuleId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersDeleteMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageRuleId** | **string** | key: id of messageRule | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersDeleteMessageRulesRequest struct via the builder pattern


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


## UsersMailFoldersDeleteMessages

> UsersMailFoldersDeleteMessages(ctx, userId, mailFolderId, messageId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersDeleteMessages(context.Background(), userId, mailFolderId, messageId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersDeleteMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersDeleteMessagesRequest struct via the builder pattern


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


## UsersMailFoldersDeleteMultiValueExtendedProperties

> UsersMailFoldersDeleteMultiValueExtendedProperties(ctx, userId, mailFolderId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersDeleteMultiValueExtendedProperties(context.Background(), userId, mailFolderId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersDeleteSingleValueExtendedProperties

> UsersMailFoldersDeleteSingleValueExtendedProperties(ctx, userId, mailFolderId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersDeleteSingleValueExtendedProperties(context.Background(), userId, mailFolderId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersGetChildFolders

> MicrosoftGraphMailFolder UsersMailFoldersGetChildFolders(ctx, userId, mailFolderId, mailFolderId1).Select_(select_).Expand(expand).Execute()

Get childFolders from users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    mailFolderId1 := "mailFolderId1_example" // string | key: id of mailFolder
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersGetChildFolders(context.Background(), userId, mailFolderId, mailFolderId1).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersGetChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersGetChildFolders`: MicrosoftGraphMailFolder
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersGetChildFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**mailFolderId1** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersGetChildFoldersRequest struct via the builder pattern


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


## UsersMailFoldersGetMessageRules

> MicrosoftGraphMessageRule UsersMailFoldersGetMessageRules(ctx, userId, mailFolderId, messageRuleId).Select_(select_).Execute()

Get messageRules from users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageRuleId := "messageRuleId_example" // string | key: id of messageRule
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersGetMessageRules(context.Background(), userId, mailFolderId, messageRuleId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersGetMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersGetMessageRules`: MicrosoftGraphMessageRule
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersGetMessageRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageRuleId** | **string** | key: id of messageRule | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersGetMessageRulesRequest struct via the builder pattern


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


## UsersMailFoldersGetMessages

> MicrosoftGraphMessage UsersMailFoldersGetMessages(ctx, userId, mailFolderId, messageId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersGetMessages(context.Background(), userId, mailFolderId, messageId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersGetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersGetMessages`: MicrosoftGraphMessage
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersGetMessagesRequest struct via the builder pattern


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


## UsersMailFoldersGetMessagesContent

> *os.File UsersMailFoldersGetMessagesContent(ctx, userId, mailFolderId, messageId).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersGetMessagesContent(context.Background(), userId, mailFolderId, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersGetMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersGetMessagesContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersGetMessagesContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersGetMessagesContentRequest struct via the builder pattern


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


## UsersMailFoldersGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersGetMultiValueExtendedProperties(ctx, userId, mailFolderId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersGetMultiValueExtendedProperties(context.Background(), userId, mailFolderId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersGetSingleValueExtendedProperties(ctx, userId, mailFolderId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersGetSingleValueExtendedProperties(context.Background(), userId, mailFolderId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersListChildFolders

> CollectionOfMailFolder UsersMailFoldersListChildFolders(ctx, userId, mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get childFolders from users



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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersListChildFolders(context.Background(), userId, mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersListChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersListChildFolders`: CollectionOfMailFolder
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersListChildFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersListChildFoldersRequest struct via the builder pattern


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


## UsersMailFoldersListMessageRules

> CollectionOfMessageRule UsersMailFoldersListMessageRules(ctx, userId, mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get messageRules from users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersListMessageRules(context.Background(), userId, mailFolderId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersListMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersListMessageRules`: CollectionOfMessageRule
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersListMessageRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersListMessageRulesRequest struct via the builder pattern


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


## UsersMailFoldersListMessages

> CollectionOfMessage UsersMailFoldersListMessages(ctx, userId, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersListMessages(context.Background(), userId, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersListMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersListMessages`: CollectionOfMessage
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersListMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersListMessagesRequest struct via the builder pattern


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


## UsersMailFoldersListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersMailFoldersListMultiValueExtendedProperties(ctx, userId, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersListMultiValueExtendedProperties(context.Background(), userId, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersMailFoldersListSingleValueExtendedProperties(ctx, userId, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersListSingleValueExtendedProperties(context.Background(), userId, mailFolderId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesCreateAttachments

> MicrosoftGraphAttachment UsersMailFoldersMessagesCreateAttachments(ctx, userId, mailFolderId, messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesCreateAttachments(context.Background(), userId, mailFolderId, messageId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesCreateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesCreateAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesCreateAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesCreateAttachmentsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesCreateExtensions

> MicrosoftGraphExtension UsersMailFoldersMessagesCreateExtensions(ctx, userId, mailFolderId, messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesCreateExtensions(context.Background(), userId, mailFolderId, messageId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesCreateExtensionsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersMessagesCreateMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesCreateMultiValueExtendedProperties(context.Background(), userId, mailFolderId, messageId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersMessagesCreateSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesCreateSingleValueExtendedProperties(context.Background(), userId, mailFolderId, messageId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesDeleteAttachments

> UsersMailFoldersMessagesDeleteAttachments(ctx, userId, mailFolderId, messageId, attachmentId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesDeleteAttachments(context.Background(), userId, mailFolderId, messageId, attachmentId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesDeleteAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesDeleteAttachmentsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesDeleteExtensions

> UsersMailFoldersMessagesDeleteExtensions(ctx, userId, mailFolderId, messageId, extensionId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesDeleteExtensions(context.Background(), userId, mailFolderId, messageId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesDeleteExtensionsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesDeleteMultiValueExtendedProperties

> UsersMailFoldersMessagesDeleteMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesDeleteMultiValueExtendedProperties(context.Background(), userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesDeleteSingleValueExtendedProperties

> UsersMailFoldersMessagesDeleteSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesDeleteSingleValueExtendedProperties(context.Background(), userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesGetAttachments

> MicrosoftGraphAttachment UsersMailFoldersMessagesGetAttachments(ctx, userId, mailFolderId, messageId, attachmentId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesGetAttachments(context.Background(), userId, mailFolderId, messageId, attachmentId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesGetAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesGetAttachments`: MicrosoftGraphAttachment
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesGetAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesGetAttachmentsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesGetExtensions

> MicrosoftGraphExtension UsersMailFoldersMessagesGetExtensions(ctx, userId, mailFolderId, messageId, extensionId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesGetExtensions(context.Background(), userId, mailFolderId, messageId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesGetExtensionsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersMailFoldersMessagesGetMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesGetMultiValueExtendedProperties(context.Background(), userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersMailFoldersMessagesGetSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesGetSingleValueExtendedProperties(context.Background(), userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesListAttachments

> CollectionOfAttachment UsersMailFoldersMessagesListAttachments(ctx, userId, mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesListAttachments(context.Background(), userId, mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesListAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesListAttachments`: CollectionOfAttachment
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesListAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesListAttachmentsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesListExtensions

> CollectionOfExtension UsersMailFoldersMessagesListExtensions(ctx, userId, mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesListExtensions(context.Background(), userId, mailFolderId, messageId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesListExtensionsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersMailFoldersMessagesListMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesListMultiValueExtendedProperties(context.Background(), userId, mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersMailFoldersMessagesListSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesListSingleValueExtendedProperties(context.Background(), userId, mailFolderId, messageId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersMailFoldersMessagesListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersMailFolderApi.UsersMailFoldersMessagesListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesUpdateAttachments

> UsersMailFoldersMessagesUpdateAttachments(ctx, userId, mailFolderId, messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    attachmentId := "attachmentId_example" // string | key: id of attachment
    microsoftGraphAttachment := *openapiclient.NewMicrosoftGraphAttachment() // MicrosoftGraphAttachment | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesUpdateAttachments(context.Background(), userId, mailFolderId, messageId, attachmentId).MicrosoftGraphAttachment(microsoftGraphAttachment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesUpdateAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**attachmentId** | **string** | key: id of attachment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesUpdateAttachmentsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesUpdateExtensions

> UsersMailFoldersMessagesUpdateExtensions(ctx, userId, mailFolderId, messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesUpdateExtensions(context.Background(), userId, mailFolderId, messageId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesUpdateExtensionsRequest struct via the builder pattern


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


## UsersMailFoldersMessagesUpdateMultiValueExtendedProperties

> UsersMailFoldersMessagesUpdateMultiValueExtendedProperties(ctx, userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesUpdateMultiValueExtendedProperties(context.Background(), userId, mailFolderId, messageId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersMessagesUpdateSingleValueExtendedProperties

> UsersMailFoldersMessagesUpdateSingleValueExtendedProperties(ctx, userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersMessagesUpdateSingleValueExtendedProperties(context.Background(), userId, mailFolderId, messageId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersMessagesUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersMessagesUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersUpdateChildFolders

> UsersMailFoldersUpdateChildFolders(ctx, userId, mailFolderId, mailFolderId1).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Update the navigation property childFolders in users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    mailFolderId1 := "mailFolderId1_example" // string | key: id of mailFolder
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersUpdateChildFolders(context.Background(), userId, mailFolderId, mailFolderId1).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersUpdateChildFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**mailFolderId1** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersUpdateChildFoldersRequest struct via the builder pattern


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


## UsersMailFoldersUpdateMessageRules

> UsersMailFoldersUpdateMessageRules(ctx, userId, mailFolderId, messageRuleId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()

Update the navigation property messageRules in users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageRuleId := "messageRuleId_example" // string | key: id of messageRule
    microsoftGraphMessageRule := *openapiclient.NewMicrosoftGraphMessageRule() // MicrosoftGraphMessageRule | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersUpdateMessageRules(context.Background(), userId, mailFolderId, messageRuleId).MicrosoftGraphMessageRule(microsoftGraphMessageRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersUpdateMessageRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageRuleId** | **string** | key: id of messageRule | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersUpdateMessageRulesRequest struct via the builder pattern


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


## UsersMailFoldersUpdateMessages

> UsersMailFoldersUpdateMessages(ctx, userId, mailFolderId, messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    microsoftGraphMessage := *openapiclient.NewMicrosoftGraphMessage() // MicrosoftGraphMessage | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersUpdateMessages(context.Background(), userId, mailFolderId, messageId).MicrosoftGraphMessage(microsoftGraphMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersUpdateMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersUpdateMessagesRequest struct via the builder pattern


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


## UsersMailFoldersUpdateMessagesContent

> UsersMailFoldersUpdateMessagesContent(ctx, userId, mailFolderId, messageId).Body(body).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    messageId := "messageId_example" // string | key: id of message
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersUpdateMessagesContent(context.Background(), userId, mailFolderId, messageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersUpdateMessagesContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**messageId** | **string** | key: id of message | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersUpdateMessagesContentRequest struct via the builder pattern


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


## UsersMailFoldersUpdateMultiValueExtendedProperties

> UsersMailFoldersUpdateMultiValueExtendedProperties(ctx, userId, mailFolderId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersUpdateMultiValueExtendedProperties(context.Background(), userId, mailFolderId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersMailFoldersUpdateSingleValueExtendedProperties

> UsersMailFoldersUpdateSingleValueExtendedProperties(ctx, userId, mailFolderId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersMailFoldersUpdateSingleValueExtendedProperties(context.Background(), userId, mailFolderId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersMailFoldersUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMailFoldersUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersUpdateMailFolders

> UsersUpdateMailFolders(ctx, userId, mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()

Update the navigation property mailFolders in users



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
    mailFolderId := "mailFolderId_example" // string | key: id of mailFolder
    microsoftGraphMailFolder := *openapiclient.NewMicrosoftGraphMailFolder() // MicrosoftGraphMailFolder | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersMailFolderApi.UsersUpdateMailFolders(context.Background(), userId, mailFolderId).MicrosoftGraphMailFolder(microsoftGraphMailFolder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMailFolderApi.UsersUpdateMailFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**mailFolderId** | **string** | key: id of mailFolder | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateMailFoldersRequest struct via the builder pattern


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

