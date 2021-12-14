# \DeviceManagementManagedDeviceOverviewApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeleteRefManagedDeviceOverview**](DeviceManagementManagedDeviceOverviewApi.md#DeviceManagementDeleteRefManagedDeviceOverview) | **Delete** /deviceManagement/managedDeviceOverview/$ref | Delete ref of navigation property managedDeviceOverview for deviceManagement
[**DeviceManagementGetManagedDeviceOverview**](DeviceManagementManagedDeviceOverviewApi.md#DeviceManagementGetManagedDeviceOverview) | **Get** /deviceManagement/managedDeviceOverview | Get managedDeviceOverview from deviceManagement
[**DeviceManagementGetRefManagedDeviceOverview**](DeviceManagementManagedDeviceOverviewApi.md#DeviceManagementGetRefManagedDeviceOverview) | **Get** /deviceManagement/managedDeviceOverview/$ref | Get ref of managedDeviceOverview from deviceManagement
[**DeviceManagementUpdateRefManagedDeviceOverview**](DeviceManagementManagedDeviceOverviewApi.md#DeviceManagementUpdateRefManagedDeviceOverview) | **Put** /deviceManagement/managedDeviceOverview/$ref | Update the ref of navigation property managedDeviceOverview in deviceManagement



## DeviceManagementDeleteRefManagedDeviceOverview

> DeviceManagementDeleteRefManagedDeviceOverview(ctx).IfMatch(ifMatch).Execute()

Delete ref of navigation property managedDeviceOverview for deviceManagement



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
    resp, r, err := api_client.DeviceManagementManagedDeviceOverviewApi.DeviceManagementDeleteRefManagedDeviceOverview(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceOverviewApi.DeviceManagementDeleteRefManagedDeviceOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteRefManagedDeviceOverviewRequest struct via the builder pattern


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


## DeviceManagementGetManagedDeviceOverview

> MicrosoftGraphManagedDeviceOverview DeviceManagementGetManagedDeviceOverview(ctx).Select_(select_).Expand(expand).Execute()

Get managedDeviceOverview from deviceManagement



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
    resp, r, err := api_client.DeviceManagementManagedDeviceOverviewApi.DeviceManagementGetManagedDeviceOverview(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceOverviewApi.DeviceManagementGetManagedDeviceOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetManagedDeviceOverview`: MicrosoftGraphManagedDeviceOverview
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceOverviewApi.DeviceManagementGetManagedDeviceOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetManagedDeviceOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphManagedDeviceOverview**](MicrosoftGraphManagedDeviceOverview.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementGetRefManagedDeviceOverview

> string DeviceManagementGetRefManagedDeviceOverview(ctx).Execute()

Get ref of managedDeviceOverview from deviceManagement



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
    resp, r, err := api_client.DeviceManagementManagedDeviceOverviewApi.DeviceManagementGetRefManagedDeviceOverview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceOverviewApi.DeviceManagementGetRefManagedDeviceOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetRefManagedDeviceOverview`: string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementManagedDeviceOverviewApi.DeviceManagementGetRefManagedDeviceOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetRefManagedDeviceOverviewRequest struct via the builder pattern


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


## DeviceManagementUpdateRefManagedDeviceOverview

> DeviceManagementUpdateRefManagedDeviceOverview(ctx).RequestBody(requestBody).Execute()

Update the ref of navigation property managedDeviceOverview in deviceManagement



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
    resp, r, err := api_client.DeviceManagementManagedDeviceOverviewApi.DeviceManagementUpdateRefManagedDeviceOverview(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementManagedDeviceOverviewApi.DeviceManagementUpdateRefManagedDeviceOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateRefManagedDeviceOverviewRequest struct via the builder pattern


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

