# \DeviceManagementSoftwareUpdateStatusSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeleteRefSoftwareUpdateStatusSummary**](DeviceManagementSoftwareUpdateStatusSummaryApi.md#DeviceManagementDeleteRefSoftwareUpdateStatusSummary) | **Delete** /deviceManagement/softwareUpdateStatusSummary/$ref | Delete ref of navigation property softwareUpdateStatusSummary for deviceManagement
[**DeviceManagementGetRefSoftwareUpdateStatusSummary**](DeviceManagementSoftwareUpdateStatusSummaryApi.md#DeviceManagementGetRefSoftwareUpdateStatusSummary) | **Get** /deviceManagement/softwareUpdateStatusSummary/$ref | Get ref of softwareUpdateStatusSummary from deviceManagement
[**DeviceManagementGetSoftwareUpdateStatusSummary**](DeviceManagementSoftwareUpdateStatusSummaryApi.md#DeviceManagementGetSoftwareUpdateStatusSummary) | **Get** /deviceManagement/softwareUpdateStatusSummary | Get softwareUpdateStatusSummary from deviceManagement
[**DeviceManagementUpdateRefSoftwareUpdateStatusSummary**](DeviceManagementSoftwareUpdateStatusSummaryApi.md#DeviceManagementUpdateRefSoftwareUpdateStatusSummary) | **Put** /deviceManagement/softwareUpdateStatusSummary/$ref | Update the ref of navigation property softwareUpdateStatusSummary in deviceManagement



## DeviceManagementDeleteRefSoftwareUpdateStatusSummary

> DeviceManagementDeleteRefSoftwareUpdateStatusSummary(ctx).IfMatch(ifMatch).Execute()

Delete ref of navigation property softwareUpdateStatusSummary for deviceManagement



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
    resp, r, err := api_client.DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementDeleteRefSoftwareUpdateStatusSummary(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementDeleteRefSoftwareUpdateStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteRefSoftwareUpdateStatusSummaryRequest struct via the builder pattern


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


## DeviceManagementGetRefSoftwareUpdateStatusSummary

> string DeviceManagementGetRefSoftwareUpdateStatusSummary(ctx).Execute()

Get ref of softwareUpdateStatusSummary from deviceManagement



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementGetRefSoftwareUpdateStatusSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementGetRefSoftwareUpdateStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetRefSoftwareUpdateStatusSummary`: string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementGetRefSoftwareUpdateStatusSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetRefSoftwareUpdateStatusSummaryRequest struct via the builder pattern


### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetSoftwareUpdateStatusSummary

> MicrosoftGraphSoftwareUpdateStatusSummary DeviceManagementGetSoftwareUpdateStatusSummary(ctx).Select_(select_).Expand(expand).Execute()

Get softwareUpdateStatusSummary from deviceManagement



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
    resp, r, err := api_client.DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementGetSoftwareUpdateStatusSummary(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementGetSoftwareUpdateStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetSoftwareUpdateStatusSummary`: MicrosoftGraphSoftwareUpdateStatusSummary
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementGetSoftwareUpdateStatusSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetSoftwareUpdateStatusSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSoftwareUpdateStatusSummary**](MicrosoftGraphSoftwareUpdateStatusSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateRefSoftwareUpdateStatusSummary

> DeviceManagementUpdateRefSoftwareUpdateStatusSummary(ctx).RequestBody(requestBody).Execute()

Update the ref of navigation property softwareUpdateStatusSummary in deviceManagement



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementUpdateRefSoftwareUpdateStatusSummary(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementSoftwareUpdateStatusSummaryApi.DeviceManagementUpdateRefSoftwareUpdateStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateRefSoftwareUpdateStatusSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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

