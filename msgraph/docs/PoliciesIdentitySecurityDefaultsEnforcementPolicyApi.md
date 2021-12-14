# \PoliciesIdentitySecurityDefaultsEnforcementPolicyApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesDeleteIdentitySecurityDefaultsEnforcementPolicy**](PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.md#PoliciesDeleteIdentitySecurityDefaultsEnforcementPolicy) | **Delete** /policies/identitySecurityDefaultsEnforcementPolicy | Delete navigation property identitySecurityDefaultsEnforcementPolicy for policies
[**PoliciesGetIdentitySecurityDefaultsEnforcementPolicy**](PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.md#PoliciesGetIdentitySecurityDefaultsEnforcementPolicy) | **Get** /policies/identitySecurityDefaultsEnforcementPolicy | Get identitySecurityDefaultsEnforcementPolicy from policies
[**PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy**](PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.md#PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy) | **Patch** /policies/identitySecurityDefaultsEnforcementPolicy | Update the navigation property identitySecurityDefaultsEnforcementPolicy in policies



## PoliciesDeleteIdentitySecurityDefaultsEnforcementPolicy

> PoliciesDeleteIdentitySecurityDefaultsEnforcementPolicy(ctx).IfMatch(ifMatch).Execute()

Delete navigation property identitySecurityDefaultsEnforcementPolicy for policies



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
    resp, r, err := api_client.PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.PoliciesDeleteIdentitySecurityDefaultsEnforcementPolicy(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.PoliciesDeleteIdentitySecurityDefaultsEnforcementPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesDeleteIdentitySecurityDefaultsEnforcementPolicyRequest struct via the builder pattern


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


## PoliciesGetIdentitySecurityDefaultsEnforcementPolicy

> MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy PoliciesGetIdentitySecurityDefaultsEnforcementPolicy(ctx).Select_(select_).Expand(expand).Execute()

Get identitySecurityDefaultsEnforcementPolicy from policies



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
    resp, r, err := api_client.PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.PoliciesGetIdentitySecurityDefaultsEnforcementPolicy(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.PoliciesGetIdentitySecurityDefaultsEnforcementPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesGetIdentitySecurityDefaultsEnforcementPolicy`: MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.PoliciesGetIdentitySecurityDefaultsEnforcementPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesGetIdentitySecurityDefaultsEnforcementPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy**](MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy

> PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy(ctx).MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy(microsoftGraphIdentitySecurityDefaultsEnforcementPolicy).Execute()

Update the navigation property identitySecurityDefaultsEnforcementPolicy in policies



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
    microsoftGraphIdentitySecurityDefaultsEnforcementPolicy := *openapiclient.NewMicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy() // MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy(context.Background()).MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy(microsoftGraphIdentitySecurityDefaultsEnforcementPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesIdentitySecurityDefaultsEnforcementPolicyApi.PoliciesUpdateIdentitySecurityDefaultsEnforcementPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesUpdateIdentitySecurityDefaultsEnforcementPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIdentitySecurityDefaultsEnforcementPolicy** | [**MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy**](MicrosoftGraphIdentitySecurityDefaultsEnforcementPolicy.md) | New navigation property values | 

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

