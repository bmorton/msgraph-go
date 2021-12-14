# \Oauth2PermissionGrantsFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2PermissionGrantsDelta**](Oauth2PermissionGrantsFunctionsApi.md#Oauth2PermissionGrantsDelta) | **Get** /oauth2PermissionGrants/microsoft.graph.delta() | Invoke function delta



## Oauth2PermissionGrantsDelta

> []*AnyOfmicrosoftGraphOAuth2PermissionGrant Oauth2PermissionGrantsDelta(ctx).Execute()

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
    resp, r, err := api_client.Oauth2PermissionGrantsFunctionsApi.Oauth2PermissionGrantsDelta(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2PermissionGrantsFunctionsApi.Oauth2PermissionGrantsDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2PermissionGrantsDelta`: []*AnyOfmicrosoftGraphOAuth2PermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `Oauth2PermissionGrantsFunctionsApi.Oauth2PermissionGrantsDelta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOauth2PermissionGrantsDeltaRequest struct via the builder pattern


### Return type

[**[]*AnyOfmicrosoftGraphOAuth2PermissionGrant**](anyOf&lt;microsoft.graph.oAuth2PermissionGrant&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

