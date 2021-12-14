# \WorkbooksItemAnalyticsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksDeleteRefAnalytics**](WorkbooksItemAnalyticsApi.md#WorkbooksDeleteRefAnalytics) | **Delete** /workbooks/{driveItem-id}/analytics/$ref | Delete ref of navigation property analytics for workbooks
[**WorkbooksGetAnalytics**](WorkbooksItemAnalyticsApi.md#WorkbooksGetAnalytics) | **Get** /workbooks/{driveItem-id}/analytics | Get analytics from workbooks
[**WorkbooksGetRefAnalytics**](WorkbooksItemAnalyticsApi.md#WorkbooksGetRefAnalytics) | **Get** /workbooks/{driveItem-id}/analytics/$ref | Get ref of analytics from workbooks
[**WorkbooksUpdateRefAnalytics**](WorkbooksItemAnalyticsApi.md#WorkbooksUpdateRefAnalytics) | **Put** /workbooks/{driveItem-id}/analytics/$ref | Update the ref of navigation property analytics in workbooks



## WorkbooksDeleteRefAnalytics

> WorkbooksDeleteRefAnalytics(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete ref of navigation property analytics for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksItemAnalyticsApi.WorkbooksDeleteRefAnalytics(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksItemAnalyticsApi.WorkbooksDeleteRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDeleteRefAnalyticsRequest struct via the builder pattern


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


## WorkbooksGetAnalytics

> MicrosoftGraphItemAnalytics WorkbooksGetAnalytics(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get analytics from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksItemAnalyticsApi.WorkbooksGetAnalytics(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksItemAnalyticsApi.WorkbooksGetAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetAnalytics`: MicrosoftGraphItemAnalytics
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksItemAnalyticsApi.WorkbooksGetAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphItemAnalytics**](MicrosoftGraphItemAnalytics.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksGetRefAnalytics

> string WorkbooksGetRefAnalytics(ctx, driveItemId).Execute()

Get ref of analytics from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksItemAnalyticsApi.WorkbooksGetRefAnalytics(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksItemAnalyticsApi.WorkbooksGetRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetRefAnalytics`: string
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksItemAnalyticsApi.WorkbooksGetRefAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetRefAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## WorkbooksUpdateRefAnalytics

> WorkbooksUpdateRefAnalytics(ctx, driveItemId).RequestBody(requestBody).Execute()

Update the ref of navigation property analytics in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksItemAnalyticsApi.WorkbooksUpdateRefAnalytics(context.Background(), driveItemId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksItemAnalyticsApi.WorkbooksUpdateRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateRefAnalyticsRequest struct via the builder pattern


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

