# \IdentityProvidersFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityProvidersAvailableProviderTypes**](IdentityProvidersFunctionsApi.md#IdentityProvidersAvailableProviderTypes) | **Get** /identityProviders/microsoft.graph.availableProviderTypes() | Invoke function availableProviderTypes



## IdentityProvidersAvailableProviderTypes

> []*string IdentityProvidersAvailableProviderTypes(ctx).Execute()

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
    resp, r, err := api_client.IdentityProvidersFunctionsApi.IdentityProvidersAvailableProviderTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersFunctionsApi.IdentityProvidersAvailableProviderTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityProvidersAvailableProviderTypes`: []*string
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersFunctionsApi.IdentityProvidersAvailableProviderTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityProvidersAvailableProviderTypesRequest struct via the builder pattern


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

