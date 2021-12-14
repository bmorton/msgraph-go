# \ApplicationsTokenLifetimePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsCreateRefTokenLifetimePolicies**](ApplicationsTokenLifetimePolicyApi.md#ApplicationsCreateRefTokenLifetimePolicies) | **Post** /applications/{application-id}/tokenLifetimePolicies/$ref | Create new navigation property ref to tokenLifetimePolicies for applications
[**ApplicationsListRefTokenLifetimePolicies**](ApplicationsTokenLifetimePolicyApi.md#ApplicationsListRefTokenLifetimePolicies) | **Get** /applications/{application-id}/tokenLifetimePolicies/$ref | Get ref of tokenLifetimePolicies from applications
[**ApplicationsListTokenLifetimePolicies**](ApplicationsTokenLifetimePolicyApi.md#ApplicationsListTokenLifetimePolicies) | **Get** /applications/{application-id}/tokenLifetimePolicies | Get tokenLifetimePolicies from applications



## ApplicationsCreateRefTokenLifetimePolicies

> map[string]interface{} ApplicationsCreateRefTokenLifetimePolicies(ctx, applicationId).RequestBody(requestBody).Execute()

Create new navigation property ref to tokenLifetimePolicies for applications



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
    applicationId := "applicationId_example" // string | key: id of application
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsTokenLifetimePolicyApi.ApplicationsCreateRefTokenLifetimePolicies(context.Background(), applicationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTokenLifetimePolicyApi.ApplicationsCreateRefTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsCreateRefTokenLifetimePolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTokenLifetimePolicyApi.ApplicationsCreateRefTokenLifetimePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsCreateRefTokenLifetimePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsListRefTokenLifetimePolicies

> CollectionOfLinksOfTokenLifetimePolicy ApplicationsListRefTokenLifetimePolicies(ctx, applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of tokenLifetimePolicies from applications



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
    applicationId := "applicationId_example" // string | key: id of application
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsTokenLifetimePolicyApi.ApplicationsListRefTokenLifetimePolicies(context.Background(), applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTokenLifetimePolicyApi.ApplicationsListRefTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsListRefTokenLifetimePolicies`: CollectionOfLinksOfTokenLifetimePolicy
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTokenLifetimePolicyApi.ApplicationsListRefTokenLifetimePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsListRefTokenLifetimePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfTokenLifetimePolicy**](CollectionOfLinksOfTokenLifetimePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsListTokenLifetimePolicies

> CollectionOfTokenLifetimePolicy ApplicationsListTokenLifetimePolicies(ctx, applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tokenLifetimePolicies from applications



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
    applicationId := "applicationId_example" // string | key: id of application
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
    resp, r, err := api_client.ApplicationsTokenLifetimePolicyApi.ApplicationsListTokenLifetimePolicies(context.Background(), applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsTokenLifetimePolicyApi.ApplicationsListTokenLifetimePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsListTokenLifetimePolicies`: CollectionOfTokenLifetimePolicy
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsTokenLifetimePolicyApi.ApplicationsListTokenLifetimePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsListTokenLifetimePoliciesRequest struct via the builder pattern


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

