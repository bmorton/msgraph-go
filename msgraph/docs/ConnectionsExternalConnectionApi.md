# \ConnectionsExternalConnectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionsExternalConnectionCreateExternalConnection**](ConnectionsExternalConnectionApi.md#ConnectionsExternalConnectionCreateExternalConnection) | **Post** /connections | Add new entity to connections
[**ConnectionsExternalConnectionDeleteExternalConnection**](ConnectionsExternalConnectionApi.md#ConnectionsExternalConnectionDeleteExternalConnection) | **Delete** /connections/{externalConnection-id} | Delete entity from connections
[**ConnectionsExternalConnectionGetExternalConnection**](ConnectionsExternalConnectionApi.md#ConnectionsExternalConnectionGetExternalConnection) | **Get** /connections/{externalConnection-id} | Get entity from connections by key
[**ConnectionsExternalConnectionListExternalConnection**](ConnectionsExternalConnectionApi.md#ConnectionsExternalConnectionListExternalConnection) | **Get** /connections | Get entities from connections
[**ConnectionsExternalConnectionUpdateExternalConnection**](ConnectionsExternalConnectionApi.md#ConnectionsExternalConnectionUpdateExternalConnection) | **Patch** /connections/{externalConnection-id} | Update entity in connections



## ConnectionsExternalConnectionCreateExternalConnection

> MicrosoftGraphExternalConnectorsExternalConnection ConnectionsExternalConnectionCreateExternalConnection(ctx).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()

Add new entity to connections

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
    microsoftGraphExternalConnectorsExternalConnection := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalConnection() // MicrosoftGraphExternalConnectorsExternalConnection | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalConnectionApi.ConnectionsExternalConnectionCreateExternalConnection(context.Background()).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionCreateExternalConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsExternalConnectionCreateExternalConnection`: MicrosoftGraphExternalConnectorsExternalConnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionCreateExternalConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsExternalConnectionCreateExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphExternalConnectorsExternalConnection** | [**MicrosoftGraphExternalConnectorsExternalConnection**](MicrosoftGraphExternalConnectorsExternalConnection.md) | New entity | 

### Return type

[**MicrosoftGraphExternalConnectorsExternalConnection**](MicrosoftGraphExternalConnectorsExternalConnection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsExternalConnectionDeleteExternalConnection

> ConnectionsExternalConnectionDeleteExternalConnection(ctx, externalConnectionId).IfMatch(ifMatch).Execute()

Delete entity from connections

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalConnectionApi.ConnectionsExternalConnectionDeleteExternalConnection(context.Background(), externalConnectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionDeleteExternalConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsExternalConnectionDeleteExternalConnectionRequest struct via the builder pattern


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


## ConnectionsExternalConnectionGetExternalConnection

> MicrosoftGraphExternalConnectorsExternalConnection ConnectionsExternalConnectionGetExternalConnection(ctx, externalConnectionId).Select_(select_).Expand(expand).Execute()

Get entity from connections by key

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalConnectionApi.ConnectionsExternalConnectionGetExternalConnection(context.Background(), externalConnectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionGetExternalConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsExternalConnectionGetExternalConnection`: MicrosoftGraphExternalConnectorsExternalConnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionGetExternalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsExternalConnectionGetExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExternalConnectorsExternalConnection**](MicrosoftGraphExternalConnectorsExternalConnection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsExternalConnectionListExternalConnection

> CollectionOfExternalConnection ConnectionsExternalConnectionListExternalConnection(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from connections

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
    resp, r, err := api_client.ConnectionsExternalConnectionApi.ConnectionsExternalConnectionListExternalConnection(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionListExternalConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsExternalConnectionListExternalConnection`: CollectionOfExternalConnection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionListExternalConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsExternalConnectionListExternalConnectionRequest struct via the builder pattern


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

[**CollectionOfExternalConnection**](CollectionOfExternalConnection.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsExternalConnectionUpdateExternalConnection

> ConnectionsExternalConnectionUpdateExternalConnection(ctx, externalConnectionId).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()

Update entity in connections

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
    microsoftGraphExternalConnectorsExternalConnection := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalConnection() // MicrosoftGraphExternalConnectorsExternalConnection | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsExternalConnectionApi.ConnectionsExternalConnectionUpdateExternalConnection(context.Background(), externalConnectionId).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsExternalConnectionApi.ConnectionsExternalConnectionUpdateExternalConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsExternalConnectionUpdateExternalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExternalConnectorsExternalConnection** | [**MicrosoftGraphExternalConnectorsExternalConnection**](MicrosoftGraphExternalConnectorsExternalConnection.md) | New property values | 

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

