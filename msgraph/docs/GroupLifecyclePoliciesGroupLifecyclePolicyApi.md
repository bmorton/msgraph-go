# \GroupLifecyclePoliciesGroupLifecyclePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy) | **Post** /groupLifecyclePolicies | Add new entity to groupLifecyclePolicies
[**GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy) | **Delete** /groupLifecyclePolicies/{groupLifecyclePolicy-id} | Delete entity from groupLifecyclePolicies
[**GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy) | **Get** /groupLifecyclePolicies/{groupLifecyclePolicy-id} | Get entity from groupLifecyclePolicies by key
[**GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy) | **Get** /groupLifecyclePolicies | Get entities from groupLifecyclePolicies
[**GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy**](GroupLifecyclePoliciesGroupLifecyclePolicyApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy) | **Patch** /groupLifecyclePolicies/{groupLifecyclePolicy-id} | Update entity in groupLifecyclePolicies



## GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy

> MicrosoftGraphGroupLifecyclePolicy GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy(ctx).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()

Add new entity to groupLifecyclePolicies

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
    microsoftGraphGroupLifecyclePolicy := *openapiclient.NewMicrosoftGraphGroupLifecyclePolicy() // MicrosoftGraphGroupLifecyclePolicy | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy(context.Background()).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy`: MicrosoftGraphGroupLifecyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupLifecyclePoliciesGroupLifecyclePolicyCreateGroupLifecyclePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md) | New entity | 

### Return type

[**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy

> GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy(ctx, groupLifecyclePolicyId).IfMatch(ifMatch).Execute()

Delete entity from groupLifecyclePolicies

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
    groupLifecyclePolicyId := "groupLifecyclePolicyId_example" // string | key: id of groupLifecyclePolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy(context.Background(), groupLifecyclePolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupLifecyclePoliciesGroupLifecyclePolicyDeleteGroupLifecyclePolicyRequest struct via the builder pattern


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


## GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy

> MicrosoftGraphGroupLifecyclePolicy GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy(ctx, groupLifecyclePolicyId).Select_(select_).Expand(expand).Execute()

Get entity from groupLifecyclePolicies by key

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
    groupLifecyclePolicyId := "groupLifecyclePolicyId_example" // string | key: id of groupLifecyclePolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy(context.Background(), groupLifecyclePolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy`: MicrosoftGraphGroupLifecyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupLifecyclePoliciesGroupLifecyclePolicyGetGroupLifecyclePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy

> CollectionOfGroupLifecyclePolicy GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from groupLifecyclePolicies

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
    resp, r, err := api_client.GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy`: CollectionOfGroupLifecyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupLifecyclePoliciesGroupLifecyclePolicyListGroupLifecyclePolicyRequest struct via the builder pattern


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

[**CollectionOfGroupLifecyclePolicy**](CollectionOfGroupLifecyclePolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy

> GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy(ctx, groupLifecyclePolicyId).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()

Update entity in groupLifecyclePolicies

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
    groupLifecyclePolicyId := "groupLifecyclePolicyId_example" // string | key: id of groupLifecyclePolicy
    microsoftGraphGroupLifecyclePolicy := *openapiclient.NewMicrosoftGraphGroupLifecyclePolicy() // MicrosoftGraphGroupLifecyclePolicy | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy(context.Background(), groupLifecyclePolicyId).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupLifecyclePoliciesGroupLifecyclePolicyApi.GroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupLifecyclePoliciesGroupLifecyclePolicyUpdateGroupLifecyclePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md) | New property values | 

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

