# \ApplicationTemplatesActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationTemplatesApplicationTemplateInstantiate**](ApplicationTemplatesActionsApi.md#ApplicationTemplatesApplicationTemplateInstantiate) | **Post** /applicationTemplates/{applicationTemplate-id}/microsoft.graph.instantiate | Invoke action instantiate



## ApplicationTemplatesApplicationTemplateInstantiate

> AnyOfmicrosoftGraphApplicationServicePrincipal ApplicationTemplatesApplicationTemplateInstantiate(ctx, applicationTemplateId).InlineObject18(inlineObject18).Execute()

Invoke action instantiate

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
    applicationTemplateId := "applicationTemplateId_example" // string | key: id of applicationTemplate
    inlineObject18 := *openapiclient.NewInlineObject18() // InlineObject18 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationTemplatesActionsApi.ApplicationTemplatesApplicationTemplateInstantiate(context.Background(), applicationTemplateId).InlineObject18(inlineObject18).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTemplatesActionsApi.ApplicationTemplatesApplicationTemplateInstantiate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationTemplatesApplicationTemplateInstantiate`: AnyOfmicrosoftGraphApplicationServicePrincipal
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTemplatesActionsApi.ApplicationTemplatesApplicationTemplateInstantiate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationTemplateId** | **string** | key: id of applicationTemplate | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationTemplatesApplicationTemplateInstantiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject18** | [**InlineObject18**](InlineObject18.md) |  | 

### Return type

[**AnyOfmicrosoftGraphApplicationServicePrincipal**](anyOf&lt;microsoft.graph.applicationServicePrincipal&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

