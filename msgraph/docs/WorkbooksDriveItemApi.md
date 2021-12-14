# \WorkbooksDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreateChildren**](WorkbooksDriveItemApi.md#WorkbooksCreateChildren) | **Post** /workbooks/{driveItem-id}/children | Create new navigation property to children for workbooks
[**WorkbooksDeleteChildren**](WorkbooksDriveItemApi.md#WorkbooksDeleteChildren) | **Delete** /workbooks/{driveItem-id}/children/{driveItem-id1} | Delete navigation property children for workbooks
[**WorkbooksDriveItemCreateDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemCreateDriveItem) | **Post** /workbooks | Add new entity to workbooks
[**WorkbooksDriveItemDeleteDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemDeleteDriveItem) | **Delete** /workbooks/{driveItem-id} | Delete entity from workbooks
[**WorkbooksDriveItemGetContent**](WorkbooksDriveItemApi.md#WorkbooksDriveItemGetContent) | **Get** /workbooks/{driveItem-id}/content | Get media content for driveItem from workbooks
[**WorkbooksDriveItemGetDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemGetDriveItem) | **Get** /workbooks/{driveItem-id} | Get entity from workbooks by key
[**WorkbooksDriveItemListDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemListDriveItem) | **Get** /workbooks | Get entities from workbooks
[**WorkbooksDriveItemUpdateContent**](WorkbooksDriveItemApi.md#WorkbooksDriveItemUpdateContent) | **Put** /workbooks/{driveItem-id}/content | Update media content for driveItem in workbooks
[**WorkbooksDriveItemUpdateDriveItem**](WorkbooksDriveItemApi.md#WorkbooksDriveItemUpdateDriveItem) | **Patch** /workbooks/{driveItem-id} | Update entity in workbooks
[**WorkbooksGetChildren**](WorkbooksDriveItemApi.md#WorkbooksGetChildren) | **Get** /workbooks/{driveItem-id}/children/{driveItem-id1} | Get children from workbooks
[**WorkbooksGetChildrenContent**](WorkbooksDriveItemApi.md#WorkbooksGetChildrenContent) | **Get** /workbooks/{driveItem-id}/children/{driveItem-id1}/content | Get media content for the navigation property children from workbooks
[**WorkbooksListChildren**](WorkbooksDriveItemApi.md#WorkbooksListChildren) | **Get** /workbooks/{driveItem-id}/children | Get children from workbooks
[**WorkbooksUpdateChildren**](WorkbooksDriveItemApi.md#WorkbooksUpdateChildren) | **Patch** /workbooks/{driveItem-id}/children/{driveItem-id1} | Update the navigation property children in workbooks
[**WorkbooksUpdateChildrenContent**](WorkbooksDriveItemApi.md#WorkbooksUpdateChildrenContent) | **Put** /workbooks/{driveItem-id}/children/{driveItem-id1}/content | Update media content for the navigation property children in workbooks



## WorkbooksCreateChildren

> MicrosoftGraphDriveItem WorkbooksCreateChildren(ctx, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Create new navigation property to children for workbooks



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksCreateChildren(context.Background(), driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksCreateChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksCreateChildren`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksCreateChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksCreateChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDeleteChildren

> WorkbooksDeleteChildren(ctx, driveItemId, driveItemId1).IfMatch(ifMatch).Execute()

Delete navigation property children for workbooks



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
    driveItemId1 := "driveItemId1_example" // string | key: id of driveItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDeleteChildren(context.Background(), driveItemId, driveItemId1).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDeleteChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemId1** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDeleteChildrenRequest struct via the builder pattern


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


## WorkbooksDriveItemCreateDriveItem

> MicrosoftGraphDriveItem WorkbooksDriveItemCreateDriveItem(ctx).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Add new entity to workbooks

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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDriveItemCreateDriveItem(context.Background()).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDriveItemCreateDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksDriveItemCreateDriveItem`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksDriveItemCreateDriveItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDriveItemCreateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New entity | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemDeleteDriveItem

> WorkbooksDriveItemDeleteDriveItem(ctx, driveItemId).IfMatch(ifMatch).Execute()

Delete entity from workbooks

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
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDriveItemDeleteDriveItem(context.Background(), driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDriveItemDeleteDriveItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkbooksDriveItemDeleteDriveItemRequest struct via the builder pattern


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


## WorkbooksDriveItemGetContent

> *os.File WorkbooksDriveItemGetContent(ctx, driveItemId).Execute()

Get media content for driveItem from workbooks



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
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDriveItemGetContent(context.Background(), driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDriveItemGetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksDriveItemGetContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksDriveItemGetContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDriveItemGetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemGetDriveItem

> MicrosoftGraphDriveItem WorkbooksDriveItemGetDriveItem(ctx, driveItemId).Select_(select_).Expand(expand).Execute()

Get entity from workbooks by key

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
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDriveItemGetDriveItem(context.Background(), driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDriveItemGetDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksDriveItemGetDriveItem`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksDriveItemGetDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDriveItemGetDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemListDriveItem

> CollectionOfDriveItem WorkbooksDriveItemListDriveItem(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from workbooks

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
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDriveItemListDriveItem(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDriveItemListDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksDriveItemListDriveItem`: CollectionOfDriveItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksDriveItemListDriveItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDriveItemListDriveItemRequest struct via the builder pattern


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

[**CollectionOfDriveItem**](CollectionOfDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemUpdateContent

> WorkbooksDriveItemUpdateContent(ctx, driveItemId).Body(body).Execute()

Update media content for driveItem in workbooks



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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDriveItemUpdateContent(context.Background(), driveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDriveItemUpdateContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkbooksDriveItemUpdateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | New media content. | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksDriveItemUpdateDriveItem

> WorkbooksDriveItemUpdateDriveItem(ctx, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update entity in workbooks

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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksDriveItemUpdateDriveItem(context.Background(), driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksDriveItemUpdateDriveItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWorkbooksDriveItemUpdateDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New property values | 

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


## WorkbooksGetChildren

> MicrosoftGraphDriveItem WorkbooksGetChildren(ctx, driveItemId, driveItemId1).Select_(select_).Expand(expand).Execute()

Get children from workbooks



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
    driveItemId1 := "driveItemId1_example" // string | key: id of driveItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksGetChildren(context.Background(), driveItemId, driveItemId1).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksGetChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetChildren`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksGetChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemId1** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksGetChildrenContent

> *os.File WorkbooksGetChildrenContent(ctx, driveItemId, driveItemId1).Execute()

Get media content for the navigation property children from workbooks



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
    driveItemId1 := "driveItemId1_example" // string | key: id of driveItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksGetChildrenContent(context.Background(), driveItemId, driveItemId1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksGetChildrenContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetChildrenContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksGetChildrenContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemId1** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetChildrenContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksListChildren

> CollectionOfDriveItem WorkbooksListChildren(ctx, driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get children from workbooks



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
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksListChildren(context.Background(), driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksListChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListChildren`: CollectionOfDriveItem
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksDriveItemApi.WorkbooksListChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListChildrenRequest struct via the builder pattern


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

[**CollectionOfDriveItem**](CollectionOfDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdateChildren

> WorkbooksUpdateChildren(ctx, driveItemId, driveItemId1).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property children in workbooks



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
    driveItemId1 := "driveItemId1_example" // string | key: id of driveItem
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksUpdateChildren(context.Background(), driveItemId, driveItemId1).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksUpdateChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemId1** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDriveItem** | [**MicrosoftGraphDriveItem**](MicrosoftGraphDriveItem.md) | New navigation property values | 

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


## WorkbooksUpdateChildrenContent

> WorkbooksUpdateChildrenContent(ctx, driveItemId, driveItemId1).Body(body).Execute()

Update media content for the navigation property children in workbooks



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
    driveItemId1 := "driveItemId1_example" // string | key: id of driveItem
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksDriveItemApi.WorkbooksUpdateChildrenContent(context.Background(), driveItemId, driveItemId1).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksDriveItemApi.WorkbooksUpdateChildrenContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**driveItemId1** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateChildrenContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | New media content. | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

