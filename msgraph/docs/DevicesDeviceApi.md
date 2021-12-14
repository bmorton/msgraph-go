# \DevicesDeviceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesDeviceCreateDevice**](DevicesDeviceApi.md#DevicesDeviceCreateDevice) | **Post** /devices | Add new entity to devices
[**DevicesDeviceDeleteDevice**](DevicesDeviceApi.md#DevicesDeviceDeleteDevice) | **Delete** /devices/{device-id} | Delete entity from devices
[**DevicesDeviceGetDevice**](DevicesDeviceApi.md#DevicesDeviceGetDevice) | **Get** /devices/{device-id} | Get entity from devices by key
[**DevicesDeviceListDevice**](DevicesDeviceApi.md#DevicesDeviceListDevice) | **Get** /devices | Get entities from devices
[**DevicesDeviceUpdateDevice**](DevicesDeviceApi.md#DevicesDeviceUpdateDevice) | **Patch** /devices/{device-id} | Update entity in devices



## DevicesDeviceCreateDevice

> MicrosoftGraphDevice DevicesDeviceCreateDevice(ctx).MicrosoftGraphDevice(microsoftGraphDevice).Execute()

Add new entity to devices

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
    microsoftGraphDevice := *openapiclient.NewMicrosoftGraphDevice() // MicrosoftGraphDevice | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesDeviceApi.DevicesDeviceCreateDevice(context.Background()).MicrosoftGraphDevice(microsoftGraphDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesDeviceApi.DevicesDeviceCreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceCreateDevice`: MicrosoftGraphDevice
    fmt.Fprintf(os.Stdout, "Response from `DevicesDeviceApi.DevicesDeviceCreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDevice** | [**MicrosoftGraphDevice**](MicrosoftGraphDevice.md) | New entity | 

### Return type

[**MicrosoftGraphDevice**](MicrosoftGraphDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceDeleteDevice

> DevicesDeviceDeleteDevice(ctx, deviceId).IfMatch(ifMatch).Execute()

Delete entity from devices

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesDeviceApi.DevicesDeviceDeleteDevice(context.Background(), deviceId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesDeviceApi.DevicesDeviceDeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceDeleteDeviceRequest struct via the builder pattern


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


## DevicesDeviceGetDevice

> MicrosoftGraphDevice DevicesDeviceGetDevice(ctx, deviceId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()

Get entity from devices by key

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
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesDeviceApi.DevicesDeviceGetDevice(context.Background(), deviceId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesDeviceApi.DevicesDeviceGetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceGetDevice`: MicrosoftGraphDevice
    fmt.Fprintf(os.Stdout, "Response from `DevicesDeviceApi.DevicesDeviceGetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDevice**](MicrosoftGraphDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceListDevice

> CollectionOfDevice DevicesDeviceListDevice(ctx).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from devices

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
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
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
    resp, r, err := api_client.DevicesDeviceApi.DevicesDeviceListDevice(context.Background()).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesDeviceApi.DevicesDeviceListDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DevicesDeviceListDevice`: CollectionOfDevice
    fmt.Fprintf(os.Stdout, "Response from `DevicesDeviceApi.DevicesDeviceListDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceListDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDevice**](CollectionOfDevice.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceUpdateDevice

> DevicesDeviceUpdateDevice(ctx, deviceId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()

Update entity in devices

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
    microsoftGraphDevice := *openapiclient.NewMicrosoftGraphDevice() // MicrosoftGraphDevice | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DevicesDeviceApi.DevicesDeviceUpdateDevice(context.Background(), deviceId).MicrosoftGraphDevice(microsoftGraphDevice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DevicesDeviceApi.DevicesDeviceUpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | key: id of device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDevice** | [**MicrosoftGraphDevice**](MicrosoftGraphDevice.md) | New property values | 

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

