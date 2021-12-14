# \DevicesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesDeviceCheckMemberGroups**](DevicesActionsApi.md#DevicesDeviceCheckMemberGroups) | **Post** /devices/{device-id}/microsoft.graph.checkMemberGroups | Invoke action checkMemberGroups
[**DevicesDeviceCheckMemberObjects**](DevicesActionsApi.md#DevicesDeviceCheckMemberObjects) | **Post** /devices/{device-id}/microsoft.graph.checkMemberObjects | Invoke action checkMemberObjects
[**DevicesDeviceGetMemberGroups**](DevicesActionsApi.md#DevicesDeviceGetMemberGroups) | **Post** /devices/{device-id}/microsoft.graph.getMemberGroups | Invoke action getMemberGroups
[**DevicesDeviceGetMemberObjects**](DevicesActionsApi.md#DevicesDeviceGetMemberObjects) | **Post** /devices/{device-id}/microsoft.graph.getMemberObjects | Invoke action getMemberObjects
[**DevicesDeviceRestore**](DevicesActionsApi.md#DevicesDeviceRestore) | **Post** /devices/{device-id}/microsoft.graph.restore | Invoke action restore
[**DevicesGetAvailableExtensionProperties**](DevicesActionsApi.md#DevicesGetAvailableExtensionProperties) | **Post** /devices/microsoft.graph.getAvailableExtensionProperties | Invoke action getAvailableExtensionProperties
[**DevicesGetByIds**](DevicesActionsApi.md#DevicesGetByIds) | **Post** /devices/microsoft.graph.getByIds | Invoke action getByIds
[**DevicesValidateProperties**](DevicesActionsApi.md#DevicesValidateProperties) | **Post** /devices/microsoft.graph.validateProperties | Invoke action validateProperties



## DevicesDeviceCheckMemberGroups

> []string DevicesDeviceCheckMemberGroups(ctx, deviceId).InlineObject106(inlineObject106).Execute()

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
    deviceId := "deviceId_example" // string | key: id of device
    inlineObject106 := *openapiclient.NewInlineObject106() // InlineObject106 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesDeviceCheckMemberGroups(context.Background(), deviceId).InlineObject106(inlineObject106).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesDeviceCheckMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceCheckMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DevicesActionsApi.DevicesDeviceCheckMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceCheckMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject106** | [**InlineObject106**](InlineObject106.md) |  | 

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


## DevicesDeviceCheckMemberObjects

> []string DevicesDeviceCheckMemberObjects(ctx, deviceId).InlineObject107(inlineObject107).Execute()

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
    deviceId := "deviceId_example" // string | key: id of device
    inlineObject107 := *openapiclient.NewInlineObject107() // InlineObject107 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesDeviceCheckMemberObjects(context.Background(), deviceId).InlineObject107(inlineObject107).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesDeviceCheckMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceCheckMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DevicesActionsApi.DevicesDeviceCheckMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceCheckMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject107** | [**InlineObject107**](InlineObject107.md) |  | 

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


## DevicesDeviceGetMemberGroups

> []string DevicesDeviceGetMemberGroups(ctx, deviceId).InlineObject108(inlineObject108).Execute()

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
    deviceId := "deviceId_example" // string | key: id of device
    inlineObject108 := *openapiclient.NewInlineObject108() // InlineObject108 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesDeviceGetMemberGroups(context.Background(), deviceId).InlineObject108(inlineObject108).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesDeviceGetMemberGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceGetMemberGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `DevicesActionsApi.DevicesDeviceGetMemberGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceGetMemberGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject108** | [**InlineObject108**](InlineObject108.md) |  | 

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


## DevicesDeviceGetMemberObjects

> []string DevicesDeviceGetMemberObjects(ctx, deviceId).InlineObject109(inlineObject109).Execute()

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
    deviceId := "deviceId_example" // string | key: id of device
    inlineObject109 := *openapiclient.NewInlineObject109() // InlineObject109 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesDeviceGetMemberObjects(context.Background(), deviceId).InlineObject109(inlineObject109).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesDeviceGetMemberObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceGetMemberObjects`: []string
    fmt.Fprintf(os.Stdout, "Response from `DevicesActionsApi.DevicesDeviceGetMemberObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceGetMemberObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject109** | [**InlineObject109**](InlineObject109.md) |  | 

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


## DevicesDeviceRestore

> AnyOfmicrosoftGraphDirectoryObject DevicesDeviceRestore(ctx, deviceId).Execute()

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
    deviceId := "deviceId_example" // string | key: id of device

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesDeviceRestore(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesDeviceRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceRestore`: AnyOfmicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DevicesActionsApi.DevicesDeviceRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceRestoreRequest struct via the builder pattern


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


## DevicesGetAvailableExtensionProperties

> []MicrosoftGraphExtensionProperty DevicesGetAvailableExtensionProperties(ctx).InlineObject110(inlineObject110).Execute()

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
    inlineObject110 := *openapiclient.NewInlineObject110() // InlineObject110 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesGetAvailableExtensionProperties(context.Background()).InlineObject110(inlineObject110).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesGetAvailableExtensionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGetAvailableExtensionProperties`: []MicrosoftGraphExtensionProperty
    fmt.Fprintf(os.Stdout, "Response from `DevicesActionsApi.DevicesGetAvailableExtensionProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetAvailableExtensionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject110** | [**InlineObject110**](InlineObject110.md) |  | 

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


## DevicesGetByIds

> []MicrosoftGraphDirectoryObject DevicesGetByIds(ctx).InlineObject111(inlineObject111).Execute()

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
    inlineObject111 := *openapiclient.NewInlineObject111() // InlineObject111 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesGetByIds(context.Background()).InlineObject111(inlineObject111).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesGetByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGetByIds`: []MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DevicesActionsApi.DevicesGetByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject111** | [**InlineObject111**](InlineObject111.md) |  | 

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


## DevicesValidateProperties

> DevicesValidateProperties(ctx).InlineObject112(inlineObject112).Execute()

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
    inlineObject112 := *openapiclient.NewInlineObject112() // InlineObject112 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesActionsApi.DevicesValidateProperties(context.Background()).InlineObject112(inlineObject112).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesActionsApi.DevicesValidateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesValidatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject112** | [**InlineObject112**](InlineObject112.md) |  | 

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

