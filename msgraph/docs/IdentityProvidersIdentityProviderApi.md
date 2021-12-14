# \IdentityProvidersIdentityProviderApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityProvidersIdentityProviderCreateIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderCreateIdentityProvider) | **Post** /identityProviders | Add new entity to identityProviders
[**IdentityProvidersIdentityProviderDeleteIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderDeleteIdentityProvider) | **Delete** /identityProviders/{identityProvider-id} | Delete entity from identityProviders
[**IdentityProvidersIdentityProviderGetIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderGetIdentityProvider) | **Get** /identityProviders/{identityProvider-id} | Get entity from identityProviders by key
[**IdentityProvidersIdentityProviderListIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderListIdentityProvider) | **Get** /identityProviders | Get entities from identityProviders
[**IdentityProvidersIdentityProviderUpdateIdentityProvider**](IdentityProvidersIdentityProviderApi.md#IdentityProvidersIdentityProviderUpdateIdentityProvider) | **Patch** /identityProviders/{identityProvider-id} | Update entity in identityProviders



## IdentityProvidersIdentityProviderCreateIdentityProvider

> MicrosoftGraphIdentityProvider IdentityProvidersIdentityProviderCreateIdentityProvider(ctx).MicrosoftGraphIdentityProvider(microsoftGraphIdentityProvider).Execute()

Add new entity to identityProviders

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
    microsoftGraphIdentityProvider := *openapiclient.NewMicrosoftGraphIdentityProvider() // MicrosoftGraphIdentityProvider | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderCreateIdentityProvider(context.Background()).MicrosoftGraphIdentityProvider(microsoftGraphIdentityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderCreateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersIdentityProviderCreateIdentityProvider`: MicrosoftGraphIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderCreateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersIdentityProviderCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIdentityProvider** | [**MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md) | New entity | 

### Return type

[**MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdentityProviderDeleteIdentityProvider

> IdentityProvidersIdentityProviderDeleteIdentityProvider(ctx, identityProviderId).IfMatch(ifMatch).Execute()

Delete entity from identityProviders

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
    identityProviderId := "identityProviderId_example" // string | key: id of identityProvider
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderDeleteIdentityProvider(context.Background(), identityProviderId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderDeleteIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderId** | **string** | key: id of identityProvider | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersIdentityProviderDeleteIdentityProviderRequest struct via the builder pattern


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


## IdentityProvidersIdentityProviderGetIdentityProvider

> MicrosoftGraphIdentityProvider IdentityProvidersIdentityProviderGetIdentityProvider(ctx, identityProviderId).Select_(select_).Expand(expand).Execute()

Get entity from identityProviders by key

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
    identityProviderId := "identityProviderId_example" // string | key: id of identityProvider
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderGetIdentityProvider(context.Background(), identityProviderId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderGetIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersIdentityProviderGetIdentityProvider`: MicrosoftGraphIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderGetIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderId** | **string** | key: id of identityProvider | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersIdentityProviderGetIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdentityProviderListIdentityProvider

> CollectionOfIdentityProvider IdentityProvidersIdentityProviderListIdentityProvider(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from identityProviders

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
    resp, r, err := api_client.IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderListIdentityProvider(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderListIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersIdentityProviderListIdentityProvider`: CollectionOfIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderListIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersIdentityProviderListIdentityProviderRequest struct via the builder pattern


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

[**CollectionOfIdentityProvider**](CollectionOfIdentityProvider.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityProvidersIdentityProviderUpdateIdentityProvider

> IdentityProvidersIdentityProviderUpdateIdentityProvider(ctx, identityProviderId).MicrosoftGraphIdentityProvider(microsoftGraphIdentityProvider).Execute()

Update entity in identityProviders

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
    identityProviderId := "identityProviderId_example" // string | key: id of identityProvider
    microsoftGraphIdentityProvider := *openapiclient.NewMicrosoftGraphIdentityProvider() // MicrosoftGraphIdentityProvider | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderUpdateIdentityProvider(context.Background(), identityProviderId).MicrosoftGraphIdentityProvider(microsoftGraphIdentityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersIdentityProviderApi.IdentityProvidersIdentityProviderUpdateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProviderId** | **string** | key: id of identityProvider | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersIdentityProviderUpdateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphIdentityProvider** | [**MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md) | New property values | 

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

