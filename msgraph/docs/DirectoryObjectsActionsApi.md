# \DirectoryObjectsActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryObjectsDirectoryObjectCheckMemberGroups**](DirectoryObjectsActionsApi.md#DirectoryObjectsDirectoryObjectCheckMemberGroups) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**DirectoryObjectsDirectoryObjectCheckMemberObjects**](DirectoryObjectsActionsApi.md#DirectoryObjectsDirectoryObjectCheckMemberObjects) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**DirectoryObjectsDirectoryObjectGetMemberGroups**](DirectoryObjectsActionsApi.md#DirectoryObjectsDirectoryObjectGetMemberGroups) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**DirectoryObjectsDirectoryObjectGetMemberObjects**](DirectoryObjectsActionsApi.md#DirectoryObjectsDirectoryObjectGetMemberObjects) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**DirectoryObjectsDirectoryObjectRestore**](DirectoryObjectsActionsApi.md#DirectoryObjectsDirectoryObjectRestore) | **Post** /directoryObjects/{directoryObject-id}/microsoft.graph.restore | Invoke action restore
[**DirectoryObjectsGetAvailableExtensionProperties**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetAvailableExtensionProperties) | **Post** /directoryObjects/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**DirectoryObjectsGetByIds**](DirectoryObjectsActionsApi.md#DirectoryObjectsGetByIds) | **Post** /directoryObjects/microsoft.graph.getByIds | Invoke action getByIds
[**DirectoryObjectsValidateProperties**](DirectoryObjectsActionsApi.md#DirectoryObjectsValidateProperties) | **Post** /directoryObjects/microsoft.graph.validateProperties | Invoke action validateProperties



## DirectoryObjectsDirectoryObjectCheckMemberGroups

> []string DirectoryObjectsDirectoryObjectCheckMemberGroups(ctx, directoryObjectId).InlineObject113(inlineObject113).Execute()

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    inlineObject113 := *openapiclient.NewInlineObject113() // InlineObject113 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectCheckMemberGroups(context.Background(), directoryObjectId).InlineObject113(inlineObject113).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject113** | [**InlineObject113**](InlineObject113.md) |  | 

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


## DirectoryObjectsDirectoryObjectCheckMemberObjects

> []string DirectoryObjectsDirectoryObjectCheckMemberObjects(ctx, directoryObjectId).InlineObject114(inlineObject114).Execute()

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    inlineObject114 := *openapiclient.NewInlineObject114() // InlineObject114 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectCheckMemberObjects(context.Background(), directoryObjectId).InlineObject114(inlineObject114).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject114** | [**InlineObject114**](InlineObject114.md) |  | 

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


## DirectoryObjectsDirectoryObjectGetMemberGroups

> []string DirectoryObjectsDirectoryObjectGetMemberGroups(ctx, directoryObjectId).InlineObject115(inlineObject115).Execute()

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    inlineObject115 := *openapiclient.NewInlineObject115() // InlineObject115 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectGetMemberGroups(context.Background(), directoryObjectId).InlineObject115(inlineObject115).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject115** | [**InlineObject115**](InlineObject115.md) |  | 

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


## DirectoryObjectsDirectoryObjectGetMemberObjects

> []string DirectoryObjectsDirectoryObjectGetMemberObjects(ctx, directoryObjectId).InlineObject116(inlineObject116).Execute()

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    inlineObject116 := *openapiclient.NewInlineObject116() // InlineObject116 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectGetMemberObjects(context.Background(), directoryObjectId).InlineObject116(inlineObject116).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject116** | [**InlineObject116**](InlineObject116.md) |  | 

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


## DirectoryObjectsDirectoryObjectRestore

> AnyOfmicrosoftGraphDirectoryObject DirectoryObjectsDirectoryObjectRestore(ctx, directoryObjectId).Execute()

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectRestore(context.Background(), directoryObjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsActionsApi.DirectoryObjectsDirectoryObjectRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectRestoreRequest struct via the builder pattern


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


## DirectoryObjectsGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty DirectoryObjectsGetAvailableExtensionProperties(ctx).InlineObject117(inlineObject117).Execute()

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
    inlineObject117 := *openapiclient.NewInlineObject117() // InlineObject117 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsGetAvailableExtensionProperties(context.Background()).InlineObject117(inlineObject117).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsActionsApi.DirectoryObjectsGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject117** | [**InlineObject117**](InlineObject117.md) |  | 

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


## DirectoryObjectsGetByIds

> []MicrosoftGraphDirectoryObject DirectoryObjectsGetByIds(ctx).InlineObject118(inlineObject118).Execute()

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
    inlineObject118 := *openapiclient.NewInlineObject118() // InlineObject118 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsGetByIds(context.Background()).InlineObject118(inlineObject118).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsActionsApi.DirectoryObjectsGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject118** | [**InlineObject118**](InlineObject118.md) |  | 

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


## DirectoryObjectsValidateProperties

> DirectoryObjectsValidateProperties(ctx).InlineObject119(inlineObject119).Execute()

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
    inlineObject119 := *openapiclient.NewInlineObject119() // InlineObject119 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsActionsApi.DirectoryObjectsValidateProperties(context.Background()).InlineObject119(inlineObject119).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsActionsApi.DirectoryObjectsValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject119** | [**InlineObject119**](InlineObject119.md) |  | 

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

