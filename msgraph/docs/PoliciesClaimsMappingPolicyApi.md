# \PoliciesClaimsMappingPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesCreateClaimsMappingPolicies) | **Post** /policies/claimsMappingPolicies | Create new navigation property to claimsMappingPolicies for policies
[**PoliciesDeleteClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesDeleteClaimsMappingPolicies) | **Delete** /policies/claimsMappingPolicies/{claimsMappingPolicy-id} | Delete navigation property claimsMappingPolicies for policies
[**PoliciesGetClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesGetClaimsMappingPolicies) | **Get** /policies/claimsMappingPolicies/{claimsMappingPolicy-id} | Get claimsMappingPolicies from policies
[**PoliciesListClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesListClaimsMappingPolicies) | **Get** /policies/claimsMappingPolicies | Get claimsMappingPolicies from policies
[**PoliciesUpdateClaimsMappingPolicies**](PoliciesClaimsMappingPolicyApi.md#PoliciesUpdateClaimsMappingPolicies) | **Patch** /policies/claimsMappingPolicies/{claimsMappingPolicy-id} | Update the navigation property claimsMappingPolicies in policies



## PoliciesCreateClaimsMappingPolicies

> MicrosoftGraphClaimsMappingPolicy PoliciesCreateClaimsMappingPolicies(ctx).MicrosoftGraphClaimsMappingPolicy(microsoftGraphClaimsMappingPolicy).Execute()

Create new navigation property to claimsMappingPolicies for policies



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
    microsoftGraphClaimsMappingPolicy := *openapiclient.NewMicrosoftGraphClaimsMappingPolicy() // MicrosoftGraphClaimsMappingPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesClaimsMappingPolicyApi.PoliciesCreateClaimsMappingPolicies(context.Background()).MicrosoftGraphClaimsMappingPolicy(microsoftGraphClaimsMappingPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesClaimsMappingPolicyApi.PoliciesCreateClaimsMappingPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreateClaimsMappingPolicies`: MicrosoftGraphClaimsMappingPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesClaimsMappingPolicyApi.PoliciesCreateClaimsMappingPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreateClaimsMappingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphClaimsMappingPolicy** | [**MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeleteClaimsMappingPolicies

> PoliciesDeleteClaimsMappingPolicies(ctx, claimsMappingPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property claimsMappingPolicies for policies



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
    claimsMappingPolicyId := "claimsMappingPolicyId_example" // string | key: id of claimsMappingPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesClaimsMappingPolicyApi.PoliciesDeleteClaimsMappingPolicies(context.Background(), claimsMappingPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesClaimsMappingPolicyApi.PoliciesDeleteClaimsMappingPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**claimsMappingPolicyId** | **string** | key: id of claimsMappingPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteClaimsMappingPoliciesRequest struct via the builder pattern


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


## PoliciesGetClaimsMappingPolicies

> MicrosoftGraphClaimsMappingPolicy PoliciesGetClaimsMappingPolicies(ctx, claimsMappingPolicyId).Select_(select_).Expand(expand).Execute()

Get claimsMappingPolicies from policies



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
    claimsMappingPolicyId := "claimsMappingPolicyId_example" // string | key: id of claimsMappingPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesClaimsMappingPolicyApi.PoliciesGetClaimsMappingPolicies(context.Background(), claimsMappingPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesClaimsMappingPolicyApi.PoliciesGetClaimsMappingPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetClaimsMappingPolicies`: MicrosoftGraphClaimsMappingPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesClaimsMappingPolicyApi.PoliciesGetClaimsMappingPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**claimsMappingPolicyId** | **string** | key: id of claimsMappingPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetClaimsMappingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListClaimsMappingPolicies

> CollectionOfClaimsMappingPolicy PoliciesListClaimsMappingPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get claimsMappingPolicies from policies



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
    resp, r, err := api_client.PoliciesClaimsMappingPolicyApi.PoliciesListClaimsMappingPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesClaimsMappingPolicyApi.PoliciesListClaimsMappingPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListClaimsMappingPolicies`: CollectionOfClaimsMappingPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesClaimsMappingPolicyApi.PoliciesListClaimsMappingPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListClaimsMappingPoliciesRequest struct via the builder pattern


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

[**CollectionOfClaimsMappingPolicy**](CollectionOfClaimsMappingPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateClaimsMappingPolicies

> PoliciesUpdateClaimsMappingPolicies(ctx, claimsMappingPolicyId).MicrosoftGraphClaimsMappingPolicy(microsoftGraphClaimsMappingPolicy).Execute()

Update the navigation property claimsMappingPolicies in policies



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
    claimsMappingPolicyId := "claimsMappingPolicyId_example" // string | key: id of claimsMappingPolicy
    microsoftGraphClaimsMappingPolicy := *openapiclient.NewMicrosoftGraphClaimsMappingPolicy() // MicrosoftGraphClaimsMappingPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesClaimsMappingPolicyApi.PoliciesUpdateClaimsMappingPolicies(context.Background(), claimsMappingPolicyId).MicrosoftGraphClaimsMappingPolicy(microsoftGraphClaimsMappingPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesClaimsMappingPolicyApi.PoliciesUpdateClaimsMappingPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**claimsMappingPolicyId** | **string** | key: id of claimsMappingPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateClaimsMappingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphClaimsMappingPolicy** | [**MicrosoftGraphClaimsMappingPolicy**](MicrosoftGraphClaimsMappingPolicy.md) | New navigation property values | 

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

