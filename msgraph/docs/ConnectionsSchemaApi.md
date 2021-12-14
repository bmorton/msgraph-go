# \ConnectionsSchemaApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionsDeleteSchema**](ConnectionsSchemaApi.md#ConnectionsDeleteSchema) | **Delete** /connections/{externalConnection-id}/schema | Delete navigation property schema for connections
[**ConnectionsGetSchema**](ConnectionsSchemaApi.md#ConnectionsGetSchema) | **Get** /connections/{externalConnection-id}/schema | Get schema from connections
[**ConnectionsUpdateSchema**](ConnectionsSchemaApi.md#ConnectionsUpdateSchema) | **Patch** /connections/{externalConnection-id}/schema | Update the navigation property schema in connections



## ConnectionsDeleteSchema

> ConnectionsDeleteSchema(ctx, externalConnectionId).IfMatch(ifMatch).Execute()

Delete navigation property schema for connections



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
    resp, r, err := api_client.ConnectionsSchemaApi.ConnectionsDeleteSchema(context.Background(), externalConnectionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSchemaApi.ConnectionsDeleteSchema``: %v\n", err)
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

Other parameters are passed through a pointer to a apiConnectionsDeleteSchemaRequest struct via the builder pattern


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


## ConnectionsGetSchema

> MicrosoftGraphExternalConnectorsSchema ConnectionsGetSchema(ctx, externalConnectionId).Select_(select_).Expand(expand).Execute()

Get schema from connections



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
    resp, r, err := api_client.ConnectionsSchemaApi.ConnectionsGetSchema(context.Background(), externalConnectionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSchemaApi.ConnectionsGetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGetSchema`: MicrosoftGraphExternalConnectorsSchema
    fmt.Fprintf(os.Stdout, "Response from `ConnectionsSchemaApi.ConnectionsGetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalConnectionId** | **string** | key: id of externalConnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExternalConnectorsSchema**](MicrosoftGraphExternalConnectorsSchema.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsUpdateSchema

> ConnectionsUpdateSchema(ctx, externalConnectionId).MicrosoftGraphExternalConnectorsSchema(microsoftGraphExternalConnectorsSchema).Execute()

Update the navigation property schema in connections



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
    microsoftGraphExternalConnectorsSchema := *openapiclient.NewMicrosoftGraphExternalConnectorsSchema() // MicrosoftGraphExternalConnectorsSchema | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionsSchemaApi.ConnectionsUpdateSchema(context.Background(), externalConnectionId).MicrosoftGraphExternalConnectorsSchema(microsoftGraphExternalConnectorsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsSchemaApi.ConnectionsUpdateSchema``: %v\n", err)
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

Other parameters are passed through a pointer to a apiConnectionsUpdateSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExternalConnectorsSchema** | [**MicrosoftGraphExternalConnectorsSchema**](MicrosoftGraphExternalConnectorsSchema.md) | New navigation property values | 

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

