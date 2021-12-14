# \DirectoryRolesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryRolesDirectoryRoleCheckMemberGroups**](DirectoryRolesActionsApi.md#DirectoryRolesDirectoryRoleCheckMemberGroups) | **Post** /directoryRoles/{directoryRole-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**DirectoryRolesDirectoryRoleCheckMemberObjects**](DirectoryRolesActionsApi.md#DirectoryRolesDirectoryRoleCheckMemberObjects) | **Post** /directoryRoles/{directoryRole-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**DirectoryRolesDirectoryRoleGetMemberGroups**](DirectoryRolesActionsApi.md#DirectoryRolesDirectoryRoleGetMemberGroups) | **Post** /directoryRoles/{directoryRole-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**DirectoryRolesDirectoryRoleGetMemberObjects**](DirectoryRolesActionsApi.md#DirectoryRolesDirectoryRoleGetMemberObjects) | **Post** /directoryRoles/{directoryRole-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**DirectoryRolesDirectoryRoleRestore**](DirectoryRolesActionsApi.md#DirectoryRolesDirectoryRoleRestore) | **Post** /directoryRoles/{directoryRole-id}/microsoft.graph.restore | Invoke action restore
[**DirectoryRolesGetAvailableExtensionProperties**](DirectoryRolesActionsApi.md#DirectoryRolesGetAvailableExtensionProperties) | **Post** /directoryRoles/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**DirectoryRolesGetByIds**](DirectoryRolesActionsApi.md#DirectoryRolesGetByIds) | **Post** /directoryRoles/microsoft.graph.getByIds | Invoke action getByIds
[**DirectoryRolesValidateProperties**](DirectoryRolesActionsApi.md#DirectoryRolesValidateProperties) | **Post** /directoryRoles/microsoft.graph.validateProperties | Invoke action validateProperties



## DirectoryRolesDirectoryRoleCheckMemberGroups

> []string DirectoryRolesDirectoryRoleCheckMemberGroups(ctx, directoryRoleId).InlineObject120(inlineObject120).Execute()

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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole
    inlineObject120 := *openapiclient.NewInlineObject120() // InlineObject120 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleCheckMemberGroups(context.Background(), directoryRoleId).InlineObject120(inlineObject120).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject120** | [**InlineObject120**](InlineObject120.md) |  | 

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


## DirectoryRolesDirectoryRoleCheckMemberObjects

> []string DirectoryRolesDirectoryRoleCheckMemberObjects(ctx, directoryRoleId).InlineObject121(inlineObject121).Execute()

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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole
    inlineObject121 := *openapiclient.NewInlineObject121() // InlineObject121 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleCheckMemberObjects(context.Background(), directoryRoleId).InlineObject121(inlineObject121).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject121** | [**InlineObject121**](InlineObject121.md) |  | 

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


## DirectoryRolesDirectoryRoleGetMemberGroups

> []string DirectoryRolesDirectoryRoleGetMemberGroups(ctx, directoryRoleId).InlineObject122(inlineObject122).Execute()

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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole
    inlineObject122 := *openapiclient.NewInlineObject122() // InlineObject122 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleGetMemberGroups(context.Background(), directoryRoleId).InlineObject122(inlineObject122).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject122** | [**InlineObject122**](InlineObject122.md) |  | 

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


## DirectoryRolesDirectoryRoleGetMemberObjects

> []string DirectoryRolesDirectoryRoleGetMemberObjects(ctx, directoryRoleId).InlineObject123(inlineObject123).Execute()

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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole
    inlineObject123 := *openapiclient.NewInlineObject123() // InlineObject123 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleGetMemberObjects(context.Background(), directoryRoleId).InlineObject123(inlineObject123).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject123** | [**InlineObject123**](InlineObject123.md) |  | 

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


## DirectoryRolesDirectoryRoleRestore

> AnyOfmicrosoftGraphDirectoryObject DirectoryRolesDirectoryRoleRestore(ctx, directoryRoleId).Execute()

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
    directoryRoleId := "directoryRoleId_example" // string | key: id of directoryRole

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleRestore(context.Background(), directoryRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesDirectoryRoleRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesActionsApi.DirectoryRolesDirectoryRoleRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryRoleId** | **string** | key: id of directoryRole | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesDirectoryRoleRestoreRequest struct via the builder pattern


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


## DirectoryRolesGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty DirectoryRolesGetAvailableExtensionProperties(ctx).InlineObject124(inlineObject124).Execute()

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
    inlineObject124 := *openapiclient.NewInlineObject124() // InlineObject124 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesGetAvailableExtensionProperties(context.Background()).InlineObject124(inlineObject124).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesActionsApi.DirectoryRolesGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject124** | [**InlineObject124**](InlineObject124.md) |  | 

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


## DirectoryRolesGetByIds

> []MicrosoftGraphDirectoryObject DirectoryRolesGetByIds(ctx).InlineObject125(inlineObject125).Execute()

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
    inlineObject125 := *openapiclient.NewInlineObject125() // InlineObject125 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesGetByIds(context.Background()).InlineObject125(inlineObject125).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryRolesGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryRolesActionsApi.DirectoryRolesGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject125** | [**InlineObject125**](InlineObject125.md) |  | 

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


## DirectoryRolesValidateProperties

> DirectoryRolesValidateProperties(ctx).InlineObject126(inlineObject126).Execute()

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
    inlineObject126 := *openapiclient.NewInlineObject126() // InlineObject126 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryRolesActionsApi.DirectoryRolesValidateProperties(context.Background()).InlineObject126(inlineObject126).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryRolesActionsApi.DirectoryRolesValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryRolesValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject126** | [**InlineObject126**](InlineObject126.md) |  | 

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

