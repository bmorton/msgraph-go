# \PoliciesAuthenticationMethodsPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesDeleteAuthenticationMethodsPolicy**](PoliciesAuthenticationMethodsPolicyApi.md#PoliciesDeleteAuthenticationMethodsPolicy) | **Delete** /policies/authenticationMethodsPolicy | Delete navigation property authenticationMethodsPolicy for policies
[**PoliciesGetAuthenticationMethodsPolicy**](PoliciesAuthenticationMethodsPolicyApi.md#PoliciesGetAuthenticationMethodsPolicy) | **Get** /policies/authenticationMethodsPolicy | Get authenticationMethodsPolicy from policies
[**PoliciesUpdateAuthenticationMethodsPolicy**](PoliciesAuthenticationMethodsPolicyApi.md#PoliciesUpdateAuthenticationMethodsPolicy) | **Patch** /policies/authenticationMethodsPolicy | Update the navigation property authenticationMethodsPolicy in policies



## PoliciesDeleteAuthenticationMethodsPolicy

> PoliciesDeleteAuthenticationMethodsPolicy(ctx).IfMatch(ifMatch).Execute()

Delete navigation property authenticationMethodsPolicy for policies



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesAuthenticationMethodsPolicyApi.PoliciesDeleteAuthenticationMethodsPolicy(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthenticationMethodsPolicyApi.PoliciesDeleteAuthenticationMethodsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteAuthenticationMethodsPolicyRequest struct via the builder pattern


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


## PoliciesGetAuthenticationMethodsPolicy

> MicrosoftGraphAuthenticationMethodsPolicy PoliciesGetAuthenticationMethodsPolicy(ctx).Select_(select_).Expand(expand).Execute()

Get authenticationMethodsPolicy from policies



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesAuthenticationMethodsPolicyApi.PoliciesGetAuthenticationMethodsPolicy(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthenticationMethodsPolicyApi.PoliciesGetAuthenticationMethodsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetAuthenticationMethodsPolicy`: MicrosoftGraphAuthenticationMethodsPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAuthenticationMethodsPolicyApi.PoliciesGetAuthenticationMethodsPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetAuthenticationMethodsPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAuthenticationMethodsPolicy**](MicrosoftGraphAuthenticationMethodsPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateAuthenticationMethodsPolicy

> PoliciesUpdateAuthenticationMethodsPolicy(ctx).MicrosoftGraphAuthenticationMethodsPolicy(microsoftGraphAuthenticationMethodsPolicy).Execute()

Update the navigation property authenticationMethodsPolicy in policies



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
    microsoftGraphAuthenticationMethodsPolicy := *openapiclient.NewMicrosoftGraphAuthenticationMethodsPolicy() // MicrosoftGraphAuthenticationMethodsPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesAuthenticationMethodsPolicyApi.PoliciesUpdateAuthenticationMethodsPolicy(context.Background()).MicrosoftGraphAuthenticationMethodsPolicy(microsoftGraphAuthenticationMethodsPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthenticationMethodsPolicyApi.PoliciesUpdateAuthenticationMethodsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateAuthenticationMethodsPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAuthenticationMethodsPolicy** | [**MicrosoftGraphAuthenticationMethodsPolicy**](MicrosoftGraphAuthenticationMethodsPolicy.md) | New navigation property values | 

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

