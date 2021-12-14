# \DirectoryDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryCreateDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryCreateDeletedItems) | **Post** /directory/deletedItems | Create new navigation property to deletedItems for directory
[**DirectoryDeleteDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryDeleteDeletedItems) | **Delete** /directory/deletedItems/{directoryObject-id} | Delete navigation property deletedItems for directory
[**DirectoryGetDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryGetDeletedItems) | **Get** /directory/deletedItems/{directoryObject-id} | Get deletedItems from directory
[**DirectoryListDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryListDeletedItems) | **Get** /directory/deletedItems | Get deletedItems from directory
[**DirectoryUpdateDeletedItems**](DirectoryDirectoryObjectApi.md#DirectoryUpdateDeletedItems) | **Patch** /directory/deletedItems/{directoryObject-id} | Update the navigation property deletedItems in directory



## DirectoryCreateDeletedItems

> MicrosoftGraphDirectoryObject DirectoryCreateDeletedItems(ctx).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()

Create new navigation property to deletedItems for directory



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
    microsoftGraphDirectoryObject := *openapiclient.NewMicrosoftGraphDirectoryObject() // MicrosoftGraphDirectoryObject | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryDirectoryObjectApi.DirectoryCreateDeletedItems(context.Background()).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryDirectoryObjectApi.DirectoryCreateDeletedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryCreateDeletedItems`: MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryDirectoryObjectApi.DirectoryCreateDeletedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryCreateDeletedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | New navigation property | 

### Return type

[**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryDeleteDeletedItems

> DirectoryDeleteDeletedItems(ctx, directoryObjectId).IfMatch(ifMatch).Execute()

Delete navigation property deletedItems for directory



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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryDirectoryObjectApi.DirectoryDeleteDeletedItems(context.Background(), directoryObjectId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryDirectoryObjectApi.DirectoryDeleteDeletedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryDeleteDeletedItemsRequest struct via the builder pattern


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


## DirectoryGetDeletedItems

> MicrosoftGraphDirectoryObject DirectoryGetDeletedItems(ctx, directoryObjectId).Select_(select_).Expand(expand).Execute()

Get deletedItems from directory



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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryDirectoryObjectApi.DirectoryGetDeletedItems(context.Background(), directoryObjectId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryDirectoryObjectApi.DirectoryGetDeletedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryGetDeletedItems`: MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryDirectoryObjectApi.DirectoryGetDeletedItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryGetDeletedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryListDeletedItems

> CollectionOfDirectoryObject DirectoryListDeletedItems(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get deletedItems from directory



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
    resp, r, err := api_client.DirectoryDirectoryObjectApi.DirectoryListDeletedItems(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryDirectoryObjectApi.DirectoryListDeletedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryListDeletedItems`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryDirectoryObjectApi.DirectoryListDeletedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryListDeletedItemsRequest struct via the builder pattern


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

[**CollectionOfDirectoryObject**](CollectionOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryUpdateDeletedItems

> DirectoryUpdateDeletedItems(ctx, directoryObjectId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()

Update the navigation property deletedItems in directory



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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    microsoftGraphDirectoryObject := *openapiclient.NewMicrosoftGraphDirectoryObject() // MicrosoftGraphDirectoryObject | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryDirectoryObjectApi.DirectoryUpdateDeletedItems(context.Background(), directoryObjectId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryDirectoryObjectApi.DirectoryUpdateDeletedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryUpdateDeletedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | New navigation property values | 

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

