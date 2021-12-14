# \ServicePrincipalsServicePrincipalApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsServicePrincipalCreateServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalCreateServicePrincipal) | **Post** /servicePrincipals | Add new entity to servicePrincipals
[**ServicePrincipalsServicePrincipalDeleteServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalDeleteServicePrincipal) | **Delete** /servicePrincipals/{servicePrincipal-id} | Delete entity from servicePrincipals
[**ServicePrincipalsServicePrincipalGetServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalGetServicePrincipal) | **Get** /servicePrincipals/{servicePrincipal-id} | Get entity from servicePrincipals by key
[**ServicePrincipalsServicePrincipalListServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalListServicePrincipal) | **Get** /servicePrincipals | Get entities from servicePrincipals
[**ServicePrincipalsServicePrincipalUpdateServicePrincipal**](ServicePrincipalsServicePrincipalApi.md#ServicePrincipalsServicePrincipalUpdateServicePrincipal) | **Patch** /servicePrincipals/{servicePrincipal-id} | Update entity in servicePrincipals



## ServicePrincipalsServicePrincipalCreateServicePrincipal

> MicrosoftGraphServicePrincipal ServicePrincipalsServicePrincipalCreateServicePrincipal(ctx).MicrosoftGraphServicePrincipal(microsoftGraphServicePrincipal).Execute()

Add new entity to servicePrincipals

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
    microsoftGraphServicePrincipal := *openapiclient.NewMicrosoftGraphServicePrincipal() // MicrosoftGraphServicePrincipal | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalCreateServicePrincipal(context.Background()).MicrosoftGraphServicePrincipal(microsoftGraphServicePrincipal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalCreateServicePrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalCreateServicePrincipal`: MicrosoftGraphServicePrincipal
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalCreateServicePrincipal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalCreateServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphServicePrincipal** | [**MicrosoftGraphServicePrincipal**](MicrosoftGraphServicePrincipal.md) | New entity | 

### Return type

[**MicrosoftGraphServicePrincipal**](MicrosoftGraphServicePrincipal.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalDeleteServicePrincipal

> ServicePrincipalsServicePrincipalDeleteServicePrincipal(ctx, servicePrincipalId).IfMatch(ifMatch).Execute()

Delete entity from servicePrincipals

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalDeleteServicePrincipal(context.Background(), servicePrincipalId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalDeleteServicePrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalDeleteServicePrincipalRequest struct via the builder pattern


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


## ServicePrincipalsServicePrincipalGetServicePrincipal

> MicrosoftGraphServicePrincipal ServicePrincipalsServicePrincipalGetServicePrincipal(ctx, servicePrincipalId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()

Get entity from servicePrincipals by key

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalGetServicePrincipal(context.Background(), servicePrincipalId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalGetServicePrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalGetServicePrincipal`: MicrosoftGraphServicePrincipal
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalGetServicePrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalGetServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphServicePrincipal**](MicrosoftGraphServicePrincipal.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalListServicePrincipal

> CollectionOfServicePrincipal ServicePrincipalsServicePrincipalListServicePrincipal(ctx).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from servicePrincipals

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
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
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
    resp, r, err := api_client.ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalListServicePrincipal(context.Background()).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalListServicePrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsServicePrincipalListServicePrincipal`: CollectionOfServicePrincipal
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalListServicePrincipal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalListServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfServicePrincipal**](CollectionOfServicePrincipal.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServicePrincipalsServicePrincipalUpdateServicePrincipal

> ServicePrincipalsServicePrincipalUpdateServicePrincipal(ctx, servicePrincipalId).MicrosoftGraphServicePrincipal(microsoftGraphServicePrincipal).Execute()

Update entity in servicePrincipals

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
    servicePrincipalId := "servicePrincipalId_example" // string | key: id of servicePrincipal
    microsoftGraphServicePrincipal := *openapiclient.NewMicrosoftGraphServicePrincipal() // MicrosoftGraphServicePrincipal | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalUpdateServicePrincipal(context.Background(), servicePrincipalId).MicrosoftGraphServicePrincipal(microsoftGraphServicePrincipal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsServicePrincipalApi.ServicePrincipalsServicePrincipalUpdateServicePrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicePrincipalId** | **string** | key: id of servicePrincipal | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsServicePrincipalUpdateServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphServicePrincipal** | [**MicrosoftGraphServicePrincipal**](MicrosoftGraphServicePrincipal.md) | New property values | 

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

