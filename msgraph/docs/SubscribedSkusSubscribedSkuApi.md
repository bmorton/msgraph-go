# \SubscribedSkusSubscribedSkuApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribedSkusSubscribedSkuCreateSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuCreateSubscribedSku) | **Post** /subscribedSkus | Add new entity to subscribedSkus
[**SubscribedSkusSubscribedSkuDeleteSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuDeleteSubscribedSku) | **Delete** /subscribedSkus/{subscribedSku-id} | Delete entity from subscribedSkus
[**SubscribedSkusSubscribedSkuGetSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuGetSubscribedSku) | **Get** /subscribedSkus/{subscribedSku-id} | Get entity from subscribedSkus by key
[**SubscribedSkusSubscribedSkuListSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuListSubscribedSku) | **Get** /subscribedSkus | Get entities from subscribedSkus
[**SubscribedSkusSubscribedSkuUpdateSubscribedSku**](SubscribedSkusSubscribedSkuApi.md#SubscribedSkusSubscribedSkuUpdateSubscribedSku) | **Patch** /subscribedSkus/{subscribedSku-id} | Update entity in subscribedSkus



## SubscribedSkusSubscribedSkuCreateSubscribedSku

> MicrosoftGraphSubscribedSku SubscribedSkusSubscribedSkuCreateSubscribedSku(ctx).MicrosoftGraphSubscribedSku(microsoftGraphSubscribedSku).Execute()

Add new entity to subscribedSkus

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
    microsoftGraphSubscribedSku := *openapiclient.NewMicrosoftGraphSubscribedSku() // MicrosoftGraphSubscribedSku | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuCreateSubscribedSku(context.Background()).MicrosoftGraphSubscribedSku(microsoftGraphSubscribedSku).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuCreateSubscribedSku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribedSkusSubscribedSkuCreateSubscribedSku`: MicrosoftGraphSubscribedSku
    fmt.Fprintf(os.Stdout, "Response from `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuCreateSubscribedSku`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribedSkusSubscribedSkuCreateSubscribedSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSubscribedSku** | [**MicrosoftGraphSubscribedSku**](MicrosoftGraphSubscribedSku.md) | New entity | 

### Return type

[**MicrosoftGraphSubscribedSku**](MicrosoftGraphSubscribedSku.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribedSkusSubscribedSkuDeleteSubscribedSku

> SubscribedSkusSubscribedSkuDeleteSubscribedSku(ctx, subscribedSkuId).IfMatch(ifMatch).Execute()

Delete entity from subscribedSkus

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
    subscribedSkuId := "subscribedSkuId_example" // string | key: id of subscribedSku
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuDeleteSubscribedSku(context.Background(), subscribedSkuId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuDeleteSubscribedSku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscribedSkuId** | **string** | key: id of subscribedSku | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribedSkusSubscribedSkuDeleteSubscribedSkuRequest struct via the builder pattern


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


## SubscribedSkusSubscribedSkuGetSubscribedSku

> MicrosoftGraphSubscribedSku SubscribedSkusSubscribedSkuGetSubscribedSku(ctx, subscribedSkuId).Select_(select_).Execute()

Get entity from subscribedSkus by key

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
    subscribedSkuId := "subscribedSkuId_example" // string | key: id of subscribedSku
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuGetSubscribedSku(context.Background(), subscribedSkuId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuGetSubscribedSku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribedSkusSubscribedSkuGetSubscribedSku`: MicrosoftGraphSubscribedSku
    fmt.Fprintf(os.Stdout, "Response from `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuGetSubscribedSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscribedSkuId** | **string** | key: id of subscribedSku | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribedSkusSubscribedSkuGetSubscribedSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphSubscribedSku**](MicrosoftGraphSubscribedSku.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribedSkusSubscribedSkuListSubscribedSku

> CollectionOfSubscribedSku SubscribedSkusSubscribedSkuListSubscribedSku(ctx).Search(search).Orderby(orderby).Select_(select_).Execute()

Get entities from subscribedSkus

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
    search := "search_example" // string | Search items by search phrases (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuListSubscribedSku(context.Background()).Search(search).Orderby(orderby).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuListSubscribedSku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribedSkusSubscribedSkuListSubscribedSku`: CollectionOfSubscribedSku
    fmt.Fprintf(os.Stdout, "Response from `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuListSubscribedSku`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribedSkusSubscribedSkuListSubscribedSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search items by search phrases | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfSubscribedSku**](CollectionOfSubscribedSku.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribedSkusSubscribedSkuUpdateSubscribedSku

> SubscribedSkusSubscribedSkuUpdateSubscribedSku(ctx, subscribedSkuId).MicrosoftGraphSubscribedSku(microsoftGraphSubscribedSku).Execute()

Update entity in subscribedSkus

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
    subscribedSkuId := "subscribedSkuId_example" // string | key: id of subscribedSku
    microsoftGraphSubscribedSku := *openapiclient.NewMicrosoftGraphSubscribedSku() // MicrosoftGraphSubscribedSku | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuUpdateSubscribedSku(context.Background(), subscribedSkuId).MicrosoftGraphSubscribedSku(microsoftGraphSubscribedSku).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribedSkusSubscribedSkuApi.SubscribedSkusSubscribedSkuUpdateSubscribedSku``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscribedSkuId** | **string** | key: id of subscribedSku | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribedSkusSubscribedSkuUpdateSubscribedSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSubscribedSku** | [**MicrosoftGraphSubscribedSku**](MicrosoftGraphSubscribedSku.md) | New property values | 

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

