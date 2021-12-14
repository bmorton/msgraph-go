# \DeviceManagementDeviceManagementReportsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementDeleteReports**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementDeleteReports) | **Delete** /deviceManagement/reports | Delete navigation property reports for deviceManagement
[**DeviceManagementGetReports**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementGetReports) | **Get** /deviceManagement/reports | Get reports from deviceManagement
[**DeviceManagementReportsCreateExportJobs**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementReportsCreateExportJobs) | **Post** /deviceManagement/reports/exportJobs | Create new navigation property to exportJobs for deviceManagement
[**DeviceManagementReportsDeleteExportJobs**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementReportsDeleteExportJobs) | **Delete** /deviceManagement/reports/exportJobs/{deviceManagementExportJob-id} | Delete navigation property exportJobs for deviceManagement
[**DeviceManagementReportsGetExportJobs**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementReportsGetExportJobs) | **Get** /deviceManagement/reports/exportJobs/{deviceManagementExportJob-id} | Get exportJobs from deviceManagement
[**DeviceManagementReportsListExportJobs**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementReportsListExportJobs) | **Get** /deviceManagement/reports/exportJobs | Get exportJobs from deviceManagement
[**DeviceManagementReportsUpdateExportJobs**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementReportsUpdateExportJobs) | **Patch** /deviceManagement/reports/exportJobs/{deviceManagementExportJob-id} | Update the navigation property exportJobs in deviceManagement
[**DeviceManagementUpdateReports**](DeviceManagementDeviceManagementReportsApi.md#DeviceManagementUpdateReports) | **Patch** /deviceManagement/reports | Update the navigation property reports in deviceManagement



## DeviceManagementDeleteReports

> DeviceManagementDeleteReports(ctx).IfMatch(ifMatch).Execute()

Delete navigation property reports for deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementDeleteReports(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementDeleteReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteReportsRequest struct via the builder pattern


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


## DeviceManagementGetReports

> MicrosoftGraphDeviceManagementReports DeviceManagementGetReports(ctx).Select_(select_).Expand(expand).Execute()

Get reports from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementGetReports(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementGetReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetReports`: MicrosoftGraphDeviceManagementReports
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementReportsApi.DeviceManagementGetReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementReports**](MicrosoftGraphDeviceManagementReports.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementReportsCreateExportJobs

> MicrosoftGraphDeviceManagementExportJob DeviceManagementReportsCreateExportJobs(ctx).MicrosoftGraphDeviceManagementExportJob(microsoftGraphDeviceManagementExportJob).Execute()

Create new navigation property to exportJobs for deviceManagement



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
    microsoftGraphDeviceManagementExportJob := *openapiclient.NewMicrosoftGraphDeviceManagementExportJob() // MicrosoftGraphDeviceManagementExportJob | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsCreateExportJobs(context.Background()).MicrosoftGraphDeviceManagementExportJob(microsoftGraphDeviceManagementExportJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsCreateExportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementReportsCreateExportJobs`: MicrosoftGraphDeviceManagementExportJob
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsCreateExportJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementReportsCreateExportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceManagementExportJob** | [**MicrosoftGraphDeviceManagementExportJob**](MicrosoftGraphDeviceManagementExportJob.md) | New navigation property | 

### Return type

[**MicrosoftGraphDeviceManagementExportJob**](MicrosoftGraphDeviceManagementExportJob.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementReportsDeleteExportJobs

> DeviceManagementReportsDeleteExportJobs(ctx, deviceManagementExportJobId).IfMatch(ifMatch).Execute()

Delete navigation property exportJobs for deviceManagement



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
    deviceManagementExportJobId := "deviceManagementExportJobId_example" // string | key: id of deviceManagementExportJob
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsDeleteExportJobs(context.Background(), deviceManagementExportJobId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsDeleteExportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExportJobId** | **string** | key: id of deviceManagementExportJob | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementReportsDeleteExportJobsRequest struct via the builder pattern


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


## DeviceManagementReportsGetExportJobs

> MicrosoftGraphDeviceManagementExportJob DeviceManagementReportsGetExportJobs(ctx, deviceManagementExportJobId).Select_(select_).Expand(expand).Execute()

Get exportJobs from deviceManagement



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
    deviceManagementExportJobId := "deviceManagementExportJobId_example" // string | key: id of deviceManagementExportJob
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsGetExportJobs(context.Background(), deviceManagementExportJobId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsGetExportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementReportsGetExportJobs`: MicrosoftGraphDeviceManagementExportJob
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsGetExportJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExportJobId** | **string** | key: id of deviceManagementExportJob | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementReportsGetExportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDeviceManagementExportJob**](MicrosoftGraphDeviceManagementExportJob.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementReportsListExportJobs

> CollectionOfDeviceManagementExportJob DeviceManagementReportsListExportJobs(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get exportJobs from deviceManagement



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
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsListExportJobs(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsListExportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementReportsListExportJobs`: CollectionOfDeviceManagementExportJob
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsListExportJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementReportsListExportJobsRequest struct via the builder pattern


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

[**CollectionOfDeviceManagementExportJob**](CollectionOfDeviceManagementExportJob.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementReportsUpdateExportJobs

> DeviceManagementReportsUpdateExportJobs(ctx, deviceManagementExportJobId).MicrosoftGraphDeviceManagementExportJob(microsoftGraphDeviceManagementExportJob).Execute()

Update the navigation property exportJobs in deviceManagement



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
    deviceManagementExportJobId := "deviceManagementExportJobId_example" // string | key: id of deviceManagementExportJob
    microsoftGraphDeviceManagementExportJob := *openapiclient.NewMicrosoftGraphDeviceManagementExportJob() // MicrosoftGraphDeviceManagementExportJob | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsUpdateExportJobs(context.Background(), deviceManagementExportJobId).MicrosoftGraphDeviceManagementExportJob(microsoftGraphDeviceManagementExportJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementReportsUpdateExportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceManagementExportJobId** | **string** | key: id of deviceManagementExportJob | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementReportsUpdateExportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDeviceManagementExportJob** | [**MicrosoftGraphDeviceManagementExportJob**](MicrosoftGraphDeviceManagementExportJob.md) | New navigation property values | 

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


## DeviceManagementUpdateReports

> DeviceManagementUpdateReports(ctx).MicrosoftGraphDeviceManagementReports(microsoftGraphDeviceManagementReports).Execute()

Update the navigation property reports in deviceManagement



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
    microsoftGraphDeviceManagementReports := *openapiclient.NewMicrosoftGraphDeviceManagementReports() // MicrosoftGraphDeviceManagementReports | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementDeviceManagementReportsApi.DeviceManagementUpdateReports(context.Background()).MicrosoftGraphDeviceManagementReports(microsoftGraphDeviceManagementReports).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementDeviceManagementReportsApi.DeviceManagementUpdateReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDeviceManagementReports** | [**MicrosoftGraphDeviceManagementReports**](MicrosoftGraphDeviceManagementReports.md) | New navigation property values | 

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

