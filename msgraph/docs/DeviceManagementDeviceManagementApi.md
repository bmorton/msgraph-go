# \DeviceManagementDeviceManagementApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeviceManagementGetDeviceManagement**](DeviceManagementDeviceManagementApi.md#DeviceManagementDeviceManagementGetDeviceManagement) | **Get** /deviceManagement | Get deviceManagement
[**DeviceManagementDeviceManagementUpdateDeviceManagement**](DeviceManagementDeviceManagementApi.md#DeviceManagementDeviceManagementUpdateDeviceManagement) | **Patch** /deviceManagement | Update deviceManagement



## DeviceManagementDeviceManagementGetDeviceManagement

> MicrosoftGraphDeviceManagement DeviceManagementDeviceManagementGetDeviceManagement(ctx).Select_(select_).Expand(expand).Execute()

Get deviceManagement

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementApi.DeviceManagementDeviceManagementGetDeviceManagement(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementApi.DeviceManagementDeviceManagementGetDeviceManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementDeviceManagementGetDeviceManagement`: MicrosoftGraphDeviceManagement
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementApi.DeviceManagementDeviceManagementGetDeviceManagement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceManagementGetDeviceManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagement**](MicrosoftGraphDeviceManagement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeviceManagementUpdateDeviceManagement

> DeviceManagementDeviceManagementUpdateDeviceManagement(ctx).MicrosoftGraphDeviceManagement(microsoftGraphDeviceManagement).Execute()

Update deviceManagement

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
    microsoftGraphDeviceManagement := *openapiclient.NewMicrosoftGraphDeviceManagement() // MicrosoftGraphDeviceManagement | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementApi.DeviceManagementDeviceManagementUpdateDeviceManagement(context.Background()).MicrosoftGraphDeviceManagement(microsoftGraphDeviceManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementApi.DeviceManagementDeviceManagementUpdateDeviceManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeviceManagementUpdateDeviceManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceManagement** | [**MicrosoftGraphDeviceManagement**](MicrosoftGraphDeviceManagement.md) | New property values | 

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

