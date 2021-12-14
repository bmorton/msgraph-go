# \PrintPrintOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrintCreateOperations**](PrintPrintOperationApi.md#PrintCreateOperations) | **Post** /print/operations | Create new navigation property to operations for print
[**PrintDeleteOperations**](PrintPrintOperationApi.md#PrintDeleteOperations) | **Delete** /print/operations/{printOperation-id} | Delete navigation property operations for print
[**PrintGetOperations**](PrintPrintOperationApi.md#PrintGetOperations) | **Get** /print/operations/{printOperation-id} | Get operations from print
[**PrintListOperations**](PrintPrintOperationApi.md#PrintListOperations) | **Get** /print/operations | Get operations from print
[**PrintUpdateOperations**](PrintPrintOperationApi.md#PrintUpdateOperations) | **Patch** /print/operations/{printOperation-id} | Update the navigation property operations in print



## PrintCreateOperations

> MicrosoftGraphPrintOperation PrintCreateOperations(ctx).MicrosoftGraphPrintOperation(microsoftGraphPrintOperation).Execute()

Create new navigation property to operations for print



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
    microsoftGraphPrintOperation := *openapiclient.NewMicrosoftGraphPrintOperation() // MicrosoftGraphPrintOperation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintOperationApi.PrintCreateOperations(context.Background()).MicrosoftGraphPrintOperation(microsoftGraphPrintOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintOperationApi.PrintCreateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintCreateOperations`: MicrosoftGraphPrintOperation
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintOperationApi.PrintCreateOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintCreateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPrintOperation** | [**MicrosoftGraphPrintOperation**](MicrosoftGraphPrintOperation.md) | New navigation property | 

### Return type

[**MicrosoftGraphPrintOperation**](MicrosoftGraphPrintOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintDeleteOperations

> PrintDeleteOperations(ctx, printOperationId).IfMatch(ifMatch).Execute()

Delete navigation property operations for print



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
    printOperationId := "printOperationId_example" // string | key: id of printOperation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintOperationApi.PrintDeleteOperations(context.Background(), printOperationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintOperationApi.PrintDeleteOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printOperationId** | **string** | key: id of printOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintDeleteOperationsRequest struct via the builder pattern


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


## PrintGetOperations

> MicrosoftGraphPrintOperation PrintGetOperations(ctx, printOperationId).Select_(select_).Expand(expand).Execute()

Get operations from print



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
    printOperationId := "printOperationId_example" // string | key: id of printOperation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintOperationApi.PrintGetOperations(context.Background(), printOperationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintOperationApi.PrintGetOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintGetOperations`: MicrosoftGraphPrintOperation
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintOperationApi.PrintGetOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printOperationId** | **string** | key: id of printOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintGetOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPrintOperation**](MicrosoftGraphPrintOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintListOperations

> CollectionOfPrintOperation PrintListOperations(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get operations from print



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
    resp, r, err := api_client.PrintPrintOperationApi.PrintListOperations(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintOperationApi.PrintListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintListOperations`: CollectionOfPrintOperation
    fmt.Fprintf(os.Stdout, "Response from `PrintPrintOperationApi.PrintListOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintListOperationsRequest struct via the builder pattern


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

[**CollectionOfPrintOperation**](CollectionOfPrintOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintUpdateOperations

> PrintUpdateOperations(ctx, printOperationId).MicrosoftGraphPrintOperation(microsoftGraphPrintOperation).Execute()

Update the navigation property operations in print



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
    printOperationId := "printOperationId_example" // string | key: id of printOperation
    microsoftGraphPrintOperation := *openapiclient.NewMicrosoftGraphPrintOperation() // MicrosoftGraphPrintOperation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrintPrintOperationApi.PrintUpdateOperations(context.Background(), printOperationId).MicrosoftGraphPrintOperation(microsoftGraphPrintOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintPrintOperationApi.PrintUpdateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printOperationId** | **string** | key: id of printOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintUpdateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPrintOperation** | [**MicrosoftGraphPrintOperation**](MicrosoftGraphPrintOperation.md) | New navigation property values | 

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

