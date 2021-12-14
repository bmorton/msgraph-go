# \PoliciesTokenLifetimePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesCreateTokenLifetimePolicies) | **Post** /policies/tokenLifetimePolicies | Create new navigation property to tokenLifetimePolicies for policies
[**PoliciesDeleteTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesDeleteTokenLifetimePolicies) | **Delete** /policies/tokenLifetimePolicies/{tokenLifetimePolicy-id} | Delete navigation property tokenLifetimePolicies for policies
[**PoliciesGetTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesGetTokenLifetimePolicies) | **Get** /policies/tokenLifetimePolicies/{tokenLifetimePolicy-id} | Get tokenLifetimePolicies from policies
[**PoliciesListTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesListTokenLifetimePolicies) | **Get** /policies/tokenLifetimePolicies | Get tokenLifetimePolicies from policies
[**PoliciesUpdateTokenLifetimePolicies**](PoliciesTokenLifetimePolicyApi.md#PoliciesUpdateTokenLifetimePolicies) | **Patch** /policies/tokenLifetimePolicies/{tokenLifetimePolicy-id} | Update the navigation property tokenLifetimePolicies in policies



## PoliciesCreateTokenLifetimePolicies

> MicrosoftGraphTokenLifetimePolicy PoliciesCreateTokenLifetimePolicies(ctx).MicrosoftGraphTokenLifetimePolicy(microsoftGraphTokenLifetimePolicy).Execute()

Create new navigation property to tokenLifetimePolicies for policies



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
    microsoftGraphTokenLifetimePolicy := *openapiclient.NewMicrosoftGraphTokenLifetimePolicy() // MicrosoftGraphTokenLifetimePolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenLifetimePolicyApi.PoliciesCreateTokenLifetimePolicies(context.Background()).MicrosoftGraphTokenLifetimePolicy(microsoftGraphTokenLifetimePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenLifetimePolicyApi.PoliciesCreateTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreateTokenLifetimePolicies`: MicrosoftGraphTokenLifetimePolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesTokenLifetimePolicyApi.PoliciesCreateTokenLifetimePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreateTokenLifetimePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTokenLifetimePolicy** | [**MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeleteTokenLifetimePolicies

> PoliciesDeleteTokenLifetimePolicies(ctx, tokenLifetimePolicyId).IfMatch(ifMatch).Execute()

Delete navigation property tokenLifetimePolicies for policies



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
    tokenLifetimePolicyId := "tokenLifetimePolicyId_example" // string | key: id of tokenLifetimePolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenLifetimePolicyApi.PoliciesDeleteTokenLifetimePolicies(context.Background(), tokenLifetimePolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenLifetimePolicyApi.PoliciesDeleteTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenLifetimePolicyId** | **string** | key: id of tokenLifetimePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteTokenLifetimePoliciesRequest struct via the builder pattern


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


## PoliciesGetTokenLifetimePolicies

> MicrosoftGraphTokenLifetimePolicy PoliciesGetTokenLifetimePolicies(ctx, tokenLifetimePolicyId).Select_(select_).Expand(expand).Execute()

Get tokenLifetimePolicies from policies



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
    tokenLifetimePolicyId := "tokenLifetimePolicyId_example" // string | key: id of tokenLifetimePolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenLifetimePolicyApi.PoliciesGetTokenLifetimePolicies(context.Background(), tokenLifetimePolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenLifetimePolicyApi.PoliciesGetTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetTokenLifetimePolicies`: MicrosoftGraphTokenLifetimePolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesTokenLifetimePolicyApi.PoliciesGetTokenLifetimePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenLifetimePolicyId** | **string** | key: id of tokenLifetimePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetTokenLifetimePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListTokenLifetimePolicies

> CollectionOfTokenLifetimePolicy PoliciesListTokenLifetimePolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tokenLifetimePolicies from policies



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
    resp, r, err := api_client.PoliciesTokenLifetimePolicyApi.PoliciesListTokenLifetimePolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenLifetimePolicyApi.PoliciesListTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListTokenLifetimePolicies`: CollectionOfTokenLifetimePolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesTokenLifetimePolicyApi.PoliciesListTokenLifetimePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListTokenLifetimePoliciesRequest struct via the builder pattern


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

[**CollectionOfTokenLifetimePolicy**](CollectionOfTokenLifetimePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateTokenLifetimePolicies

> PoliciesUpdateTokenLifetimePolicies(ctx, tokenLifetimePolicyId).MicrosoftGraphTokenLifetimePolicy(microsoftGraphTokenLifetimePolicy).Execute()

Update the navigation property tokenLifetimePolicies in policies



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
    tokenLifetimePolicyId := "tokenLifetimePolicyId_example" // string | key: id of tokenLifetimePolicy
    microsoftGraphTokenLifetimePolicy := *openapiclient.NewMicrosoftGraphTokenLifetimePolicy() // MicrosoftGraphTokenLifetimePolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenLifetimePolicyApi.PoliciesUpdateTokenLifetimePolicies(context.Background(), tokenLifetimePolicyId).MicrosoftGraphTokenLifetimePolicy(microsoftGraphTokenLifetimePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenLifetimePolicyApi.PoliciesUpdateTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenLifetimePolicyId** | **string** | key: id of tokenLifetimePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateTokenLifetimePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTokenLifetimePolicy** | [**MicrosoftGraphTokenLifetimePolicy**](MicrosoftGraphTokenLifetimePolicy.md) | New navigation property values | 

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

