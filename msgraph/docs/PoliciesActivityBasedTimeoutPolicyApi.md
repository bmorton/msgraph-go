# \PoliciesActivityBasedTimeoutPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesCreateActivityBasedTimeoutPolicies) | **Post** /policies/activityBasedTimeoutPolicies | Create new navigation property to activityBasedTimeoutPolicies for policies
[**PoliciesDeleteActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesDeleteActivityBasedTimeoutPolicies) | **Delete** /policies/activityBasedTimeoutPolicies/{activityBasedTimeoutPolicy-id} | Delete navigation property activityBasedTimeoutPolicies for policies
[**PoliciesGetActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesGetActivityBasedTimeoutPolicies) | **Get** /policies/activityBasedTimeoutPolicies/{activityBasedTimeoutPolicy-id} | Get activityBasedTimeoutPolicies from policies
[**PoliciesListActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesListActivityBasedTimeoutPolicies) | **Get** /policies/activityBasedTimeoutPolicies | Get activityBasedTimeoutPolicies from policies
[**PoliciesUpdateActivityBasedTimeoutPolicies**](PoliciesActivityBasedTimeoutPolicyApi.md#PoliciesUpdateActivityBasedTimeoutPolicies) | **Patch** /policies/activityBasedTimeoutPolicies/{activityBasedTimeoutPolicy-id} | Update the navigation property activityBasedTimeoutPolicies in policies



## PoliciesCreateActivityBasedTimeoutPolicies

> MicrosoftGraphActivityBasedTimeoutPolicy PoliciesCreateActivityBasedTimeoutPolicies(ctx).MicrosoftGraphActivityBasedTimeoutPolicy(microsoftGraphActivityBasedTimeoutPolicy).Execute()

Create new navigation property to activityBasedTimeoutPolicies for policies



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
    microsoftGraphActivityBasedTimeoutPolicy := *openapiclient.NewMicrosoftGraphActivityBasedTimeoutPolicy() // MicrosoftGraphActivityBasedTimeoutPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesActivityBasedTimeoutPolicyApi.PoliciesCreateActivityBasedTimeoutPolicies(context.Background()).MicrosoftGraphActivityBasedTimeoutPolicy(microsoftGraphActivityBasedTimeoutPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesActivityBasedTimeoutPolicyApi.PoliciesCreateActivityBasedTimeoutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreateActivityBasedTimeoutPolicies`: MicrosoftGraphActivityBasedTimeoutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesActivityBasedTimeoutPolicyApi.PoliciesCreateActivityBasedTimeoutPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreateActivityBasedTimeoutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphActivityBasedTimeoutPolicy** | [**MicrosoftGraphActivityBasedTimeoutPolicy**](MicrosoftGraphActivityBasedTimeoutPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphActivityBasedTimeoutPolicy**](MicrosoftGraphActivityBasedTimeoutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeleteActivityBasedTimeoutPolicies

> PoliciesDeleteActivityBasedTimeoutPolicies(ctx, activityBasedTimeoutPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property activityBasedTimeoutPolicies for policies



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
    activityBasedTimeoutPolicyId := "activityBasedTimeoutPolicyId_example" // string | key: id of activityBasedTimeoutPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesActivityBasedTimeoutPolicyApi.PoliciesDeleteActivityBasedTimeoutPolicies(context.Background(), activityBasedTimeoutPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesActivityBasedTimeoutPolicyApi.PoliciesDeleteActivityBasedTimeoutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityBasedTimeoutPolicyId** | **string** | key: id of activityBasedTimeoutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteActivityBasedTimeoutPoliciesRequest struct via the builder pattern


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


## PoliciesGetActivityBasedTimeoutPolicies

> MicrosoftGraphActivityBasedTimeoutPolicy PoliciesGetActivityBasedTimeoutPolicies(ctx, activityBasedTimeoutPolicyId).Select_(select_).Expand(expand).Execute()

Get activityBasedTimeoutPolicies from policies



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
    activityBasedTimeoutPolicyId := "activityBasedTimeoutPolicyId_example" // string | key: id of activityBasedTimeoutPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesActivityBasedTimeoutPolicyApi.PoliciesGetActivityBasedTimeoutPolicies(context.Background(), activityBasedTimeoutPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesActivityBasedTimeoutPolicyApi.PoliciesGetActivityBasedTimeoutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetActivityBasedTimeoutPolicies`: MicrosoftGraphActivityBasedTimeoutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesActivityBasedTimeoutPolicyApi.PoliciesGetActivityBasedTimeoutPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityBasedTimeoutPolicyId** | **string** | key: id of activityBasedTimeoutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetActivityBasedTimeoutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphActivityBasedTimeoutPolicy**](MicrosoftGraphActivityBasedTimeoutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListActivityBasedTimeoutPolicies

> CollectionOfActivityBasedTimeoutPolicy PoliciesListActivityBasedTimeoutPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get activityBasedTimeoutPolicies from policies



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
    resp, r, err := api_client.PoliciesActivityBasedTimeoutPolicyApi.PoliciesListActivityBasedTimeoutPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesActivityBasedTimeoutPolicyApi.PoliciesListActivityBasedTimeoutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListActivityBasedTimeoutPolicies`: CollectionOfActivityBasedTimeoutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesActivityBasedTimeoutPolicyApi.PoliciesListActivityBasedTimeoutPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListActivityBasedTimeoutPoliciesRequest struct via the builder pattern


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

[**CollectionOfActivityBasedTimeoutPolicy**](CollectionOfActivityBasedTimeoutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateActivityBasedTimeoutPolicies

> PoliciesUpdateActivityBasedTimeoutPolicies(ctx, activityBasedTimeoutPolicyId).MicrosoftGraphActivityBasedTimeoutPolicy(microsoftGraphActivityBasedTimeoutPolicy).Execute()

Update the navigation property activityBasedTimeoutPolicies in policies



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
    activityBasedTimeoutPolicyId := "activityBasedTimeoutPolicyId_example" // string | key: id of activityBasedTimeoutPolicy
    microsoftGraphActivityBasedTimeoutPolicy := *openapiclient.NewMicrosoftGraphActivityBasedTimeoutPolicy() // MicrosoftGraphActivityBasedTimeoutPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesActivityBasedTimeoutPolicyApi.PoliciesUpdateActivityBasedTimeoutPolicies(context.Background(), activityBasedTimeoutPolicyId).MicrosoftGraphActivityBasedTimeoutPolicy(microsoftGraphActivityBasedTimeoutPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesActivityBasedTimeoutPolicyApi.PoliciesUpdateActivityBasedTimeoutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityBasedTimeoutPolicyId** | **string** | key: id of activityBasedTimeoutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateActivityBasedTimeoutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphActivityBasedTimeoutPolicy** | [**MicrosoftGraphActivityBasedTimeoutPolicy**](MicrosoftGraphActivityBasedTimeoutPolicy.md) | New navigation property values | 

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

