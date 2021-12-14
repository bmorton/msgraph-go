# \DeviceManagementDeviceConfigurationDeviceStateSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeleteDeviceConfigurationDeviceStateSummaries**](DeviceManagementDeviceConfigurationDeviceStateSummaryApi.md#DeviceManagementDeleteDeviceConfigurationDeviceStateSummaries) | **Delete** /deviceManagement/deviceConfigurationDeviceStateSummaries | Delete navigation property deviceConfigurationDeviceStateSummaries for deviceManagement
[**DeviceManagementGetDeviceConfigurationDeviceStateSummaries**](DeviceManagementDeviceConfigurationDeviceStateSummaryApi.md#DeviceManagementGetDeviceConfigurationDeviceStateSummaries) | **Get** /deviceManagement/deviceConfigurationDeviceStateSummaries | Get deviceConfigurationDeviceStateSummaries from deviceManagement
[**DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries**](DeviceManagementDeviceConfigurationDeviceStateSummaryApi.md#DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries) | **Patch** /deviceManagement/deviceConfigurationDeviceStateSummaries | Update the navigation property deviceConfigurationDeviceStateSummaries in deviceManagement



## DeviceManagementDeleteDeviceConfigurationDeviceStateSummaries

> DeviceManagementDeleteDeviceConfigurationDeviceStateSummaries(ctx).IfMatch(ifMatch).Execute()

Delete navigation property deviceConfigurationDeviceStateSummaries for deviceManagement



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceConfigurationDeviceStateSummaryApi.DeviceManagementDeleteDeviceConfigurationDeviceStateSummaries(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceConfigurationDeviceStateSummaryApi.DeviceManagementDeleteDeviceConfigurationDeviceStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteDeviceConfigurationDeviceStateSummariesRequest struct via the builder pattern


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


## DeviceManagementGetDeviceConfigurationDeviceStateSummaries

> MicrosoftGraphDeviceConfigurationDeviceStateSummary DeviceManagementGetDeviceConfigurationDeviceStateSummaries(ctx).Select_(select_).Expand(expand).Execute()

Get deviceConfigurationDeviceStateSummaries from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceConfigurationDeviceStateSummaryApi.DeviceManagementGetDeviceConfigurationDeviceStateSummaries(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceConfigurationDeviceStateSummaryApi.DeviceManagementGetDeviceConfigurationDeviceStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetDeviceConfigurationDeviceStateSummaries`: MicrosoftGraphDeviceConfigurationDeviceStateSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceConfigurationDeviceStateSummaryApi.DeviceManagementGetDeviceConfigurationDeviceStateSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetDeviceConfigurationDeviceStateSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceConfigurationDeviceStateSummary**](MicrosoftGraphDeviceConfigurationDeviceStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries

> DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries(ctx).MicrosoftGraphDeviceConfigurationDeviceStateSummary(microsoftGraphDeviceConfigurationDeviceStateSummary).Execute()

Update the navigation property deviceConfigurationDeviceStateSummaries in deviceManagement



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
    microsoftGraphDeviceConfigurationDeviceStateSummary := *openapiclient.NewMicrosoftGraphDeviceConfigurationDeviceStateSummary() // MicrosoftGraphDeviceConfigurationDeviceStateSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceConfigurationDeviceStateSummaryApi.DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries(context.Background()).MicrosoftGraphDeviceConfigurationDeviceStateSummary(microsoftGraphDeviceConfigurationDeviceStateSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceConfigurationDeviceStateSummaryApi.DeviceManagementUpdateDeviceConfigurationDeviceStateSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateDeviceConfigurationDeviceStateSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceConfigurationDeviceStateSummary** | [**MicrosoftGraphDeviceConfigurationDeviceStateSummary**](MicrosoftGraphDeviceConfigurationDeviceStateSummary.md) | New navigation property values | 

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

