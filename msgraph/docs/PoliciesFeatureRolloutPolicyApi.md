# \PoliciesFeatureRolloutPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreateFeatureRolloutPolicies**](PoliciesFeatureRolloutPolicyApi.md#PoliciesCreateFeatureRolloutPolicies) | **Post** /policies/featureRolloutPolicies | Create new navigation property to featureRolloutPolicies for policies
[**PoliciesDeleteFeatureRolloutPolicies**](PoliciesFeatureRolloutPolicyApi.md#PoliciesDeleteFeatureRolloutPolicies) | **Delete** /policies/featureRolloutPolicies/{featureRolloutPolicy-id} | Delete navigation property featureRolloutPolicies for policies
[**PoliciesFeatureRolloutPoliciesCreateAppliesTo**](PoliciesFeatureRolloutPolicyApi.md#PoliciesFeatureRolloutPoliciesCreateAppliesTo) | **Post** /policies/featureRolloutPolicies/{featureRolloutPolicy-id}/appliesTo | Create new navigation property to appliesTo for policies
[**PoliciesFeatureRolloutPoliciesDeleteAppliesTo**](PoliciesFeatureRolloutPolicyApi.md#PoliciesFeatureRolloutPoliciesDeleteAppliesTo) | **Delete** /policies/featureRolloutPolicies/{featureRolloutPolicy-id}/appliesTo/{directoryObject-id} | Delete navigation property appliesTo for policies
[**PoliciesFeatureRolloutPoliciesGetAppliesTo**](PoliciesFeatureRolloutPolicyApi.md#PoliciesFeatureRolloutPoliciesGetAppliesTo) | **Get** /policies/featureRolloutPolicies/{featureRolloutPolicy-id}/appliesTo/{directoryObject-id} | Get appliesTo from policies
[**PoliciesFeatureRolloutPoliciesListAppliesTo**](PoliciesFeatureRolloutPolicyApi.md#PoliciesFeatureRolloutPoliciesListAppliesTo) | **Get** /policies/featureRolloutPolicies/{featureRolloutPolicy-id}/appliesTo | Get appliesTo from policies
[**PoliciesFeatureRolloutPoliciesUpdateAppliesTo**](PoliciesFeatureRolloutPolicyApi.md#PoliciesFeatureRolloutPoliciesUpdateAppliesTo) | **Patch** /policies/featureRolloutPolicies/{featureRolloutPolicy-id}/appliesTo/{directoryObject-id} | Update the navigation property appliesTo in policies
[**PoliciesGetFeatureRolloutPolicies**](PoliciesFeatureRolloutPolicyApi.md#PoliciesGetFeatureRolloutPolicies) | **Get** /policies/featureRolloutPolicies/{featureRolloutPolicy-id} | Get featureRolloutPolicies from policies
[**PoliciesListFeatureRolloutPolicies**](PoliciesFeatureRolloutPolicyApi.md#PoliciesListFeatureRolloutPolicies) | **Get** /policies/featureRolloutPolicies | Get featureRolloutPolicies from policies
[**PoliciesUpdateFeatureRolloutPolicies**](PoliciesFeatureRolloutPolicyApi.md#PoliciesUpdateFeatureRolloutPolicies) | **Patch** /policies/featureRolloutPolicies/{featureRolloutPolicy-id} | Update the navigation property featureRolloutPolicies in policies



## PoliciesCreateFeatureRolloutPolicies

> MicrosoftGraphFeatureRolloutPolicy PoliciesCreateFeatureRolloutPolicies(ctx).MicrosoftGraphFeatureRolloutPolicy(microsoftGraphFeatureRolloutPolicy).Execute()

Create new navigation property to featureRolloutPolicies for policies



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
    microsoftGraphFeatureRolloutPolicy := *openapiclient.NewMicrosoftGraphFeatureRolloutPolicy() // MicrosoftGraphFeatureRolloutPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesCreateFeatureRolloutPolicies(context.Background()).MicrosoftGraphFeatureRolloutPolicy(microsoftGraphFeatureRolloutPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesCreateFeatureRolloutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreateFeatureRolloutPolicies`: MicrosoftGraphFeatureRolloutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesFeatureRolloutPolicyApi.PoliciesCreateFeatureRolloutPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreateFeatureRolloutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphFeatureRolloutPolicy** | [**MicrosoftGraphFeatureRolloutPolicy**](MicrosoftGraphFeatureRolloutPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphFeatureRolloutPolicy**](MicrosoftGraphFeatureRolloutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeleteFeatureRolloutPolicies

> PoliciesDeleteFeatureRolloutPolicies(ctx, featureRolloutPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property featureRolloutPolicies for policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesDeleteFeatureRolloutPolicies(context.Background(), featureRolloutPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesDeleteFeatureRolloutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteFeatureRolloutPoliciesRequest struct via the builder pattern


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


## PoliciesFeatureRolloutPoliciesCreateAppliesTo

> MicrosoftGraphDirectoryObject PoliciesFeatureRolloutPoliciesCreateAppliesTo(ctx, featureRolloutPolicyId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()

Create new navigation property to appliesTo for policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
    microsoftGraphDirectoryObject := *openapiclient.NewMicrosoftGraphDirectoryObject() // MicrosoftGraphDirectoryObject | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesCreateAppliesTo(context.Background(), featureRolloutPolicyId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesCreateAppliesTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesFeatureRolloutPoliciesCreateAppliesTo`: MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesCreateAppliesTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesFeatureRolloutPoliciesCreateAppliesToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | New navigation property | 

### Return type

[**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesFeatureRolloutPoliciesDeleteAppliesTo

> PoliciesFeatureRolloutPoliciesDeleteAppliesTo(ctx, featureRolloutPolicyId, directoryObjectId).IfMatch(ifMatch).Execute()

Delete navigation property appliesTo for policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesDeleteAppliesTo(context.Background(), featureRolloutPolicyId, directoryObjectId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesDeleteAppliesTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesFeatureRolloutPoliciesDeleteAppliesToRequest struct via the builder pattern


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


## PoliciesFeatureRolloutPoliciesGetAppliesTo

> MicrosoftGraphDirectoryObject PoliciesFeatureRolloutPoliciesGetAppliesTo(ctx, featureRolloutPolicyId, directoryObjectId).Select_(select_).Expand(expand).Execute()

Get appliesTo from policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesGetAppliesTo(context.Background(), featureRolloutPolicyId, directoryObjectId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesGetAppliesTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesFeatureRolloutPoliciesGetAppliesTo`: MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesGetAppliesTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesFeatureRolloutPoliciesGetAppliesToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesFeatureRolloutPoliciesListAppliesTo

> CollectionOfDirectoryObject PoliciesFeatureRolloutPoliciesListAppliesTo(ctx, featureRolloutPolicyId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get appliesTo from policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
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
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesListAppliesTo(context.Background(), featureRolloutPolicyId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesListAppliesTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesFeatureRolloutPoliciesListAppliesTo`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesListAppliesTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesFeatureRolloutPoliciesListAppliesToRequest struct via the builder pattern


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

[**CollectionOfDirectoryObject**](CollectionOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesFeatureRolloutPoliciesUpdateAppliesTo

> PoliciesFeatureRolloutPoliciesUpdateAppliesTo(ctx, featureRolloutPolicyId, directoryObjectId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()

Update the navigation property appliesTo in policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    microsoftGraphDirectoryObject := *openapiclient.NewMicrosoftGraphDirectoryObject() // MicrosoftGraphDirectoryObject | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesUpdateAppliesTo(context.Background(), featureRolloutPolicyId, directoryObjectId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesFeatureRolloutPoliciesUpdateAppliesTo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesFeatureRolloutPoliciesUpdateAppliesToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | New navigation property values | 

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


## PoliciesGetFeatureRolloutPolicies

> MicrosoftGraphFeatureRolloutPolicy PoliciesGetFeatureRolloutPolicies(ctx, featureRolloutPolicyId).Select_(select_).Expand(expand).Execute()

Get featureRolloutPolicies from policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesGetFeatureRolloutPolicies(context.Background(), featureRolloutPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesGetFeatureRolloutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetFeatureRolloutPolicies`: MicrosoftGraphFeatureRolloutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesFeatureRolloutPolicyApi.PoliciesGetFeatureRolloutPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetFeatureRolloutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphFeatureRolloutPolicy**](MicrosoftGraphFeatureRolloutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListFeatureRolloutPolicies

> CollectionOfFeatureRolloutPolicy PoliciesListFeatureRolloutPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get featureRolloutPolicies from policies



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
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesListFeatureRolloutPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesListFeatureRolloutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListFeatureRolloutPolicies`: CollectionOfFeatureRolloutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesFeatureRolloutPolicyApi.PoliciesListFeatureRolloutPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListFeatureRolloutPoliciesRequest struct via the builder pattern


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

[**CollectionOfFeatureRolloutPolicy**](CollectionOfFeatureRolloutPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateFeatureRolloutPolicies

> PoliciesUpdateFeatureRolloutPolicies(ctx, featureRolloutPolicyId).MicrosoftGraphFeatureRolloutPolicy(microsoftGraphFeatureRolloutPolicy).Execute()

Update the navigation property featureRolloutPolicies in policies



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
    featureRolloutPolicyId := "featureRolloutPolicyId_example" // string | key: id of featureRolloutPolicy
    microsoftGraphFeatureRolloutPolicy := *openapiclient.NewMicrosoftGraphFeatureRolloutPolicy() // MicrosoftGraphFeatureRolloutPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesFeatureRolloutPolicyApi.PoliciesUpdateFeatureRolloutPolicies(context.Background(), featureRolloutPolicyId).MicrosoftGraphFeatureRolloutPolicy(microsoftGraphFeatureRolloutPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesFeatureRolloutPolicyApi.PoliciesUpdateFeatureRolloutPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureRolloutPolicyId** | **string** | key: id of featureRolloutPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateFeatureRolloutPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphFeatureRolloutPolicy** | [**MicrosoftGraphFeatureRolloutPolicy**](MicrosoftGraphFeatureRolloutPolicy.md) | New navigation property values | 

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

