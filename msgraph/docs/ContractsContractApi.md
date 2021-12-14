# \ContractsContractApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractsContractCreateContract**](ContractsContractApi.md#ContractsContractCreateContract) | **Post** /contracts | Add new entity to contracts
[**ContractsContractDeleteContract**](ContractsContractApi.md#ContractsContractDeleteContract) | **Delete** /contracts/{contract-id} | Delete entity from contracts
[**ContractsContractGetContract**](ContractsContractApi.md#ContractsContractGetContract) | **Get** /contracts/{contract-id} | Get entity from contracts by key
[**ContractsContractListContract**](ContractsContractApi.md#ContractsContractListContract) | **Get** /contracts | Get entities from contracts
[**ContractsContractUpdateContract**](ContractsContractApi.md#ContractsContractUpdateContract) | **Patch** /contracts/{contract-id} | Update entity in contracts



## ContractsContractCreateContract

> MicrosoftGraphContract ContractsContractCreateContract(ctx).MicrosoftGraphContract(microsoftGraphContract).Execute()

Add new entity to contracts

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
    microsoftGraphContract := *openapiclient.NewMicrosoftGraphContract() // MicrosoftGraphContract | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsContractApi.ContractsContractCreateContract(context.Background()).MicrosoftGraphContract(microsoftGraphContract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsContractApi.ContractsContractCreateContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractCreateContract`: MicrosoftGraphContract
    fmt.Fprintf(os.Stdout, "Response from `ContractsContractApi.ContractsContractCreateContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractCreateContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphContract** | [**MicrosoftGraphContract**](MicrosoftGraphContract.md) | New entity | 

### Return type

[**MicrosoftGraphContract**](MicrosoftGraphContract.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractDeleteContract

> ContractsContractDeleteContract(ctx, contractId).IfMatch(ifMatch).Execute()

Delete entity from contracts

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
    contractId := "contractId_example" // string | key: id of contract
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsContractApi.ContractsContractDeleteContract(context.Background(), contractId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsContractApi.ContractsContractDeleteContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractDeleteContractRequest struct via the builder pattern


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


## ContractsContractGetContract

> MicrosoftGraphContract ContractsContractGetContract(ctx, contractId).Select_(select_).Expand(expand).Execute()

Get entity from contracts by key

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
    contractId := "contractId_example" // string | key: id of contract
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsContractApi.ContractsContractGetContract(context.Background(), contractId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsContractApi.ContractsContractGetContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractGetContract`: MicrosoftGraphContract
    fmt.Fprintf(os.Stdout, "Response from `ContractsContractApi.ContractsContractGetContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractGetContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphContract**](MicrosoftGraphContract.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractListContract

> CollectionOfContract ContractsContractListContract(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from contracts

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
    resp, r, err := api_client.ContractsContractApi.ContractsContractListContract(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsContractApi.ContractsContractListContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsContractListContract`: CollectionOfContract
    fmt.Fprintf(os.Stdout, "Response from `ContractsContractApi.ContractsContractListContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractListContractRequest struct via the builder pattern


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

[**CollectionOfContract**](CollectionOfContract.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsContractUpdateContract

> ContractsContractUpdateContract(ctx, contractId).MicrosoftGraphContract(microsoftGraphContract).Execute()

Update entity in contracts

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
    contractId := "contractId_example" // string | key: id of contract
    microsoftGraphContract := *openapiclient.NewMicrosoftGraphContract() // MicrosoftGraphContract | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractsContractApi.ContractsContractUpdateContract(context.Background(), contractId).MicrosoftGraphContract(microsoftGraphContract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsContractApi.ContractsContractUpdateContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | key: id of contract | 

### Other Parameters

Other parameters are passed through a pointer to a apiContractsContractUpdateContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphContract** | [**MicrosoftGraphContract**](MicrosoftGraphContract.md) | New property values | 

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

