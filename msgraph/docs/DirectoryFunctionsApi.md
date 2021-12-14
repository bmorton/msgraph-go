# \DirectoryFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryAdministrativeUnitsDelta**](DirectoryFunctionsApi.md#DirectoryAdministrativeUnitsDelta) | **Get** /directory/administrativeUnits/microsoft.graph.delta() | Invoke function delta



## DirectoryAdministrativeUnitsDelta

> []*AnyOfmicrosoftGraphAdministrativeUnit DirectoryAdministrativeUnitsDelta(ctx).Execute()

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
    resp, r, err := api_client.DirectoryFunctionsApi.DirectoryAdministrativeUnitsDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryFunctionsApi.DirectoryAdministrativeUnitsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryAdministrativeUnitsDelta`: []*AnyOfmicrosoftGraphAdministrativeUnit
    fmt.Fprintf(os.Stdout, "Response from `DirectoryFunctionsApi.DirectoryAdministrativeUnitsDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryAdministrativeUnitsDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphAdministrativeUnit**](anyOf&lt;microsoft.graph.administrativeUnit&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

