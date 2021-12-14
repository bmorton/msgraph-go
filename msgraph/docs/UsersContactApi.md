# \UsersContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersContactsCreateExtensions**](UsersContactApi.md#UsersContactsCreateExtensions) | **Post** /users/{user-id}/contacts/{contact-id}/extensions | Create new navigation property to extensions for users
[**UsersContactsCreateMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsCreateMultiValueExtendedProperties) | **Post** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for users
[**UsersContactsCreateSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsCreateSingleValueExtendedProperties) | **Post** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for users
[**UsersContactsDeleteExtensions**](UsersContactApi.md#UsersContactsDeleteExtensions) | **Delete** /users/{user-id}/contacts/{contact-id}/extensions/{extension-id} | Delete navigation property extensions for users
[**UsersContactsDeleteMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsDeleteMultiValueExtendedProperties) | **Delete** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for users
[**UsersContactsDeletePhoto**](UsersContactApi.md#UsersContactsDeletePhoto) | **Delete** /users/{user-id}/contacts/{contact-id}/photo | Delete navigation property photo for users
[**UsersContactsDeleteSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsDeleteSingleValueExtendedProperties) | **Delete** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for users
[**UsersContactsGetExtensions**](UsersContactApi.md#UsersContactsGetExtensions) | **Get** /users/{user-id}/contacts/{contact-id}/extensions/{extension-id} | Get extensions from users
[**UsersContactsGetMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsGetMultiValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from users
[**UsersContactsGetPhoto**](UsersContactApi.md#UsersContactsGetPhoto) | **Get** /users/{user-id}/contacts/{contact-id}/photo | Get photo from users
[**UsersContactsGetPhotoContent**](UsersContactApi.md#UsersContactsGetPhotoContent) | **Get** /users/{user-id}/contacts/{contact-id}/photo/$value | Get media content for the navigation property photo from users
[**UsersContactsGetSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsGetSingleValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from users
[**UsersContactsListExtensions**](UsersContactApi.md#UsersContactsListExtensions) | **Get** /users/{user-id}/contacts/{contact-id}/extensions | Get extensions from users
[**UsersContactsListMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsListMultiValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from users
[**UsersContactsListSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsListSingleValueExtendedProperties) | **Get** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from users
[**UsersContactsUpdateExtensions**](UsersContactApi.md#UsersContactsUpdateExtensions) | **Patch** /users/{user-id}/contacts/{contact-id}/extensions/{extension-id} | Update the navigation property extensions in users
[**UsersContactsUpdateMultiValueExtendedProperties**](UsersContactApi.md#UsersContactsUpdateMultiValueExtendedProperties) | **Patch** /users/{user-id}/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in users
[**UsersContactsUpdatePhoto**](UsersContactApi.md#UsersContactsUpdatePhoto) | **Patch** /users/{user-id}/contacts/{contact-id}/photo | Update the navigation property photo in users
[**UsersContactsUpdatePhotoContent**](UsersContactApi.md#UsersContactsUpdatePhotoContent) | **Put** /users/{user-id}/contacts/{contact-id}/photo/$value | Update media content for the navigation property photo in users
[**UsersContactsUpdateSingleValueExtendedProperties**](UsersContactApi.md#UsersContactsUpdateSingleValueExtendedProperties) | **Patch** /users/{user-id}/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in users
[**UsersCreateContacts**](UsersContactApi.md#UsersCreateContacts) | **Post** /users/{user-id}/contacts | Create new navigation property to contacts for users
[**UsersDeleteContacts**](UsersContactApi.md#UsersDeleteContacts) | **Delete** /users/{user-id}/contacts/{contact-id} | Delete navigation property contacts for users
[**UsersGetContacts**](UsersContactApi.md#UsersGetContacts) | **Get** /users/{user-id}/contacts/{contact-id} | Get contacts from users
[**UsersListContacts**](UsersContactApi.md#UsersListContacts) | **Get** /users/{user-id}/contacts | Get contacts from users
[**UsersUpdateContacts**](UsersContactApi.md#UsersUpdateContacts) | **Patch** /users/{user-id}/contacts/{contact-id} | Update the navigation property contacts in users



## UsersContactsCreateExtensions

> MicrosoftGraphExtension UsersContactsCreateExtensions(ctx, userId, contactId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsCreateExtensions(context.Background(), userId, contactId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsCreateExtensionsRequest struct via the builder pattern


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


## UsersContactsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersContactsCreateMultiValueExtendedProperties(ctx, userId, contactId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsCreateMultiValueExtendedProperties(context.Background(), userId, contactId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersContactsCreateSingleValueExtendedProperties(ctx, userId, contactId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsCreateSingleValueExtendedProperties(context.Background(), userId, contactId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsDeleteExtensions

> UsersContactsDeleteExtensions(ctx, userId, contactId, extensionId).IfMatch(ifMatch).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsDeleteExtensions(context.Background(), userId, contactId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsDeleteExtensionsRequest struct via the builder pattern


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


## UsersContactsDeleteMultiValueExtendedProperties

> UsersContactsDeleteMultiValueExtendedProperties(ctx, userId, contactId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsDeleteMultiValueExtendedProperties(context.Background(), userId, contactId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsDeletePhoto

> UsersContactsDeletePhoto(ctx, userId, contactId).IfMatch(ifMatch).Execute()

Delete navigation property photo for users



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
    contactId := "contactId_example" // string | key: id of contact
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsDeletePhoto(context.Background(), userId, contactId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsDeletePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsDeletePhotoRequest struct via the builder pattern


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


## UsersContactsDeleteSingleValueExtendedProperties

> UsersContactsDeleteSingleValueExtendedProperties(ctx, userId, contactId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsDeleteSingleValueExtendedProperties(context.Background(), userId, contactId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsGetExtensions

> MicrosoftGraphExtension UsersContactsGetExtensions(ctx, userId, contactId, extensionId).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsGetExtensions(context.Background(), userId, contactId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsGetExtensionsRequest struct via the builder pattern


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


## UsersContactsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty UsersContactsGetMultiValueExtendedProperties(ctx, userId, contactId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsGetMultiValueExtendedProperties(context.Background(), userId, contactId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsGetPhoto

> MicrosoftGraphProfilePhoto UsersContactsGetPhoto(ctx, userId, contactId).Select_(select_).Execute()

Get photo from users



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
    contactId := "contactId_example" // string | key: id of contact
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsGetPhoto(context.Background(), userId, contactId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsGetPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsGetPhoto`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsGetPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsGetPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersContactsGetPhotoContent

> *os.File UsersContactsGetPhotoContent(ctx, userId, contactId).Execute()

Get media content for the navigation property photo from users



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
    contactId := "contactId_example" // string | key: id of contact

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsGetPhotoContent(context.Background(), userId, contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsGetPhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsGetPhotoContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsGetPhotoContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsGetPhotoContentRequest struct via the builder pattern


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


## UsersContactsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty UsersContactsGetSingleValueExtendedProperties(ctx, userId, contactId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsGetSingleValueExtendedProperties(context.Background(), userId, contactId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsListExtensions

> CollectionOfExtension UsersContactsListExtensions(ctx, userId, contactId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsListExtensions(context.Background(), userId, contactId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsListExtensionsRequest struct via the builder pattern


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


## UsersContactsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty UsersContactsListMultiValueExtendedProperties(ctx, userId, contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
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
    resp, r, err := api_client.UsersContactApi.UsersContactsListMultiValueExtendedProperties(context.Background(), userId, contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty UsersContactsListSingleValueExtendedProperties(ctx, userId, contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
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
    resp, r, err := api_client.UsersContactApi.UsersContactsListSingleValueExtendedProperties(context.Background(), userId, contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersContactsListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersContactsListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsUpdateExtensions

> UsersContactsUpdateExtensions(ctx, userId, contactId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsUpdateExtensions(context.Background(), userId, contactId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsUpdateExtensionsRequest struct via the builder pattern


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


## UsersContactsUpdateMultiValueExtendedProperties

> UsersContactsUpdateMultiValueExtendedProperties(ctx, userId, contactId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsUpdateMultiValueExtendedProperties(context.Background(), userId, contactId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersContactsUpdatePhoto

> UsersContactsUpdatePhoto(ctx, userId, contactId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photo in users



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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsUpdatePhoto(context.Background(), userId, contactId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsUpdatePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsUpdatePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphProfilePhoto** | [**MicrosoftGraphProfilePhoto**](MicrosoftGraphProfilePhoto.md) | New navigation property values | 

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


## UsersContactsUpdatePhotoContent

> UsersContactsUpdatePhotoContent(ctx, userId, contactId).Body(body).Execute()

Update media content for the navigation property photo in users



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
    contactId := "contactId_example" // string | key: id of contact
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsUpdatePhotoContent(context.Background(), userId, contactId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsUpdatePhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsUpdatePhotoContentRequest struct via the builder pattern


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


## UsersContactsUpdateSingleValueExtendedProperties

> UsersContactsUpdateSingleValueExtendedProperties(ctx, userId, contactId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersContactsUpdateSingleValueExtendedProperties(context.Background(), userId, contactId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersContactsUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersContactsUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## UsersCreateContacts

> MicrosoftGraphContact UsersCreateContacts(ctx, userId).MicrosoftGraphContact(microsoftGraphContact).Execute()

Create new navigation property to contacts for users



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
    microsoftGraphContact := *openapiclient.NewMicrosoftGraphContact() // MicrosoftGraphContact | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersCreateContacts(context.Background(), userId).MicrosoftGraphContact(microsoftGraphContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersCreateContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateContacts`: MicrosoftGraphContact
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersCreateContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphContact** | [**MicrosoftGraphContact**](MicrosoftGraphContact.md) | New navigation property | 

### Return type

[**MicrosoftGraphContact**](MicrosoftGraphContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteContacts

> UsersDeleteContacts(ctx, userId, contactId).IfMatch(ifMatch).Execute()

Delete navigation property contacts for users



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
    contactId := "contactId_example" // string | key: id of contact
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersDeleteContacts(context.Background(), userId, contactId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersDeleteContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteContactsRequest struct via the builder pattern


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


## UsersGetContacts

> MicrosoftGraphContact UsersGetContacts(ctx, userId, contactId).Select_(select_).Execute()

Get contacts from users



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
    contactId := "contactId_example" // string | key: id of contact
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersGetContacts(context.Background(), userId, contactId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersGetContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetContacts`: MicrosoftGraphContact
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersGetContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphContact**](MicrosoftGraphContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListContacts

> CollectionOfContact UsersListContacts(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get contacts from users



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
    resp, r, err := api_client.UsersContactApi.UsersListContacts(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersListContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListContacts`: CollectionOfContact
    fmt.Fprintf(os.Stdout, "Response from `UsersContactApi.UsersListContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListContactsRequest struct via the builder pattern


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

[**CollectionOfContact**](CollectionOfContact.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateContacts

> UsersUpdateContacts(ctx, userId, contactId).MicrosoftGraphContact(microsoftGraphContact).Execute()

Update the navigation property contacts in users



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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphContact := *openapiclient.NewMicrosoftGraphContact() // MicrosoftGraphContact | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersContactApi.UsersUpdateContacts(context.Background(), userId, contactId).MicrosoftGraphContact(microsoftGraphContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersContactApi.UsersUpdateContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphContact** | [**MicrosoftGraphContact**](MicrosoftGraphContact.md) | New navigation property values | 

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

