# \DeviceAppManagementDeviceAppManagementApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAppManagementDeviceAppManagementGetDeviceAppManagement**](DeviceAppManagementDeviceAppManagementApi.md#DeviceAppManagementDeviceAppManagementGetDeviceAppManagement) | **Get** /deviceAppManagement | Get deviceAppManagement
[**DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement**](DeviceAppManagementDeviceAppManagementApi.md#DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement) | **Patch** /deviceAppManagement | Update deviceAppManagement



## DeviceAppManagementDeviceAppManagementGetDeviceAppManagement

> MicrosoftGraphDeviceAppManagement DeviceAppManagementDeviceAppManagementGetDeviceAppManagement(ctx).Select_(select_).Expand(expand).Execute()

Get deviceAppManagement

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
    resp, r, err := api_client.DeviceAppManagementDeviceAppManagementApi.DeviceAppManagementDeviceAppManagementGetDeviceAppManagement(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDeviceAppManagementApi.DeviceAppManagementDeviceAppManagementGetDeviceAppManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAppManagementDeviceAppManagementGetDeviceAppManagement`: MicrosoftGraphDeviceAppManagement
    fmt.Fprintf(os.Stdout, "Response from `DeviceAppManagementDeviceAppManagementApi.DeviceAppManagementDeviceAppManagementGetDeviceAppManagement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeviceAppManagementGetDeviceAppManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceAppManagement**](MicrosoftGraphDeviceAppManagement.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement

> DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement(ctx).MicrosoftGraphDeviceAppManagement(microsoftGraphDeviceAppManagement).Execute()

Update deviceAppManagement

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
    microsoftGraphDeviceAppManagement := *openapiclient.NewMicrosoftGraphDeviceAppManagement() // MicrosoftGraphDeviceAppManagement | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceAppManagementDeviceAppManagementApi.DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement(context.Background()).MicrosoftGraphDeviceAppManagement(microsoftGraphDeviceAppManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAppManagementDeviceAppManagementApi.DeviceAppManagementDeviceAppManagementUpdateDeviceAppManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAppManagementDeviceAppManagementUpdateDeviceAppManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceAppManagement** | [**MicrosoftGraphDeviceAppManagement**](MicrosoftGraphDeviceAppManagement.md) | New property values | 

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

