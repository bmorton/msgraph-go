# \ApplicationsFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsDelta**](ApplicationsFunctionsApi.md#ApplicationsDelta) | **Get** /applications/microsoft.graph.delta() | Invoke function delta



## ApplicationsDelta

> []*AnyOfmicrosoftGraphApplication ApplicationsDelta(ctx).Execute()

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
    resp, r, err := api_client.ApplicationsFunctionsApi.ApplicationsDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsFunctionsApi.ApplicationsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationsDelta`: []*AnyOfmicrosoftGraphApplication
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsFunctionsApi.ApplicationsDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphApplication**](anyOf&lt;microsoft.graph.application&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

