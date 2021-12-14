# \SubscriptionsSubscriptionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsSubscriptionCreateSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionCreateSubscription) | **Post** /subscriptions | Add new entity to subscriptions
[**SubscriptionsSubscriptionDeleteSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionDeleteSubscription) | **Delete** /subscriptions/{subscription-id} | Delete entity from subscriptions
[**SubscriptionsSubscriptionGetSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionGetSubscription) | **Get** /subscriptions/{subscription-id} | Get entity from subscriptions by key
[**SubscriptionsSubscriptionListSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionListSubscription) | **Get** /subscriptions | Get entities from subscriptions
[**SubscriptionsSubscriptionUpdateSubscription**](SubscriptionsSubscriptionApi.md#SubscriptionsSubscriptionUpdateSubscription) | **Patch** /subscriptions/{subscription-id} | Update entity in subscriptions



## SubscriptionsSubscriptionCreateSubscription

> MicrosoftGraphSubscription SubscriptionsSubscriptionCreateSubscription(ctx).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()

Add new entity to subscriptions

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
    microsoftGraphSubscription := *openapiclient.NewMicrosoftGraphSubscription() // MicrosoftGraphSubscription | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsSubscriptionApi.SubscriptionsSubscriptionCreateSubscription(context.Background()).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionCreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionCreateSubscription`: MicrosoftGraphSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionCreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | New entity | 

### Return type

[**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionDeleteSubscription

> SubscriptionsSubscriptionDeleteSubscription(ctx, subscriptionId).IfMatch(ifMatch).Execute()

Delete entity from subscriptions

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
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsSubscriptionApi.SubscriptionsSubscriptionDeleteSubscription(context.Background(), subscriptionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionDeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionDeleteSubscriptionRequest struct via the builder pattern


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


## SubscriptionsSubscriptionGetSubscription

> MicrosoftGraphSubscription SubscriptionsSubscriptionGetSubscription(ctx, subscriptionId).Select_(select_).Execute()

Get entity from subscriptions by key

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
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsSubscriptionApi.SubscriptionsSubscriptionGetSubscription(context.Background(), subscriptionId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionGetSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionGetSubscription`: MicrosoftGraphSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionGetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionListSubscription

> CollectionOfSubscription SubscriptionsSubscriptionListSubscription(ctx).Search(search).Select_(select_).Execute()

Get entities from subscriptions

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsSubscriptionApi.SubscriptionsSubscriptionListSubscription(context.Background()).Search(search).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionListSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionListSubscription`: CollectionOfSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionListSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionListSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search items by search phrases | 
 **select_** | **[]string** | Select properties to be returned | 

### Return type

[**CollectionOfSubscription**](CollectionOfSubscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionUpdateSubscription

> SubscriptionsSubscriptionUpdateSubscription(ctx, subscriptionId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()

Update entity in subscriptions

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
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    microsoftGraphSubscription := *openapiclient.NewMicrosoftGraphSubscription() // MicrosoftGraphSubscription | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsSubscriptionApi.SubscriptionsSubscriptionUpdateSubscription(context.Background(), subscriptionId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsSubscriptionApi.SubscriptionsSubscriptionUpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | New property values | 

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

