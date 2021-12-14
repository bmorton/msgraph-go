# \PoliciesTokenIssuancePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesCreateTokenIssuancePolicies) | **Post** /policies/tokenIssuancePolicies | Create new navigation property to tokenIssuancePolicies for policies
[**PoliciesDeleteTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesDeleteTokenIssuancePolicies) | **Delete** /policies/tokenIssuancePolicies/{tokenIssuancePolicy-id} | Delete navigation property tokenIssuancePolicies for policies
[**PoliciesGetTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesGetTokenIssuancePolicies) | **Get** /policies/tokenIssuancePolicies/{tokenIssuancePolicy-id} | Get tokenIssuancePolicies from policies
[**PoliciesListTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesListTokenIssuancePolicies) | **Get** /policies/tokenIssuancePolicies | Get tokenIssuancePolicies from policies
[**PoliciesUpdateTokenIssuancePolicies**](PoliciesTokenIssuancePolicyApi.md#PoliciesUpdateTokenIssuancePolicies) | **Patch** /policies/tokenIssuancePolicies/{tokenIssuancePolicy-id} | Update the navigation property tokenIssuancePolicies in policies



## PoliciesCreateTokenIssuancePolicies

> MicrosoftGraphTokenIssuancePolicy PoliciesCreateTokenIssuancePolicies(ctx).MicrosoftGraphTokenIssuancePolicy(microsoftGraphTokenIssuancePolicy).Execute()

Create new navigation property to tokenIssuancePolicies for policies



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
    microsoftGraphTokenIssuancePolicy := *openapiclient.NewMicrosoftGraphTokenIssuancePolicy() // MicrosoftGraphTokenIssuancePolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenIssuancePolicyApi.PoliciesCreateTokenIssuancePolicies(context.Background()).MicrosoftGraphTokenIssuancePolicy(microsoftGraphTokenIssuancePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenIssuancePolicyApi.PoliciesCreateTokenIssuancePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreateTokenIssuancePolicies`: MicrosoftGraphTokenIssuancePolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesTokenIssuancePolicyApi.PoliciesCreateTokenIssuancePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreateTokenIssuancePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTokenIssuancePolicy** | [**MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeleteTokenIssuancePolicies

> PoliciesDeleteTokenIssuancePolicies(ctx, tokenIssuancePolicyId).IfMatch(ifMatch).Execute()

Delete navigation property tokenIssuancePolicies for policies



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
    tokenIssuancePolicyId := "tokenIssuancePolicyId_example" // string | key: id of tokenIssuancePolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenIssuancePolicyApi.PoliciesDeleteTokenIssuancePolicies(context.Background(), tokenIssuancePolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenIssuancePolicyApi.PoliciesDeleteTokenIssuancePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenIssuancePolicyId** | **string** | key: id of tokenIssuancePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteTokenIssuancePoliciesRequest struct via the builder pattern


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


## PoliciesGetTokenIssuancePolicies

> MicrosoftGraphTokenIssuancePolicy PoliciesGetTokenIssuancePolicies(ctx, tokenIssuancePolicyId).Select_(select_).Expand(expand).Execute()

Get tokenIssuancePolicies from policies



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
    tokenIssuancePolicyId := "tokenIssuancePolicyId_example" // string | key: id of tokenIssuancePolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenIssuancePolicyApi.PoliciesGetTokenIssuancePolicies(context.Background(), tokenIssuancePolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenIssuancePolicyApi.PoliciesGetTokenIssuancePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetTokenIssuancePolicies`: MicrosoftGraphTokenIssuancePolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesTokenIssuancePolicyApi.PoliciesGetTokenIssuancePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenIssuancePolicyId** | **string** | key: id of tokenIssuancePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetTokenIssuancePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListTokenIssuancePolicies

> CollectionOfTokenIssuancePolicy PoliciesListTokenIssuancePolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tokenIssuancePolicies from policies



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
    resp, r, err := api_client.PoliciesTokenIssuancePolicyApi.PoliciesListTokenIssuancePolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenIssuancePolicyApi.PoliciesListTokenIssuancePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListTokenIssuancePolicies`: CollectionOfTokenIssuancePolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesTokenIssuancePolicyApi.PoliciesListTokenIssuancePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListTokenIssuancePoliciesRequest struct via the builder pattern


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

[**CollectionOfTokenIssuancePolicy**](CollectionOfTokenIssuancePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateTokenIssuancePolicies

> PoliciesUpdateTokenIssuancePolicies(ctx, tokenIssuancePolicyId).MicrosoftGraphTokenIssuancePolicy(microsoftGraphTokenIssuancePolicy).Execute()

Update the navigation property tokenIssuancePolicies in policies



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
    tokenIssuancePolicyId := "tokenIssuancePolicyId_example" // string | key: id of tokenIssuancePolicy
    microsoftGraphTokenIssuancePolicy := *openapiclient.NewMicrosoftGraphTokenIssuancePolicy() // MicrosoftGraphTokenIssuancePolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesTokenIssuancePolicyApi.PoliciesUpdateTokenIssuancePolicies(context.Background(), tokenIssuancePolicyId).MicrosoftGraphTokenIssuancePolicy(microsoftGraphTokenIssuancePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesTokenIssuancePolicyApi.PoliciesUpdateTokenIssuancePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenIssuancePolicyId** | **string** | key: id of tokenIssuancePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateTokenIssuancePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTokenIssuancePolicy** | [**MicrosoftGraphTokenIssuancePolicy**](MicrosoftGraphTokenIssuancePolicy.md) | New navigation property values | 

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

