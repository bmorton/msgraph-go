# \DataPolicyOperationsDataPolicyOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation) | **Post** /dataPolicyOperations | Add new entity to dataPolicyOperations
[**DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation) | **Delete** /dataPolicyOperations/{dataPolicyOperation-id} | Delete entity from dataPolicyOperations
[**DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation) | **Get** /dataPolicyOperations/{dataPolicyOperation-id} | Get entity from dataPolicyOperations by key
[**DataPolicyOperationsDataPolicyOperationListDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationListDataPolicyOperation) | **Get** /dataPolicyOperations | Get entities from dataPolicyOperations
[**DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation**](DataPolicyOperationsDataPolicyOperationApi.md#DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation) | **Patch** /dataPolicyOperations/{dataPolicyOperation-id} | Update entity in dataPolicyOperations



## DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation

> MicrosoftGraphDataPolicyOperation DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation(ctx).MicrosoftGraphDataPolicyOperation(microsoftGraphDataPolicyOperation).Execute()

Add new entity to dataPolicyOperations

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
    microsoftGraphDataPolicyOperation := *openapiclient.NewMicrosoftGraphDataPolicyOperation() // MicrosoftGraphDataPolicyOperation | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation(context.Background()).MicrosoftGraphDataPolicyOperation(microsoftGraphDataPolicyOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation`: MicrosoftGraphDataPolicyOperation
    fmt.Fprintf(os.Stdout, "Response from `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationCreateDataPolicyOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataPolicyOperationsDataPolicyOperationCreateDataPolicyOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDataPolicyOperation** | [**MicrosoftGraphDataPolicyOperation**](MicrosoftGraphDataPolicyOperation.md) | New entity | 

### Return type

[**MicrosoftGraphDataPolicyOperation**](MicrosoftGraphDataPolicyOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation

> DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation(ctx, dataPolicyOperationId).IfMatch(ifMatch).Execute()

Delete entity from dataPolicyOperations

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
    dataPolicyOperationId := "dataPolicyOperationId_example" // string | key: id of dataPolicyOperation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation(context.Background(), dataPolicyOperationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPolicyOperationId** | **string** | key: id of dataPolicyOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataPolicyOperationsDataPolicyOperationDeleteDataPolicyOperationRequest struct via the builder pattern


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


## DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation

> MicrosoftGraphDataPolicyOperation DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation(ctx, dataPolicyOperationId).Select_(select_).Expand(expand).Execute()

Get entity from dataPolicyOperations by key

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
    dataPolicyOperationId := "dataPolicyOperationId_example" // string | key: id of dataPolicyOperation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation(context.Background(), dataPolicyOperationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation`: MicrosoftGraphDataPolicyOperation
    fmt.Fprintf(os.Stdout, "Response from `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationGetDataPolicyOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPolicyOperationId** | **string** | key: id of dataPolicyOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataPolicyOperationsDataPolicyOperationGetDataPolicyOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDataPolicyOperation**](MicrosoftGraphDataPolicyOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataPolicyOperationsDataPolicyOperationListDataPolicyOperation

> CollectionOfDataPolicyOperation DataPolicyOperationsDataPolicyOperationListDataPolicyOperation(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from dataPolicyOperations

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
    resp, r, err := api_client.DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationListDataPolicyOperation(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationListDataPolicyOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataPolicyOperationsDataPolicyOperationListDataPolicyOperation`: CollectionOfDataPolicyOperation
    fmt.Fprintf(os.Stdout, "Response from `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationListDataPolicyOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataPolicyOperationsDataPolicyOperationListDataPolicyOperationRequest struct via the builder pattern


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

[**CollectionOfDataPolicyOperation**](CollectionOfDataPolicyOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation

> DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation(ctx, dataPolicyOperationId).MicrosoftGraphDataPolicyOperation(microsoftGraphDataPolicyOperation).Execute()

Update entity in dataPolicyOperations

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
    dataPolicyOperationId := "dataPolicyOperationId_example" // string | key: id of dataPolicyOperation
    microsoftGraphDataPolicyOperation := *openapiclient.NewMicrosoftGraphDataPolicyOperation() // MicrosoftGraphDataPolicyOperation | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation(context.Background(), dataPolicyOperationId).MicrosoftGraphDataPolicyOperation(microsoftGraphDataPolicyOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPolicyOperationsDataPolicyOperationApi.DataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPolicyOperationId** | **string** | key: id of dataPolicyOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataPolicyOperationsDataPolicyOperationUpdateDataPolicyOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDataPolicyOperation** | [**MicrosoftGraphDataPolicyOperation**](MicrosoftGraphDataPolicyOperation.md) | New property values | 

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

