# \MeScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeCreateScopedRoleMemberOf) | **Post** /me/scopedRoleMemberOf | Create new navigation property to scopedRoleMemberOf for me
[**MeDeleteScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeDeleteScopedRoleMemberOf) | **Delete** /me/scopedRoleMemberOf/{scopedRoleMembership-id} | Delete navigation property scopedRoleMemberOf for me
[**MeGetScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeGetScopedRoleMemberOf) | **Get** /me/scopedRoleMemberOf/{scopedRoleMembership-id} | Get scopedRoleMemberOf from me
[**MeListScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeListScopedRoleMemberOf) | **Get** /me/scopedRoleMemberOf | Get scopedRoleMemberOf from me
[**MeUpdateScopedRoleMemberOf**](MeScopedRoleMembershipApi.md#MeUpdateScopedRoleMemberOf) | **Patch** /me/scopedRoleMemberOf/{scopedRoleMembership-id} | Update the navigation property scopedRoleMemberOf in me



## MeCreateScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership MeCreateScopedRoleMemberOf(ctx).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Create new navigation property to scopedRoleMemberOf for me



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
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeScopedRoleMembershipApi.MeCreateScopedRoleMemberOf(context.Background()).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeScopedRoleMembershipApi.MeCreateScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateScopedRoleMemberOf`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `MeScopedRoleMembershipApi.MeCreateScopedRoleMemberOf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateScopedRoleMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | New navigation property | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDeleteScopedRoleMemberOf

> MeDeleteScopedRoleMemberOf(ctx, scopedRoleMembershipId).IfMatch(ifMatch).Execute()

Delete navigation property scopedRoleMemberOf for me



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeScopedRoleMembershipApi.MeDeleteScopedRoleMemberOf(context.Background(), scopedRoleMembershipId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeScopedRoleMembershipApi.MeDeleteScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteScopedRoleMemberOfRequest struct via the builder pattern


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


## MeGetScopedRoleMemberOf

> MicrosoftGraphScopedRoleMembership MeGetScopedRoleMemberOf(ctx, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()

Get scopedRoleMemberOf from me



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeScopedRoleMembershipApi.MeGetScopedRoleMemberOf(context.Background(), scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeScopedRoleMembershipApi.MeGetScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetScopedRoleMemberOf`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `MeScopedRoleMembershipApi.MeGetScopedRoleMemberOf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetScopedRoleMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListScopedRoleMemberOf

> CollectionOfScopedRoleMembership MeListScopedRoleMemberOf(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get scopedRoleMemberOf from me



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
    resp, r, err := api_client.MeScopedRoleMembershipApi.MeListScopedRoleMemberOf(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeScopedRoleMembershipApi.MeListScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListScopedRoleMemberOf`: CollectionOfScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `MeScopedRoleMembershipApi.MeListScopedRoleMemberOf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListScopedRoleMemberOfRequest struct via the builder pattern


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

[**CollectionOfScopedRoleMembership**](CollectionOfScopedRoleMembership.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateScopedRoleMemberOf

> MeUpdateScopedRoleMemberOf(ctx, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Update the navigation property scopedRoleMemberOf in me



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
    scopedRoleMembershipId := "scopedRoleMembershipId_example" // string | key: id of scopedRoleMembership
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeScopedRoleMembershipApi.MeUpdateScopedRoleMemberOf(context.Background(), scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeScopedRoleMembershipApi.MeUpdateScopedRoleMemberOf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateScopedRoleMemberOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | New navigation property values | 

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

