# \MeOfficeGraphInsightsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDeleteInsights**](MeOfficeGraphInsightsApi.md#MeDeleteInsights) | **Delete** /me/insights | Delete navigation property insights for me
[**MeGetInsights**](MeOfficeGraphInsightsApi.md#MeGetInsights) | **Get** /me/insights | Get insights from me
[**MeInsightsCreateShared**](MeOfficeGraphInsightsApi.md#MeInsightsCreateShared) | **Post** /me/insights/shared | Create new navigation property to shared for me
[**MeInsightsCreateTrending**](MeOfficeGraphInsightsApi.md#MeInsightsCreateTrending) | **Post** /me/insights/trending | Create new navigation property to trending for me
[**MeInsightsCreateUsed**](MeOfficeGraphInsightsApi.md#MeInsightsCreateUsed) | **Post** /me/insights/used | Create new navigation property to used for me
[**MeInsightsDeleteShared**](MeOfficeGraphInsightsApi.md#MeInsightsDeleteShared) | **Delete** /me/insights/shared/{sharedInsight-id} | Delete navigation property shared for me
[**MeInsightsDeleteTrending**](MeOfficeGraphInsightsApi.md#MeInsightsDeleteTrending) | **Delete** /me/insights/trending/{trending-id} | Delete navigation property trending for me
[**MeInsightsDeleteUsed**](MeOfficeGraphInsightsApi.md#MeInsightsDeleteUsed) | **Delete** /me/insights/used/{usedInsight-id} | Delete navigation property used for me
[**MeInsightsGetShared**](MeOfficeGraphInsightsApi.md#MeInsightsGetShared) | **Get** /me/insights/shared/{sharedInsight-id} | Get shared from me
[**MeInsightsGetTrending**](MeOfficeGraphInsightsApi.md#MeInsightsGetTrending) | **Get** /me/insights/trending/{trending-id} | Get trending from me
[**MeInsightsGetUsed**](MeOfficeGraphInsightsApi.md#MeInsightsGetUsed) | **Get** /me/insights/used/{usedInsight-id} | Get used from me
[**MeInsightsListShared**](MeOfficeGraphInsightsApi.md#MeInsightsListShared) | **Get** /me/insights/shared | Get shared from me
[**MeInsightsListTrending**](MeOfficeGraphInsightsApi.md#MeInsightsListTrending) | **Get** /me/insights/trending | Get trending from me
[**MeInsightsListUsed**](MeOfficeGraphInsightsApi.md#MeInsightsListUsed) | **Get** /me/insights/used | Get used from me
[**MeInsightsSharedDeleteRefLastSharedMethod**](MeOfficeGraphInsightsApi.md#MeInsightsSharedDeleteRefLastSharedMethod) | **Delete** /me/insights/shared/{sharedInsight-id}/lastSharedMethod/$ref | Delete ref of navigation property lastSharedMethod for me
[**MeInsightsSharedDeleteRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsSharedDeleteRefResource) | **Delete** /me/insights/shared/{sharedInsight-id}/resource/$ref | Delete ref of navigation property resource for me
[**MeInsightsSharedGetLastSharedMethod**](MeOfficeGraphInsightsApi.md#MeInsightsSharedGetLastSharedMethod) | **Get** /me/insights/shared/{sharedInsight-id}/lastSharedMethod | Get lastSharedMethod from me
[**MeInsightsSharedGetRefLastSharedMethod**](MeOfficeGraphInsightsApi.md#MeInsightsSharedGetRefLastSharedMethod) | **Get** /me/insights/shared/{sharedInsight-id}/lastSharedMethod/$ref | Get ref of lastSharedMethod from me
[**MeInsightsSharedGetRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsSharedGetRefResource) | **Get** /me/insights/shared/{sharedInsight-id}/resource/$ref | Get ref of resource from me
[**MeInsightsSharedGetResource**](MeOfficeGraphInsightsApi.md#MeInsightsSharedGetResource) | **Get** /me/insights/shared/{sharedInsight-id}/resource | Get resource from me
[**MeInsightsSharedUpdateRefLastSharedMethod**](MeOfficeGraphInsightsApi.md#MeInsightsSharedUpdateRefLastSharedMethod) | **Put** /me/insights/shared/{sharedInsight-id}/lastSharedMethod/$ref | Update the ref of navigation property lastSharedMethod in me
[**MeInsightsSharedUpdateRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsSharedUpdateRefResource) | **Put** /me/insights/shared/{sharedInsight-id}/resource/$ref | Update the ref of navigation property resource in me
[**MeInsightsTrendingDeleteRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsTrendingDeleteRefResource) | **Delete** /me/insights/trending/{trending-id}/resource/$ref | Delete ref of navigation property resource for me
[**MeInsightsTrendingGetRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsTrendingGetRefResource) | **Get** /me/insights/trending/{trending-id}/resource/$ref | Get ref of resource from me
[**MeInsightsTrendingGetResource**](MeOfficeGraphInsightsApi.md#MeInsightsTrendingGetResource) | **Get** /me/insights/trending/{trending-id}/resource | Get resource from me
[**MeInsightsTrendingUpdateRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsTrendingUpdateRefResource) | **Put** /me/insights/trending/{trending-id}/resource/$ref | Update the ref of navigation property resource in me
[**MeInsightsUpdateShared**](MeOfficeGraphInsightsApi.md#MeInsightsUpdateShared) | **Patch** /me/insights/shared/{sharedInsight-id} | Update the navigation property shared in me
[**MeInsightsUpdateTrending**](MeOfficeGraphInsightsApi.md#MeInsightsUpdateTrending) | **Patch** /me/insights/trending/{trending-id} | Update the navigation property trending in me
[**MeInsightsUpdateUsed**](MeOfficeGraphInsightsApi.md#MeInsightsUpdateUsed) | **Patch** /me/insights/used/{usedInsight-id} | Update the navigation property used in me
[**MeInsightsUsedDeleteRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsUsedDeleteRefResource) | **Delete** /me/insights/used/{usedInsight-id}/resource/$ref | Delete ref of navigation property resource for me
[**MeInsightsUsedGetRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsUsedGetRefResource) | **Get** /me/insights/used/{usedInsight-id}/resource/$ref | Get ref of resource from me
[**MeInsightsUsedGetResource**](MeOfficeGraphInsightsApi.md#MeInsightsUsedGetResource) | **Get** /me/insights/used/{usedInsight-id}/resource | Get resource from me
[**MeInsightsUsedUpdateRefResource**](MeOfficeGraphInsightsApi.md#MeInsightsUsedUpdateRefResource) | **Put** /me/insights/used/{usedInsight-id}/resource/$ref | Update the ref of navigation property resource in me
[**MeUpdateInsights**](MeOfficeGraphInsightsApi.md#MeUpdateInsights) | **Patch** /me/insights | Update the navigation property insights in me



## MeDeleteInsights

> MeDeleteInsights(ctx).IfMatch(ifMatch).Execute()

Delete navigation property insights for me



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
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeDeleteInsights(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeDeleteInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteInsightsRequest struct via the builder pattern


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


## MeGetInsights

> MicrosoftGraphOfficeGraphInsights MeGetInsights(ctx).Select_(select_).Expand(expand).Execute()

Get insights from me



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
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeGetInsights(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeGetInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetInsights`: MicrosoftGraphOfficeGraphInsights
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeGetInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphOfficeGraphInsights**](MicrosoftGraphOfficeGraphInsights.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsCreateShared

> MicrosoftGraphSharedInsight MeInsightsCreateShared(ctx).MicrosoftGraphSharedInsight(microsoftGraphSharedInsight).Execute()

Create new navigation property to shared for me



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
    microsoftGraphSharedInsight := *openapiclient.NewMicrosoftGraphSharedInsight() // MicrosoftGraphSharedInsight | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsCreateShared(context.Background()).MicrosoftGraphSharedInsight(microsoftGraphSharedInsight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsCreateShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsCreateShared`: MicrosoftGraphSharedInsight
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsCreateShared`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsCreateSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSharedInsight** | [**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md) | New navigation property | 

### Return type

[**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsCreateTrending

> MicrosoftGraphTrending MeInsightsCreateTrending(ctx).MicrosoftGraphTrending(microsoftGraphTrending).Execute()

Create new navigation property to trending for me



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
    microsoftGraphTrending := *openapiclient.NewMicrosoftGraphTrending() // MicrosoftGraphTrending | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsCreateTrending(context.Background()).MicrosoftGraphTrending(microsoftGraphTrending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsCreateTrending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsCreateTrending`: MicrosoftGraphTrending
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsCreateTrending`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsCreateTrendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTrending** | [**MicrosoftGraphTrending**](MicrosoftGraphTrending.md) | New navigation property | 

### Return type

[**MicrosoftGraphTrending**](MicrosoftGraphTrending.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsCreateUsed

> MicrosoftGraphUsedInsight MeInsightsCreateUsed(ctx).MicrosoftGraphUsedInsight(microsoftGraphUsedInsight).Execute()

Create new navigation property to used for me



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
    microsoftGraphUsedInsight := *openapiclient.NewMicrosoftGraphUsedInsight() // MicrosoftGraphUsedInsight | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsCreateUsed(context.Background()).MicrosoftGraphUsedInsight(microsoftGraphUsedInsight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsCreateUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsCreateUsed`: MicrosoftGraphUsedInsight
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsCreateUsed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsCreateUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphUsedInsight** | [**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md) | New navigation property | 

### Return type

[**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsDeleteShared

> MeInsightsDeleteShared(ctx, sharedInsightId).IfMatch(ifMatch).Execute()

Delete navigation property shared for me



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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsDeleteShared(context.Background(), sharedInsightId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsDeleteShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsDeleteSharedRequest struct via the builder pattern


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


## MeInsightsDeleteTrending

> MeInsightsDeleteTrending(ctx, trendingId).IfMatch(ifMatch).Execute()

Delete navigation property trending for me



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
    trendingId := "trendingId_example" // string | key: id of trending
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsDeleteTrending(context.Background(), trendingId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsDeleteTrending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string** | key: id of trending | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsDeleteTrendingRequest struct via the builder pattern


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


## MeInsightsDeleteUsed

> MeInsightsDeleteUsed(ctx, usedInsightId).IfMatch(ifMatch).Execute()

Delete navigation property used for me



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
    usedInsightId := "usedInsightId_example" // string | key: id of usedInsight
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsDeleteUsed(context.Background(), usedInsightId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsDeleteUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string** | key: id of usedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsDeleteUsedRequest struct via the builder pattern


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


## MeInsightsGetShared

> MicrosoftGraphSharedInsight MeInsightsGetShared(ctx, sharedInsightId).Select_(select_).Expand(expand).Execute()

Get shared from me



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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsGetShared(context.Background(), sharedInsightId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsGetShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsGetShared`: MicrosoftGraphSharedInsight
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsGetShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsGetSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsGetTrending

> MicrosoftGraphTrending MeInsightsGetTrending(ctx, trendingId).Select_(select_).Expand(expand).Execute()

Get trending from me



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
    trendingId := "trendingId_example" // string | key: id of trending
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsGetTrending(context.Background(), trendingId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsGetTrending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsGetTrending`: MicrosoftGraphTrending
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsGetTrending`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string** | key: id of trending | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsGetTrendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTrending**](MicrosoftGraphTrending.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsGetUsed

> MicrosoftGraphUsedInsight MeInsightsGetUsed(ctx, usedInsightId).Select_(select_).Expand(expand).Execute()

Get used from me



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
    usedInsightId := "usedInsightId_example" // string | key: id of usedInsight
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsGetUsed(context.Background(), usedInsightId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsGetUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsGetUsed`: MicrosoftGraphUsedInsight
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsGetUsed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string** | key: id of usedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsGetUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsListShared

> CollectionOfSharedInsight MeInsightsListShared(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get shared from me



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
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsListShared(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsListShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsListShared`: CollectionOfSharedInsight
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsListShared`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsListSharedRequest struct via the builder pattern


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

[**CollectionOfSharedInsight**](CollectionOfSharedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsListTrending

> CollectionOfTrending MeInsightsListTrending(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get trending from me



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
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsListTrending(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsListTrending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsListTrending`: CollectionOfTrending
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsListTrending`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsListTrendingRequest struct via the builder pattern


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

[**CollectionOfTrending**](CollectionOfTrending.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsListUsed

> CollectionOfUsedInsight MeInsightsListUsed(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get used from me



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
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsListUsed(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsListUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsListUsed`: CollectionOfUsedInsight
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsListUsed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsListUsedRequest struct via the builder pattern


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

[**CollectionOfUsedInsight**](CollectionOfUsedInsight.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsSharedDeleteRefLastSharedMethod

> MeInsightsSharedDeleteRefLastSharedMethod(ctx, sharedInsightId).IfMatch(ifMatch).Execute()

Delete ref of navigation property lastSharedMethod for me

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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedDeleteRefLastSharedMethod(context.Background(), sharedInsightId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedDeleteRefLastSharedMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedDeleteRefLastSharedMethodRequest struct via the builder pattern


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


## MeInsightsSharedDeleteRefResource

> MeInsightsSharedDeleteRefResource(ctx, sharedInsightId).IfMatch(ifMatch).Execute()

Delete ref of navigation property resource for me



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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedDeleteRefResource(context.Background(), sharedInsightId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedDeleteRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedDeleteRefResourceRequest struct via the builder pattern


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


## MeInsightsSharedGetLastSharedMethod

> MicrosoftGraphEntity MeInsightsSharedGetLastSharedMethod(ctx, sharedInsightId).Select_(select_).Expand(expand).Execute()

Get lastSharedMethod from me

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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedGetLastSharedMethod(context.Background(), sharedInsightId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedGetLastSharedMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsSharedGetLastSharedMethod`: MicrosoftGraphEntity
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsSharedGetLastSharedMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedGetLastSharedMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](MicrosoftGraphEntity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsSharedGetRefLastSharedMethod

> string MeInsightsSharedGetRefLastSharedMethod(ctx, sharedInsightId).Execute()

Get ref of lastSharedMethod from me

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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedGetRefLastSharedMethod(context.Background(), sharedInsightId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedGetRefLastSharedMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsSharedGetRefLastSharedMethod`: string
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsSharedGetRefLastSharedMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedGetRefLastSharedMethodRequest struct via the builder pattern


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


## MeInsightsSharedGetRefResource

> string MeInsightsSharedGetRefResource(ctx, sharedInsightId).Execute()

Get ref of resource from me



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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedGetRefResource(context.Background(), sharedInsightId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedGetRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsSharedGetRefResource`: string
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsSharedGetRefResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedGetRefResourceRequest struct via the builder pattern


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


## MeInsightsSharedGetResource

> MicrosoftGraphEntity MeInsightsSharedGetResource(ctx, sharedInsightId).Select_(select_).Expand(expand).Execute()

Get resource from me



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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedGetResource(context.Background(), sharedInsightId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedGetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsSharedGetResource`: MicrosoftGraphEntity
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsSharedGetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](MicrosoftGraphEntity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsSharedUpdateRefLastSharedMethod

> MeInsightsSharedUpdateRefLastSharedMethod(ctx, sharedInsightId).RequestBody(requestBody).Execute()

Update the ref of navigation property lastSharedMethod in me

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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedUpdateRefLastSharedMethod(context.Background(), sharedInsightId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedUpdateRefLastSharedMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedUpdateRefLastSharedMethodRequest struct via the builder pattern


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


## MeInsightsSharedUpdateRefResource

> MeInsightsSharedUpdateRefResource(ctx, sharedInsightId).RequestBody(requestBody).Execute()

Update the ref of navigation property resource in me



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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsSharedUpdateRefResource(context.Background(), sharedInsightId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsSharedUpdateRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsSharedUpdateRefResourceRequest struct via the builder pattern


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


## MeInsightsTrendingDeleteRefResource

> MeInsightsTrendingDeleteRefResource(ctx, trendingId).IfMatch(ifMatch).Execute()

Delete ref of navigation property resource for me



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
    trendingId := "trendingId_example" // string | key: id of trending
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsTrendingDeleteRefResource(context.Background(), trendingId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsTrendingDeleteRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string** | key: id of trending | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsTrendingDeleteRefResourceRequest struct via the builder pattern


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


## MeInsightsTrendingGetRefResource

> string MeInsightsTrendingGetRefResource(ctx, trendingId).Execute()

Get ref of resource from me



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
    trendingId := "trendingId_example" // string | key: id of trending

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsTrendingGetRefResource(context.Background(), trendingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsTrendingGetRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsTrendingGetRefResource`: string
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsTrendingGetRefResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string** | key: id of trending | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsTrendingGetRefResourceRequest struct via the builder pattern


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


## MeInsightsTrendingGetResource

> MicrosoftGraphEntity MeInsightsTrendingGetResource(ctx, trendingId).Select_(select_).Expand(expand).Execute()

Get resource from me



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
    trendingId := "trendingId_example" // string | key: id of trending
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsTrendingGetResource(context.Background(), trendingId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsTrendingGetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsTrendingGetResource`: MicrosoftGraphEntity
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsTrendingGetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string** | key: id of trending | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsTrendingGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](MicrosoftGraphEntity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsTrendingUpdateRefResource

> MeInsightsTrendingUpdateRefResource(ctx, trendingId).RequestBody(requestBody).Execute()

Update the ref of navigation property resource in me



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
    trendingId := "trendingId_example" // string | key: id of trending
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsTrendingUpdateRefResource(context.Background(), trendingId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsTrendingUpdateRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string** | key: id of trending | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsTrendingUpdateRefResourceRequest struct via the builder pattern


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


## MeInsightsUpdateShared

> MeInsightsUpdateShared(ctx, sharedInsightId).MicrosoftGraphSharedInsight(microsoftGraphSharedInsight).Execute()

Update the navigation property shared in me



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
    sharedInsightId := "sharedInsightId_example" // string | key: id of sharedInsight
    microsoftGraphSharedInsight := *openapiclient.NewMicrosoftGraphSharedInsight() // MicrosoftGraphSharedInsight | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsUpdateShared(context.Background(), sharedInsightId).MicrosoftGraphSharedInsight(microsoftGraphSharedInsight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsUpdateShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedInsightId** | **string** | key: id of sharedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsUpdateSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSharedInsight** | [**MicrosoftGraphSharedInsight**](MicrosoftGraphSharedInsight.md) | New navigation property values | 

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


## MeInsightsUpdateTrending

> MeInsightsUpdateTrending(ctx, trendingId).MicrosoftGraphTrending(microsoftGraphTrending).Execute()

Update the navigation property trending in me



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
    trendingId := "trendingId_example" // string | key: id of trending
    microsoftGraphTrending := *openapiclient.NewMicrosoftGraphTrending() // MicrosoftGraphTrending | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsUpdateTrending(context.Background(), trendingId).MicrosoftGraphTrending(microsoftGraphTrending).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsUpdateTrending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trendingId** | **string** | key: id of trending | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsUpdateTrendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTrending** | [**MicrosoftGraphTrending**](MicrosoftGraphTrending.md) | New navigation property values | 

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


## MeInsightsUpdateUsed

> MeInsightsUpdateUsed(ctx, usedInsightId).MicrosoftGraphUsedInsight(microsoftGraphUsedInsight).Execute()

Update the navigation property used in me



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
    usedInsightId := "usedInsightId_example" // string | key: id of usedInsight
    microsoftGraphUsedInsight := *openapiclient.NewMicrosoftGraphUsedInsight() // MicrosoftGraphUsedInsight | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsUpdateUsed(context.Background(), usedInsightId).MicrosoftGraphUsedInsight(microsoftGraphUsedInsight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsUpdateUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string** | key: id of usedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsUpdateUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphUsedInsight** | [**MicrosoftGraphUsedInsight**](MicrosoftGraphUsedInsight.md) | New navigation property values | 

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


## MeInsightsUsedDeleteRefResource

> MeInsightsUsedDeleteRefResource(ctx, usedInsightId).IfMatch(ifMatch).Execute()

Delete ref of navigation property resource for me



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
    usedInsightId := "usedInsightId_example" // string | key: id of usedInsight
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsUsedDeleteRefResource(context.Background(), usedInsightId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsUsedDeleteRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string** | key: id of usedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsUsedDeleteRefResourceRequest struct via the builder pattern


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


## MeInsightsUsedGetRefResource

> string MeInsightsUsedGetRefResource(ctx, usedInsightId).Execute()

Get ref of resource from me



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
    usedInsightId := "usedInsightId_example" // string | key: id of usedInsight

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsUsedGetRefResource(context.Background(), usedInsightId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsUsedGetRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsUsedGetRefResource`: string
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsUsedGetRefResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string** | key: id of usedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsUsedGetRefResourceRequest struct via the builder pattern


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


## MeInsightsUsedGetResource

> MicrosoftGraphEntity MeInsightsUsedGetResource(ctx, usedInsightId).Select_(select_).Expand(expand).Execute()

Get resource from me



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
    usedInsightId := "usedInsightId_example" // string | key: id of usedInsight
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsUsedGetResource(context.Background(), usedInsightId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsUsedGetResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeInsightsUsedGetResource`: MicrosoftGraphEntity
    fmt.Fprintf(os.Stdout, "Response from `MeOfficeGraphInsightsApi.MeInsightsUsedGetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string** | key: id of usedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsUsedGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphEntity**](MicrosoftGraphEntity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeInsightsUsedUpdateRefResource

> MeInsightsUsedUpdateRefResource(ctx, usedInsightId).RequestBody(requestBody).Execute()

Update the ref of navigation property resource in me



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
    usedInsightId := "usedInsightId_example" // string | key: id of usedInsight
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeInsightsUsedUpdateRefResource(context.Background(), usedInsightId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeInsightsUsedUpdateRefResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usedInsightId** | **string** | key: id of usedInsight | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeInsightsUsedUpdateRefResourceRequest struct via the builder pattern


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


## MeUpdateInsights

> MeUpdateInsights(ctx).MicrosoftGraphOfficeGraphInsights(microsoftGraphOfficeGraphInsights).Execute()

Update the navigation property insights in me



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
    microsoftGraphOfficeGraphInsights := *openapiclient.NewMicrosoftGraphOfficeGraphInsights() // MicrosoftGraphOfficeGraphInsights | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOfficeGraphInsightsApi.MeUpdateInsights(context.Background()).MicrosoftGraphOfficeGraphInsights(microsoftGraphOfficeGraphInsights).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOfficeGraphInsightsApi.MeUpdateInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphOfficeGraphInsights** | [**MicrosoftGraphOfficeGraphInsights**](MicrosoftGraphOfficeGraphInsights.md) | New navigation property values | 

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

