# \PoliciesAdminConsentRequestPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesDeleteAdminConsentRequestPolicy**](PoliciesAdminConsentRequestPolicyApi.md#PoliciesDeleteAdminConsentRequestPolicy) | **Delete** /policies/adminConsentRequestPolicy | Delete navigation property adminConsentRequestPolicy for policies
[**PoliciesGetAdminConsentRequestPolicy**](PoliciesAdminConsentRequestPolicyApi.md#PoliciesGetAdminConsentRequestPolicy) | **Get** /policies/adminConsentRequestPolicy | Get adminConsentRequestPolicy from policies
[**PoliciesUpdateAdminConsentRequestPolicy**](PoliciesAdminConsentRequestPolicyApi.md#PoliciesUpdateAdminConsentRequestPolicy) | **Patch** /policies/adminConsentRequestPolicy | Update the navigation property adminConsentRequestPolicy in policies



## PoliciesDeleteAdminConsentRequestPolicy

> PoliciesDeleteAdminConsentRequestPolicy(ctx).IfMatch(ifMatch).Execute()

Delete navigation property adminConsentRequestPolicy for policies



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
    resp, r, err := api_client.PoliciesAdminConsentRequestPolicyApi.PoliciesDeleteAdminConsentRequestPolicy(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAdminConsentRequestPolicyApi.PoliciesDeleteAdminConsentRequestPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteAdminConsentRequestPolicyRequest struct via the builder pattern


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


## PoliciesGetAdminConsentRequestPolicy

> MicrosoftGraphAdminConsentRequestPolicy PoliciesGetAdminConsentRequestPolicy(ctx).Select_(select_).Expand(expand).Execute()

Get adminConsentRequestPolicy from policies



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
    resp, r, err := api_client.PoliciesAdminConsentRequestPolicyApi.PoliciesGetAdminConsentRequestPolicy(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAdminConsentRequestPolicyApi.PoliciesGetAdminConsentRequestPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetAdminConsentRequestPolicy`: MicrosoftGraphAdminConsentRequestPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAdminConsentRequestPolicyApi.PoliciesGetAdminConsentRequestPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetAdminConsentRequestPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphAdminConsentRequestPolicy**](MicrosoftGraphAdminConsentRequestPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateAdminConsentRequestPolicy

> PoliciesUpdateAdminConsentRequestPolicy(ctx).MicrosoftGraphAdminConsentRequestPolicy(microsoftGraphAdminConsentRequestPolicy).Execute()

Update the navigation property adminConsentRequestPolicy in policies



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
    microsoftGraphAdminConsentRequestPolicy := *openapiclient.NewMicrosoftGraphAdminConsentRequestPolicy() // MicrosoftGraphAdminConsentRequestPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesAdminConsentRequestPolicyApi.PoliciesUpdateAdminConsentRequestPolicy(context.Background()).MicrosoftGraphAdminConsentRequestPolicy(microsoftGraphAdminConsentRequestPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAdminConsentRequestPolicyApi.PoliciesUpdateAdminConsentRequestPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateAdminConsentRequestPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphAdminConsentRequestPolicy** | [**MicrosoftGraphAdminConsentRequestPolicy**](MicrosoftGraphAdminConsentRequestPolicy.md) | New navigation property values | 

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

