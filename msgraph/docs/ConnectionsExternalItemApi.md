# \ConnectionsExternalItemApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionsCreateItems**](ConnectionsExternalItemApi.md#ConnectionsCreateItems) | **Post** /connections/{externalConnection-id}/items | Create new navigation property to items for connections
[**ConnectionsDeleteItems**](ConnectionsExternalItemApi.md#ConnectionsDeleteItems) | **Delete** /connections/{externalConnection-id}/items/{externalItem-id} | Delete navigation property items for connections
[**ConnectionsGetItems**](ConnectionsExternalItemApi.md#ConnectionsGetItems) | **Get** /connections/{externalConnection-id}/items/{externalItem-id} | Get items from connections
[**ConnectionsListItems**](ConnectionsExternalItemApi.md#ConnectionsListItems) | **Get** /connections/{externalConnection-id}/items | Get items from connections
[**ConnectionsUpdateItems**](ConnectionsExternalItemApi.md#ConnectionsUpdateItems) | **Patch** /connections/{externalConnection-id}/items/{externalItem-id} | Update the navigation property items in connections



## ConnectionsCreateItems

> MicrosoftGraphExternalConnectorsExternalItem ConnectionsCreateItems(ctx, externalConnectionId).MicrosoftGraphExternalConnectorsExternalItem(microsoftGraphExternalConnectorsExternalItem).Execute()

Create new navigation property to items for connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    microsoftGraphExternalConnectorsExternalItem := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalItem() // MicrosoftGraphExternalConnectorsExternalItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalItemApi.ConnectionsCreateItems(context.Background(), externalConnectionId).MicrosoftGraphExternalConnectorsExternalItem(microsoftGraphExternalConnectorsExternalItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalItemApi.ConnectionsCreateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsCreateItems`: MicrosoftGraphExternalConnectorsExternalItem
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalItemApi.ConnectionsCreateItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsCreateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExternalConnectorsExternalItem** | [**MicrosoftGraphExternalConnectorsExternalItem**](MicrosoftGraphExternalConnectorsExternalItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphExternalConnectorsExternalItem**](MicrosoftGraphExternalConnectorsExternalItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsDeleteItems

> ConnectionsDeleteItems(ctx, externalConnectionId, externalItemId).IfMatch(ifMatch).Execute()

Delete navigation property items for connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalItemId := "externalItemId_example" // string | key: id of externalItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalItemApi.ConnectionsDeleteItems(context.Background(), externalConnectionId, externalItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalItemApi.ConnectionsDeleteItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalItemId** | **string** | key: id of externalItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsDeleteItemsRequest struct via the builder pattern


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


## ConnectionsGetItems

> MicrosoftGraphExternalConnectorsExternalItem ConnectionsGetItems(ctx, externalConnectionId, externalItemId).Select_(select_).Expand(expand).Execute()

Get items from connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalItemId := "externalItemId_example" // string | key: id of externalItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalItemApi.ConnectionsGetItems(context.Background(), externalConnectionId, externalItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalItemApi.ConnectionsGetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGetItems`: MicrosoftGraphExternalConnectorsExternalItem
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalItemApi.ConnectionsGetItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalItemId** | **string** | key: id of externalItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExternalConnectorsExternalItem**](MicrosoftGraphExternalConnectorsExternalItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsListItems

> CollectionOfExternalItem ConnectionsListItems(ctx, externalConnectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get items from connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
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
    resp, r, err := api_client.ConnectionsExternalItemApi.ConnectionsListItems(context.Background(), externalConnectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalItemApi.ConnectionsListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsListItems`: CollectionOfExternalItem
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalItemApi.ConnectionsListItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsListItemsRequest struct via the builder pattern


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

[**CollectionOfExternalItem**](CollectionOfExternalItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsUpdateItems

> ConnectionsUpdateItems(ctx, externalConnectionId, externalItemId).MicrosoftGraphExternalConnectorsExternalItem(microsoftGraphExternalConnectorsExternalItem).Execute()

Update the navigation property items in connections



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
    externalConnectionId := "externalConnectionId_example" // string | key: id of externalConnection
    externalItemId := "externalItemId_example" // string | key: id of externalItem
    microsoftGraphExternalConnectorsExternalItem := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalItem() // MicrosoftGraphExternalConnectorsExternalItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalItemApi.ConnectionsUpdateItems(context.Background(), externalConnectionId, externalItemId).MicrosoftGraphExternalConnectorsExternalItem(microsoftGraphExternalConnectorsExternalItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalItemApi.ConnectionsUpdateItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**externalItemId** | **string** | key: id of externalItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsUpdateItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExternalConnectorsExternalItem** | [**MicrosoftGraphExternalConnectorsExternalItem**](MicrosoftGraphExternalConnectorsExternalItem.md) | New navigation property values | 

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

