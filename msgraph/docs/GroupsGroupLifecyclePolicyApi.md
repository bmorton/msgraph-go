# \GroupsGroupLifecyclePolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsCreateGroupLifecyclePolicies) | **Post** /groups/{group-id}/groupLifecyclePolicies | Create new navigation property to groupLifecyclePolicies for groups
[**GroupsDeleteGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsDeleteGroupLifecyclePolicies) | **Delete** /groups/{group-id}/groupLifecyclePolicies/{groupLifecyclePolicy-id} | Delete navigation property groupLifecyclePolicies for groups
[**GroupsGetGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsGetGroupLifecyclePolicies) | **Get** /groups/{group-id}/groupLifecyclePolicies/{groupLifecyclePolicy-id} | Get groupLifecyclePolicies from groups
[**GroupsListGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsListGroupLifecyclePolicies) | **Get** /groups/{group-id}/groupLifecyclePolicies | Get groupLifecyclePolicies from groups
[**GroupsUpdateGroupLifecyclePolicies**](GroupsGroupLifecyclePolicyApi.md#GroupsUpdateGroupLifecyclePolicies) | **Patch** /groups/{group-id}/groupLifecyclePolicies/{groupLifecyclePolicy-id} | Update the navigation property groupLifecyclePolicies in groups



## GroupsCreateGroupLifecyclePolicies

> MicrosoftGraphGroupLifecyclePolicy GroupsCreateGroupLifecyclePolicies(ctx, groupId).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()

Create new navigation property to groupLifecyclePolicies for groups



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
    groupId := "groupId_example" // string | key: id of group
    microsoftGraphGroupLifecyclePolicy := *openapiclient.NewMicrosoftGraphGroupLifecyclePolicy() // MicrosoftGraphGroupLifecyclePolicy | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupLifecyclePolicyApi.GroupsCreateGroupLifecyclePolicies(context.Background(), groupId).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupLifecyclePolicyApi.GroupsCreateGroupLifecyclePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateGroupLifecyclePolicies`: MicrosoftGraphGroupLifecyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `GroupsGroupLifecyclePolicyApi.GroupsCreateGroupLifecyclePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateGroupLifecyclePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md) | New navigation property | 

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


## GroupsDeleteGroupLifecyclePolicies

> GroupsDeleteGroupLifecyclePolicies(ctx, groupId, groupLifecyclePolicyId).IfMatch(ifMatch).Execute()

Delete navigation property groupLifecyclePolicies for groups



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
    groupId := "groupId_example" // string | key: id of group
    groupLifecyclePolicyId := "groupLifecyclePolicyId_example" // string | key: id of groupLifecyclePolicy
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupLifecyclePolicyApi.GroupsDeleteGroupLifecyclePolicies(context.Background(), groupId, groupLifecyclePolicyId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupLifecyclePolicyApi.GroupsDeleteGroupLifecyclePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteGroupLifecyclePoliciesRequest struct via the builder pattern


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


## GroupsGetGroupLifecyclePolicies

> MicrosoftGraphGroupLifecyclePolicy GroupsGetGroupLifecyclePolicies(ctx, groupId, groupLifecyclePolicyId).Select_(select_).Expand(expand).Execute()

Get groupLifecyclePolicies from groups



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
    groupId := "groupId_example" // string | key: id of group
    groupLifecyclePolicyId := "groupLifecyclePolicyId_example" // string | key: id of groupLifecyclePolicy
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupLifecyclePolicyApi.GroupsGetGroupLifecyclePolicies(context.Background(), groupId, groupLifecyclePolicyId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupLifecyclePolicyApi.GroupsGetGroupLifecyclePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetGroupLifecyclePolicies`: MicrosoftGraphGroupLifecyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `GroupsGroupLifecyclePolicyApi.GroupsGetGroupLifecyclePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetGroupLifecyclePoliciesRequest struct via the builder pattern


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


## GroupsListGroupLifecyclePolicies

> CollectionOfGroupLifecyclePolicy GroupsListGroupLifecyclePolicies(ctx, groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get groupLifecyclePolicies from groups



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
    groupId := "groupId_example" // string | key: id of group
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
    resp, r, err := api_client.GroupsGroupLifecyclePolicyApi.GroupsListGroupLifecyclePolicies(context.Background(), groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupLifecyclePolicyApi.GroupsListGroupLifecyclePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsListGroupLifecyclePolicies`: CollectionOfGroupLifecyclePolicy
    fmt.Fprintf(os.Stdout, "Response from `GroupsGroupLifecyclePolicyApi.GroupsListGroupLifecyclePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListGroupLifecyclePoliciesRequest struct via the builder pattern


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


## GroupsUpdateGroupLifecyclePolicies

> GroupsUpdateGroupLifecyclePolicies(ctx, groupId, groupLifecyclePolicyId).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()

Update the navigation property groupLifecyclePolicies in groups



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
    groupId := "groupId_example" // string | key: id of group
    groupLifecyclePolicyId := "groupLifecyclePolicyId_example" // string | key: id of groupLifecyclePolicy
    microsoftGraphGroupLifecyclePolicy := *openapiclient.NewMicrosoftGraphGroupLifecyclePolicy() // MicrosoftGraphGroupLifecyclePolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupLifecyclePolicyApi.GroupsUpdateGroupLifecyclePolicies(context.Background(), groupId, groupLifecyclePolicyId).MicrosoftGraphGroupLifecyclePolicy(microsoftGraphGroupLifecyclePolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupLifecyclePolicyApi.GroupsUpdateGroupLifecyclePolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateGroupLifecyclePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphGroupLifecyclePolicy** | [**MicrosoftGraphGroupLifecyclePolicy**](MicrosoftGraphGroupLifecyclePolicy.md) | New navigation property values | 

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

