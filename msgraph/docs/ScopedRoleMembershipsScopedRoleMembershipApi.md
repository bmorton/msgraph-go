# \ScopedRoleMembershipsScopedRoleMembershipApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership) | **Post** /scopedRoleMemberships | Add new entity to scopedRoleMemberships
[**ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership) | **Delete** /scopedRoleMemberships/{scopedRoleMembership-id} | Delete entity from scopedRoleMemberships
[**ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership) | **Get** /scopedRoleMemberships/{scopedRoleMembership-id} | Get entity from scopedRoleMemberships by key
[**ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership) | **Get** /scopedRoleMemberships | Get entities from scopedRoleMemberships
[**ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership**](ScopedRoleMembershipsScopedRoleMembershipApi.md#ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership) | **Patch** /scopedRoleMemberships/{scopedRoleMembership-id} | Update entity in scopedRoleMemberships



## ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership

> MicrosoftGraphScopedRoleMembership ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership(ctx).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Add new entity to scopedRoleMemberships

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
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership(context.Background()).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScopedRoleMembershipsScopedRoleMembershipCreateScopedRoleMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | New entity | 

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


## ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership

> ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership(ctx, scopedRoleMembershipId).IfMatch(ifMatch).Execute()

Delete entity from scopedRoleMemberships

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
    resp, r, err := api_client.ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership(context.Background(), scopedRoleMembershipId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembership``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScopedRoleMembershipsScopedRoleMembershipDeleteScopedRoleMembershipRequest struct via the builder pattern


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


## ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership

> MicrosoftGraphScopedRoleMembership ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership(ctx, scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()

Get entity from scopedRoleMemberships by key

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
    resp, r, err := api_client.ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership(context.Background(), scopedRoleMembershipId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership`: MicrosoftGraphScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scopedRoleMembershipId** | **string** | key: id of scopedRoleMembership | 

### Other Parameters

Other parameters are passed through a pointer to a apiScopedRoleMembershipsScopedRoleMembershipGetScopedRoleMembershipRequest struct via the builder pattern


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


## ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership

> CollectionOfScopedRoleMembership ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from scopedRoleMemberships

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
    resp, r, err := api_client.ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership`: CollectionOfScopedRoleMembership
    fmt.Fprintf(os.Stdout, "Response from `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScopedRoleMembershipsScopedRoleMembershipListScopedRoleMembershipRequest struct via the builder pattern


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


## ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership

> ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership(ctx, scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()

Update entity in scopedRoleMemberships

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
    microsoftGraphScopedRoleMembership := *openapiclient.NewMicrosoftGraphScopedRoleMembership() // MicrosoftGraphScopedRoleMembership | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership(context.Background(), scopedRoleMembershipId).MicrosoftGraphScopedRoleMembership(microsoftGraphScopedRoleMembership).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScopedRoleMembershipsScopedRoleMembershipApi.ScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembership``: %v\n", err)
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

Other parameters are passed through a pointer to a apiScopedRoleMembershipsScopedRoleMembershipUpdateScopedRoleMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphScopedRoleMembership** | [**MicrosoftGraphScopedRoleMembership**](MicrosoftGraphScopedRoleMembership.md) | New property values | 

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

