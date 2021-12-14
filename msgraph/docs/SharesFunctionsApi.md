# \SharesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished**](SharesFunctionsApi.md#SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished) | **Get** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.isPublished() | Invoke function isPublished
[**SharesSharedDriveItemListContentTypesContentTypeIsPublished**](SharesFunctionsApi.md#SharesSharedDriveItemListContentTypesContentTypeIsPublished) | **Get** /shares/{sharedDriveItem-id}/list/contentTypes/{contentType-id}/microsoft.graph.isPublished() | Invoke function isPublished
[**SharesSharedDriveItemListItemGetActivitiesByInterval53ee**](SharesFunctionsApi.md#SharesSharedDriveItemListItemGetActivitiesByInterval53ee) | **Get** /shares/{sharedDriveItem-id}/listItem/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;&#39;{startDateTime}&#39;,endDateTime&#x3D;&#39;{endDateTime}&#39;,interval&#x3D;&#39;{interval}&#39;) | Invoke function getActivitiesByInterval
[**SharesSharedDriveItemListItemGetActivitiesByInterval96b0**](SharesFunctionsApi.md#SharesSharedDriveItemListItemGetActivitiesByInterval96b0) | **Get** /shares/{sharedDriveItem-id}/listItem/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee**](SharesFunctionsApi.md#SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee) | **Get** /shares/{sharedDriveItem-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;&#39;{startDateTime}&#39;,endDateTime&#x3D;&#39;{endDateTime}&#39;,interval&#x3D;&#39;{interval}&#39;) | Invoke function getActivitiesByInterval
[**SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0**](SharesFunctionsApi.md#SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0) | **Get** /shares/{sharedDriveItem-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval



## SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished

> bool SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished(ctx, sharedDriveItemId, contentTypeId).Execute()

Invoke function isPublished

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesFunctionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished(context.Background(), sharedDriveItemId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesFunctionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished`: bool
    fmt.Fprintf(os.Stdout, "Response from `SharesFunctionsApi.SharesSharedDriveItemListContentTypesContentTypeBaseIsPublished`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeBaseIsPublishedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemListContentTypesContentTypeIsPublished

> bool SharesSharedDriveItemListContentTypesContentTypeIsPublished(ctx, sharedDriveItemId, contentTypeId).Execute()

Invoke function isPublished

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesFunctionsApi.SharesSharedDriveItemListContentTypesContentTypeIsPublished(context.Background(), sharedDriveItemId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesFunctionsApi.SharesSharedDriveItemListContentTypesContentTypeIsPublished``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListContentTypesContentTypeIsPublished`: bool
    fmt.Fprintf(os.Stdout, "Response from `SharesFunctionsApi.SharesSharedDriveItemListContentTypesContentTypeIsPublished`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListContentTypesContentTypeIsPublishedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemListItemGetActivitiesByInterval53ee

> []*AnyOfmicrosoftGraphItemActivityStat SharesSharedDriveItemListItemGetActivitiesByInterval53ee(ctx, sharedDriveItemId, startDateTime, endDateTime, interval).Execute()

Invoke function getActivitiesByInterval

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    startDateTime := "startDateTime_example" // string | Usage: startDateTime={startDateTime}
    endDateTime := "endDateTime_example" // string | Usage: endDateTime={endDateTime}
    interval := "interval_example" // string | Usage: interval={interval}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesFunctionsApi.SharesSharedDriveItemListItemGetActivitiesByInterval53ee(context.Background(), sharedDriveItemId, startDateTime, endDateTime, interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesFunctionsApi.SharesSharedDriveItemListItemGetActivitiesByInterval53ee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListItemGetActivitiesByInterval53ee`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `SharesFunctionsApi.SharesSharedDriveItemListItemGetActivitiesByInterval53ee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**startDateTime** | **string** | Usage: startDateTime&#x3D;{startDateTime} | 
**endDateTime** | **string** | Usage: endDateTime&#x3D;{endDateTime} | 
**interval** | **string** | Usage: interval&#x3D;{interval} | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListItemGetActivitiesByInterval53eeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]*AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemListItemGetActivitiesByInterval96b0

> []*AnyOfmicrosoftGraphItemActivityStat SharesSharedDriveItemListItemGetActivitiesByInterval96b0(ctx, sharedDriveItemId).Execute()

Invoke function getActivitiesByInterval

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesFunctionsApi.SharesSharedDriveItemListItemGetActivitiesByInterval96b0(context.Background(), sharedDriveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesFunctionsApi.SharesSharedDriveItemListItemGetActivitiesByInterval96b0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListItemGetActivitiesByInterval96b0`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `SharesFunctionsApi.SharesSharedDriveItemListItemGetActivitiesByInterval96b0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListItemGetActivitiesByInterval96b0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee

> []*AnyOfmicrosoftGraphItemActivityStat SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee(ctx, sharedDriveItemId, listItemId, startDateTime, endDateTime, interval).Execute()

Invoke function getActivitiesByInterval

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    listItemId := "listItemId_example" // string | key: id of listItem
    startDateTime := "startDateTime_example" // string | Usage: startDateTime={startDateTime}
    endDateTime := "endDateTime_example" // string | Usage: endDateTime={endDateTime}
    interval := "interval_example" // string | Usage: interval={interval}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesFunctionsApi.SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee(context.Background(), sharedDriveItemId, listItemId, startDateTime, endDateTime, interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesFunctionsApi.SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `SharesFunctionsApi.SharesSharedDriveItemListItemsListItemGetActivitiesByInterval53ee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**listItemId** | **string** | key: id of listItem | 
**startDateTime** | **string** | Usage: startDateTime&#x3D;{startDateTime} | 
**endDateTime** | **string** | Usage: endDateTime&#x3D;{endDateTime} | 
**interval** | **string** | Usage: interval&#x3D;{interval} | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListItemsListItemGetActivitiesByInterval53eeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**[]*AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0

> []*AnyOfmicrosoftGraphItemActivityStat SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0(ctx, sharedDriveItemId, listItemId).Execute()

Invoke function getActivitiesByInterval

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
    sharedDriveItemId := "sharedDriveItemId_example" // string | key: id of sharedDriveItem
    listItemId := "listItemId_example" // string | key: id of listItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesFunctionsApi.SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0(context.Background(), sharedDriveItemId, listItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesFunctionsApi.SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `SharesFunctionsApi.SharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListItemsListItemGetActivitiesByInterval96b0Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphItemActivityStat**](anyOf&lt;microsoft.graph.itemActivityStat&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

