# \DrivesFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DrivesDriveListContentTypesContentTypeBaseIsPublished**](DrivesFunctionsApi.md#DrivesDriveListContentTypesContentTypeBaseIsPublished) | **Get** /drives/{drive-id}/list/contentTypes/{contentType-id}/base/microsoft.graph.isPublished() | Invoke function isPublished
[**DrivesDriveListContentTypesContentTypeIsPublished**](DrivesFunctionsApi.md#DrivesDriveListContentTypesContentTypeIsPublished) | **Get** /drives/{drive-id}/list/contentTypes/{contentType-id}/microsoft.graph.isPublished() | Invoke function isPublished
[**DrivesDriveListItemsListItemGetActivitiesByInterval53ee**](DrivesFunctionsApi.md#DrivesDriveListItemsListItemGetActivitiesByInterval53ee) | **Get** /drives/{drive-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;&#39;{startDateTime}&#39;,endDateTime&#x3D;&#39;{endDateTime}&#39;,interval&#x3D;&#39;{interval}&#39;) | Invoke function getActivitiesByInterval
[**DrivesDriveListItemsListItemGetActivitiesByInterval96b0**](DrivesFunctionsApi.md#DrivesDriveListItemsListItemGetActivitiesByInterval96b0) | **Get** /drives/{drive-id}/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**DrivesDriveRecent**](DrivesFunctionsApi.md#DrivesDriveRecent) | **Get** /drives/{drive-id}/microsoft.graph.recent() | Invoke function recent
[**DrivesDriveSearch**](DrivesFunctionsApi.md#DrivesDriveSearch) | **Get** /drives/{drive-id}/microsoft.graph.search(q&#x3D;&#39;{q}&#39;) | Invoke function search
[**DrivesDriveSharedWithMe**](DrivesFunctionsApi.md#DrivesDriveSharedWithMe) | **Get** /drives/{drive-id}/microsoft.graph.sharedWithMe() | Invoke function sharedWithMe



## DrivesDriveListContentTypesContentTypeBaseIsPublished

> bool DrivesDriveListContentTypesContentTypeBaseIsPublished(ctx, driveId, contentTypeId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesFunctionsApi.DrivesDriveListContentTypesContentTypeBaseIsPublished(context.Background(), driveId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesFunctionsApi.DrivesDriveListContentTypesContentTypeBaseIsPublished``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveListContentTypesContentTypeBaseIsPublished`: bool
    fmt.Fprintf(os.Stdout, "Response from `DrivesFunctionsApi.DrivesDriveListContentTypesContentTypeBaseIsPublished`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeBaseIsPublishedRequest struct via the builder pattern


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


## DrivesDriveListContentTypesContentTypeIsPublished

> bool DrivesDriveListContentTypesContentTypeIsPublished(ctx, driveId, contentTypeId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesFunctionsApi.DrivesDriveListContentTypesContentTypeIsPublished(context.Background(), driveId, contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesFunctionsApi.DrivesDriveListContentTypesContentTypeIsPublished``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveListContentTypesContentTypeIsPublished`: bool
    fmt.Fprintf(os.Stdout, "Response from `DrivesFunctionsApi.DrivesDriveListContentTypesContentTypeIsPublished`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListContentTypesContentTypeIsPublishedRequest struct via the builder pattern


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


## DrivesDriveListItemsListItemGetActivitiesByInterval53ee

> []*AnyOfmicrosoftGraphItemActivityStat DrivesDriveListItemsListItemGetActivitiesByInterval53ee(ctx, driveId, listItemId, startDateTime, endDateTime, interval).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    listItemId := "listItemId_example" // string | key: id of listItem
    startDateTime := "startDateTime_example" // string | Usage: startDateTime={startDateTime}
    endDateTime := "endDateTime_example" // string | Usage: endDateTime={endDateTime}
    interval := "interval_example" // string | Usage: interval={interval}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesFunctionsApi.DrivesDriveListItemsListItemGetActivitiesByInterval53ee(context.Background(), driveId, listItemId, startDateTime, endDateTime, interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesFunctionsApi.DrivesDriveListItemsListItemGetActivitiesByInterval53ee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveListItemsListItemGetActivitiesByInterval53ee`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `DrivesFunctionsApi.DrivesDriveListItemsListItemGetActivitiesByInterval53ee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**listItemId** | **string** | key: id of listItem | 
**startDateTime** | **string** | Usage: startDateTime&#x3D;{startDateTime} | 
**endDateTime** | **string** | Usage: endDateTime&#x3D;{endDateTime} | 
**interval** | **string** | Usage: interval&#x3D;{interval} | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListItemsListItemGetActivitiesByInterval53eeRequest struct via the builder pattern


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


## DrivesDriveListItemsListItemGetActivitiesByInterval96b0

> []*AnyOfmicrosoftGraphItemActivityStat DrivesDriveListItemsListItemGetActivitiesByInterval96b0(ctx, driveId, listItemId).Execute()

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
    driveId := "driveId_example" // string | key: id of drive
    listItemId := "listItemId_example" // string | key: id of listItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesFunctionsApi.DrivesDriveListItemsListItemGetActivitiesByInterval96b0(context.Background(), driveId, listItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesFunctionsApi.DrivesDriveListItemsListItemGetActivitiesByInterval96b0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveListItemsListItemGetActivitiesByInterval96b0`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `DrivesFunctionsApi.DrivesDriveListItemsListItemGetActivitiesByInterval96b0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveListItemsListItemGetActivitiesByInterval96b0Request struct via the builder pattern


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


## DrivesDriveRecent

> []*AnyOfmicrosoftGraphDriveItem DrivesDriveRecent(ctx, driveId).Execute()

Invoke function recent

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
    driveId := "driveId_example" // string | key: id of drive

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesFunctionsApi.DrivesDriveRecent(context.Background(), driveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesFunctionsApi.DrivesDriveRecent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveRecent`: []*AnyOfmicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DrivesFunctionsApi.DrivesDriveRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesDriveSearch

> []*AnyOfmicrosoftGraphDriveItem DrivesDriveSearch(ctx, driveId, q).Execute()

Invoke function search

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
    driveId := "driveId_example" // string | key: id of drive
    q := "q_example" // string | Usage: q={q}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesFunctionsApi.DrivesDriveSearch(context.Background(), driveId, q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesFunctionsApi.DrivesDriveSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveSearch`: []*AnyOfmicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DrivesFunctionsApi.DrivesDriveSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 
**q** | **string** | Usage: q&#x3D;{q} | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DrivesDriveSharedWithMe

> []*AnyOfmicrosoftGraphDriveItem DrivesDriveSharedWithMe(ctx, driveId).Execute()

Invoke function sharedWithMe

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
    driveId := "driveId_example" // string | key: id of drive

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrivesFunctionsApi.DrivesDriveSharedWithMe(context.Background(), driveId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrivesFunctionsApi.DrivesDriveSharedWithMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DrivesDriveSharedWithMe`: []*AnyOfmicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DrivesFunctionsApi.DrivesDriveSharedWithMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveId** | **string** | key: id of drive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDrivesDriveSharedWithMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

