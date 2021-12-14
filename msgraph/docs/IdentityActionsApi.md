# \IdentityActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate**](IdentityActionsApi.md#IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate) | **Post** /identity/apiConnectors/{identityApiConnector-id}/microsoft.graph.uploadClientCertificate | Invoke action uploadClientCertificate
[**IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder**](IdentityActionsApi.md#IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder) | **Post** /identity/b2xUserFlows/{b2xIdentityUserFlow-id}/userAttributeAssignments/microsoft.graph.setOrder | Invoke action setOrder



## IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate

> AnyOfmicrosoftGraphIdentityApiConnector IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate(ctx, identityApiConnectorId).InlineObject356(inlineObject356).Execute()

Invoke action uploadClientCertificate

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
    identityApiConnectorId := "identityApiConnectorId_example" // string | key: id of identityApiConnector
    inlineObject356 := *openapiclient.NewInlineObject356() // InlineObject356 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityActionsApi.IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate(context.Background(), identityApiConnectorId).InlineObject356(inlineObject356).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityActionsApi.IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate`: AnyOfmicrosoftGraphIdentityApiConnector
    fmt.Fprintf(os.Stdout, "Response from `IdentityActionsApi.IdentityApiConnectorsIdentityApiConnectorUploadClientCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityApiConnectorId** | **string** | key: id of identityApiConnector | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityApiConnectorsIdentityApiConnectorUploadClientCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject356** | [**InlineObject356**](InlineObject356.md) |  | 

### Return type

[**AnyOfmicrosoftGraphIdentityApiConnector**](anyOf&lt;microsoft.graph.identityApiConnector&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder

> IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder(ctx, b2xIdentityUserFlowId).InlineObject357(inlineObject357).Execute()

Invoke action setOrder

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
    inlineObject357 := *openapiclient.NewInlineObject357() // InlineObject357 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityActionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder(context.Background(), b2xIdentityUserFlowId).InlineObject357(inlineObject357).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityActionsApi.IdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**b2xIdentityUserFlowId** | **string** | key: id of b2xIdentityUserFlow | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityB2xUserFlowsB2xIdentityUserFlowUserAttributeAssignmentsSetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject357** | [**InlineObject357**](InlineObject357.md) |  | 

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

