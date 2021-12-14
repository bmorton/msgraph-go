# \DeviceManagementResourceOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceManagementCreateResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementCreateResourceOperations) | **Post** /deviceManagement/resourceOperations | Create new navigation property to resourceOperations for deviceManagement
[**DeviceManagementDeleteResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementDeleteResourceOperations) | **Delete** /deviceManagement/resourceOperations/{resourceOperation-id} | Delete navigation property resourceOperations for deviceManagement
[**DeviceManagementGetResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementGetResourceOperations) | **Get** /deviceManagement/resourceOperations/{resourceOperation-id} | Get resourceOperations from deviceManagement
[**DeviceManagementListResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementListResourceOperations) | **Get** /deviceManagement/resourceOperations | Get resourceOperations from deviceManagement
[**DeviceManagementUpdateResourceOperations**](DeviceManagementResourceOperationApi.md#DeviceManagementUpdateResourceOperations) | **Patch** /deviceManagement/resourceOperations/{resourceOperation-id} | Update the navigation property resourceOperations in deviceManagement



## DeviceManagementCreateResourceOperations

> MicrosoftGraphResourceOperation DeviceManagementCreateResourceOperations(ctx).MicrosoftGraphResourceOperation(microsoftGraphResourceOperation).Execute()

Create new navigation property to resourceOperations for deviceManagement



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
    microsoftGraphResourceOperation := *openapiclient.NewMicrosoftGraphResourceOperation() // MicrosoftGraphResourceOperation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourceOperationApi.DeviceManagementCreateResourceOperations(context.Background()).MicrosoftGraphResourceOperation(microsoftGraphResourceOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourceOperationApi.DeviceManagementCreateResourceOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementCreateResourceOperations`: MicrosoftGraphResourceOperation
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourceOperationApi.DeviceManagementCreateResourceOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementCreateResourceOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphResourceOperation** | [**MicrosoftGraphResourceOperation**](MicrosoftGraphResourceOperation.md) | New navigation property | 

### Return type

[**MicrosoftGraphResourceOperation**](MicrosoftGraphResourceOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementDeleteResourceOperations

> DeviceManagementDeleteResourceOperations(ctx, resourceOperationId).IfMatch(ifMatch).Execute()

Delete navigation property resourceOperations for deviceManagement



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
    resourceOperationId := "resourceOperationId_example" // string | key: id of resourceOperation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourceOperationApi.DeviceManagementDeleteResourceOperations(context.Background(), resourceOperationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourceOperationApi.DeviceManagementDeleteResourceOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOperationId** | **string** | key: id of resourceOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementDeleteResourceOperationsRequest struct via the builder pattern


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


## DeviceManagementGetResourceOperations

> MicrosoftGraphResourceOperation DeviceManagementGetResourceOperations(ctx, resourceOperationId).Select_(select_).Expand(expand).Execute()

Get resourceOperations from deviceManagement



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
    resourceOperationId := "resourceOperationId_example" // string | key: id of resourceOperation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourceOperationApi.DeviceManagementGetResourceOperations(context.Background(), resourceOperationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourceOperationApi.DeviceManagementGetResourceOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementGetResourceOperations`: MicrosoftGraphResourceOperation
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourceOperationApi.DeviceManagementGetResourceOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOperationId** | **string** | key: id of resourceOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementGetResourceOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphResourceOperation**](MicrosoftGraphResourceOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementListResourceOperations

> CollectionOfResourceOperation DeviceManagementListResourceOperations(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get resourceOperations from deviceManagement



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
    resp, r, err := api_client.DeviceManagementResourceOperationApi.DeviceManagementListResourceOperations(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourceOperationApi.DeviceManagementListResourceOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceManagementListResourceOperations`: CollectionOfResourceOperation
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourceOperationApi.DeviceManagementListResourceOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementListResourceOperationsRequest struct via the builder pattern


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

[**CollectionOfResourceOperation**](CollectionOfResourceOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceManagementUpdateResourceOperations

> DeviceManagementUpdateResourceOperations(ctx, resourceOperationId).MicrosoftGraphResourceOperation(microsoftGraphResourceOperation).Execute()

Update the navigation property resourceOperations in deviceManagement



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
    resourceOperationId := "resourceOperationId_example" // string | key: id of resourceOperation
    microsoftGraphResourceOperation := *openapiclient.NewMicrosoftGraphResourceOperation() // MicrosoftGraphResourceOperation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourceOperationApi.DeviceManagementUpdateResourceOperations(context.Background(), resourceOperationId).MicrosoftGraphResourceOperation(microsoftGraphResourceOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourceOperationApi.DeviceManagementUpdateResourceOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOperationId** | **string** | key: id of resourceOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceManagementUpdateResourceOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphResourceOperation** | [**MicrosoftGraphResourceOperation**](MicrosoftGraphResourceOperation.md) | New navigation property values | 

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

