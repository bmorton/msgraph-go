# \ApplicationsHomeRealmDiscoveryPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsCreateRefHomeRealmDiscoveryPolicies**](ApplicationsHomeRealmDiscoveryPolicyApi.md#ApplicationsCreateRefHomeRealmDiscoveryPolicies) | **Post** /applications/{application-id}/homeRealmDiscoveryPolicies/$ref | Create new navigation property ref to homeRealmDiscoveryPolicies for applications
[**ApplicationsListHomeRealmDiscoveryPolicies**](ApplicationsHomeRealmDiscoveryPolicyApi.md#ApplicationsListHomeRealmDiscoveryPolicies) | **Get** /applications/{application-id}/homeRealmDiscoveryPolicies | Get homeRealmDiscoveryPolicies from applications
[**ApplicationsListRefHomeRealmDiscoveryPolicies**](ApplicationsHomeRealmDiscoveryPolicyApi.md#ApplicationsListRefHomeRealmDiscoveryPolicies) | **Get** /applications/{application-id}/homeRealmDiscoveryPolicies/$ref | Get ref of homeRealmDiscoveryPolicies from applications



## ApplicationsCreateRefHomeRealmDiscoveryPolicies

> map[string]interface{} ApplicationsCreateRefHomeRealmDiscoveryPolicies(ctx, applicationId).RequestBody(requestBody).Execute()

Create new navigation property ref to homeRealmDiscoveryPolicies for applications

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
    resp, r, err := api_client.ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsCreateRefHomeRealmDiscoveryPolicies(context.Background(), applicationId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsCreateRefHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsCreateRefHomeRealmDiscoveryPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsCreateRefHomeRealmDiscoveryPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsCreateRefHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


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


## ApplicationsListHomeRealmDiscoveryPolicies

> CollectionOfHomeRealmDiscoveryPolicy ApplicationsListHomeRealmDiscoveryPolicies(ctx, applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get homeRealmDiscoveryPolicies from applications

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
    resp, r, err := api_client.ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsListHomeRealmDiscoveryPolicies(context.Background(), applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsListHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsListHomeRealmDiscoveryPolicies`: CollectionOfHomeRealmDiscoveryPolicy
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsListHomeRealmDiscoveryPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsListHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


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

[**CollectionOfHomeRealmDiscoveryPolicy**](CollectionOfHomeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationsListRefHomeRealmDiscoveryPolicies

> CollectionOfLinksOfHomeRealmDiscoveryPolicy ApplicationsListRefHomeRealmDiscoveryPolicies(ctx, applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of homeRealmDiscoveryPolicies from applications

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
    resp, r, err := api_client.ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsListRefHomeRealmDiscoveryPolicies(context.Background(), applicationId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsListRefHomeRealmDiscoveryPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsListRefHomeRealmDiscoveryPolicies`: CollectionOfLinksOfHomeRealmDiscoveryPolicy
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsHomeRealmDiscoveryPolicyApi.ApplicationsListRefHomeRealmDiscoveryPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | key: id of application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsListRefHomeRealmDiscoveryPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfHomeRealmDiscoveryPolicy**](CollectionOfLinksOfHomeRealmDiscoveryPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

