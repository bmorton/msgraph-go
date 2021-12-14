# \DevicesExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesCreateExtensions**](DevicesExtensionApi.md#DevicesCreateExtensions) | **Post** /devices/{device-id}/extensions | Create new navigation property to extensions for devices
[**DevicesDeleteExtensions**](DevicesExtensionApi.md#DevicesDeleteExtensions) | **Delete** /devices/{device-id}/extensions/{extension-id} | Delete navigation property extensions for devices
[**DevicesGetExtensions**](DevicesExtensionApi.md#DevicesGetExtensions) | **Get** /devices/{device-id}/extensions/{extension-id} | Get extensions from devices
[**DevicesListExtensions**](DevicesExtensionApi.md#DevicesListExtensions) | **Get** /devices/{device-id}/extensions | Get extensions from devices
[**DevicesUpdateExtensions**](DevicesExtensionApi.md#DevicesUpdateExtensions) | **Patch** /devices/{device-id}/extensions/{extension-id} | Update the navigation property extensions in devices



## DevicesCreateExtensions

> MicrosoftGraphExtension DevicesCreateExtensions(ctx, deviceId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for devices



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
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesExtensionApi.DevicesCreateExtensions(context.Background(), deviceId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesExtensionApi.DevicesCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `DevicesExtensionApi.DevicesCreateExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesCreateExtensionsRequest struct via the builder pattern


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


## DevicesDeleteExtensions

> DevicesDeleteExtensions(ctx, deviceId, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for devices



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
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesExtensionApi.DevicesDeleteExtensions(context.Background(), deviceId, extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesExtensionApi.DevicesDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeleteExtensionsRequest struct via the builder pattern


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


## DevicesGetExtensions

> MicrosoftGraphExtension DevicesGetExtensions(ctx, deviceId, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from devices



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
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesExtensionApi.DevicesGetExtensions(context.Background(), deviceId, extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesExtensionApi.DevicesGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `DevicesExtensionApi.DevicesGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesGetExtensionsRequest struct via the builder pattern


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


## DevicesListExtensions

> CollectionOfExtension DevicesListExtensions(ctx, deviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from devices



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
    resp, r, err := api_client.DevicesExtensionApi.DevicesListExtensions(context.Background(), deviceId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesExtensionApi.DevicesListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `DevicesExtensionApi.DevicesListExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesListExtensionsRequest struct via the builder pattern


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

[**CollectionOfExtension**](CollectionOfExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesUpdateExtensions

> DevicesUpdateExtensions(ctx, deviceId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in devices



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
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesExtensionApi.DevicesUpdateExtensions(context.Background(), deviceId, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesExtensionApi.DevicesUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesUpdateExtensionsRequest struct via the builder pattern


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

