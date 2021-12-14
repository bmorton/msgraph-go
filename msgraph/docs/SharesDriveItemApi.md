# \SharesDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesCreateItems**](SharesDriveItemApi.md#SharesCreateItems) | **Post** /shares/{sharedDriveItem-id}/items | Create new navigation property to items for shares
[**SharesDeleteDriveItem**](SharesDriveItemApi.md#SharesDeleteDriveItem) | **Delete** /shares/{sharedDriveItem-id}/driveItem | Delete navigation property driveItem for shares
[**SharesDeleteItems**](SharesDriveItemApi.md#SharesDeleteItems) | **Delete** /shares/{sharedDriveItem-id}/items/{driveItem-id} | Delete navigation property items for shares
[**SharesDeleteRoot**](SharesDriveItemApi.md#SharesDeleteRoot) | **Delete** /shares/{sharedDriveItem-id}/root | Delete navigation property root for shares
[**SharesGetDriveItem**](SharesDriveItemApi.md#SharesGetDriveItem) | **Get** /shares/{sharedDriveItem-id}/driveItem | Get driveItem from shares
[**SharesGetDriveItemContent**](SharesDriveItemApi.md#SharesGetDriveItemContent) | **Get** /shares/{sharedDriveItem-id}/driveItem/content | Get media content for the navigation property driveItem from shares
[**SharesGetItems**](SharesDriveItemApi.md#SharesGetItems) | **Get** /shares/{sharedDriveItem-id}/items/{driveItem-id} | Get items from shares
[**SharesGetItemsContent**](SharesDriveItemApi.md#SharesGetItemsContent) | **Get** /shares/{sharedDriveItem-id}/items/{driveItem-id}/content | Get media content for the navigation property items from shares
[**SharesGetRoot**](SharesDriveItemApi.md#SharesGetRoot) | **Get** /shares/{sharedDriveItem-id}/root | Get root from shares
[**SharesGetRootContent**](SharesDriveItemApi.md#SharesGetRootContent) | **Get** /shares/{sharedDriveItem-id}/root/content | Get media content for the navigation property root from shares
[**SharesListItems**](SharesDriveItemApi.md#SharesListItems) | **Get** /shares/{sharedDriveItem-id}/items | Get items from shares
[**SharesUpdateDriveItem**](SharesDriveItemApi.md#SharesUpdateDriveItem) | **Patch** /shares/{sharedDriveItem-id}/driveItem | Update the navigation property driveItem in shares
[**SharesUpdateDriveItemContent**](SharesDriveItemApi.md#SharesUpdateDriveItemContent) | **Put** /shares/{sharedDriveItem-id}/driveItem/content | Update media content for the navigation property driveItem in shares
[**SharesUpdateItems**](SharesDriveItemApi.md#SharesUpdateItems) | **Patch** /shares/{sharedDriveItem-id}/items/{driveItem-id} | Update the navigation property items in shares
[**SharesUpdateItemsContent**](SharesDriveItemApi.md#SharesUpdateItemsContent) | **Put** /shares/{sharedDriveItem-id}/items/{driveItem-id}/content | Update media content for the navigation property items in shares
[**SharesUpdateRoot**](SharesDriveItemApi.md#SharesUpdateRoot) | **Patch** /shares/{sharedDriveItem-id}/root | Update the navigation property root in shares
[**SharesUpdateRootContent**](SharesDriveItemApi.md#SharesUpdateRootContent) | **Put** /shares/{sharedDriveItem-id}/root/content | Update media content for the navigation property root in shares



## SharesCreateItems

> MicrosoftGraphDriveItem SharesCreateItems(ctx, sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Create new navigation property to items for shares



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesCreateItems(context.Background(), sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesCreateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesCreateItems`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesCreateItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesCreateItemsRequest struct via the builder pattern


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


## SharesDeleteDriveItem

> SharesDeleteDriveItem(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete navigation property driveItem for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesDeleteDriveItem(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesDeleteDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesDeleteDriveItemRequest struct via the builder pattern


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


## SharesDeleteItems

> SharesDeleteItems(ctx, sharedDriveItemId, driveItemId).IfMatch(ifMatch).Execute()

Delete navigation property items for shares



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesDeleteItems(context.Background(), sharedDriveItemId, driveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesDeleteItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesDeleteItemsRequest struct via the builder pattern


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


## SharesDeleteRoot

> SharesDeleteRoot(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete navigation property root for shares



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesDeleteRoot(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesDeleteRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesDeleteRootRequest struct via the builder pattern


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


## SharesGetDriveItem

> MicrosoftGraphDriveItem SharesGetDriveItem(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get driveItem from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesGetDriveItem(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesGetDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetDriveItem`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesGetDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetDriveItemRequest struct via the builder pattern


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


## SharesGetDriveItemContent

> *os.File SharesGetDriveItemContent(ctx, sharedDriveItemId).Execute()

Get media content for the navigation property driveItem from shares

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
    resp, r, err := api_client.SharesDriveItemApi.SharesGetDriveItemContent(context.Background(), sharedDriveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesGetDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetDriveItemContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesGetDriveItemContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetDriveItemContentRequest struct via the builder pattern


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


## SharesGetItems

> MicrosoftGraphDriveItem SharesGetItems(ctx, sharedDriveItemId, driveItemId).Select_(select_).Expand(expand).Execute()

Get items from shares



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesGetItems(context.Background(), sharedDriveItemId, driveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesGetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetItems`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesGetItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetItemsRequest struct via the builder pattern


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


## SharesGetItemsContent

> *os.File SharesGetItemsContent(ctx, sharedDriveItemId, driveItemId).Execute()

Get media content for the navigation property items from shares

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
    driveItemId := "driveItemId_example" // string | key: id of driveItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesGetItemsContent(context.Background(), sharedDriveItemId, driveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesGetItemsContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetItemsContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesGetItemsContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetItemsContentRequest struct via the builder pattern


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


## SharesGetRoot

> MicrosoftGraphDriveItem SharesGetRoot(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get root from shares



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesGetRoot(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesGetRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetRoot`: MicrosoftGraphDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesGetRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetRootRequest struct via the builder pattern


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


## SharesGetRootContent

> *os.File SharesGetRootContent(ctx, sharedDriveItemId).Execute()

Get media content for the navigation property root from shares

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
    resp, r, err := api_client.SharesDriveItemApi.SharesGetRootContent(context.Background(), sharedDriveItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesGetRootContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesGetRootContent`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesGetRootContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesGetRootContentRequest struct via the builder pattern


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


## SharesListItems

> CollectionOfDriveItem SharesListItems(ctx, sharedDriveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get items from shares



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
    resp, r, err := api_client.SharesDriveItemApi.SharesListItems(context.Background(), sharedDriveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesListItems`: CollectionOfDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesDriveItemApi.SharesListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesListItemsRequest struct via the builder pattern


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


## SharesUpdateDriveItem

> SharesUpdateDriveItem(ctx, sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property driveItem in shares



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesUpdateDriveItem(context.Background(), sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesUpdateDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesUpdateDriveItemRequest struct via the builder pattern


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


## SharesUpdateDriveItemContent

> SharesUpdateDriveItemContent(ctx, sharedDriveItemId).Body(body).Execute()

Update media content for the navigation property driveItem in shares

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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesUpdateDriveItemContent(context.Background(), sharedDriveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesUpdateDriveItemContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesUpdateDriveItemContentRequest struct via the builder pattern


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


## SharesUpdateItems

> SharesUpdateItems(ctx, sharedDriveItemId, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property items in shares



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesUpdateItems(context.Background(), sharedDriveItemId, driveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesUpdateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesUpdateItemsRequest struct via the builder pattern


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


## SharesUpdateItemsContent

> SharesUpdateItemsContent(ctx, sharedDriveItemId, driveItemId).Body(body).Execute()

Update media content for the navigation property items in shares

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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesUpdateItemsContent(context.Background(), sharedDriveItemId, driveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesUpdateItemsContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesUpdateItemsContentRequest struct via the builder pattern


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


## SharesUpdateRoot

> SharesUpdateRoot(ctx, sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()

Update the navigation property root in shares



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
    microsoftGraphDriveItem := *openapiclient.NewMicrosoftGraphDriveItem() // MicrosoftGraphDriveItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesUpdateRoot(context.Background(), sharedDriveItemId).MicrosoftGraphDriveItem(microsoftGraphDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesUpdateRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesUpdateRootRequest struct via the builder pattern


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


## SharesUpdateRootContent

> SharesUpdateRootContent(ctx, sharedDriveItemId).Body(body).Execute()

Update media content for the navigation property root in shares

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
    body := os.NewFile(1234, "some_file") // *os.File | New media content.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesDriveItemApi.SharesUpdateRootContent(context.Background(), sharedDriveItemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesDriveItemApi.SharesUpdateRootContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesUpdateRootContentRequest struct via the builder pattern


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

