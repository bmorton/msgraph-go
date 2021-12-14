# \PermissionGrantsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionGrantsGetAvailableExtensionProperties**](PermissionGrantsActionsApi.md#PermissionGrantsGetAvailableExtensionProperties) | **Post** /permissionGrants/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**PermissionGrantsGetByIds**](PermissionGrantsActionsApi.md#PermissionGrantsGetByIds) | **Post** /permissionGrants/microsoft.graph.getByIds | Invoke action getByIds
[**PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups**](PermissionGrantsActionsApi.md#PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups) | **Post** /permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects**](PermissionGrantsActionsApi.md#PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects) | **Post** /permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups**](PermissionGrantsActionsApi.md#PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups) | **Post** /permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects**](PermissionGrantsActionsApi.md#PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects) | **Post** /permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**PermissionGrantsResourceSpecificPermissionGrantRestore**](PermissionGrantsActionsApi.md#PermissionGrantsResourceSpecificPermissionGrantRestore) | **Post** /permissionGrants/{resourceSpecificPermissionGrant-id}/microsoft.graph.restore | Invoke action restore
[**PermissionGrantsValidateProperties**](PermissionGrantsActionsApi.md#PermissionGrantsValidateProperties) | **Post** /permissionGrants/microsoft.graph.validateProperties | Invoke action validateProperties



## PermissionGrantsGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty PermissionGrantsGetAvailableExtensionProperties(ctx).InlineObject710(inlineObject710).Execute()

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
    inlineObject710 := *openapiclient.NewInlineObject710() // InlineObject710 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsGetAvailableExtensionProperties(context.Background()).InlineObject710(inlineObject710).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsActionsApi.PermissionGrantsGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject710** | [**InlineObject710**](InlineObject710.md) |  | 

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


## PermissionGrantsGetByIds

> []MicrosoftGraphDirectoryObject PermissionGrantsGetByIds(ctx).InlineObject711(inlineObject711).Execute()

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
    inlineObject711 := *openapiclient.NewInlineObject711() // InlineObject711 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsGetByIds(context.Background()).InlineObject711(inlineObject711).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsActionsApi.PermissionGrantsGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject711** | [**InlineObject711**](InlineObject711.md) |  | 

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


## PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups

> []string PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups(ctx, resourceSpecificPermissionGrantId).InlineObject706(inlineObject706).Execute()

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    inlineObject706 := *openapiclient.NewInlineObject706() // InlineObject706 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups(context.Background(), resourceSpecificPermissionGrantId).InlineObject706(inlineObject706).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject706** | [**InlineObject706**](InlineObject706.md) |  | 

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


## PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects

> []string PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects(ctx, resourceSpecificPermissionGrantId).InlineObject707(inlineObject707).Execute()

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    inlineObject707 := *openapiclient.NewInlineObject707() // InlineObject707 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects(context.Background(), resourceSpecificPermissionGrantId).InlineObject707(inlineObject707).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject707** | [**InlineObject707**](InlineObject707.md) |  | 

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


## PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups

> []string PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups(ctx, resourceSpecificPermissionGrantId).InlineObject708(inlineObject708).Execute()

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    inlineObject708 := *openapiclient.NewInlineObject708() // InlineObject708 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups(context.Background(), resourceSpecificPermissionGrantId).InlineObject708(inlineObject708).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject708** | [**InlineObject708**](InlineObject708.md) |  | 

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


## PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects

> []string PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects(ctx, resourceSpecificPermissionGrantId).InlineObject709(inlineObject709).Execute()

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant
    inlineObject709 := *openapiclient.NewInlineObject709() // InlineObject709 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects(context.Background(), resourceSpecificPermissionGrantId).InlineObject709(inlineObject709).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject709** | [**InlineObject709**](InlineObject709.md) |  | 

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


## PermissionGrantsResourceSpecificPermissionGrantRestore

> AnyOfmicrosoftGraphDirectoryObject PermissionGrantsResourceSpecificPermissionGrantRestore(ctx, resourceSpecificPermissionGrantId).Execute()

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
    resourceSpecificPermissionGrantId := "resourceSpecificPermissionGrantId_example" // string | key: id of resourceSpecificPermissionGrant

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantRestore(context.Background(), resourceSpecificPermissionGrantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionGrantsResourceSpecificPermissionGrantRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `PermissionGrantsActionsApi.PermissionGrantsResourceSpecificPermissionGrantRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSpecificPermissionGrantId** | **string** | key: id of resourceSpecificPermissionGrant | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsResourceSpecificPermissionGrantRestoreRequest struct via the builder pattern


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


## PermissionGrantsValidateProperties

> PermissionGrantsValidateProperties(ctx).InlineObject712(inlineObject712).Execute()

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
    inlineObject712 := *openapiclient.NewInlineObject712() // InlineObject712 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionGrantsActionsApi.PermissionGrantsValidateProperties(context.Background()).InlineObject712(inlineObject712).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGrantsActionsApi.PermissionGrantsValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionGrantsValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject712** | [**InlineObject712**](InlineObject712.md) |  | 

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

