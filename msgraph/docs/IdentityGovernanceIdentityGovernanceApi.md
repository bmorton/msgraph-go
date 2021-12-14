# \IdentityGovernanceIdentityGovernanceApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityGovernanceIdentityGovernanceGetIdentityGovernance**](IdentityGovernanceIdentityGovernanceApi.md#IdentityGovernanceIdentityGovernanceGetIdentityGovernance) | **Get** /identityGovernance | Get identityGovernance
[**IdentityGovernanceIdentityGovernanceUpdateIdentityGovernance**](IdentityGovernanceIdentityGovernanceApi.md#IdentityGovernanceIdentityGovernanceUpdateIdentityGovernance) | **Patch** /identityGovernance | Update identityGovernance



## IdentityGovernanceIdentityGovernanceGetIdentityGovernance

> MicrosoftGraphIdentityGovernance IdentityGovernanceIdentityGovernanceGetIdentityGovernance(ctx).Select_(select_).Expand(expand).Execute()

Get identityGovernance

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
    resp, r, err := api_client.IdentityGovernanceIdentityGovernanceApi.IdentityGovernanceIdentityGovernanceGetIdentityGovernance(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceIdentityGovernanceApi.IdentityGovernanceIdentityGovernanceGetIdentityGovernance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceIdentityGovernanceGetIdentityGovernance`: MicrosoftGraphIdentityGovernance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceIdentityGovernanceApi.IdentityGovernanceIdentityGovernanceGetIdentityGovernance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceIdentityGovernanceGetIdentityGovernanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphIdentityGovernance**](MicrosoftGraphIdentityGovernance.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceIdentityGovernanceUpdateIdentityGovernance

> IdentityGovernanceIdentityGovernanceUpdateIdentityGovernance(ctx).MicrosoftGraphIdentityGovernance(microsoftGraphIdentityGovernance).Execute()

Update identityGovernance

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
    microsoftGraphIdentityGovernance := *openapiclient.NewMicrosoftGraphIdentityGovernance() // MicrosoftGraphIdentityGovernance | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceIdentityGovernanceApi.IdentityGovernanceIdentityGovernanceUpdateIdentityGovernance(context.Background()).MicrosoftGraphIdentityGovernance(microsoftGraphIdentityGovernance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceIdentityGovernanceApi.IdentityGovernanceIdentityGovernanceUpdateIdentityGovernance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceIdentityGovernanceUpdateIdentityGovernanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphIdentityGovernance** | [**MicrosoftGraphIdentityGovernance**](MicrosoftGraphIdentityGovernance.md) | New property values | 

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

