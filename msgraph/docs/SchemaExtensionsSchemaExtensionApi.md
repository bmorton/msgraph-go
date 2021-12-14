# \SchemaExtensionsSchemaExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemaExtensionsSchemaExtensionCreateSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionCreateSchemaExtension) | **Post** /schemaExtensions | Add new entity to schemaExtensions
[**SchemaExtensionsSchemaExtensionDeleteSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionDeleteSchemaExtension) | **Delete** /schemaExtensions/{schemaExtension-id} | Delete entity from schemaExtensions
[**SchemaExtensionsSchemaExtensionGetSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionGetSchemaExtension) | **Get** /schemaExtensions/{schemaExtension-id} | Get entity from schemaExtensions by key
[**SchemaExtensionsSchemaExtensionListSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionListSchemaExtension) | **Get** /schemaExtensions | Get entities from schemaExtensions
[**SchemaExtensionsSchemaExtensionUpdateSchemaExtension**](SchemaExtensionsSchemaExtensionApi.md#SchemaExtensionsSchemaExtensionUpdateSchemaExtension) | **Patch** /schemaExtensions/{schemaExtension-id} | Update entity in schemaExtensions



## SchemaExtensionsSchemaExtensionCreateSchemaExtension

> MicrosoftGraphSchemaExtension SchemaExtensionsSchemaExtensionCreateSchemaExtension(ctx).MicrosoftGraphSchemaExtension(microsoftGraphSchemaExtension).Execute()

Add new entity to schemaExtensions

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
    microsoftGraphSchemaExtension := *openapiclient.NewMicrosoftGraphSchemaExtension() // MicrosoftGraphSchemaExtension | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionCreateSchemaExtension(context.Background()).MicrosoftGraphSchemaExtension(microsoftGraphSchemaExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionCreateSchemaExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaExtensionsSchemaExtensionCreateSchemaExtension`: MicrosoftGraphSchemaExtension
    fmt.Fprintf(os.Stdout, "Response from `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionCreateSchemaExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchemaExtensionsSchemaExtensionCreateSchemaExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSchemaExtension** | [**MicrosoftGraphSchemaExtension**](MicrosoftGraphSchemaExtension.md) | New entity | 

### Return type

[**MicrosoftGraphSchemaExtension**](MicrosoftGraphSchemaExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaExtensionsSchemaExtensionDeleteSchemaExtension

> SchemaExtensionsSchemaExtensionDeleteSchemaExtension(ctx, schemaExtensionId).IfMatch(ifMatch).Execute()

Delete entity from schemaExtensions

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
    schemaExtensionId := "schemaExtensionId_example" // string | key: id of schemaExtension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionDeleteSchemaExtension(context.Background(), schemaExtensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionDeleteSchemaExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaExtensionId** | **string** | key: id of schemaExtension | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaExtensionsSchemaExtensionDeleteSchemaExtensionRequest struct via the builder pattern


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


## SchemaExtensionsSchemaExtensionGetSchemaExtension

> MicrosoftGraphSchemaExtension SchemaExtensionsSchemaExtensionGetSchemaExtension(ctx, schemaExtensionId).Select_(select_).Expand(expand).Execute()

Get entity from schemaExtensions by key

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
    schemaExtensionId := "schemaExtensionId_example" // string | key: id of schemaExtension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionGetSchemaExtension(context.Background(), schemaExtensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionGetSchemaExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaExtensionsSchemaExtensionGetSchemaExtension`: MicrosoftGraphSchemaExtension
    fmt.Fprintf(os.Stdout, "Response from `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionGetSchemaExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaExtensionId** | **string** | key: id of schemaExtension | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaExtensionsSchemaExtensionGetSchemaExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSchemaExtension**](MicrosoftGraphSchemaExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaExtensionsSchemaExtensionListSchemaExtension

> CollectionOfSchemaExtension SchemaExtensionsSchemaExtensionListSchemaExtension(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from schemaExtensions

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
    resp, r, err := api_client.SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionListSchemaExtension(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionListSchemaExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaExtensionsSchemaExtensionListSchemaExtension`: CollectionOfSchemaExtension
    fmt.Fprintf(os.Stdout, "Response from `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionListSchemaExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchemaExtensionsSchemaExtensionListSchemaExtensionRequest struct via the builder pattern


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

[**CollectionOfSchemaExtension**](CollectionOfSchemaExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaExtensionsSchemaExtensionUpdateSchemaExtension

> SchemaExtensionsSchemaExtensionUpdateSchemaExtension(ctx, schemaExtensionId).MicrosoftGraphSchemaExtension(microsoftGraphSchemaExtension).Execute()

Update entity in schemaExtensions

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
    schemaExtensionId := "schemaExtensionId_example" // string | key: id of schemaExtension
    microsoftGraphSchemaExtension := *openapiclient.NewMicrosoftGraphSchemaExtension() // MicrosoftGraphSchemaExtension | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionUpdateSchemaExtension(context.Background(), schemaExtensionId).MicrosoftGraphSchemaExtension(microsoftGraphSchemaExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemaExtensionsSchemaExtensionApi.SchemaExtensionsSchemaExtensionUpdateSchemaExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaExtensionId** | **string** | key: id of schemaExtension | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaExtensionsSchemaExtensionUpdateSchemaExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSchemaExtension** | [**MicrosoftGraphSchemaExtension**](MicrosoftGraphSchemaExtension.md) | New property values | 

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

