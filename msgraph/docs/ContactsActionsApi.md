# \ContactsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactsGetAvailableExtensionProperties**](ContactsActionsApi.md#ContactsGetAvailableExtensionProperties) | **Post** /contacts/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**ContactsGetByIds**](ContactsActionsApi.md#ContactsGetByIds) | **Post** /contacts/microsoft.graph.getByIds | Invoke action getByIds
[**ContactsOrgContactCheckMemberGroups**](ContactsActionsApi.md#ContactsOrgContactCheckMemberGroups) | **Post** /contacts/{orgContact-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**ContactsOrgContactCheckMemberObjects**](ContactsActionsApi.md#ContactsOrgContactCheckMemberObjects) | **Post** /contacts/{orgContact-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**ContactsOrgContactGetMemberGroups**](ContactsActionsApi.md#ContactsOrgContactGetMemberGroups) | **Post** /contacts/{orgContact-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**ContactsOrgContactGetMemberObjects**](ContactsActionsApi.md#ContactsOrgContactGetMemberObjects) | **Post** /contacts/{orgContact-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**ContactsOrgContactRestore**](ContactsActionsApi.md#ContactsOrgContactRestore) | **Post** /contacts/{orgContact-id}/microsoft.graph.restore | Invoke action restore
[**ContactsValidateProperties**](ContactsActionsApi.md#ContactsValidateProperties) | **Post** /contacts/microsoft.graph.validateProperties | Invoke action validateProperties



## ContactsGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty ContactsGetAvailableExtensionProperties(ctx).InlineObject46(inlineObject46).Execute()

Invoke action getAvailableExtensionProperties

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject46 := *openapiclient.NewInlineObject46() // InlineObject46 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsGetAvailableExtensionProperties(context.Background()).InlineObject46(inlineObject46).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `ContactsActionsApi.ContactsGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject46** | [**InlineObject46**](InlineObject46.md) |  | 

### Return type

[**[]MicrosoftGraphExtensionProperty**](MicrosoftGraphExtensionProperty.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsGetByIds

> []MicrosoftGraphDirectoryObject ContactsGetByIds(ctx).InlineObject47(inlineObject47).Execute()

Invoke action getByIds

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject47 := *openapiclient.NewInlineObject47() // InlineObject47 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsGetByIds(context.Background()).InlineObject47(inlineObject47).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsActionsApi.ContactsGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject47** | [**InlineObject47**](InlineObject47.md) |  | 

### Return type

[**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactCheckMemberGroups

> []string ContactsOrgContactCheckMemberGroups(ctx, orgContactId).InlineObject42(inlineObject42).Execute()

Invoke action checkMemberGroups

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgContactId := "orgContactId_example" // string | key: id of orgContact
    inlineObject42 := *openapiclient.NewInlineObject42() // InlineObject42 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsOrgContactCheckMemberGroups(context.Background(), orgContactId).InlineObject42(inlineObject42).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsOrgContactCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContactsActionsApi.ContactsOrgContactCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject42** | [**InlineObject42**](InlineObject42.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactCheckMemberObjects

> []string ContactsOrgContactCheckMemberObjects(ctx, orgContactId).InlineObject43(inlineObject43).Execute()

Invoke action checkMemberObjects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgContactId := "orgContactId_example" // string | key: id of orgContact
    inlineObject43 := *openapiclient.NewInlineObject43() // InlineObject43 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsOrgContactCheckMemberObjects(context.Background(), orgContactId).InlineObject43(inlineObject43).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsOrgContactCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContactsActionsApi.ContactsOrgContactCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject43** | [**InlineObject43**](InlineObject43.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactGetMemberGroups

> []string ContactsOrgContactGetMemberGroups(ctx, orgContactId).InlineObject44(inlineObject44).Execute()

Invoke action getMemberGroups

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgContactId := "orgContactId_example" // string | key: id of orgContact
    inlineObject44 := *openapiclient.NewInlineObject44() // InlineObject44 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsOrgContactGetMemberGroups(context.Background(), orgContactId).InlineObject44(inlineObject44).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsOrgContactGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContactsActionsApi.ContactsOrgContactGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject44** | [**InlineObject44**](InlineObject44.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactGetMemberObjects

> []string ContactsOrgContactGetMemberObjects(ctx, orgContactId).InlineObject45(inlineObject45).Execute()

Invoke action getMemberObjects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgContactId := "orgContactId_example" // string | key: id of orgContact
    inlineObject45 := *openapiclient.NewInlineObject45() // InlineObject45 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsOrgContactGetMemberObjects(context.Background(), orgContactId).InlineObject45(inlineObject45).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsOrgContactGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `ContactsActionsApi.ContactsOrgContactGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject45** | [**InlineObject45**](InlineObject45.md) |  | 

### Return type

**[]string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsOrgContactRestore

> AnyOfmicrosoftGraphDirectoryObject ContactsOrgContactRestore(ctx, orgContactId).Execute()

Invoke action restore

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgContactId := "orgContactId_example" // string | key: id of orgContact

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsOrgContactRestore(context.Background(), orgContactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsOrgContactRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContactsOrgContactRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `ContactsActionsApi.ContactsOrgContactRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgContactId** | **string** | key: id of orgContact | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactsOrgContactRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsValidateProperties

> ContactsValidateProperties(ctx).InlineObject48(inlineObject48).Execute()

Invoke action validateProperties

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject48 := *openapiclient.NewInlineObject48() // InlineObject48 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContactsActionsApi.ContactsValidateProperties(context.Background()).InlineObject48(inlineObject48).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsActionsApi.ContactsValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject48** | [**InlineObject48**](InlineObject48.md) |  | 

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

