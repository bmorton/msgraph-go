# \PoliciesAuthenticationFlowsPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesDeleteAuthenticationFlowsPolicy**](PoliciesAuthenticationFlowsPolicyApi.md#PoliciesDeleteAuthenticationFlowsPolicy) | **Delete** /policies/authenticationFlowsPolicy | Delete navigation property authenticationFlowsPolicy for policies
[**PoliciesGetAuthenticationFlowsPolicy**](PoliciesAuthenticationFlowsPolicyApi.md#PoliciesGetAuthenticationFlowsPolicy) | **Get** /policies/authenticationFlowsPolicy | Get authenticationFlowsPolicy from policies
[**PoliciesUpdateAuthenticationFlowsPolicy**](PoliciesAuthenticationFlowsPolicyApi.md#PoliciesUpdateAuthenticationFlowsPolicy) | **Patch** /policies/authenticationFlowsPolicy | Update the navigation property authenticationFlowsPolicy in policies



## PoliciesDeleteAuthenticationFlowsPolicy

> PoliciesDeleteAuthenticationFlowsPolicy(ctx).IfMatch(ifMatch).Execute()

Delete navigation property authenticationFlowsPolicy for policies



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
    resp, r, err := api_client.PoliciesAuthenticationFlowsPolicyApi.PoliciesDeleteAuthenticationFlowsPolicy(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthenticationFlowsPolicyApi.PoliciesDeleteAuthenticationFlowsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteAuthenticationFlowsPolicyRequest struct via the builder pattern


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


## PoliciesGetAuthenticationFlowsPolicy

> MicrosoftGraphAuthenticationFlowsPolicy PoliciesGetAuthenticationFlowsPolicy(ctx).Select_(select_).Expand(expand).Execute()

Get authenticationFlowsPolicy from policies



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
    resp, r, err := api_client.PoliciesAuthenticationFlowsPolicyApi.PoliciesGetAuthenticationFlowsPolicy(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthenticationFlowsPolicyApi.PoliciesGetAuthenticationFlowsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetAuthenticationFlowsPolicy`: MicrosoftGraphAuthenticationFlowsPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAuthenticationFlowsPolicyApi.PoliciesGetAuthenticationFlowsPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetAuthenticationFlowsPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAuthenticationFlowsPolicy**](MicrosoftGraphAuthenticationFlowsPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateAuthenticationFlowsPolicy

> PoliciesUpdateAuthenticationFlowsPolicy(ctx).MicrosoftGraphAuthenticationFlowsPolicy(microsoftGraphAuthenticationFlowsPolicy).Execute()

Update the navigation property authenticationFlowsPolicy in policies



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
    microsoftGraphAuthenticationFlowsPolicy := *openapiclient.NewMicrosoftGraphAuthenticationFlowsPolicy() // MicrosoftGraphAuthenticationFlowsPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesAuthenticationFlowsPolicyApi.PoliciesUpdateAuthenticationFlowsPolicy(context.Background()).MicrosoftGraphAuthenticationFlowsPolicy(microsoftGraphAuthenticationFlowsPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthenticationFlowsPolicyApi.PoliciesUpdateAuthenticationFlowsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateAuthenticationFlowsPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAuthenticationFlowsPolicy** | [**MicrosoftGraphAuthenticationFlowsPolicy**](MicrosoftGraphAuthenticationFlowsPolicy.md) | New navigation property values | 

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

