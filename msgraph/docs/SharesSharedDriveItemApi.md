# \SharesSharedDriveItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharesSharedDriveItemCreateSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemCreateSharedDriveItem) | **Post** /shares | Add new entity to shares
[**SharesSharedDriveItemDeleteSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemDeleteSharedDriveItem) | **Delete** /shares/{sharedDriveItem-id} | Delete entity from shares
[**SharesSharedDriveItemGetSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemGetSharedDriveItem) | **Get** /shares/{sharedDriveItem-id} | Get entity from shares by key
[**SharesSharedDriveItemListSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemListSharedDriveItem) | **Get** /shares | Get entities from shares
[**SharesSharedDriveItemUpdateSharedDriveItem**](SharesSharedDriveItemApi.md#SharesSharedDriveItemUpdateSharedDriveItem) | **Patch** /shares/{sharedDriveItem-id} | Update entity in shares



## SharesSharedDriveItemCreateSharedDriveItem

> MicrosoftGraphSharedDriveItem SharesSharedDriveItemCreateSharedDriveItem(ctx).MicrosoftGraphSharedDriveItem(microsoftGraphSharedDriveItem).Execute()

Add new entity to shares

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
    microsoftGraphSharedDriveItem := *openapiclient.NewMicrosoftGraphSharedDriveItem() // MicrosoftGraphSharedDriveItem | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesSharedDriveItemApi.SharesSharedDriveItemCreateSharedDriveItem(context.Background()).MicrosoftGraphSharedDriveItem(microsoftGraphSharedDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesSharedDriveItemApi.SharesSharedDriveItemCreateSharedDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemCreateSharedDriveItem`: MicrosoftGraphSharedDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesSharedDriveItemApi.SharesSharedDriveItemCreateSharedDriveItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemCreateSharedDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSharedDriveItem** | [**MicrosoftGraphSharedDriveItem**](MicrosoftGraphSharedDriveItem.md) | New entity | 

### Return type

[**MicrosoftGraphSharedDriveItem**](MicrosoftGraphSharedDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemDeleteSharedDriveItem

> SharesSharedDriveItemDeleteSharedDriveItem(ctx, sharedDriveItemId).IfMatch(ifMatch).Execute()

Delete entity from shares

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
    resp, r, err := api_client.SharesSharedDriveItemApi.SharesSharedDriveItemDeleteSharedDriveItem(context.Background(), sharedDriveItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesSharedDriveItemApi.SharesSharedDriveItemDeleteSharedDriveItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesSharedDriveItemDeleteSharedDriveItemRequest struct via the builder pattern


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


## SharesSharedDriveItemGetSharedDriveItem

> MicrosoftGraphSharedDriveItem SharesSharedDriveItemGetSharedDriveItem(ctx, sharedDriveItemId).Select_(select_).Expand(expand).Execute()

Get entity from shares by key

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
    resp, r, err := api_client.SharesSharedDriveItemApi.SharesSharedDriveItemGetSharedDriveItem(context.Background(), sharedDriveItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesSharedDriveItemApi.SharesSharedDriveItemGetSharedDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemGetSharedDriveItem`: MicrosoftGraphSharedDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesSharedDriveItemApi.SharesSharedDriveItemGetSharedDriveItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedDriveItemId** | **string** | key: id of sharedDriveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemGetSharedDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSharedDriveItem**](MicrosoftGraphSharedDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemListSharedDriveItem

> CollectionOfSharedDriveItem SharesSharedDriveItemListSharedDriveItem(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from shares

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
    resp, r, err := api_client.SharesSharedDriveItemApi.SharesSharedDriveItemListSharedDriveItem(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesSharedDriveItemApi.SharesSharedDriveItemListSharedDriveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharesSharedDriveItemListSharedDriveItem`: CollectionOfSharedDriveItem
    fmt.Fprintf(os.Stdout, "Response from `SharesSharedDriveItemApi.SharesSharedDriveItemListSharedDriveItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSharesSharedDriveItemListSharedDriveItemRequest struct via the builder pattern


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

[**CollectionOfSharedDriveItem**](CollectionOfSharedDriveItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharesSharedDriveItemUpdateSharedDriveItem

> SharesSharedDriveItemUpdateSharedDriveItem(ctx, sharedDriveItemId).MicrosoftGraphSharedDriveItem(microsoftGraphSharedDriveItem).Execute()

Update entity in shares

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
    microsoftGraphSharedDriveItem := *openapiclient.NewMicrosoftGraphSharedDriveItem() // MicrosoftGraphSharedDriveItem | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharesSharedDriveItemApi.SharesSharedDriveItemUpdateSharedDriveItem(context.Background(), sharedDriveItemId).MicrosoftGraphSharedDriveItem(microsoftGraphSharedDriveItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharesSharedDriveItemApi.SharesSharedDriveItemUpdateSharedDriveItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSharesSharedDriveItemUpdateSharedDriveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSharedDriveItem** | [**MicrosoftGraphSharedDriveItem**](MicrosoftGraphSharedDriveItem.md) | New property values | 

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

