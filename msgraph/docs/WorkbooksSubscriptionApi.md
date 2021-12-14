# \WorkbooksSubscriptionApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkbooksCreateSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksCreateSubscriptions) | **Post** /workbooks/{driveItem-id}/subscriptions | Create new navigation property to subscriptions for workbooks
[**WorkbooksDeleteSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksDeleteSubscriptions) | **Delete** /workbooks/{driveItem-id}/subscriptions/{subscription-id} | Delete navigation property subscriptions for workbooks
[**WorkbooksGetSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksGetSubscriptions) | **Get** /workbooks/{driveItem-id}/subscriptions/{subscription-id} | Get subscriptions from workbooks
[**WorkbooksListSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksListSubscriptions) | **Get** /workbooks/{driveItem-id}/subscriptions | Get subscriptions from workbooks
[**WorkbooksUpdateSubscriptions**](WorkbooksSubscriptionApi.md#WorkbooksUpdateSubscriptions) | **Patch** /workbooks/{driveItem-id}/subscriptions/{subscription-id} | Update the navigation property subscriptions in workbooks



## WorkbooksCreateSubscriptions

> MicrosoftGraphSubscription WorkbooksCreateSubscriptions(ctx, driveItemId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()

Create new navigation property to subscriptions for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    microsoftGraphSubscription := *openapiclient.NewMicrosoftGraphSubscription() // MicrosoftGraphSubscription | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksSubscriptionApi.WorkbooksCreateSubscriptions(context.Background(), driveItemId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksSubscriptionApi.WorkbooksCreateSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksCreateSubscriptions`: MicrosoftGraphSubscription
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksSubscriptionApi.WorkbooksCreateSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksCreateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | New navigation property | 

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


## WorkbooksDeleteSubscriptions

> WorkbooksDeleteSubscriptions(ctx, driveItemId, subscriptionId).IfMatch(ifMatch).Execute()

Delete navigation property subscriptions for workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksSubscriptionApi.WorkbooksDeleteSubscriptions(context.Background(), driveItemId, subscriptionId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksSubscriptionApi.WorkbooksDeleteSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksDeleteSubscriptionsRequest struct via the builder pattern


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


## WorkbooksGetSubscriptions

> MicrosoftGraphSubscription WorkbooksGetSubscriptions(ctx, driveItemId, subscriptionId).Select_(select_).Expand(expand).Execute()

Get subscriptions from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksSubscriptionApi.WorkbooksGetSubscriptions(context.Background(), driveItemId, subscriptionId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksSubscriptionApi.WorkbooksGetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksGetSubscriptions`: MicrosoftGraphSubscription
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksSubscriptionApi.WorkbooksGetSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

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


## WorkbooksListSubscriptions

> CollectionOfSubscription WorkbooksListSubscriptions(ctx, driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get subscriptions from workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
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
    resp, r, err := api_client.WorkbooksSubscriptionApi.WorkbooksListSubscriptions(context.Background(), driveItemId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksSubscriptionApi.WorkbooksListSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkbooksListSubscriptions`: CollectionOfSubscription
    fmt.Fprintf(os.Stdout, "Response from `WorkbooksSubscriptionApi.WorkbooksListSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksListSubscriptionsRequest struct via the builder pattern


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

[**CollectionOfSubscription**](CollectionOfSubscription.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkbooksUpdateSubscriptions

> WorkbooksUpdateSubscriptions(ctx, driveItemId, subscriptionId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()

Update the navigation property subscriptions in workbooks



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
    driveItemId := "driveItemId_example" // string | key: id of driveItem
    subscriptionId := "subscriptionId_example" // string | key: id of subscription
    microsoftGraphSubscription := *openapiclient.NewMicrosoftGraphSubscription() // MicrosoftGraphSubscription | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkbooksSubscriptionApi.WorkbooksUpdateSubscriptions(context.Background(), driveItemId, subscriptionId).MicrosoftGraphSubscription(microsoftGraphSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkbooksSubscriptionApi.WorkbooksUpdateSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**driveItemId** | **string** | key: id of driveItem | 
**subscriptionId** | **string** | key: id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkbooksUpdateSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphSubscription** | [**MicrosoftGraphSubscription**](MicrosoftGraphSubscription.md) | New navigation property values | 

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

