# \IdentityIdentityProviderBaseApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityCreateIdentityProviders**](IdentityIdentityProviderBaseApi.md#IdentityCreateIdentityProviders) | **Post** /identity/identityProviders | Create new navigation property to identityProviders for identity
[**IdentityDeleteIdentityProviders**](IdentityIdentityProviderBaseApi.md#IdentityDeleteIdentityProviders) | **Delete** /identity/identityProviders/{identityProviderBase-id} | Delete navigation property identityProviders for identity
[**IdentityGetIdentityProviders**](IdentityIdentityProviderBaseApi.md#IdentityGetIdentityProviders) | **Get** /identity/identityProviders/{identityProviderBase-id} | Get identityProviders from identity
[**IdentityListIdentityProviders**](IdentityIdentityProviderBaseApi.md#IdentityListIdentityProviders) | **Get** /identity/identityProviders | Get identityProviders from identity
[**IdentityUpdateIdentityProviders**](IdentityIdentityProviderBaseApi.md#IdentityUpdateIdentityProviders) | **Patch** /identity/identityProviders/{identityProviderBase-id} | Update the navigation property identityProviders in identity



## IdentityCreateIdentityProviders

> MicrosoftGraphIdentityProviderBase IdentityCreateIdentityProviders(ctx).MicrosoftGraphIdentityProviderBase(microsoftGraphIdentityProviderBase).Execute()

Create new navigation property to identityProviders for identity



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
    microsoftGraphIdentityProviderBase := *openapiclient.NewMicrosoftGraphIdentityProviderBase() // MicrosoftGraphIdentityProviderBase | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityProviderBaseApi.IdentityCreateIdentityProviders(context.Background()).MicrosoftGraphIdentityProviderBase(microsoftGraphIdentityProviderBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityProviderBaseApi.IdentityCreateIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityCreateIdentityProviders`: MicrosoftGraphIdentityProviderBase
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityProviderBaseApi.IdentityCreateIdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityCreateIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIdentityProviderBase** | [**MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md) | New navigation property | 

### Return type

[**MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityDeleteIdentityProviders

> IdentityDeleteIdentityProviders(ctx, identityProviderBaseId).IfMatch(ifMatch).Execute()

Delete navigation property identityProviders for identity



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
    identityProviderBaseId := "identityProviderBaseId_example" // string | key: id of identityProviderBase
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityProviderBaseApi.IdentityDeleteIdentityProviders(context.Background(), identityProviderBaseId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityProviderBaseApi.IdentityDeleteIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderBaseId** | **string** | key: id of identityProviderBase | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityDeleteIdentityProvidersRequest struct via the builder pattern


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


## IdentityGetIdentityProviders

> MicrosoftGraphIdentityProviderBase IdentityGetIdentityProviders(ctx, identityProviderBaseId).Select_(select_).Expand(expand).Execute()

Get identityProviders from identity



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
    identityProviderBaseId := "identityProviderBaseId_example" // string | key: id of identityProviderBase
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityProviderBaseApi.IdentityGetIdentityProviders(context.Background(), identityProviderBaseId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityProviderBaseApi.IdentityGetIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetIdentityProviders`: MicrosoftGraphIdentityProviderBase
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityProviderBaseApi.IdentityGetIdentityProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderBaseId** | **string** | key: id of identityProviderBase | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityListIdentityProviders

> CollectionOfIdentityProviderBase IdentityListIdentityProviders(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get identityProviders from identity



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
    resp, r, err := api_client.IdentityIdentityProviderBaseApi.IdentityListIdentityProviders(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityProviderBaseApi.IdentityListIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityListIdentityProviders`: CollectionOfIdentityProviderBase
    fmt.Fprintf(os.Stdout, "Response from `IdentityIdentityProviderBaseApi.IdentityListIdentityProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityListIdentityProvidersRequest struct via the builder pattern


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

[**CollectionOfIdentityProviderBase**](CollectionOfIdentityProviderBase.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityUpdateIdentityProviders

> IdentityUpdateIdentityProviders(ctx, identityProviderBaseId).MicrosoftGraphIdentityProviderBase(microsoftGraphIdentityProviderBase).Execute()

Update the navigation property identityProviders in identity



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
    identityProviderBaseId := "identityProviderBaseId_example" // string | key: id of identityProviderBase
    microsoftGraphIdentityProviderBase := *openapiclient.NewMicrosoftGraphIdentityProviderBase() // MicrosoftGraphIdentityProviderBase | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityIdentityProviderBaseApi.IdentityUpdateIdentityProviders(context.Background(), identityProviderBaseId).MicrosoftGraphIdentityProviderBase(microsoftGraphIdentityProviderBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityIdentityProviderBaseApi.IdentityUpdateIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderBaseId** | **string** | key: id of identityProviderBase | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityUpdateIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphIdentityProviderBase** | [**MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md) | New navigation property values | 

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

