# \ServicePrincipalsFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServicePrincipalsDelta**](ServicePrincipalsFunctionsApi.md#ServicePrincipalsDelta) | **Get** /servicePrincipals/microsoft.graph.delta() | Invoke function delta



## ServicePrincipalsDelta

> []*AnyOfmicrosoftGraphServicePrincipal ServicePrincipalsDelta(ctx).Execute()

Invoke function delta

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
    resp, r, err := api_client.ServicePrincipalsFunctionsApi.ServicePrincipalsDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicePrincipalsFunctionsApi.ServicePrincipalsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServicePrincipalsDelta`: []*AnyOfmicrosoftGraphServicePrincipal
    fmt.Fprintf(os.Stdout, "Response from `ServicePrincipalsFunctionsApi.ServicePrincipalsDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServicePrincipalsDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphServicePrincipal**](anyOf&lt;microsoft.graph.servicePrincipal&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

