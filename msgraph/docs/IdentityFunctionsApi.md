# \IdentityFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder**](IdentityFunctionsApi.md#IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder) | **Get** /identity/b2xUserFlows/{b2xIdentityUserFlow-id}/userAttributeAssignments/microsoft.graph.getOrder() | Invoke function getOrder
[**IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes**](IdentityFunctionsApi.md#IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes) | **Get** /identity/b2xUserFlows/{b2xIdentityUserFlow-id}/userFlowIdentityProviders/microsoft.graph.availableProviderTypes() | Invoke function availableProviderTypes
[**IdentityIdentityProvidersAvailableProviderTypes**](IdentityFunctionsApi.md#IdentityIdentityProvidersAvailableProviderTypes) | **Get** /identity/identityProviders/microsoft.graph.availableProviderTypes() | Invoke function availableProviderTypes



## IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder

> AnyOfmicrosoftGraphAssignmentOrder IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder(ctx, b2xIdentityUserFlowId).Execute()

Invoke function getOrder

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
    b2xIdentityUserFlowId := "b2xIdentityUserFlowId_example" // string | key: id of b2xIdentityUserFlow

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityFunctionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder(context.Background(), b2xIdentityUserFlowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityFunctionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder`: AnyOfmicrosoftGraphAssignmentOrder
    fmt.Fprintf(os.Stdout, "Response from `IdentityFunctionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**b2xIdentityUserFlowId** | **string** | key: id of b2xIdentityUserFlow | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnyOfmicrosoftGraphAssignmentOrder**](anyOf&lt;microsoft.graph.assignmentOrder&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes

> []*string IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes(ctx, b2xIdentityUserFlowId).Execute()

Invoke function availableProviderTypes

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
    b2xIdentityUserFlowId := "b2xIdentityUserFlowId_example" // string | key: id of b2xIdentityUserFlow

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityFunctionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes(context.Background(), b2xIdentityUserFlowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityFunctionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes`: []*string
    fmt.Fprintf(os.Stdout, "Response from `IdentityFunctionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**b2xIdentityUserFlowId** | **string** | key: id of b2xIdentityUserFlow | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityB2xUserFlowsB2xIdentityUserFlowUserFlowIdentityProvidersAvailableProviderTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]*string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityIdentityProvidersAvailableProviderTypes

> []*string IdentityIdentityProvidersAvailableProviderTypes(ctx).Execute()

Invoke function availableProviderTypes

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityFunctionsApi.IdentityIdentityProvidersAvailableProviderTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityFunctionsApi.IdentityIdentityProvidersAvailableProviderTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityIdentityProvidersAvailableProviderTypes`: []*string
    fmt.Fprintf(os.Stdout, "Response from `IdentityFunctionsApi.IdentityIdentityProvidersAvailableProviderTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityIdentityProvidersAvailableProviderTypesRequest struct via the builder pattern


### Return type

**[]*string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

