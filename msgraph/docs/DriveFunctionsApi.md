# \DriveFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DriveListContentTypesContentTypeBaseIsPublished**](DriveFunctionsApi.md#DriveListContentTypesContentTypeBaseIsPublished) | **Get** /drive/list/contentTypes/{contentType-id}/base/microsoft.graph.isPublished() | Invoke function isPublished
[**DriveListContentTypesContentTypeIsPublished**](DriveFunctionsApi.md#DriveListContentTypesContentTypeIsPublished) | **Get** /drive/list/contentTypes/{contentType-id}/microsoft.graph.isPublished() | Invoke function isPublished
[**DriveListItemsListItemGetActivitiesByInterval53ee**](DriveFunctionsApi.md#DriveListItemsListItemGetActivitiesByInterval53ee) | **Get** /drive/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval(startDateTime&#x3D;&#39;{startDateTime}&#39;,endDateTime&#x3D;&#39;{endDateTime}&#39;,interval&#x3D;&#39;{interval}&#39;) | Invoke function getActivitiesByInterval
[**DriveListItemsListItemGetActivitiesByInterval96b0**](DriveFunctionsApi.md#DriveListItemsListItemGetActivitiesByInterval96b0) | **Get** /drive/list/items/{listItem-id}/microsoft.graph.getActivitiesByInterval() | Invoke function getActivitiesByInterval
[**DriveRecent**](DriveFunctionsApi.md#DriveRecent) | **Get** /drive/microsoft.graph.recent() | Invoke function recent
[**DriveSearch**](DriveFunctionsApi.md#DriveSearch) | **Get** /drive/microsoft.graph.search(q&#x3D;&#39;{q}&#39;) | Invoke function search
[**DriveSharedWithMe**](DriveFunctionsApi.md#DriveSharedWithMe) | **Get** /drive/microsoft.graph.sharedWithMe() | Invoke function sharedWithMe



## DriveListContentTypesContentTypeBaseIsPublished

> bool DriveListContentTypesContentTypeBaseIsPublished(ctx, contentTypeId).Execute()

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveFunctionsApi.DriveListContentTypesContentTypeBaseIsPublished(context.Background(), contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveFunctionsApi.DriveListContentTypesContentTypeBaseIsPublished``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListContentTypesContentTypeBaseIsPublished`: bool
    fmt.Fprintf(os.Stdout, "Response from `DriveFunctionsApi.DriveListContentTypesContentTypeBaseIsPublished`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeBaseIsPublishedRequest struct via the builder pattern


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


## DriveListContentTypesContentTypeIsPublished

> bool DriveListContentTypesContentTypeIsPublished(ctx, contentTypeId).Execute()

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
    contentTypeId := "contentTypeId_example" // string | key: id of contentType

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveFunctionsApi.DriveListContentTypesContentTypeIsPublished(context.Background(), contentTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveFunctionsApi.DriveListContentTypesContentTypeIsPublished``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListContentTypesContentTypeIsPublished`: bool
    fmt.Fprintf(os.Stdout, "Response from `DriveFunctionsApi.DriveListContentTypesContentTypeIsPublished`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentTypeId** | **string** | key: id of contentType | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListContentTypesContentTypeIsPublishedRequest struct via the builder pattern


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


## DriveListItemsListItemGetActivitiesByInterval53ee

> []*AnyOfmicrosoftGraphItemActivityStat DriveListItemsListItemGetActivitiesByInterval53ee(ctx, listItemId, startDateTime, endDateTime, interval).Execute()

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
    listItemId := "listItemId_example" // string | key: id of listItem
    startDateTime := "startDateTime_example" // string | Usage: startDateTime={startDateTime}
    endDateTime := "endDateTime_example" // string | Usage: endDateTime={endDateTime}
    interval := "interval_example" // string | Usage: interval={interval}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveFunctionsApi.DriveListItemsListItemGetActivitiesByInterval53ee(context.Background(), listItemId, startDateTime, endDateTime, interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveFunctionsApi.DriveListItemsListItemGetActivitiesByInterval53ee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListItemsListItemGetActivitiesByInterval53ee`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `DriveFunctionsApi.DriveListItemsListItemGetActivitiesByInterval53ee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string** | key: id of listItem | 
**startDateTime** | **string** | Usage: startDateTime&#x3D;{startDateTime} | 
**endDateTime** | **string** | Usage: endDateTime&#x3D;{endDateTime} | 
**interval** | **string** | Usage: interval&#x3D;{interval} | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListItemsListItemGetActivitiesByInterval53eeRequest struct via the builder pattern


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


## DriveListItemsListItemGetActivitiesByInterval96b0

> []*AnyOfmicrosoftGraphItemActivityStat DriveListItemsListItemGetActivitiesByInterval96b0(ctx, listItemId).Execute()

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
    listItemId := "listItemId_example" // string | key: id of listItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveFunctionsApi.DriveListItemsListItemGetActivitiesByInterval96b0(context.Background(), listItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveFunctionsApi.DriveListItemsListItemGetActivitiesByInterval96b0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveListItemsListItemGetActivitiesByInterval96b0`: []*AnyOfmicrosoftGraphItemActivityStat
    fmt.Fprintf(os.Stdout, "Response from `DriveFunctionsApi.DriveListItemsListItemGetActivitiesByInterval96b0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listItemId** | **string** | key: id of listItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveListItemsListItemGetActivitiesByInterval96b0Request struct via the builder pattern


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


## DriveRecent

> []*AnyOfmicrosoftGraphDriveItem DriveRecent(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveFunctionsApi.DriveRecent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveFunctionsApi.DriveRecent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveRecent`: []*AnyOfmicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveFunctionsApi.DriveRecent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDriveRecentRequest struct via the builder pattern


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


## DriveSearch

> []*AnyOfmicrosoftGraphDriveItem DriveSearch(ctx, q).Execute()

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
    q := "q_example" // string | Usage: q={q}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveFunctionsApi.DriveSearch(context.Background(), q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveFunctionsApi.DriveSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveSearch`: []*AnyOfmicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveFunctionsApi.DriveSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**q** | **string** | Usage: q&#x3D;{q} | 

### Other Parameters

Other parameters are passed through a pointer to a apiDriveSearchRequest struct via the builder pattern


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


## DriveSharedWithMe

> []*AnyOfmicrosoftGraphDriveItem DriveSharedWithMe(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DriveFunctionsApi.DriveSharedWithMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DriveFunctionsApi.DriveSharedWithMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DriveSharedWithMe`: []*AnyOfmicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `DriveFunctionsApi.DriveSharedWithMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDriveSharedWithMeRequest struct via the builder pattern


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

