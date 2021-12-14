# \PoliciesPermissionGrantPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesCreatePermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesCreatePermissionGrantPolicies) | **Post** /policies/permissionGrantPolicies | Create new navigation property to permissionGrantPolicies for policies
[**PoliciesDeletePermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesDeletePermissionGrantPolicies) | **Delete** /policies/permissionGrantPolicies/{permissionGrantPolicy-id} | Delete navigation property permissionGrantPolicies for policies
[**PoliciesGetPermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesGetPermissionGrantPolicies) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id} | Get permissionGrantPolicies from policies
[**PoliciesListPermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesListPermissionGrantPolicies) | **Get** /policies/permissionGrantPolicies | Get permissionGrantPolicies from policies
[**PoliciesPermissionGrantPoliciesCreateExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesCreateExcludes) | **Post** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes | Create new navigation property to excludes for policies
[**PoliciesPermissionGrantPoliciesCreateIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesCreateIncludes) | **Post** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes | Create new navigation property to includes for policies
[**PoliciesPermissionGrantPoliciesDeleteExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesDeleteExcludes) | **Delete** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes/{permissionGrantConditionSet-id} | Delete navigation property excludes for policies
[**PoliciesPermissionGrantPoliciesDeleteIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesDeleteIncludes) | **Delete** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes/{permissionGrantConditionSet-id} | Delete navigation property includes for policies
[**PoliciesPermissionGrantPoliciesGetExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesGetExcludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes/{permissionGrantConditionSet-id} | Get excludes from policies
[**PoliciesPermissionGrantPoliciesGetIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesGetIncludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes/{permissionGrantConditionSet-id} | Get includes from policies
[**PoliciesPermissionGrantPoliciesListExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesListExcludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes | Get excludes from policies
[**PoliciesPermissionGrantPoliciesListIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesListIncludes) | **Get** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes | Get includes from policies
[**PoliciesPermissionGrantPoliciesUpdateExcludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesUpdateExcludes) | **Patch** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/excludes/{permissionGrantConditionSet-id} | Update the navigation property excludes in policies
[**PoliciesPermissionGrantPoliciesUpdateIncludes**](PoliciesPermissionGrantPolicyApi.md#PoliciesPermissionGrantPoliciesUpdateIncludes) | **Patch** /policies/permissionGrantPolicies/{permissionGrantPolicy-id}/includes/{permissionGrantConditionSet-id} | Update the navigation property includes in policies
[**PoliciesUpdatePermissionGrantPolicies**](PoliciesPermissionGrantPolicyApi.md#PoliciesUpdatePermissionGrantPolicies) | **Patch** /policies/permissionGrantPolicies/{permissionGrantPolicy-id} | Update the navigation property permissionGrantPolicies in policies



## PoliciesCreatePermissionGrantPolicies

> MicrosoftGraphPermissionGrantPolicy PoliciesCreatePermissionGrantPolicies(ctx).MicrosoftGraphPermissionGrantPolicy(microsoftGraphPermissionGrantPolicy).Execute()

Create new navigation property to permissionGrantPolicies for policies



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
    microsoftGraphPermissionGrantPolicy := *openapiclient.NewMicrosoftGraphPermissionGrantPolicy() // MicrosoftGraphPermissionGrantPolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesCreatePermissionGrantPolicies(context.Background()).MicrosoftGraphPermissionGrantPolicy(microsoftGraphPermissionGrantPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesCreatePermissionGrantPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesCreatePermissionGrantPolicies`: MicrosoftGraphPermissionGrantPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesCreatePermissionGrantPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesCreatePermissionGrantPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPermissionGrantPolicy** | [**MicrosoftGraphPermissionGrantPolicy**](MicrosoftGraphPermissionGrantPolicy.md) | New navigation property | 

### Return type

[**MicrosoftGraphPermissionGrantPolicy**](MicrosoftGraphPermissionGrantPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesDeletePermissionGrantPolicies

> PoliciesDeletePermissionGrantPolicies(ctx, permissionGrantPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property permissionGrantPolicies for policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesDeletePermissionGrantPolicies(context.Background(), permissionGrantPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesDeletePermissionGrantPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeletePermissionGrantPoliciesRequest struct via the builder pattern


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


## PoliciesGetPermissionGrantPolicies

> MicrosoftGraphPermissionGrantPolicy PoliciesGetPermissionGrantPolicies(ctx, permissionGrantPolicyId).Select_(select_).Expand(expand).Execute()

Get permissionGrantPolicies from policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesGetPermissionGrantPolicies(context.Background(), permissionGrantPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesGetPermissionGrantPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetPermissionGrantPolicies`: MicrosoftGraphPermissionGrantPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesGetPermissionGrantPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetPermissionGrantPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPermissionGrantPolicy**](MicrosoftGraphPermissionGrantPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesListPermissionGrantPolicies

> CollectionOfPermissionGrantPolicy PoliciesListPermissionGrantPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get permissionGrantPolicies from policies



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
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesListPermissionGrantPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesListPermissionGrantPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesListPermissionGrantPolicies`: CollectionOfPermissionGrantPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesListPermissionGrantPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesListPermissionGrantPoliciesRequest struct via the builder pattern


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

[**CollectionOfPermissionGrantPolicy**](CollectionOfPermissionGrantPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesCreateExcludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesCreateExcludes(ctx, permissionGrantPolicyId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()

Create new navigation property to excludes for policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    microsoftGraphPermissionGrantConditionSet := *openapiclient.NewMicrosoftGraphPermissionGrantConditionSet() // MicrosoftGraphPermissionGrantConditionSet | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesCreateExcludes(context.Background(), permissionGrantPolicyId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesCreateExcludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPermissionGrantPoliciesCreateExcludes`: MicrosoftGraphPermissionGrantConditionSet
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesCreateExcludes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesCreateExcludesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | New navigation property | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesCreateIncludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesCreateIncludes(ctx, permissionGrantPolicyId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()

Create new navigation property to includes for policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    microsoftGraphPermissionGrantConditionSet := *openapiclient.NewMicrosoftGraphPermissionGrantConditionSet() // MicrosoftGraphPermissionGrantConditionSet | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesCreateIncludes(context.Background(), permissionGrantPolicyId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesCreateIncludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPermissionGrantPoliciesCreateIncludes`: MicrosoftGraphPermissionGrantConditionSet
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesCreateIncludes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesCreateIncludesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | New navigation property | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesDeleteExcludes

> PoliciesPermissionGrantPoliciesDeleteExcludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId).IfMatch(ifMatch).Execute()

Delete navigation property excludes for policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    permissionGrantConditionSetId := "permissionGrantConditionSetId_example" // string | key: id of permissionGrantConditionSet
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesDeleteExcludes(context.Background(), permissionGrantPolicyId, permissionGrantConditionSetId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesDeleteExcludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string** | key: id of permissionGrantConditionSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesDeleteExcludesRequest struct via the builder pattern


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


## PoliciesPermissionGrantPoliciesDeleteIncludes

> PoliciesPermissionGrantPoliciesDeleteIncludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId).IfMatch(ifMatch).Execute()

Delete navigation property includes for policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    permissionGrantConditionSetId := "permissionGrantConditionSetId_example" // string | key: id of permissionGrantConditionSet
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesDeleteIncludes(context.Background(), permissionGrantPolicyId, permissionGrantConditionSetId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesDeleteIncludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string** | key: id of permissionGrantConditionSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesDeleteIncludesRequest struct via the builder pattern


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


## PoliciesPermissionGrantPoliciesGetExcludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesGetExcludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId).Select_(select_).Expand(expand).Execute()

Get excludes from policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    permissionGrantConditionSetId := "permissionGrantConditionSetId_example" // string | key: id of permissionGrantConditionSet
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesGetExcludes(context.Background(), permissionGrantPolicyId, permissionGrantConditionSetId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesGetExcludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPermissionGrantPoliciesGetExcludes`: MicrosoftGraphPermissionGrantConditionSet
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesGetExcludes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string** | key: id of permissionGrantConditionSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesGetExcludesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesGetIncludes

> MicrosoftGraphPermissionGrantConditionSet PoliciesPermissionGrantPoliciesGetIncludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId).Select_(select_).Expand(expand).Execute()

Get includes from policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    permissionGrantConditionSetId := "permissionGrantConditionSetId_example" // string | key: id of permissionGrantConditionSet
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesGetIncludes(context.Background(), permissionGrantPolicyId, permissionGrantConditionSetId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesGetIncludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPermissionGrantPoliciesGetIncludes`: MicrosoftGraphPermissionGrantConditionSet
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesGetIncludes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string** | key: id of permissionGrantConditionSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesGetIncludesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesListExcludes

> CollectionOfPermissionGrantConditionSet PoliciesPermissionGrantPoliciesListExcludes(ctx, permissionGrantPolicyId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get excludes from policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
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
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesListExcludes(context.Background(), permissionGrantPolicyId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesListExcludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPermissionGrantPoliciesListExcludes`: CollectionOfPermissionGrantConditionSet
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesListExcludes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesListExcludesRequest struct via the builder pattern


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

[**CollectionOfPermissionGrantConditionSet**](CollectionOfPermissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesListIncludes

> CollectionOfPermissionGrantConditionSet PoliciesPermissionGrantPoliciesListIncludes(ctx, permissionGrantPolicyId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get includes from policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
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
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesListIncludes(context.Background(), permissionGrantPolicyId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesListIncludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPermissionGrantPoliciesListIncludes`: CollectionOfPermissionGrantConditionSet
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesListIncludes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesListIncludesRequest struct via the builder pattern


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

[**CollectionOfPermissionGrantConditionSet**](CollectionOfPermissionGrantConditionSet.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPermissionGrantPoliciesUpdateExcludes

> PoliciesPermissionGrantPoliciesUpdateExcludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()

Update the navigation property excludes in policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    permissionGrantConditionSetId := "permissionGrantConditionSetId_example" // string | key: id of permissionGrantConditionSet
    microsoftGraphPermissionGrantConditionSet := *openapiclient.NewMicrosoftGraphPermissionGrantConditionSet() // MicrosoftGraphPermissionGrantConditionSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesUpdateExcludes(context.Background(), permissionGrantPolicyId, permissionGrantConditionSetId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesUpdateExcludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string** | key: id of permissionGrantConditionSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesUpdateExcludesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | New navigation property values | 

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


## PoliciesPermissionGrantPoliciesUpdateIncludes

> PoliciesPermissionGrantPoliciesUpdateIncludes(ctx, permissionGrantPolicyId, permissionGrantConditionSetId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()

Update the navigation property includes in policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    permissionGrantConditionSetId := "permissionGrantConditionSetId_example" // string | key: id of permissionGrantConditionSet
    microsoftGraphPermissionGrantConditionSet := *openapiclient.NewMicrosoftGraphPermissionGrantConditionSet() // MicrosoftGraphPermissionGrantConditionSet | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesUpdateIncludes(context.Background(), permissionGrantPolicyId, permissionGrantConditionSetId).MicrosoftGraphPermissionGrantConditionSet(microsoftGraphPermissionGrantConditionSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesPermissionGrantPoliciesUpdateIncludes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 
**permissionGrantConditionSetId** | **string** | key: id of permissionGrantConditionSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPermissionGrantPoliciesUpdateIncludesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphPermissionGrantConditionSet** | [**MicrosoftGraphPermissionGrantConditionSet**](MicrosoftGraphPermissionGrantConditionSet.md) | New navigation property values | 

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


## PoliciesUpdatePermissionGrantPolicies

> PoliciesUpdatePermissionGrantPolicies(ctx, permissionGrantPolicyId).MicrosoftGraphPermissionGrantPolicy(microsoftGraphPermissionGrantPolicy).Execute()

Update the navigation property permissionGrantPolicies in policies



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
    permissionGrantPolicyId := "permissionGrantPolicyId_example" // string | key: id of permissionGrantPolicy
    microsoftGraphPermissionGrantPolicy := *openapiclient.NewMicrosoftGraphPermissionGrantPolicy() // MicrosoftGraphPermissionGrantPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPermissionGrantPolicyApi.PoliciesUpdatePermissionGrantPolicies(context.Background(), permissionGrantPolicyId).MicrosoftGraphPermissionGrantPolicy(microsoftGraphPermissionGrantPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPermissionGrantPolicyApi.PoliciesUpdatePermissionGrantPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGrantPolicyId** | **string** | key: id of permissionGrantPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdatePermissionGrantPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPermissionGrantPolicy** | [**MicrosoftGraphPermissionGrantPolicy**](MicrosoftGraphPermissionGrantPolicy.md) | New navigation property values | 

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

