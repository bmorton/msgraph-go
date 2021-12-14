# \PoliciesConditionalAccessPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesCreateConditionalAccessPolicies) | **Post** /policies/conditionalAccessPolicies | Create new navigation property to conditionalAccessPolicies for policies
[**PoliciesDeleteConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesDeleteConditionalAccessPolicies) | **Delete** /policies/conditionalAccessPolicies/{conditionalAccessPolicy-id} | Delete navigation property conditionalAccessPolicies for policies
[**PoliciesGetConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesGetConditionalAccessPolicies) | **Get** /policies/conditionalAccessPolicies/{conditionalAccessPolicy-id} | Get conditionalAccessPolicies from policies
[**PoliciesListConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesListConditionalAccessPolicies) | **Get** /policies/conditionalAccessPolicies | Get conditionalAccessPolicies from policies
[**PoliciesUpdateConditionalAccessPolicies**](PoliciesConditionalAccessPolicyApi.md#PoliciesUpdateConditionalAccessPolicies) | **Patch** /policies/conditionalAccessPolicies/{conditionalAccessPolicy-id} | Update the navigation property conditionalAccessPolicies in policies



## PoliciesCreateConditionalAccessPolicies

> MicrosoftGraphConditionalAccessPolicy PoliciesCreateConditionalAccessPolicies(ctx).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()

Create new navigation property to conditionalAccessPolicies for policies



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
    microsoftGraphConditionalAccessPolicy := *openapiclient.NewMicrosoftGraphConditionalAccessPolicy() // MicrosoftGraphConditionalAccessPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConditionalAccessPolicyApi.PoliciesCreateConditionalAccessPolicies(context.Background()).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConditionalAccessPolicyApi.PoliciesCreateConditionalAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreateConditionalAccessPolicies`: MicrosoftGraphConditionalAccessPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesConditionalAccessPolicyApi.PoliciesCreateConditionalAccessPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreateConditionalAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphConditionalAccessPolicy** | [**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeleteConditionalAccessPolicies

> PoliciesDeleteConditionalAccessPolicies(ctx, conditionalAccessPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property conditionalAccessPolicies for policies



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
    conditionalAccessPolicyId := "conditionalAccessPolicyId_example" // string | key: id of conditionalAccessPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConditionalAccessPolicyApi.PoliciesDeleteConditionalAccessPolicies(context.Background(), conditionalAccessPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConditionalAccessPolicyApi.PoliciesDeleteConditionalAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string** | key: id of conditionalAccessPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteConditionalAccessPoliciesRequest struct via the builder pattern


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


## PoliciesGetConditionalAccessPolicies

> MicrosoftGraphConditionalAccessPolicy PoliciesGetConditionalAccessPolicies(ctx, conditionalAccessPolicyId).Select_(select_).Expand(expand).Execute()

Get conditionalAccessPolicies from policies



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
    conditionalAccessPolicyId := "conditionalAccessPolicyId_example" // string | key: id of conditionalAccessPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConditionalAccessPolicyApi.PoliciesGetConditionalAccessPolicies(context.Background(), conditionalAccessPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConditionalAccessPolicyApi.PoliciesGetConditionalAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetConditionalAccessPolicies`: MicrosoftGraphConditionalAccessPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesConditionalAccessPolicyApi.PoliciesGetConditionalAccessPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string** | key: id of conditionalAccessPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetConditionalAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListConditionalAccessPolicies

> CollectionOfConditionalAccessPolicy PoliciesListConditionalAccessPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get conditionalAccessPolicies from policies



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
    resp, r, err := api_client.PoliciesConditionalAccessPolicyApi.PoliciesListConditionalAccessPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConditionalAccessPolicyApi.PoliciesListConditionalAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListConditionalAccessPolicies`: CollectionOfConditionalAccessPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesConditionalAccessPolicyApi.PoliciesListConditionalAccessPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListConditionalAccessPoliciesRequest struct via the builder pattern


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

[**CollectionOfConditionalAccessPolicy**](CollectionOfConditionalAccessPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateConditionalAccessPolicies

> PoliciesUpdateConditionalAccessPolicies(ctx, conditionalAccessPolicyId).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()

Update the navigation property conditionalAccessPolicies in policies



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
    conditionalAccessPolicyId := "conditionalAccessPolicyId_example" // string | key: id of conditionalAccessPolicy
    microsoftGraphConditionalAccessPolicy := *openapiclient.NewMicrosoftGraphConditionalAccessPolicy() // MicrosoftGraphConditionalAccessPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConditionalAccessPolicyApi.PoliciesUpdateConditionalAccessPolicies(context.Background(), conditionalAccessPolicyId).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConditionalAccessPolicyApi.PoliciesUpdateConditionalAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string** | key: id of conditionalAccessPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateConditionalAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphConditionalAccessPolicy** | [**MicrosoftGraphConditionalAccessPolicy**](MicrosoftGraphConditionalAccessPolicy.md) | New navigation property values | 

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

