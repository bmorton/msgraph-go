# \PoliciesPolicyRootApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesPolicyRootGetPolicyRoot**](PoliciesPolicyRootApi.md#PoliciesPolicyRootGetPolicyRoot) | **Get** /policies | Get policies
[**PoliciesPolicyRootUpdatePolicyRoot**](PoliciesPolicyRootApi.md#PoliciesPolicyRootUpdatePolicyRoot) | **Patch** /policies | Update policies



## PoliciesPolicyRootGetPolicyRoot

> MicrosoftGraphPolicyRoot PoliciesPolicyRootGetPolicyRoot(ctx).Select_(select_).Expand(expand).Execute()

Get policies

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
    resp, r, err := api_client.PoliciesPolicyRootApi.PoliciesPolicyRootGetPolicyRoot(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPolicyRootApi.PoliciesPolicyRootGetPolicyRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PoliciesPolicyRootGetPolicyRoot`: MicrosoftGraphPolicyRoot
    fmt.Fprintf(os.Stdout, "Response from `PoliciesPolicyRootApi.PoliciesPolicyRootGetPolicyRoot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPolicyRootGetPolicyRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPolicyRoot**](MicrosoftGraphPolicyRoot.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyRootUpdatePolicyRoot

> PoliciesPolicyRootUpdatePolicyRoot(ctx).MicrosoftGraphPolicyRoot(microsoftGraphPolicyRoot).Execute()

Update policies

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
    microsoftGraphPolicyRoot := *openapiclient.NewMicrosoftGraphPolicyRoot() // MicrosoftGraphPolicyRoot | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesPolicyRootApi.PoliciesPolicyRootUpdatePolicyRoot(context.Background()).MicrosoftGraphPolicyRoot(microsoftGraphPolicyRoot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesPolicyRootApi.PoliciesPolicyRootUpdatePolicyRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPoliciesPolicyRootUpdatePolicyRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPolicyRoot** | [**MicrosoftGraphPolicyRoot**](MicrosoftGraphPolicyRoot.md) | New property values | 

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

