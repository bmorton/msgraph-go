# \PoliciesAuthorizationPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesDeleteAuthorizationPolicy**](PoliciesAuthorizationPolicyApi.md#PoliciesDeleteAuthorizationPolicy) | **Delete** /policies/authorizationPolicy | Delete navigation property authorizationPolicy for policies
[**PoliciesGetAuthorizationPolicy**](PoliciesAuthorizationPolicyApi.md#PoliciesGetAuthorizationPolicy) | **Get** /policies/authorizationPolicy | Get authorizationPolicy from policies
[**PoliciesUpdateAuthorizationPolicy**](PoliciesAuthorizationPolicyApi.md#PoliciesUpdateAuthorizationPolicy) | **Patch** /policies/authorizationPolicy | Update the navigation property authorizationPolicy in policies



## PoliciesDeleteAuthorizationPolicy

> PoliciesDeleteAuthorizationPolicy(ctx).IfMatch(ifMatch).Execute()

Delete navigation property authorizationPolicy for policies



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
    resp, r, err := api_client.PoliciesAuthorizationPolicyApi.PoliciesDeleteAuthorizationPolicy(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthorizationPolicyApi.PoliciesDeleteAuthorizationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteAuthorizationPolicyRequest struct via the builder pattern


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


## PoliciesGetAuthorizationPolicy

> MicrosoftGraphAuthorizationPolicy PoliciesGetAuthorizationPolicy(ctx).Select_(select_).Expand(expand).Execute()

Get authorizationPolicy from policies



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
    resp, r, err := api_client.PoliciesAuthorizationPolicyApi.PoliciesGetAuthorizationPolicy(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthorizationPolicyApi.PoliciesGetAuthorizationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetAuthorizationPolicy`: MicrosoftGraphAuthorizationPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAuthorizationPolicyApi.PoliciesGetAuthorizationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetAuthorizationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAuthorizationPolicy**](MicrosoftGraphAuthorizationPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateAuthorizationPolicy

> PoliciesUpdateAuthorizationPolicy(ctx).MicrosoftGraphAuthorizationPolicy(microsoftGraphAuthorizationPolicy).Execute()

Update the navigation property authorizationPolicy in policies



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
    microsoftGraphAuthorizationPolicy := *openapiclient.NewMicrosoftGraphAuthorizationPolicy() // MicrosoftGraphAuthorizationPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesAuthorizationPolicyApi.PoliciesUpdateAuthorizationPolicy(context.Background()).MicrosoftGraphAuthorizationPolicy(microsoftGraphAuthorizationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAuthorizationPolicyApi.PoliciesUpdateAuthorizationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateAuthorizationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAuthorizationPolicy** | [**MicrosoftGraphAuthorizationPolicy**](MicrosoftGraphAuthorizationPolicy.md) | New navigation property values | 

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

