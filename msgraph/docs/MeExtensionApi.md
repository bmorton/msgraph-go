# \MeExtensionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateExtensions**](MeExtensionApi.md#MeCreateExtensions) | **Post** /me/extensions | Create new navigation property to extensions for me
[**MeDeleteExtensions**](MeExtensionApi.md#MeDeleteExtensions) | **Delete** /me/extensions/{extension-id} | Delete navigation property extensions for me
[**MeGetExtensions**](MeExtensionApi.md#MeGetExtensions) | **Get** /me/extensions/{extension-id} | Get extensions from me
[**MeListExtensions**](MeExtensionApi.md#MeListExtensions) | **Get** /me/extensions | Get extensions from me
[**MeUpdateExtensions**](MeExtensionApi.md#MeUpdateExtensions) | **Patch** /me/extensions/{extension-id} | Update the navigation property extensions in me



## MeCreateExtensions

> MicrosoftGraphExtension MeCreateExtensions(ctx).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Create new navigation property to extensions for me



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
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeExtensionApi.MeCreateExtensions(context.Background()).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeExtensionApi.MeCreateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeExtensionApi.MeCreateExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDeleteExtensions

> MeDeleteExtensions(ctx, extensionId).IfMatch(ifMatch).Execute()

Delete navigation property extensions for me



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
    extensionId := "extensionId_example" // string | key: id of extension
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeExtensionApi.MeDeleteExtensions(context.Background(), extensionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeExtensionApi.MeDeleteExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteExtensionsRequest struct via the builder pattern


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


## MeGetExtensions

> MicrosoftGraphExtension MeGetExtensions(ctx, extensionId).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    extensionId := "extensionId_example" // string | key: id of extension
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeExtensionApi.MeGetExtensions(context.Background(), extensionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeExtensionApi.MeGetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetExtensions`: MicrosoftGraphExtension
    fmt.Fprintf(os.Stdout, "Response from `MeExtensionApi.MeGetExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphExtension**](MicrosoftGraphExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListExtensions

> CollectionOfExtension MeListExtensions(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get extensions from me



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
    resp, r, err := api_client.MeExtensionApi.MeListExtensions(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeExtensionApi.MeListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListExtensions`: CollectionOfExtension
    fmt.Fprintf(os.Stdout, "Response from `MeExtensionApi.MeListExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListExtensionsRequest struct via the builder pattern


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

[**CollectionOfExtension**](CollectionOfExtension.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateExtensions

> MeUpdateExtensions(ctx, extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()

Update the navigation property extensions in me



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
    extensionId := "extensionId_example" // string | key: id of extension
    microsoftGraphExtension := *openapiclient.NewMicrosoftGraphExtension() // MicrosoftGraphExtension | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeExtensionApi.MeUpdateExtensions(context.Background(), extensionId).MicrosoftGraphExtension(microsoftGraphExtension).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeExtensionApi.MeUpdateExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | key: id of extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphExtension** | [**MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | New navigation property values | 

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

