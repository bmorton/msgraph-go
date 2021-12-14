# \GroupLifecyclePoliciesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup**](GroupLifecyclePoliciesActionsApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup) | **Post** /groupLifecyclePolicies/{groupLifecyclePolicy-id}/microsoft.graph.addGroup | Invoke action addGroup
[**GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup**](GroupLifecyclePoliciesActionsApi.md#GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup) | **Post** /groupLifecyclePolicies/{groupLifecyclePolicy-id}/microsoft.graph.removeGroup | Invoke action removeGroup



## GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup

> bool GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup(ctx, groupLifecyclePolicyId).InlineObject147(inlineObject147).Execute()

Invoke action addGroup

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
    inlineObject147 := *openapiclient.NewInlineObject147() // InlineObject147 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupLifecyclePoliciesActionsApi.GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup(context.Background(), groupLifecyclePolicyId).InlineObject147(inlineObject147).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupLifecyclePoliciesActionsApi.GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupLifecyclePoliciesActionsApi.GroupLifecyclePoliciesGroupLifecyclePolicyAddGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupLifecyclePoliciesGroupLifecyclePolicyAddGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject147** | [**InlineObject147**](InlineObject147.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup

> bool GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup(ctx, groupLifecyclePolicyId).InlineObject148(inlineObject148).Execute()

Invoke action removeGroup

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
    inlineObject148 := *openapiclient.NewInlineObject148() // InlineObject148 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupLifecyclePoliciesActionsApi.GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup(context.Background(), groupLifecyclePolicyId).InlineObject148(inlineObject148).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupLifecyclePoliciesActionsApi.GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup`: bool
    fmt.Fprintf(os.Stdout, "Response from `GroupLifecyclePoliciesActionsApi.GroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupLifecyclePolicyId** | **string** | key: id of groupLifecyclePolicy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupLifecyclePoliciesGroupLifecyclePolicyRemoveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject148** | [**InlineObject148**](InlineObject148.md) |  | 

### Return type

**bool**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

