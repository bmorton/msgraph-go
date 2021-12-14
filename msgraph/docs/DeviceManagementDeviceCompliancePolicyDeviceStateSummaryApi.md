# \DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary**](DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.md#DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary) | **Delete** /deviceManagement/deviceCompliancePolicyDeviceStateSummary | Delete navigation property deviceCompliancePolicyDeviceStateSummary for deviceManagement
[**DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary**](DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.md#DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary) | **Get** /deviceManagement/deviceCompliancePolicyDeviceStateSummary | Get deviceCompliancePolicyDeviceStateSummary from deviceManagement
[**DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary**](DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.md#DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary) | **Patch** /deviceManagement/deviceCompliancePolicyDeviceStateSummary | Update the navigation property deviceCompliancePolicyDeviceStateSummary in deviceManagement



## DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary

> DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary(ctx).IfMatch(ifMatch).Execute()

Delete navigation property deviceCompliancePolicyDeviceStateSummary for deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.DeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteDeviceCompliancePolicyDeviceStateSummaryRequest struct via the builder pattern


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


## DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary

> MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary(ctx).Select_(select_).Expand(expand).Execute()

Get deviceCompliancePolicyDeviceStateSummary from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary`: MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.DeviceManagementGetDeviceCompliancePolicyDeviceStateSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetDeviceCompliancePolicyDeviceStateSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary**](MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary

> DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary(ctx).MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary(microsoftGraphDeviceCompliancePolicyDeviceStateSummary).Execute()

Update the navigation property deviceCompliancePolicyDeviceStateSummary in deviceManagement



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
    microsoftGraphDeviceCompliancePolicyDeviceStateSummary := *openapiclient.NewMicrosoftGraphDeviceCompliancePolicyDeviceStateSummary() // MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary(context.Background()).MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary(microsoftGraphDeviceCompliancePolicyDeviceStateSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceCompliancePolicyDeviceStateSummaryApi.DeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateDeviceCompliancePolicyDeviceStateSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceCompliancePolicyDeviceStateSummary** | [**MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary**](MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary.md) | New navigation property values | 

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

