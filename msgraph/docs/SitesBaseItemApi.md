# \SitesBaseItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesCreateItems**](SitesBaseItemApi.md#SitesCreateItems) | **Post** /sites/{site-id}/items | Create new navigation property to items for sites
[**SitesDeleteItems**](SitesBaseItemApi.md#SitesDeleteItems) | **Delete** /sites/{site-id}/items/{baseItem-id} | Delete navigation property items for sites
[**SitesGetItems**](SitesBaseItemApi.md#SitesGetItems) | **Get** /sites/{site-id}/items/{baseItem-id} | Get items from sites
[**SitesListItems**](SitesBaseItemApi.md#SitesListItems) | **Get** /sites/{site-id}/items | Get items from sites
[**SitesUpdateItems**](SitesBaseItemApi.md#SitesUpdateItems) | **Patch** /sites/{site-id}/items/{baseItem-id} | Update the navigation property items in sites



## SitesCreateItems

> MicrosoftGraphBaseItem SitesCreateItems(ctx, siteId).MicrosoftGraphBaseItem(microsoftGraphBaseItem).Execute()

Create new navigation property to items for sites



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
    microsoftGraphBaseItem := *openapiclient.NewMicrosoftGraphBaseItem() // MicrosoftGraphBaseItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesBaseItemApi.SitesCreateItems(context.Background(), siteId).MicrosoftGraphBaseItem(microsoftGraphBaseItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesBaseItemApi.SitesCreateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesCreateItems`: MicrosoftGraphBaseItem
    fmt.Fprintf(os.Stdout, "Response from `SitesBaseItemApi.SitesCreateItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesCreateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphBaseItem** | [**MicrosoftGraphBaseItem**](MicrosoftGraphBaseItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphBaseItem**](MicrosoftGraphBaseItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesDeleteItems

> SitesDeleteItems(ctx, siteId, baseItemId).IfMatch(ifMatch).Execute()

Delete navigation property items for sites



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
    baseItemId := "baseItemId_example" // string | key: id of baseItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesBaseItemApi.SitesDeleteItems(context.Background(), siteId, baseItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesBaseItemApi.SitesDeleteItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**baseItemId** | **string** | key: id of baseItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesDeleteItemsRequest struct via the builder pattern


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


## SitesGetItems

> MicrosoftGraphBaseItem SitesGetItems(ctx, siteId, baseItemId).Select_(select_).Expand(expand).Execute()

Get items from sites



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
    baseItemId := "baseItemId_example" // string | key: id of baseItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesBaseItemApi.SitesGetItems(context.Background(), siteId, baseItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesBaseItemApi.SitesGetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesGetItems`: MicrosoftGraphBaseItem
    fmt.Fprintf(os.Stdout, "Response from `SitesBaseItemApi.SitesGetItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**baseItemId** | **string** | key: id of baseItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphBaseItem**](MicrosoftGraphBaseItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesListItems

> CollectionOfBaseItem SitesListItems(ctx, siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get items from sites



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
    resp, r, err := api_client.SitesBaseItemApi.SitesListItems(context.Background(), siteId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesBaseItemApi.SitesListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SitesListItems`: CollectionOfBaseItem
    fmt.Fprintf(os.Stdout, "Response from `SitesBaseItemApi.SitesListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesListItemsRequest struct via the builder pattern


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

[**CollectionOfBaseItem**](CollectionOfBaseItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SitesUpdateItems

> SitesUpdateItems(ctx, siteId, baseItemId).MicrosoftGraphBaseItem(microsoftGraphBaseItem).Execute()

Update the navigation property items in sites



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
    baseItemId := "baseItemId_example" // string | key: id of baseItem
    microsoftGraphBaseItem := *openapiclient.NewMicrosoftGraphBaseItem() // MicrosoftGraphBaseItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SitesBaseItemApi.SitesUpdateItems(context.Background(), siteId, baseItemId).MicrosoftGraphBaseItem(microsoftGraphBaseItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SitesBaseItemApi.SitesUpdateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | key: id of site | 
**baseItemId** | **string** | key: id of baseItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiSitesUpdateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphBaseItem** | [**MicrosoftGraphBaseItem**](MicrosoftGraphBaseItem.md) | New navigation property values | 

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

