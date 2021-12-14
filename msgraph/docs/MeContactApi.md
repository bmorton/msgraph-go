# \MeContactApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeContactsCreateExtensions**](MeContactApi.md#MeContactsCreateExtensions) | **Post** /me/contacts/{contact-id}/extensions | Create new navigation property to extensions for me
[**MeContactsCreateMultiValueExtendedProperties**](MeContactApi.md#MeContactsCreateMultiValueExtendedProperties) | **Post** /me/contacts/{contact-id}/multiValueExtendedProperties | Create new navigation property to multiValueExtendedProperties for me
[**MeContactsCreateSingleValueExtendedProperties**](MeContactApi.md#MeContactsCreateSingleValueExtendedProperties) | **Post** /me/contacts/{contact-id}/singleValueExtendedProperties | Create new navigation property to singleValueExtendedProperties for me
[**MeContactsDeleteExtensions**](MeContactApi.md#MeContactsDeleteExtensions) | **Delete** /me/contacts/{contact-id}/extensions/{extension-id} | Delete navigation property extensions for me
[**MeContactsDeleteMultiValueExtendedProperties**](MeContactApi.md#MeContactsDeleteMultiValueExtendedProperties) | **Delete** /me/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Delete navigation property multiValueExtendedProperties for me
[**MeContactsDeletePhoto**](MeContactApi.md#MeContactsDeletePhoto) | **Delete** /me/contacts/{contact-id}/photo | Delete navigation property photo for me
[**MeContactsDeleteSingleValueExtendedProperties**](MeContactApi.md#MeContactsDeleteSingleValueExtendedProperties) | **Delete** /me/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Delete navigation property singleValueExtendedProperties for me
[**MeContactsGetExtensions**](MeContactApi.md#MeContactsGetExtensions) | **Get** /me/contacts/{contact-id}/extensions/{extension-id} | Get extensions from me
[**MeContactsGetMultiValueExtendedProperties**](MeContactApi.md#MeContactsGetMultiValueExtendedProperties) | **Get** /me/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Get multiValueExtendedProperties from me
[**MeContactsGetPhoto**](MeContactApi.md#MeContactsGetPhoto) | **Get** /me/contacts/{contact-id}/photo | Get photo from me
[**MeContactsGetPhotoContent**](MeContactApi.md#MeContactsGetPhotoContent) | **Get** /me/contacts/{contact-id}/photo/$value | Get media content for the navigation property photo from me
[**MeContactsGetSingleValueExtendedProperties**](MeContactApi.md#MeContactsGetSingleValueExtendedProperties) | **Get** /me/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Get singleValueExtendedProperties from me
[**MeContactsListExtensions**](MeContactApi.md#MeContactsListExtensions) | **Get** /me/contacts/{contact-id}/extensions | Get extensions from me
[**MeContactsListMultiValueExtendedProperties**](MeContactApi.md#MeContactsListMultiValueExtendedProperties) | **Get** /me/contacts/{contact-id}/multiValueExtendedProperties | Get multiValueExtendedProperties from me
[**MeContactsListSingleValueExtendedProperties**](MeContactApi.md#MeContactsListSingleValueExtendedProperties) | **Get** /me/contacts/{contact-id}/singleValueExtendedProperties | Get singleValueExtendedProperties from me
[**MeContactsUpdateExtensions**](MeContactApi.md#MeContactsUpdateExtensions) | **Patch** /me/contacts/{contact-id}/extensions/{extension-id} | Update the navigation property extensions in me
[**MeContactsUpdateMultiValueExtendedProperties**](MeContactApi.md#MeContactsUpdateMultiValueExtendedProperties) | **Patch** /me/contacts/{contact-id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty-id} | Update the navigation property multiValueExtendedProperties in me
[**MeContactsUpdatePhoto**](MeContactApi.md#MeContactsUpdatePhoto) | **Patch** /me/contacts/{contact-id}/photo | Update the navigation property photo in me
[**MeContactsUpdatePhotoContent**](MeContactApi.md#MeContactsUpdatePhotoContent) | **Put** /me/contacts/{contact-id}/photo/$value | Update media content for the navigation property photo in me
[**MeContactsUpdateSingleValueExtendedProperties**](MeContactApi.md#MeContactsUpdateSingleValueExtendedProperties) | **Patch** /me/contacts/{contact-id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty-id} | Update the navigation property singleValueExtendedProperties in me
[**MeCreateContacts**](MeContactApi.md#MeCreateContacts) | **Post** /me/contacts | Create new navigation property to contacts for me
[**MeDeleteContacts**](MeContactApi.md#MeDeleteContacts) | **Delete** /me/contacts/{contact-id} | Delete navigation property contacts for me
[**MeGetContacts**](MeContactApi.md#MeGetContacts) | **Get** /me/contacts/{contact-id} | Get contacts from me
[**MeListContacts**](MeContactApi.md#MeListContacts) | **Get** /me/contacts | Get contacts from me
[**MeUpdateContacts**](MeContactApi.md#MeUpdateContacts) | **Patch** /me/contacts/{contact-id} | Update the navigation property contacts in me



## MeContactsCreateExtensions

> MicrosoftGraphExtension MeContactsCreateExtensions(ctx, contactId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsCreateExtensions(context.Background(), contactId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsCreateExtensionsRequest struct via the builder pattern


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


## MeContactsCreateMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeContactsCreateMultiValueExtendedProperties(ctx, contactId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsCreateMultiValueExtendedProperties(context.Background(), contactId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsCreateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsCreateMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsCreateMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsCreateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsCreateSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeContactsCreateSingleValueExtendedProperties(ctx, contactId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsCreateSingleValueExtendedProperties(context.Background(), contactId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsCreateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsCreateSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsCreateSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsCreateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsDeleteExtensions

> MeContactsDeleteExtensions(ctx, contactId, extensionId).IfMatch(ifMatch).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsDeleteExtensions(context.Background(), contactId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsDeleteExtensionsRequest struct via the builder pattern


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


## MeContactsDeleteMultiValueExtendedProperties

> MeContactsDeleteMultiValueExtendedProperties(ctx, contactId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsDeleteMultiValueExtendedProperties(context.Background(), contactId, multiValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsDeleteMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsDeleteMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsDeletePhoto

> MeContactsDeletePhoto(ctx, contactId).IfMatch(ifMatch).Execute()

Delete navigation property photo for me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsDeletePhoto(context.Background(), contactId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsDeletePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsDeletePhotoRequest struct via the builder pattern


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


## MeContactsDeleteSingleValueExtendedProperties

> MeContactsDeleteSingleValueExtendedProperties(ctx, contactId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsDeleteSingleValueExtendedProperties(context.Background(), contactId, singleValueLegacyExtendedPropertyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsDeleteSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsDeleteSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsGetExtensions

> MicrosoftGraphExtension MeContactsGetExtensions(ctx, contactId, extensionId).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsGetExtensions(context.Background(), contactId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsGetExtensionsRequest struct via the builder pattern


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


## MeContactsGetMultiValueExtendedProperties

> MicrosoftGraphMultiValueLegacyExtendedProperty MeContactsGetMultiValueExtendedProperties(ctx, contactId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsGetMultiValueExtendedProperties(context.Background(), contactId, multiValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsGetMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsGetMultiValueExtendedProperties`: MicrosoftGraphMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsGetMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsGetMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsGetPhoto

> MicrosoftGraphProfilePhoto MeContactsGetPhoto(ctx, contactId).Select_(select_).Execute()

Get photo from me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsGetPhoto(context.Background(), contactId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsGetPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsGetPhoto`: MicrosoftGraphProfilePhoto
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsGetPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsGetPhotoRequest struct via the builder pattern


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


## MeContactsGetPhotoContent

> *os.File MeContactsGetPhotoContent(ctx, contactId).Execute()

Get media content for the navigation property photo from me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsGetPhotoContent(context.Background(), contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsGetPhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsGetPhotoContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsGetPhotoContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsGetPhotoContentRequest struct via the builder pattern


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


## MeContactsGetSingleValueExtendedProperties

> MicrosoftGraphSingleValueLegacyExtendedProperty MeContactsGetSingleValueExtendedProperties(ctx, contactId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsGetSingleValueExtendedProperties(context.Background(), contactId, singleValueLegacyExtendedPropertyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsGetSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsGetSingleValueExtendedProperties`: MicrosoftGraphSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsGetSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsGetSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsListExtensions

> CollectionOfExtension MeContactsListExtensions(ctx, contactId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.MeContactApi.MeContactsListExtensions(context.Background(), contactId).Top(top).Skip(skip).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsListExtensionsRequest struct via the builder pattern


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


## MeContactsListMultiValueExtendedProperties

> CollectionOfMultiValueLegacyExtendedProperty MeContactsListMultiValueExtendedProperties(ctx, contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.MeContactApi.MeContactsListMultiValueExtendedProperties(context.Background(), contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsListMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsListMultiValueExtendedProperties`: CollectionOfMultiValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsListMultiValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsListMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsListSingleValueExtendedProperties

> CollectionOfSingleValueLegacyExtendedProperty MeContactsListSingleValueExtendedProperties(ctx, contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    resp, r, err := api_client.MeContactApi.MeContactsListSingleValueExtendedProperties(context.Background(), contactId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsListSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeContactsListSingleValueExtendedProperties`: CollectionOfSingleValueLegacyExtendedProperty
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeContactsListSingleValueExtendedProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsListSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsUpdateExtensions

> MeContactsUpdateExtensions(ctx, contactId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsUpdateExtensions(context.Background(), contactId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsUpdateExtensionsRequest struct via the builder pattern


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


## MeContactsUpdateMultiValueExtendedProperties

> MeContactsUpdateMultiValueExtendedProperties(ctx, contactId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    multiValueLegacyExtendedPropertyId := "multiValueLegacyExtendedPropertyId_example" // string | key: id of multiValueLegacyExtendedProperty
    microsoftGraphMultiValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphMultiValueLegacyExtendedProperty() // MicrosoftGraphMultiValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsUpdateMultiValueExtendedProperties(context.Background(), contactId, multiValueLegacyExtendedPropertyId).MicrosoftGraphMultiValueLegacyExtendedProperty(microsoftGraphMultiValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsUpdateMultiValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**multiValueLegacyExtendedPropertyId** | **string** | key: id of multiValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsUpdateMultiValueExtendedPropertiesRequest struct via the builder pattern


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


## MeContactsUpdatePhoto

> MeContactsUpdatePhoto(ctx, contactId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()

Update the navigation property photo in me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphProfilePhoto := *openapiclient.NewMicrosoftGraphProfilePhoto() // MicrosoftGraphProfilePhoto | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsUpdatePhoto(context.Background(), contactId).MicrosoftGraphProfilePhoto(microsoftGraphProfilePhoto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsUpdatePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsUpdatePhotoRequest struct via the builder pattern


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


## MeContactsUpdatePhotoContent

> MeContactsUpdatePhotoContent(ctx, contactId).Body(body).Execute()

Update media content for the navigation property photo in me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsUpdatePhotoContent(context.Background(), contactId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsUpdatePhotoContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsUpdatePhotoContentRequest struct via the builder pattern


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


## MeContactsUpdateSingleValueExtendedProperties

> MeContactsUpdateSingleValueExtendedProperties(ctx, contactId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()

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
    contactId := "contactId_example" // string | key: id of contact
    singleValueLegacyExtendedPropertyId := "singleValueLegacyExtendedPropertyId_example" // string | key: id of singleValueLegacyExtendedProperty
    microsoftGraphSingleValueLegacyExtendedProperty := *openapiclient.NewMicrosoftGraphSingleValueLegacyExtendedProperty() // MicrosoftGraphSingleValueLegacyExtendedProperty | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeContactsUpdateSingleValueExtendedProperties(context.Background(), contactId, singleValueLegacyExtendedPropertyId).MicrosoftGraphSingleValueLegacyExtendedProperty(microsoftGraphSingleValueLegacyExtendedProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeContactsUpdateSingleValueExtendedProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 
**singleValueLegacyExtendedPropertyId** | **string** | key: id of singleValueLegacyExtendedProperty | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeContactsUpdateSingleValueExtendedPropertiesRequest struct via the builder pattern


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


## MeCreateContacts

> MicrosoftGraphContact MeCreateContacts(ctx).MicrosoftGraphContact(microsoftGraphContact).Execute()

Create new navigation property to contacts for me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    microsoftGraphContact := *openapiclient.NewMicrosoftGraphContact() // MicrosoftGraphContact | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeCreateContacts(context.Background()).MicrosoftGraphContact(microsoftGraphContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeCreateContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateContacts`: MicrosoftGraphContact
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeCreateContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateContactsRequest struct via the builder pattern


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


## MeDeleteContacts

> MeDeleteContacts(ctx, contactId).IfMatch(ifMatch).Execute()

Delete navigation property contacts for me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeDeleteContacts(context.Background(), contactId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeDeleteContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteContactsRequest struct via the builder pattern


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


## MeGetContacts

> MicrosoftGraphContact MeGetContacts(ctx, contactId).Select_(select_).Execute()

Get contacts from me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeGetContacts(context.Background(), contactId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeGetContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetContacts`: MicrosoftGraphContact
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeGetContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetContactsRequest struct via the builder pattern


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


## MeListContacts

> CollectionOfContact MeListContacts(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()

Get contacts from me



### Example

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
    resp, r, err := api_client.MeContactApi.MeListContacts(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeListContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListContacts`: CollectionOfContact
    fmt.Fprintf(os.Stdout, "Response from `MeContactApi.MeListContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListContactsRequest struct via the builder pattern


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


## MeUpdateContacts

> MeUpdateContacts(ctx, contactId).MicrosoftGraphContact(microsoftGraphContact).Execute()

Update the navigation property contacts in me



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contactId := "contactId_example" // string | key: id of contact
    microsoftGraphContact := *openapiclient.NewMicrosoftGraphContact() // MicrosoftGraphContact | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeContactApi.MeUpdateContacts(context.Background(), contactId).MicrosoftGraphContact(microsoftGraphContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeContactApi.MeUpdateContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** | key: id of contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateContactsRequest struct via the builder pattern


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

