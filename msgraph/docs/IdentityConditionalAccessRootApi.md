# \IdentityConditionalAccessRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityConditionalAccessCreateNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessCreateNamedLocations) | **Post** /identity/conditionalAccess/namedLocations | Create new navigation property to namedLocations for identity
[**IdentityConditionalAccessCreatePolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessCreatePolicies) | **Post** /identity/conditionalAccess/policies | Create new navigation property to policies for identity
[**IdentityConditionalAccessDeleteNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessDeleteNamedLocations) | **Delete** /identity/conditionalAccess/namedLocations/{namedLocation-id} | Delete navigation property namedLocations for identity
[**IdentityConditionalAccessDeletePolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessDeletePolicies) | **Delete** /identity/conditionalAccess/policies/{conditionalAccessPolicy-id} | Delete navigation property policies for identity
[**IdentityConditionalAccessGetNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessGetNamedLocations) | **Get** /identity/conditionalAccess/namedLocations/{namedLocation-id} | Get namedLocations from identity
[**IdentityConditionalAccessGetPolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessGetPolicies) | **Get** /identity/conditionalAccess/policies/{conditionalAccessPolicy-id} | Get policies from identity
[**IdentityConditionalAccessListNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessListNamedLocations) | **Get** /identity/conditionalAccess/namedLocations | Get namedLocations from identity
[**IdentityConditionalAccessListPolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessListPolicies) | **Get** /identity/conditionalAccess/policies | Get policies from identity
[**IdentityConditionalAccessUpdateNamedLocations**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessUpdateNamedLocations) | **Patch** /identity/conditionalAccess/namedLocations/{namedLocation-id} | Update the navigation property namedLocations in identity
[**IdentityConditionalAccessUpdatePolicies**](IdentityConditionalAccessRootApi.md#IdentityConditionalAccessUpdatePolicies) | **Patch** /identity/conditionalAccess/policies/{conditionalAccessPolicy-id} | Update the navigation property policies in identity
[**IdentityDeleteConditionalAccess**](IdentityConditionalAccessRootApi.md#IdentityDeleteConditionalAccess) | **Delete** /identity/conditionalAccess | Delete navigation property conditionalAccess for identity
[**IdentityGetConditionalAccess**](IdentityConditionalAccessRootApi.md#IdentityGetConditionalAccess) | **Get** /identity/conditionalAccess | Get conditionalAccess from identity
[**IdentityUpdateConditionalAccess**](IdentityConditionalAccessRootApi.md#IdentityUpdateConditionalAccess) | **Patch** /identity/conditionalAccess | Update the navigation property conditionalAccess in identity



## IdentityConditionalAccessCreateNamedLocations

> MicrosoftGraphNamedLocation IdentityConditionalAccessCreateNamedLocations(ctx).MicrosoftGraphNamedLocation(microsoftGraphNamedLocation).Execute()

Create new navigation property to namedLocations for identity



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
    microsoftGraphNamedLocation := *openapiclient.NewMicrosoftGraphNamedLocation() // MicrosoftGraphNamedLocation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessCreateNamedLocations(context.Background()).MicrosoftGraphNamedLocation(microsoftGraphNamedLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessCreateNamedLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityConditionalAccessCreateNamedLocations`: MicrosoftGraphNamedLocation
    fmt.Fprintf(os.Stdout, "Response from `IdentityConditionalAccessRootApi.IdentityConditionalAccessCreateNamedLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessCreateNamedLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphNamedLocation** | [**MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md) | New navigation property | 

### Return type

[**MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessCreatePolicies

> MicrosoftGraphConditionalAccessPolicy IdentityConditionalAccessCreatePolicies(ctx).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()

Create new navigation property to policies for identity



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
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessCreatePolicies(context.Background()).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessCreatePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityConditionalAccessCreatePolicies`: MicrosoftGraphConditionalAccessPolicy
    fmt.Fprintf(os.Stdout, "Response from `IdentityConditionalAccessRootApi.IdentityConditionalAccessCreatePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessCreatePoliciesRequest struct via the builder pattern


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


## IdentityConditionalAccessDeleteNamedLocations

> IdentityConditionalAccessDeleteNamedLocations(ctx, namedLocationId).IfMatch(ifMatch).Execute()

Delete navigation property namedLocations for identity



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
    namedLocationId := "namedLocationId_example" // string | key: id of namedLocation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessDeleteNamedLocations(context.Background(), namedLocationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessDeleteNamedLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namedLocationId** | **string** | key: id of namedLocation | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessDeleteNamedLocationsRequest struct via the builder pattern


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


## IdentityConditionalAccessDeletePolicies

> IdentityConditionalAccessDeletePolicies(ctx, conditionalAccessPolicyId).IfMatch(ifMatch).Execute()

Delete navigation property policies for identity



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
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessDeletePolicies(context.Background(), conditionalAccessPolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessDeletePolicies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIdentityConditionalAccessDeletePoliciesRequest struct via the builder pattern


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


## IdentityConditionalAccessGetNamedLocations

> MicrosoftGraphNamedLocation IdentityConditionalAccessGetNamedLocations(ctx, namedLocationId).Select_(select_).Expand(expand).Execute()

Get namedLocations from identity



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
    namedLocationId := "namedLocationId_example" // string | key: id of namedLocation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessGetNamedLocations(context.Background(), namedLocationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessGetNamedLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityConditionalAccessGetNamedLocations`: MicrosoftGraphNamedLocation
    fmt.Fprintf(os.Stdout, "Response from `IdentityConditionalAccessRootApi.IdentityConditionalAccessGetNamedLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namedLocationId** | **string** | key: id of namedLocation | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessGetNamedLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessGetPolicies

> MicrosoftGraphConditionalAccessPolicy IdentityConditionalAccessGetPolicies(ctx, conditionalAccessPolicyId).Select_(select_).Expand(expand).Execute()

Get policies from identity



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
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessGetPolicies(context.Background(), conditionalAccessPolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessGetPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityConditionalAccessGetPolicies`: MicrosoftGraphConditionalAccessPolicy
    fmt.Fprintf(os.Stdout, "Response from `IdentityConditionalAccessRootApi.IdentityConditionalAccessGetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conditionalAccessPolicyId** | **string** | key: id of conditionalAccessPolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessGetPoliciesRequest struct via the builder pattern


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


## IdentityConditionalAccessListNamedLocations

> CollectionOfNamedLocation IdentityConditionalAccessListNamedLocations(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get namedLocations from identity



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
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessListNamedLocations(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessListNamedLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityConditionalAccessListNamedLocations`: CollectionOfNamedLocation
    fmt.Fprintf(os.Stdout, "Response from `IdentityConditionalAccessRootApi.IdentityConditionalAccessListNamedLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessListNamedLocationsRequest struct via the builder pattern


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

[**CollectionOfNamedLocation**](CollectionOfNamedLocation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityConditionalAccessListPolicies

> CollectionOfConditionalAccessPolicy IdentityConditionalAccessListPolicies(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get policies from identity



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
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessListPolicies(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityConditionalAccessListPolicies`: CollectionOfConditionalAccessPolicy
    fmt.Fprintf(os.Stdout, "Response from `IdentityConditionalAccessRootApi.IdentityConditionalAccessListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessListPoliciesRequest struct via the builder pattern


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


## IdentityConditionalAccessUpdateNamedLocations

> IdentityConditionalAccessUpdateNamedLocations(ctx, namedLocationId).MicrosoftGraphNamedLocation(microsoftGraphNamedLocation).Execute()

Update the navigation property namedLocations in identity



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
    namedLocationId := "namedLocationId_example" // string | key: id of namedLocation
    microsoftGraphNamedLocation := *openapiclient.NewMicrosoftGraphNamedLocation() // MicrosoftGraphNamedLocation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessUpdateNamedLocations(context.Background(), namedLocationId).MicrosoftGraphNamedLocation(microsoftGraphNamedLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessUpdateNamedLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namedLocationId** | **string** | key: id of namedLocation | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityConditionalAccessUpdateNamedLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphNamedLocation** | [**MicrosoftGraphNamedLocation**](MicrosoftGraphNamedLocation.md) | New navigation property values | 

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


## IdentityConditionalAccessUpdatePolicies

> IdentityConditionalAccessUpdatePolicies(ctx, conditionalAccessPolicyId).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()

Update the navigation property policies in identity



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
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityConditionalAccessUpdatePolicies(context.Background(), conditionalAccessPolicyId).MicrosoftGraphConditionalAccessPolicy(microsoftGraphConditionalAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityConditionalAccessUpdatePolicies``: %v\n", err)
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

Other parameters are passed through a pointer to a apiIdentityConditionalAccessUpdatePoliciesRequest struct via the builder pattern


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


## IdentityDeleteConditionalAccess

> IdentityDeleteConditionalAccess(ctx).IfMatch(ifMatch).Execute()

Delete navigation property conditionalAccess for identity



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityDeleteConditionalAccess(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityDeleteConditionalAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityDeleteConditionalAccessRequest struct via the builder pattern


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


## IdentityGetConditionalAccess

> MicrosoftGraphConditionalAccessRoot IdentityGetConditionalAccess(ctx).Select_(select_).Expand(expand).Execute()

Get conditionalAccess from identity



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityGetConditionalAccess(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityGetConditionalAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGetConditionalAccess`: MicrosoftGraphConditionalAccessRoot
    fmt.Fprintf(os.Stdout, "Response from `IdentityConditionalAccessRootApi.IdentityGetConditionalAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetConditionalAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphConditionalAccessRoot**](MicrosoftGraphConditionalAccessRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityUpdateConditionalAccess

> IdentityUpdateConditionalAccess(ctx).MicrosoftGraphConditionalAccessRoot(microsoftGraphConditionalAccessRoot).Execute()

Update the navigation property conditionalAccess in identity



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
    microsoftGraphConditionalAccessRoot := *openapiclient.NewMicrosoftGraphConditionalAccessRoot() // MicrosoftGraphConditionalAccessRoot | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityConditionalAccessRootApi.IdentityUpdateConditionalAccess(context.Background()).MicrosoftGraphConditionalAccessRoot(microsoftGraphConditionalAccessRoot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityConditionalAccessRootApi.IdentityUpdateConditionalAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityUpdateConditionalAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphConditionalAccessRoot** | [**MicrosoftGraphConditionalAccessRoot**](MicrosoftGraphConditionalAccessRoot.md) | New navigation property values | 

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

