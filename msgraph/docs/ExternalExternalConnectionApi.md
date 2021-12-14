# \ExternalExternalConnectionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalCreateConnections**](ExternalExternalConnectionApi.md#ExternalCreateConnections) | **Post** /external/connections | Create new navigation property to connections for external
[**ExternalDeleteConnections**](ExternalExternalConnectionApi.md#ExternalDeleteConnections) | **Delete** /external/connections/{externalConnection-id} | Delete navigation property connections for external
[**ExternalGetConnections**](ExternalExternalConnectionApi.md#ExternalGetConnections) | **Get** /external/connections/{externalConnection-id} | Get connections from external
[**ExternalListConnections**](ExternalExternalConnectionApi.md#ExternalListConnections) | **Get** /external/connections | Get connections from external
[**ExternalUpdateConnections**](ExternalExternalConnectionApi.md#ExternalUpdateConnections) | **Patch** /external/connections/{externalConnection-id} | Update the navigation property connections in external



## ExternalCreateConnections

> MicrosoftGraphExternalConnectorsExternalConnection ExternalCreateConnections(ctx).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()

Create new navigation property to connections for external

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
    microsoftGraphExternalConnectorsExternalConnection := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalConnection() // MicrosoftGraphExternalConnectorsExternalConnection | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalExternalConnectionApi.ExternalCreateConnections(context.Background()).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalExternalConnectionApi.ExternalCreateConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalCreateConnections`: MicrosoftGraphExternalConnectorsExternalConnection
    fmt.Fprintf(os.Stdout, "Response from `ExternalExternalConnectionApi.ExternalCreateConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalCreateConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphExternalConnectorsExternalConnection** | [**MicrosoftGraphExternalConnectorsExternalConnection**](MicrosoftGraphExternalConnectorsExternalConnection.md) | New navigation property | 

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


## ExternalDeleteConnections

> ExternalDeleteConnections(ctx, externalConnectionId).IfMatch(ifMatch).Execute()

Delete navigation property connections for external

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
    resp, r, err := api_client.ExternalExternalConnectionApi.ExternalDeleteConnections(context.Background(), externalConnectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalExternalConnectionApi.ExternalDeleteConnections``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExternalDeleteConnectionsRequest struct via the builder pattern


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


## ExternalGetConnections

> MicrosoftGraphExternalConnectorsExternalConnection ExternalGetConnections(ctx, externalConnectionId).Select_(select_).Expand(expand).Execute()

Get connections from external

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
    resp, r, err := api_client.ExternalExternalConnectionApi.ExternalGetConnections(context.Background(), externalConnectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalExternalConnectionApi.ExternalGetConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalGetConnections`: MicrosoftGraphExternalConnectorsExternalConnection
    fmt.Fprintf(os.Stdout, "Response from `ExternalExternalConnectionApi.ExternalGetConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalGetConnectionsRequest struct via the builder pattern


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


## ExternalListConnections

> CollectionOfExternalConnection ExternalListConnections(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get connections from external

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
    resp, r, err := api_client.ExternalExternalConnectionApi.ExternalListConnections(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalExternalConnectionApi.ExternalListConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalListConnections`: CollectionOfExternalConnection
    fmt.Fprintf(os.Stdout, "Response from `ExternalExternalConnectionApi.ExternalListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalListConnectionsRequest struct via the builder pattern


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


## ExternalUpdateConnections

> ExternalUpdateConnections(ctx, externalConnectionId).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()

Update the navigation property connections in external

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
    microsoftGraphExternalConnectorsExternalConnection := *openapiclient.NewMicrosoftGraphExternalConnectorsExternalConnection() // MicrosoftGraphExternalConnectorsExternalConnection | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalExternalConnectionApi.ExternalUpdateConnections(context.Background(), externalConnectionId).MicrosoftGraphExternalConnectorsExternalConnection(microsoftGraphExternalConnectorsExternalConnection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalExternalConnectionApi.ExternalUpdateConnections``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExternalUpdateConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExternalConnectorsExternalConnection** | [**MicrosoftGraphExternalConnectorsExternalConnection**](MicrosoftGraphExternalConnectorsExternalConnection.md) | New navigation property values | 

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

