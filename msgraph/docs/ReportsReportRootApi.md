# \ReportsReportRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportsReportRootGetReportRoot**](ReportsReportRootApi.md#ReportsReportRootGetReportRoot) | **Get** /reports | Get reports
[**ReportsReportRootUpdateReportRoot**](ReportsReportRootApi.md#ReportsReportRootUpdateReportRoot) | **Patch** /reports | Update reports



## ReportsReportRootGetReportRoot

> MicrosoftGraphReportRoot ReportsReportRootGetReportRoot(ctx).Select_(select_).Expand(expand).Execute()

Get reports

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
    resp, r, err := api_client.ReportsReportRootApi.ReportsReportRootGetReportRoot(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsReportRootApi.ReportsReportRootGetReportRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsReportRootGetReportRoot`: MicrosoftGraphReportRoot
    fmt.Fprintf(os.Stdout, "Response from `ReportsReportRootApi.ReportsReportRootGetReportRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsReportRootGetReportRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphReportRoot**](MicrosoftGraphReportRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsReportRootUpdateReportRoot

> ReportsReportRootUpdateReportRoot(ctx).MicrosoftGraphReportRoot(microsoftGraphReportRoot).Execute()

Update reports

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
    microsoftGraphReportRoot := *openapiclient.NewMicrosoftGraphReportRoot() // MicrosoftGraphReportRoot | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsReportRootApi.ReportsReportRootUpdateReportRoot(context.Background()).MicrosoftGraphReportRoot(microsoftGraphReportRoot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsReportRootApi.ReportsReportRootUpdateReportRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportsReportRootUpdateReportRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphReportRoot** | [**MicrosoftGraphReportRoot**](MicrosoftGraphReportRoot.md) | New property values | 

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

