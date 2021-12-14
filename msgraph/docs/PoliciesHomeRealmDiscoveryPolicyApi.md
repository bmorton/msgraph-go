# \PoliciesHomeRealmDiscoveryPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesCreateHomeRealmDiscoveryPolicies) | **Post** /policies/homeRealmDiscoveryPolicies | Create new navigation property to homeRealmDiscoveryPolicies for policies
[**PoliciesDeleteHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesDeleteHomeRealmDiscoveryPolicies) | **Delete** /policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id} | Delete navigation property homeRealmDiscoveryPolicies for policies
[**PoliciesGetHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesGetHomeRealmDiscoveryPolicies) | **Get** /policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id} | Get homeRealmDiscoveryPolicies from policies
[**PoliciesListHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesListHomeRealmDiscoveryPolicies) | **Get** /policies/homeRealmDiscoveryPolicies | Get homeRealmDiscoveryPolicies from policies
[**PoliciesUpdateHomeRealmDiscoveryPolicies**](PoliciesHomeRealmDiscoveryPolicyApi.md#PoliciesUpdateHomeRealmDiscoveryPolicies) | **Patch** /policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy-id} | Update the navigation property homeRealmDiscoveryPolicies in policies



## PoliciesCreateHomeRealmDiscoveryPolicies

> MicrosoftGraphHomeRealmDiscoveryPolicy PoliciesCreateHomeRealmDiscoveryPolicies(ctx).MicrosoftGraphHomeRealmDiscoveryPolicy(microsoftGraphHomeRealmDiscoveryPolicy).Execute()

Create new navigation property to homeRealmDiscoveryPolicies for policies



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
    microsoftGraphHomeRealmDiscoveryPolicy := *openapiclient.NewMicrosoftGraphHomeRealmDiscoveryPolicy() // MicrosoftGraphHomeRealmDiscoveryPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesHomeRealmDiscoveryPolicyApi.PoliciesCreateHomeRealmDiscoveryPolicies(context.Background()).MicrosoftGraphHomeRealmDiscoveryPolicy(microsoftGraphHomeRealmDiscoveryPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesCreateHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreateHomeRealmDiscoveryPolicies`: MicrosoftGraphHomeRealmDiscoveryPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesCreateHomeRealmDiscoveryPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreateHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphHomeRealmDiscoveryPolicy** | [**MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeleteHomeRealmDiscoveryPolicies

> PoliciesDeleteHomeRealmDiscoveryPolicies(ctx, homeRealmDiscoveryPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property homeRealmDiscoveryPolicies for policies



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
    homeRealmDiscoveryPolicyId := "homeRealmDiscoveryPolicyId_example" // string | key: id of homeRealmDiscoveryPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesHomeRealmDiscoveryPolicyApi.PoliciesDeleteHomeRealmDiscoveryPolicies(context.Background(), homeRealmDiscoveryPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesDeleteHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homeRealmDiscoveryPolicyId** | **string** | key: id of homeRealmDiscoveryPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


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


## PoliciesGetHomeRealmDiscoveryPolicies

> MicrosoftGraphHomeRealmDiscoveryPolicy PoliciesGetHomeRealmDiscoveryPolicies(ctx, homeRealmDiscoveryPolicyId).Select_(select_).Expand(expand).Execute()

Get homeRealmDiscoveryPolicies from policies



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
    homeRealmDiscoveryPolicyId := "homeRealmDiscoveryPolicyId_example" // string | key: id of homeRealmDiscoveryPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesHomeRealmDiscoveryPolicyApi.PoliciesGetHomeRealmDiscoveryPolicies(context.Background(), homeRealmDiscoveryPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesGetHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetHomeRealmDiscoveryPolicies`: MicrosoftGraphHomeRealmDiscoveryPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesGetHomeRealmDiscoveryPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homeRealmDiscoveryPolicyId** | **string** | key: id of homeRealmDiscoveryPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListHomeRealmDiscoveryPolicies

> CollectionOfHomeRealmDiscoveryPolicy PoliciesListHomeRealmDiscoveryPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get homeRealmDiscoveryPolicies from policies



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
    resp, r, err := api_client.PoliciesHomeRealmDiscoveryPolicyApi.PoliciesListHomeRealmDiscoveryPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesListHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListHomeRealmDiscoveryPolicies`: CollectionOfHomeRealmDiscoveryPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesListHomeRealmDiscoveryPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


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

[**CollectionOfHomeRealmDiscoveryPolicy**](CollectionOfHomeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateHomeRealmDiscoveryPolicies

> PoliciesUpdateHomeRealmDiscoveryPolicies(ctx, homeRealmDiscoveryPolicyId).MicrosoftGraphHomeRealmDiscoveryPolicy(microsoftGraphHomeRealmDiscoveryPolicy).Execute()

Update the navigation property homeRealmDiscoveryPolicies in policies



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
    homeRealmDiscoveryPolicyId := "homeRealmDiscoveryPolicyId_example" // string | key: id of homeRealmDiscoveryPolicy
    microsoftGraphHomeRealmDiscoveryPolicy := *openapiclient.NewMicrosoftGraphHomeRealmDiscoveryPolicy() // MicrosoftGraphHomeRealmDiscoveryPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesHomeRealmDiscoveryPolicyApi.PoliciesUpdateHomeRealmDiscoveryPolicies(context.Background(), homeRealmDiscoveryPolicyId).MicrosoftGraphHomeRealmDiscoveryPolicy(microsoftGraphHomeRealmDiscoveryPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesHomeRealmDiscoveryPolicyApi.PoliciesUpdateHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**homeRealmDiscoveryPolicyId** | **string** | key: id of homeRealmDiscoveryPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphHomeRealmDiscoveryPolicy** | [**MicrosoftGraphHomeRealmDiscoveryPolicy**](MicrosoftGraphHomeRealmDiscoveryPolicy.md) | New navigation property values | 

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

