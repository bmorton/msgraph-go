# \SitesItemAnalyticsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesDeleteRefAnalytics**](SitesItemAnalyticsApi.md#SitesDeleteRefAnalytics) | **Delete** /sites/{site-id}/analytics/$ref | Delete ref of navigation property analytics for sites
[**SitesGetAnalytics**](SitesItemAnalyticsApi.md#SitesGetAnalytics) | **Get** /sites/{site-id}/analytics | Get analytics from sites
[**SitesGetRefAnalytics**](SitesItemAnalyticsApi.md#SitesGetRefAnalytics) | **Get** /sites/{site-id}/analytics/$ref | Get ref of analytics from sites
[**SitesUpdateRefAnalytics**](SitesItemAnalyticsApi.md#SitesUpdateRefAnalytics) | **Put** /sites/{site-id}/analytics/$ref | Update the ref of navigation property analytics in sites



## SitesDeleteRefAnalytics

> SitesDeleteRefAnalytics(ctx, siteId).IfMatch(ifMatch).Execute()

Delete ref of navigation property analytics for sites



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
    siteId := "siteId_example" // string | key: id of site
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesItemAnalyticsApi.SitesDeleteRefAnalytics(context.Background(), siteId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesItemAnalyticsApi.SitesDeleteRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesDeleteRefAnalyticsRequest struct via the builder pattern


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


## SitesGetAnalytics

> MicrosoftGraphItemAnalytics SitesGetAnalytics(ctx, siteId).Select_(select_).Expand(expand).Execute()

Get analytics from sites



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
    siteId := "siteId_example" // string | key: id of site
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesItemAnalyticsApi.SitesGetAnalytics(context.Background(), siteId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesItemAnalyticsApi.SitesGetAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetAnalytics`: MicrosoftGraphItemAnalytics
    fmt.Fprintf(os.Stdout, "Response from `SitesItemAnalyticsApi.SitesGetAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetAnalyticsRequest struct via the builder pattern


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


## SitesGetRefAnalytics

> string SitesGetRefAnalytics(ctx, siteId).Execute()

Get ref of analytics from sites



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
    siteId := "siteId_example" // string | key: id of site

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesItemAnalyticsApi.SitesGetRefAnalytics(context.Background(), siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesItemAnalyticsApi.SitesGetRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetRefAnalytics`: string
    fmt.Fprintf(os.Stdout, "Response from `SitesItemAnalyticsApi.SitesGetRefAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetRefAnalyticsRequest struct via the builder pattern


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


## SitesUpdateRefAnalytics

> SitesUpdateRefAnalytics(ctx, siteId).RequestBody(requestBody).Execute()

Update the ref of navigation property analytics in sites



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
    siteId := "siteId_example" // string | key: id of site
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesItemAnalyticsApi.SitesUpdateRefAnalytics(context.Background(), siteId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesItemAnalyticsApi.SitesUpdateRefAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesUpdateRefAnalyticsRequest struct via the builder pattern


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

