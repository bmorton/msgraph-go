# \WorkbooksThumbnailSetApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreateThumbnails**](WorkbooksThumbnailSetApi.md#WorkbooksCreateThumbnails) | **Post** /workbooks/{driveItem-id}/thumbnails | Create new navigation property to thumbnails for workbooks
[**WorkbooksDeleteThumbnails**](WorkbooksThumbnailSetApi.md#WorkbooksDeleteThumbnails) | **Delete** /workbooks/{driveItem-id}/thumbnails/{thumbnailSet-id} | Delete navigation property thumbnails for workbooks
[**WorkbooksGetThumbnails**](WorkbooksThumbnailSetApi.md#WorkbooksGetThumbnails) | **Get** /workbooks/{driveItem-id}/thumbnails/{thumbnailSet-id} | Get thumbnails from workbooks
[**WorkbooksListThumbnails**](WorkbooksThumbnailSetApi.md#WorkbooksListThumbnails) | **Get** /workbooks/{driveItem-id}/thumbnails | Get thumbnails from workbooks
[**WorkbooksUpdateThumbnails**](WorkbooksThumbnailSetApi.md#WorkbooksUpdateThumbnails) | **Patch** /workbooks/{driveItem-id}/thumbnails/{thumbnailSet-id} | Update the navigation property thumbnails in workbooks



## WorkbooksCreateThumbnails

> MicrosoftGraphThumbnailSet WorkbooksCreateThumbnails(ctx, driveItemId).MicrosoftGraphThumbnailSet(microsoftGraphThumbnailSet).Execute()

Create new navigation property to thumbnails for workbooks



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
    microsoftGraphThumbnailSet := *openapiclient.NewMicrosoftGraphThumbnailSet() // MicrosoftGraphThumbnailSet | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksThumbnailSetApi.WorkbooksCreateThumbnails(context.Background(), driveItemId).MicrosoftGraphThumbnailSet(microsoftGraphThumbnailSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksThumbnailSetApi.WorkbooksCreateThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksCreateThumbnails`: MicrosoftGraphThumbnailSet
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksThumbnailSetApi.WorkbooksCreateThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksCreateThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphThumbnailSet** | [**MicrosoftGraphThumbnailSet**](MicrosoftGraphThumbnailSet.md) | New navigation property | 

### Return type

[**MicrosoftGraphThumbnailSet**](MicrosoftGraphThumbnailSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDeleteThumbnails

> WorkbooksDeleteThumbnails(ctx, driveItemId, thumbnailSetId).IfMatch(ifMatch).Execute()

Delete navigation property thumbnails for workbooks



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
    thumbnailSetId := "thumbnailSetId_example" // string | key: id of thumbnailSet
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksThumbnailSetApi.WorkbooksDeleteThumbnails(context.Background(), driveItemId, thumbnailSetId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksThumbnailSetApi.WorkbooksDeleteThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**thumbnailSetId** | **string** | key: id of thumbnailSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDeleteThumbnailsRequest struct via the builder pattern


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


## WorkbooksGetThumbnails

> MicrosoftGraphThumbnailSet WorkbooksGetThumbnails(ctx, driveItemId, thumbnailSetId).Select_(select_).Expand(expand).Execute()

Get thumbnails from workbooks



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
    thumbnailSetId := "thumbnailSetId_example" // string | key: id of thumbnailSet
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksThumbnailSetApi.WorkbooksGetThumbnails(context.Background(), driveItemId, thumbnailSetId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksThumbnailSetApi.WorkbooksGetThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetThumbnails`: MicrosoftGraphThumbnailSet
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksThumbnailSetApi.WorkbooksGetThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**thumbnailSetId** | **string** | key: id of thumbnailSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphThumbnailSet**](MicrosoftGraphThumbnailSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListThumbnails

> CollectionOfThumbnailSet WorkbooksListThumbnails(ctx, driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get thumbnails from workbooks



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
    resp, r, err := api_client.WorkbooksThumbnailSetApi.WorkbooksListThumbnails(context.Background(), driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksThumbnailSetApi.WorkbooksListThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListThumbnails`: CollectionOfThumbnailSet
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksThumbnailSetApi.WorkbooksListThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListThumbnailsRequest struct via the builder pattern


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

[**CollectionOfThumbnailSet**](CollectionOfThumbnailSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdateThumbnails

> WorkbooksUpdateThumbnails(ctx, driveItemId, thumbnailSetId).MicrosoftGraphThumbnailSet(microsoftGraphThumbnailSet).Execute()

Update the navigation property thumbnails in workbooks



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
    thumbnailSetId := "thumbnailSetId_example" // string | key: id of thumbnailSet
    microsoftGraphThumbnailSet := *openapiclient.NewMicrosoftGraphThumbnailSet() // MicrosoftGraphThumbnailSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksThumbnailSetApi.WorkbooksUpdateThumbnails(context.Background(), driveItemId, thumbnailSetId).MicrosoftGraphThumbnailSet(microsoftGraphThumbnailSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksThumbnailSetApi.WorkbooksUpdateThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**thumbnailSetId** | **string** | key: id of thumbnailSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphThumbnailSet** | [**MicrosoftGraphThumbnailSet**](MicrosoftGraphThumbnailSet.md) | New navigation property values | 

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

