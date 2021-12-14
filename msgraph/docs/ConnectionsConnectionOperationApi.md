# \ConnectionsConnectionOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionsCreateOperations**](ConnectionsConnectionOperationApi.md#ConnectionsCreateOperations) | **Post** /connections/{externalConnection-id}/operations | Create new navigation property to operations for connections
[**ConnectionsDeleteOperations**](ConnectionsConnectionOperationApi.md#ConnectionsDeleteOperations) | **Delete** /connections/{externalConnection-id}/operations/{connectionOperation-id} | Delete navigation property operations for connections
[**ConnectionsGetOperations**](ConnectionsConnectionOperationApi.md#ConnectionsGetOperations) | **Get** /connections/{externalConnection-id}/operations/{connectionOperation-id} | Get operations from connections
[**ConnectionsListOperations**](ConnectionsConnectionOperationApi.md#ConnectionsListOperations) | **Get** /connections/{externalConnection-id}/operations | Get operations from connections
[**ConnectionsUpdateOperations**](ConnectionsConnectionOperationApi.md#ConnectionsUpdateOperations) | **Patch** /connections/{externalConnection-id}/operations/{connectionOperation-id} | Update the navigation property operations in connections



## ConnectionsCreateOperations

> MicrosoftGraphExternalConnectorsConnectionOperation ConnectionsCreateOperations(ctx, externalConnectionId).MicrosoftGraphExternalConnectorsConnectionOperation(microsoftGraphExternalConnectorsConnectionOperation).Execute()

Create new navigation property to operations for connections



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
    microsoftGraphExternalConnectorsConnectionOperation := *openapiclient.NewMicrosoftGraphExternalConnectorsConnectionOperation() // MicrosoftGraphExternalConnectorsConnectionOperation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsConnectionOperationApi.ConnectionsCreateOperations(context.Background(), externalConnectionId).MicrosoftGraphExternalConnectorsConnectionOperation(microsoftGraphExternalConnectorsConnectionOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsConnectionOperationApi.ConnectionsCreateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsCreateOperations`: MicrosoftGraphExternalConnectorsConnectionOperation
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsConnectionOperationApi.ConnectionsCreateOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsCreateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExternalConnectorsConnectionOperation** | [**MicrosoftGraphExternalConnectorsConnectionOperation**](MicrosoftGraphExternalConnectorsConnectionOperation.md) | New navigation property | 

### Return type

[**MicrosoftGraphExternalConnectorsConnectionOperation**](MicrosoftGraphExternalConnectorsConnectionOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsDeleteOperations

> ConnectionsDeleteOperations(ctx, externalConnectionId, connectionOperationId).IfMatch(ifMatch).Execute()

Delete navigation property operations for connections



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
    connectionOperationId := "connectionOperationId_example" // string | key: id of connectionOperation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsConnectionOperationApi.ConnectionsDeleteOperations(context.Background(), externalConnectionId, connectionOperationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsConnectionOperationApi.ConnectionsDeleteOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**connectionOperationId** | **string** | key: id of connectionOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsDeleteOperationsRequest struct via the builder pattern


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


## ConnectionsGetOperations

> MicrosoftGraphExternalConnectorsConnectionOperation ConnectionsGetOperations(ctx, externalConnectionId, connectionOperationId).Select_(select_).Expand(expand).Execute()

Get operations from connections



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
    connectionOperationId := "connectionOperationId_example" // string | key: id of connectionOperation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsConnectionOperationApi.ConnectionsGetOperations(context.Background(), externalConnectionId, connectionOperationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsConnectionOperationApi.ConnectionsGetOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGetOperations`: MicrosoftGraphExternalConnectorsConnectionOperation
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsConnectionOperationApi.ConnectionsGetOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**connectionOperationId** | **string** | key: id of connectionOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGetOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExternalConnectorsConnectionOperation**](MicrosoftGraphExternalConnectorsConnectionOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsListOperations

> CollectionOfConnectionOperation ConnectionsListOperations(ctx, externalConnectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get operations from connections



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
    resp, r, err := api_client.ConnectionsConnectionOperationApi.ConnectionsListOperations(context.Background(), externalConnectionId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsConnectionOperationApi.ConnectionsListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsListOperations`: CollectionOfConnectionOperation
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsConnectionOperationApi.ConnectionsListOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsListOperationsRequest struct via the builder pattern


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

[**CollectionOfConnectionOperation**](CollectionOfConnectionOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsUpdateOperations

> ConnectionsUpdateOperations(ctx, externalConnectionId, connectionOperationId).MicrosoftGraphExternalConnectorsConnectionOperation(microsoftGraphExternalConnectorsConnectionOperation).Execute()

Update the navigation property operations in connections



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
    connectionOperationId := "connectionOperationId_example" // string | key: id of connectionOperation
    microsoftGraphExternalConnectorsConnectionOperation := *openapiclient.NewMicrosoftGraphExternalConnectorsConnectionOperation() // MicrosoftGraphExternalConnectorsConnectionOperation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsConnectionOperationApi.ConnectionsUpdateOperations(context.Background(), externalConnectionId, connectionOperationId).MicrosoftGraphExternalConnectorsConnectionOperation(microsoftGraphExternalConnectorsConnectionOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsConnectionOperationApi.ConnectionsUpdateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 
**connectionOperationId** | **string** | key: id of connectionOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsUpdateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphExternalConnectorsConnectionOperation** | [**MicrosoftGraphExternalConnectorsConnectionOperation**](MicrosoftGraphExternalConnectorsConnectionOperation.md) | New navigation property values | 

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

